package main

import (
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/migrate"
)

type handler struct {
	migrate migrate.App

	cookie cookie.App
}
