package repository

import (
	"context"
)

// Delete deletes a subscription by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM subscriptions WHERE agent_id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
