package user

import (
	"context"
	"time"

	"github.com/elojah/trax/pkg/ulid"
)

type App interface {
	Store

	CreateJWT(context.Context, U, string, time.Duration) (string, error)
	ReadJWT(context.Context, string) (Claims, error)

	Auth(context.Context, string) (U, error)
}

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	GoogleID *string
	TwitchID *string
}

type Store interface {
	Insert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
	FetchMany(context.Context, Filter) ([]U, error)
	Delete(context.Context, Filter) error
}
