package repository

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// Fetch retrieves users with pagination
func (r *Repository) Fetch(ctx context.Context, limit, offset int) ([]*model.User, error) {
	query := `SELECT id, name, agent_id, subscription_id, telegram_id FROM users ORDER BY telegram_id LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var userModel UserModel
		err := rows.Scan(
			&userModel.ID,
			&userModel.Name,
			&userModel.AgentID,
			&userModel.SubscriptionID,
			&userModel.TelegramID,
		)
		if err != nil {
			return nil, err
		}

		user, err := r.mapToModel(&userModel)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
