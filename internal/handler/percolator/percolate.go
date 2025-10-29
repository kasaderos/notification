package percolator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Handler) PercolateEvents(ctx context.Context) error {
	msgs, err := s.consumer.Consume(ctx)
	if err != nil {
		return fmt.Errorf("failed to consume events: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg := <-msgs:
			var event Event

			if err := json.Unmarshal(msg.Body, &event); err != nil {
				return fmt.Errorf("failed to unmarshal event: %w", err)
			}

			err = s.service.PercolateEvents(ctx, model.Event{
				ID:      event.ID,
				Domain:  event.Domain,
				URL:     event.URL,
				Title:   event.Title,
				Content: event.Content,
			})
			if err != nil {
				return fmt.Errorf("failed to percolate event: %w", err)
			}

			msg.Ack(false)
		}
	}
}
