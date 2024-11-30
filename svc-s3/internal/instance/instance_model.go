package instance

import "time"

type Instance struct {
	ID        int       `db:"id"`
	PlanID    string    `db:"plan_id"`
	UserID    string    `db:"user_id"`
	IsActive  int       `db:"is_active"`
	AccessKey string    `db:"access_key"`
	SecretKey string    `db:"secret_key"`
	Endpoint  string    `db:"endpoint"`
	Region    string    `db:"region"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
