package auth

import (
	"project/go-fiber-boilerplate/interfaces"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils/constants"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) interfaces.NewAuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	data := &dto.Register{}

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			constants.NewBadRequestError(err.Error()),
		)
	}

	err = h.authService.Register(data)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(
		constants.NewCreated("register success", nil),
	)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	req := &dto.Login{}

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	token, err := h.authService.Login(req)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		MaxAge:   int(token.AccessTokenExpired),
		SameSite: "disabled",
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		MaxAge:   int(token.RefreshTokenExpired),
		SameSite: "disabled",
		Domain:   "localhost",
	})

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "login success",
		"result":  token,
	})
}
