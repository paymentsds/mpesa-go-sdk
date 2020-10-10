package errors

import "fmt"

type ValidationError struct {
	properties []string
}

func NewValidationError(props []string) *ValidationError {
	return &ValidationError{
		properties: props,
	}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("The following fields have validation errors: %s", v.properties)
}
