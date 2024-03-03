package postgres

import (
	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/postgres"
)

var _ user.Store = (*Store)(nil)
var _ user.StoreProfile = (*Store)(nil)

type Store struct {
	*postgres.Service
}
