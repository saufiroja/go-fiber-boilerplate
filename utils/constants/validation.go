package constants

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type IValidation interface {
	Validate(data interface{}) error
	ValidationMessage(err error) error
}

type Validation struct {
	validation *validator.Validate
}

func NewValidationError() IValidation {
	return &Validation{
		validation: validator.New(),
	}
}

func (v *Validation) Validate(data interface{}) error {
	err := v.validation.Struct(data)
	if err != nil {
		return err
	}

	return nil
}

func (v *Validation) ValidationMessage(err error) error {
	var message string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", e.Field())
		case "email":
			message = fmt.Sprintf("%s is not valid", e.Field())
		case "min":
			message = fmt.Sprintf("%s is too short", e.Field())
		case "max":
			message = fmt.Sprintf("%s is too long", e.Field())
		}
	}

	return NewBadRequestError(message)
}
