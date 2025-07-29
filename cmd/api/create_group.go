package main

import (
	"context"
	"errors"
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

func (h *handler) CreateGroup(ctx context.Context, req *dto.CreateGroupReq) (*user.Group, error) {
	logger := log.With().Str("method", "create_group").Logger()

	if req == nil {
		return &user.Group{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Group{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	now := time.Now().Unix()
	g := user.Group{
		ID:          ulid.NewID(),
		Name:        req.Name,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// #MARK:Check request
	if err := g.Check(); err != nil {
		return &user.Group{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.InsertGroup(ctx, g); err != nil {
			if errors.As(err, &gerrors.ErrConflict{}) {
				logger.Error().Err(err).Msg("failed to insert group")

				return transaction.Rollback, status.New(codes.AlreadyExists, err.Error()).Err()
			}

			logger.Error().Err(err).Msg("failed to insert group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// insert default admin role with all permissions
		role := user.Role{
			ID:        ulid.NewID(),
			GroupID:   g.ID,
			Name:      "admin",
			UpdatedAt: now,
			CreatedAt: now,
		}
		if err := h.user.InsertRole(ctx, role); err != nil {
			logger.Error().Err(err).Msg("failed to insert role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if err := h.user.InsertBatchPermission(ctx, user.AllPermissions(role.ID)...); err != nil {
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
		return &user.Group{}, err
	}

	logger.Info().Msg("success")

	return &user.Group{
		ID:   g.ID,
		Name: g.Name,
	}, nil
}
