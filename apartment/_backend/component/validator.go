package component

import (
	gpvalidator "github.com/go-playground/validator"
)

type validator struct {
	validator *gpvalidator.Validate
}

func NewValidator() *validator {
	return &validator{
		validator: gpvalidator.New(),
	}
}

func (v *validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
