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

func (h *handler) DeleteInvitation(ctx context.Context, req *dto.DeleteInvitationReq) (*user.Invitation, error) {
	logger := log.With().Str("method", "delete_invitation").Logger()

	if req == nil {
		return &user.Invitation{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Invitation{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{
		GroupID:  req.GroupID,
		Resource: user.R_user,
		Command:  user.C_write,
	}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &user.Invitation{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err := h.user.DeleteInvitationRoleByGroup(ctx, user.FilterInvitationRole{
			InvitationID: req.ID,
			GroupIDs:     []ulid.ID{req.GroupID},
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete invitation roles")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		counts, err := h.user.CountInvitationRoleByInvitation(ctx, user.FilterInvitationRole{
			InvitationID: req.ID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to count invitation roles")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if count, ok := counts[req.ID.String()]; ok && count == 0 {
			if err := h.user.DeleteInvitation(ctx, user.FilterInvitation{
				ID: req.ID,
			}); err != nil {
				logger.Error().Err(err).Msg("failed to delete invitation")

				return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
			}
		}

		return transaction.Commit, nil
	}); err != nil {
		logger.Error().Err(err).Msg("transaction failed")

		return &user.Invitation{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("invitation deleted successfully")

	return &user.Invitation{
		ID: req.ID,
	}, nil
}
