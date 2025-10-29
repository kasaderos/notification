package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kasaderos/notification/internal/model"
)

func (r *Repository) CreatePercolation(ctx context.Context, rule model.NotificationRule) error {
	rule.ID = uuid.New().String()

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
		"createdAt": rule.CreatedAt.UnixMilli(),
		"updatedAt": rule.UpdatedAt.UnixMilli(),
	}

	if err := r.client.IndexDocument(
		ctx,
		rule.ID,
		repoRule,
		r.eventPercolationIndex,
	); err != nil {
		return fmt.Errorf("failed to index percolation rule: %w", err)
	}

	return nil
}
