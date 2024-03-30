package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) SigninGoogle(ctx context.Context, req *pbtypes.String) (*dto.SigninResp, error) {
	logger := log.With().Str("method", "signin_google").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &dto.SigninResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Validate token
	gid, claims, err := h.google.Signin(ctx, req.Value)
	if err != nil {
		logger.Error().Err(err).Msg("failed to validate token")

		return &dto.SigninResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	var u user.U
	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		logger = logger.With().Str("google_id", gid).Logger()

		// #Check if user exist
		u, err = h.user.Fetch(ctx, user.Filter{GoogleID: &gid})
		if err != nil {
			if !errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("failed to fetch user")

				return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
			}
			// user not found, create
		} else {
			logger.Info().Msg("found")

			return transaction.Commit, nil
		}

		var p user.Profile
		u, p, err = claims.CreateUser(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("failed to create user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		logger = logger.With().Str("user_id", u.ID.String()).Logger()

		if err := h.user.Insert(ctx, u); err != nil {
			logger.Error().Err(err).Msg("failed to insert user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if err := h.user.InsertProfile(ctx, p); err != nil {
			if errors.As(err, &gerrors.ErrConflict{}) {
				logger.Error().Err(err).Msg("failed to insert profile")

				return transaction.Rollback, status.New(codes.AlreadyExists, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to insert profile")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		logger.Info().Msg("created")

		return transaction.Commit, nil
	}); err != nil {
		return &dto.SigninResp{}, err
	}

	// #Create JWT
	jwt, err := h.user.CreateJWT(ctx, u, "access", time.Hour)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create JWT")

		return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Create refresh token
	rt, err := h.user.CreateJWT(ctx, u, "refresh", 24*time.Hour)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create JWT")

		return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.SigninResp{
		AccessToken:  jwt,
		RefreshToken: rt,
	}, nil
}
