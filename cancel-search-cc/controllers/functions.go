package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"xorm.io/xorm"
)

type UserController struct {
	DB *xorm.Engine
}

// @Summary		GetDelRecord
// @Description	輸入病歷號碼來獲取取消掛號紀錄
// @Tags			Search PatientService
// @Produce		x-www-form-urlencoded
// @Param			pno	query	string	true	"病歷號碼"
// @Router			/search/del/{pno} [get]
func (uc *UserController) GetDelRecord(c *gin.Context) {
	pno := c.DefaultQuery("pno", `88001111`)
	log.Println(pno)
	sql := "select pno, operdhm, dept, dr, schdate, schap, room, digseq, udate, xuser from ODRH_DEL where PNO=? order by schdate desc"
	result, err := uc.DB.QueryString(sql, pno)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "獲取錯誤",
			"log":  err,
		})

		return
	}
	if result == nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "沒有獲取到資料或找不到該病患",
			"log":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":   200,
		"status": "紀錄已獲取",
		"msg":    result,
	})

}
