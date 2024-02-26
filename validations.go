package inkstone

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	VALIDATION_EMAIL               string = "inkEmail"
	VALIDATION_NOT_ALL_FIELDS_ZERO string = "notAllFieldsZero"
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

func ValidationNotAllFieldsZero(sl validator.StructLevel) {
	value := sl.Current().Interface()

	t := reflect.TypeOf(value)

	structName := t.Name()

	v := reflect.ValueOf(value)

	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)

		if !fieldValue.IsZero() {
			return
		}
	}

	sl.ReportError(value, structName, "", VALIDATION_NOT_ALL_FIELDS_ZERO, "")
}
