package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	ppostgres "github.com/elojah/trax/pkg/paginate/postgres"
	tpostgres "github.com/elojah/trax/pkg/pbtypes/postgres"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

var (
	sortGroup = map[string]string{
		"name":       "g.name",
		"created_at": "g.created_at",
		"updated_at": "g.updated_at",
	}
)

type sqlGroup struct {
	ID          ulid.ID
	Name        string
	Description sql.NullString
	AvatarURL   sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func newGroup(g user.Group) sqlGroup {
	return sqlGroup{
		ID:          g.ID,
		Name:        g.Name,
		Description: tpostgres.NullStringEmpty(g.Description),
		AvatarURL:   tpostgres.NullStringEmpty(g.AvatarURL),
		CreatedAt:   time.Unix(g.CreatedAt, 0),
		UpdatedAt:   time.Unix(g.UpdatedAt, 0),
	}
}

func (sqlp sqlGroup) group() user.Group {
	return user.Group{
		ID:          sqlp.ID,
		Name:        sqlp.Name,
		AvatarURL:   sqlp.AvatarURL.String,
		Description: sqlp.Description.String,
		CreatedAt:   sqlp.CreatedAt.Unix(),
		UpdatedAt:   sqlp.UpdatedAt.Unix(),
	}
}

type filterGroup user.FilterGroup

func (f filterGroup) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.ID != nil {
		clause = append(clause, fmt.Sprintf(`g.id = $%d`, n))
		args = append(args, f.ID)
		n++
	}

	if len(f.IDs) > 0 {
		clause = append(clause, fmt.Sprintf(`g.id IN (%s)`, postgres.Array(n, len(f.IDs))))
		args = append(args, ulid.IDs(f.IDs).Any()...)
		n += len(f.IDs)
	}

	if len(f.Search) > 0 {
		clause = append(clause, fmt.Sprintf(`g.name ILIKE $%d`, n))
		args = append(args, "%"+f.Search+"%")
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

func (f filterGroup) index() string {
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

type patchGroup user.PatchGroup

func (p patchGroup) set() (string, []any, int) {
	var cols []string
	var args []any
	n := 1

	if p.Name != nil {
		cols = append(cols, fmt.Sprintf(`name = $%d`, n))
		args = append(args, *p.Name)
		n++
	}

	if p.Description != nil {
		cols = append(cols, fmt.Sprintf(`description = $%d`, n))
		args = append(args, *p.Description)
		n++
	}

	if p.AvatarURL != nil {
		cols = append(cols, fmt.Sprintf(`avatar_url = $%d`, n))
		args = append(args, *p.AvatarURL)
		n++
	}

	if p.UpdatedAt != nil {
		cols = append(cols, fmt.Sprintf(`updated_at = $%d`, n))
		args = append(args, time.Unix(*p.UpdatedAt, 0))
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

func (s Store) InsertGroup(ctx context.Context, group user.Group) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	g := newGroup(group)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."group" (id, name, description, avatar_url, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 6))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), g.ID, g.Name, g.Description, g.AvatarURL, g.CreatedAt, g.UpdatedAt); err != nil {
		return postgres.Error(err, "group", g.Name)
	}

	return nil
}

func (s Store) InsertBatchGroup(ctx context.Context, groups ...user.Group) error {
	if len(groups) == 0 {
		return nil
	}

	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	gs := make([]sqlGroup, len(groups))
	for i, group := range groups {
		gs[i] = newGroup(group)
	}

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."group" (id, name, description, avatar_url, created_at, updated_at) VALUES `)
	b.WriteString(postgres.BatchInsert(6, len(groups)))

	args := make([]any, 0, len(gs)*6)
	for _, g := range gs {
		args = append(args, g.ID, g.Name, g.Description, g.AvatarURL, g.CreatedAt, g.UpdatedAt)
	}

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "group", fmt.Sprintf("batch insert %d groups", len(groups)))
	}

	return nil
}

func (s Store) FetchGroup(ctx context.Context, f user.FilterGroup) (user.Group, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Group{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT g.id, g.name, g.description, g.avatar_url, g.created_at, g.updated_at FROM "user"."group" g `)

	clause, args := filterGroup(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var g sqlGroup
	if err := q.Scan(&g.ID, &g.Name, &g.CreatedAt, &g.UpdatedAt); err != nil {
		return user.Group{}, postgres.Error(err, "group", filterGroup(f).index())
	}

	return g.group(), nil
}

func (s Store) ListGroup(ctx context.Context, f user.FilterGroup) ([]user.Group, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT g.id, g.name, g.description, g.avatar_url, g.created_at, g.updated_at, COUNT(1) OVER() `)
	if f.Paginate != nil {
		b.WriteString(ppostgres.Paginate(*f.Paginate).Row(sortGroup))
	} else {
		b.WriteString(`, 0 `)
	}
	b.WriteString(` FROM "user"."group" g `)

	clause, args := filterGroup(f).where(1)
	b.WriteString(clause)

	if f.Paginate != nil {
		pag := ppostgres.Paginate(*f.Paginate).CTE(b.String())
		b.Reset()
		b.WriteString(pag)
	}

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "group", filterGroup(f).index())
	}

	var groups []user.Group

	var count uint64
	var row_number int

	for rows.Next() {
		var g sqlGroup
		if err := rows.Scan(&g.ID, &g.Name, &g.Description, &g.AvatarURL, &g.CreatedAt, &g.UpdatedAt, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "group", filterGroup(f).index())
		}

		groups = append(groups, g.group())
	}

	return groups, count, nil
}

func (s Store) UpdateGroup(ctx context.Context, f user.FilterGroup, p user.PatchGroup) ([]user.Group, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."group" g `)

	cols, args, n := patchGroup(p).set()
	b.WriteString(cols)

	clause, wargs := filterGroup(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING id, name, description, avatar_url, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "group", filterGroup(f).index())
	}

	var groups []user.Group

	for rows.Next() {
		var g sqlGroup
		if err := rows.Scan(&g.ID, &g.Name, &g.Description, &g.AvatarURL, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "group", filterGroup(f).index())
		}

		groups = append(groups, g.group())
	}

	return groups, nil
}

func (s Store) DeleteGroup(ctx context.Context, f user.FilterGroup) ([]user.Group, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."group" g `)

	clause, args := filterGroup(f).where(1)
	b.WriteString(clause)

	b.WriteString(` RETURNING id, name, description, avatar_url, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "group", filterGroup(f).index())
	}

	var groups []user.Group

	for rows.Next() {
		var g sqlGroup
		if err := rows.Scan(&g.ID, &g.Name, &g.Description, &g.AvatarURL, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "group", filterGroup(f).index())
		}

		groups = append(groups, g.group())
	}

	return groups, nil
}
