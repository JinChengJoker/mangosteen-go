package controller

import (
	"fmt"
	"mangosteen/internal/database"
	"mangosteen/internal/database/model"
	"mangosteen/internal/email"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EmailBody struct {
	Email string `json:"email" binding:"required"`
}

func CreateValidationCode(ctx *gin.Context) {
	var json EmailBody

	// 获取 body 数据
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 创建记录
	code := genCode()
	v_code := model.ValidationCode{Email: json.Email, Code: code}
	result := database.DB.Create(&v_code)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	// 发送邮件
	err := email.Send(
		json.Email,
		"山竹记账验证码",
		fmt.Sprintf(`<body><h3>正在登录或注册山竹记账</h3><p>验证码：</p><h3>%s</h3><p>请勿与任何人共享此验证码，谢谢！</p><body>`, code),
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "发送成功",
	})
}

func genCode() string {
	// 设置随机数种子，以确保每次运行生成不同的随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(900000) + 100000
	s := strconv.Itoa(n)

	return s
}
