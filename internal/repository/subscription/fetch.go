package repository

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// GetByID retrieves a subscription by ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.Subscription, error) {
	query := `SELECT agent_id, keywords, sources, ttl, created_at, updated_at FROM subscriptions WHERE id = $1`

	var subscriptionModel SubscriptionModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&subscriptionModel.AgentID,
		&subscriptionModel.Keywords,
		&subscriptionModel.Sources,
		&subscriptionModel.Enabled,
		&subscriptionModel.ExpiredAt,
		&subscriptionModel.CreatedAt,
		&subscriptionModel.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return r.mapToModel(&subscriptionModel)
}
