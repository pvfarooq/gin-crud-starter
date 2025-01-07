package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError formats validation errors into human readable messages.
// It takes a validator.FieldError parameter and returns a formatted error string.
// The function handles specific validation cases:
//   - "required": When a required field is missing
//   - "email": When an email format is invalid
//   - "min": When a field doesn't meet minimum length requirement
//   - Other cases return a generic invalid field message
func FormatValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is not valid", err.Field())
	}
}
