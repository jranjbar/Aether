package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/config"
	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/database"
	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/logger"
	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/repository"
	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/server"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log := logger.New(cfg)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.Open(cfg.Database)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error("database close failed", slog.String("error", err.Error()))
		}
	}()

	healthRepository := repository.NewHealthRepository(db)
	router := server.NewRouter(log, healthRepository)
	httpServer := server.New(cfg, log, router)
	errCh := make(chan error, 1)

	go func() {
		errCh <- httpServer.Start()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return err
		}

		return nil
	case err := <-errCh:
		if err != nil {
			return err
		}

		return nil
	}
}
