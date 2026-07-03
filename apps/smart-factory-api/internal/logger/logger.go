package logger

import (
	"log/slog"
	"os"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/config"
)

// New creates a structured logger configured for the current environment.
func New(cfg config.Config) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if cfg.Env == "development" {
		opts.Level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	return slog.New(handler).With(
		slog.String("app", cfg.AppName),
		slog.String("env", cfg.Env),
	)
}
