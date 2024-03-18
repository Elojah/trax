package postgres

import (
	"github.com/elojah/trax/internal/user"
)

var _ user.Store = (*Store)(nil)
var _ user.StoreProfile = (*Store)(nil)

type Store struct{}
