package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

// Update updates an existing notification rule
func (r *Repository) Update(ctx context.Context, rule model.NotificationRule) error {
	ruleJSON, _ := json.Marshal(rule.Rule)

	query := `UPDATE notification_rules SET rule = $2, updated_at = $3, version = $4 WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query,
		rule.ID,
		string(ruleJSON),
		rule.UpdatedAt.UnixMilli(),
		rule.Version,
	)
	if err != nil {
		return fmt.Errorf("failed to update notification rule: %w", err)
	}

	return nil
}
