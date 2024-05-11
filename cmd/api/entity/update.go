package entity

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

func (h *HandlerEntity) UpdateEntity(ctx context.Context, req *dto.UpdateEntityReq) (*user.Entity, error) {
	logger := log.With().Str("method", "update_entity").Logger()

	if req == nil {
		return &user.Entity{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	claims, err := h.User.Auth(ctx, "access")
	if err != nil {
		return &user.Entity{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := claims.Require(user.Requirement{EntityID: req.ID, Resource: user.R_entity, Command: user.C_update}); err != nil {
		return &user.Entity{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	var e user.Entity

	if err := h.User.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		now := time.Now().Unix()
		es, err := h.User.UpdateEntity(ctx, user.FilterEntity{
			ID: req.ID,
		}, user.PatchEntity{
			Name:        pbtypes.GetString(req.Name),
			Description: pbtypes.GetString(req.Description),
			AvatarURL:   pbtypes.GetString(req.AvatarURL),
			UpdatedAt:   &now,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to update entity")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(es) == 0 {
			err := gerrors.ErrNotFound{Resource: "entity", Index: req.ID.String()}
			logger.Error().Err(err).Msg("failed to update entity")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		e = es[0]

		return transaction.Commit, nil
	}); err != nil {
		return &user.Entity{}, err
	}

	logger.Info().Msg("success")

	return &e, nil
}
