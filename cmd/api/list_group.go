package main

import (
	"context"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/internal/user/dto"
	gerrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/paginate"
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

	var groupViews []dto.GroupView
	var total uint64

	if err := h.user.Tx(ctx, transaction.Write, func(ctx context.Context) (transaction.Operation, error) {
		groups, totalGroup, err := h.user.ListGroup(ctx, user.FilterGroup{
			IDs:      ids,
			Paginate: req.Paginate,
			Search:   req.Search,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		if len(groups) == 0 {
			logger.Info().Msg("no groups found")
			return transaction.Commit, nil
		}

		users, _, err := h.user.ListByGroup(ctx, user.Filter{
			GroupIDs: ids,
			Paginate: &paginate.Paginate{
				Start: 0,
				End:   9,
				Sort:  "created_at",
				Order: false,
			},
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to list users by group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		countUser, err := h.user.CountByGroup(ctx, user.Filter{
			GroupIDs: ids,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to count users by group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		countRole, err := h.user.CountRoleByGroup(ctx, user.FilterRole{
			GroupIDs: ids,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to count roles by group")

			return transaction.Rollback, status.New(codes.Internal, err.Error()).Err()
		}

		total = totalGroup
		groupViews = make([]dto.GroupView, 0, len(groups))
		for _, group := range groups {
			groupView := dto.GroupView{
				Group: group,
				UserSample: func() []user.U {
					if sample, ok := users[group.ID.String()]; ok {
						return sample
					}
					return []user.U{}
				}(),
				UserCount: func() uint64 {
					if count, ok := countUser[group.ID.String()]; ok {
						return count
					}
					return 0
				}(),
				RoleCount: func() uint64 {
					if count, ok := countRole[group.ID.String()]; ok {
						return count
					}
					return 0
				}(),
			}
			groupViews = append(groupViews, groupView)
		}

		return transaction.Commit, nil
	}); err != nil {
		logger.Error().Err(err).Msg("failed transaction")

		return &dto.ListGroupResp{}, err
	}

	logger.Info().Msg("success")

	return &dto.ListGroupResp{
		Groups: groupViews,
		Total:  total,
	}, nil
}
