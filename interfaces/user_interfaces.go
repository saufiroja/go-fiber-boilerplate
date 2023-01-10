package interfaces

import (
	"project/go-fiber-boilerplate/dto"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Register(user *dto.Register) error
	// Login(email string) error
}

type UserService interface {
	Register(user *dto.Register) error
	// Login(email, password string) error
}

type UserControllers interface {
	Register(c *fiber.Ctx) error
}
