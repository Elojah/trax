package user

import (
	"context"

	"github.com/elojah/trax/pkg/paginate"
	"github.com/elojah/trax/pkg/ulid"
)

type FilterInvitation struct {
	ID  ulid.ID
	IDs []ulid.ID

	Email  *string
	Emails []string

	*paginate.Paginate
	Search string

	GroupIDs []ulid.ID
}

type PatchInvitation struct {
	Email     *string
	UpdatedAt *int64
}

type StoreInvitation interface {
	InsertInvitation(context.Context, Invitation) error
	InsertBatchInvitation(context.Context, ...Invitation) error
	UpdateInvitation(context.Context, FilterInvitation, PatchInvitation) ([]Invitation, error)
	FetchInvitation(context.Context, FilterInvitation) (Invitation, error)
	ListInvitation(context.Context, FilterInvitation) ([]Invitation, uint64, error)
	ListInvitationByGroup(context.Context, FilterInvitation) (map[string][]Invitation, uint64, error)
	DeleteInvitation(context.Context, FilterInvitation) error
}

type FilterInvitationRole struct {
	InvitationID  ulid.ID
	InvitationIDs []ulid.ID

	RoleID  ulid.ID
	RoleIDs []ulid.ID
}

type StoreInvitationRole interface {
	InsertInvitationRole(context.Context, InvitationRole) error
	InsertBatchInvitationRole(context.Context, ...InvitationRole) error
	FetchInvitationRole(context.Context, FilterInvitationRole) (InvitationRole, error)
	ListInvitationRole(context.Context, FilterInvitationRole) ([]InvitationRole, error)
	DeleteInvitationRole(context.Context, FilterInvitationRole) error
}

func (i Invitation) Check() error {
	if i.Email == "" {
		return ErrEmptyEmail{}
	} else if len(i.Email) > 255 {
		return ErrInvalidEmailLength{
			Email:  i.Email,
			Length: 255,
		}
	}

	return nil
}
