package repository

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

// Update updates an existing user
func (r *Repository) Update(ctx context.Context, user *model.User) error {
	var agentID, subscriptionID *string
	if user.AgentID != nil {
		agentIDStr := *user.AgentID
		agentID = &agentIDStr
	}

	if user.SubscriptionID != nil {
		subscriptionIDStr := *user.SubscriptionID
		subscriptionID = &subscriptionIDStr
	}

	query := `UPDATE users SET name = $2, agent_id = $3, subscription_id = $4 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		agentID,
		subscriptionID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
