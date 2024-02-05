package postgres

import (
	"github.com/elojah/trax/pkg/user"
)

var _ user.Store = (*Store)(nil)

type Store struct{}
