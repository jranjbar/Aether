package database

import (
	"testing"
	"time"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/config"
)

func TestOpenConfiguresPool(t *testing.T) {
	db, err := Open(config.DatabaseConfig{
		URL:             "postgres://user:pass@localhost:5432/db?sslmode=disable",
		MaxOpenConns:    7,
		MaxIdleConns:    3,
		ConnMaxLifetime: time.Minute,
		PingTimeout:     time.Second,
	})
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer db.Close()

	stats := db.Pool().Stats()
	if stats.MaxOpenConnections != 7 {
		t.Fatalf("MaxOpenConnections = %d, want %d", stats.MaxOpenConnections, 7)
	}
}
