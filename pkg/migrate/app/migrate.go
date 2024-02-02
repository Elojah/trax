package app

import (
	"context"

	"github.com/elojah/go-postgres"
	"github.com/elojah/trax/pkg/errors"
	"github.com/postgresdb/gocqlx/migrate"
)

type App struct {
	*postgres.Service
}

func (a App) Up(ctx context.Context, dir string) error {
	return migrate.Migrate(ctx, a.Session, dir)
}

func (a App) Down(context.Context, string) error {
	return errors.ErrNotImplemented{Version: "soontm"}
}
