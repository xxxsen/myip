package config

import (
	"encoding/json"
	"os"

	"github.com/xxxsen/common/logger"
)

type Config struct {
	Bind      string           `json:"bind"`
	Headers   []string         `json:"headers"`
	LogConfig logger.LogConfig `json:"log_config"`
}

func Parse(f string) (*Config, error) {
	raw, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var conf Config
	if err := json.Unmarshal(raw, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
