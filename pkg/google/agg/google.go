package agg

import (
	"context"

	"github.com/elojah/trax/pkg/google"
	"golang.org/x/oauth2"
	o2google "golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

var _ google.Agg = (*Agg)(nil)

type Agg struct {
	config oauth2.Config
}

func (a *Agg) Dial(ctx context.Context, cfg oauth2.Config) error {
	a.config = cfg

	a.config.Endpoint = o2google.Endpoint

	return nil
}

func (a *Agg) Close(ctx context.Context) error {
	a.config = oauth2.Config{}

	return nil
}

func (a Agg) OAuth() *oauth2.Config {
	return &a.config
}

func (a Agg) Signin(ctx context.Context, token string) (string, google.Claims, error) {
	// #MARK:Validate token
	p, err := idtoken.Validate(ctx, token, a.config.ClientID)
	if err != nil {
		return "", nil, err
	}

	gid, err := google.Claims(p.Claims).GetString("sub")
	if err != nil {
		return "", nil, err
	}

	return gid, p.Claims, nil
}
