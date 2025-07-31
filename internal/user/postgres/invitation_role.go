package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

type sqlInvitationRole struct {
	InvitationID ulid.ID
	RoleID       ulid.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func newInvitationRole(ir user.InvitationRole) sqlInvitationRole {
	return sqlInvitationRole{
		InvitationID: ir.InvitationID,
		RoleID:       ir.RoleID,
		CreatedAt:    time.Unix(ir.CreatedAt, 0),
		UpdatedAt:    time.Unix(ir.UpdatedAt, 0),
	}
}

func (sqlir sqlInvitationRole) invitationRole() user.InvitationRole {
	return user.InvitationRole{
		InvitationID: sqlir.InvitationID,
		RoleID:       sqlir.RoleID,
		CreatedAt:    sqlir.CreatedAt.Unix(),
		UpdatedAt:    sqlir.UpdatedAt.Unix(),
	}
}

type filterInvitationRole user.FilterInvitationRole

func (f filterInvitationRole) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.InvitationID != nil {
		clause = append(clause, fmt.Sprintf(`ir.invitation_id = $%d`, n))
		args = append(args, f.InvitationID)
		n++
	}

	if len(f.InvitationIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`ir.invitation_id IN (%s)`, postgres.Array(n, len(f.InvitationIDs))))
		args = append(args, ulid.IDs(f.InvitationIDs).Any()...)
		n += len(f.InvitationIDs)
	}

	if f.RoleID != nil {
		clause = append(clause, fmt.Sprintf(`ir.role_id = $%d`, n))
		args = append(args, f.RoleID)
		n++
	}

	// !!! Only available if role r is joined
	if len(f.GroupIDs) > 0 {
		// Only happens on groupIDs, change if needed
		clause = append(clause, `r.id = ir.role_id`)
		clause = append(clause, fmt.Sprintf(`r.group_id IN (%s)`, postgres.Array(n, len(f.GroupIDs))))
		args = append(args, ulid.IDs(f.GroupIDs).Any()...)
		n += len(f.GroupIDs)
	}

	if len(f.RoleIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`ir.role_id IN (%s)`, postgres.Array(n, len(f.RoleIDs))))
		args = append(args, ulid.IDs(f.RoleIDs).Any()...)
		// n += len(f.RoleIDs)
	}

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filterInvitationRole) index() string {
	var cols []string

	if f.InvitationID != nil {
		cols = append(cols, f.InvitationID.String())
	}

	if f.InvitationIDs != nil {
		uids := ulid.IDs(f.InvitationIDs).String()
		cols = append(cols, strings.Join(uids, "|"))
	}

	if f.RoleID != nil {
		cols = append(cols, f.RoleID.String())
	}

	if f.RoleIDs != nil {
		uids := ulid.IDs(f.RoleIDs).String()
		cols = append(cols, strings.Join(uids, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertInvitationRole(ctx context.Context, invitationRole user.InvitationRole) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	ir := newInvitationRole(invitationRole)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."invitation_role" (invitation_id, role_id, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 4))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), ir.InvitationID, ir.RoleID, ir.CreatedAt, ir.UpdatedAt); err != nil {
		return postgres.Error(err, "invitation_role", fmt.Sprintf("%s-%s", ir.InvitationID.String(), ir.RoleID.String()))
	}

	return nil
}

func (s Store) InsertBatchInvitationRole(ctx context.Context, invitationRoles ...user.InvitationRole) error {
	if len(invitationRoles) == 0 {
		return nil
	}

	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	irs := make([]sqlInvitationRole, len(invitationRoles))
	for i, ir := range invitationRoles {
		irs[i] = newInvitationRole(ir)
	}

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."invitation_role" (invitation_id, role_id, created_at, updated_at) VALUES `)
	b.WriteString(postgres.BatchInsert(4, len(invitationRoles)))

	args := make([]any, 0, len(irs)*4)
	for _, ir := range irs {
		args = append(args, ir.InvitationID, ir.RoleID, ir.CreatedAt, ir.UpdatedAt)
	}

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "invitation_role", fmt.Sprintf("batch insert %d roles", len(invitationRoles)))
	}

	return nil
}

func (s Store) FetchInvitationRole(ctx context.Context, f user.FilterInvitationRole) (user.InvitationRole, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.InvitationRole{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT ir.invitation_id, ir.role_id, ir.created_at, ir.updated_at FROM "user"."invitation_role" ir `)

	clause, args := filterInvitationRole(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var ir sqlInvitationRole
	if err := q.Scan(&ir.InvitationID, &ir.RoleID, &ir.CreatedAt, &ir.UpdatedAt); err != nil {
		return user.InvitationRole{}, postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
	}

	return ir.invitationRole(), nil
}

func (s Store) ListInvitationRole(ctx context.Context, f user.FilterInvitationRole) ([]user.InvitationRole, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT ir.invitation_id, ir.role_id, ir.created_at, ir.updated_at FROM "user"."invitation_role" ir `)

	clause, args := filterInvitationRole(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
	}

	var invitationRoles []user.InvitationRole

	for rows.Next() {
		var ir sqlInvitationRole
		if err := rows.Scan(&ir.InvitationID, &ir.RoleID, &ir.CreatedAt, &ir.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
		}

		invitationRoles = append(invitationRoles, ir.invitationRole())
	}

	return invitationRoles, nil
}

func (s Store) CountInvitationRoleByInvitation(ctx context.Context, f user.FilterInvitationRole) (map[string]uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT ir.invitation_id, COUNT(1) FROM "user"."invitation_role" ir `)

	clause, args := filterInvitationRole(f).where(1)
	b.WriteString(clause)
	b.WriteString(` GROUP BY ir.invitation_id`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
	}

	counts := make(map[string]uint64)

	for rows.Next() {
		var id ulid.ID
		var count uint64
		if err := rows.Scan(&id, &count); err != nil {
			return nil, postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
		}

		counts[id.String()] = count
	}

	return counts, nil
}

func (s Store) DeleteInvitationRole(ctx context.Context, f user.FilterInvitationRole) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."invitation_role" ir `)

	clause, args := filterInvitationRole(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
	}

	return nil
}

func (s Store) DeleteInvitationRoleByGroup(ctx context.Context, f user.FilterInvitationRole) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."invitation_role" ir USING "user"."role" r `)

	clause, args := filterInvitationRole(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "invitation_role", filterInvitationRole(f).index())
	}

	return nil
}
