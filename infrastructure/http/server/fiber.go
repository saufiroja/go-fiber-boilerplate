package server

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Server() *fiber.App {
	conf := config.NewAppConfig()
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2 Jan 2006 15:04:05",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "APP Boilerplate is up!!!!",
		})
	})

	routes.AuthRoutes(app, *conf)
	routes.UserRoutes(app, *conf)

	return app
}
