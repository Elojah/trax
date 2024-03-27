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

func (h *handler) UpdateProfile(ctx context.Context, req *dto.UpdateProfileReq) (*pbtypes.Empty, error) {
	logger := log.With().Str("method", "update_profile").Logger()

	if req == nil {
		return &pbtypes.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &pbtypes.Empty{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	p := user.PatchProfile{
		FirstName: func(s string) *string {
			if len(s) == 0 {
				return nil
			}
			return &s
		}(req.Firstname),
		LastName: func(s string) *string {
			if len(s) == 0 {
				return nil
			}
			return &s
		}(req.Lastname),
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.UpdateProfile(ctx, user.FilterProfile{
			UserID: u.ID,
		}, p); err != nil {
			logger.Error().Err(err).Msg("failed to update user profile")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &pbtypes.Empty{}, err
	}

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
