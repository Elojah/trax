package main

import (
	"context"

	"github.com/elojah/trax/pkg/grpc"
	"github.com/elojah/trax/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Redis redis.Config `json:"redis"`
	GRPC  grpc.Config  `json:"grpc"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
