package rest_test

import (
	"encoding/json"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assert"
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

	// Timezone is set dynamically, so we just check it's not nil
	if defaultConfig.Timezone == nil {
		t.Error("Expected DefaultConfig.Timezone to be non-nil")
	}

	// StrictSearchParameters should default to false
	if defaultConfig.StrictSearchParameters != false {
		t.Errorf("Expected DefaultConfig.StrictSearchParameters to be false, got %v", defaultConfig.StrictSearchParameters)
	}
}

func TestConfigMarshalJSON(t *testing.T) {
	// Create a config with known values
	timezone, _ := time.LoadLocation("UTC")

	config := rest.Config{
		Timezone:               timezone,
		MaxCount:               100,
		DefaultCount:           50,
		DefaultFormat:          rest.Format("application/fhir+json"),
		StrictSearchParameters: true,
	}

	// Marshal to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal Config: %v", err)
	}

	assert.JSONEqual(t, `{
		"timezone": "UTC",
		"maxCount": 100,
		"defaultCount": 50,
		"defaultFormat": "application/fhir+json",
		"strictSearchParameters": true
	}`, string(data))
}

func TestConfigUnmarshalJSON(t *testing.T) {
	jsonData := `{
		"timezone": "America/New_York",
		"maxCount": 200,
		"defaultCount": 100,
		"defaultFormat": "application/fhir+xml",
		"strictSearchParameters": true
	}`

	var config rest.Config
	err := json.Unmarshal([]byte(jsonData), &config)
	if err != nil {
		t.Fatalf("Failed to unmarshal valid JSON: %v", err)
	}

	// Check parsed values
	expectedTimezone, _ := time.LoadLocation("America/New_York")
	if config.Timezone.String() != expectedTimezone.String() {
		t.Errorf("Expected Timezone %s, got %s", expectedTimezone.String(), config.Timezone.String())
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

	if config.StrictSearchParameters != true {
		t.Errorf("Expected StrictSearchParameters true, got %v", config.StrictSearchParameters)
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
			name:    "Invalid Timezone",
			json:    `{"timezone": "NonExistentTimezone"}`,
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
