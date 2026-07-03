package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/config"
)

// Server owns the HTTP server lifecycle.
type Server struct {
	httpServer *http.Server
	log        *slog.Logger
}

// New creates an HTTP server from configuration and dependencies.
func New(cfg config.Config, log *slog.Logger, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
		log: log,
	}
}

// Start begins serving HTTP requests.
func (s *Server) Start() error {
	s.log.Info("server starting", slog.String("addr", s.httpServer.Addr))

	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.log.InfoContext(ctx, "server shutting down")
	return s.httpServer.Shutdown(ctx)
}
