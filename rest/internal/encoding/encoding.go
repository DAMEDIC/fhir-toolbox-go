package encoding

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"io"
)

type Format string

const (
	FormatJSON Format = "application/fhir+json"
	FormatXML  Format = "application/fhir+xml"
)

func disabledErr[R model.Release]() error {
	r := model.ReleaseName[R]()
	return fmt.Errorf("release %s disabled by build tag; remove all build tags or add %s", r, r)
}

func DecodeResource[R model.Release](r io.Reader, format Format) (model.Resource, error) {
	var release R
	switch any(release).(type) {
	case model.R4:
		return decodeR4Resource(r, format)
	case model.R4B:
		return decodeR4BResource(r, format)
	case model.R5:
		return decodeR5Resource(r, format)
	default:
		// This should never happen as long as we control all implementations of the Release interface.
		// This is achieved by sealing the interface. See the interface definition for more information.
		panic("unsupported release")
	}
}

var decodeR4Resource = func(r io.Reader, format Format) (model.Resource, error) {
	return nil, disabledErr[model.R4]()
}
var decodeR4BResource = func(r io.Reader, format Format) (model.Resource, error) {
	return nil, disabledErr[model.R4B]()
}
var decodeR5Resource = func(r io.Reader, format Format) (model.Resource, error) {
	return nil, disabledErr[model.R5]()
}

func decodingError(encoding string) basic.OperationOutcome {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{
				Severity:    basic.Code{Value: ptr.To("fatal")},
				Code:        basic.Code{Value: ptr.To("processing")},
				Diagnostics: &basic.String{Value: ptr.To(fmt.Sprintf("error parsing %s body", encoding))},
			},
		},
	}
}

func Decode[T any](r io.Reader, format Format) (T, error) {
	switch format {
	case FormatJSON:
		return decodeJSON[T](r)
	case FormatXML:
		return decodeXML[T](r)
	default:
		return *new(T), fmt.Errorf("unsupported format: %s", format)
	}
}

func Encode[T any](w io.Writer, v T, format Format) error {
	switch format {
	case FormatJSON:
		return encodeJSON(w, v)
	case FormatXML:
		return encodeXML(w, v)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func encodeJSON[T any](w io.Writer, v T) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	return encoder.Encode(v)
}

func decodeJSON[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, decodingError("json")
	}
	return v, nil
}

func encodeXML[T any](w io.Writer, v T) error {
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return fmt.Errorf("encode xml: %w", err)
	}
	return xml.NewEncoder(w).Encode(v)
}

func decodeXML[T any](r io.Reader) (T, error) {
	var v T
	if err := xml.NewDecoder(r).Decode(&v); err != nil {
		return v, decodingError("xml")
	}
	return v, nil
}
