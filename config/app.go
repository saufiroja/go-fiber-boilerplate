package config

import (
	"log"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func initApp(conf *AppConfig) {
	env := os.Getenv("GO_ENV")

	switch cases.Lower(language.English).String(env) {
	case "development":
		conf.App.Env = "development"
		log.Println("Running in development mode")
	case "production":
		conf.App.Env = "production"
		log.Println("Running in production mode")
	default:
		conf.App.Env = "development"
		log.Println("Running in development mode")
	}
}
