package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/trax/internal/user/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h handler) signin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	var req dto.SigninReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error().Err(err).Msg("failed to decode request")

		http.Error(w, "failed to signup", http.StatusBadRequest)

		return
	}

	jwt, err := h.Signin(ctx, &req)
	if err != nil {
		logger.Error().Err(err).Msg("failed to signin")

		if e, ok := status.FromError(err); ok {
			switch e.Code() { //nolint: exhaustive
			case codes.Unauthenticated:
				http.Error(w, e.Message(), http.StatusUnauthorized)
			case codes.Internal:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			default:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "failed to signin", http.StatusInternalServerError)
		}

		return
	}

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

	w.Write([]byte(jwt.AccessToken))

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")
}
