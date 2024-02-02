package google

import (
	"context"

	"golang.org/x/oauth2"
)

type App interface {
	Signin(context.Context, string) (string, error)
	OAuth() oauth2.Config
}
