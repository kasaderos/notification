package repository

import (
	"context"
	"encoding/json"

	"github.com/kasaderos/notification/internal/model"
)

// Update updates an existing subscription
func (r *Repository) Update(ctx context.Context, subscription *model.Subscription) error {
	keywordsJSON, _ := json.Marshal(subscription.Keywords)
	sourcesJSON, _ := json.Marshal(subscription.Sources)

	query := `UPDATE subscriptions SET keywords = $2, sources = $3, ttl = $4, updated_at = $5 WHERE agent_id = $1`
	_, err := r.db.ExecContext(ctx, query,
		subscription.AgentID,
		string(keywordsJSON),
		string(sourcesJSON),
		subscription.Enabled,
		subscription.ExpiredAt,
		subscription.UpdatedAt,
	)
	return err
}
