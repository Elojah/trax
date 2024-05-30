package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h handler) signout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = ctx

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	// Set refresh token
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    "",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Domain:   ".legacyfactory.com",
		MaxAge:   -1,
	})

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")
}
