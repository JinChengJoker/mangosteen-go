package cmd

import (
	"mangosteen/internal/router"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	router.Setup(r)
	r.Run()
}
