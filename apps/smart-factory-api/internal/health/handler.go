package health

import (
	"context"
	"net/http"
	"time"
)

// Checker verifies a dependency needed for serving traffic.
type Checker interface {
	Check(ctx context.Context) error
}

// Handler serves health, readiness, and liveness probes.
type Handler struct {
	checker Checker
	timeout time.Duration
}

// NewHandler creates a health probe handler.
func NewHandler(checker Checker) Handler {
	return Handler{
		checker: checker,
		timeout: 2 * time.Second,
	}
}

// Health responds to broad health checks.
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if err := h.check(r.Context()); err != nil {
		writeUnavailable(w)
		return
	}

	writeOK(w)
}

// Ready responds to readiness checks.
func (h Handler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.check(r.Context()); err != nil {
		writeUnavailable(w)
		return
	}

	writeOK(w)
}

// Live responds to liveness checks.
func (h Handler) Live(w http.ResponseWriter, r *http.Request) {
	writeOK(w)
}

func (h Handler) check(ctx context.Context) error {
	checkCtx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	return h.checker.Check(checkCtx)
}

func writeOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK\n"))
}

func writeUnavailable(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusServiceUnavailable)
	_, _ = w.Write([]byte("UNAVAILABLE\n"))
}
