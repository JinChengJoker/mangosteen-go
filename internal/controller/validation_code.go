package controller

import (
	"mangosteen/internal/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailBody struct {
	Email string `json:"email" binding:"required"`
}

func CreateValidationCode(ctx *gin.Context) {
	var json EmailBody

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := email.Send(json.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "发送成功",
	})
}
