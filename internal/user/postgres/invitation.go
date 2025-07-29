package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elojah/trax/internal/user"
	ppostgres "github.com/elojah/trax/pkg/paginate/postgres"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/ulid"
	_ "github.com/lib/pq"
)

var (
	sortInvitation = map[string]string{
		"email":      "i.email",
		"created_at": "i.created_at",
		"updated_at": "i.updated_at",
	}
)

type sqlInvitation struct {
	ID        ulid.ID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newInvitation(i user.Invitation) sqlInvitation {
	return sqlInvitation{
		ID:        i.ID,
		Email:     i.Email,
		CreatedAt: time.Unix(i.CreatedAt, 0),
		UpdatedAt: time.Unix(i.UpdatedAt, 0),
	}
}

func (sqli sqlInvitation) invitation() user.Invitation {
	return user.Invitation{
		ID:        sqli.ID,
		Email:     sqli.Email,
		CreatedAt: sqli.CreatedAt.Unix(),
		UpdatedAt: sqli.UpdatedAt.Unix(),
	}
}

type filterInvitation user.FilterInvitation

func (f filterInvitation) where(n int) (string, []any) {
	var clause []string
	var args []any

	if f.ID != nil {
		clause = append(clause, fmt.Sprintf(`i.id = $%d`, n))
		args = append(args, f.ID)
		n++
	}

	if len(f.IDs) > 0 {
		clause = append(clause, fmt.Sprintf(`i.id IN (%s)`, postgres.Array(n, len(f.IDs))))
		args = append(args, ulid.IDs(f.IDs).Any()...)
		n += len(f.IDs)
	}

	if f.Email != nil {
		clause = append(clause, fmt.Sprintf(`i.email = $%d`, n))
		args = append(args, *f.Email)
		n++
	}

	if len(f.Emails) > 0 {
		clause = append(clause, fmt.Sprintf(`i.email IN (%s)`, postgres.Array(n, len(f.Emails))))
		for _, email := range f.Emails {
			args = append(args, email)
		}
		n += len(f.Emails)
	}

	if len(f.Search) > 0 {
		clause = append(clause, fmt.Sprintf(`i.email ILIKE $%d`, n))
		args = append(args, "%"+f.Search+"%")
		// n++
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

func (f filterInvitation) index() string {
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

	if f.Emails != nil {
		cols = append(cols, strings.Join(f.Emails, "|"))
	}

	if f.Search != "" {
		cols = append(cols, f.Search)
	}

	return strings.Join(cols, " - ")
}

type patchInvitation user.PatchInvitation

func (p patchInvitation) set() (string, []any, int) {
	var cols []string
	var args []any
	n := 1

	if p.Email != nil {
		cols = append(cols, fmt.Sprintf(`email = $%d`, n))
		args = append(args, *p.Email)
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
		b.WriteString("id = id")
	} else {
		b.WriteString(strings.Join(cols, " , "))
	}

	return b.String(), args, n
}

func (s Store) InsertInvitation(ctx context.Context, invitation user.Invitation) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	i := newInvitation(invitation)

	b := strings.Builder{}
	b.WriteString(`INSERT INTO "user"."invitation" (id, email, created_at, updated_at) VALUES (`)
	b.WriteString(postgres.Array(1, 4))
	b.WriteString(`)`)

	if _, err := tx.Exec(ctx, b.String(), i.ID, i.Email, i.CreatedAt, i.UpdatedAt); err != nil {
		return postgres.Error(err, "invitation", i.ID.String())
	}

	return nil
}

func (s Store) UpdateInvitation(ctx context.Context, f user.FilterInvitation, p user.PatchInvitation) ([]user.Invitation, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, err
	}

	b := strings.Builder{}
	b.WriteString(`UPDATE "user"."invitation" i `)

	cols, args, n := patchInvitation(p).set()
	b.WriteString(cols)

	clause, wargs := filterInvitation(f).where(n)
	b.WriteString(clause)

	args = append(args, wargs...)

	b.WriteString(` RETURNING id, email, created_at, updated_at`)

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, postgres.Error(err, "invitation", filterInvitation(f).index())
	}

	var invitations []user.Invitation

	for rows.Next() {
		var i sqlInvitation
		if err := rows.Scan(&i.ID, &i.Email, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, postgres.Error(err, "invitation", filterInvitation(f).index())
		}

		invitations = append(invitations, i.invitation())
	}

	return invitations, nil
}

func (s Store) FetchInvitation(ctx context.Context, f user.FilterInvitation) (user.Invitation, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return user.Invitation{}, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT i.id, i.email, i.created_at, i.updated_at FROM "user"."invitation" i `)

	clause, args := filterInvitation(f).where(1)
	b.WriteString(clause)

	q := tx.QueryRow(ctx, b.String(), args...)

	var i sqlInvitation
	if err := q.Scan(&i.ID, &i.Email, &i.CreatedAt, &i.UpdatedAt); err != nil {
		return user.Invitation{}, postgres.Error(err, "invitation", filterInvitation(f).index())
	}

	return i.invitation(), nil
}

func (s Store) ListInvitation(ctx context.Context, f user.FilterInvitation) ([]user.Invitation, uint64, error) {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return nil, 0, err
	}

	b := strings.Builder{}
	b.WriteString(`SELECT i.id, i.email, i.created_at, i.updated_at, COUNT(1) OVER() `)
	if f.Paginate != nil {
		b.WriteString(ppostgres.Paginate(*f.Paginate).Row(sortInvitation))
	} else {
		b.WriteString(`, 0 `)
	}
	b.WriteString(` FROM "user"."invitation" i `)

	clause, args := filterInvitation(f).where(1)
	b.WriteString(clause)

	if f.Paginate != nil {
		pag := ppostgres.Paginate(*f.Paginate).CTE(b.String())
		b.Reset()
		b.WriteString(pag)
	}

	rows, err := tx.Query(ctx, b.String(), args...)
	if err != nil {
		return nil, 0, postgres.Error(err, "invitation", filterInvitation(f).index())
	}

	var invitations []user.Invitation

	var count uint64
	var row_number int

	for rows.Next() {
		var i sqlInvitation
		if err := rows.Scan(&i.ID, &i.Email, &i.CreatedAt, &i.UpdatedAt, &count, &row_number); err != nil {
			return nil, 0, postgres.Error(err, "invitation", filterInvitation(f).index())
		}

		invitations = append(invitations, i.invitation())
	}

	return invitations, count, nil
}

func (s Store) DeleteInvitation(ctx context.Context, f user.FilterInvitation) error {
	tx, err := postgres.Tx(ctx)
	if err != nil {
		return err
	}

	b := strings.Builder{}
	b.WriteString(`DELETE FROM "user"."invitation" i `)

	clause, args := filterInvitation(f).where(1)
	b.WriteString(clause)

	if _, err := tx.Exec(ctx, b.String(), args...); err != nil {
		return postgres.Error(err, "invitation", filterInvitation(f).index())
	}

	return nil
}
