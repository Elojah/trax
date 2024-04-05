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

	// #MARK:Create user
	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		now := time.Now().Unix()
		hash, salt := user.Encrypt(req.Password)
		u := user.U{
			ID:           ulid.NewID(),
			Email:        req.Email,
			PasswordHash: hash,
			PasswordSalt: salt,
			GoogleID:     "",
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		p := user.Profile{
			UserID:    u.ID,
			FirstName: req.Firstname,
			LastName:  req.Lastname,
			CreatedAt: now,
			UpdatedAt: now,
		}

		logger = logger.With().Str("user_id", u.ID.String()).Str("user_email", u.Email).Logger()

		if err := h.user.Insert(ctx, u); err != nil {
			if errors.As(err, &gerrors.ErrConflict{}) {
				logger.Error().Err(err).Msg("failed to insert user")

				return transaction.Rollback, status.New(codes.AlreadyExists, err.Error()).Err()
			}
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
		return &pbtypes.Empty{}, err
	}

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
