package service

import (
	"context"

	"github.com/elojah/trax/pkg/google"
	"golang.org/x/oauth2"
	o2google "golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

var _ google.Service = (*S)(nil)

type S struct {
	config oauth2.Config
}

func (s *S) Dial(ctx context.Context, cfg oauth2.Config) error {
	s.config = cfg

	s.config.Endpoint = o2google.Endpoint

	return nil
}

func (s *S) Close(ctx context.Context) error {
	s.config = oauth2.Config{}

	return nil
}

func (s S) OAuth() *oauth2.Config {
	return &s.config
}

func (s S) Signin(ctx context.Context, token string) (string, google.Claims, error) {
	// #MARK:Validate token
	p, err := idtoken.Validate(ctx, token, s.config.ClientID)
	if err != nil {
		return "", nil, err
	}

	gid, err := google.Claims(p.Claims).GetString("sub")
	if err != nil {
		return "", nil, err
	}

	return gid, p.Claims, nil
}
