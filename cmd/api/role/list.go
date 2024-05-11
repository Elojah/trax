package role

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

func (h *HandlerRole) ListRole(ctx context.Context, req *dto.ListRoleReq) (*dto.ListRoleResp, error) {
	logger := log.With().Str("method", "list_role").Logger()

	if req == nil {
		return &dto.ListRoleResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.User.Auth(ctx, "access")
	if err != nil {
		return &dto.ListRoleResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var entityIDs []ulid.ID
	var ids []ulid.ID

	if req.UserIDs {
		ids = claims.RoleIDs()
	} else if len(req.IDs) > 0 {
		ids = req.IDs
	}

	if req.UserEntityIDs {
		entityIDs = claims.EntityIDs(user.Requirement{Resource: user.R_role, Command: user.C_read})
	} else if len(req.EntityIDs) > 0 {
		if err := claims.Require(user.NewRequirements(req.EntityIDs, user.R_role, user.C_read)...); err != nil {
			logger.Error().Err(err).Msg("permission denied")

			return &dto.ListRoleResp{}, status.New(codes.PermissionDenied, err.Error()).Err()
		}

		entityIDs = req.EntityIDs
	} else if !req.UserIDs {
		err := gerrors.ErrMissingAtLeast{
			AtLeast: 1,
			Fields:  []string{"user_ids", "user_entity_ids", "entity_ids"},
		}

		logger.Error().Err(err).Msg("invalid argument")

		return &dto.ListRoleResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	var roles []user.Role
	var total uint64

	if err := h.User.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		roles, total, err = h.User.ListRole(ctx, user.FilterRole{
			IDs:       ids,
			EntityIDs: entityIDs,
			Paginate:  req.Paginate,
			Search:    req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list role")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		return transaction.Commit, nil
	}); err != nil {
		logger.Error().Err(err).Msg("failed transaction")

		return &dto.ListRoleResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListRoleResp{
		Roles: roles,
		Total: total,
	}, nil
}