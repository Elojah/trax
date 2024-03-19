package postgres

import (
	"errors"

	terrors "github.com/elojah/trax/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func Error(err error, resource string, index string) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return terrors.ErrUnknown{}
	}

	switch pgErr.Code {
	case "23505":
		return terrors.ErrConflict{Resource: resource, Index: index}
	}

	return terrors.ErrUnknown{}
}
