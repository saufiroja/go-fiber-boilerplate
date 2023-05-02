package routes

import (
	"project/go-fiber-boilerplate/infrastructure/http/middlewares"
	"project/go-fiber-boilerplate/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	// every new handler
	authControllers interfaces.AuthControllers
	userControllers interfaces.UserControllers
}

func NewRoutes(
	authControllers interfaces.AuthControllers,
	userControllers interfaces.UserControllers,
) *Routes {
	return &Routes{
		authControllers: authControllers,
		userControllers: userControllers,
	}
}

func (r *Routes) InitRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Post("/register", r.authControllers.Register)
	v1.Post("/login", r.authControllers.Login)

	v1.Get("/users", r.userControllers.FindAllUsers, middlewares.MiddlewaresUser)
	v1.Get("/users/:id", r.userControllers.FindUserByID, middlewares.MiddlewaresUser)
	v1.Put("/users/:id", r.userControllers.UpdateUserByID, middlewares.MiddlewaresUser)
	v1.Delete("/users/:id", r.userControllers.DeleteUserByID, middlewares.MiddlewaresUser)
}
