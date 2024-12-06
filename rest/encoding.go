package rest

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"slices"
)

type Format string

const (
	FormatJSON Format = "application/fhir+json"
	FormatXML  Format = "application/fhir+xml"
)

var (
	alternateFormatsJSON = []string{"application/json", "text/json", "json"}
	alternateFormatsXML  = []string{"application/xml", "text/xml", "xml"}
)

func detectFormat(r *http.Request, defaultFormat Format) Format {
	// url parameter overrides the Accept header
	formatQuery := r.URL.Query()["_format"]
	if len(formatQuery) > 0 {
		format := matchFormat(formatQuery[0])
		if format != "" {
			return format
		}
	}

	for _, accept := range r.Header["Accept"] {
		format := matchFormat(accept)
		if format != "" {
			return format
		}
	}
	return defaultFormat
}

func matchFormat(requested string) Format {
	switch {
	case requested == string(FormatJSON) || slices.Contains(alternateFormatsJSON, requested):
		return FormatJSON
	case requested == string(FormatXML) || slices.Contains(alternateFormatsXML, requested):
		return FormatXML
	}
	return ""
}

func encodeJSON[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decodeJSON[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func encodeXML[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/fhir+xml")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return fmt.Errorf("encode xml: %w", err)
	}
	if err := xml.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode xml: %w", err)
	}
	return nil
}

func decodeXML[T any](r *http.Request) (T, error) {
	var v T
	if err := xml.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode xml: %w", err)
	}
	return v, nil
}
