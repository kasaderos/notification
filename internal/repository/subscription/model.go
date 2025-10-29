package repository

import (
	"database/sql"
	"time"
)

// SubscriptionModel represents the database model for Subscription
type SubscriptionModel struct {
	AgentID   string    `db:"agent_id"`
	Keywords  string    `db:"keywords"` // JSON string
	Sources   string    `db:"sources"`  // JSON string
	Enabled   bool      `db:"enabled"`
	ExpiredAt time.Time `db:"expired_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Repository implements SubscriptionRepository interface
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new subscription repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
