package interfaces

import (
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/models/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	FindAllUsers() ([]dto.FindAllUsers, error)
	FindUserByID(id string) (*dto.FindUserByID, error)
	UpdateUserByID(id string, user *dto.UpdateUserByID) error
	DeleteUserByID(id string) error
	InsertUser(user *dto.Register) error
	FindUserByEmail(email string) (*entity.User, error)
}

type UserService interface {
	FindAllUsers() ([]dto.FindAllUsers, error)
	FindUserByID(id string) (*dto.FindUserByID, error)
	UpdateUserByID(id string, user *dto.UpdateUserByID) error
	DeleteUserByID(id string) error
}

type UserHandler interface {
	FindAllUsers(c *fiber.Ctx) error
	FindUserByID(c *fiber.Ctx) error
	UpdateUserByID(c *fiber.Ctx) error
	DeleteUserByID(c *fiber.Ctx) error
}
