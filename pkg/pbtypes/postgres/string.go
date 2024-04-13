package postgres

import (
	"database/sql"

	"github.com/elojah/trax/pkg/pbtypes"
)

func NullString(s *pbtypes.String) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{String: s.Value, Valid: true}
}

func NewString(s sql.NullString) *pbtypes.String {
	if !s.Valid {
		return nil
	}

	return &pbtypes.String{Value: s.String}
}
