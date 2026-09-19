package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenRouterAPIKey string
	OpenRouterModel  string
}

func Load() (Config, error) {
	err := godotenv.Load()

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := Config{
		OpenRouterAPIKey: strings.TrimSpace(os.Getenv("OPENROUTER_API_KEY")),
		OpenRouterModel:  strings.TrimSpace(os.Getenv("OPENROUTER_MODEL")),
	}

	if cfg.OpenRouterAPIKey == "" {
		return Config{}, errors.New("missing OPENROUTER_API_KEY")
	}

	if cfg.OpenRouterModel == "" {
		return Config{}, errors.New("missing OPENROUTER_MODEL")
	}

	return cfg, nil
}
