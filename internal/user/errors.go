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

type ErrEmptyName struct{}

func (e ErrEmptyName) Error() string {
	return fmt.Sprintf("empty name")
}

type ErrInvalidNameLength struct {
	Name   string
	Length int
}

func (e ErrInvalidNameLength) Error() string {
	return fmt.Sprintf("invalid name length %d : %s", e.Length, e.Name)
}
