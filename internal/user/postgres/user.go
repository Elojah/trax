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

type filter user.Filter

func (f filter) where() (string, []any) {
	var clause []string
	var args []any
	n := 1

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

	if f.GoogleID != nil {
		clause = append(clause, fmt.Sprintf(`google_id = $%d`, n))
		args = append(args, *f.GoogleID)
		n++
	}

	if f.TwitchID != nil {
		clause = append(clause, fmt.Sprintf(`twitch_id = $%d`, n))
		args = append(args, *f.TwitchID)
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

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	if f.IDs != nil {
		ss := ulid.IDs(f.IDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.GoogleID != nil {
		cols = append(cols, *f.GoogleID)
	}

	if f.TwitchID != nil {
		cols = append(cols, *f.TwitchID)
	}

	return strings.Join(cols, " - ")
}

func (s Store) Insert(ctx context.Context, u user.U) error {
	b := strings.Builder{}
	b.WriteString(`INSERT INTO main.user (id, email, password, google_id, twitch_id, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 7))
	b.WriteString(`)`)

	_, err := s.Service.DB.ExecContext(
		ctx,
		b.String(),
		u.ID, u.Email, u.Password, u.GoogleID, u.TwitchID, u.CreatedAt, u.UpdatedAt,
	)

	return err
}

func (s Store) Fetch(ctx context.Context, f user.Filter) (user.U, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, email, password, google_id, twitch_id, created_at, updated_at FROM main.user `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Service.DB.QueryRowContext(ctx, b.String(), args...)

	var u user.U
	if err := q.Scan(&u.ID, &u.Email, &u.Password, &u.GoogleID, &u.TwitchID, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.U{}, terrors.ErrNotFound{Resource: "user", Index: filter(f).index()}
		}

		return user.U{}, err
	}

	return u, nil
}

func (s Store) FetchMany(ctx context.Context, f user.Filter) ([]user.U, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, email, password, google_id, twitch_id, created_at, updated_at FROM main.user `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	rows, err := s.Service.DB.QueryContext(ctx, b.String(), args...)
	if err != nil {
		return nil, err
	}

	var users []user.U

	for rows.Next() {
		var u user.U
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.GoogleID, &u.TwitchID, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (s Store) Delete(ctx context.Context, f user.Filter) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.user `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	_, err := s.Service.DB.ExecContext(ctx, b.String(), args...)
	if err != nil {
		return err
	}

	return err
}
