package repository

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// UserRepository defines the interface for user operations
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]*model.User, error)
}
