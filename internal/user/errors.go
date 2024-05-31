package user

import (
	"fmt"
	"strings"
)

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

type ErrUnauthorizedRole struct {
	Roles []string
}

func (e ErrUnauthorizedRole) Error() string {
	return fmt.Sprintf("unauthorized roles: %s", strings.Join(e.Roles, ", "))
}

type ErrEmptyPermissions struct{}

func (e ErrEmptyPermissions) Error() string {
	return fmt.Sprintf("missing permissions")
}

type ErrUnknownResource struct {
	Resource string
}

func (e ErrUnknownResource) Error() string {
	return fmt.Sprintf("unknown resource: %s", e.Resource)
}

type ErrUnknownCommand struct {
	Command string
}

func (e ErrUnknownCommand) Error() string {
	return fmt.Sprintf("unknown command: %s", e.Command)
}

type ErrForbiddenAdminRole struct{}

func (e ErrForbiddenAdminRole) Error() string {
	return fmt.Sprintf("admin role cannot be updated")
}
