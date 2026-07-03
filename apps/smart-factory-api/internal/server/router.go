package server

import (
	"log/slog"
	"net/http"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/health"
	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/middleware"
)

// NewRouter builds the API routes and middleware stack.
func NewRouter(log *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	healthHandler := health.NewHandler()

	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("GET /ready", healthHandler.Ready)
	mux.HandleFunc("GET /live", healthHandler.Live)

	return middleware.Logging(log)(mux)
}
