package crawler

import "time"

type Handler struct {
	sourceService  SourceService
	crawler        Crawler
	publisher      Publisher
	collectWorkers int
	eventTTL       time.Duration
}

func New(
	sourceService SourceService,
	crawler Crawler,
	publisher Publisher,
	collectWorkers int,
	eventTTL time.Duration,
) *Handler {
	return &Handler{
		sourceService:  sourceService,
		crawler:        crawler,
		publisher:      publisher,
		collectWorkers: collectWorkers,
		eventTTL:       eventTTL,
	}
}
