package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	database.Connect()
	defer database.Close(database.DB)

	r := gin.Default()
	router.Setup(r)
	r.Run()
}
