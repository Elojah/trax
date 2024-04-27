package user

import (
	"context"
	"time"

	"github.com/elojah/trax/pkg/paginate"
	"github.com/elojah/trax/pkg/ulid"
)

var (
	Resources = map[Resource]struct{}{
		R_asset:     {},
		R_entity:    {},
		R_operation: {},
		R_role:      {},
		R_user:      {},
	}

	Commands = map[Command]struct{}{
		C_read:   {},
		C_create: {},
		C_update: {},
		C_delete: {},
	}
)

type FilterRole struct {
	ID  ulid.ID
	IDs []ulid.ID

	EntityID ulid.ID

	*paginate.Paginate
	Search string
}

type StoreRole interface {
	InsertRole(context.Context, Role) error
	FetchRole(context.Context, FilterRole) (Role, error)
	ListRole(context.Context, FilterRole) ([]Role, uint64, error)
	DeleteRole(context.Context, FilterRole) error
}

type FilterPermission struct {
	RoleID  ulid.ID
	RoleIDs []ulid.ID

	Resource *Resource
	Command  *Command
}

type StorePermission interface {
	InsertPermission(context.Context, Permission) error
	InsertPermissions(context.Context, []Permission) error
	FetchPermission(context.Context, FilterPermission) (Permission, error)
	ListPermission(context.Context, FilterPermission) ([]Permission, error)
	DeletePermission(context.Context, FilterPermission) error
}

type FilterRoleUser struct {
	RoleID  ulid.ID
	RoleIDs []ulid.ID

	UserID  ulid.ID
	UserIDs []ulid.ID
}

type StoreRoleUser interface {
	InsertRoleUser(context.Context, RoleUser) error
	FetchRoleUser(context.Context, FilterRoleUser) (RoleUser, error)
	ListRoleUser(context.Context, FilterRoleUser) ([]RoleUser, error)
	DeleteRoleUser(context.Context, FilterRoleUser) error

	ListClaims(context.Context, ulid.ID) (ClaimAuth, error)
}

func AllPermissions(roleID ulid.ID) []Permission {
	var perms []Permission

	now := time.Now().Unix()

	for r := range Resources {
		for c := range Commands {
			perms = append(perms, Permission{
				RoleID:    roleID,
				Resource:  r,
				Command:   c,
				CreatedAt: now,
				UpdatedAt: now,
			})
		}
	}

	return perms
}
