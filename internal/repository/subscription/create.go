package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

// Create creates a new subscription
func (r *Repository) Create(ctx context.Context, subscription *model.Subscription) error {
	keywordsJSON, _ := json.Marshal(subscription.Keywords)
	sourcesJSON, _ := json.Marshal(subscription.Sources)

	query := `INSERT INTO subscriptions (agent_id, keywords, sources, enabled, expired_at, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query,
		subscription.AgentID,
		string(keywordsJSON),
		string(sourcesJSON),
		subscription.Enabled,
		subscription.ExpiredAt,
		subscription.CreatedAt,
		subscription.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create subscription: %w", err)
	}

	return nil
}
