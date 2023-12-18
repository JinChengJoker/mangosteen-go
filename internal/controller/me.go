package controller

import (
	"context"
	"mangosteen/dal/query"
	"mangosteen/internal/auth"
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
	at := ctx.GetHeader("Authorization")
	if len(at) < 8 || at[0:7] != "Bearer " {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "无效的令牌",
		})
		return
	}
	tokenString := at[7:]
	claims, err := auth.ParseJWT(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "无效的令牌",
		})
		return
	}
	uid := claims.UserID
	qu := query.User
	user, err := qu.WithContext(context.Background()).Where(qu.ID.Eq(uid)).First()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "无效的令牌",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (m *Me) List(ctx *gin.Context) {}
