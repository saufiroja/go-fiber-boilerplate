package main

import (
	"log"
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/http/server"
)

func main() {
	conf := config.NewAppConfig()
	app := server.Server()

	host := conf.Fiber.Host
	port := conf.Fiber.Port

	err := app.Listen(host + ":" + port)
	if err != nil {
		log.Panic(err)
	}
}
