package routes

import "project/go-fiber-boilerplate/interfaces"

type routes struct {
	// every new handler
	authControllers interfaces.AuthControllers
	userControllers interfaces.UserControllers
}

func NewRoutes (
	authControllers interfaces.AuthControllers,
	userControllers interfaces.UserControllers
){
	return &routes{
		authControllers: authControllers,
		userControllers: userControllers,
	}
}

func (r *routes) InitRoutes(app *fiber.App) {
	app.Post("/register", r.authControllers.Register)
	app.Post("/login", r.authControllers.Login)

	app.Get("/users", r.userControllers.FindAllUsers)
	app.Get("/users/:id", r.userControllers.FindUserByID)
	app.Put("/users/:id", r.userControllers.UpdateUserByID)
	app.Delete("/users/:id", r.userControllers.DeleteUserByID)
}