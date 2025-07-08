package main

import (
	"context"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateGroup(ctx context.Context, req *dto.UpdateGroupReq) (*user.Group, error) {
	logger := log.With().Str("method", "update_group").Logger()

	if req == nil {
		return &user.Group{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	claims, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &user.Group{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{GroupID: req.ID, Resource: user.R_group, Command: user.C_edit}); err != nil {
		return &user.Group{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	var g user.Group

	// #MARK:Check request
	if req.Name != nil {
		g.Name = req.Name.Value
		if err := g.Check(); err != nil {
			return &user.Group{}, status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		now := time.Now().Unix()
		es, err := h.user.UpdateGroup(ctx, user.FilterGroup{
			ID: req.ID,
		}, user.PatchGroup{
			Name:        pbtypes.GetString(req.Name),
			Description: pbtypes.GetString(req.Description),
			AvatarURL:   pbtypes.GetString(req.AvatarURL),
			UpdatedAt:   &now,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to update group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(es) == 0 {
			err := gerrors.ErrNotFound{Resource: "group", Index: req.ID.String()}
			logger.Error().Err(err).Msg("failed to update group")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		g = es[0]

		return transaction.Commit, nil
	}); err != nil {
		return &user.Group{}, err
	}

	logger.Info().Msg("success")

	return &g, nil
}
