package main

import (
	"github.com/elojah/trax/internal/user"
	"github.com/elojah/trax/pkg/google"
)

type handler struct {
	google google.Service

	user user.Service
}
