package main

import (
	"project/go-fiber-boilerplate/infrastructure/http/server"
)

func main() {
	server.NewServer().Run()
}
