package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	client   *Client
	queue    string
	exchange string
}

func NewPublisher(client *Client, exchange, routingKey string) (*Publisher, error) {
	ch := client.Channel()

	// Declare exchange
	err := ch.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare queue (routing key is typically the queue name in direct exchanges)
	_, err = ch.QueueDeclare(
		routingKey, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind queue to exchange with routing key
	err = ch.QueueBind(
		routingKey, // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &Publisher{
		client:   client,
		queue:    routingKey,
		exchange: exchange,
	}, nil
}

// PublishWithTTL publishes any JSON-serializable object to the queue with a specific TTL
// If ttl is 0, the message will not expire
func (p *Publisher) PublishWithTTL(ctx context.Context, v any, ttl time.Duration) error {
	ch := p.client.Channel()

	body, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %w", err)
	}

	publishing := amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent, // Make message persistent
	}

	// Set expiration if TTL is provided
	if ttl > 0 {
		// Expiration must be a string representing milliseconds
		publishing.Expiration = fmt.Sprintf("%d", ttl.Milliseconds())
	}

	err = ch.PublishWithContext(
		ctx,
		p.exchange, // exchange
		p.queue,    // routing key
		false,      // mandatory
		false,      // immediate
		publishing,
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (p *Publisher) Close() error {
	return p.client.Close()
}
