package user

import (
	"context"

	"github.com/elojah/trax/pkg/ulid"
)

type PatchEntity struct {
	Name      *string
	AvatarURL *string
	UpdatedAt *int64
}

type FilterEntity struct {
	ID  ulid.ID
	IDs []ulid.ID
}

type StoreEntity interface {
	InsertEntity(context.Context, Entity) error
	UpdateEntity(context.Context, FilterEntity, PatchEntity) ([]Entity, error)
	FetchEntity(context.Context, FilterEntity) (Entity, error)
	FetchManyEntity(context.Context, FilterEntity) ([]Entity, error)
	DeleteEntity(context.Context, FilterEntity) error
}

func (e Entity) Check() error {
	if e.Name == "" {
		return ErrEmptyName{}
	} else if len(e.Name) > 255 {
		return ErrInvalidNameLength{
			Name:   e.Name,
			Length: 255,
		}
	}

	return nil
}
