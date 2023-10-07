package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(validate *validator.Validate, src interface{}) ([]interface{}, error) {
	if err := validate.Struct(src); err != nil {
		var errs []interface{}
		validatonErr := err.(validator.ValidationErrors)
		for _, i := range validatonErr {
			msg := fmt.Sprintf("[%s] failed on the '%s' validation", strings.ToLower(i.Field()), i.Tag())
			errs = append(errs, msg)
		}

		return errs, err
	}

	return nil, nil
}
