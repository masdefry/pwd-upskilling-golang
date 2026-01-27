package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorResponse(err error) map[string]string {
	result := make(map[string]string)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())

			switch fe.Tag() {
			case "required":
				result[field] = field + " is required"
			case "email":
				result[field] = "invalid email format"
			case "min":
				result[field] = field + " must be at least " + fe.Param() + " characters"
			default:
				result[field] = "invalid value"
			}
		}
	}

	return result
}
