package agg

import (
	"context"

	"github.com/elojah/trax/pkg/errors"
)

type Agg struct {
	// *postgres.Service
}

func (a Agg) Up(ctx context.Context, dir string) error {
	return nil
	// return migrate.Migrate(ctx, a.Session, dir)
}

func (a Agg) Down(context.Context, string) error {
	return errors.ErrNotImplemented{Version: "soontm"}
}
