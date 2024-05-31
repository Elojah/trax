package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) FetchClaims(ctx context.Context, req *pbtypes.Empty) (*user.ClaimAuth, error) {
	logger := log.With().Str("method", "fetch_claims").Logger()

	if req == nil {
		return &user.ClaimAuth{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.ClaimAuth{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &claims.Auth, nil
}
