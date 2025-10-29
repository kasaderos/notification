package repository

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// GetByID retrieves a user by ID
func (r *Repository) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT id, name, agent_id, subscription_id, telegram_id FROM users WHERE id = $1`

	var userModel UserModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&userModel.ID,
		&userModel.Name,
		&userModel.AgentID,
		&userModel.SubscriptionID,
		&userModel.TelegramID,
	)
	if err != nil {
		return nil, err
	}

	return r.mapToModel(&userModel)
}

// GetByTelegramID retrieves a user by Telegram ID
func (r *Repository) GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error) {
	query := `SELECT id, name, agent_id, subscription_id, telegram_id FROM users WHERE telegram_id = $1`

	var userModel UserModel
	err := r.db.QueryRowContext(ctx, query, telegramID).Scan(
		&userModel.ID,
		&userModel.Name,
		&userModel.AgentID,
		&userModel.SubscriptionID,
		&userModel.TelegramID,
	)
	if err != nil {
		return nil, err
	}

	return r.mapToModel(&userModel)
}
