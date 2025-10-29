package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/kasaderos/notification/internal/model"
)

func (r *Repository) UpdatePercolation(ctx context.Context, id string, rule model.NotificationRule) error {
	repoRule := map[string]any{
		"query": map[string]any{
			"bool": map[string]any{
				"must": []map[string]any{
					{
						"term": map[string]any{
							"domain": rule.Rule.Domains(),
						},
					},
					{
						"match": map[string]any{
							"content": strings.Join(rule.Rule.Keywords, " "),
						},
					},
				},
			},
		},
		"domain":    rule.Rule.Domains(),
		"user_id":   rule.UserID,
		"version":   rule.Version,
		"updatedAt": rule.UpdatedAt.UnixMilli(),
	}

	if err := r.client.IndexDocument(
		ctx,
		id,
		repoRule,
		r.eventPercolationIndex,
	); err != nil {
		return fmt.Errorf("failed to update percolation rule: %w", err)
	}

	return nil
}
