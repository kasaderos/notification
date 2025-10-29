package agent

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) Create(ctx context.Context, agent model.Agent) error {
	err := s.repo.Create(ctx, agent)
	if err != nil {
		return fmt.Errorf("failed to create agent: %w", err)
	}

	err = s.subscriptionService.Create(ctx, model.Subscription{
		AgentID: agent.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create subscription: %w", err)
	}

	return nil
}
