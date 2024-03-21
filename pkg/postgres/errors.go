package postgres

import (
	"errors"

	terrors "github.com/elojah/trax/pkg/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func Error(err error, resource string, index string) error {
	if err == pgx.ErrNoRows {
		return terrors.ErrNotFound{Resource: resource, Index: index}
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return terrors.ErrConflict{Resource: resource, Index: index}
		}
	}

	return terrors.ErrUnknown{Err: err}
}
