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

	var entityIDs []ulid.ID
	if req.EntityIDs != nil {
		if err := claims.Require(user.NewRequirements(req.EntityIDs, user.R_user, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListUserResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		entityIDs = req.EntityIDs
	} else {
		entityIDs = claims.EntityIDs(user.Requirement{Resource: user.R_user, Command: user.C_read})
	}

	var users []user.U
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		users, total, err = h.user.ListByEntity(ctx, user.Filter{
			IDs:       req.IDs,
			EntityIDs: entityIDs,
			Paginate:  req.Paginate,
			Search:    req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.ListUserResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListUserResp{
		Users: users,
		Total: total,
	}, nil
}
