package repository

import (
	"context"
	"fmt"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
)

// ErrNotFound is returned when a record is not found in the repository.
var ErrNotFound = fmt.Errorf("not found")

// WidgetRepository defines the storage interface for widgets.
type WidgetRepository interface {
	Create(ctx context.Context, w domain.Widget) (domain.Widget, error)
	ByID(ctx context.Context, id string) (domain.Widget, error)
	GetAll(ctx context.Context) ([]domain.Widget, error)
	DeleteByID(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) (int64, error)
	UpdateByID(ctx context.Context, id string, upd domain.WidgetUpdate) (domain.Widget, error)
}
