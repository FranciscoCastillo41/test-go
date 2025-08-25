package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
	"github.com/FranciscoCastillo41/test-go/backend/internal/repository"
)

// ErrInvalidInput is returned when name/price are invalid.
var ErrInvalidInput = fmt.Errorf("invalid input")

// WidgetService handles business logic for widgets.
type WidgetService struct {
	repo repository.WidgetRepository
}

// NewWidgetService creates a new service using the provided repository.
func NewWidgetService(r repository.WidgetRepository) *WidgetService {
	return &WidgetService{repo: r}
}

// Create validates input, generates an ID and timestamp, then saves a widget.
func (s *WidgetService) Create(ctx context.Context, name string, price float64) (domain.Widget, error) {
	if name == "" || price < 0 {
		return domain.Widget{}, ErrInvalidInput
	}

	w := domain.Widget{
		ID:        uuid.NewString(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().UTC(),
	}
	return s.repo.Create(ctx, w)
}

// ByID fetches a widget by its ID.
func (s *WidgetService) ByID(ctx context.Context, id string) (domain.Widget, error) {
	return s.repo.ByID(ctx, id)
}

// GetAll retrieves all widgets from the repository.
func (s *WidgetService) GetAll(ctx context.Context) ([]domain.Widget, error) {
	return s.repo.GetAll(ctx)
}

// Delete a single widget.
func (s *WidgetService) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidInput
	}
	return s.repo.DeleteByID(ctx, id)
}

// Delete all widgets, returning how many were removed.
func (s *WidgetService) DeleteAll(ctx context.Context) (int64, error) {
	return s.repo.DeleteAll(ctx)
}

// internal/services/widget_service.go
func (s *WidgetService) UpdateByID(ctx context.Context, id string, upd domain.WidgetUpdate) (domain.Widget, error) {
	if id == "" {
		return domain.Widget{}, ErrInvalidInput
	}
	// Require at least one field
	if upd.Name == nil && upd.Price == nil {
		return domain.Widget{}, ErrInvalidInput
	}
	// Validate if provided
	if upd.Name != nil && *upd.Name == "" {
		return domain.Widget{}, ErrInvalidInput
	}
	if upd.Price != nil && *upd.Price < 0 {
		return domain.Widget{}, ErrInvalidInput
	}
	return s.repo.UpdateByID(ctx, id, upd)
}
