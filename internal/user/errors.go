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

type ErrUnknownEntity struct {
	EntityID string
}

func (e ErrUnknownEntity) Error() string {
	return fmt.Sprintf("unknown entity: %s", e.EntityID)
}

type ErrUnauthorizedResource struct {
	Resource string
}

func (e ErrUnauthorizedResource) Error() string {
	return fmt.Sprintf("unauthorized resource: %s", e.Resource)
}

type ErrUnauthorizedCommand struct {
	Resource string
	Command  string
}

func (e ErrUnauthorizedCommand) Error() string {
	return fmt.Sprintf("unauthorized command: %s on resource %s", e.Command, e.Resource)
}
