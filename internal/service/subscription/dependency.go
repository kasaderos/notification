package subscription

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

type Repository interface {
	Create(ctx context.Context, subscription model.Subscription) error
	GetByID(ctx context.Context, id string) (model.Subscription, error)
	Update(ctx context.Context, subscription model.Subscription) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]model.Subscription, error)
}
