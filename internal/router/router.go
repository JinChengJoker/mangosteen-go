package router

import (
	"mangosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

var controllers = []controller.Controller{
	&controller.ValidationCode{},
	&controller.Session{},
	&controller.Me{},
}

func Setup(r *gin.Engine) {
	rg := r.Group("/api")

	rg.GET("/v1/ping", controller.Ping)

	for _, c := range controllers {
		c.RegisterRoutes(rg)
	}
}
