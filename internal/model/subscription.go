package model

import (
	"time"
)

// Subscription represents a subscription to events events
type Subscription struct {
	ID        string
	AgentID   string
	Keywords  []string
	Sources   []Source
	Enabled   bool
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Source struct {
	ID     string
	Domain string

	CreatedAt time.Time
}
