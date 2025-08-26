package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
	"github.com/FranciscoCastillo41/test-go/backend/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UserService handles user business logic
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateOrGetUser creates a user if they don't exist, or returns existing user
// This is called when a user first signs up via Supabase Auth
func (s *UserService) CreateOrGetUser(ctx context.Context, authUserID uuid.UUID, email string) (*domain.User, error) {
	// First try to get existing user
	user, err := s.repo.GetByAuthUserID(ctx, authUserID)
	if err == nil {
		return user, nil // User already exists
	}
	
	// If user not found, create new user
	if errors.Is(err, pgx.ErrNoRows) {
		req := domain.CreateUserRequest{
			AuthUserID: authUserID,
			Email:      email,
		}
		return s.repo.Create(ctx, req)
	}
	
	// Some other error occurred
	return nil, fmt.Errorf("failed to check existing user: %w", err)
}

// GetUserProfile gets user profile by auth user ID
func (s *UserService) GetUserProfile(ctx context.Context, authUserID uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetByAuthUserID(ctx, authUserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}
	return user, nil
}

// UpdateUserProfile updates user profile information
func (s *UserService) UpdateUserProfile(ctx context.Context, authUserID uuid.UUID, req domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.Update(ctx, authUserID, req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}
	return user, nil
}

// DeleteUser deletes a user account
func (s *UserService) DeleteUser(ctx context.Context, authUserID uuid.UUID) error {
	err := s.repo.Delete(ctx, authUserID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}