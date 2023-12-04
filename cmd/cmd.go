package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "mangosteen",
	}

	serverCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}

	dbCmd := &cobra.Command{
		Use: "db",
	}

	mgrtCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}

	rootCmd.AddCommand(dbCmd, serverCmd)
	dbCmd.AddCommand(mgrtCmd)

	database.Connect()
	defer database.Close()

	rootCmd.Execute()
}

func RunServer() {
	r := gin.Default()
	router.Setup(r)
	r.Run()
}
