package event

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) PercolateEvents(ctx context.Context, event model.Event) error {
	userIDs, err := s.repo.PercolateEvent(ctx, event)
	if err != nil {
		return fmt.Errorf("failed to percolate event: %w", err)
	}

	if len(userIDs) == 0 {
		return nil
	}

	// todo: send notification to users
	return nil
}
