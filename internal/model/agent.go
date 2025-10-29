package model

import (
	"time"

	"github.com/google/uuid"
)

// Agent represents a events agent
type Agent struct {
	ID               string       `json:"id" db:"id"`
	UserID           string       `json:"userID" db:"user_id"`
	NotificationType string       `json:"notificationType" db:"notification_type"`
	Name             string       `json:"name" db:"name"`
	Envs             []string     `json:"envs" db:"envs"`
	Subscription     Subscription `json:"subscription" db:"subscription"`
	EventsCount      int          `json:"eventsCount" db:"events_count"`
	CreatedAt        int64        `json:"createdAt" db:"created_at"`
	UpdatedAt        int64        `json:"updatedAt" db:"updated_at"`
}

// CreateNotification creates a new notification for this agent
func (a *Agent) CreateNotification() *Notification {
	return &Notification{
		ID:        uuid.New().String(),
		Type:      a.NotificationType,
		Events:    []Event{},
		CreatedAt: time.Now().UnixMilli(),
	}
}
