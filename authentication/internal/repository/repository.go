package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}


func (userRepository *UserRepository) Create(ctx context.Context, u *User) (*User, error) {
	// Using sqlx.Named to bind parameters for PostgreSQL
	query := `INSERT INTO users (firstname, lastname, email, password, is_admin)
              VALUES (:firstname, :lastname, :email, :password, :is_admin)
              RETURNING id`

	// Use sqlx.Named to ensure the named parameters are properly handled
	stmt, args, err := sqlx.Named(query, u)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	// Convert the named query to a prepared statement
	stmt = userRepository.db.Rebind(stmt)

	// Execute the query and scan the result into u.ID
	err = userRepository.db.GetContext(ctx, &u.ID, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	return u, nil
}

func (userRepository *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	// SQL query with named parameter for PostgreSQL
	query := "SELECT * FROM users WHERE email = :email"

	// Use sqlx.Named to bind the named parameters
	stmt, args, err := sqlx.Named(query, map[string]interface{}{"email": email})
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	// Rebind to convert the named parameters to positional ones for PostgreSQL
	stmt = userRepository.db.Rebind(stmt)

	// Execute the query and scan the result into u
	var u User
	err = userRepository.db.GetContext(ctx, &u, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &u, nil
}

func (userRepository *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	var users []User
	err := userRepository.db.SelectContext(ctx, &users, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return users, nil
}

func (userRepository *UserRepository) Update(ctx context.Context, u *User) (*User, error) {
	// Use sqlx.Named to handle named parameters in PostgreSQL
	query := `UPDATE users
              SET firstname = :firstname, lastname = :lastname, email = :email,
                  password = :password, is_admin = :is_admin, updated_at = :updated_at
              WHERE id = :id`

	// Bind the named parameters to the struct u
	stmt, args, err := sqlx.Named(query, u)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	// Rebind the query to PostgreSQL's positional parameters (i.e., $1, $2, ...)
	stmt = userRepository.db.Rebind(stmt)

	// Execute the update query
	_, err = userRepository.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return u, nil
}

func (userRepository *UserRepository) Delete(ctx context.Context, id int64) error {
	// Use Rebind to convert `?` into PostgreSQL's positional parameter `$1`
	query := "DELETE FROM users WHERE id = ?"
	stmt := userRepository.db.Rebind(query)

	// Execute the deletion query with the id as a positional parameter
	_, err := userRepository.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}


type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (sessionRepository *SessionRepository) CreateSession(ctx context.Context, s *Session) (*Session, error) {
	// SQL query with named parameters for PostgreSQL
	query := `INSERT INTO sessions (id, user_email, refresh_token, is_revoked, expires_at)
              VALUES (:id, :user_email, :refresh_token, :is_revoked, :expires_at)`

	// Use sqlx.Named to bind parameters to the struct
	stmt, args, err := sqlx.Named(query, s)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	// Rebind the query to use positional parameters for PostgreSQL
	stmt = sessionRepository.db.Rebind(stmt)

	// Execute the insert query
	_, err = sessionRepository.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting session: %w", err)
	}

	return s, nil
}

func (sessionRepository *SessionRepository) GetSession(ctx context.Context, id string) (*Session, error) {
	var s Session

	// Query for PostgreSQL with positional parameters
	query := "SELECT * FROM sessions WHERE id = ?"

	// Rebind the query to convert `?` to PostgreSQL's positional parameters
	stmt := sessionRepository.db.Rebind(query)

	// Execute the query
	err := sessionRepository.db.GetContext(ctx, &s, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("error getting session: %w", err)
	}

	return &s, nil
}

func (sessionRepository *SessionRepository) RevokeSession(ctx context.Context, id string) error {
	// SQL query with named parameters for PostgreSQL
	query := "UPDATE sessions SET is_revoked = true WHERE id = :id"

	// Use sqlx.Named to bind parameters to the map
	stmt, args, err := sqlx.Named(query, map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("error preparing named parameters: %w", err)
	}

	// Rebind the query for PostgreSQL's positional parameters
	stmt = sessionRepository.db.Rebind(stmt)

	// Execute the update query
	_, err = sessionRepository.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("error revoking session: %v", err)
	}

	return nil
}

func (sessionRepository *SessionRepository) DeleteSession(ctx context.Context, id string) error {
	// Query for PostgreSQL with positional parameters
	query := "DELETE FROM sessions WHERE id = ?"

	// Rebind the query to convert `?` to PostgreSQL's positional parameters
	stmt := sessionRepository.db.Rebind(query)

	// Execute the delete query
	_, err := sessionRepository.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}
