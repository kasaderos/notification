package tgbotapi

// import (
// 	"context"
// 	"database/sql"
// 	"log"
// 	"os/signal"
// 	"syscall"

// 	"github.com/elastic/go-elasticsearch/v8"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// 	"github.com/kasaderos/notification/internal/config"
// 	handler "github.com/kasaderos/notification/internal/handler/tgbot"
// )

// func Run() error {
// 	ctx := context.Background()

// 	// Load configuration
// 	cfg := config.LoadConfig()

// 	// Initialize database connection
// 	db, err := sql.Open("postgres", cfg.PostgresDSN)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	defer db.Close()

// 	// Test database connection
// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Failed to ping database: %v", err)
// 	}

// 	// Initialize Elasticsearch client
// 	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
// 		Addresses: []string{cfg.ElasticsearchURL},
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to create Elasticsearch client: %v", err)
// 	}

// 	// Initialize repositories
// 	userRepo := userR.NewUserRepository(db)
// 	agentRepo := agentR.NewAgentRepository(db)
// 	ruleRepo := ruleR.NewNotificationRuleRepository(db)
// 	subscriptionRepo := subscriptionR.NewSubscriptionRepository(db)
// 	eventsRepo := eventsR.NewEventsRepository(esClient)
// 	eventRepo := eventR.NewEventRepository()

// 	// Initialize services
// 	agentService := service.NewAgentService(agentRepo)
// 	notificationRuleService := service.NewNotificationRuleService(
// 		ruleRepo,
// 		agentRepo,
// 		userRepo,
// 		subscriptionRepo,
// 		eventsRepo,
// 	)

// 	// Initialize Telegram bot
// 	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
// 	if err != nil {
// 		log.Fatalf("Failed to create bot: %v", err)
// 	}

// 	bot.Debug = true
// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	// Initialize bot handler
// 	botHandler := handler.NewTelegramBotHandler(bot, notificationRuleService, userRepo)

// 	// Set up webhook or polling
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)

// 	// Create context for graceful shutdown
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// Handle graceful shutdown
// 	signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

// 	// Process updates
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Bot stopped")
// 			return nil
// 		case update := <-updates:
// 			if update.Message != nil {
// 				botHandler.HandleUpdate(ctx, update)
// 			}
// 		}
// 	}
// }
