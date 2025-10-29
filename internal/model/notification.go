package model

// Notification represents a notification with events
type Notification struct {
	ID        string  `json:"id" db:"id"`
	Type      string  `json:"type" db:"type"`
	Events    []Event `json:"events" db:"events"`
	CreatedAt int64   `json:"createdAt" db:"created_at"`
}
