# Events Notification Telegram Bot

A Telegram bot for creating and managing events notification rules using Elasticsearch percolator queries.

## Features

- Create notification rules with domain and keyword filters
- Update and delete existing rules
- List user's notification rules
- Elasticsearch percolator for efficient events filtering
- PostgreSQL for data persistence

## Architecture

The bot follows a clean architecture pattern with:

- **Models**: Core business entities (Agent, User, NotificationRule, etc.)
- **Repositories**: Data access layer (PostgreSQL + Elasticsearch)
- **Services**: Business logic layer
- **Handlers**: Telegram bot command handlers

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Start PostgreSQL and Elasticsearch (using the provided `compose.yaml`)
4. Run database migrations
5. Build and run the bot

## Environment Variables

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token from BotFather
- `POSTGRES_DSN`: PostgreSQL connection string
- `ELASTICSEARCH_URL`: Elasticsearch server URL

## Commands

- `/start` - Welcome message and available commands
- `/create_rule` - Create a new notification rule
- `/update_rule` - Update an existing rule
- `/delete_rule` - Delete a rule
- `/list_rules` - List your rules

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

# Build application
make build

# Run application
make run

# Build Docker image
make docker-build
```

## Database Schema

The application uses PostgreSQL with the following main tables:
- `users` - Telegram users
- `agents` - Events agents bound to users
- `notification_rules` - User-defined notification rules
- `subscriptions` - Agent subscriptions to events sources

## Elasticsearch Integration

The bot uses Elasticsearch percolator queries to efficiently filter events articles based on user-defined rules. This allows for real-time matching of incoming events against stored queries.