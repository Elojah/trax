package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) FetchProfile(ctx context.Context, req *pbtypes.Empty) (*user.Profile, error) {
	logger := log.With().Str("method", "fetch_profile").Logger()

	if req == nil {
		return &user.Profile{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Profile{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var profile user.Profile

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		profile, err = h.user.FetchProfile(ctx, user.FilterProfile{
			UserID: u.ID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to fetch profile")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.Profile{}, err
	}

	logger.Info().Msg("success")

	return &profile, nil
}
