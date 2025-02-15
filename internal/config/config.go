package config

import (
	"log"
	"os"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
)

func Read() *model.Config {
	var config model.Config
	config.AppConfig = readAppConfig()
	config.DBConfig = readDBConfig()
	config.AuthConfig = readAuthConfig()
	return &config
}

func readDBConfig() model.DBConfig {
	var config model.DBConfig
	config.DBConnection = getEnv("DB_CONNECTION_STRING")
	return config
}

func readAppConfig() model.AppConfig {
	var config model.AppConfig
	config.Host = getEnv("HOST")
	config.ThisURL = getEnv("THIS_URL")
	config.Debug = getEnv("DEBUG")
	return config
}

func readAuthConfig() model.AuthConfig {
	var config model.AuthConfig
	config.Host = getEnv("AUTH_HOST")
	return config
}

func getEnv(key string) string {
	var data string
	if data = os.Getenv(key); data == "" {
		log.Fatalf("не указан - %s", key)
	}
	return data
}
