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
	sortEntity = map[string]string{
		"name":       "e.name",
		"created_at": "e.created_at",
		"updated_at": "e.updated_at",
	}
)

type sqlEntity struct {
	ID          ulid.ID
	Name        string
	Description sql.NullString
	AvatarURL   sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func newEntity(e user.Entity) sqlEntity {
	return sqlEntity{
		ID:          e.ID,
		Name:        e.Name,
		Description: tpostgres.NullStringEmpty(e.Description),
		AvatarURL:   tpostgres.NullStringEmpty(e.AvatarURL),
		CreatedAt:   time.Unix(e.CreatedAt, 0),
		UpdatedAt:   time.Unix(e.UpdatedAt, 0),
	}
}

func (sqlp sqlEntity) entity() user.Entity {
	return user.Entity{
		ID:          sqlp.ID,
		Name:        sqlp.Name,
		AvatarURL:   sqlp.AvatarURL.String,
		Description: sqlp.Description.String,
		CreatedAt:   sqlp.CreatedAt.Unix(),
		UpdatedAt:   sqlp.UpdatedAt.Unix(),
	}
}

type filterEntity user.FilterEntity

func (f filterEntity) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.ID != nil {
		clause = append(clause, fmt.Sprintf(`e.id = $%d`, n))
		args = append(args, f.ID)
		n++
	}

	if len(f.IDs) > 0 {
		clause = append(clause, fmt.Sprintf(`e.id IN (%s)`, postgres.Array(n, len(f.IDs))))
		args = append(args, ulid.IDs(f.IDs).Any()...)
		n += len(f.IDs)
	}

	if len(f.Search) > 0 {
		clause = append(clause, fmt.Sprintf(`e.name ILIKE $%d`, n))
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

func (f filterEntity) index() string {
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

type patchEntity user.PatchEntity

func (e patchEntity) set() (string, []any, int) {
	var cols []string
	var args []any
	n := 1

	if e.Name != nil {
		cols = append(cols, fmt.Sprintf(`name = $%d`, n))
		args = append(args, *e.Name)
		n++
	}

	if e.Description != nil {
		cols = append(cols, fmt.Sprintf(`description = $%d`, n))
		args = append(args, *e.Description)
		n++
	}

	if e.AvatarURL != nil {
		cols = append(cols, fmt.Sprintf(`avatar_url = $%d`, n))
		args = append(args, *e.AvatarURL)
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

func (s Store) InsertEntity(ctx context.Context, entity user.Entity) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	e := newEntity(entity)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."entity" (id, name, description, avatar_url, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 6))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), e.ID, e.Name, e.Description, e.AvatarURL, e.CreatedAt, e.UpdatedAt); err != nil {
		return postgres.Error(err, "entity", e.Name)
	}

	return nil
}

func (s Store) FetchEntity(ctx context.Context, f user.FilterEntity) (user.Entity, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Entity{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT e.id, e.name, e.description, e.avatar_url, e.created_at, e.updated_at FROM "user"."entity" e `)

	clause, args := filterEntity(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var e sqlEntity
	if err := q.Scan(&e.ID, &e.Name, &e.CreatedAt, &e.UpdatedAt); err != nil {
		return user.Entity{}, postgres.Error(err, "entity", filterEntity(f).index())
	}

	return e.entity(), nil
}

func (s Store) ListEntity(ctx context.Context, f user.FilterEntity) ([]user.Entity, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT e.id, e.name, e.description, e.avatar_url, e.created_at, e.updated_at, COUNT(1) OVER()`)
	if f.Paginate != nil {
		b.WriteString(ppostgres.Paginate(*f.Paginate).Row(sortEntity))
	} else {
		b.WriteString(`, 0`)
	}
	b.WriteString(` FROM "user"."entity" e `)

	clause, args := filterEntity(f).where(1)
	b.WriteString(clause)

	if f.Paginate != nil {
		pag := ppostgres.Paginate(*f.Paginate).CTE(b.String())
		b.Reset()
		b.WriteString(pag)
	}

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "entity", filterEntity(f).index())
	}

	var entities []user.Entity

	var count uint64
	var row_number int

	for rows.Next() {
		var e sqlEntity
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.AvatarURL, &e.CreatedAt, &e.UpdatedAt, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "entity", filterEntity(f).index())
		}

		entities = append(entities, e.entity())
	}

	return entities, count, nil
}

func (s Store) UpdateEntity(ctx context.Context, f user.FilterEntity, p user.PatchEntity) ([]user.Entity, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."entity" e `)

	cols, args, n := patchEntity(p).set()
	b.WriteString(cols)

	clause, wargs := filterEntity(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING id, name, description, avatar_url, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "entity", filterEntity(f).index())
	}

	var entities []user.Entity

	for rows.Next() {
		var e sqlEntity
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.AvatarURL, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "entity", filterEntity(f).index())
		}

		entities = append(entities, e.entity())
	}

	return entities, nil
}

func (s Store) DeleteEntity(ctx context.Context, f user.FilterEntity) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."entity" e `)

	clause, args := filterEntity(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "entity", filterEntity(f).index())
	}

	return err
}
