package config

import (
	"os"
)

func initFiber(conf *AppConfig) {
	port := os.Getenv("PORT")

	conf.Fiber.Port = port
}
