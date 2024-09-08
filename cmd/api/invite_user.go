package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) InviteUser(ctx context.Context, req *dto.InviteUserReq) (*user.U, error) {
	logger := log.With().Str("method", "invite_user").Logger()

	if req == nil {
		return &user.U{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.U{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var roles []user.Role
	var permissions user.Permissions

	// Fetch role and permissions
	if err := h.user.Tx(ctx, transaction.Read, func(ctx context.Context) (transaction.Operation, error) {
		var err error

		roles, _, err = h.user.ListRole(ctx, user.FilterRole{
			IDs:      req.RoleIDs,
			EntityID: req.EntityID,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to fetch role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// overwrite roles ids with entity filter
		req.RoleIDs = make([]ulid.ID, 0, len(roles))
		for _, role := range roles {
			req.RoleIDs = append(req.RoleIDs, role.ID)
		}

		permissions, err = h.user.ListPermission(ctx, user.FilterPermission{
			RoleIDs: req.RoleIDs,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list role permissions")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.U{}, err
	}

	if len(roles) == 0 {
		err := gerrors.ErrNotFound{Resource: "role", Index: ""}
		logger.Error().Err(err).Msg("roles not found")

		return &user.U{}, status.New(codes.NotFound, err.Error()).Err()
	}

	// check permissions
	if err := claims.Require(user.Requirement{EntityID: req.EntityID, Resource: user.R_user, Command: user.C_create}); err != nil {
		logger.Error().Err(err).Msg("permission denied")

		return &user.U{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}
	// current user need to have all assigned permissions to assign a role
	for _, perm := range permissions {
		if err := claims.Require(user.Requirement{EntityID: req.EntityID, Resource: perm.Resource, Command: perm.Command}); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &user.U{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	// now := time.Now().Unix()
	// ru := user.RoleUser{
	// 	RoleID:    req.RoleID,
	// 	UserID:    req.UserID,
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }

	// var u user.U
	// var roles []user.Role

	// if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
	// 	rolesByUser, _, err := h.user.ListRoleByUser(ctx, user.FilterRole{
	// 		UserID:   req.UserID,
	// 		EntityID: role.EntityID,
	// 	})
	// 	if err != nil {
	// 		logger.Error().Err(err).Msg("failed to list roles")

	// 		return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
	// 	}

	// 	// user needs AT LEAST ONE ROLE in current entity to be assigned a new role
	// 	// if user has no role in current entity, it means he has no access to this entity
	// 	// if you want to assign a role to a user, you need to invite him with a role in the entity first
	// 	var ok bool

	// 	roles, ok = rolesByUser[req.UserID.String()]
	// 	if !ok || len(roles) == 0 {
	// 		err := gerrors.ErrNotFound{Resource: "user", Index: req.UserID.String()}
	// 		logger.Error().Err(err).Msg("user not found")

	// 		return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
	// 	}

	// 	if err = h.user.InsertRoleUser(ctx, ru); err != nil {
	// 		logger.Error().Err(err).Msg("failed to insert role user")

	// 		return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
	// 	}

	// 	// To build response
	// 	u, err = h.user.Fetch(ctx, user.Filter{
	// 		ID: req.UserID,
	// 	})
	// 	if err != nil {
	// 		if !errors.As(err, &gerrors.ErrNotFound{}) {
	// 			logger.Error().Err(err).Msg("user not found")

	// 			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
	// 		}
	// 		logger.Error().Err(err).Msg("failed to fetch role")

	// 		return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
	// 	}

	// 	return transaction.Commit, nil
	// }); err != nil {
	// 	return &user.U{}, err
	// }

	logger.Info().Msg("success")

	return &user.U{
		// User:  u,
		// Roles: append(roles, role),
	}, nil
}
