package interfaces

import (
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/models/entity"

	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	Register(user *dto.Register) error
	Login(email string) (*entity.User, error)
}

type AuthService interface {
	Register(user *dto.Register) error
	Login(user *dto.Login) (*dto.LoginResponse, error)
}

type NewAuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
