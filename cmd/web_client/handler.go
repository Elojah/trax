package main

import (
	"github.com/elojah/trax/cmd/auth/grpc"
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/google"
)

type handler struct {
	cookie cookie.Agg
	google google.Agg

	grpc.AuthClient
}
