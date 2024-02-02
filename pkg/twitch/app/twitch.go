package app

import (
	"context"

	"github.com/elojah/trax/pkg/twitch"
	"golang.org/x/oauth2"
)

var _ twitch.App = (*App)(nil)

type App struct {
	twitch.Client

	config oauth2.Config
}

func (a *App) Dial(ctx context.Context, cfg oauth2.Config) error {
	a.config = cfg

	return nil
}

func (a *App) Close(ctx context.Context) error {
	a.config = oauth2.Config{}

	return nil
}

func (a App) OAuth() oauth2.Config {
	return a.config
}
