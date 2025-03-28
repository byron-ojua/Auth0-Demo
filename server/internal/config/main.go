package config

import (
	_ "embed"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

//go:embed config.json
var configJSON []byte

const (
	// MAX_PAGE_SIZE is the maximum page size for calls to the database
	MAX_PAGE_SIZE = 500

	PRODUCTION_ENV = "PROD"
)

// Config is the configuration for the application
type Config struct {
	Auth Auth `json:"auth"`
}

// Auth is the configuration for the authentication
type Auth struct {
	ClientSecret string `json:"client_secret"`
}

// New creates a new Config instance
func New(logger *zap.SugaredLogger) (*Config, error) {
	logger.Info("Setting up configuration")

	var c Config

	err := json.Unmarshal(configJSON, &c)
	if err != nil {
		logger.Errorf("failed to unmarshal config: %v", err)
		return nil, err
	}

	// Get env variable APP_ENV to determine which database to use
	// Default to development
	env := os.Getenv("APP_ENV")

	if env == PRODUCTION_ENV {
		logger.Debug("Running with production configuration")
	} else {
		logger.Debug("Running with development configuration")
	}

	return &c, nil
}
