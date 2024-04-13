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

func (h *handler) FetchEntityProfile(ctx context.Context, req *dto.FetchEntityProfileReq) (*user.EntityProfile, error) {
	logger := log.With().Str("method", "fetch_entity_profile").Logger()

	if req == nil {
		return &user.EntityProfile{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.EntityProfile{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(req.EntityID, user.R_entity, user.C_read); err != nil {
		logger.Error().Err(err).Msg("failed to require permission")

		return &user.EntityProfile{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	var profile user.EntityProfile

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		profile, err = h.user.FetchEntityProfile(ctx, user.FilterEntityProfile{
			EntityID: req.EntityID,
		})
		if err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("failed to fetch user entity profile")

				return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to fetch user entity profile")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.EntityProfile{}, err
	}

	logger.Info().Msg("success")

	return &profile, nil
}
