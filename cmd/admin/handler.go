package main

import (
	"github.com/elojah/trax/pkg/cookie"
	"github.com/elojah/trax/pkg/migrate"
)

type handler struct {
	migrate migrate.Agg

	cookie cookie.Agg
}
