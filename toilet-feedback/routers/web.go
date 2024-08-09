package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// 构建 index.ts 文件的完整路径
	indexPath := filepath.Join(currentDir, "routers", "index.ts")

	// 检查文件是否存在
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		log.Fatalf("index.ts file does not exist at path: %s", indexPath)
	}

	validLocationCodes, err = ParseLocationCodes(indexPath)
	if err != nil {
		log.Fatalf("Failed to parse location codes: %v", err)
	}

	log.Printf("Successfully parsed %d location codes from %s", len(validLocationCodes), indexPath)
}
func (w *WebService) Run() {
	var dsnHC string

	dsnHC = os.Getenv("ORA_CONNECT_HC")
	log.Println("Now connecting to HC database.")
	dsnSMS := os.Getenv("ORA_CONNECT_SMS")
	log.Println("Ready to send SMS message.")

	engineHC, err := w.createDBEngine(dsnHC)
	if err != nil {
		log.Println("Connecting to HC database failed:", err)
		return
	}
	engineSMS, err := w.createDBEngine(dsnSMS)
	if err != nil {
		log.Println("Connecting to SMS database failed:", err)
		return
	}

	// Start service, listening on default port
	w.routing(engineHC, engineSMS)
}

func (w *WebService) createDBEngine(dsn string) (*xorm.Engine, error) {

	engine, err := xorm.NewEngine("godror", dsn)
	if err != nil {
		return nil, err
	}
	//日志打印SQL
	engine.ShowSQL(true)
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(15)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(names.SnakeMapper{})
	return engine, nil
}

func (w *WebService) routing(dbHC, dbSMS *xorm.Engine) {
	userController := controllers.UserController{DB: dbHC}
	userControllerSMS := controllers.UserController{DB: dbSMS}

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
	v1.POST("/sms", userControllerSMS.SMSSend)

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
		log.Printf("Server failed to start: %v", err)
	}
}

// isValidLocationCode 函數用於驗證地點代號是否有效
func isValidLocationCode(code string) bool {
	return validLocationCodes[code]
}
