package source

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

var fetchQuery = `SELECT id, domain, created_at FROM sources`

func (r *Repository) Fetch(ctx context.Context) ([]model.Source, error) {
	rows, err := r.db.QueryContext(ctx, fetchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sources []model.Source
	for rows.Next() {
		var source Source
		err := rows.Scan(&source.ID, &source.Domain, &source.CreatedAt)
		if err != nil {
			return nil, err
		}

		sources = append(sources, model.Source{
			ID:        source.ID,
			Domain:    source.Domain,
			CreatedAt: source.CreatedAt,
		})
	}

	return sources, nil
}
