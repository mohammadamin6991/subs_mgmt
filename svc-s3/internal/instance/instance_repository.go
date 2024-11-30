package instance

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type InstanceRepository struct {
	db *sqlx.DB
}

func NewInstanceRepository(db *sqlx.DB) *InstanceRepository {
	return &InstanceRepository{
		db: db,
	}
}

// CreateInstance creates a new instance in the database
func (repo *InstanceRepository) CreateInstance(ctx context.Context, i *Instance) (*Instance, error) {
	query := `INSERT INTO public.instances (plan_id, user_id, is_active, access_key, secret_key, endpoint, region)
              VALUES (:plan_id, :user_id, :is_active, :access_key, :secret_key, :endpoint, :region)
              RETURNING id, created_at, updated_at`

	stmt, args, err := sqlx.Named(query, i)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	stmt = repo.db.Rebind(stmt)

	var result Instance
	err = repo.db.GetContext(ctx, &result, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting instance: %w", err)
	}

	// Return the instance including the generated ID and timestamps
	i.ID = result.ID
	i.CreatedAt = result.CreatedAt
	i.UpdatedAt = result.UpdatedAt

	return i, nil
}

// GetInstance fetches an instance by its ID
func (repo *InstanceRepository) GetInstance(ctx context.Context, id int) (*Instance, error) {
	var i Instance

	query := "SELECT * FROM public.instances WHERE id = ?"
	stmt := repo.db.Rebind(query)

	err := repo.db.GetContext(ctx, &i, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("error getting instance: %w", err)
	}

	return &i, nil
}

// UpdateInstance updates an existing instance
func (repo *InstanceRepository) UpdateInstance(ctx context.Context, i *Instance) (*Instance, error) {
	query := `UPDATE public.instances
              SET plan_id = :plan_id, user_id = :user_id, is_active = :is_active, access_key = :access_key,
                  secret_key = :secret_key, endpoint = :endpoint, region = :region, updated_at = NOW()
              WHERE id = :id
              RETURNING updated_at`

	stmt, args, err := sqlx.Named(query, i)
	if err != nil {
		return nil, fmt.Errorf("error preparing named parameters: %w", err)
	}

	stmt = repo.db.Rebind(stmt)

	var updated Instance
	err = repo.db.GetContext(ctx, &updated, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error updating instance: %w", err)
	}

	i.UpdatedAt = updated.UpdatedAt

	return i, nil
}

// DeleteInstance deletes an instance by its ID
func (repo *InstanceRepository) DeleteInstance(ctx context.Context, id int) error {
	query := "DELETE FROM public.instances WHERE id = ?"

	stmt := repo.db.Rebind(query)

	_, err := repo.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting instance: %w", err)
	}

	return nil
}

// GetAllInstances fetches all instances
func (repo *InstanceRepository) GetAllInstances(ctx context.Context) ([]Instance, error) {
	var instances []Instance

	query := "SELECT * FROM public.instances"
	stmt := repo.db.Rebind(query)

	err := repo.db.SelectContext(ctx, &instances, stmt)
	if err != nil {
		return nil, fmt.Errorf("error getting instances: %w", err)
	}

	return instances, nil
}
