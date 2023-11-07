package user

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils/constants"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) interfaces.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.FindAllUsers()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(
		constants.NewSuccess("success get all users", users),
	)
}

func (h *userHandler) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.FindUserByID(id)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(
		constants.NewSuccess("success get user by id", user),
	)
}

func (h *userHandler) UpdateUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	data := &dto.UpdateUserByID{}

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(400).JSON(
			constants.NewBadRequestError(err.Error()),
		)
	}

	err = h.userService.UpdateUserByID(id, data)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(
		constants.NewSuccess("success update user by id", nil),
	)
}

func (h *userHandler) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userService.DeleteUserByID(id)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(
		constants.NewSuccess("success delete user by id", nil),
	)
}
