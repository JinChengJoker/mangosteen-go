package router

import (
	"mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

var controllers = []controller.Controller{
	&controller.ValidationCode{},
	&controller.Session{},
}

func Setup(r *gin.Engine) {
	r.GET("/api/v1/ping", controller.Ping)

	for _, c := range controllers {
		c.RegisterRoutes(r)
	}
}
