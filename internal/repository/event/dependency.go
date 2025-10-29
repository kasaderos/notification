package event

import (
	"context"

	"github.com/kasaderos/notification/pkg/elastic"
)

type ElasticClient interface {
	IndexDocument(ctx context.Context, docID string, document any, indexName string) error
	Search(ctx context.Context, indexName string, query any) (*elastic.Response, error)
	DeleteDocument(ctx context.Context, docID string, indexName string) error
}
