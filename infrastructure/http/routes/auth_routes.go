package routes

import (
	"project/go-fiber-boilerplate/config"
	controllers "project/go-fiber-boilerplate/controllers/auth"
	"project/go-fiber-boilerplate/infrastructure/database/postgres"
	repo "project/go-fiber-boilerplate/repository/auth"
	service "project/go-fiber-boilerplate/service/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, conf config.AppConfig) {
	db := postgres.NewPostgres(&conf)

	userRepository := repo.NewAuthRepository(db)
	userService := service.NewAuthService(userRepository)
	userControllers := controllers.NewAuthControllers(userService)

	app.Post("/register", userControllers.Register)
}
