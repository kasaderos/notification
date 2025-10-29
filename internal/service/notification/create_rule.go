package notification

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) CreateRule(ctx context.Context, rule model.NotificationRule) error {
	return s.ruleRepo.Create(ctx, rule)
}
