package errors

import "fmt"

type ErrMissingTransaction struct{}

func (e ErrMissingTransaction) Error() string {
	return "transaction not initialized"
}

type ErrNotFound struct {
	Resource string
	Index    string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("resource %s not found for %s", e.Resource, e.Index)
}

type ErrConflict struct {
	Resource string
	Index    string
}

func (e ErrConflict) Error() string {
	return fmt.Sprintf("resource %s already exists for %s", e.Resource, e.Index)
}
