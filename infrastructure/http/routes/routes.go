package routes

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/http/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(app *fiber.App, conf *config.AppConfig) {
	auth.NewAuthRoutes(app, conf)
}
