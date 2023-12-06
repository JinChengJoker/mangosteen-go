package cmd

import (
	"fmt"
	"mangosteen/internal/database"
	"mangosteen/internal/email"
	"mangosteen/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Run() {
	// 创建相关命令
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
	emailCmd := &cobra.Command{
		Use: "email",
		Run: func(cmd *cobra.Command, args []string) {
			email.Send()
		},
	}
	rootCmd.AddCommand(dbCmd, serverCmd, emailCmd)
	dbCmd.AddCommand(mgrtCmd)

	// 读取配置文件，包含密钥等内容
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 连接数据库
	database.Connect()
	defer database.Close()

	rootCmd.Execute()
}

func RunServer() {
	r := gin.Default()
	router.Setup(r)
	r.Run()
}
