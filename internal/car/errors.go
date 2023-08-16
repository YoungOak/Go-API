package car

import (
	"fmt"
)

type ErrorFieldMissing struct {
	Field string
}

func (e ErrorFieldMissing) Error() string {
	return fmt.Sprintf("car field '%s' missing", e.Field)
}

type ErrorFieldInvalid struct {
	Field string
	Value any
}

func (e ErrorFieldInvalid) Error() string {
	return fmt.Sprintf("car field '%s' invalid value: '%v'", e.Field, e.Value)
}
