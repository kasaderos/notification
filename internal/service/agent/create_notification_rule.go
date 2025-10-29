package agent

import (
	"context"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) CreateNotificationRule(
	ctx context.Context,
	agentID string,
	rule model.NotificationRule,
) error {
	subscription, err := s.subscriptionService.GetByAgentID(ctx, agentID)
	if err != nil {
		return fmt.Errorf("failed to get subscription: %w", err)
	}

	err = s.notificationService.CreateRule(ctx, rule)
	if err != nil {
		return fmt.Errorf("failed to create notification rule: %w", err)
	}

	err = s.subscriptionService.BindRule(ctx, subscription.ID, rule)
	if err != nil {
		return fmt.Errorf("failed to bind rule: %w", err)
	}

	return nil
}
