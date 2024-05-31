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

func (h *handler) UpdateRole(ctx context.Context, req *dto.UpdateRoleReq) (*user.Role, error) {
	logger := log.With().Str("method", "update_role").Logger()

	if req == nil {
		return &user.Role{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Role{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{EntityID: req.RoleID, Resource: user.R_role, Command: user.C_create}); err != nil {
		return &user.Role{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	// check permissions
	// current user need to have all assigned permissions to create a role
	for _, perm := range req.Permissions {
		if err := claims.Require(user.Requirement{EntityID: req.RoleID, Resource: perm.Resource, Command: perm.Command}); err != nil {
			return &user.Role{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	now := time.Now().Unix()
	role := user.Role{
		ID:        ulid.NewID(),
		Name:      "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	permissions := make([]user.Permission, 0, len(req.Permissions))
	for _, perm := range req.Permissions {
		permissions = append(permissions, user.Permission{
			RoleID:    role.ID,
			Resource:  perm.Resource,
			Command:   perm.Command,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.user.InsertRole(ctx, role); err != nil {
			logger.Error().Err(err).Msg("failed to insert role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if err = h.user.InsertPermissions(ctx, permissions); err != nil {
			logger.Error().Err(err).Msg("failed to insert role permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.Role{}, err
	}

	logger.Info().Msg("success")

	return &user.Role{
		ID:   role.ID,
		Name: role.Name,
	}, nil
}
