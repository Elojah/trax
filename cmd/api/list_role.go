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

	var roleIDs []ulid.ID
	if req.IDs != nil {
		if err := claims.RequireRoles(req.IDs...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListRoleResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		roleIDs = req.IDs
	} else {
		roleIDs = claims.RoleIDs()
	}

	var roles []user.Role
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		roles, total, err = h.user.ListRole(ctx, user.FilterRole{
			IDs:      roleIDs,
			Paginate: req.Paginate,
			Search:   req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.ListRoleResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListRoleResp{
		Roles: roles,
		Total: total,
	}, nil
}
