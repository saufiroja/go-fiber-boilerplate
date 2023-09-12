package config

import (
	"github.com/joho/godotenv"
)

// struct for the config file
type AppConfig struct {
	App struct {
		Env string
	}
	Fiber struct {
		Port string
	}
	Postgres struct {
		Host string
		Port string
		User string
		Pass string
		Name string
		Ssl  string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	_ = godotenv.Load()

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initFiber(appConfig)
		initPostgres(appConfig)
	}

	return appConfig
}
