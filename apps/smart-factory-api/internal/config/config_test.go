package config

import (
	"testing"
	"time"
)

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("APP_NAME", "")
	t.Setenv("APP_ENV", "")
	t.Setenv("HTTP_HOST", "")
	t.Setenv("HTTP_PORT", "")
	t.Setenv("HTTP_READ_TIMEOUT", "")
	t.Setenv("HTTP_WRITE_TIMEOUT", "")
	t.Setenv("HTTP_IDLE_TIMEOUT", "")
	t.Setenv("HTTP_SHUTDOWN_TIMEOUT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.AppName != "smart-factory-api" {
		t.Fatalf("AppName = %q, want %q", cfg.AppName, "smart-factory-api")
	}

	if cfg.Env != "development" {
		t.Fatalf("Env = %q, want %q", cfg.Env, "development")
	}

	if cfg.Host != "0.0.0.0" {
		t.Fatalf("Host = %q, want %q", cfg.Host, "0.0.0.0")
	}

	if cfg.Port != "8080" {
		t.Fatalf("Port = %q, want %q", cfg.Port, "8080")
	}
}

func TestLoadUsesEnvironment(t *testing.T) {
	t.Setenv("APP_NAME", "factory")
	t.Setenv("APP_ENV", "test")
	t.Setenv("HTTP_HOST", "127.0.0.1")
	t.Setenv("HTTP_PORT", "9090")
	t.Setenv("HTTP_READ_TIMEOUT", "1s")
	t.Setenv("HTTP_WRITE_TIMEOUT", "2s")
	t.Setenv("HTTP_IDLE_TIMEOUT", "3s")
	t.Setenv("HTTP_SHUTDOWN_TIMEOUT", "4s")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.AppName != "factory" || cfg.Env != "test" || cfg.Host != "127.0.0.1" || cfg.Port != "9090" {
		t.Fatalf("Load() = %+v", cfg)
	}

	if cfg.ReadTimeout != time.Second || cfg.WriteTimeout != 2*time.Second || cfg.IdleTimeout != 3*time.Second || cfg.ShutdownTimeout != 4*time.Second {
		t.Fatalf("Load() timeouts = %+v", cfg)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	t.Setenv("HTTP_PORT", "invalid")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}
