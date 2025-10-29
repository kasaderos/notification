package model

import (
	"time"
)

// NotificationRule represents a rule for notifications
type NotificationRule struct {
	ID        string
	Rule      Rule
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID string
}

// Rule contains the notification rule configuration
type Rule struct {
	Sources  []Source
	Keywords []string
	Prompt   string
	Schedule Schedule
}

func (r *Rule) Domains() []string {
	domains := make([]string, len(r.Sources))
	for i, source := range r.Sources {
		domains[i] = source.Domain
	}

	return domains
}

type Schedule struct {
	ClockTime string
}
