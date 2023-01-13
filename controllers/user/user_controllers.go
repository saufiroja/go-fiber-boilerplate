package user

import (
	"project/go-fiber-boilerplate/dto"
	"project/go-fiber-boilerplate/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
	service interfaces.UserService
}

func NewUserControllers(service interfaces.UserService) interfaces.UserControllers {
	return &Controllers{
		service: service,
	}
}

func (controller *Controllers) FindAllUsers(c *fiber.Ctx) error {
	users, err := controller.service.FindAllUsers()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "200",
		"message": "success get all users",
		"result":  users,
	})
}

func (controller *Controllers) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := controller.service.FindUserByID(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "200",
		"message": "success get user by id",
		"result":  user,
	})
}

func (controller *Controllers) UpdateUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	data := &dto.UpdateUserByID{}

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	err = controller.service.UpdateUserByID(id, data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "200",
		"message": "success update user by id",
	})
}

func (controller *Controllers) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.service.DeleteUserByID(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "200",
		"message": "success delete user by id",
	})
}
