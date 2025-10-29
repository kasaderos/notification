package source

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

type Repository interface {
	Fetch(ctx context.Context) ([]model.Source, error)
}
