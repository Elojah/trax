package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListInvitation(ctx context.Context, req *dto.ListInvitationReq) (*dto.ListInvitationResp, error) {
	logger := log.With().Str("method", "list_invitation").Logger()

	if req == nil {
		return &dto.ListInvitationResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.ListInvitationResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var ids []ulid.ID
	var groupIDs []ulid.ID

	if req.OwnGroup {
		groupIDs = claims.GroupIDs(user.Requirement{Resource: user.R_user, Command: user.C_read})
	} else if len(req.GroupIDs) > 0 {
		if err := claims.Require(user.NewRequirements(req.GroupIDs, user.R_user, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListInvitationResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		groupIDs = req.GroupIDs
	} else {
		err := gerrors.ErrMissingAtLeast{
			AtLeast: 1,
			Fields:  []string{"own_group", "group_ids"},
		}

		logger.Error().Err(err).Msg("invalid argument")

		return &dto.ListInvitationResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if len(req.IDs) == 0 {
		ids = req.IDs
	}

	var invitations []user.Invitation
	var total uint64

	// Fetch invitations
	if err := h.user.Tx(ctx, transaction.Read, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		invitationsByGroup, totalByGroup, err := h.user.ListInvitationByGroup(ctx, user.FilterInvitation{
			IDs:      ids,
			GroupIDs: groupIDs,
			Paginate: req.Paginate,
			Search:   req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list invitations")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		total = totalByGroup

		invitations = make([]user.Invitation, 0)
		for _, i := range invitationsByGroup {
			invitations = append(invitations, i...)
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.ListInvitationResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListInvitationResp{
		Invitations: invitations,
		Total:       total,
	}, nil
}
