package crawler

import (
	"context"
	"time"

	"github.com/kasaderos/notification/internal/model"
)

type SourceService interface {
	Fetch(ctx context.Context) ([]model.Source, error)
}

type Crawler interface {
	Crawl(ctx context.Context, source model.Source) ([]model.Event, error)
}

type Publisher interface {
	PublishWithTTL(ctx context.Context, event any, ttl time.Duration) error
}
