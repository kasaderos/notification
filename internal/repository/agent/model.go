package agent

import "time"

// AgentModel represents the database model for Agent
type AgentModel struct {
	ID               string    `db:"id"`
	UserID           string    `db:"user_id"`
	NotificationType string    `db:"notification_type"`
	Name             string    `db:"name"`
	Envs             string    `db:"envs"`         // JSON string
	Subscription     string    `db:"subscription"` // JSON string
	EventsCount      int       `db:"events_count"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
