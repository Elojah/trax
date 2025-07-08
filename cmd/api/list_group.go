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

func (h *handler) ListGroup(ctx context.Context, req *dto.ListGroupReq) (*dto.ListGroupResp, error) {
	logger := log.With().Str("method", "list_group").Logger()

	if req == nil {
		return &dto.ListGroupResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.ListGroupResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var ids []ulid.ID

	if req.Own {
		ids = claims.GroupIDs()
	} else if req.IDs != nil {
		if err := claims.Require(user.NewRequirements(req.IDs, user.R_group, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListGroupResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		ids = req.IDs
	} else {
		err := gerrors.ErrMissingAtLeast{AtLeast: 1, Fields: []string{"user_all", "ids"}}

		logger.Error().Err(err).Msg("invalid argument")

		return &dto.ListGroupResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	var groups []user.Group
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		groups, total, err = h.user.ListGroup(ctx, user.FilterGroup{
			IDs:      ids,
			Paginate: req.Paginate,
			Search:   req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		logger.Error().Err(err).Msg("failed transaction")

		return &dto.ListGroupResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListGroupResp{
		Groups: groups,
		Total:  total,
	}, nil
}
