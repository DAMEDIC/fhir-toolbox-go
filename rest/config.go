package rest

import (
	"encoding/json"
	"net/url"
	"time"
)

type Config struct {
	// Base path of the FHIR server, used e.g. for building bundles.
	Base *url.URL
	// Timezone used for parsing date parameters without timezones.
	// Defaults to current server timezone.
	Timezone *time.Location
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

// DefaultConfig sets reasonable defaults that can be used by the package user.
var DefaultConfig = Config{
	Timezone:      time.Now().Location(),
	Date:          time.Now().Format(time.RFC3339),
	MaxCount:      500,
	DefaultCount:  500,
	DefaultFormat: FormatJSON,
}

type jsonConfig struct {
	Base     string `json:"base"`
	Timezone string `json:"timezone"`
	jsonConfigInner
}

// type alias to prevent infinite recursion when (un)marshalling
type jsonConfigInner Config

func (c Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonConfig{
		Base:            c.Base.String(),
		Timezone:        c.Timezone.String(),
		jsonConfigInner: jsonConfigInner(c),
	})
}

func (c *Config) UnmarshalJSON(b []byte) error {
	jc := jsonConfig{
		jsonConfigInner: jsonConfigInner(*c),
	}

	err := json.Unmarshal(b, &jc)
	if err != nil {
		return err
	}

	c.Base, err = url.Parse(jc.Base)
	if err != nil {
		return err
	}

	c.Timezone, err = time.LoadLocation(jc.Timezone)
	if err != nil {
		return err
	}

	return nil
}
