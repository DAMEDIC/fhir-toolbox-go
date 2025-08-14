package rest

import (
	"encoding/json"
	"time"
)

type Config struct {
	// Timezone used for parsing date parameters without timezones.
	// Defaults to current server timezone.
	Timezone *time.Location `json:"-"`
	// MaxCount of search bundle entries.
	// Defaults to 500.
	MaxCount int `json:"maxCount"`
	// DefaultCount of search bundle entries.
	// Defaults to 500.
	DefaultCount int `json:"defaultCount"`
	// DefaultFormat of the server.
	// Defaults to JSON.
	DefaultFormat Format `json:"defaultFormat"`
	// StrictSearchParameters when true causes the server to return an error
	// if unsupported search parameters are used. When false (default),
	// unsupported search parameters are silently ignored.
	StrictSearchParameters bool `json:"strictSearchParameters"`
}

// DefaultConfig sets reasonable defaults that can be used by the package user.
var DefaultConfig = Config{
	Timezone:               time.Local,
	MaxCount:               500,
	DefaultCount:           500,
	DefaultFormat:          FormatJSON,
	StrictSearchParameters: false,
}

type jsonConfig struct {
	Timezone string `json:"timezone"`
	jsonConfigInner
}

// type alias to prevent infinite recursion when unmarshalling
type jsonConfigInner Config

func (c Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonConfig{
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
	*c = Config(jc.jsonConfigInner)

	c.Timezone, err = time.LoadLocation(jc.Timezone)
	if err != nil {
		return err
	}

	return nil
}
