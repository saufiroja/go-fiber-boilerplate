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

	initilized(app)

	return app
}

func initilized(app *fiber.App) {
	db := postgres.NewPostgres(&conf)
	// repository
	authRepository := repo.NewAuthRepository(db)
	userRepository := repo.NewUserRepository(db)
	// service
	authService := service.NewAuthService(authRepository)
	userService := service.NewUserService(userRepository)
	
	// controllers
	authControllers := controllers.NewAuthControllers(authService)
	userControllers := controller.NewUserControllers(userService)

	routes.NewRoutes(
		authControllers,
		userControllers
	)
}