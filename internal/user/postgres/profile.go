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

type sqlProfile struct {
	UserID    ulid.ID
	FirstName string
	LastName  string
	AvatarURL sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newProfile(p user.Profile) sqlProfile {
	return sqlProfile{
		UserID:    p.UserID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		AvatarURL: ppostgres.NullString(p.AvatarURL),
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlProfile) profile() user.Profile {
	return user.Profile{
		UserID:    sqlp.UserID,
		FirstName: sqlp.FirstName,
		LastName:  sqlp.LastName,
		AvatarURL: ppostgres.NewString(sqlp.AvatarURL),
		CreatedAt: sqlp.CreatedAt.Unix(),
		UpdatedAt: sqlp.UpdatedAt.Unix(),
	}
}

type filterProfile user.FilterProfile

func (f filterProfile) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.UserID != nil {
		clause = append(clause, fmt.Sprintf(`user_id = $%d`, n))
		args = append(args, f.UserID)
		n++
	}

	if len(f.UserIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`user_id IN (%s)`, postgres.Array(n, len(f.UserIDs))))
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

func (f filterProfile) index() string {
	var cols []string

	if f.UserID != nil {
		cols = append(cols, f.UserID.String())
	}

	if f.UserIDs != nil {
		ss := ulid.IDs(f.UserIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

type patchProfile user.PatchProfile

func (p patchProfile) set() (string, []any, int) {
	var cols []string
	var args []any
	n := 1

	if p.FirstName != nil {
		cols = append(cols, fmt.Sprintf(`first_name = $%d`, n))
		args = append(args, *p.FirstName)
		n++
	}

	if p.LastName != nil {
		cols = append(cols, fmt.Sprintf(`last_name = $%d`, n))
		args = append(args, *p.LastName)
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

func (s Store) InsertProfile(ctx context.Context, profile user.Profile) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newProfile(profile)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."profile" (user_id, first_name, last_name, avatar_url, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 6))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.UserID, p.FirstName, p.LastName, p.AvatarURL, p.CreatedAt, p.UpdatedAt); err != nil {
		return postgres.Error(err, "profile", p.UserID.String())
	}

	return nil
}

func (s Store) FetchProfile(ctx context.Context, f user.FilterProfile) (user.Profile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Profile{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT user_id, first_name, last_name, avatar_url, created_at, updated_at FROM "user"."profile" `)

	clause, args := filterProfile(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlProfile
	if err := q.Scan(&p.UserID, &p.FirstName, &p.LastName, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.Profile{}, postgres.Error(err, "profile", filterProfile(f).index())
	}

	return p.profile(), nil
}

func (s Store) ListProfile(ctx context.Context, f user.FilterProfile) ([]user.Profile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT user_id, first_name, last_name, avatar_url, created_at, updated_at FROM "user"."profile" `)

	clause, args := filterProfile(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "profile", filterProfile(f).index())
	}

	var profiles []user.Profile

	for rows.Next() {
		var p sqlProfile
		if err := rows.Scan(&p.UserID, &p.FirstName, &p.LastName, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "profile", filterProfile(f).index())
		}

		profiles = append(profiles, p.profile())
	}

	return profiles, nil
}

func (s Store) UpdateProfile(ctx context.Context, f user.FilterProfile, p user.PatchProfile) ([]user.Profile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."profile" `)

	cols, args, n := patchProfile(p).set()
	b.WriteString(cols)

	clause, wargs := filterProfile(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING user_id, first_name, last_name, avatar_url, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "profile", filterProfile(f).index())
	}

	var profiles []user.Profile

	for rows.Next() {
		var p sqlProfile
		if err := rows.Scan(&p.UserID, &p.FirstName, &p.LastName, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "profile", filterProfile(f).index())
		}

		profiles = append(profiles, p.profile())
	}

	return profiles, nil
}

func (s Store) DeleteProfile(ctx context.Context, f user.FilterProfile) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."profile" `)

	clause, args := filterProfile(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "profile", filterProfile(f).index())
	}

	return err
}
