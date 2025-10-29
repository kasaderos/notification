package event

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
)

func (r *Repository) SaveEvent(ctx context.Context, event model.Event) error {
	event.ID = uuid.New().String()

	repoEvent := map[string]any{
		"id":        event.ID,
		"domain":    event.Domain,
		"url":       event.URL,
		"title":     event.Title,
		"content":   event.Content,
		"createdAt": time.Now().UnixMilli(),
	}

	err := r.client.IndexDocument(ctx, event.ID, repoEvent, r.eventIndex)
	if err != nil {
		return fmt.Errorf("failed to save event: %w", err)
	}

	return nil
}
