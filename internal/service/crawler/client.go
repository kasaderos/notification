package crawler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/kasaderos/notification/internal/model"
	"github.com/kasaderos/notification/pkg/html"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{client: &http.Client{}}
}

func (c *Client) Crawl(ctx context.Context, source model.Source) ([]model.Event, error) {
	url := fmt.Sprintf("https://%s", source.Domain)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	content := html.ProcessContent(body)
	title := html.ExtractTitle(body)

	return []model.Event{
		{
			Domain:  source.Domain,
			URL:     url,
			Title:   title,
			Content: content,
		},
	}, nil
}
