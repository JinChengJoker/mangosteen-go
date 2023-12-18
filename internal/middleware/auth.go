package middleware

import (
	"context"
	"mangosteen/dal/query"
	"mangosteen/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		at := ctx.GetHeader("Authorization")
		if len(at) < 8 || at[0:7] != "Bearer " {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效的令牌",
			})
			return
		}
		tokenString := at[7:]
		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效的令牌",
			})
			return
		}
		uid := claims.UserID
		qu := query.User
		user, err := qu.WithContext(context.Background()).Where(qu.ID.Eq(uid)).First()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效的令牌",
			})
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
