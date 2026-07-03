package service

import (
	"context"
	"errors"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/domain/product"
	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/repository"
)

var (
	ErrProductCodeRequired = errors.New("product code is required")
	ErrProductNameRequired = errors.New("product name is required")
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (s *ProductService) Create(ctx context.Context, p product.Product) error {

	if p.Code == "" {
		return ErrProductCodeRequired
	}

	if p.Name == "" {
		return ErrProductNameRequired
	}

	return s.repository.Create(ctx, p)
}
