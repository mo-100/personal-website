package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string `envconfig:"SERVER_PORT"`

	DatabaseDriverName string `envconfig:"DATABASE_DRIVER_NAME"`
	DatabaseUsername   string `envconfig:"DATABASE_USERNAME"`
	DatabasePassword   string `envconfig:"DATABASE_PASSWORD"`
	DatabaseName       string `envconfig:"DATABASE_NAME"`
	DatabaseURL        string `envconfig:"DATABASE_URL"`
	DatabasePort       string `envconfig:"DATABASE_PORT"`
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
