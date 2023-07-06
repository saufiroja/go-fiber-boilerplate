package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandlerError(err error) error {
	var message string
	var code int

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

	return HandlerErrorCustom(code, message)
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *Response) Error() string {
	return r.Message
}

func (r *Response) GetCode() int {
	return r.Code
}

func HandlerErrorCustom(code int, message string) error {
	return &Response{
		Message: message,
		Code:    code,
	}
}
