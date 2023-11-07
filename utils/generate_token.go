package utils

import (
	"os"
	"project/go-fiber-boilerplate/utils/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type GenerateToken struct{}

func NewGenerateToken() *GenerateToken {
	return &GenerateToken{}
}

func (g *GenerateToken) GenerateAccessToken(id, email, fullname string) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")
	expired := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"fullname": fullname,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, constants.NewBadRequestError(err.Error())
	}

	return tokenString, expired, nil
}

func (g *GenerateToken) GenerateRefreshToken(id, email, fullname string) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")

	expired := time.Now().Add(time.Hour * 24 * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"fullname": fullname,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, constants.NewBadRequestError(err.Error())
	}

	return tokenString, expired, nil
}
