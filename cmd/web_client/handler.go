package main

import (
	"github.com/elojah/trax/cmd/auth/grpc"
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/google"
	"github.com/elojah/trax/pkg/twitch"
)

type handler struct {
	cookie cookie.App
	twitch twitch.App
	google google.App

	grpc.AuthClient
}
