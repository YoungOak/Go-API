package data

import "fmt"

type ErrorAlreadyExists struct {
	ID string
}

func (e ErrorAlreadyExists) Error() string {
	return fmt.Sprintf("entry with ID '%s' alrady exists", e.ID)
}

type ErrorRecordNotFound struct {
	ID string
}

func (e ErrorRecordNotFound) Error() string {
	return fmt.Sprintf("no record in store with ID: '%s'", e.ID)
}
