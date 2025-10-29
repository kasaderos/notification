package event

// Repository implements EventsRepository interface
type Repository struct {
	client ElasticClient

	eventIndex            string
	eventPercolationIndex string
}

// NewRepository creates a new events repository
func New(
	client ElasticClient,
	eventIndex string,
	eventPercolationIndex string,
) *Repository {
	return &Repository{
		client:                client,
		eventIndex:            eventIndex,
		eventPercolationIndex: eventPercolationIndex,
	}
}
