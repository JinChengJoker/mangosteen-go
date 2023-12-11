package controller

import (
	"errors"
	"mangosteen/internal/auth"
	"mangosteen/internal/database"
	"mangosteen/internal/database/model"
	"net/http"
	"time"

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
				"message": "请先发送邮箱验证码",
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
		var user model.User
		// 查找用户
		result := database.DB.Where("email = ?", rBody.Email).First(&user)
		if result.Error != nil {
			// 未找到对应的用户
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// 使用 email 创建用户
				user.Email = rBody.Email
				result := database.DB.Create(&user)
				if result.Error != nil {
					ctx.JSON(500, gin.H{
						"message": result.Error.Error(),
					})
					return
				}
			}
		}

		// 生成 jwt
		jwt, err := auth.NewJWT(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// 标记验证码为已使用
		currentTime := time.Now()
		vCode.UsedAt = &currentTime
		result = database.DB.Save(&vCode)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": result.Error.Error(),
			})
			return
		}

		// 接口返回 jwt
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
