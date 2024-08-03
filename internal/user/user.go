package user

import (
	"context"
	"crypto/subtle"
	"time"

	"github.com/elojah/trax/pkg/paginate"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"golang.org/x/crypto/argon2"
)

type Agg interface {
	transaction.Transactioner

	Store

	StoreEntity

	StoreRole
	StorePermission
	StoreRoleUser

	CreateJWT(context.Context, U, string, time.Duration) (string, error)
	ReadJWT(context.Context, string) (Claims, error)

	Auth(context.Context, string) (Claims, error)
}

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	Email    *string
	GoogleID *string

	// ListByEntity only
	EntityIDs []ulid.ID
	RoleIDs   []ulid.ID

	*paginate.Paginate
	Search string
}

type Patch struct {
	FirstName *string
	LastName  *string
	AvatarURL *string
	UpdatedAt *int64
}

type Store interface {
	Insert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
	List(context.Context, Filter) ([]U, uint64, error)
	Update(context.Context, Filter, Patch) ([]U, error)
	Delete(context.Context, Filter) error

	FetchWithPassword(context.Context, Filter) (U, error)
	ListByEntity(context.Context, Filter) ([]U, uint64, error)
	ListByRole(context.Context, Filter) (map[string][]U, uint64, error)
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
