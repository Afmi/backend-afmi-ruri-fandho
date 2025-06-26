package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	Port                 string
	Secret               string
	LimitCountPerRequest int64
}

func ServerConfig() string {
	// Load environment variables from .env file
	viper.SetConfigFile(".env") // or you can use viper.AddConfigPath() if you have a config folder
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file, %s", err)
	}

	// Automatically read environment variables if defined
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "5001")

	// Return formatted string with server address
	appServer := fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
	log.Println("Server Running at:", appServer)
	return appServer
}
