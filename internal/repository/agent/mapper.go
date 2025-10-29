package agent

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
)

// mapToModel converts AgentModel to model.Agent
func (r *Repository) mapToModel(agentModel *AgentModel) (*model.Agent, error) {
	agentID, err := uuid.Parse(agentModel.ID)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(agentModel.UserID)
	if err != nil {
		return nil, err
	}

	var envs []string
	if err := json.Unmarshal([]byte(agentModel.Envs), &envs); err != nil {
		return nil, err
	}

	var subscription model.Subscription
	if err := json.Unmarshal([]byte(agentModel.Subscription), &subscription); err != nil {
		return nil, err
	}

	return &model.Agent{
		ID:               agentID.String(),
		UserID:           userID.String(),
		NotificationType: agentModel.NotificationType,
		Name:             agentModel.Name,
		Envs:             envs,
		Subscription:     subscription,
		EventsCount:      agentModel.EventsCount,
		CreatedAt:        agentModel.CreatedAt.UnixMilli(),
		UpdatedAt:        agentModel.UpdatedAt.UnixMilli(),
	}, nil
}
