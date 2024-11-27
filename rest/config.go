package rest

import (
	"time"
)

type Config struct {
	// Base path of the FHIR server, used e.g. for building bundles.
	Base string `json:"base"`
	// Timezone used for parsing date parameters without timezones.
	// Defaults to current server timezone.
	Timezone string `json:"timezone"`
	// Date of last capability change, used for building the CapabilityStatement.
	// Defaults to current time.
	Date string `json:"date"`
	// MaxCount of search bundle entries.
	// Defaults to 500.
	MaxCount int `json:"maxCount"`
	// DefaultCount of search bundle entries.
	// Defaults to 500.
	DefaultCount int `json:"defaultCount"`
	// DefaultFormat of the server.
	// Defaults to JSON.
	DefaultFormat Format `json:"defaultFormat"`
}

var DefaultConfig = Config{
	Timezone:      time.Now().Location().String(),
	Date:          time.Now().Format(time.RFC3339),
	MaxCount:      500,
	DefaultCount:  500,
	DefaultFormat: FormatJSON,
}
