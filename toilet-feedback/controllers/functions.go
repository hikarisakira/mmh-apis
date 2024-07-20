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
	_, err := uc.DB.Exec("INSERT INTO MANNOU (An_Code, An_Name, An_Cout1, An_Cout2, An_Cout3, An_Cout4, An_Cout5, An_Cout6) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		data.An_Code, data.An_Name, data.An_Cout1, data.An_Cout2, data.An_Cout3, data.An_Cout4, data.An_Cout5, data.An_Cout6)

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
