package config

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fhir-toolbox/rest"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type config struct {
	LogLevel     string       `json:"logLevel"`
	ServerConfig serverConfig `json:"server"`
}

var defaultConfig = config{
	LogLevel:     cmp.Or(os.Getenv("LOG_LEVEL"), "INFO"),
	ServerConfig: defaultServerConfig,
}

func Load(configFile string) (config, error) {
	// first thing? initialize logging
	level := slog.LevelVar{}
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: &level,
	})
	requestContextHandler := rest.NewRequestContextSlogHandler(jsonHandler)
	logger := slog.New(requestContextHandler)
	slog.SetDefault(logger)

	var c = defaultConfig

	if configFile != "" {
		err := loadConfigFile(&c, configFile)
		if err != nil {
			return c, fmt.Errorf("Unable to decode JSON: %w", err)
		}
	}

	// set log level
	switch strings.ToUpper(c.LogLevel) {
	case "DEBUG":
		level.Set(slog.LevelDebug)
	case "INFO":
		level.Set(slog.LevelInfo)
	case "WARN":
		level.Set(slog.LevelWarn)
	case "ERROR":
		level.Set(slog.LevelError)
	default:
		return c, fmt.Errorf("unknown log level: %v", c.LogLevel)
	}

	return c, nil
}

func loadConfigFile(c *config, configFile string) error {
	if configFile != "" {
		configJsonBytes, _ := os.ReadFile(configFile)

		configJsonDecoder := json.NewDecoder(bytes.NewReader(configJsonBytes))
		configJsonDecoder.DisallowUnknownFields()

		err := configJsonDecoder.Decode(c)
		if err != nil {
			return fmt.Errorf("Unable to decode JSON: %w", err)
		}
	}
	return nil
}
