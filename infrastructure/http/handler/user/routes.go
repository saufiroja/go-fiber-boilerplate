package user

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/http/middlewares"
	repo "project/go-fiber-boilerplate/repository/postgres/user"
	service "project/go-fiber-boilerplate/service/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewUserRoutes(app *fiber.App, conf *config.AppConfig, db *gorm.DB) {
	repoUser := repo.NewUserRepository(db)

	serviceUser := service.NewUserService(repoUser)

	handlerUser := NewUserHandler(serviceUser)

	// routes
	user := app.Group("/user")
	user.Use(middlewares.MiddlewaresUser)
	user.Get("/", handlerUser.FindAllUsers)
	user.Get("/:id", handlerUser.FindUserByID)
	user.Put("/:id", handlerUser.UpdateUserByID)
	user.Delete("/:id", handlerUser.DeleteUserByID)
}
