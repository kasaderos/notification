package event

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

type NotificationService interface {
	SendNotification(ctx context.Context, userID string, notification model.Notification) error
}

type Repository interface {
	SaveEvent(ctx context.Context, event model.Event) error
	PercolateEvent(ctx context.Context, event model.Event) ([]string, error)
}
