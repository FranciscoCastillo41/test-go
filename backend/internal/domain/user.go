package domain

import (
	"time"

	"github.com/google/uuid"
)

// User is the core domain model. Pure Go (no framework/DB types).
type User struct {
	ID           uuid.UUID  `json:"id"`
	AuthUserID   uuid.UUID  `json:"auth_user_id"` // Supabase Auth user ID (from JWT 'sub')
	Email        string     `json:"email"`
	FullName     *string    `json:"full_name"`    // Optional
	AvatarURL    *string    `json:"avatar_url"`   // Optional  
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// CreateUserRequest represents the data needed to create a new user
type CreateUserRequest struct {
	AuthUserID uuid.UUID `json:"auth_user_id"`
	Email      string    `json:"email"`
	FullName   *string   `json:"full_name"`
	AvatarURL  *string   `json:"avatar_url"`
}

// UpdateUserRequest represents the data that can be updated for a user
type UpdateUserRequest struct {
	FullName  *string `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
}


