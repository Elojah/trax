package main

import (
	"context"
	"errors"
	"time"

	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/elojah/trax/pkg/user"
	"github.com/elojah/trax/pkg/user/dto"
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
	gid, err := h.google.Signin(ctx, req.Value)
	if err != nil {
		logger.Error().Err(err).Msg("failed to validate token")

		return &dto.SigninResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	// #Check if user exist
	u, err := h.user.Fetch(ctx, user.Filter{GoogleID: &gid})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			// #Create user
			u = user.U{
				ID:       ulid.NewID(),
				GoogleID: gid,
			}
			if err := h.user.Insert(ctx, u); err != nil {
				logger.Error().Err(err).Msg("failed to create user")

				return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
			}
		} else {
			logger.Error().Err(err).Msg("failed to fetch user")

			return &dto.SigninResp{}, status.New(codes.Internal, err.Error()).Err()
		}
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
