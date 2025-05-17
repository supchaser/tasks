package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string // номер порта без двоеточия, например "8000"
	Author string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	author := os.Getenv("AUTHOR")
	if author == "" {
		author = "Unknown Author"
	}

	return &Config{
		Port:   port,
		Author: author,
	}, nil
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%s", c.Port)
}
