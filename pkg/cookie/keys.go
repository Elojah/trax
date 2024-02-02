package cookie

import (
	"context"

	"github.com/gorilla/securecookie"
)

type FilterKeys struct {
	All bool
}

type CacheKeys interface {
	ReadKeys(context.Context, FilterKeys) ([]Keys, error)

	CreateKeys(context.Context, ...Keys) error
	RotateKeys(context.Context, Keys, int64) error
	DeleteKeys(context.Context, FilterKeys) error
}

func NewKeys() Keys {
	return Keys{
		Hash:  securecookie.GenerateRandomKey(keyLength),
		Block: securecookie.GenerateRandomKey(32), //nolint: gomnd
	}
}
