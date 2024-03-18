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
	"github.com/jackc/pgconn"
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

	if f.Email != nil {
		clause = append(clause, fmt.Sprintf(`email = $%d`, n))
		args = append(args, *f.Email)
		n++
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

	if f.Email != nil {
		cols = append(cols, *f.Email)
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
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."user" (id, email, password, google_id, twitch_id, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 7))
	b.WriteString(`)`)

	if _, err := tx.Exec(
		ctx,
		b.String(),
		u.ID, u.Email, u.Password, u.GoogleID, u.TwitchID, u.CreatedAt, u.UpdatedAt,
	); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return terrors.ErrConflict{Resource: "user", Index: u.ID.String()}
		}
	}

	return err
}

func (s Store) Fetch(ctx context.Context, f user.Filter) (user.U, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.U{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, email, password, google_id, twitch_id, created_at, updated_at FROM "user"."user" `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

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
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT id, email, password, google_id, twitch_id, created_at, updated_at FROM "user"."user" `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
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
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."user" `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return err
	}

	return err
}
