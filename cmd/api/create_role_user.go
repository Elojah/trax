package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/paginate"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateRoleUser(ctx context.Context, req *dto.CreateRoleUserReq) (*dto.RoleUserResp, error) {
	logger := log.With().Str("method", "create_role_user").Logger()

	if req == nil {
		return &dto.RoleUserResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.RoleUserResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
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
		return &dto.RoleUserResp{}, err
	}

	// #MARK:Check permissions
	if err := claims.Require(user.Requirement{GroupID: role.GroupID, Resource: user.R_user, Command: user.C_write}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &dto.RoleUserResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	for _, perm := range permissions {
		if err := claims.Require(user.Requirement{GroupID: role.GroupID, Resource: perm.Resource, Command: perm.Command}); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.RoleUserResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	now := time.Now().Unix()
	ru := user.RoleUser{
		RoleID:    req.RoleID,
		UserID:    req.UserID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	var u user.U

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		_, total, err := h.user.ListRole(ctx, user.FilterRole{
			UserID:  req.UserID,
			GroupID: role.GroupID,
			Paginate: &paginate.Paginate{
				Start: 0,
				End:   1,
			},
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list roles")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// user needs AT LEAST ONE ROLE in current group to be assigned a new role
		// if user has no role in current group, it means he has no access to this group
		// if you want to assign a role to a user, you need to invite him with a role in the group first
		if total == 0 {
			err := gerrors.ErrNotFound{Resource: "user", Index: req.UserID.String()}
			logger.Error().Err(err).Msg("user not found")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		if err = h.user.InsertRoleUser(ctx, ru); err != nil {
			logger.Error().Err(err).Msg("failed to insert role user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

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
		return &dto.RoleUserResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.RoleUserResp{
		User: u,
		Role: dto.RolePermission{
			Role:        role,
			Permissions: permissions,
		},
	}, nil
}
