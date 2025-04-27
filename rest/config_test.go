package rest_test

import (
	"encoding/json"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assert"
	"net/url"
	"testing"
	"time"
)

func TestConfigDefaultValues(t *testing.T) {
	// Test that DefaultConfig has expected values
	defaultConfig := rest.DefaultConfig

	if defaultConfig.MaxCount != 500 {
		t.Errorf("Expected DefaultConfig.MaxCount to be 500, got %d", defaultConfig.MaxCount)
	}

	if defaultConfig.DefaultCount != 500 {
		t.Errorf("Expected DefaultConfig.DefaultCount to be 500, got %d", defaultConfig.DefaultCount)
	}

	// DefaultFormat should be JSON format
	expectedFormat := rest.Format("application/fhir+json")
	if defaultConfig.DefaultFormat != expectedFormat {
		t.Errorf("Expected DefaultConfig.DefaultFormat to be %v, got %v", expectedFormat, defaultConfig.DefaultFormat)
	}

	// Timezone and Date are set dynamically, so we just check they're not nil
	if defaultConfig.Timezone == nil {
		t.Error("Expected DefaultConfig.Timezone to be non-nil")
	}

	// Date should be set to a non-zero time
	if defaultConfig.Date.IsZero() {
		t.Error("Expected DefaultConfig.Date to be non-zero")
	}
}

func TestConfigMarshalJSON(t *testing.T) {
	// Create a config with known values
	baseURL, _ := url.Parse("http://example.com/fhir")
	timezone, _ := time.LoadLocation("UTC")
	date := time.Date(2023, 1, 1, 12, 0, 0, 0, timezone)

	config := rest.Config{
		Base:          baseURL,
		Timezone:      timezone,
		Date:          date,
		MaxCount:      100,
		DefaultCount:  50,
		DefaultFormat: rest.Format("application/fhir+json"),
	}

	// Marshal to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal Config: %v", err)
	}

	assert.JSONEqual(t, `{
		"base": "http://example.com/fhir",
		"timezone": "UTC",
		"date": "2023-01-01T12:00:00Z", 
		"maxCount": 100,
		"defaultCount": 50,
		"defaultFormat": "application/fhir+json"
	}`, string(data))
}

func TestConfigUnmarshalJSON(t *testing.T) {
	jsonData := `{
		"base": "http://example.com/fhir",
		"timezone": "America/New_York",
		"date": "2023-01-01T12:00:00Z",
		"maxCount": 200,
		"defaultCount": 100,
		"defaultFormat": "application/fhir+xml"
	}`

	var config rest.Config
	err := json.Unmarshal([]byte(jsonData), &config)
	if err != nil {
		t.Fatalf("Failed to unmarshal valid JSON: %v", err)
	}

	// Check parsed values
	expectedBase, _ := url.Parse("http://example.com/fhir")
	if config.Base.String() != expectedBase.String() {
		t.Errorf("Expected Base URL %s, got %s", expectedBase.String(), config.Base.String())
	}

	expectedTimezone, _ := time.LoadLocation("America/New_York")
	if config.Timezone.String() != expectedTimezone.String() {
		t.Errorf("Expected Timezone %s, got %s", expectedTimezone.String(), config.Timezone.String())
	}

	expectedDate, _ := time.Parse(time.RFC3339, "2023-01-01T12:00:00Z")
	if !config.Date.Equal(expectedDate) {
		t.Errorf("Expected Date %v, got %v", expectedDate, config.Date)
	}

	if config.MaxCount != 200 {
		t.Errorf("Expected MaxCount 200, got %d", config.MaxCount)
	}

	if config.DefaultCount != 100 {
		t.Errorf("Expected DefaultCount 100, got %d", config.DefaultCount)
	}

	if config.DefaultFormat != rest.FormatXML {
		t.Errorf("Expected DefaultFormat application/fhir+xml, got %v", config.DefaultFormat)
	}
}

func TestConfigUnmarshalJSONErrors(t *testing.T) {
	testCases := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name:    "Invalid JSON",
			json:    `{invalid json`,
			wantErr: true,
		},
		{
			name:    "Invalid Base URL",
			json:    `{"base": "://invalid"}`,
			wantErr: true,
		},
		{
			name:    "Invalid Timezone",
			json:    `{"base": "http://example.com", "timezone": "NonExistentTimezone"}`,
			wantErr: true,
		},
		{
			name:    "Invalid Date Format",
			json:    `{"base": "http://example.com", "timezone": "UTC", "date": "not-a-date"}`,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var config rest.Config
			err := json.Unmarshal([]byte(tc.json), &config)

			if tc.wantErr && err == nil {
				t.Errorf("Expected error for %s, got nil", tc.name)
			} else if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error for %s: %v", tc.name, err)
			}
		})
	}
}
