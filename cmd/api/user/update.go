package user

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

func (h *HandlerUser) UpdateUser(ctx context.Context, req *dto.UpdateUserReq) (*user.U, error) {
	logger := log.With().Str("method", "update_user").Logger()

	if req == nil {
		return &user.U{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #MARK:Authenticate
	claims, err := h.User.Auth(ctx, "access")
	if err != nil {
		return &user.U{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	var u user.U

	if err := h.User.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		now := time.Now().Unix()
		us, err := h.User.Update(ctx, user.Filter{
			ID: claims.UserID,
		}, user.Patch{
			FirstName: pbtypes.GetString(req.Firstname),
			LastName:  pbtypes.GetString(req.Lastname),
			AvatarURL: pbtypes.GetString(req.AvatarURL),
			UpdatedAt: &now,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to update user")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(us) == 0 {
			err := gerrors.ErrNotFound{Resource: "user", Index: claims.UserID.String()}
			logger.Error().Err(err).Msg("failed to update user")

			return transaction.Rollback, status.New(codes.NotFound, err.Error()).Err()
		}

		u = us[0]

		return transaction.Commit, nil
	}); err != nil {
		return &user.U{}, err
	}

	logger.Info().Msg("success")

	return &u, nil
}
