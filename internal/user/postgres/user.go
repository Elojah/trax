package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	pagpostgres "github.com/elojah/trax/pkg/paginate/postgres"
	ppostgres "github.com/elojah/trax/pkg/pbtypes/postgres"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

var (
	sortUser = map[string]string{
		"email":      "u.email",
		"name":       "u.last_name",
		"created_at": "u.created_at",
		"updated_at": "u.updated_at",
	}
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

	// !!! Only available if role r is joined
	if len(f.EntityIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`r.entity_id IN (%s)`, postgres.Array(n, len(f.EntityIDs))))
		args = append(args, ulid.IDs(f.EntityIDs).Any()...)
		n += len(f.EntityIDs)
	}

	// !!! Only available if role r is joined
	if len(f.RoleIDs) > 0 {
		clause = append(clause, fmt.Sprintf(`r.id IN (%s)`, postgres.Array(n, len(f.RoleIDs))))
		args = append(args, ulid.IDs(f.RoleIDs).Any()...)
		n += len(f.RoleIDs)
	}

	if len(f.Search) > 0 {
		clause = append(clause, fmt.Sprintf(`( u.email ILIKE $%d OR u.first_name ILIKE $%d OR u.last_name ILIKE $%d )`, n, n+1, n+2))
		args = append(args, "%"+f.Search+"%", "%"+f.Search+"%", "%"+f.Search+"%")
		n += 3
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

	if f.Search != "" {
		cols = append(cols, f.Search)
	}

	if f.EntityIDs != nil {
		ss := ulid.IDs(f.EntityIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.RoleIDs != nil {
		ss := ulid.IDs(f.RoleIDs).String()
		cols = append(cols, strings.Join(ss, "|"))
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

func (s Store) FetchWithPassword(ctx context.Context, f user.Filter) (user.U, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.U{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT u.id, u.email, u.password_hash, u.password_salt, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at FROM "user"."user" u `)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var u sqlUser
	if err := q.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.PasswordSalt, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return user.U{}, postgres.Error(err, "user", filter(f).index())
	}

	return u.user(), nil
}

func (s Store) List(ctx context.Context, f user.Filter) ([]user.U, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT u.id, u.email, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at, COUNT(1) OVER() `)
	if f.Paginate != nil {
		b.WriteString(pagpostgres.Paginate(*f.Paginate).Row(sortUser))
	} else {
		b.WriteString(`, 0`)
	}
	b.WriteString(`FROM "user"."user" u `)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "user", filter(f).index())
	}

	var users []user.U

	var count uint64
	var row_number int

	for rows.Next() {
		var u sqlUser
		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "user", filter(f).index())
		}

		users = append(users, u.user())
	}

	return users, count, nil
}

func (s Store) ListByEntity(ctx context.Context, f user.Filter) ([]user.U, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT DISTINCT ON (u.id) u.id, u.email, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at, COUNT(1) OVER() `)
	if f.Paginate != nil {
		b.WriteString(pagpostgres.Paginate(*f.Paginate).Row(sortUser))
	} else {
		b.WriteString(`, 0`)
	}
	b.WriteString(`
	FROM "user"."user" u
	JOIN "user"."role_user" ru ON u.id = ru.user_id
	JOIN "user"."role" r ON ru.role_id = r.id
	`)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "user+role", filter(f).index())
	}

	var users []user.U

	var count uint64
	var row_number int

	for rows.Next() {
		var u sqlUser
		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "user+role", filter(f).index())
		}

		users = append(users, u.user())
	}

	return users, count, nil
}

func (s Store) ListByRole(ctx context.Context, f user.Filter) (map[string][]user.U, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT DISTINCT ON (u.id) u.id, u.email, u.first_name, u.last_name, u.avatar_url, u.created_at, u.updated_at, COUNT(1) OVER() `)
	if f.Paginate != nil {
		b.WriteString(pagpostgres.Paginate(*f.Paginate).Row(sortUser))
	} else {
		b.WriteString(`, 0`)
	}
	b.WriteString(`
	FROM "user"."user" u
	JOIN "user"."role_user" ru ON u.id = ru.user_id
	JOIN "user"."role" r ON ru.role_id = r.id
	`)

	clause, args := filter(f).where(1)
	b.WriteString(clause)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "user+role", filter(f).index())
	}

	users := make(map[string][]user.U)

	var count uint64
	var row_number int

	for rows.Next() {
		var u sqlUser
		var roleID ulid.ID

		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.AvatarURL, &u.CreatedAt, &u.UpdatedAt, &roleID, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "user+role", filter(f).index())
		}

		users[roleID.String()] = append(users[roleID.String()], u.user())
	}

	return users, count, nil
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
