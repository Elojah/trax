package main

import (
	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/google"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/elojah/trax/pkg/twitch"
)

type handler struct {
	transaction.Transactioner

	twitch twitch.App
	google google.App

	user user.App
}
