package plan

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PlanRepository struct {
	db *sqlx.DB
}

func NewPlanRepository(db *sqlx.DB) *PlanRepository {
	return &PlanRepository{
		db: db,
	}
}

// CreatePlan creates a new plan in the database
func (repo *PlanRepository) CreatePlan(ctx context.Context, p *Plan) (*Plan, error) {
	query := `INSERT INTO public.plans (name, description, days_per_interval, storage_type, storage_size, price)
              VALUES (:name, :description, :days_per_interval, :storage_type, :storage_size, :price)
              RETURNING id, created_at, updated_at`

	stmt, args, err := sqlx.Named(query, p)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	stmt = repo.db.Rebind(stmt)

	var result Plan
	err = repo.db.GetContext(ctx, &result, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting plan: %w", err)
	}

	// Return the plan including the generated ID and timestamps
	p.ID = result.ID
	p.CreatedAt = result.CreatedAt
	p.UpdatedAt = result.UpdatedAt

	return p, nil
}

// GetPlan fetches a plan by its ID
func (repo *PlanRepository) GetPlan(ctx context.Context, id int) (*Plan, error) {
	var p Plan

	query := "SELECT * FROM public.plans WHERE id = ?"
	stmt := repo.db.Rebind(query)

	err := repo.db.GetContext(ctx, &p, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("error getting plan: %w", err)
	}

	return &p, nil
}

// UpdatePlan updates an existing plan
func (repo *PlanRepository) UpdatePlan(ctx context.Context, p *Plan) (*Plan, error) {
	query := `UPDATE public.plans
              SET name = :name, description = :description, days_per_interval = :days_per_interval,
                  storage_type = :storage_type, storage_size = :storage_size, price = :price,
                  updated_at = NOW()
              WHERE id = :id
              RETURNING updated_at`

	stmt, args, err := sqlx.Named(query, p)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	stmt = repo.db.Rebind(stmt)

	var updated Plan
	err = repo.db.GetContext(ctx, &updated, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error updating plan: %w", err)
	}

	p.UpdatedAt = updated.UpdatedAt

	return p, nil
}

// DeletePlan deletes a plan by its ID
func (repo *PlanRepository) DeletePlan(ctx context.Context, id int) error {
	query := "DELETE FROM public.plans WHERE id = ?"

	stmt := repo.db.Rebind(query)

	_, err := repo.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting plan: %w", err)
	}

	return nil
}

// GetAllPlans fetches all plans (for list-like functionality)
func (repo *PlanRepository) GetAllPlans(ctx context.Context) ([]Plan, error) {
	var plans []Plan

	query := "SELECT * FROM public.plans"
	stmt := repo.db.Rebind(query)

	err := repo.db.SelectContext(ctx, &plans, stmt)
	if err != nil {
		return nil, fmt.Errorf("error getting plans: %w", err)
	}

	return plans, nil
}
