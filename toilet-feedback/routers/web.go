package routers

import (
	"fmt"
	"git.hcmmh.linux/k861/mmh-apis/toilet-feedback/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"log"
	"net/http"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebService struct{}

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

	// 设置静态文件服务
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.png", "./dist/favicon.png")
	r.StaticFile("/logo_hc.png", "./dist/logo_hc.png")
	r.StaticFile("/logo_cc.png", "./dist/logo_cc.png")

	// 处理根路径和带有地点代号的路径
	r.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "/" || path == "" {
			// 根路径，返回 index.html
			c.File("./dist/index.html")
		} else {
			// 检查是否为有效的地点代号（这里您需要实现自己的验证逻辑）
			if isValidLocationCode(path[1:]) {
				// 有效的地点代号，返回 index.html
				c.File("./dist/index.html")
			} else {
				// 无效的路径，返回 404
				c.String(http.StatusNotFound, "404 page not found")
			}
		}
	})

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

// isValidLocationCode 函数用于验证地点代号是否有效
func isValidLocationCode(code string) bool {
	// 这里您需要实现自己的验证逻辑
	// 例如，检查代号是否在预定义的列表中，或者查询数据库
	// 这里只是一个示例实现
	validCodes := map[string]bool{
		"P3F301M": true,
		"F5F503F": true,
		"TESTING": true,
	}
	return validCodes[code]
}
