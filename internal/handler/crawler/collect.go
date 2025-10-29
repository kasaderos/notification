package crawler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
	"github.com/kasaderos/notification/pkg/workerpool"
)

// CollectEvents collects events from all domains in notification rules
func (s *Handler) CollectEvents(ctx context.Context) error {
	sources, err := s.sourceService.Fetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch sources: %w", err)
	}

	jobs := make([]workerpool.Job, len(sources))
	for i, source := range sources {
		jobs[i] = s.createJob(source)
	}

	wp := workerpool.New(s.collectWorkers, jobs)

	err = wp.Run(ctx)
	if err != nil {
		return fmt.Errorf("failed to run worker pool: %w", err)
	}

	return nil
}

func (s *Handler) createJob(source model.Source) workerpool.Job {
	return func(ctx context.Context) error {
		events, err := s.crawler.Crawl(ctx, source)
		if err != nil {
			return fmt.Errorf("failed to crawl source: %w", err)
		}

		for _, event := range events {
			jsonEvent, err := json.Marshal(Event{
				ID:      uuid.New().String(),
				Domain:  event.Domain,
				URL:     event.URL,
				Title:   event.Title,
				Content: event.Content,
			})
			if err != nil {
				return fmt.Errorf("failed to marshal event: %w", err)
			}

			err = s.publisher.PublishWithTTL(ctx, jsonEvent, s.eventTTL)
			if err != nil {
				return fmt.Errorf("failed to publish event: %w", err)
			}
		}

		return nil
	}
}
