package config

import (
	"log"

	"github.com/joho/godotenv"
)

// struct for the config file
type AppConfig struct {
	App struct {
		Env string
	}
	Fiber struct {
		Host string
		Port string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initFiber(appConfig)
	}

	return appConfig
}
