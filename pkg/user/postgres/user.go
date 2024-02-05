package postgres

import (
	"context"

	"github.com/elojah/trax/pkg/user"
)

func (s Store) Insert(context.Context, user.U) error {
	return nil
}
func (s Store) Fetch(context.Context, user.Filter) (user.U, error) {
	return user.U{}, nil
}
func (s Store) FetchMany(context.Context, user.Filter) ([]user.U, []byte, error) {
	return nil, nil, nil
}
func (s Store) Delete(context.Context, user.Filter) error {
	return nil
}
