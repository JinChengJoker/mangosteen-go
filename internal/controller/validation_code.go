package controller

import (
	"context"
	"fmt"
	"mangosteen/dal/model"
	"mangosteen/dal/query"
	"mangosteen/internal/email"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ValidationCode struct{}

func (vc *ValidationCode) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/v1/validation_code", vc.Create)
}

func (vc *ValidationCode) Create(ctx *gin.Context) {
	var rBody struct {
		Email string `json:"email" binding:"required"`
	}

	// 获取 body 数据
	if err := ctx.ShouldBindJSON(&rBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 创建记录
	code := genCode()
	v := model.ValidationCode{Email: rBody.Email, Code: code}
	err := query.ValidationCode.WithContext(context.Background()).Create(&v)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 发送邮件
	err = email.Send(
		rBody.Email,
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

func (vc *ValidationCode) Delete(ctx *gin.Context) {}
func (vc *ValidationCode) Update(ctx *gin.Context) {}
func (vc *ValidationCode) Get(ctx *gin.Context)    {}
func (vc *ValidationCode) List(ctx *gin.Context)   {}

func genCode() string {
	// 设置随机数种子，以确保每次运行生成不同的随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(900000) + 100000
	s := strconv.Itoa(n)

	return s
}
