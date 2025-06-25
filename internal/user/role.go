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
		C_read:  {},
		C_edit:  {},
		C_write: {},
	}
)

type FilterRole struct {
	ID  ulid.ID
	IDs []ulid.ID

	EntityID  ulid.ID
	EntityIDs []ulid.ID

	UserID ulid.ID

	*paginate.Paginate
	Search string
}

type PatchRole struct {
	Name      *string
	UpdatedAt *int64
}

type StoreRole interface {
	InsertRole(context.Context, Role) error
	UpdateRole(context.Context, FilterRole, PatchRole) ([]Role, error)
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

func (r Role) Check() error {
	if r.Name == "" {
		return ErrEmptyName{}
	} else if len(r.Name) > 255 {
		return ErrInvalidNameLength{
			Name:   r.Name,
			Length: 255,
		}
	}

	return nil
}

type Roles []Role

func (rs Roles) IDs() []ulid.ID {
	ids := make([]ulid.ID, 0, len(rs))

	for _, r := range rs {
		ids = append(ids, r.ID)
	}

	return ids
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

type Permissions []Permission

func (ps Permissions) Check() error {
	if len(ps) == 0 {
		return ErrEmptyPermissions{}
	}

	for _, p := range ps {
		if _, ok := Resources[p.Resource]; !ok {
			return ErrUnknownResource{Resource: p.Resource.String()}
		}

		if _, ok := Commands[p.Command]; !ok {
			return ErrUnknownCommand{Command: p.Command.String()}
		}
	}

	return nil
}

func (ps Permissions) ByRole() map[string][]Permission {
	perms := make(map[string][]Permission)

	for _, p := range ps {
		perms[p.RoleID.String()] = append(perms[p.RoleID.String()], p)
	}

	return perms
}
