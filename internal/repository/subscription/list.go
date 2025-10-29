package repository

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// Fetch retrieves subscriptions with pagination
func (r *Repository) Fetch(ctx context.Context, limit, offset int) ([]*model.Subscription, error) {
	query := `SELECT agent_id, keywords, sources, ttl, created_at, updated_at FROM subscriptions ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*model.Subscription
	for rows.Next() {
		var subscriptionModel SubscriptionModel
		err := rows.Scan(
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

		subscription, err := r.mapToModel(&subscriptionModel)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}
