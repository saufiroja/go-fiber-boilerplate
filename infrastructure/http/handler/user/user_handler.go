package user

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) interfaces.UserHandler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) FindAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.FindAllUsers()
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

func (h *Handler) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.FindUserByID(id)
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

func (h *Handler) UpdateUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	data := &dto.UpdateUserByID{}

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    "400",
			"message": err.Error(),
		})
	}

	err = h.userService.UpdateUserByID(id, data)
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

func (h *Handler) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userService.DeleteUserByID(id)
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
