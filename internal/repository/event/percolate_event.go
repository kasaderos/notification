package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kasaderos/notification/internal/model"
)

func (r *Repository) PercolateEvent(ctx context.Context, event model.Event) ([]string, error) {
	query := map[string]interface{}{
		"percolate": map[string]interface{}{
			"field":    "query",
			"document": event,
		},
	}

	res, err := r.client.Search(ctx, r.eventPercolationIndex, query)
	if err != nil {
		return nil, fmt.Errorf("failed to percolate event: %w", err)
	}

	userIDs := make([]string, 0, len(res.Hits))
	for _, hit := range res.Hits {
		var event Event
		if err := json.Unmarshal(hit.Source, &event); err != nil {
			return nil, fmt.Errorf("failed to unmarshal hit source: %w", err)
		}

		userIDs = append(userIDs, event.UserID)
	}

	return userIDs, nil
}
