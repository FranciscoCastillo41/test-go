package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
)

// PGWidgetRepo implements WidgetRepository using Postgres (pgxpool).
type PGWidgetRepo struct {
	pool *pgxpool.Pool
}

// NewPGWidgetRepo wires a new repo with a shared connection pool.
func NewPGWidgetRepo(pool *pgxpool.Pool) *PGWidgetRepo {
	return &PGWidgetRepo{pool: pool}
}

// Create inserts a new widget row. It returns the same widget on success.
func (r *PGWidgetRepo) Create(ctx context.Context, w domain.Widget) (domain.Widget, error) {
	const q = `
		INSERT INTO widgets (id, name, price, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.pool.Exec(ctx, q, w.ID, w.Name, w.Price, w.CreatedAt.UTC())
	return w, err
}

// ByID loads a widget by its primary key. If not found, returns ErrNotFound.
func (r *PGWidgetRepo) ByID(ctx context.Context, id string) (domain.Widget, error) {
	const q = `
		SELECT id, name, price, created_at
		FROM widgets
		WHERE id = $1
	`

	var w domain.Widget
	var createdAt time.Time

	row := r.pool.QueryRow(ctx, q, id)
	if err := row.Scan(&w.ID, &w.Name, &w.Price, &createdAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Widget{}, ErrNotFound
		}
		return domain.Widget{}, err
	}

	w.CreatedAt = createdAt.UTC()
	return w, nil
}

// GetAll retrieves all widgets from the database, ordered by creation date descending.
func (r *PGWidgetRepo) GetAll(ctx context.Context) ([]domain.Widget, error) {
	const q = `
		SELECT id, name, price, created_at
		FROM widgets
		ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var widgets []domain.Widget
	for rows.Next() {
		var w domain.Widget
		var createdAt time.Time

		if err := rows.Scan(&w.ID, &w.Name, &w.Price, &createdAt); err != nil {
			return nil, err
		}

		w.CreatedAt = createdAt.UTC()
		widgets = append(widgets, w)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return widgets, nil
}

// DeleteByID delets a single widget by ID.
func (r *PGWidgetRepo) DeleteByID(ctx context.Context, id string) error {
	const q = `DELETE FROM widgets WHERE id = $1`

	tag, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// DeleteAll deletes all widgets and returns the number of rows deleted.
func (r *PGWidgetRepo) DeleteAll(ctx context.Context) (int64, error) {
	const q = `DELETE FROM widgets`
	tag, err := r.pool.Exec(ctx, q)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func (r *PGWidgetRepo) UpdateByID(ctx context.Context, id string, upd domain.WidgetUpdate) (domain.Widget, error) {
	const q = `
        UPDATE widgets
        SET name  = COALESCE($2, name),
            price = COALESCE($3, price)
        WHERE id = $1
        RETURNING id, name, price, created_at
    `
	var w domain.Widget
	var createdAt time.Time

	err := r.pool.QueryRow(ctx, q, id, upd.Name, upd.Price).
		Scan(&w.ID, &w.Name, &w.Price, &createdAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Widget{}, ErrNotFound
	}
	if err != nil {
		return domain.Widget{}, err
	}
	w.CreatedAt = createdAt.UTC()
	return w, nil
}
