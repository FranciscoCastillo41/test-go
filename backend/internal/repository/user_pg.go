package repository

import (
	"context"
	"fmt"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PGUserRepo implements UserRepository using PostgreSQL
type PGUserRepo struct {
	pool *pgxpool.Pool
}

// NewPGUserRepo creates a new PostgreSQL user repository
func NewPGUserRepo(pool *pgxpool.Pool) *PGUserRepo {
	return &PGUserRepo{pool: pool}
}

// Create creates a new user record
func (r *PGUserRepo) Create(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	query := `
		INSERT INTO users (auth_user_id, email, full_name, avatar_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id, auth_user_id, email, full_name, avatar_url, created_at, updated_at`

	var user domain.User
	err := r.pool.QueryRow(ctx, query, req.AuthUserID, req.Email, req.FullName, req.AvatarURL).Scan(
		&user.ID,
		&user.AuthUserID,
		&user.Email,
		&user.FullName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

// GetByAuthUserID finds a user by their Supabase Auth user ID
func (r *PGUserRepo) GetByAuthUserID(ctx context.Context, authUserID uuid.UUID) (*domain.User, error) {
	query := `
		SELECT id, auth_user_id, email, full_name, avatar_url, created_at, updated_at
		FROM users 
		WHERE auth_user_id = $1`

	var user domain.User
	err := r.pool.QueryRow(ctx, query, authUserID).Scan(
		&user.ID,
		&user.AuthUserID,
		&user.Email,
		&user.FullName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by auth_user_id %s: %w", authUserID, err)
	}

	return &user, nil
}

// Update updates user profile data
func (r *PGUserRepo) Update(ctx context.Context, authUserID uuid.UUID, req domain.UpdateUserRequest) (*domain.User, error) {
	query := `
		UPDATE users 
		SET full_name = $2, avatar_url = $3, updated_at = now()
		WHERE auth_user_id = $1
		RETURNING id, auth_user_id, email, full_name, avatar_url, created_at, updated_at`

	var user domain.User
	err := r.pool.QueryRow(ctx, query, authUserID, req.FullName, req.AvatarURL).Scan(
		&user.ID,
		&user.AuthUserID,
		&user.Email,
		&user.FullName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update user %s: %w", authUserID, err)
	}

	return &user, nil
}

// Delete removes a user record
func (r *PGUserRepo) Delete(ctx context.Context, authUserID uuid.UUID) error {
	query := `DELETE FROM users WHERE auth_user_id = $1`
	
	result, err := r.pool.Exec(ctx, query, authUserID)
	if err != nil {
		return fmt.Errorf("failed to delete user %s: %w", authUserID, err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with auth_user_id %s not found", authUserID)
	}

	return nil
}