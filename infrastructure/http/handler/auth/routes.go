package auth

import (
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/database"
	repo "project/go-fiber-boilerplate/repository/postgres/user"
	service "project/go-fiber-boilerplate/service/auth"
	"project/go-fiber-boilerplate/utils"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(app *fiber.App, conf *config.AppConfig) {
	// config
	db := database.NewPostgres(conf)

	repoAuth := repo.NewUserRepository(db)
	token := utils.NewGenerateToken()
	password := utils.NewPassword()
	serviceAuth := service.NewAuthService(repoAuth, token, password)
	handlerAuth := NewAuthHandler(serviceAuth)

	// routes
	auth := app.Group("/auth")
	auth.Post("/register", handlerAuth.Register)
	auth.Post("/login", handlerAuth.Login)
}
