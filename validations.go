package inkstone

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	VALIDATION_EMAIL string = "inkEmail"
)

func emailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}

	return matched
}
