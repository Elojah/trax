package main

import (
	"context"

	"github.com/elojah/trax/pkg/grpc"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
)

type config struct {
	Postgres postgres.Config `json:"postgres"`
	GRPC     grpc.Config     `json:"grpc"`
	Redis    redis.Config    `json:"redis"`

	Twitch oauth2.Config `json:"twitch"`
	Google oauth2.Config `json:"google"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
