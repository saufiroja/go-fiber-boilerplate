package interfaces

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Register(user *dto.Register) error
	Login(email string) (*entity.User, error)
}

type UserService interface {
	Register(user *dto.Register) error
	Login(email, password string) (*dto.LoginResponse, error)
}

type UserControllers interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
