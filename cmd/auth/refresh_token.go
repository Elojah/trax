package main

import (
	"context"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) RefreshToken(ctx context.Context, req *pbtypes.String) (*dto.SigninResp, error) {
	logger := log.With().Str("method", "refresh_token").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &dto.SigninResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "refresh")
	if err != nil {
		return &dto.SigninResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #MARK:Create JWT
	jwt, err := h.user.CreateJWT(ctx, user.U{ID: claims.UserID}, "access", time.Hour)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create JWT")

		return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #MARK:Create refresh token
	rt, err := h.user.CreateJWT(ctx, user.U{ID: claims.UserID}, "refresh", 24*time.Hour)
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
