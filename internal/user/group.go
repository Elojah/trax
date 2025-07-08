package user

import (
	"context"

	"github.com/elojah/trax/pkg/paginate"
	"github.com/elojah/trax/pkg/ulid"
)

type FilterGroup struct {
	ID  ulid.ID
	IDs []ulid.ID

	*paginate.Paginate
	Search string
}

type PatchGroup struct {
	Name        *string
	Description *string
	AvatarURL   *string
	UpdatedAt   *int64
}

type StoreGroup interface {
	InsertGroup(context.Context, Group) error
	UpdateGroup(context.Context, FilterGroup, PatchGroup) ([]Group, error)
	FetchGroup(context.Context, FilterGroup) (Group, error)
	ListGroup(context.Context, FilterGroup) ([]Group, uint64, error)
	DeleteGroup(context.Context, FilterGroup) ([]Group, error)
}

func (g Group) Check() error {
	if g.Name == "" {
		return ErrEmptyName{}
	} else if len(g.Name) > 255 {
		return ErrInvalidNameLength{
			Name:   g.Name,
			Length: 255,
		}
	}

	return nil
}
