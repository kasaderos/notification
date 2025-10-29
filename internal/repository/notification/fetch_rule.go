package notification

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// GetByID retrieves a notification rule by ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.NotificationRule, error) {
	query := `SELECT id, rule, created_at, updated_at, version FROM notification_rules WHERE id = $1`

	var ruleModel NotificationRuleModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&ruleModel.ID,
		&ruleModel.Rule,
		&ruleModel.CreatedAt,
		&ruleModel.UpdatedAt,
		&ruleModel.Version,
	)
	if err != nil {
		return nil, err
	}

	return r.mapToModel(&ruleModel)
}
