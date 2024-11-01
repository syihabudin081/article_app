package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ParseValidationErrors(err error) string {
	var errorMsg []string
	for _, err := range err.(validator.ValidationErrors) {
		errorMsg = append(errorMsg, err.Field()+": "+err.ActualTag())
	}
	return strings.Join(errorMsg, ", ")
}
