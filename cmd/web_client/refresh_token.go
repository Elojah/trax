package main

import (
	"net/http"

	"github.com/elojah/trax/pkg/pbtypes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	rc, err := r.Cookie("refresh")
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")

		http.Error(w, "failed to retrieve token", http.StatusUnauthorized)
		return
	}

	// Refresh with auth.
	jwt, err := h.RefreshToken(
		metadata.AppendToOutgoingContext(ctx, "token", rc.Value),
		&pbtypes.String{Value: ""},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to refresh")

		if e, ok := status.FromError(err); ok {
			switch e.Code() { //nolint: exhaustive
			case codes.Internal:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			case codes.InvalidArgument:
				http.Error(w, e.Message(), http.StatusBadRequest)
			default:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "failed to refresh", http.StatusInternalServerError)
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
		SameSite: http.SameSiteStrictMode,
		Domain:   ".legacyfactory.com",
		MaxAge:   24 * 60 * 60,
	})

	if _, err := w.Write([]byte(jwt.AccessToken)); err != nil {
		logger.Error().Err(err).Msg("failed to write response")

		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")
}
