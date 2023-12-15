package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(rg *gin.RouterGroup)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
}
