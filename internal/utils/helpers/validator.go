package helpers

import "github.com/go-playground/validator"

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidtor() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if cv == nil || cv.validator == nil {
		return nil
	}

	return cv.validator.Struct(i)
}
