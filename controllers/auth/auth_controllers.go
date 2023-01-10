package auth

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
	service interfaces.UserService
}

func NewAuthControllers(service interfaces.UserService) interfaces.UserControllers {
	return &Controllers{
		service: service,
	}
}

func (controllers *Controllers) Register(c *fiber.Ctx) error {
	data := &dto.Register{}

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": err.Error(),
		})
	}

	err = controllers.service.Register(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Register success",
	})
}
