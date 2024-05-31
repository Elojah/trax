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

func (h *handler) FetchUser(ctx context.Context, req *dto.FetchUserReq) (*user.U, error) {
	logger := log.With().Str("method", "fetch_user").Logger()

	if req == nil {
		return &user.U{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.U{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var id ulid.ID

	if req.Me {
		id = claims.UserID
	} else {
		if err := claims.Require(user.Requirement{EntityID: req.EntityID, Resource: user.R_user, Command: user.C_read}); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &user.U{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		id = req.ID
	}

	var u user.U

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		u, err = h.user.Fetch(ctx, user.Filter{
			ID: id,
		})
		if err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("failed to fetch user")

				return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to fetch user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.U{}, err
	}

	logger.Info().Msg("success")

	return &u, nil
}
