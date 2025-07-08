package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

type sqlRoleUser struct {
	RoleID    ulid.ID
	UserID    ulid.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newRoleUser(p user.RoleUser) sqlRoleUser {
	return sqlRoleUser{
		RoleID:    p.RoleID,
		UserID:    p.UserID,
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlRoleUser) roleUser() user.RoleUser {
	return user.RoleUser{
		RoleID:    sqlp.RoleID,
		UserID:    sqlp.UserID,
		CreatedAt: sqlp.CreatedAt.Unix(),
		UpdatedAt: sqlp.UpdatedAt.Unix(),
	}
}

type filterRoleUser user.FilterRoleUser

func (f filterRoleUser) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.RoleID != nil {
		clause = append(clause, fmt.Sprintf(`ru.role_id = $%d`, n))
		args = append(args, f.RoleID)
		n++
	}

	if len(f.RoleIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`ru.role_id IN (%s)`, postgres.Array(n, len(f.RoleIDs))))
		args = append(args, ulid.IDs(f.RoleIDs).Any()...)
		n += len(f.RoleIDs)
	}

	if f.UserID != nil {
		clause = append(clause, fmt.Sprintf(`ru.user_id = $%d`, n))
		args = append(args, f.UserID)
		n++
	}

	if len(f.UserIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`ru.user_id IN (%s)`, postgres.Array(n, len(f.UserIDs))))
		args = append(args, ulid.IDs(f.UserIDs).Any()...)
		n += len(f.UserIDs)
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

func (f filterRoleUser) index() string {
	var cols []string

	if f.RoleID != nil {
		cols = append(cols, f.RoleID.String())
	}

	if f.RoleIDs != nil {
		ss := ulid.IDs(f.RoleIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.UserID != nil {
		cols = append(cols, f.UserID.String())
	}

	if f.UserIDs != nil {
		ss := ulid.IDs(f.UserIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertRoleUser(ctx context.Context, roleUser user.RoleUser) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newRoleUser(roleUser)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."role_user" (role_id, user_id, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 4))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.RoleID, p.UserID, p.CreatedAt, p.UpdatedAt); err != nil {
		return postgres.Error(err, "role_user", p.RoleID.String())
	}

	return nil
}

func (s Store) FetchRoleUser(ctx context.Context, f user.FilterRoleUser) (user.RoleUser, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.RoleUser{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT ru.role_id, ru.user_id, ru.created_at, ru.updated_at FROM "user"."role_user" ru `)

	clause, args := filterRoleUser(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlRoleUser
	if err := q.Scan(&p.RoleID, &p.UserID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.RoleUser{}, postgres.Error(err, "role_user", filterRoleUser(f).index())
	}

	return p.roleUser(), nil
}

func (s Store) ListRoleUser(ctx context.Context, f user.FilterRoleUser) ([]user.RoleUser, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT ru.role_id, ru.user_id, ru.created_at, ru.updated_at FROM "user"."role_user" ru `)

	clause, args := filterRoleUser(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "role_user", filterRoleUser(f).index())
	}

	var roleUsers []user.RoleUser

	for rows.Next() {
		var p sqlRoleUser
		if err := rows.Scan(&p.RoleID, &p.UserID, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "role_user", filterRoleUser(f).index())
		}

		roleUsers = append(roleUsers, p.roleUser())
	}

	return roleUsers, nil
}

func (s Store) DeleteRoleUser(ctx context.Context, f user.FilterRoleUser) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."role_user" ru `)

	clause, args := filterRoleUser(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "role_user", filterRoleUser(f).index())
	}

	return err
}

func (s Store) ListClaims(ctx context.Context, userID ulid.ID) (user.ClaimAuth, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.ClaimAuth{}, err
	}

	b := strings.Builder{}
	b.WriteString(`
		SELECT r.group_id, r.id, p.resource, p.command
		FROM "user"."role_user" ru
		JOIN "user"."role" r ON ru.role_id = r.id
		JOIN "user"."permission" p ON r.id = p.role_id
		WHERE ru.user_id = $1
	`)

	rows, err := tx.Query(ctx, b.String(), userID)
	if err != nil {
		return user.ClaimAuth{}, postgres.Error(err, "role_user", userID.String())
	}

	var claim user.ClaimAuth

	for rows.Next() {
		var groupID, roleID ulid.ID
		var resource, command string

		if err := rows.Scan(&groupID, &roleID, &resource, &command); err != nil {
			return user.ClaimAuth{}, postgres.Error(err, "role_user+claim", userID.String())
		}

		eid := groupID.String()
		rid := roleID.String()

		if claim.Groups == nil {
			claim.Groups = make(map[string]user.ClaimGroup)
		}

		if _, ok := claim.Groups[eid]; !ok {
			claim.Groups[eid] = user.ClaimGroup{
				ID:        groupID,
				Roles:     make(map[string]pbtypes.Empty),
				Resources: make(map[string]user.ClaimResources),
			}
		}

		if _, ok := claim.Groups[eid].Roles[rid]; !ok {
			claim.Groups[eid].Roles[rid] = pbtypes.Empty{}
		}

		if _, ok := claim.Groups[eid].Resources[resource]; !ok {
			claim.Groups[eid].Resources[resource] = user.ClaimResources{
				Commands: make(map[string]pbtypes.Empty),
			}
		}

		if _, ok := claim.Groups[eid].Resources[resource].Commands[command]; !ok {
			claim.Groups[eid].Resources[resource].Commands[command] = pbtypes.Empty{}
		}
	}

	return claim, nil
}
