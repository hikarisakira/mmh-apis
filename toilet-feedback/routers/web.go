package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"xorm.io/xorm/names"

	"git.hcmmh.linux/k861/mmh-apis/toilet-feedback/controllers"
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
	var dsnHC string

	dsnHC = os.Getenv("ORA_CONNECT_HC")
	log.Println("Now connecting to HC database.")

	engineHC, err := xorm.NewEngine("godror", dsnHC)
	if err != nil {
		log.Println("Connecting to HC database failed:", err)
		return
	}
	//日志打印SQL
	engineHC.ShowSQL(true)
	//设置连接池的空闲数大小
	engineHC.SetMaxIdleConns(5)
	//设置最大打开连接数
	engineHC.SetMaxOpenConns(15)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engineHC.SetTableMapper(names.SnakeMapper{})
	// Start service, listening on default port
	w.routing(engineHC)
}

func (w *WebService) routing(db *xorm.Engine) {

	userController := controllers.UserController{DB: db}

	r := gin.Default()
	r.LoadHTMLGlob("dist/*.html")
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.png", "./dist/favicon.png")
	r.StaticFile("/logo_hc.png", "./dist/logo_hc.png")
	r.StaticFile("/logo_cc.png", "./dist/logo_cc.png")
	r.GET("/", HttpWeb)

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

	v1 := r.Group("/ins")
	v1.POST("/ins", userController.MannouInsert)

	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 6003))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	err := r.Run(":6003")
	if err != nil {
		log.Printf("HC database failed to start: %v", err)
	}
}

func HttpWeb(c *gin.Context) {
	data := new(IndexData)
	data.title = "公共區域通報系統"
	data.context = "新竹馬偕紀念醫院/公共區域通報系統"
	c.HTML(http.StatusOK, "index.html", data)
}
