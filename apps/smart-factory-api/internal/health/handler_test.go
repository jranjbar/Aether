package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerEndpointsReturnOK(t *testing.T) {
	handler := NewHandler()
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
