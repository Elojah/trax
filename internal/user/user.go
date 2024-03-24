package user

import (
	"context"
	"crypto/subtle"
	"time"

	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"golang.org/x/crypto/argon2"
)

type App interface {
	transaction.Transactioner

	Store
	StoreProfile

	CreateJWT(context.Context, U, string, time.Duration) (string, error)
	ReadJWT(context.Context, string) (Claims, error)

	Auth(context.Context, string) (U, error)
}

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	Email    *string
	GoogleID *string
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

const (
	ntime   uint32 = 1
	memory  uint32 = 64 * 1024
	threads uint8  = 4
	keylen  uint32 = 16 // salt ulid length in bytes
)

func Encrypt(password string) ([]byte, []byte) {
	salt := ulid.NewID()
	hash := argon2.IDKey([]byte(password), salt, ntime, memory, threads, keylen)

	return hash, salt
}

func (u U) Check(password string) error {
	hash := argon2.IDKey([]byte(password), u.PasswordSalt, ntime, memory, threads, keylen)

	if subtle.ConstantTimeCompare(u.PasswordHash, hash) != 1 {
		return ErrInvalidPassword{}
	}

	return nil
}
