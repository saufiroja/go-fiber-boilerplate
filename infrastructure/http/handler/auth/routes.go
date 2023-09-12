package auth

import (
	"project/go-fiber-boilerplate/config"
	repo "project/go-fiber-boilerplate/repository/postgres/user"
	service "project/go-fiber-boilerplate/service/auth"
	"project/go-fiber-boilerplate/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewAuthRoutes(app *fiber.App, conf *config.AppConfig, db *gorm.DB) {

	token := utils.NewGenerateToken()
	password := utils.NewPassword()

	repoAuth := repo.NewUserRepository(db)

	serviceAuth := service.NewAuthService(repoAuth, token, password)

	handlerAuth := NewAuthHandler(serviceAuth)

	// routes
	auth := app.Group("/auth")
	auth.Post("/register", handlerAuth.Register)
	auth.Post("/login", handlerAuth.Login)
}
