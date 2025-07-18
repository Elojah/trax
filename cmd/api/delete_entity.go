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

func (h *handler) DeleteEntity(ctx context.Context, req *dto.DeleteEntityReq) (*user.Entity, error) {
	logger := log.With().Str("method", "delete_entity").Logger()

	if req == nil {
		return &user.Entity{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Entity{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{
		EntityID: req.ID,
		Resource: user.R_entity,
		Command:  user.C_write,
	}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &user.Entity{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	var e user.Entity

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		entities, err := h.user.DeleteEntity(ctx, user.FilterEntity{
			ID: req.ID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to delete entity")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(entities) == 0 {
			err := gerrors.ErrNotFound{Resource: "entity", Index: req.ID.String()}
			logger.Error().Err(err).Msg("failed to delete entity")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		e = entities[0]

		return transaction.Commit, nil
	}); err != nil {
		return &user.Entity{}, err
	}

	logger.Info().Msg("success")

	return &e, nil
}
