package role

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

func (h *HandlerRole) CreateRole(ctx context.Context, req *dto.CreateRoleReq) (*user.Role, error) {
	logger := log.With().Str("method", "create_role").Logger()

	if req == nil {
		return &user.Role{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.User.Auth(ctx, "access")
	if err != nil {
		return &user.Role{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// TODO: check if user has permission to create role
	_ = claims

	now := time.Now().Unix()
	e := user.Role{
		ID:        ulid.NewID(),
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := h.User.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		if err = h.User.InsertRole(ctx, e); err != nil {
			logger.Error().Err(err).Msg("failed to insert role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		return &user.Role{}, err
	}

	logger.Info().Msg("success")

	return &user.Role{
		ID:   e.ID,
		Name: e.Name,
	}, nil
}
