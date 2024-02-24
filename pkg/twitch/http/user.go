package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/elojah/trax/pkg/twitch"
)

type userFilter twitch.UserFilter

func (f userFilter) set(req *http.Request) {
	q := req.URL.Query()

	// optional add user ids
	for _, id := range f.IDs {
		q.Add("id", id)
	}

	req.URL.RawQuery = q.Encode()
}

func (c Client) GetUsers(
	ctx context.Context,
	a twitch.Auth,
	f twitch.UserFilter,
	callback func(twitch.User) error,
) error {
	b, err := url.Parse(twitchURL)
	if err != nil {
		return err
	}

	r, err := url.Parse("/helix/users")
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.ResolveReference(r).String(), nil)
	if err != nil {
		return err
	}

	auth(a).set(req)

	userFilter(f).set(req)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return twitch.ErrInvalidStatus{Status: resp.Status}
	}

	var result struct {
		Data []twitch.User
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	for _, u := range result.Data {
		if err := callback(u); err != nil {
			return err
		}
	}

	return nil
}
