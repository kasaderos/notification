package crawler

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`

	Crawler  Crawler  `yaml:"crawler"`
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
	Postgres Postgres `yaml:"postgres"`
	Elastic  Elastic  `yaml:"elastic"`

	EventTTL              time.Duration `yaml:"event_ttl" env:"EVENT_TTL" env-default:"1h"`
	EventIndex            string        `yaml:"event_index" env:"EVENT_INDEX" env-default:"event"`
	EventPercolationIndex string        `yaml:"event_percolation_index" env:"EVENT_PERCOLATION_INDEX" env-default:"event_percolation"`
	CollectWorkers        int           `yaml:"collect_workers" env:"COLLECT_WORKERS" env-default:"10"`
}

type RabbitMQ struct {
	URL        string `yaml:"url" env:"RABBITMQ_URL"`
	Exchange   string `yaml:"exchange" env:"RABBITMQ_EXCHANGE"`
	RoutingKey string `yaml:"routing_key" env:"RABBITMQ_ROUTING_KEY"`
}

type Crawler struct {
	Retries int `yaml:"retries" env:"CRAWLER_RETRIES" env-default:"3"`
}

type Postgres struct {
	Database string `yaml:"database" env:"POSTGRES_DATABASE"`
	User     string `yaml:"user" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" env:"POSTGRES_HOST"`
	Port     string `yaml:"port" env:"POSTGRES_PORT"`
	SSLMode  string `yaml:"ssl_mode" env:"POSTGRES_SSL_MODE" env-default:"disable"`
}

func (p *Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.Database, p.SSLMode,
	)
}

type Elastic struct {
	URL string `yaml:"url" env:"ELASTIC_URL"`
}

// LoadConfig loads the configuration using cleanenv.
func LoadConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("configs/event-crawler.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
