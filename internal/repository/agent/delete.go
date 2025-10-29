package agent

import "context"

var deleteQuery = `DELETE FROM agents WHERE id = $1`

// Delete deletes an agent by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, deleteQuery, id)
	return err
}
