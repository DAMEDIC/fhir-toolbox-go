package rest

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"net/http"
	"slices"
)

type format string

const (
	formatJSON format = "application/fhir+json"
	formatXML  format = "application/fhir+xml"
)

var (
	alternateFormatsJSON = []string{"application/json", "text/json", "json"}
	alternateFormatsXML  = []string{"application/xml", "text/xml", "xml"}
)

func detectFormat(r *http.Request, headerName string, defaultFormat format) format {
	// url parameter overrides the Accept header
	formatQuery := r.URL.Query()["_format"]
	if len(formatQuery) > 0 {
		format := matchFormat(formatQuery[0])
		if format != "" {
			return format
		}
	}

	for _, accept := range r.Header[headerName] {
		format := matchFormat(accept)
		if format != "" {
			return format
		}
	}
	return defaultFormat
}

func matchFormat(requested string) format {
	switch {
	case requested == string(formatJSON) || slices.Contains(alternateFormatsJSON, requested):
		return formatJSON
	case requested == string(formatXML) || slices.Contains(alternateFormatsXML, requested):
		return formatXML
	}
	return ""
}

type decodingError struct {
	encoding string
}

func (e decodingError) Error() string {
	return fmt.Sprintf("error parsing %s body", e.encoding)
}

func (e decodingError) StatusCode() int {
	return 400
}

func (e decodingError) OperationOutcome() basic.OperationOutcome {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{
				Severity:    basic.Code{Value: utils.Ptr("fatal")},
				Code:        basic.Code{Value: utils.Ptr("processing")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
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

func decodeResourceJSON[R model.Release](r *http.Request) (model.Resource, capabilities.FHIRError) {
	var release R
	switch any(release).(type) {
	case model.R4:
		resource, err := decodeJSON[r4.ContainedResource](r)
		return resource.Resource, err
	case model.R4B:
		resource, err := decodeJSON[r4b.ContainedResource](r)
		return resource.Resource, err
	case model.R5:
		resource, err := decodeJSON[r5.ContainedResource](r)
		return resource.Resource, err
	default:
		// This should never happen as long as we control all implementations of the Release interface.
		// This is achieved by sealing the interface. See the interface definition for more information.
		panic("unsupported release")
	}
}

func decodeJSON[T any](r *http.Request) (T, capabilities.FHIRError) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, decodingError{
			encoding: "json",
		}
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

func decodeResourceXML[R model.Release](r *http.Request) (model.Resource, capabilities.FHIRError) {
	var release R
	switch any(release).(type) {
	case model.R4:
		resource, err := decodeXML[r4.ContainedResource](r)
		return resource.Resource, err
	case model.R4B:
		resource, err := decodeXML[r4b.ContainedResource](r)
		return resource.Resource, err
	case model.R5:
		resource, err := decodeXML[r5.ContainedResource](r)
		return resource.Resource, err
	default:
		// This should never happen as long as we control all implementations of the Release interface.
		// This is achieved by sealing the interface. See the interface definition for more information.
		panic("unsupported release")
	}
}

func decodeXML[T any](r *http.Request) (T, capabilities.FHIRError) {
	var v T
	if err := xml.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, decodingError{
			encoding: "xml",
		}
	}
	return v, nil
}
