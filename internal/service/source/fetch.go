package source

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

func (s *Service) Fetch(ctx context.Context) ([]model.Source, error) {
	return s.repo.Fetch(ctx)
}
