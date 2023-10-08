package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *Response) Error() string {
	return r.Message
}

func HandlerError(err error) error {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		message := formatValidationErrorMessage(validationErr)
		return NewCustomError(400, message)
	}

	// Handle other types of errors if needed
	return NewCustomError(500, "Internal Server Error")
}

func formatValidationErrorMessage(errs validator.ValidationErrors) string {
	fieldErrors := make([]string, len(errs))
	for i, err := range errs {
		switch err.Tag() {
		case "required":
			fieldErrors[i] = fmt.Sprintf("%s is required", err.Field())
		case "email":
			fieldErrors[i] = fmt.Sprintf("%s is not valid", err.Field())
		case "min":
			fieldErrors[i] = fmt.Sprintf("%s is too short", err.Field())
		case "max":
			fieldErrors[i] = fmt.Sprintf("%s is too long", err.Field())
		}
	}
	return fmt.Sprintf("Validation error(s): %s", joinErrors(fieldErrors, "; "))
}

func joinErrors(errors []string, separator string) string {
	return strings.Join(errors, separator)
}

func NewCustomError(code int, message string) error {
	return &Response{
		Message: message,
		Code:    code,
	}
}
