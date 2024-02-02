package errors

import "fmt"

type ErrInvalidClaims struct {
	Err error
}

func (e ErrInvalidClaims) Error() string {
	return fmt.Sprintf("invalid claims: %s", e.Err)
}
