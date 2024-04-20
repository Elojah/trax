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

func (h *handler) ListEntity(ctx context.Context, req *dto.ListEntityReq) (*dto.ListEntityResp, error) {
	logger := log.With().Str("method", "list_entity").Logger()

	if req == nil {
		return &dto.ListEntityResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.ListEntityResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var entityIDs []ulid.ID

	if req.IDs != nil {
		if err := claims.Require(user.NewRequirements(req.IDs, user.R_entity, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListEntityResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		entityIDs = req.IDs
	} else {
		entityIDs = claims.EntityIDs()
	}

	var entities []user.Entity
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		entities, total, err = h.user.ListEntity(ctx, user.FilterEntity{
			IDs:      entityIDs,
			Paginate: req.Paginate,
			Search:   req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list entity")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &dto.ListEntityResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListEntityResp{
		Entities: entities,
		Total:    total,
	}, nil
}
