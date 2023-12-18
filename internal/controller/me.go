package controller

import (
	"mangosteen/dal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Me struct{}

func (m *Me) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/v1/me", m.Get)
}

func (m *Me) Create(ctx *gin.Context) {}

func (m *Me) Delete(ctx *gin.Context) {}

func (m *Me) Update(ctx *gin.Context) {}

func (m *Me) Get(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	me, ok := user.(*model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "请求处理失败，请稍后再试。",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data": me,
		})
	}
}

func (m *Me) List(ctx *gin.Context) {}
