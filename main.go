package main

import (
	"fmt"
	"log"
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/server"
)

func main() {
	conf := config.NewAppConfig()
	app := server.Server()

	host := conf.Fiber.Host
	port := conf.Fiber.Port

	fmt.Println("Server is running on port: " + port)

	err := app.Listen(host + ":" + port)
	if err != nil {
		log.Panic(err)
	}
}
