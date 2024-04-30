package controllers

import (
	"github.com/gin-gonic/gin"
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
	//log.Println(pno)
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

// @Summary		GetPatientInfoViaIdno
// @Description	輸入身分證字號，取得病歷號、姓名、性別、生日、身分證字號
// @Tags			Search
// @Produce		x-www-form-urlencoded
// @Param			idno	query	string	true  "身分證字號"
// @Router			/search/id/{idno} [get]

func (uc *UserController) GetPatientInfoViaIdno(c *gin.Context) {
	idno := c.DefaultQuery("idno", `O123456789`)
	sql := "select pno, name, sex, BIRTH, idno from idp where idno=?"
	result, err := uc.DB.QueryString(sql, idno)
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

// @Summary		GetPatientInfoViaPno
// @Description	輸入病歷號碼，取得病歷號、姓名、性別、生日、身分證字號
// @Tags			Search	Information
// @Produce		x-www-form-urlencoded
// @Param			pno		query	string	true  "病歷號碼"
// @Router			/search/pno/{pno} [get]
func (uc *UserController) GetPatientInfoViaPno(c *gin.Context) {
	pno := c.DefaultQuery("pno", `88001555`)
	sql := "select pno, name, sex, BIRTH, idno from idp where pno=?"
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
