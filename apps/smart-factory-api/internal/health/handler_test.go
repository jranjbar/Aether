package health

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerEndpointsReturnOK(t *testing.T) {
	handler := NewHandler(stubChecker{})
	tests := map[string]http.HandlerFunc{
		"health": handler.Health,
		"ready":  handler.Ready,
		"live":   handler.Live,
	}

	for name, endpoint := range tests {
		t.Run(name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/"+name, nil)
			response := httptest.NewRecorder()

			endpoint(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
			}

			if response.Body.String() != "OK\n" {
				t.Fatalf("body = %q, want %q", response.Body.String(), "OK\n")
			}
		})
	}
}

func TestHandlerReturnsUnavailableWhenCheckFails(t *testing.T) {
	handler := NewHandler(stubChecker{err: errors.New("offline")})

	for _, endpoint := range []struct {
		name    string
		handler http.HandlerFunc
	}{
		{name: "health", handler: handler.Health},
		{name: "ready", handler: handler.Ready},
	} {
		t.Run(endpoint.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/"+endpoint.name, nil)
			response := httptest.NewRecorder()

			endpoint.handler(response, request)

			if response.Code != http.StatusServiceUnavailable {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusServiceUnavailable)
			}
		})
	}
}

type stubChecker struct {
	err error
}

func (s stubChecker) Check(ctx context.Context) error {
	return s.err
}
