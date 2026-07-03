package quality

import "time"

type RecordID string

type Record struct {
	ID RecordID

	BatchID string

	Passed bool

	Inspector string

	Notes string

	CreatedAt time.Time
}
