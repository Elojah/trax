package main

import (
	"github.com/elojah/trax/cmd/api/entity"
	"github.com/elojah/trax/cmd/api/role"
	"github.com/elojah/trax/cmd/api/user"
)

type handler struct {
	entity.HandlerEntity
	role.HandlerRole
	user.HandlerUser
}
