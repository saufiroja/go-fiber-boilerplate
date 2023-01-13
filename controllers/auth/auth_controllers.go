package auth

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
	service interfaces.AuthService
}

func NewAuthControllers(service interfaces.AuthService) interfaces.AuthControllers {
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
		"message": "register success",
	})
}

func (controller *Controllers) Login(c *fiber.Ctx) error {
	req := &dto.Login{}

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	token, err := controller.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "200",
		"message": "login success",
		"result":  token,
	})
}
