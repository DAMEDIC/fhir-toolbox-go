package search

import (
	"encoding/json"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"github.com/cockroachdb/apd/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"net/url"
	"testing"
	"time"
)

func TestParseAndToString(t *testing.T) {
	tests := []struct {
		name         string
		capabilities Capabilities[r4.SearchParameter]
		options      Options
		want         string
	}{
		{
			name: "number",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{ParameterKey{Name: "number"}: AllOf{{Number{Value: apd.New(100, -3)}}}},
			},
			want: "number=0.100",
		},
		{
			name: "number with prefix",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{ParameterKey{Name: "number"}: AllOf{{Number{Prefix: PrefixGreaterOrEqual, Value: apd.New(100, -3)}}}},
			},
			want: "number=ge0.100",
		},
		{
			name: "number with modifer",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type:     r4.Code{Value: ptr.To(TypeNumber)},
						Modifier: []r4.Code{{Value: ptr.To(ModifierMissing)}},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{ParameterKey{Name: "number", Modifier: ModifierMissing}: AllOf{{Number{Value: apd.New(100, -3)}}}},
			},
			want: "number:missing=0.100",
		},
		{
			name: "number with count",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{ParameterKey{Name: "number"}: AllOf{{Number{Value: apd.New(100, -3)}}}},
				Count:      100,
			},
			want: "number=0.100&_count=100",
		},
		{
			name: "number with max count",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{ParameterKey{Name: "number"}: AllOf{{Number{Value: apd.New(100, -3)}}}},
				Count:      1000,
			},
			want: "number=0.100&_count=500",
		},
		{
			name: "date",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"date": {
						Type: r4.Code{Value: ptr.To(TypeDate)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "date"}: AllOf{{Date{
						Value:     time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
						Precision: PrecisionDay,
					}}},
				},
			},
			want: "date=2024-12-25",
		},
		{
			name: "string",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"string": {
						Type: r4.Code{Value: ptr.To(TypeString)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "string"}: AllOf{{String("example")}},
				},
			},
			want: "string=example",
		},
		{
			name: "token",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"token": {
						Type: r4.Code{Value: ptr.To(TypeToken)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "token"}: AllOf{
						{Token{Code: "value"}},
					},
				},
			},
			want: "token=value",
		},
		{
			name: "token parameter with system",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"token": {
						Type: r4.Code{Value: ptr.To(TypeToken)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "token"}: AllOf{{Token{System: &url.URL{Scheme: "scheme", Host: "system"}, Code: "value"}}},
				},
			},
			want: "token=scheme://system|value",
		},
		{
			name: "local reference",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "ref"}: AllOf{{Reference{Type: "Patient", Id: "123"}}},
				},
			},
			want: "ref=Patient/123",
		},
		{
			name: "local reference with version",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "ref"}: AllOf{{Reference{Type: "Patient", Id: "123", Version: "456"}}},
				},
			},
			want: "ref=Patient/123/_history/456",
		},
		{
			name: "url reference",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "ref"}: AllOf{{Reference{Url: &url.URL{Scheme: "scheme", Host: "host"}}}},
				},
			},
			want: "ref=scheme://host",
		},
		{
			name: "url reference with version",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "ref"}: AllOf{{Reference{Url: &url.URL{Scheme: "scheme", Host: "host"}, Version: "456"}}},
				},
			},
			want: "ref=scheme://host|456",
		},
		{
			name: "reference identifier modifier (treated as token)",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type:     r4.Code{Value: ptr.To(TypeReference)},
						Modifier: []r4.Code{{Value: ptr.To(ModifierIdentifier)}},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "ref", Modifier: ModifierIdentifier}: AllOf{{Token{System: &url.URL{Scheme: "scheme", Host: "system"}, Code: "value"}}},
				},
			},
			want: "ref:identifier=scheme://system|value",
		},
		{
			name: "composite",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"composite": {
						Type: r4.Code{Value: ptr.To(TypeComposite)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "composite"}: AllOf{{Composite{"a", "b"}}},
				},
			},
			want: "composite=a$b",
		},
		{
			name: "quantity",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"quantity": {
						Type: r4.Code{Value: ptr.To(TypeQuantity)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "quantity"}: AllOf{{Quantity{Prefix: PrefixGreaterOrEqual, Value: apd.New(100, -3), System: &url.URL{Scheme: "scheme", Host: "host"}, Code: "code"}}},
				},
			},
			want: "quantity=ge0.100|scheme://host|code",
		},
		{
			name: "uri",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"uri": {
						Type: r4.Code{Value: ptr.To(TypeUri)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "uri"}: AllOf{{Uri{&url.URL{Scheme: "urn", Opaque: "oid:1.2.3.4.5"}}}},
				},
			},
			want: "uri=urn:oid:1.2.3.4.5",
		},
		{
			name: "special",
			capabilities: Capabilities[r4.SearchParameter]{
				Parameters: map[string]r4.SearchParameter{
					"special": {
						Type: r4.Code{Value: ptr.To(TypeSpecial)},
					},
				},
			},
			options: Options{
				Parameters: map[ParameterKey]AllOf{
					ParameterKey{Name: "special"}: AllOf{{Special("abc")}},
				},
			},
			want: "special=abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantValues, err := url.ParseQuery(tt.want)
			if err != nil {
				t.Fatalf("Failed to parse query: %v", err)
			}

			// test parse
			parsedOpts, err := ParseOptions(tt.capabilities, wantValues, time.UTC, 500, tt.options.Count)
			if err != nil {
				t.Fatalf("Failed to parse options: %v", err)
			}

			tt.options.Count = min(tt.options.Count, 500)

			if !cmp.Equal(parsedOpts, tt.options, cmpopts.EquateComparable(apd.Decimal{})) {
				t.Errorf("ParseOptions() = %v, want %v, diff: %s", parsedOpts, tt.options, cmp.Diff(parsedOpts, tt.options, cmpopts.EquateComparable(apd.Decimal{})))
			}

			// test to string
			gotValues, err := url.ParseQuery(tt.options.QueryString())
			if err != nil {
				t.Fatalf("Failed to parse query string: %v", err)
			}

			if !cmp.Equal(wantValues, gotValues, cmpopts.EquateComparable(apd.Decimal{})) {
				t.Errorf("QueryString() = %v, want %v, diff: %s", gotValues, wantValues, cmp.Diff(wantValues, gotValues, cmpopts.EquateComparable(apd.Decimal{})))
			}
		})
	}
}

func TestParametersMarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		parameter map[ParameterKey]AllOf
		expected  string
	}{
		{
			name:      "No Modifier",
			parameter: map[ParameterKey]AllOf{ParameterKey{Name: "exampleName"}: AllOf{{Number{Value: apd.New(100, -3)}}}},
			expected:  `{"exampleName":[[{"Prefix":"","Value":"0.100"}]]}`},
		{
			name:      "Modifier",
			parameter: map[ParameterKey]AllOf{ParameterKey{Name: "exampleName", Modifier: ModifierExact}: AllOf{{Number{Value: apd.New(100, -3)}}}},
			expected:  `{"exampleName:exact":[[{"Prefix":"","Value":"0.100"}]]}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.parameter)
			if err != nil {
				t.Fatalf("MarshalJSON should not return an error: %v", err)
			}

			// Compare JSON by unmarshaling both strings to ensure they're equivalent
			var expected, actual interface{}
			if err := json.Unmarshal([]byte(tt.expected), &expected); err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}
			if err := json.Unmarshal(data, &actual); err != nil {
				t.Fatalf("Failed to unmarshal actual JSON: %v", err)
			}

			if !cmp.Equal(expected, actual, cmpopts.EquateComparable(apd.Decimal{})) {
				t.Errorf("JSON output does not match expected.\nExpected: %s\nActual: %s\nDiff: %s", tt.expected, string(data), cmp.Diff(expected, actual, cmpopts.EquateComparable(apd.Decimal{})))
			}
		})
	}
}
