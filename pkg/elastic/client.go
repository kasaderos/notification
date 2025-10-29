package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9"
)

// Client wraps Elasticsearch client with helper methods
type Client struct {
	client *elasticsearch.Client
}

// NewClient creates a new Elasticsearch client
func NewClient(url string) (*Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{url},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create elastic client: %w", err)
	}

	return &Client{client: client}, nil
}

// IndexDocument indexes a document
func (c *Client) IndexDocument(ctx context.Context, docID string, document any, indexName string) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(document); err != nil {
		return fmt.Errorf("error encoding document: %w", err)
	}

	res, err := c.client.Index(
		indexName,
		&buf,
		c.client.Index.WithDocumentID(docID),
		c.client.Index.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		return fmt.Errorf("error response: %s", string(body))
	}

	return nil
}

// DeleteDocument deletes a document
func (c *Client) DeleteDocument(ctx context.Context, docID string, indexName string) error {
	res, err := c.client.Delete(
		indexName,
		docID,
		c.client.Delete.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}

		return fmt.Errorf("error response: %s", string(body))
	}

	return nil
}

// Hit represents a single search hit from Elasticsearch
type Hit struct {
	Source json.RawMessage `json:"_source"`
}

// Response represents a search response from Elasticsearch
type Response struct {
	Hits       []Hit
	TotalCount int64
}

// Search performs a search query
func (c *Client) Search(ctx context.Context, indexName string, query any) (*Response, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	res, err := c.client.Search(
		c.client.Search.WithIndex(indexName),
		c.client.Search.WithBody(&buf),
		c.client.Search.WithContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("error performing search: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error response: %s", string(body))
	}

	var sr struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []Hit `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		return nil, fmt.Errorf("failed to decode search result: %w", err)
	}

	return &Response{
		Hits:       sr.Hits.Hits,
		TotalCount: sr.Hits.Total.Value,
	}, nil
}
