package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	Conn *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) (*PostgresRepository) {
	return &PostgresRepository {
		Conn: db,
	}
}


type User struct {
	ID        int64      `db:"id"`
	Firstname string     `db:"firstname"`
	Lastname  string     `db:"lastname"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	IsAdmin   bool       `db:"is_admin"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}


type Session struct {
	ID           string    `db:"id"`
	UserEmail    string    `db:"user_email"`
	RefreshToken string    `db:"refresh_token"`
	IsRevoked    bool      `db:"is_revoked"`
	CreatedAt    time.Time `db:"created_at"`
	ExpiresAt    time.Time `db:"expires_at"`
}
