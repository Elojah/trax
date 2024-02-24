package main

import (
	"github.com/elojah/trax/pkg/google"
	"github.com/elojah/trax/pkg/twitch"
	"github.com/elojah/trax/internal/user"
)

type handler struct {
	twitch twitch.App
	google google.App

	user user.App
}
