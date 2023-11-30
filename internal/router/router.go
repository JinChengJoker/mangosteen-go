package router

import (
	"mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/ping", controller.Ping)
}
