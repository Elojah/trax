package main

import (
	"context"

	"github.com/elojah/trax/pkg/grpc"
	"github.com/elojah/trax/pkg/http"
	"github.com/ilyakaznacheev/cleanenv"
)

type web struct {
	Static string `json:"static"`
}

type config struct {
	HTTP       http.Config       `json:"http"`
	Web        web               `json:"web"`
	AuthClient grpc.ConfigClient `json:"auth_client"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
