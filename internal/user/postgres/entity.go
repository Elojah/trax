package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

type sqlEntity struct {
	ID        ulid.ID
	Name      string
	AvatarURL sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newEntity(p user.Entity) sqlEntity {
	return sqlEntity{
		ID:        p.ID,
		Name:      p.Name,
		AvatarURL: sql.NullString{String: p.AvatarURL, Valid: p.AvatarURL != ""},
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlEntity) entity() user.Entity {
	return user.Entity{
		ID:        sqlp.ID,
		Name:      sqlp.Name,
		AvatarURL: sqlp.AvatarURL.String,
		CreatedAt: sqlp.CreatedAt.Unix(),
		UpdatedAt: sqlp.UpdatedAt.Unix(),
	}
}

type filterEntity user.FilterEntity

func (f filterEntity) where(n int) (string, []any) {
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

func (s Store) InsertEntity(ctx context.Context, entity user.Entity) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newEntity(entity)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."entity" (id, name, avatar_url, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 6))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.ID, p.Name, p.AvatarURL, p.CreatedAt, p.UpdatedAt); err != nil {
		return postgres.Error(err, "entity", p.ID.String())
	}

	return nil
}

func (s Store) FetchEntity(ctx context.Context, f user.FilterEntity) (user.Entity, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Entity{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, name, avatar_url, created_at, updated_at FROM "user"."entity" `)

	clause, args := filterEntity(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlEntity
	if err := q.Scan(&p.ID, &p.Name, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.Entity{}, postgres.Error(err, "entity", filterEntity(f).index())
	}

	return p.entity(), nil
}

func (s Store) FetchManyEntity(ctx context.Context, f user.FilterEntity) ([]user.Entity, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, name, avatar_url, created_at, updated_at FROM "user"."entity" `)

	clause, args := filterEntity(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "entity", filterEntity(f).index())
	}

	var entities []user.Entity

	for rows.Next() {
		var p sqlEntity
		if err := rows.Scan(&p.ID, &p.Name, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "entity", filterEntity(f).index())
		}

		entities = append(entities, p.entity())
	}

	return entities, nil
}

func (s Store) DeleteEntity(ctx context.Context, f user.FilterEntity) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."entity" `)

	clause, args := filterEntity(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "entity", filterEntity(f).index())
	}

	return err
}
