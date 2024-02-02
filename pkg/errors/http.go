package errors

import "fmt"

type ErrInvalidStatus struct {
	Status string
}

func (e ErrInvalidStatus) Error() string {
	return fmt.Sprintf("invalid HTTP status %s", e.Status)
}

type ErrInvalidClaim struct {
	Claim string
}

func (e ErrInvalidClaim) Error() string {
	return fmt.Sprintf("invalid claim %s", e.Claim)
}
