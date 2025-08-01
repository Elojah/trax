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
			FirstName:    req.Firstname,
			LastName:     req.Lastname,
			GoogleID:     "",
			CreatedAt:    now,
			UpdatedAt:    now,
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

		logger.Info().Msg("created")

		// Assign user roles if invitation exists for this email
		invitation, err := h.user.FetchInvitation(ctx, user.FilterInvitation{
			Email: &u.Email,
		})
		if err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Info().Msg("no invitation found for user")

				return transaction.Commit, nil
			} else {
				logger.Error().Err(err).Msg("failed to fetch invitation")

				return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
			}
		}

		iroles, err := h.user.ListInvitationRole(ctx, user.FilterInvitationRole{
			InvitationID: invitation.ID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list invitation roles")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(iroles) == 0 {
			logger.Info().Msg("no roles to assign from invitation")

			return transaction.Commit, nil
		}

		roleUsers := make([]user.RoleUser, 0, len(iroles))
		for _, irole := range iroles {
			roleUsers = append(roleUsers, user.RoleUser{
				UserID:    u.ID,
				RoleID:    irole.RoleID,
				CreatedAt: now,
				UpdatedAt: now,
			})
		}
		if err := h.user.InsertBatchRoleUser(ctx, roleUsers...); err != nil {
			logger.Error().Err(err).Msg("failed to insert role users")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &pbtypes.Empty{}, err
	}

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
