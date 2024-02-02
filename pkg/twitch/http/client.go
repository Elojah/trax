package http

import (
	"net/http"

	"github.com/elojah/trax/pkg/twitch"
)

const (
	twitchURL = "https://api.twitch.tv"
)

var _ twitch.Client = (*Client)(nil)

type Client struct {
	*http.Client
}

type auth twitch.Auth

func (a auth) set(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+a.Token)
	req.Header.Set("Client-Id", a.ClientID)
}
