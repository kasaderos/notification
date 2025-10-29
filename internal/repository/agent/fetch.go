package agent

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

var fetchQuery = `SELECT id, user_id, notification_type, name, envs, subscription, events_count, created_at, updated_at 
			  FROM agents WHERE id = $1`

var fetchByUserIDQuery = `SELECT id, user_id, notification_type, name, envs, subscription, events_count, created_at, updated_at 
			  FROM agents WHERE user_id = $1`

// Fetch retrieves an agent by ID
func (r *Repository) Fetch(ctx context.Context, id string) (*model.Agent, error) {
	var agentModel AgentModel

	err := r.db.QueryRowContext(ctx, fetchQuery, id).Scan(
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

	return r.mapToModel(&agentModel)
}

// FetchByUserID retrieves an agent by user ID
func (r *Repository) FetchByUserID(ctx context.Context, userID string) (*model.Agent, error) {

	var agentModel AgentModel
	err := r.db.QueryRowContext(ctx, fetchByUserIDQuery, userID).Scan(
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

	return r.mapToModel(&agentModel)
}
