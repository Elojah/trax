package twitch

import (
	"context"

	"golang.org/x/oauth2"
)

type Auth struct {
	Token    string
	ClientID string
}

type Cursor struct {
	Total  uint
	Cursor string
}

type Client interface {
	// User
	GetUsers(context.Context, Auth, UserFilter, func(User) error) error

	// Follow
	GetFollows(context.Context, Auth, FollowFilter, func(Follow) error) (Cursor, error)
}

type App interface {
	Client
	OAuth() oauth2.Config
}
