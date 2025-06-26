package main

import (
	"time"
	"user-test/config"
	"user-test/infra/database"
	"user-test/infra/logger"
	"user-test/routers"

	"github.com/spf13/viper"
)

func main() {
	// Set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	// Load configuration
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	// Connect to DB and run migrations
	if err := database.DbConnection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	// Setup router
	router := routers.SetupRoute()

	// Get server address
	serverAddress := config.ServerConfig()

	// Run server
	if err := router.Run(serverAddress); err != nil {
		logger.Fatalf("Failed to start server: %s", err)
	}
}
