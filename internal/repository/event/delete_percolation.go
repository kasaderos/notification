package event

import (
	"context"
	"fmt"
)

// Delete deletes an events query (placeholder for future implementation)
func (r *Repository) DeletePercolation(ctx context.Context, id string) error {
	if err := r.client.DeleteDocument(ctx, id, r.eventPercolationIndex); err != nil {
		return fmt.Errorf("failed to delete percolation rule: %w", err)
	}

	return nil
}
