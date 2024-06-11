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

func (h *handler) DeleteRoleUser(ctx context.Context, req *dto.DeleteRoleUserReq) (*dto.UserRoles, error) {
	logger := log.With().Str("method", "delete_role_user").Logger()

	if req == nil {
		return &dto.UserRoles{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.UserRoles{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var role user.Role
	var permissions user.Permissions

	// Fetch role and permissions
	if err := h.user.Tx(ctx, transaction.Read, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		role, err = h.user.FetchRole(ctx, user.FilterRole{
			ID: req.RoleID,
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
			RoleID: req.RoleID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list role permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.UserRoles{}, err
	}

	// check permissions
	if err := claims.Require(user.Requirement{EntityID: role.EntityID, Resource: user.R_user, Command: user.C_update}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &dto.UserRoles{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}
	// current user need to have all assigned permissions to delete a role
	for _, perm := range permissions {
		if err := claims.Require(user.Requirement{EntityID: role.EntityID, Resource: perm.Resource, Command: perm.Command}); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.UserRoles{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	var u user.U
	var roles []user.Role

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		rolesByUser, _, err := h.user.ListRoleByUser(ctx, user.FilterRole{
			UserID:   req.UserID,
			EntityID: role.EntityID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list roles")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// user needs AT LEAST ONE ROLE in current entity to be assigned a new role
		// if user has no role in current entity, it means he has no access to this entity
		// if you want to assign a role to a user, you need to invite him with a role in the entity first
		var ok bool

		roles, ok = rolesByUser[req.UserID.String()]
		if !ok || len(roles) == 0 {
			err := gerrors.ErrNotFound{Resource: "user", Index: req.UserID.String()}
			logger.Error().Err(err).Msg("user not found")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		index := -1
		found := func() bool {
			for i, r := range roles {
				if r.ID.Compare(req.RoleID) == 0 {
					index = i
					return true
				}
			}

			return false
		}()
		if !found {
			err := gerrors.ErrNotFound{Resource: "role", Index: req.RoleID.String()}
			logger.Error().Err(err).Msg("role not found")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		if err = h.user.DeleteRoleUser(ctx, user.FilterRoleUser{
			RoleID: req.RoleID,
			UserID: req.UserID,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete role user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// remove role from roles slice
		roles = append(roles[:index], roles[index+1:]...)

		// To build response
		u, err = h.user.Fetch(ctx, user.Filter{
			ID: req.UserID,
		})
		if err != nil {
			if !errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("user not found")

				return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
			}
			logger.Error().Err(err).Msg("failed to fetch role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.UserRoles{}, err
	}

	logger.Info().Msg("success")

	return &dto.UserRoles{
		User:  u,
		Roles: roles,
	}, nil
}
