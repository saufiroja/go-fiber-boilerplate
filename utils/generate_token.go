package utils

import (
	"encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id, email, fullname string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"fullname": fullname,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(id, email, fullname string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"fullname": fullname,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
