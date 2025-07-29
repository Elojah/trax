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

type sqlPermission struct {
	RoleID    ulid.ID
	Resource  string
	Command   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newPermission(p user.Permission) sqlPermission {
	return sqlPermission{
		RoleID:    p.RoleID,
		Resource:  p.Resource.String(),
		Command:   p.Command.String(),
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlPermission) permission() (user.Permission, error) {
	resource, ok := user.Resource_value[sqlp.Resource]
	if !ok {
		return user.Permission{}, postgres.ErrUnknownEnumValue{Enum: "resource", Value: sqlp.Resource}
	}

	command, ok := user.Command_value[sqlp.Command]
	if !ok {
		return user.Permission{}, postgres.ErrUnknownEnumValue{Enum: "command", Value: sqlp.Command}
	}

	return user.Permission{
		RoleID:    sqlp.RoleID,
		Resource:  user.Resource(resource),
		Command:   user.Command(command),
		CreatedAt: sqlp.CreatedAt.Unix(),
		UpdatedAt: sqlp.UpdatedAt.Unix(),
	}, nil
}

type filterPermission user.FilterPermission

func (f filterPermission) where() (string, []any) {
	var clause []string
	var args []any
	n := 1

	if f.RoleID != nil {
		clause = append(clause, fmt.Sprintf(`p.role_id = $%d`, n))
		args = append(args, f.RoleID)
		n++
	}

	if len(f.RoleIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`p.role_id IN (%s)`, postgres.Array(n, len(f.RoleIDs))))
		args = append(args, ulid.IDs(f.RoleIDs).Any()...)
		n += len(f.RoleIDs)
	}

	if f.Resource != nil {
		clause = append(clause, fmt.Sprintf(`p.resource = $%d`, n))
		args = append(args, f.Resource.String())
		n++
	}

	if f.Command != nil {
		clause = append(clause, fmt.Sprintf(`p.command = $%d`, n))
		args = append(args, f.Command.String())
		// n++
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

func (f filterPermission) index() string {
	var cols []string

	if f.RoleID != nil {
		cols = append(cols, f.RoleID.String())
	}

	if f.RoleIDs != nil {
		ss := ulid.IDs(f.RoleIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.Resource != nil {
		cols = append(cols, f.Resource.String())
	}
	if f.Command != nil {
		cols = append(cols, f.Command.String())
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertPermission(ctx context.Context, permission user.Permission) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newPermission(permission)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."permission" (role_id, resource, command, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 5))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.RoleID, p.Resource, p.Command, p.CreatedAt, p.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (s Store) InsertBatchPermission(ctx context.Context, permissions ...user.Permission) error {
	if len(permissions) == 0 {
		return nil
	}

	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	ps := make([]sqlPermission, len(permissions))
	for i, permission := range permissions {
		ps[i] = newPermission(permission)
	}

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."permission" (role_id, resource, command, created_at, updated_at) VALUES `)
	b.WriteString(postgres.BatchInsert(5, len(permissions)))

	args := make([]any, 0, len(ps)*5)
	for _, p := range ps {
		args = append(args, p.RoleID, p.Resource, p.Command, p.CreatedAt, p.UpdatedAt)
	}

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "permission", fmt.Sprintf("batch insert %d permissions", len(permissions)))
	}

	return nil
}

func (s Store) FetchPermission(ctx context.Context, f user.FilterPermission) (user.Permission, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Permission{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT p.role_id, p.resource, p.command, p.created_at, p.updated_at FROM "user"."permission" p `)

	clause, args := filterPermission(f).where()
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlPermission
	if err := q.Scan(&p.RoleID, &p.Resource, &p.Command, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.Permission{}, postgres.Error(err, "permission", filterPermission(f).index())
	}

	return p.permission()
}

func (s Store) ListPermission(ctx context.Context, f user.FilterPermission) ([]user.Permission, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT p.role_id, p.resource, p.command, p.created_at, p.updated_at FROM "user"."permission" p `)

	clause, args := filterPermission(f).where()
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, err
	}

	var permissions []user.Permission

	for rows.Next() {
		var p sqlPermission
		if err := rows.Scan(&p.RoleID, &p.Resource, &p.Command, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}

		sqlP, err := p.permission()
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, sqlP)
	}

	return permissions, nil
}

func (s Store) DeletePermission(ctx context.Context, f user.FilterPermission) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."permission" p `)

	clause, args := filterPermission(f).where()
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return err
	}

	return err
}
