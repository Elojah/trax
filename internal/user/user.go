package user

import (
	"context"
	"time"

	"github.com/elojah/trax/pkg/ulid"
)

type App interface {
	Store
	StoreProfile

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

type FilterProfile struct {
	UserID  ulid.ID
	UserIDs []ulid.ID
}

type StoreProfile interface {
	InsertProfile(context.Context, Profile) error
	FetchProfile(context.Context, FilterProfile) (Profile, error)
	FetchManyProfile(context.Context, FilterProfile) ([]Profile, error)
	DeleteProfile(context.Context, FilterProfile) error
}
