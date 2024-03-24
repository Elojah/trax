package main

import (
	"context"

	"github.com/elojah/trax/pkg/grpc"
	"github.com/elojah/trax/pkg/http"
	"github.com/elojah/trax/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
)

type web struct {
	Static string `json:"static"`
}

type config struct {
	HTTP       http.Config       `json:"http"`
	Web        web               `json:"web"`
	Redis      redis.Config      `json:"redis"`
	AuthClient grpc.ConfigClient `json:"auth_client"`

	Google oauth2.Config `json:"google"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
