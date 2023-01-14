package routes

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/database/postgres"
	controller "project/go-fiber-boilerplate/infrastructure/http/controllers/user"
	"project/go-fiber-boilerplate/infrastructure/http/middlewares"
	repo "project/go-fiber-boilerplate/repository/user"
	service "project/go-fiber-boilerplate/service/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, conf config.AppConfig) {
	db := postgres.NewPostgres(&conf)

	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userControllers := controller.NewUserControllers(userService)

	// middleware
	app.Use(middlewares.MiddlewaresUser)

	app.Get("/users", userControllers.FindAllUsers)
	app.Get("/users/:id", userControllers.FindUserByID)
	app.Put("/users/:id", userControllers.UpdateUserByID)
	app.Delete("/users/:id", userControllers.DeleteUserByID)
}
