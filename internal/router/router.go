package router

import (
	"mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/api/v1/ping", controller.Ping)
	r.POST("/api/v1/validation_code", controller.CreateValidationCode)
	r.POST("/api/v1/login", controller.CreateSession)
}
