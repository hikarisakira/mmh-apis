package controllers

import (
	"git.hcmmh.linux/k861/mmh-apis/toilet-feedback/models"
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type UserController struct {
	DB *xorm.Engine
}

// @Summary		MannouInsert
// @Description	輸入通報資料
// @Tags			Insert
// @Accept			json
// @Produce		x-www-form-urlencoded
// @Param			request	body	models.FeedBackFormat	true	"由上而下值為：問卷代號(請全部大寫)、項目代碼、答案、病歷號碼"
// @Router			/ins/ins [post]
func (uc *UserController) MannouInsert(c *gin.Context) {
	var data []models.FeedBackFormat

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for _, item := range data {
		_, err := uc.DB.Exec("INSERT INTO MANNOU (An_Code, An_Name, An_Cout1, An_Cout2, An_Cout3, An_Cout4, An_Cout5, An_Cout6) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", item.An_Code, item.An_Name, item.An_Cout1, item.An_Cout2, item.An_Cout3, item.An_Cout4, item.An_Cout5, item.An_Cout6)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status": "表單數據已成功插入",
			"data":   data,
		})
	}
}
