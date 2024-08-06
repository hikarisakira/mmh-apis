package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"git.hcmmh.linux/k861/mmh-apis/toilet-feedback/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebService struct{}

var validLocationCodes map[string]bool

func init() {
	var err error
	validLocationCodes, err = ParseLocationCodes("./index.ts")
	if err != nil {
		log.Fatalf("Failed to parse location codes: %v", err)
	}
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

	// CORS middleware
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

	// 設置靜態文件路由
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.png", "./dist/favicon.png")
	r.StaticFile("/logo_hc.png", "./dist/logo_hc.png")
	r.StaticFile("/logo_cc.png", "./dist/logo_cc.png")

	// API 路由
	v1 := r.Group("/ins")
	v1.POST("/ins", userController.MannouInsert)

	// Swagger 路由
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 6003))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// 處理所有其他路徑
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/" || isValidLocationCode(path[1:]) {
			// 根路徑或有效的地點代號，返回 index.html
			c.File("./dist/index.html")
		} else {
			// 其他所有路徑，重定向到根路徑
			c.Redirect(http.StatusFound, "/")
		}
	})

	err := r.Run(":6003")
	if err != nil {
		log.Printf("HC database failed to start: %v", err)
	}
}

// isValidLocationCode 函數用於驗證地點代號是否有效
func isValidLocationCode(code string) bool {
	return validLocationCodes[code]
}
