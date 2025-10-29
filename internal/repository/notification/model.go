package notification

import (
	"database/sql"
	"time"
)

// NotificationRuleModel represents the database model for NotificationRule
type NotificationRuleModel struct {
	ID        string    `db:"id"`
	Rule      string    `db:"rule"` // JSON string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Version   int       `db:"version"`
}

// Repository implements NotificationRuleRepository interface
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new notification rule repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
