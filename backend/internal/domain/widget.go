package domain

import "time"

// Widget is the core domain model. Pure Go (no framework/DB types).
type Widget struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// For partial updates (PATCH)
type WidgetUpdate struct {
	Name  *string  `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
}
