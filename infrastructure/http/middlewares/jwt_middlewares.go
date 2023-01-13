package middlewares

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func MiddlewaresUser(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	refreshToken := c.Cookies("refresh_token")
	secret := os.Getenv("JWT_SECRET")

	if accessToken == "" && refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    "401",
			"message": "Unauthorized",
		})
	}

	if accessToken != "" {
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "401",
				"message": "Unauthorized",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("id", claims["id"])
			c.Locals("email", claims["email"])
			c.Locals("fullname", claims["fullname"])
			return c.Next()
		}
	}

	if refreshToken != "" {
		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "401",
				"message": "Unauthorized",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("id", claims["id"])
			c.Locals("email", claims["email"])
			c.Locals("fullname", claims["fullname"])
			return c.Next()
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    "401",
		"message": "Unauthorized",
	})
}
