package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kasaderos/notification/internal/model"
)

var createQuery = `INSERT INTO agents
(id, user_id, notification_type, name, envs, subscription, events_count, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id`

// Create creates a new agent
func (r *Repository) Create(ctx context.Context, agent *model.Agent) error {
	envsJSON, err := json.Marshal(agent.Envs)
	if err != nil {
		return fmt.Errorf("failed to marshal envs: %w", err)
	}

	subscriptionJSON, err := json.Marshal(agent.Subscription)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription: %w", err)
	}

	_, err = r.db.ExecContext(ctx, createQuery,
		agent.ID,
		agent.UserID,
		agent.NotificationType,
		agent.Name,
		string(envsJSON),
		string(subscriptionJSON),
		agent.EventsCount,
		time.UnixMilli(agent.CreatedAt),
		time.UnixMilli(agent.UpdatedAt),
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
