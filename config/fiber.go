package config

import (
	"log"
	"os"
)

const (
	DefaultHost = "localhost"
	DefaultPort = "3000"
)

func initFiber(conf *AppConfig) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	switch {
	case host == "":
		host = DefaultHost
		log.Printf("HOST is not set, using default: %s", DefaultHost)
	case port == "":
		port = DefaultPort
		log.Printf("PORT is not set, using default: %s", DefaultPort)
	}

	conf.Fiber.Host = host
	conf.Fiber.Port = port
}
