package main

import (
	"context"
	"time"

	"github.com/elojah/trax/internal/user"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateEntity(ctx context.Context, req *user.Entity) (*user.Entity, error) {
	logger := log.With().Str("method", "create_entity").Logger()

	if req == nil {
		return &user.Entity{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Entity{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #MARK:Check request
	if err := req.Check(); err != nil {
		return &user.Entity{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	now := time.Now().Unix()
	req.ID = ulid.NewID()
	req.CreatedAt = now
	req.UpdatedAt = now

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.InsertEntity(ctx, *req); err != nil {
			logger.Error().Err(err).Msg("failed to insert entity")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// insert default admin role with all permissions
		role := user.Role{
			ID:        ulid.NewID(),
			EntityID:  req.ID,
			Name:      "admin",
			UpdatedAt: now,
			CreatedAt: now,
		}
		if err := h.user.InsertRole(ctx, role); err != nil {
			logger.Error().Err(err).Msg("failed to insert role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if err := h.user.InsertPermissions(ctx, user.AllPermissions(role.ID)); err != nil {
			logger.Error().Err(err).Msg("failed to insert permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if err := h.user.InsertRoleUser(ctx, user.RoleUser{
			RoleID:    role.ID,
			UserID:    u.ID,
			CreatedAt: now,
			UpdatedAt: now,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to insert role user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.Entity{}, err
	}

	logger.Info().Msg("success")

	return &user.Entity{}, nil
}
