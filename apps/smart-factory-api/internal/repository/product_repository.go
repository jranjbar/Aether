package repository

import (
	"context"

	"github.com/jranjbar/Aether/apps/smart-factory-api/internal/domain/product"
)

type ProductRepository interface {
	Create(ctx context.Context, p product.Product) error
	List(ctx context.Context) ([]product.Product, error)
	GetByID(ctx context.Context, id product.ProductID) (product.Product, error)
}
