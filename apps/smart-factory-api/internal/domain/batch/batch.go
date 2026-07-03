package batch

import "time"

type BatchID string

type Batch struct {
	ID BatchID

	ProductID string

	BatchNumber string

	Quantity float64

	Status string

	ManufacturedAt time.Time
	ExpiresAt      time.Time
}
