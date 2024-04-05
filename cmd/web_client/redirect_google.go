package main

import (
	"net/http"

	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/rs/zerolog/log"
)

func (h handler) redirectGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	state := r.FormValue("state")

	osc, err := r.Cookie("oauth-state-callback")
	if err != nil || state == "" {
		http.Error(w, "missing state challenge", http.StatusBadRequest)

		return
	}

	challenge, err := h.cookie.Decode(ctx, "oauth-state-callback", osc.Value)
	if err != nil {
		logger.Error().Err(err).Msg("failed to decode cookie")
		http.Error(w, "failed to read cookie", http.StatusInternalServerError)

		return
	}

	if challenge != state {
		http.Error(w, "invalid oauth state", http.StatusBadRequest)

		return
	}

	// #MARK:Fetch token
	oa := h.google.OAuth()

	token, err := oa.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		logger.Error().Err(err).Msg("invalid oauth state")
		http.Error(w, "invalid oauth state", http.StatusBadRequest)

		return
	}

	// #MARK:Fetch JWT
	jwt, err := h.AuthClient.SigninGoogle(ctx, &pbtypes.String{Value: token.AccessToken})
	if err != nil {
		logger.Error().Err(err).Msg("failed to signin")
		http.Error(w, "failed to signin", http.StatusInternalServerError)

		return
	}

	// clear previous oauth cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth-state-callback",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: false,
	})

	// Set access token
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    jwt.AccessToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		Domain:   ".legacyfactory.com",
	})

	// Set refresh token
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    jwt.RefreshToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
		Domain:   ".legacyfactory.com",
	})

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
