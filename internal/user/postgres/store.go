package postgres

import (
	"github.com/elojah/trax/internal/user"
)

var _ user.Store = (*Store)(nil)

var _ user.StoreGroup = (*Store)(nil)

var _ user.StoreRole = (*Store)(nil)
var _ user.StorePermission = (*Store)(nil)
var _ user.StoreRoleUser = (*Store)(nil)

type Store struct{}
