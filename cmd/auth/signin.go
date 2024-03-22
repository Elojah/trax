package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Signin(ctx context.Context, req *dto.SigninReq) (*dto.SigninResp, error) {
	logger := log.With().Str("method", "signin").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &dto.SigninResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// Fetch user
	var u user.U

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		u, err = h.user.Fetch(ctx, user.Filter{
			Email: &req.Email,
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
		return &dto.SigninResp{}, err
	}

	// Check user password
	if err := u.Check(req.Password); err != nil {
		logger.Error().Err(err).Msg("failed to check user")

		if errors.As(err, &user.ErrInvalidPassword{}) {
			logger.Error().Err(err).Msg("invalid password")

			return &dto.SigninResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
		}
		logger.Error().Err(err).Msg("failed to check user")

		return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Create JWT
	jwt, err := h.user.CreateJWT(ctx, u, "access", time.Minute) // TODO: up to 1 hour (configurable)
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
