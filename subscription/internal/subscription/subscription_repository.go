package subscription

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Subscription struct {
	ID              int       `db:"id"`
	UserID          string    `db:"user_id"`
	ServiceID       int       `db:"service_id"`
	ServiceEndpoint string    `db:"service_endpoint"`
	StartDate       time.Time `db:"start_date"`
	EndDate         time.Time `db:"end_date"`
	InvoiceID       int       `db:"invoice_id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type SubscriptionRepository struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (repo *SubscriptionRepository) CreateSubscription(ctx context.Context, s *Subscription) (*Subscription, error) {
	query := `INSERT INTO subscriptions (user_id, service_id, service_endpoint, start_date, end_date, invoice_id)
              VALUES (:user_id, :service_id, :service_endpoint, :start_date, :end_date, :invoice_id)
              RETURNING id, created_at, updated_at`

	stmt, args, err := sqlx.Named(query, s)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}
	stmt = repo.db.Rebind(stmt)

	err = repo.db.GetContext(ctx, s, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error creating subscription: %w", err)
	}

	return s, nil
}

func (repo *SubscriptionRepository) GetSubscription(ctx context.Context, id int) (*Subscription, error) {
	var s Subscription
	query := "SELECT * FROM subscriptions WHERE id = ?"

	stmt := repo.db.Rebind(query)
	err := repo.db.GetContext(ctx, &s, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("error fetching subscription: %w", err)
	}

	return &s, nil
}

func (repo *SubscriptionRepository) UpdateSubscription(ctx context.Context, s *Subscription) (*Subscription, error) {
	query := `UPDATE subscriptions SET user_id = :user_id, service_id = :service_id, service_endpoint = :service_endpoint,
              start_date = :start_date, end_date = :end_date, invoice_id = :invoice_id, updated_at = NOW()
              WHERE id = :id RETURNING updated_at`

	stmt, args, err := sqlx.Named(query, s)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}
	stmt = repo.db.Rebind(stmt)

	err = repo.db.GetContext(ctx, s, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error updating subscription: %w", err)
	}

	return s, nil
}

func (repo *SubscriptionRepository) DeleteSubscription(ctx context.Context, id int) error {
	query := "DELETE FROM subscriptions WHERE id = ?"
	stmt := repo.db.Rebind(query)

	_, err := repo.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting subscription: %w", err)
	}

	return nil
}

func (repo *SubscriptionRepository) GetAllSubscriptions(ctx context.Context) ([]Subscription, error) {
	var subscriptions []Subscription
	query := "SELECT * FROM subscriptions"

	stmt := repo.db.Rebind(query)
	err := repo.db.SelectContext(ctx, &subscriptions, stmt)
	if err != nil {
		return nil, fmt.Errorf("error fetching subscriptions: %w", err)
	}

	return subscriptions, nil
}
