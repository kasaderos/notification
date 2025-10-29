package agent

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) StartNotification(ctx context.Context, agentID string) error {
	subscription, err := s.subscriptionService.GetByAgentID(ctx, agentID)
	if err != nil {
		return fmt.Errorf("failed to get subscription: %w", err)
	}

	subscription.Enabled = true

	err = s.subscriptionService.Update(ctx, subscription)
	if err != nil {
		return fmt.Errorf("failed to update subscription: %w", err)
	}

	err = s.userEventQueue.SendEvent(ctx, model.EventStartNotify)
	if err != nil {
		return fmt.Errorf("failed to send event: %w", err)
	}

	return nil
}
