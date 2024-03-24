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
		AvatarURL: sql.NullString{String: p.AvatarURL, Valid: p.AvatarURL != ""},
		CreatedAt: time.Unix(p.CreatedAt, 0),
		UpdatedAt: time.Unix(p.UpdatedAt, 0),
	}
}

func (sqlp sqlProfile) profile() user.Profile {
	return user.Profile{
		UserID:    sqlp.UserID,
		FirstName: sqlp.FirstName,
		LastName:  sqlp.LastName,
		AvatarURL: sqlp.AvatarURL.String,
		CreatedAt: sqlp.CreatedAt.Unix(),
		UpdatedAt: sqlp.UpdatedAt.Unix(),
	}
}

type filterProfile user.FilterProfile

func (f filterProfile) where() (string, []any) {
	var clause []string
	var args []any
	n := 1

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

func (s Store) InsertProfile(ctx context.Context, profile user.Profile) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	p := newProfile(profile)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."user_profile" (user_id, first_name, last_name, avatar_url, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 6))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), p.UserID, p.FirstName, p.LastName, p.AvatarURL, p.CreatedAt, p.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchProfile(ctx context.Context, f user.FilterProfile) (user.Profile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Profile{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT user_id, first_name, last_name, avatar_url, created_at, updated_at FROM "user"."user_profile" `)

	clause, args := filterProfile(f).where()
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var p sqlProfile
	if err := q.Scan(&p.UserID, &p.FirstName, &p.LastName, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return user.Profile{}, postgres.Error(err, "profile", filterProfile(f).index())
	}

	return p.profile(), nil
}

func (s Store) FetchManyProfile(ctx context.Context, f user.FilterProfile) ([]user.Profile, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT user_id, first_name, last_name, avatar_url, created_at, updated_at FROM "user"."user_profile" `)

	clause, args := filterProfile(f).where()
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, err
	}

	var profiles []user.Profile

	for rows.Next() {
		var p sqlProfile
		if err := rows.Scan(&p.UserID, &p.FirstName, &p.LastName, &p.AvatarURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
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
	b.WriteString(`DELETE FROM "user"."user_profile" `)

	clause, args := filterProfile(f).where()
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return err
	}

	return err
}
