package notification

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

type Repository interface {
	Create(ctx context.Context, notification model.Notification) error
	GetByID(ctx context.Context, id string) (model.Notification, error)
	Update(ctx context.Context, notification model.Notification) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]model.Notification, error)
}

type RuleRepository interface {
	Create(ctx context.Context, rule model.NotificationRule) error
	GetByID(ctx context.Context, id string) (model.NotificationRule, error)
	Update(ctx context.Context, rule model.NotificationRule) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]model.NotificationRule, error)
}
