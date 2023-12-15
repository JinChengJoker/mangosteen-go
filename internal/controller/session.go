package controller

import (
	"context"
	"errors"
	"mangosteen/dal/query"
	"mangosteen/internal/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Session struct{}

func (s *Session) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/login", s.Create)
}

func (s *Session) Create(ctx *gin.Context) {
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
	vc := query.ValidationCode
	vcode, err := vc.WithContext(context.Background()).Where(vc.Email.Eq(rBody.Email)).Last()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "登录失败，无效的邮箱或验证码",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 校对验证码
	if rBody.Code == vcode.Code {
		// 检查验证码是否已被使用
		if vcode.UsedAt != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "登录失败，验证码已失效",
			})
			return
		}

		u := query.User
		// 查找用户
		user, err := u.WithContext(context.Background()).Where(u.Email.Eq(rBody.Email)).First()
		if err != nil {
			// 未找到对应的用户
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 使用 email 创建用户
				user.Email = rBody.Email
				err := u.WithContext(context.Background()).Create(user)
				if err != nil {
					ctx.JSON(500, gin.H{
						"message": err.Error(),
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
		vcode.UsedAt = &currentTime
		err = vc.WithContext(context.Background()).Save(vcode)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
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

func (s *Session) Delete(ctx *gin.Context) {}
func (s *Session) Update(ctx *gin.Context) {}
func (s *Session) Get(ctx *gin.Context)    {}
func (s *Session) List(ctx *gin.Context)   {}
