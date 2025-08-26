package repository

import (
	"context"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
	"github.com/google/uuid"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create creates a new user record
	Create(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error)
	
	// GetByAuthUserID finds a user by their Supabase Auth user ID
	GetByAuthUserID(ctx context.Context, authUserID uuid.UUID) (*domain.User, error)
	
	// Update updates user profile data
	Update(ctx context.Context, authUserID uuid.UUID, req domain.UpdateUserRequest) (*domain.User, error)
	
	// Delete removes a user record
	Delete(ctx context.Context, authUserID uuid.UUID) error
}