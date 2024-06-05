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

func (h *handler) ListUser(ctx context.Context, req *dto.ListUserReq) (*dto.ListUserResp, error) {
	logger := log.With().Str("method", "list_user").Logger()

	if req == nil {
		return &dto.ListUserResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.ListUserResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var ids []ulid.ID
	var entityIDs []ulid.ID

	if req.OwnEntity {
		entityIDs = claims.EntityIDs(user.Requirement{Resource: user.R_user, Command: user.C_read})
	} else if len(req.EntityIDs) > 0 {
		if err := claims.Require(user.NewRequirements(req.EntityIDs, user.R_user, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListUserResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		entityIDs = req.EntityIDs
	} else {
		err := gerrors.ErrMissingAtLeast{
			AtLeast: 1,
			Fields:  []string{"user_entity_ids", "entity_ids"},
		}

		logger.Error().Err(err).Msg("invalid argument")

		return &dto.ListUserResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if len(req.IDs) == 0 {
		ids = req.IDs
	}

	var users []user.U
	var roles map[string][]user.Role
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		users, total, err = h.user.ListByEntity(ctx, user.Filter{
			IDs:       ids,
			EntityIDs: entityIDs,
			Paginate:  req.Paginate,
			Search:    req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		roles, _, err = h.user.ListRoleByUser(ctx, user.FilterRole{
			UserIDs:   ids,
			EntityIDs: entityIDs,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list roles")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		logger.Error().Err(err).Msg("failed transaction")

		return &dto.ListUserResp{}, err
	}

	resp := &dto.ListUserResp{
		Users: make([]dto.UserRoles, 0, len(users)),
		Total: total,
	}

	for _, u := range users {
		rs, ok := roles[u.ID.String()]
		if !ok {
			// should never happen, both queries should never return a user without roles
			continue
		}
		resp.Users = append(resp.Users, dto.UserRoles{
			User:  u,
			Roles: rs,
		})
	}

	logger.Info().Msg("success")

	return resp, nil
}
