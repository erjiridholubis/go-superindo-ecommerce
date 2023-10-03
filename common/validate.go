package common

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func ValidateStruct(toValidate interface{}) []*ErrorResponse {
	var errorResponse []*ErrorResponse
	err := validate.Struct(toValidate)
	if err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			for _, valErr := range valErrs {
				errorResponse = append(errorResponse, &ErrorResponse{
					Field:   valErr.Field(),
					Message: valErr.Error(),
				})
			}
		}
	}
	return errorResponse
}
