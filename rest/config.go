package rest

import (
	"encoding/json"
	"net/url"
	"time"
)

type Config struct {
	// Base path of the FHIR server, used e.g. for building bundles.
	Base *url.URL `json:"-"`
	// Timezone used for parsing date parameters without timezones.
	// Defaults to current server timezone.
	Timezone *time.Location `json:"-"`
	// Date of last capability change, used for building the CapabilityStatement (in ISO 8601/RFC 3339 format).
	// Defaults to current time.
	Date time.Time `json:"-"`
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
	Timezone:      time.Local,
	Date:          time.Now(),
	MaxCount:      500,
	DefaultCount:  500,
	DefaultFormat: FormatJSON,
}

type jsonConfig struct {
	Base     string `json:"base"`
	Timezone string `json:"timezone"`
	Date     string `json:"date"`
	jsonConfigInner
}

// type alias to prevent infinite recursion when unmarshalling
type jsonConfigInner Config

func (c Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonConfig{
		Base:            c.Base.String(),
		Timezone:        c.Timezone.String(),
		Date:            c.Date.Format(time.RFC3339),
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

	c.Base, err = url.Parse(jc.Base)
	if err != nil {
		return err
	}

	c.Timezone, err = time.LoadLocation(jc.Timezone)
	if err != nil {
		return err
	}

	if jc.Date != "" {
		c.Date, err = time.Parse(time.RFC3339, jc.Date)
		if err != nil {
			return err
		}
	}

	return nil
}
