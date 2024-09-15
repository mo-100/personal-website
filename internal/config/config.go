package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string `envconfig:"SERVER_PORT" required:"true"`

	Environment string `envconfig:"ENVIRONMENT" required:"true"`

	DatabaseDriverName string `envconfig:"DATABASE_DRIVER_NAME" required:"true"`
	DatabaseUsername   string `envconfig:"DATABASE_USERNAME" required:"true"`
	DatabasePassword   string `envconfig:"DATABASE_PASSWORD" required:"true"`
	DatabaseName       string `envconfig:"DATABASE_NAME" required:"true"`
	DatabaseURL        string `envconfig:"DATABASE_URL" required:"true"`
	DatabasePort       string `envconfig:"DATABASE_PORT" required:"true"`
}

func loadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	var cfg Config
	// Load environment variables from .env file
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		log.Panic(err)
	}
	return cfg
}
