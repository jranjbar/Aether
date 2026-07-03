package repository

import (
	"context"
)

// Pinger verifies connectivity to a backing service.
type Pinger interface {
	Ping(ctx context.Context) error
}

// HealthRepository provides persistence health checks.
type HealthRepository struct {
	db Pinger
}

// NewHealthRepository creates a persistence health repository.
func NewHealthRepository(db Pinger) HealthRepository {
	return HealthRepository{
		db: db,
	}
}

// Check verifies database connectivity.
func (r HealthRepository) Check(ctx context.Context) error {
	return r.db.Ping(ctx)
}
