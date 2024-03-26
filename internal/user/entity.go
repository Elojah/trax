package user

import (
	"context"

	"github.com/elojah/trax/pkg/ulid"
)

type FilterEntity struct {
	ID  ulid.ID
	IDs []ulid.ID
}

type StoreEntity interface {
	InsertEntity(context.Context, Entity) error
	FetchEntity(context.Context, FilterEntity) (Entity, error)
	FetchManyEntity(context.Context, FilterEntity) ([]Entity, error)
	DeleteEntity(context.Context, FilterEntity) error
}
