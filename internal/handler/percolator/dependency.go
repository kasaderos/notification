package percolator

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/kasaderos/notification/internal/model"
)

type EventService interface {
	PercolateEvents(ctx context.Context, event model.Event) error
}

type Consumer interface {
	Consume(ctx context.Context) (<-chan amqp.Delivery, error)
}
