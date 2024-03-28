package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"git.hcmmh.linux/k861/mmh-apis/cancel-search/initializers"
	"git.hcmmh.linux/k861/mmh-apis/cancel-search/routers"

	_ "github.com/godror/godror"

	//swagger
	_ "git.hcmmh.linux/k861/mmh-apis/cancel-search/docs"
)

//	@title			testApi
//	@version		24.02b7f3
//	@description	刪除預約掛號查詢表格
//	@contact.name	李承洋/K861
//	@contact.email	hsakira.k861@mmh.org.tw

//	@license.name	Private/馬偕紀念醫院 員工保密協定

// @host		10.8.41.142:8080
// @schemes	http
func main() {
	initializers.LoadEnv()
	r := routers.WebService{}
	r.Run()
	fmt.Println("System API online.")

	// 測試時請將下列此行註記掉
	gin.SetMode(gin.ReleaseMode)
}
