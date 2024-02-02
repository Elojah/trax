package errors

import "fmt"

type ErrNullRequest struct{}

func (e ErrNullRequest) Error() string {
	return fmt.Sprintf("null request")
}

type ErrMissingAuth struct{}

func (e ErrMissingAuth) Error() string {
	return fmt.Sprintf("missing authentication")
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("invalid credentials")
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
