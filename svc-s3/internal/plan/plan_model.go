package plan

import "time"

type Plan struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	DaysPerInterval int       `db:"days_per_interval"`
	StorageType     string    `db:"storage_type"`
	StorageSize     int       `db:"storage_size"`
	Price           int       `db:"price"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
