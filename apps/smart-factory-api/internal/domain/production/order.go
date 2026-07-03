package production

import "time"

type OrderID string

type Order struct {
	ID OrderID

	ProductID string

	BatchID string

	PlannedQuantity float64

	ProducedQuantity float64

	Status string

	StartedAt  time.Time
	FinishedAt *time.Time
}
