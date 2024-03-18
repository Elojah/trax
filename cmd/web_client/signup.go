package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/trax/internal/user/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h handler) signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	var req dto.SignupReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error().Err(err).Msg("failed to decode request")

		http.Error(w, "failed to signup", http.StatusBadRequest)

		return
	}

	_, err := h.Signup(ctx, &req)
	if err != nil {
		logger.Error().Err(err).Msg("failed to signup")

		if e, ok := status.FromError(err); ok {
			switch e.Code() { //nolint: exhaustive
			case codes.AlreadyExists:
				http.Error(w, e.Message(), http.StatusConflict)
			case codes.Internal:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			default:
				http.Error(w, e.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "failed to signup", http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("success")
}
