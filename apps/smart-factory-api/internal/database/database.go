package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/config"
)

// DB wraps a PostgreSQL connection pool.
type DB struct {
	pool        *sql.DB
	pingTimeout time.Duration
}

// Open creates a PostgreSQL connection pool.
func Open(cfg config.DatabaseConfig) (*DB, error) {
	pool, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	pool.SetMaxOpenConns(cfg.MaxOpenConns)
	pool.SetMaxIdleConns(cfg.MaxIdleConns)
	pool.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &DB{
		pool:        pool,
		pingTimeout: cfg.PingTimeout,
	}, nil
}

// Ping verifies database connectivity.
func (db *DB) Ping(ctx context.Context) error {
	pingCtx, cancel := context.WithTimeout(ctx, db.pingTimeout)
	defer cancel()

	if err := db.pool.PingContext(pingCtx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	return nil
}

// Close releases database resources.
func (db *DB) Close() error {
	if err := db.pool.Close(); err != nil {
		return fmt.Errorf("close database: %w", err)
	}

	return nil
}

// Pool exposes the underlying database/sql connection pool.
func (db *DB) Pool() *sql.DB {
	return db.pool
}
