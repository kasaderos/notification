package percolator

import (
	"context"
	"fmt"
	"log"

	config "github.com/kasaderos/notification/internal/config/percolator"
	percolatorHandler "github.com/kasaderos/notification/internal/handler/percolator"
	eventRepository "github.com/kasaderos/notification/internal/repository/event"
	eventService "github.com/kasaderos/notification/internal/service/event"
	"github.com/kasaderos/notification/pkg/elastic"
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

	elasticClient, err := elastic.NewClient(cfg.Elastic.URL)
	if err != nil {
		return fmt.Errorf("failed to create elastic client: %w", err)
	}

	rabbitmqClient, err := rabbitmq.NewClient(cfg.RabbitMQ.URL)
	if err != nil {
		return fmt.Errorf("failed to create rabbitmq client: %w", err)
	}

	consumer, err := rabbitmq.NewConsumer(rabbitmqClient, cfg.RabbitMQ.Queue, cfg.RabbitMQ.Exchange)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}

	eventRepo := eventRepository.New(elasticClient, cfg.EventIndex, cfg.EventPercolationIndex)

	// Initialize event collection service
	eventCollectionService := eventService.New(
		eventRepo,
	)

	eventPercolatorHandler := percolatorHandler.New(
		eventCollectionService,
		consumer,
	)

	log.Println("Starting event percolator...")

	// Run the crawler once
	if err := eventPercolatorHandler.PercolateEvents(ctx); err != nil {
		return err
	}

	log.Println("Event percolator completed successfully")
	return nil
}
