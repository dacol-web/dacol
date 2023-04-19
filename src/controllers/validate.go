package controllers

import (
	"github.com/go-playground/validator"
)

func ValidateMsg(err error) map[string]string {
	errCol := map[string]string{}

	for _, valid := range err.(validator.ValidationErrors) {
		switch valid.Tag() {
		case "required":
			errCol[valid.Field()] = "This field cannot empty"
		case "min":
			errCol[valid.Field()] = "This field must be gread than 8"
		case "email":
			errCol[valid.Field()] = "This field must be valid email"
		case "eqfield":
			errCol[valid.Field()] = "This field must be same"
		}
	}

	return errCol
}
