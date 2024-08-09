package controllers

import (
	"log"
	"net/http"

	"git.hcmmh.linux/k861/mmh-apis/toilet-feedback/models"
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type UserController struct {
	DB *xorm.Engine
}

// @Summary		MannouInsert
// @Description 公共區域通報系統，輸入所有須通報事項
// @Tags			Write	Feedback
// @Accept			json
// @Produce		x-www-form-urlencoded
// @Param			request	body	models.FeedBackFormat	true  "回報事項"
// @Router			/ins [post]

func (uc *UserController) MannouInsert(c *gin.Context) {
	var data models.FeedBackFormat

	// 尝试解析 JSON
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	log.Printf("Received data: %+v", data)

	// 执行数据库插入
	_, err := uc.DB.Exec("INSERT INTO MANNOU (PlaceCode, PlaceName, An_Cout1, An_Cout2, An_Cout3, An_Cout4, Bn_Cout1, Bn_Cout2, Bn_Cout3, Bn_Cout4, Bn_Cout5, Bn_Cout6, Bn_Cout7) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.PlaceCode, data.PlaceName, data.An_Cout1, data.An_Cout2, data.An_Cout3, data.An_Cout4, data.Bn_Cout1, data.Bn_Cout2, data.Bn_Cout3, data.Bn_Cout4, data.Bn_Cout5, data.Bn_Cout6, data.Bn_Cout7)

	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	// 成功响应
	c.JSON(http.StatusOK, gin.H{
		"status": "表單數據已成功插入",
		"data":   data,
	})
}

// @Summary		SMSSend
// @Description 傳送簡訊
// @Tags			Write	SMS
// @Accept			json
// @Produce		x-www-form-urlencoded
// @Param			request	body	models.SMSData	true "電話號碼、簡訊內容"
// @Router			/sms [post]
func (uc *UserController) SMSSend(c *gin.Context) {
	var data models.SMSData

	// 尝试解析 JSON
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	log.Printf("Received data: %+v", data)

	// 执行数据库插入
	_, err := uc.DB.Exec("INSERT INTO SMSMSG (Kind, Phsno, Message, Status, OPER, Udate) VALUES ('公共通報', ?, ?, 'M', 7330, SYSDATE)", data.Phsno, data.Message)

	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	// 成功响应
	c.JSON(http.StatusOK, gin.H{
		"status": "表單數據已成功插入",
		"data":   data,
	})
}
