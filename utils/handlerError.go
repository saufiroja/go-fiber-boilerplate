package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandlerError(err error) error {
	var message string

	if obj, ok := err.(validator.ValidationErrors); ok {
		for _, v := range obj {
			switch v.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", v.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid", v.Field())
			case "min":
				message = fmt.Sprintf("%s is too short", v.Field())
			case "max":
				message = fmt.Sprintf("%s is too long", v.Field())

			}
		}
	}

	return errors.New(message)
}
