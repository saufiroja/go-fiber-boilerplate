package main

import (
	"log"
	"os"
	"os/signal"
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/http/server"
	"syscall"
)

func main() {
	conf := config.NewAppConfig()
	app := server.Server()

	host := conf.Fiber.Host
	port := conf.Fiber.Port

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(host + ":" + port); err != nil {
			log.Fatalf("error when listening to %s:%s, %s", host, port, err)
		}
	}()

	log.Printf("server is running on %s:%s", host, port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("error when shutting down the server, %s", err)
	}

	log.Println("process clean up...")
}
