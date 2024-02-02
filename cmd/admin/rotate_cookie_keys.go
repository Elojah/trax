package main

import (
	"context"

	"github.com/elojah/trax/pkg/cookie"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) RotateCookieKeys(ctx context.Context, req *pbtypes.Empty) (*pbtypes.Empty, error) {
	logger := log.With().Str("method", "rotate_cookie_keys").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &pbtypes.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	if err := h.cookie.RotateKeys(ctx, cookie.NewKeys(), cookie.NKeys); err != nil {
		logger.Error().Err(err).Msg("failed to rotate cookie keys")

		return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
