package controller

import (
	"errors"
	"mangosteen/internal/auth"
	"mangosteen/internal/database"
	"mangosteen/internal/database/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSession(ctx *gin.Context) {
	var rBody struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}

	// 获取 request body
	if err := ctx.ShouldBindJSON(&rBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 校对邮箱
	var vCode model.ValidationCode
	result := database.DB.Where("email = ?", rBody.Email).Last(&vCode)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "邮箱未注册",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	// 校对验证码
	if rBody.Code == vCode.Code {
		// 生成 jwt
		jwt, err := auth.NewJWT(1)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "login success",
			"token":   jwt,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码不正确",
		})
	}
}
