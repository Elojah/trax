package main

import (
	"github.com/elojah/trax/cmd/auth/grpc"
)

type handler struct {
	grpc.AuthClient
}
