package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ToErrorMessages(ve validator.ValidationErrors) []string {
	errorMessages := make([]string, len(ve))

	for i, fieldError := range ve {
		errorMessages[i] = toErrorMessage(fieldError)
	}

	return errorMessages
}

func toErrorMessage(err validator.FieldError) string {
	var field = err.Field()
	var message = ""
	switch err.Tag() {
	case "required":
		message = "is required"
	case "email":
		message = "is not a valid email"
	case "min":
		message = fmt.Sprintf("should be min %s characters", err.Param())
	case "max":
		message = fmt.Sprintf("should be max %s characters", err.Param())
	default:
		message = err.Error()
	}

	return fmt.Sprintf("'%s' %s", field, message)
}
