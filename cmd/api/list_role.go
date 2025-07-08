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

func (h *handler) ListRole(ctx context.Context, req *dto.ListRoleReq) (*dto.ListRoleResp, error) {
	logger := log.With().Str("method", "list_role").Logger()

	if req == nil {
		return &dto.ListRoleResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.ListRoleResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var groupIDs []ulid.ID
	var ids []ulid.ID

	// Check valid group first
	if req.OwnGroup || req.Own {
		groupIDs = claims.GroupIDs(user.Requirement{Resource: user.R_role, Command: user.C_read})
	} else if len(req.GroupIDs) > 0 {
		if err := claims.Require(user.NewRequirements(req.GroupIDs, user.R_role, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListRoleResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		groupIDs = req.GroupIDs
	} else {
		err := gerrors.ErrMissingAtLeast{
			AtLeast: 1,
			Fields:  []string{"own", "own_group", "group_ids"},
		}

		logger.Error().Err(err).Msg("invalid argument")

		return &dto.ListRoleResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	// Role IDs
	if req.Own {
		ids = claims.RoleIDs()
	} else if len(req.IDs) > 0 {
		ids = req.IDs
	}

	var result []dto.RolePermission
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		roles, totalRole, err := h.user.ListRole(ctx, user.FilterRole{
			IDs:      ids,
			GroupIDs: groupIDs,
			UserID:   req.UserID,
			Paginate: req.Paginate,
			Search:   req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		// Populate roles with their permissions
		permissions, err := h.user.ListPermission(ctx, user.FilterPermission{
			RoleIDs: user.Roles(roles).IDs(),
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list permission")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		permissionsByRole := user.Permissions(permissions).ByRole()

		for _, r := range roles {
			ps, ok := permissionsByRole[r.ID.String()]
			if !ok {
				ps = []user.Permission{}
			}

			result = append(result, dto.RolePermission{
				Role:        r,
				Permissions: ps,
			})
		}

		total = totalRole

		return transaction.Commit, nil
	}); err != nil {
		logger.Error().Err(err).Msg("failed transaction")

		return &dto.ListRoleResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListRoleResp{
		Roles: result,
		Total: total,
	}, nil
}
