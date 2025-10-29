package agent

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

var listQuery = `SELECT id, user_id, notification_type, name, envs, subscription, events_count, created_at, updated_at 
			  FROM agents ORDER BY created_at DESC LIMIT $1 OFFSET $2`

// List retrieves agents with pagination
func (r *Repository) List(ctx context.Context, limit, offset int) ([]*model.Agent, error) {
	rows, err := r.db.QueryContext(ctx, listQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var agents []*model.Agent
	for rows.Next() {
		var agentModel AgentModel
		err := rows.Scan(
			&agentModel.ID,
			&agentModel.UserID,
			&agentModel.NotificationType,
			&agentModel.Name,
			&agentModel.Envs,
			&agentModel.Subscription,
			&agentModel.EventsCount,
			&agentModel.CreatedAt,
			&agentModel.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		agent, err := r.mapToModel(&agentModel)
		if err != nil {
			return nil, err
		}

		agents = append(agents, agent)
	}

	return agents, nil
}
