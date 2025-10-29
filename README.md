# Events Notification System

A distributed system for collecting, processing, and delivering event notifications using Telegram bot, Elasticsearch percolator queries, and RabbitMQ messaging.

## Features

- **Telegram Bot**: Create and manage notification rules with domain and keyword filters
- **Event Crawler**: Collect events from various sources
- **Event Percolator**: Process events using Elasticsearch percolator queries
- **RabbitMQ Integration**: Reliable message queuing between services
- **PostgreSQL**: Data persistence for users, rules, and subscriptions
- **Elasticsearch**: Efficient event filtering and search

## Architecture

The system consists of three microservices:

- **tgbotapi**: Telegram bot for user interaction and rule management
- **crawler**: Event collection service that crawls sources and publishes to RabbitMQ
- **percolator**: Event processing service that consumes from RabbitMQ and matches against rules

### Clean Architecture Pattern

- **Models**: Core business entities (Agent, User, NotificationRule, Event, etc.)
- **Repositories**: Data access layer (PostgreSQL + Elasticsearch)
- **Services**: Business logic layer
- **Handlers**: Service-specific handlers (Telegram, Event processing)

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Start infrastructure services (PostgreSQL, Elasticsearch, RabbitMQ) using `compose.yaml`
4. Run database migrations
5. Build and run the services

## Environment Variables

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token from BotFather
- `POSTGRES_DSN`: PostgreSQL connection string
- `ELASTICSEARCH_URL`: Elasticsearch server URL
- `RABBITMQ_URL`: RabbitMQ connection string (e.g., `amqp://guest:guest@localhost:5672/`)

## Services

### Telegram Bot (tgbotapi)
- `/start` - Welcome message and available commands
- `/create_rule` - Create a new notification rule
- `/update_rule` - Update an existing rule
- `/delete_rule` - Delete a rule
- `/list_rules` - List your rules

### Event Crawler (crawler)
- Collects events from configured sources
- Publishes events to RabbitMQ for processing
- Configurable crawling intervals and retry logic

### Event Percolator (percolator)
- Consumes events from RabbitMQ
- Matches events against user-defined rules using Elasticsearch percolator
- Sends notifications via Telegram bot

## Example Usage

```
/create_rule domains:tengrievents.kz keywords:Iran,Kazakhstan prompt:Events about Iran and Kazakhstan
```

## Development

```bash
# Install dependencies
make deps

# Run linter
make lint

# Build all services
make build

# Run individual services
make run-tgbotapi
make run-crawler
make run-percolator

# Build Docker image
make docker-build

# Run with Docker Compose
docker-compose up -d
```

## Data Flow

```
Event Sources → Crawler → RabbitMQ → Percolator → Elasticsearch → Telegram Bot → Users
     ↑                                                                    ↓
User Commands → Telegram Bot → PostgreSQL → Rule Management
```

## Database Schema

PostgreSQL tables:
- `users` - Telegram users
- `agents` - Event agents bound to users
- `notification_rules` - User-defined notification rules
- `subscriptions` - Agent subscriptions to event sources
- `sources` - Event sources configuration

## Message Queue

RabbitMQ is used for reliable message delivery between services:
- **Event Queue**: Events from crawler to percolator
- **TTL Support**: Configurable message expiration
- **Persistent Messages**: Survive broker restarts

## Elasticsearch Integration

The percolator service uses Elasticsearch percolator queries to efficiently filter events based on user-defined rules. This allows for real-time matching of incoming events against stored queries without scanning all documents.