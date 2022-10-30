package actions

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// DefaultValidator contains go-playground validator instance.
type DefaultValidator struct {
	validator *validator.Validate
}

// Validate validates structs based on go-playground validator.
func (cv *DefaultValidator) Validate(i interface{}) error {
	return fmt.Errorf("struct validation failed %w", cv.validator.Struct(i))
}
