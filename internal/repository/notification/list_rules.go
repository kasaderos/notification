package notification

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// Fetch retrieves notification rules with pagination
func (r *Repository) Fetch(ctx context.Context, limit, offset int) ([]*model.NotificationRule, error) {
	query := `SELECT id, rule, created_at, updated_at, version FROM notification_rules ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []*model.NotificationRule
	for rows.Next() {
		var ruleModel NotificationRuleModel
		err := rows.Scan(
			&ruleModel.ID,
			&ruleModel.Rule,
			&ruleModel.CreatedAt,
			&ruleModel.UpdatedAt,
			&ruleModel.Version,
		)
		if err != nil {
			return nil, err
		}

		rule, err := r.mapToModel(&ruleModel)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	return rules, nil
}
