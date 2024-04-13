package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	ppostgres "github.com/elojah/trax/pkg/pbtypes/postgres"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

type sqlEntityProfile struct {
	EntityID    ulid.ID
	Description sql.NullString
	AvatarURL   sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func newEntityProfile(p user.EntityProfile) sqlEntityProfile {
	return sqlEntityProfile{
		EntityID:    p.EntityID,
		Description: ppostgres.NullString(p.Description),
		AvatarURL:   ppostgres.NullString(p.AvatarURL),
		CreatedAt:   time.Unix(p.CreatedAt, 0),
		UpdatedAt:   time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlEntityProfile) profile() user.EntityProfile {
	return user.EntityProfile{
		EntityID:    sqlp.EntityID,
		Description: ppostgres.NewString(sqlp.Description),
		AvatarURL:   ppostgres.NewString(sqlp.AvatarURL),
		CreatedAt:   sqlp.CreatedAt.Unix(),
		UpdatedAt:   sqlp.UpdatedAt.Unix(),
	}
}

type filterEntityProfile user.FilterEntityProfile

func (f filterEntityProfile) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.EntityID != nil {
		clause = append(clause, fmt.Sprintf(`user_id = $%d`, n))
		args = append(args, f.EntityID)
		n++
	}

	if len(f.EntityIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`user_id IN (%s)`, postgres.Array(n, len(f.EntityIDs))))
		args = append(args, ulid.IDs(f.EntityIDs).Any()...)
		n += len(f.EntityIDs)
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

func (f filterEntityProfile) index() string {
	var cols []string

	if f.EntityID != nil {
		cols = append(cols, f.EntityID.String())
	}

	if f.EntityIDs != nil {
		ss := ulid.IDs(f.EntityIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

type patchEntityProfile user.PatchEntityProfile

func (p patchEntityProfile) set() (string, []any, int) {
	var cols []string
	var args []any
	n := 1

	if p.Description != nil {
		cols = append(cols, fmt.Sprintf(`first_name = $%d`, n))
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
		b.WriteString("user_id = user_id")
	} else {
		b.WriteString(strings.Join(cols, " , "))
	}

	return b.String(), args, n
}

func (s Store) InsertEntityProfile(ctx context.Context, profile user.EntityProfile) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newEntityProfile(profile)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."entity_profile" (entity_id, description, avatar_url, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 5))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.EntityID, p.Description, p.AvatarURL, p.CreatedAt, p.UpdatedAt); err != nil {
		return postgres.Error(err, "entity_profile", p.EntityID.String())
	}

	return nil
}

func (s Store) FetchEntityProfile(ctx context.Context, f user.FilterEntityProfile) (user.EntityProfile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.EntityProfile{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, description, avatar_url, created_at, updated_at FROM "user"."entity_profile" `)

	clause, args := filterEntityProfile(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlEntityProfile
	if err := q.Scan(&p.EntityID, &p.Description, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.EntityProfile{}, postgres.Error(err, "entity_profile", filterEntityProfile(f).index())
	}

	return p.profile(), nil
}

func (s Store) ListEntityProfile(ctx context.Context, f user.FilterEntityProfile) ([]user.EntityProfile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, description, avatar_url, created_at, updated_at FROM "user"."entity_profile" `)

	clause, args := filterEntityProfile(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "entity_profile", filterEntityProfile(f).index())
	}

	var profiles []user.EntityProfile

	for rows.Next() {
		var p sqlEntityProfile
		if err := rows.Scan(&p.EntityID, &p.Description, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "entity_profile", filterEntityProfile(f).index())
		}

		profiles = append(profiles, p.profile())
	}

	return profiles, nil
}

func (s Store) UpdateEntityProfile(ctx context.Context, f user.FilterEntityProfile, p user.PatchEntityProfile) ([]user.EntityProfile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."entity_profile" `)

	cols, args, n := patchEntityProfile(p).set()
	b.WriteString(cols)

	clause, wargs := filterEntityProfile(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING entity_id, description, avatar_url, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "entity_profile", filterEntityProfile(f).index())
	}

	var profiles []user.EntityProfile

	for rows.Next() {
		var p sqlEntityProfile
		if err := rows.Scan(&p.EntityID, &p.Description, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "entity_profile", filterEntityProfile(f).index())
		}

		profiles = append(profiles, p.profile())
	}

	return profiles, nil
}

func (s Store) DeleteEntityProfile(ctx context.Context, f user.FilterEntityProfile) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."entity_profile" `)

	clause, args := filterEntityProfile(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "entity_profile", filterEntityProfile(f).index())
	}

	return err
}
