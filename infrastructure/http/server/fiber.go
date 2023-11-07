package server

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/database"
	"project/go-fiber-boilerplate/infrastructure/http/handler/auth"
	"project/go-fiber-boilerplate/infrastructure/http/handler/user"
	"project/go-fiber-boilerplate/utils/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Server() *fiber.App {
	conf := config.NewAppConfig()
	app := fiber.New()

	db := database.NewPostgres(conf)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.ConfigDefault))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			constants.NewSuccess("welcome to go fiber boilerplate", nil),
		)
	})

	// routes
	auth.NewAuthRoutes(app, conf, db)
	user.NewUserRoutes(app, conf, db)

	return app
}
