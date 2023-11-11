package auth

import (
	"database/sql"
	repo "project/go-fiber-boilerplate/repository/postgres/user"
	service "project/go-fiber-boilerplate/service/auth"
	"project/go-fiber-boilerplate/utils"
	"project/go-fiber-boilerplate/utils/constants"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(app *fiber.App, db *sql.DB) {

	token := utils.NewGenerateToken()
	password := utils.NewPassword()
	validate := constants.NewValidationError()

	repoAuth := repo.NewUserRepository(db)

	serviceAuth := service.NewAuthService(repoAuth, token, password, validate)

	handlerAuth := NewAuthHandler(serviceAuth)

	// routes
	auth := app.Group("/auth")
	auth.Post("/register", handlerAuth.Register)
	auth.Post("/login", handlerAuth.Login)
}
