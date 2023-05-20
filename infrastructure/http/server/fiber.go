package server

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/database"
	authHandler "project/go-fiber-boilerplate/infrastructure/http/handler/auth"
	userHandler "project/go-fiber-boilerplate/infrastructure/http/handler/user"
	"project/go-fiber-boilerplate/infrastructure/http/routes"
	"project/go-fiber-boilerplate/repository/postgres"
	authService "project/go-fiber-boilerplate/service/auth"
	userService "project/go-fiber-boilerplate/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Server() *fiber.App {
	conf := config.NewAppConfig()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2 Jan 2006 15:04:05",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"code":    200,
			"message": "API is running",
		})
	})

	initilized(app, conf)

	return app
}

func initilized(app *fiber.App, conf *config.AppConfig) {
	db := database.NewPostgres(conf)
	// repository
	authRepository := postgres.NewAuthRepository(db)
	userRepository := postgres.NewUserRepository(db)
	// service
	authService := authService.NewAuthService(authRepository)
	userService := userService.NewUserService(userRepository)

	// controllers
	authControllers := authHandler.NewAuthControllers(authService)
	userControllers := userHandler.NewUserControllers(userService)

	routes.NewRoutes(
		authControllers,
		userControllers,
	)
}
