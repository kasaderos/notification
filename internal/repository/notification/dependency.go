package notification

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// NotificationRuleRepository defines the interface for notification rule operations
type NotificationRuleRepository interface {
	Create(ctx context.Context, rule *model.NotificationRule) error
	GetByID(ctx context.Context, id string) (*model.NotificationRule, error)
	Update(ctx context.Context, rule *model.NotificationRule) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]*model.NotificationRule, error)
}
