package invoice

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Invoice struct {
	ID        int       `db:"id"`
	UserID    string    `db:"user_id"`
	DueDate   time.Time `db:"due_date"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type InvoiceRepository struct {
	db *sqlx.DB
}

func NewInvoiceRepository(db *sqlx.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (repo *InvoiceRepository) CreateInvoice(ctx context.Context, inv *Invoice) (*Invoice, error) {
	query := `INSERT INTO invoices (user_id, due_date, status)
              VALUES (:user_id, :due_date, :status)
              RETURNING id, created_at, updated_at`

	stmt, args, err := sqlx.Named(query, inv)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}
	stmt = repo.db.Rebind(stmt)

	err = repo.db.GetContext(ctx, inv, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error creating invoice: %w", err)
	}

	return inv, nil
}

func (repo *InvoiceRepository) GetInvoice(ctx context.Context, id int) (*Invoice, error) {
	var inv Invoice
	query := "SELECT * FROM invoices WHERE id = ?"

	stmt := repo.db.Rebind(query)
	err := repo.db.GetContext(ctx, &inv, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("error fetching invoice: %w", err)
	}

	return &inv, nil
}

func (repo *InvoiceRepository) UpdateInvoice(ctx context.Context, inv *Invoice) (*Invoice, error) {
	query := `UPDATE invoices SET user_id = :user_id, due_date = :due_date, status = :status, updated_at = NOW()
              WHERE id = :id RETURNING updated_at`

	stmt, args, err := sqlx.Named(query, inv)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}
	stmt = repo.db.Rebind(stmt)

	err = repo.db.GetContext(ctx, inv, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error updating invoice: %w", err)
	}

	return inv, nil
}

func (repo *InvoiceRepository) DeleteInvoice(ctx context.Context, id int) error {
	query := "DELETE FROM invoices WHERE id = ?"
	stmt := repo.db.Rebind(query)

	_, err := repo.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting invoice: %w", err)
	}

	return nil
}

func (repo *InvoiceRepository) GetAllInvoices(ctx context.Context) ([]Invoice, error) {
	var invoices []Invoice
	query := "SELECT * FROM invoices"

	stmt := repo.db.Rebind(query)
	err := repo.db.SelectContext(ctx, &invoices, stmt)
	if err != nil {
		return nil, fmt.Errorf("error fetching invoices: %w", err)
	}

	return invoices, nil
}
