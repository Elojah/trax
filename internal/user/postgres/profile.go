package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/elojah/trax/internal/user"
	terrors "github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

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
	b := strings.Builder{}
	b.WriteString(`INSERT INTO main.profile (user_id, first_name, last_name, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 7))
	b.WriteString(`)`)

	_, err := s.Service.DB.ExecContext(
		ctx,
		b.String(),
		profile.UserID, profile.Firstname, profile.LastName, profile.CreatedAt, profile.UpdatedAt,
	)

	return err
}

func (s Store) FetchProfile(ctx context.Context, f user.FilterProfile) (user.Profile, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT user_id, first_name, last_name, created_at, updated_at FROM main.profile `)

	clause, args := filterProfile(f).where()
	b.WriteString(clause)

	q := s.Service.DB.QueryRowContext(ctx, b.String(), args...)

	var profile user.Profile
	if err := q.Scan(&profile.UserID, &profile.Firstname, &profile.LastName, &profile.CreatedAt, &profile.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.Profile{}, terrors.ErrNotFound{Resource: "profile", Index: filterProfile(f).index()}
		}

		return user.Profile{}, err
	}

	return profile, nil
}

func (s Store) FetchManyProfile(ctx context.Context, f user.FilterProfile) ([]user.Profile, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT user_id, first_name, last_name, created_at, updated_at FROM main.profile `)

	clause, args := filterProfile(f).where()
	b.WriteString(clause)

	rows, err := s.Service.DB.QueryContext(ctx, b.String(), args...)
	if err != nil {
		return nil, err
	}

	var profiles []user.Profile

	for rows.Next() {
		var profile user.Profile
		if err := rows.Scan(&profile.UserID, &profile.Firstname, &profile.LastName, &profile.CreatedAt, &profile.UpdatedAt); err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (s Store) DeleteProfile(ctx context.Context, f user.FilterProfile) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.profile `)

	clause, args := filterProfile(f).where()
	b.WriteString(clause)

	_, err := s.Service.DB.ExecContext(ctx, b.String(), args...)
	if err != nil {
		return err
	}

	return err
}
