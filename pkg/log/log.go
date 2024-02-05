package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

// Service embed a zerolog logger.
type Service struct {
	zerolog.Logger
}

// Dial start log dispatching.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	zerolog.TimeFieldFormat = ""
	s.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	return nil
}

func (s *Service) Close(ctx context.Context) error {
	return nil
}
