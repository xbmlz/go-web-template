package validator

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func Init() {
	v := validator.New()

	Validate = v
}
