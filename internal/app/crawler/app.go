package crawler

import (
	"context"
	"fmt"
	"log"

	config "github.com/kasaderos/notification/internal/config/crawler"
	crawlerHandler "github.com/kasaderos/notification/internal/handler/crawler"
	sourceRepository "github.com/kasaderos/notification/internal/repository/source"
	"github.com/kasaderos/notification/internal/service/crawler"
	"github.com/kasaderos/notification/pkg/postgres"
	"github.com/kasaderos/notification/pkg/rabbitmq"
	_ "github.com/lib/pq"
)

func Run() error {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	rabbitmqClient, err := rabbitmq.NewClient(cfg.RabbitMQ.URL)
	if err != nil {
		return fmt.Errorf("failed to create rabbitmq client: %w", err)
	}

	crawlerClient := crawler.NewClient()

	db, err := postgres.NewClient(cfg.Postgres.DSN())
	if err != nil {
		return fmt.Errorf("failed to create postgres client: %w", err)
	}

	publisher, err := rabbitmq.NewPublisher(
		rabbitmqClient,
		cfg.RabbitMQ.Exchange,
		cfg.RabbitMQ.RoutingKey,
	)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %w", err)
	}

	sourceRepo := sourceRepository.New(db)

	// Initialize event collection service
	eventCollectionService := crawlerHandler.New(
		sourceRepo,
		crawlerClient,
		publisher,
		cfg.CollectWorkers,
		cfg.EventTTL,
	)

	log.Println("Starting crawler...")

	// Run the crawler once
	if err := eventCollectionService.CollectEvents(ctx); err != nil {
		return err
	}

	log.Println("Crawling completed successfully")

	return nil
}
