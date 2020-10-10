package errors

import "fmt"

type MissingPropertiesError struct {
	missingProperties []string
}

func NewMissingPropertiesError(props []string) *ValidationError {
	return &MissingPropertiesError{
		missingProperties: props,
	}
}

func (v *MissingPropertiesError) Error() string {
	return fmt.Sprintf("The following fields have uninitialized values: %s", v.missingProperties)
}
