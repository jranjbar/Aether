package server

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouterRegistersProbeRoutes(t *testing.T) {
	router := NewRouter(slog.New(slog.NewTextHandler(io.Discard, nil)), routerChecker{})

	for _, path := range []string{"/health", "/ready", "/live"} {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
			}
		})
	}
}

type routerChecker struct{}

func (routerChecker) Check(ctx context.Context) error {
	return nil
}
