package repository

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

// Create creates a new user
func (r *Repository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id, name, agent_id, subscription_id, telegram_id) 
			  VALUES ($1, $2, $3, $4, $5)`

	var agentID, subscriptionID *string
	if user.AgentID != nil {
		agentIDStr := *user.AgentID
		agentID = &agentIDStr
	}

	if user.SubscriptionID != nil {
		subscriptionIDStr := *user.SubscriptionID
		subscriptionID = &subscriptionIDStr
	}

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		agentID,
		subscriptionID,
		user.TelegramID,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
