package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	client   *Client
	queue    string
	exchange string
}

func NewConsumer(client *Client, queue, exchange string) (*Consumer, error) {
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

	// Declare queue
	_, err = ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind queue to exchange
	err = ch.QueueBind(
		queue,    // queue name
		queue,    // routing key
		exchange, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &Consumer{
		client:   client,
		queue:    queue,
		exchange: exchange,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) (<-chan amqp.Delivery, error) {
	// Create a new channel for consuming to avoid closing the shared channel
	ch, err := c.client.Connection().Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// Set QoS to prefetch only one message at a time
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		ch.Close()
		return nil, fmt.Errorf("failed to set QoS: %w", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		c.queue, // queue
		"",      // consumer tag
		false,   // auto-ack (set to false for manual acknowledgment)
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		ch.Close()
		return nil, fmt.Errorf("failed to register consumer: %w", err)
	}

	deliveries := make(chan amqp.Delivery)

	go func() {
		defer close(deliveries)
		defer ch.Close()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-msgs:
				if !ok {
					return
				}

				select {
				case deliveries <- msg:
					// Message is forwarded to consumer channel
					// Consumer should call Ack/Nack manually
				case <-ctx.Done():
					msg.Nack(false, true) // Reject and requeue on context cancellation
					return
				}
			}
		}
	}()

	return deliveries, nil
}

func (c *Consumer) Close() error {
	return c.client.Close()
}
