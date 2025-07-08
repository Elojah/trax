package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) DeleteGroup(ctx context.Context, req *dto.DeleteGroupReq) (*user.Group, error) {
	logger := log.With().Str("method", "delete_group").Logger()

	if req == nil {
		return &user.Group{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Group{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{
		GroupID:  req.ID,
		Resource: user.R_group,
		Command:  user.C_write,
	}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &user.Group{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	var g user.Group

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		groups, err := h.user.DeleteGroup(ctx, user.FilterGroup{
			ID: req.ID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to delete group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(groups) == 0 {
			err := gerrors.ErrNotFound{Resource: "group", Index: req.ID.String()}
			logger.Error().Err(err).Msg("failed to delete group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		g = groups[0]

		return transaction.Commit, nil
	}); err != nil {
		return &user.Group{}, err
	}

	logger.Info().Msg("success")

	return &g, nil
}
