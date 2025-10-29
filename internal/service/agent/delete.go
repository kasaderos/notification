package agent

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) Delete(ctx context.Context, agentID string) error {
	err := s.repo.Delete(ctx, agentID)
	if err != nil {
		return fmt.Errorf("failed to delete agent: %w", err)
	}

	err = s.userEventQueue.SendEvent(ctx, model.EventAgentDeleted)
	if err != nil {
		return fmt.Errorf("failed to send event: %w", err)
	}

	return nil
}
