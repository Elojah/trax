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

type sqlRole struct {
	ID        ulid.ID
	EntityID  ulid.ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newRole(p user.Role) sqlRole {
	return sqlRole{
		ID:        p.ID,
		EntityID:  p.EntityID,
		Name:      p.Name,
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlRole) role() user.Role {
	return user.Role{
		ID:        sqlp.ID,
		EntityID:  sqlp.EntityID,
		Name:      sqlp.Name,
		CreatedAt: sqlp.CreatedAt.Unix(),
		UpdatedAt: sqlp.UpdatedAt.Unix(),
	}
}

type filterRole user.FilterRole

func (f filterRole) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.ID != nil {
		clause = append(clause, fmt.Sprintf(`id = $%d`, n))
		args = append(args, f.ID)
		n++
	}

	if len(f.IDs) > 0 {
		clause = append(clause, fmt.Sprintf(`id IN (%s)`, postgres.Array(n, len(f.IDs))))
		args = append(args, ulid.IDs(f.IDs).Any()...)
		n += len(f.IDs)
	}

	if f.EntityID != nil {
		clause = append(clause, fmt.Sprintf(`entity_id = $%d`, n))
		args = append(args, f.EntityID)
		n++
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

func (f filterRole) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	if f.IDs != nil {
		ss := ulid.IDs(f.IDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertRole(ctx context.Context, role user.Role) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newRole(role)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."role" (id, entity_id, name, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 5))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.ID, p.EntityID, p.Name, p.CreatedAt, p.UpdatedAt); err != nil {
		return postgres.Error(err, "role", p.ID.String())
	}

	return nil
}

func (s Store) FetchRole(ctx context.Context, f user.FilterRole) (user.Role, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Role{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, entity_id, name, created_at, updated_at FROM "user"."role" `)

	clause, args := filterRole(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlRole
	if err := q.Scan(&p.ID, &p.EntityID, &p.Name, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.Role{}, postgres.Error(err, "role", filterRole(f).index())
	}

	return p.role(), nil
}

func (s Store) FetchManyRole(ctx context.Context, f user.FilterRole) ([]user.Role, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, entity_id, name, created_at, updated_at FROM "user"."role" `)

	clause, args := filterRole(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "role", filterRole(f).index())
	}

	var roles []user.Role

	for rows.Next() {
		var p sqlRole
		if err := rows.Scan(&p.ID, &p.EntityID, &p.Name, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "role", filterRole(f).index())
		}

		roles = append(roles, p.role())
	}

	return roles, nil
}

func (s Store) DeleteRole(ctx context.Context, f user.FilterRole) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."role" `)

	clause, args := filterRole(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "role", filterRole(f).index())
	}

	return err
}
