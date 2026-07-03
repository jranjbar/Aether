package inventory

type Item struct {
	ProductID string
	BatchID   string

	Quantity float64

	Warehouse string

	Location string
}
