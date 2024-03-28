package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateProfile(ctx context.Context, req *dto.UpdateProfileReq) (*user.Profile, error) {
	logger := log.With().Str("method", "update_profile").Logger()

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
		profiles, err := h.user.UpdateProfile(ctx, user.FilterProfile{
			UserID: u.ID,
		}, user.PatchProfile{
			FirstName: pbtypes.GetString(req.Firstname),
			LastName:  pbtypes.GetString(req.Lastname),
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to update user profile")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(profiles) == 0 {
			err := gerrors.ErrNotFound{Resource: "profile", Index: u.ID.String()}
			logger.Error().Err(err).Msg("failed to update user profile")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		profile = profiles[0]

		return transaction.Commit, nil
	}); err != nil {
		return &user.Profile{}, err
	}

	logger.Info().Msg("success")

	return &profile, nil
}
