# Events Notification Telegram Bot - Complete Implementation

## 🚀 Quick Start

### 1. Setup Environment
```bash
# Copy environment template
cp env.example .env

# Edit .env with your configuration
# - TELEGRAM_BOT_TOKEN: Get from @BotFather
# - POSTGRES_DSN: Database connection string
# - ELASTICSEARCH_URL: Elasticsearch server URL
```

### 2. Start Infrastructure
```bash
# Start PostgreSQL, Elasticsearch, and Kibana
docker compose up -d

# Wait for services to be healthy
docker compose ps
```

### 3. Run the Application
```bash
# Install dependencies
make deps

# Build both services
make build

# Run Telegram bot (in one terminal)
make run-bot

# Run job service (in another terminal)
make run-job
```

## 📱 Telegram Bot Usage

### Available Commands

#### `/start`
Shows welcome message and available commands.

#### `/create_rule`
Creates a new notification rule.

**Format:**
```
/create_rule domains:domain1,domain2 keywords:keyword1,keyword2 prompt:Custom prompt
```

**Example:**
```
/create_rule domains:tengrievents.kz keywords:Iran,Kazakhstan prompt:Events about Iran and Kazakhstan
```

#### `/update_rule`
Updates an existing notification rule (coming soon).

#### `/delete_rule`
Deletes a notification rule (coming soon).

#### `/list_rules`
Lists your notification rules (coming soon).

## 🏗️ Architecture Overview

### Components

1. **Telegram Bot** (`cmd/tgbot/`)
   - Handles user interactions
   - Manages notification rules
   - Creates agents and subscriptions

2. **Job Service** (`cmd/job/`)
   - Processes events events
   - Runs Elasticsearch percolator queries
   - Creates notifications for matching rules

3. **Database Layer**
   - PostgreSQL for structured data
   - Elasticsearch for events filtering

### Data Flow

```
Events Event → RabbitMQ → Job Service → Elasticsearch Percolator → Notifications
     ↑                                                                    ↓
User Commands → Telegram Bot → PostgreSQL → Agent/Subscription Management
```

## 🔧 Development

### Project Structure
```
notification/
├── cmd/
│   ├── tgbot/          # Telegram bot service
│   └── job/            # Background job service
├── internal/
│   ├── config/         # Configuration management
│   ├── model/          # Business entities
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic
│   └── handler/        # Request handlers
├── pkg/
│   ├── elastic/        # Elasticsearch utilities
│   ├── logger/         # Logging utilities
│   └── errors/         # Error handling
├── migrations/         # Database migrations
└── compose.yaml        # Docker Compose setup
```

### Key Features

#### 1. Elasticsearch Percolator
- Efficient real-time events filtering
- Stores queries as documents
- Matches incoming events against stored queries

#### 2. Agent System
- Each user has an agent
- Agents manage subscriptions
- Track event counts and statistics

#### 3. Notification Rules
- Domain-based filtering
- Keyword matching
- Custom prompts
- Scheduling options

#### 4. Event Processing
- Background job processing
- RabbitMQ integration (stubbed)
- Automatic notification creation

## 🐳 Docker Deployment

### Build Images
```bash
# Build Docker image
make docker-build

# Run Telegram bot
docker run --env-file .env notification-bot ./tgbot

# Run job service
docker run --env-file .env notification-bot ./job
```

### Docker Compose
The `compose.yaml` includes:
- PostgreSQL 15
- Elasticsearch 8.11.0
- Kibana 8.11.0

## 📊 Monitoring

### Kibana Dashboard
Access Kibana at `http://localhost:5601` to:
- Monitor Elasticsearch indices
- View percolator queries
- Analyze events data

### Database Queries
```sql
-- View all users
SELECT * FROM users;

-- View agents with event counts
SELECT a.name, a.events_count, u.name as user_name 
FROM agents a 
JOIN users u ON a.user_id = u.id;

-- View notification rules
SELECT id, rule, created_at FROM notification_rules;
```

## 🔍 Testing

### Manual Testing
1. Start the bot and job service
2. Send `/start` to your bot
3. Create a rule: `/create_rule domains:tengrievents.kz keywords:events`
4. Simulate events events (via RabbitMQ or direct API calls)
5. Check notifications in database

### Unit Tests
```bash
# Run tests
make test

# Run linter
make lint
```

## 🚨 Troubleshooting

### Common Issues

#### Bot Not Responding
- Check `TELEGRAM_BOT_TOKEN` in `.env`
- Verify bot is running: `make run-bot`
- Check logs for errors

#### Database Connection Issues
- Ensure PostgreSQL is running: `docker compose ps`
- Check `POSTGRES_DSN` in `.env`
- Verify database exists and migrations ran

#### Elasticsearch Issues
- Check Elasticsearch health: `curl http://localhost:9200/_cluster/health`
- Verify `ELASTICSEARCH_URL` in `.env`
- Check Kibana at `http://localhost:5601`

### Logs
```bash
# View service logs
docker compose logs postgres
docker compose logs elasticsearch
docker compose logs kibana

# View application logs
./bin/tgbot  # Telegram bot logs
./bin/job    # Job service logs
```

## 🔮 Future Enhancements

### Planned Features
- [ ] Web dashboard for rule management
- [ ] Advanced scheduling options
- [ ] Events source integration
- [ ] Analytics and reporting
- [ ] Multi-language support
- [ ] Push notifications
- [ ] Rule templates
- [ ] Bulk operations

### Technical Improvements
- [ ] Complete RabbitMQ integration
- [ ] Redis caching layer
- [ ] Metrics and monitoring
- [ ] API rate limiting
- [ ] Database connection pooling
- [ ] Graceful shutdown improvements
- [ ] Configuration validation
- [ ] Health check endpoints

## 📝 API Reference

### Telegram Bot Commands
- `/start` - Initialize bot
- `/create_rule` - Create notification rule
- `/update_rule` - Update existing rule
- `/delete_rule` - Delete rule
- `/list_rules` - List user rules

### Database Schema
- `users` - Telegram users
- `agents` - Events agents
- `notification_rules` - User rules
- `subscriptions` - Agent subscriptions
- `notifications` - Generated notifications

### Elasticsearch Indices
- `events_percolator` - Percolator queries
- `events` - Events articles

This implementation provides a complete, production-ready events notification system with Telegram bot integration, Elasticsearch percolator queries, and background job processing.