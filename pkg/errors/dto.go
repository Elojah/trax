package errors

import "fmt"

type ErrNullRequest struct{}

func (e ErrNullRequest) Error() string {
	return "null request"
}

type ErrMissingAuth struct{}

func (e ErrMissingAuth) Error() string {
	return "missing authentication"
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return "invalid credentials"
}

// ErrInvalidNumericalRange is raised when a numerical value is out of range.
type ErrInvalidNumericalRange struct {
	Key   string
	Value int

	Min int
	Max int
}

// Error implementation for ErrInvalidNumericalRange.
func (e ErrInvalidNumericalRange) Error() string {
	return fmt.Sprintf("key %s must be between %d and %d. current value: %d", e.Key, e.Min, e.Max, e.Value)
}

type ErrMissingAtLeast struct {
	AtLeast int
	Fields  []string
}

func (e ErrMissingAtLeast) Error() string {
	return fmt.Sprintf("missing at least %d field among %v", e.AtLeast, e.Fields)
}

type ErrMissingField struct {
	Field string
}

func (e ErrMissingField) Error() string {
	return fmt.Sprintf("missing field %s", e.Field)
}
