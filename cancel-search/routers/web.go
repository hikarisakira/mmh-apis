package routers

import (
	"fmt"
	"log"
	"os"
	"xorm.io/xorm/names"

	"git.hcmmh.linux/k861/mmh-apis/cancel-search/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"xorm.io/xorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebService struct{}
type IndexData struct {
	title   string
	context string
}

func (w *WebService) Run() {
	var dsn string
	if gin.Mode() == gin.DebugMode {
		dsn = os.Getenv("ORA_CONNECT_TESTING")
		log.Println("Now connecting to testing database.")
	}
	if gin.Mode() == gin.ReleaseMode {
		dsn = os.Getenv("ORA_CONNECT_PRODUCTION")
		log.Println("Now connecting to production database.")
	}
	engine, err := xorm.NewEngine("godror", dsn)
	if err != nil {
		println("Oracle database linking err.")
		return
	}

	//日志打印SQL
	engine.ShowSQL(true)
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(15)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(names.SnakeMapper{})

	w.routing(engine)
}

func (w *WebService) routing(db *xorm.Engine) {

	userController := controllers.UserController{DB: db}

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})

	v1 := r.Group("/search")
	v1.GET("/del/:pno", userController.GetDelRecord)

	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8080))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	err := r.Run(":8080")
	if err != nil {
		log.Printf("fail to start")
	}
}
