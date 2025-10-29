package percolator

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`

	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
	Elastic  Elastic  `yaml:"elastic"`

	EventIndex            string `yaml:"event_index" env:"EVENT_INDEX" env-default:"event"`
	EventPercolationIndex string `yaml:"event_percolation_index" env:"EVENT_PERCOLATION_INDEX" env-default:"event_percolation"`
}

type RabbitMQ struct {
	URL      string `yaml:"url" env:"RABBITMQ_URL"`
	Exchange string `yaml:"exchange" env:"RABBITMQ_EXCHANGE"`
	Queue    string `yaml:"queue" env:"RABBITMQ_QUEUE"`
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
