package repository

import (
	"context"
	"errors"
	"testing"
)

func TestHealthRepositoryCheck(t *testing.T) {
	repo := NewHealthRepository(fakePinger{})

	if err := repo.Check(context.Background()); err != nil {
		t.Fatalf("Check() error = %v", err)
	}
}

func TestHealthRepositoryCheckReturnsPingError(t *testing.T) {
	expected := errors.New("offline")
	repo := NewHealthRepository(fakePinger{err: expected})

	err := repo.Check(context.Background())
	if !errors.Is(err, expected) {
		t.Fatalf("Check() error = %v, want %v", err, expected)
	}
}

type fakePinger struct {
	err error
}

func (f fakePinger) Ping(ctx context.Context) error {
	return f.err
}
