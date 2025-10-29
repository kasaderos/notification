package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

// Create creates a new notification rule
func (r *Repository) Create(ctx context.Context, rule *model.NotificationRule) error {
	ruleJSON, _ := json.Marshal(rule.Rule)

	query := `INSERT INTO notification_rules (id, rule, created_at, updated_at, version) 
			  VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, query,
		rule.ID,
		string(ruleJSON),
		rule.CreatedAt,
		rule.UpdatedAt,
		rule.Version,
	)
	if err != nil {
		return fmt.Errorf("failed to create notification rule: %w", err)
	}

	return err
}
