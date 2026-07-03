package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config contains runtime settings for the Smart Factory API.
type Config struct {
	AppName         string
	Env             string
	Host            string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

// Load reads configuration from environment variables.
func Load() (Config, error) {
	cfg := Config{
		AppName:         getEnv("APP_NAME", "smart-factory-api"),
		Env:             getEnv("APP_ENV", "development"),
		Host:            getEnv("HTTP_HOST", "0.0.0.0"),
		Port:            getEnv("HTTP_PORT", "8080"),
		ReadTimeout:     getDurationEnv("HTTP_READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getDurationEnv("HTTP_WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:     getDurationEnv("HTTP_IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getDurationEnv("HTTP_SHUTDOWN_TIMEOUT", 10*time.Second),
	}

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) validate() error {
	port, err := strconv.Atoi(c.Port)
	if err != nil {
		return fmt.Errorf("HTTP_PORT must be a number: %w", err)
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("HTTP_PORT must be between 1 and 65535")
	}

	if c.ReadTimeout <= 0 {
		return fmt.Errorf("HTTP_READ_TIMEOUT must be positive")
	}

	if c.WriteTimeout <= 0 {
		return fmt.Errorf("HTTP_WRITE_TIMEOUT must be positive")
	}

	if c.IdleTimeout <= 0 {
		return fmt.Errorf("HTTP_IDLE_TIMEOUT must be positive")
	}

	if c.ShutdownTimeout <= 0 {
		return fmt.Errorf("HTTP_SHUTDOWN_TIMEOUT must be positive")
	}

	return nil
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	return value
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}
