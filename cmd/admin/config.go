package main

import (
	"context"

	"github.com/elojah/trax/pkg/grpc"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Postgres postgres.Config `json:"postgres"`
	Redis    redis.Config    `json:"redis"`
	GRPC     grpc.Config     `json:"grpc"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
