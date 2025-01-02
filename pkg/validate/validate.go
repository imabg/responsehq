package validate

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func NewValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func Struct(s interface{}) error {
	return validate.Struct(s)
}
