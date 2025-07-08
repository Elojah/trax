package main

import (
	"context"
	"errors"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) FetchGroup(ctx context.Context, req *dto.FetchGroupReq) (*user.Group, error) {
	logger := log.With().Str("method", "fetch_group").Logger()

	if req == nil {
		return &user.Group{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Group{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{
		GroupID:  req.ID,
		Resource: user.R_group,
		Command:  user.C_read,
	}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &user.Group{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	var group user.Group

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		group, err = h.user.FetchGroup(ctx, user.FilterGroup{
			ID: req.ID,
		})
		if err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("failed to fetch user group")

				return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to fetch user group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.Group{}, err
	}

	logger.Info().Msg("success")

	return &group, nil
}
