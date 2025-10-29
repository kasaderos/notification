package repository

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
)

// mapToModel converts SubscriptionModel to model.Subscription
func (r *Repository) mapToModel(subscriptionModel *SubscriptionModel) (*model.Subscription, error) {
	agentID, err := uuid.Parse(subscriptionModel.AgentID)
	if err != nil {
		return nil, err
	}

	var keywords []string
	if err := json.Unmarshal([]byte(subscriptionModel.Keywords), &keywords); err != nil {
		return nil, err
	}

	var sources []model.Source
	if err := json.Unmarshal([]byte(subscriptionModel.Sources), &sources); err != nil {
		return nil, err
	}

	return &model.Subscription{
		AgentID:   agentID.String(),
		Keywords:  keywords,
		Sources:   sources,
		Enabled:   subscriptionModel.Enabled,
		ExpiredAt: subscriptionModel.ExpiredAt,
		CreatedAt: subscriptionModel.CreatedAt,
		UpdatedAt: subscriptionModel.UpdatedAt,
	}, nil
}
