package product

import "time"

type ProductID string

type Product struct {
	ID          ProductID
	Code        string
	Name        string
	Description string

	Category string
	Unit     string

	CreatedAt time.Time
	UpdatedAt time.Time
}
