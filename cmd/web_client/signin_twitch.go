package main

import (
	"net/http"

	"github.com/elojah/trax/pkg/ulid"
	"github.com/rs/zerolog/log"
)

func (h handler) signinTwitch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	id := ulid.NewID()

	// Create JWT secure cookie.
	ck, err := h.cookie.Encode(ctx, "oauth-state-callback", id.String())
	if err != nil {
		logger.Error().Err(err).Msg("failed to encode token")
		http.Error(w, "failed to encode token", http.StatusInternalServerError)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth-state-callback",
		Value:    ck,
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
	})

	logger.Info().Msg("success")

	oa := h.twitch.OAuth()

	http.Redirect(w, r, oa.AuthCodeURL(id.String()), http.StatusTemporaryRedirect)
}
