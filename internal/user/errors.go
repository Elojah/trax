package user

import "fmt"

type ErrInvalidClaims struct {
	Err error
}

func (e ErrInvalidClaims) Error() string {
	return fmt.Sprintf("invalid claims: %s", e.Err)
}

type ErrInvalidPassword struct{}

func (e ErrInvalidPassword) Error() string {
	return fmt.Sprintf("invalid password")
}
