package agent

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

type Subscribe interface {
	Subscribe(model.Subscription) (<-chan model.Event, error)
}

type NotificationService interface {
	CreateRule(ctx context.Context, rule model.NotificationRule) error
}

type SubscriptionService interface {
	Create(ctx context.Context, subscription model.Subscription) error
	GetByAgentID(ctx context.Context, agentID string) (model.Subscription, error)
	BindRule(ctx context.Context, subscriptionID string, rule model.NotificationRule) error
	Update(ctx context.Context, subscription model.Subscription) error
}

type Repository interface {
	Create(ctx context.Context, agent model.Agent) error
	GetByID(ctx context.Context, id string) (model.Agent, error)
	GetByUserID(ctx context.Context, userID string) (model.Agent, error)
	Update(ctx context.Context, agent model.Agent) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]model.Agent, error)
}

type UserEventQueue interface {
	SendEvent(ctx context.Context, event string) error
}
