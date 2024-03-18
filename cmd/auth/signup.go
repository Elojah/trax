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
	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Signup(ctx context.Context, req *dto.SignupReq) (*pbtypes.Empty, error) {
	logger := log.With().Str("method", "signup").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &pbtypes.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Create user
	now := time.Now().Unix()
	u := user.U{
		ID:        ulid.NewID(),
		Email:     req.Email,
		Password:  req.Password,
		TwitchID:  "",
		GoogleID:  "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	var result error

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) error {
		if err := h.user.Insert(ctx, u); err != nil {
			if errors.As(err, &gerrors.ErrConflict{}) {
				logger.Error().Err(err).Msg("failed to insert user")
				result = status.New(codes.AlreadyExists, err.Error()).Err()

				return nil
			} else {
				logger.Error().Err(err).Msg("failed to insert user")
				result = status.New(codes.Internal, err.Error()).Err()

				return nil
			}
		}

		return nil
	}); err != nil {
		return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
	} else if result != nil {
		return &pbtypes.Empty{}, result
	}

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
