package config

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"log/slog"
	"os"
	"strings"
)

type config struct {
	LogLevel string `json:"logLevel"`
}

func Load() config {
	configFile := flag.String("f", "", "config file path")
	flag.Parse()

	if *configFile == "" {
		log.Fatalf("No config file provided (use -f <config_file>)")
	}

	configJsonBytes, _ := os.ReadFile(*configFile)

	configJsonDecoder := json.NewDecoder(bytes.NewReader(configJsonBytes))
	configJsonDecoder.DisallowUnknownFields()

	var c = config{}

	err := configJsonDecoder.Decode(&c)
	if err != nil {
		log.Fatalf("Unable to decode JSON: %v", err)
	}

	// set log level
	var level slog.Level
	switch strings.ToUpper(c.LogLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	}

	slog.SetLogLoggerLevel(level)

	return c
}
