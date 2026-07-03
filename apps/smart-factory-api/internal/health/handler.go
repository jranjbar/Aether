package health

import (
	"net/http"
)

// Handler serves health, readiness, and liveness probes.
type Handler struct{}

// NewHandler creates a health probe handler.
func NewHandler() Handler {
	return Handler{}
}

// Health responds to broad health checks.
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeOK(w)
}

// Ready responds to readiness checks.
func (h Handler) Ready(w http.ResponseWriter, r *http.Request) {
	writeOK(w)
}

// Live responds to liveness checks.
func (h Handler) Live(w http.ResponseWriter, r *http.Request) {
	writeOK(w)
}

func writeOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK\n"))
}
