package notification

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
)

// mapToModel converts NotificationRuleModel to model.NotificationRule
func (r *Repository) mapToModel(ruleModel *NotificationRuleModel) (*model.NotificationRule, error) {
	ruleID, err := uuid.Parse(ruleModel.ID)
	if err != nil {
		return nil, err
	}

	var rule model.Rule
	if err := json.Unmarshal([]byte(ruleModel.Rule), &rule); err != nil {
		return nil, err
	}

	return &model.NotificationRule{
		ID:        ruleID.String(),
		Rule:      rule,
		CreatedAt: ruleModel.CreatedAt,
		UpdatedAt: ruleModel.UpdatedAt,
		Version:   ruleModel.Version,
	}, nil
}
