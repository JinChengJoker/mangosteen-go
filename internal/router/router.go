package router

import (
	"mangosteen/internal/controller"
	"mangosteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

var noAuthControllers = []controller.Controller{
	&controller.ValidationCode{},
	&controller.Session{},
}
var needAuthControllers = []controller.Controller{
	&controller.Me{},
}

func Setup(r *gin.Engine) {
	r.GET("/ping", controller.Ping)

	rg := r.Group("/auth")
	for _, c := range noAuthControllers {
		c.RegisterRoutes(rg)
	}

	rg = r.Group("/api", middleware.Auth())
	for _, c := range needAuthControllers {
		c.RegisterRoutes(rg)
	}
}
