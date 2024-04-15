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

type sqlUser struct {
	ID           ulid.ID
	Email        string
	PasswordHash []byte
	PasswordSalt []byte
	GoogleID     sql.NullString
	FirstName    string
	LastName     string
	AvatarURL    sql.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func newUser(u user.U) sqlUser {
	return sqlUser{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		PasswordSalt: u.PasswordSalt,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		AvatarURL:    ppostgres.NullStringEmpty(u.AvatarURL),
		GoogleID:     ppostgres.NullStringEmpty(u.GoogleID),
		CreatedAt:    time.Unix(u.CreatedAt, 0),
		UpdatedAt:    time.Unix(u.UpdatedAt, 0),
	}
}

func (sqlu sqlUser) user() user.U {
	return user.U{
		ID:           sqlu.ID,
		Email:        sqlu.Email,
		PasswordHash: sqlu.PasswordHash,
		PasswordSalt: sqlu.PasswordSalt,
		FirstName:    sqlu.FirstName,
		LastName:     sqlu.LastName,
		AvatarURL:    sqlu.AvatarURL.String,
		GoogleID:     sqlu.GoogleID.String,
		CreatedAt:    sqlu.CreatedAt.Unix(),
		UpdatedAt:    sqlu.UpdatedAt.Unix(),
	}
}

type filter user.Filter

func (f filter) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.ID != nil {
		clause = append(clause, fmt.Sprintf(`u.id = $%d`, n))
		args = append(args, f.ID)
		n++
	}

	if len(f.IDs) > 0 {
		clause = append(clause, fmt.Sprintf(`u.id IN (%s)`, postgres.Array(n, len(f.IDs))))
		args = append(args, ulid.IDs(f.IDs).Any()...)
		n += len(f.IDs)
	}

	if f.Email != nil {
		clause = append(clause, fmt.Sprintf(`u.email = $%d`, n))
		args = append(args, *f.Email)
		n++
	}

	if f.GoogleID != nil {
		clause = append(clause, fmt.Sprintf(`u.google_id = $%d`, n))
		args = append(args, *f.GoogleID)
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

	return strings.Join(cols, " - ")
}

type patch user.Patch

func (p patch) set() (string, []any, int) {
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

func (s Store) Insert(ctx context.Context, u user.U) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."user" (id, email, password_hash, password_salt, first_name, last_name, avatar_url, google_id, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 10))
	b.WriteString(`)`)

	sqlu := newUser(u)

	if _, err := tx.Exec(
		ctx,
		b.String(),
		sqlu.ID, sqlu.Email, sqlu.PasswordHash, sqlu.PasswordSalt, sqlu.FirstName, sqlu.LastName, sqlu.AvatarURL, sqlu.GoogleID, sqlu.CreatedAt, sqlu.UpdatedAt,
	); err != nil {
		return postgres.Error(err, "user", u.ID.String())
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f user.Filter) (user.U, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.U{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT u.id, u.email, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at FROM "user"."user" u `)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var u sqlUser
	if err := q.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return user.U{}, postgres.Error(err, "user", filter(f).index())
	}

	return u.user(), nil
}

func (s Store) List(ctx context.Context, f user.Filter) ([]user.U, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT u.id, u.email, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at FROM "user"."user" u `)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "user", filter(f).index())
	}

	var users []user.U

	for rows.Next() {
		var u user.U
		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "user", filter(f).index())
		}

		users = append(users, u)
	}

	return users, nil
}

func (s Store) Update(ctx context.Context, f user.Filter, p user.Patch) ([]user.U, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."user" u `)

	cols, args, n := patch(p).set()
	b.WriteString(cols)

	clause, wargs := filter(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING u.id, u.email, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "user", filter(f).index())
	}

	var users []user.U

	for rows.Next() {
		var u sqlUser
		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "user", filter(f).index())
		}

		users = append(users, u.user())
	}

	return users, nil
}

func (s Store) Delete(ctx context.Context, f user.Filter) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."user" u `)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "user", filter(f).index())
	}

	return err
}
