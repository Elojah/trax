package user

import (
	"context"

	"github.com/elojah/trax/pkg/paginate"
	"github.com/elojah/trax/pkg/ulid"
)

type PatchEntity struct {
	Name        *string
	Description *string
	AvatarURL   *string
	UpdatedAt   *int64
}

type FilterEntity struct {
	ID  ulid.ID
	IDs []ulid.ID

	*paginate.Paginate
	Search string
}

type StoreEntity interface {
	InsertEntity(context.Context, Entity) error
	UpdateEntity(context.Context, FilterEntity, PatchEntity) ([]Entity, error)
	FetchEntity(context.Context, FilterEntity) (Entity, error)
	ListEntity(context.Context, FilterEntity) ([]Entity, uint64, error)
	DeleteEntity(context.Context, FilterEntity) error
}

type FilterEntityProfile struct {
	EntityID  ulid.ID
	EntityIDs []ulid.ID
}

type PatchEntityProfile struct {
	AvatarURL   *string
	Description *string
	UpdatedAt   *int64
}

type StoreEntityProfile interface {
	InsertEntityProfile(context.Context, EntityProfile) error
	FetchEntityProfile(context.Context, FilterEntityProfile) (EntityProfile, error)
	ListEntityProfile(context.Context, FilterEntityProfile) ([]EntityProfile, error)
	UpdateEntityProfile(context.Context, FilterEntityProfile, PatchEntityProfile) ([]EntityProfile, error)
	DeleteEntityProfile(context.Context, FilterEntityProfile) error
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
