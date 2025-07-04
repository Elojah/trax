package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	ppostgres "github.com/elojah/trax/pkg/paginate/postgres"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

var (
	sortRole = map[string]string{
		"name":       "r.name",
		"created_at": "r.created_at",
		"updated_at": "r.updated_at",
	}
)

type sqlRole struct {
	ID        ulid.ID
	EntityID  ulid.ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newRole(r user.Role) sqlRole {
	return sqlRole{
		ID:        r.ID,
		EntityID:  r.EntityID,
		Name:      r.Name,
		CreatedAt: time.Unix(r.CreatedAt, 0),
		UpdatedAt: time.Unix(r.UpdatedAt, 0),
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
		clause = append(clause, fmt.Sprintf(`r.id = $%d`, n))
		args = append(args, f.ID)
		n++
	}

	if len(f.IDs) > 0 {
		clause = append(clause, fmt.Sprintf(`r.id IN (%s)`, postgres.Array(n, len(f.IDs))))
		args = append(args, ulid.IDs(f.IDs).Any()...)
		n += len(f.IDs)
	}

	if f.EntityID != nil {
		clause = append(clause, fmt.Sprintf(`r.entity_id = $%d`, n))
		args = append(args, f.EntityID)
		n++
	}

	if len(f.EntityIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`r.entity_id IN (%s)`, postgres.Array(n, len(f.EntityIDs))))
		args = append(args, ulid.IDs(f.EntityIDs).Any()...)
		n += len(f.EntityIDs)
	}

	// ListRoleByUser only
	if f.UserID != nil {
		clause = append(clause, fmt.Sprintf(`ru.user_id = $%d`, n))
		args = append(args, f.UserID)
		n++
	}

	if len(f.Search) > 0 {
		clause = append(clause, fmt.Sprintf(`r.name ILIKE $%d`, n))
		args = append(args, "%"+f.Search+"%")
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

	if f.EntityID != nil {
		cols = append(cols, f.EntityID.String())
	}

	if f.EntityIDs != nil {
		ss := ulid.IDs(f.EntityIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.Search != "" {
		cols = append(cols, f.Search)
	}

	return strings.Join(cols, " - ")
}

type patchRole user.PatchRole

func (e patchRole) set() (string, []any, int) {
	var cols []string
	var args []any
	n := 1

	if e.Name != nil {
		cols = append(cols, fmt.Sprintf(`name = $%d`, n))
		args = append(args, *e.Name)
		n++
	}

	if e.UpdatedAt != nil {
		cols = append(cols, fmt.Sprintf(`updated_at = $%d`, n))
		args = append(args, time.Unix(*e.UpdatedAt, 0))
		n++
	}

	b := strings.Builder{}
	b.WriteString(" SET ")

	if len(cols) == 0 {
		b.WriteString("id = id")
	} else {
		b.WriteString(strings.Join(cols, " , "))
	}

	return b.String(), args, n
}

func (s Store) InsertRole(ctx context.Context, role user.Role) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	r := newRole(role)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."role" (id, entity_id, name, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 5))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), r.ID, r.EntityID, r.Name, r.CreatedAt, r.UpdatedAt); err != nil {
		return postgres.Error(err, "role", r.ID.String())
	}

	return nil
}

func (s Store) UpdateRole(ctx context.Context, f user.FilterRole, p user.PatchRole) ([]user.Role, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."role" r `)

	cols, args, n := patchRole(p).set()
	b.WriteString(cols)

	clause, wargs := filterRole(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING id, entity_id, name, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "entity", filterRole(f).index())
	}

	var roles []user.Role

	for rows.Next() {
		var r sqlRole
		if err := rows.Scan(&r.ID, &r.EntityID, &r.Name, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "role", filterRole(f).index())
		}

		roles = append(roles, r.role())
	}

	return roles, nil
}

func (s Store) FetchRole(ctx context.Context, f user.FilterRole) (user.Role, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Role{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT r.id, r.entity_id, r.name, r.created_at, r.updated_at FROM "user"."role" r `)

	if f.UserID != nil {
		b.WriteString(` JOIN "user"."role_user" ru ON r.id = ru.role_id `)
	}

	clause, args := filterRole(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var r sqlRole
	if err := q.Scan(&r.ID, &r.EntityID, &r.Name, &r.CreatedAt, &r.UpdatedAt); err != nil {
		return user.Role{}, postgres.Error(err, "role", filterRole(f).index())
	}

	return r.role(), nil
}

func (s Store) ListRole(ctx context.Context, f user.FilterRole) ([]user.Role, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT r.id, r.entity_id, r.name, r.created_at, r.updated_at, COUNT(1) OVER() `)
	if f.Paginate != nil {
		b.WriteString(ppostgres.Paginate(*f.Paginate).Row(sortRole))
	} else {
		b.WriteString(`, 0`)
	}
	b.WriteString(` FROM "user"."role" r `)

	if f.UserID != nil {
		b.WriteString(` JOIN "user"."role_user" ru ON r.id = ru.role_id `)
	}

	clause, args := filterRole(f).where(1)
	b.WriteString(clause)

	if f.Paginate != nil {
		pag := ppostgres.Paginate(*f.Paginate).CTE(b.String())
		b.Reset()
		b.WriteString(pag)
	}

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "role", filterRole(f).index())
	}

	var roles []user.Role

	var count uint64
	var row_number int

	for rows.Next() {
		var r sqlRole
		if err := rows.Scan(&r.ID, &r.EntityID, &r.Name, &r.CreatedAt, &r.UpdatedAt, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "role", filterRole(f).index())
		}

		roles = append(roles, r.role())

	}

	return roles, count, nil
}

func (s Store) DeleteRole(ctx context.Context, f user.FilterRole) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."role" r `)

	clause, args := filterRole(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "role", filterRole(f).index())
	}

	return err
}
