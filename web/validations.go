package web

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	VALIDATION_EMAIL               = "inkEmail"
	VALIDATION_PHONE               = "inkPhone"
	VALIDATION_NOT_ALL_FIELDS_ZERO = "notAllFieldsZero"
)

func ValidationEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}

	return matched
}

func ValidationPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	phoneRegex := `^1[3-9]\d{9}$`

	matched, err := regexp.MatchString(phoneRegex, phone)
	if err != nil {
		return false
	}

	return matched
}

func ValidationNotAllFieldsZero(sl validator.StructLevel) {
	value := sl.Current().Interface()

	t := reflect.TypeOf(value)

	structName := t.Name()

	v := reflect.ValueOf(value)

	for i := 0; i < t.NumField(); i++ {
		if f := t.Field(i); f.Name == "Id" {
			continue
		}

		fieldValue := v.Field(i)

		if !fieldValue.IsZero() {
			return
		}
	}

	sl.ReportError(value, structName, "", VALIDATION_NOT_ALL_FIELDS_ZERO, "")
}
