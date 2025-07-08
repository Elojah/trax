package main

import (
	"context"
	"errors"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) DeleteRole(ctx context.Context, req *dto.DeleteRoleReq) (*dto.RolePermission, error) {
	logger := log.With().Str("method", "delete_role").Logger()

	if req == nil {
		return &dto.RolePermission{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.RolePermission{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var role user.Role
	var permissions user.Permissions

	// Fetch role and permissions
	if err := h.user.Tx(ctx, transaction.Read, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		role, err = h.user.FetchRole(ctx, user.FilterRole{
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

		permissions, err = h.user.ListPermission(ctx, user.FilterPermission{
			RoleID: req.ID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list role permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.RolePermission{}, err
	}

	if role.Name == "admin" {
		err := user.ErrForbiddenAdminRole{}
		logger.Error().Err(err).Msg("permission denied on admin role")

		return &dto.RolePermission{}, status.New(codes.PermissionDenied, err.Error()).Err()

	}

	// check permissions
	if err := claims.Require(user.Requirement{GroupID: role.GroupID, Resource: user.R_role, Command: user.C_write}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &dto.RolePermission{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}
	// current user need to have all assigned permissions to delete a role
	for _, perm := range permissions {
		if err := claims.Require(user.Requirement{GroupID: role.GroupID, Resource: perm.Resource, Command: perm.Command}); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.RolePermission{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.DeleteRole(ctx, user.FilterRole{
			ID: req.ID,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.RolePermission{}, err
	}

	logger.Info().Msg("success")

	return &dto.RolePermission{
		Role:        role,
		Permissions: permissions,
	}, nil
}
