package postgres

import (
	"errors"
	"fmt"

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

type ErrUnknownEnumValue struct {
	Enum  string
	Value string
}

func (e ErrUnknownEnumValue) Error() string {
	return fmt.Sprintf("unknown value %s for enum %s", e.Value, e.Enum)
}
