package repository

import (
	"database/sql"
)

// UserModel represents the database model for User
type UserModel struct {
	ID             string  `db:"id"`
	Name           string  `db:"name"`
	AgentID        *string `db:"agent_id"`
	SubscriptionID *string `db:"subscription_id"`
	TelegramID     int64   `db:"telegram_id"`
}

// Repository implements UserRepository interface
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new user repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
