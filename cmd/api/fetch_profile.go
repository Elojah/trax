package main

import (
	"context"
	"errors"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) FetchProfile(ctx context.Context, req *dto.FetchProfileReq) (*user.Profile, error) {
	logger := log.With().Str("method", "fetch_profile").Logger()

	if req == nil {
		return &user.Profile{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Profile{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var userID ulid.ID

	if req.Me {
		userID = claims.UserID
	} else {
		if err := claims.Require(req.UserID, user.R_user, user.C_read); err != nil {
			logger.Error().Err(err).Msg("failed to require permission")

			return &user.Profile{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		// TODO: Check if user has at least one role in this entity

		userID = req.UserID
	}

	var profile user.Profile

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		profile, err = h.user.FetchProfile(ctx, user.FilterProfile{
			UserID: userID,
		})
		if err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("failed to fetch user profile")

				return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to fetch user profile")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.Profile{}, err
	}

	logger.Info().Msg("success")

	return &profile, nil
}
