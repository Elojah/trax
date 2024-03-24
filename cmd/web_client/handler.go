package main

import (
	"github.com/elojah/trax/cmd/auth/grpc"
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/google"
)

type handler struct {
	cookie cookie.App
	google google.App

	grpc.AuthClient
}
