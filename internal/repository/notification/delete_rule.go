package notification

import (
	"context"
)

// Delete deletes a notification rule by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM notification_rules WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
