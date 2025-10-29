package repository

import (
	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
)

// mapToModel converts UserModel to model.User
func (r *Repository) mapToModel(userModel *UserModel) (*model.User, error) {
	userID, err := uuid.Parse(userModel.ID)
	if err != nil {
		return nil, err
	}

	var agentID, subscriptionID string
	if userModel.AgentID != nil {
		agentUUID, err := uuid.Parse(*userModel.AgentID)
		if err != nil {
			return nil, err
		}
		agentID = agentUUID.String()
	}

	if userModel.SubscriptionID != nil {
		subscriptionUUID, err := uuid.Parse(*userModel.SubscriptionID)
		if err != nil {
			return nil, err
		}
		subscriptionID = subscriptionUUID.String()
	}

	return &model.User{
		ID:             userID.String(),
		Name:           userModel.Name,
		AgentID:        emptyToNil(agentID),
		SubscriptionID: emptyToNil(subscriptionID),
		TelegramID:     userModel.TelegramID,
	}, nil
}

func emptyToNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
