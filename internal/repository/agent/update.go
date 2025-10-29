package agent

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kasaderos/notification/internal/model"
)

// Update updates an existing agent
func (r *Repository) Update(ctx context.Context, agent *model.Agent) error {
	envsJSON, _ := json.Marshal(agent.Envs)
	subscriptionJSON, _ := json.Marshal(agent.Subscription)

	query := `UPDATE agents SET 
			  notification_type = $2, name = $3, envs = $4, subscription = $5, 
			  events_count = $6, updated_at = $7 
			  WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query,
		agent.ID,
		agent.NotificationType,
		agent.Name,
		string(envsJSON),
		string(subscriptionJSON),
		agent.EventsCount,
		time.UnixMilli(agent.UpdatedAt),
	)
	return err
}
