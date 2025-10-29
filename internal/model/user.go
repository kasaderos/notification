package model

// User represents a user in the system
type User struct {
	ID             string  `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	AgentID        *string `json:"agentID" db:"agent_id"`
	SubscriptionID *string `json:"subscriptionID" db:"subscription_id"`
	TelegramID     int64   `json:"telegramID" db:"telegram_id"`
}
