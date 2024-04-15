package main

import (
	"context"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateEntity(ctx context.Context, req *dto.CreateEntityReq) (*user.Entity, error) {
	logger := log.With().Str("method", "create_entity").Logger()

	if req == nil {
		return &user.Entity{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Entity{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	now := time.Now().Unix()
	e := user.Entity{
		ID:          ulid.NewID(),
		Name:        req.Name,
		Description: req.Description.Value,
		AvatarURL:   req.AvatarURL.Value,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// #MARK:Check request
	if err := e.Check(); err != nil {
		return &user.Entity{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.InsertEntity(ctx, e); err != nil {
			logger.Error().Err(err).Msg("failed to insert entity")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// insert default admin role with all permissions
		role := user.Role{
			ID:        ulid.NewID(),
			EntityID:  e.ID,
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
			UserID:    claims.UserID,
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

	return &user.Entity{
		ID:   e.ID,
		Name: e.Name,
	}, nil
}
