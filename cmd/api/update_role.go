package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateRole(ctx context.Context, req *dto.UpdateRoleReq) (*dto.RolePermission, error) {
	logger := log.With().Str("method", "update_role").Logger()

	if req == nil {
		return &dto.RolePermission{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.RolePermission{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var r user.Role

	// Fetch entity from role
	if err := h.user.Tx(ctx, transaction.Read, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		r, err = h.user.FetchRole(ctx, user.FilterRole{
			ID: req.ID,
		})
		if err != nil {
			if !errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("role not found")

				return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to fetch role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.RolePermission{}, err
	}

	if err := claims.Require(user.Requirement{EntityID: r.EntityID, Resource: user.R_role, Command: user.C_update}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &dto.RolePermission{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	// check permissions
	// current user need to have all assigned permissions to create a role
	for _, perm := range req.Permissions {
		if err := claims.Require(user.Requirement{EntityID: r.EntityID, Resource: perm.Resource, Command: perm.Command}); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.RolePermission{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	now := time.Now().Unix()

	permissions := make([]user.Permission, 0, len(req.Permissions))
	for _, perm := range req.Permissions {
		permissions = append(permissions, user.Permission{
			RoleID:    r.ID,
			Resource:  perm.Resource,
			Command:   perm.Command,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.DeletePermission(ctx, user.FilterPermission{
			RoleID: r.ID,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if err = h.user.InsertPermissions(ctx, permissions); err != nil {
			logger.Error().Err(err).Msg("failed to insert role permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		roles, err := h.user.UpdateRole(ctx, user.FilterRole{
			ID: req.ID,
		}, user.PatchRole{
			UpdatedAt: &now,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to update role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(roles) > 0 {
			r = roles[0]
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.RolePermission{}, err
	}

	logger.Info().Msg("success")

	return &dto.RolePermission{
		Role:        r,
		Permissions: permissions,
	}, nil
}
