package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port   string
	DBPath string
	// later: API keys, SMTp creds, other secrets...
}

func Load() (*Config, error) {

	if os.Getenv("PORT") == "" {
		return nil, fmt.Errorf("PORT not set")
	}

	if os.Getenv("DB_PATH") == "" {
		return nil, fmt.Errorf("DBPath not set")
	}

	return &Config{
		Port:   os.Getenv("PORT"),
		DBPath: os.Getenv("DB_PATH"),
	}, nil
}
