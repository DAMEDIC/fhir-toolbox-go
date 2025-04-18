package search

import (
	"encoding/json"
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
		capabilities Capabilities
		options      Options
		want         string
	}{
		{
			name:         "number",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"number": {Type: TypeNumber}}},
			options: Options{
				Parameters: Parameters{ParameterKey{Name: "number"}: All{{Number{Value: apd.New(100, -3)}}}},
			},
			want: "number=0.100",
		},
		{
			name:         "number with prefix",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"number": {Type: TypeNumber}}},
			options: Options{
				Parameters: Parameters{ParameterKey{Name: "number"}: All{{Number{Prefix: PrefixGreaterOrEqual, Value: apd.New(100, -3)}}}},
			},
			want: "number=ge0.100",
		},
		{
			name:         "number with modifer",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"number": {Type: TypeNumber, Modifiers: []Modifier{ModifierMissing}}}},
			options: Options{
				Parameters: Parameters{ParameterKey{Name: "number", Modifier: ModifierMissing}: All{{Number{Value: apd.New(100, -3)}}}},
			},
			want: "number:missing=0.100",
		},
		{
			name:         "number with count",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"number": {Type: TypeNumber}}},
			options: Options{
				Parameters: Parameters{ParameterKey{Name: "number"}: All{{Number{Value: apd.New(100, -3)}}}},
				Count:      100,
			},
			want: "number=0.100&_count=100",
		},
		{
			name:         "number with max count",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"number": {Type: TypeNumber}}},
			options: Options{
				Parameters: Parameters{ParameterKey{Name: "number"}: All{{Number{Value: apd.New(100, -3)}}}},
				Count:      1000,
			},
			want: "number=0.100&_count=500",
		},
		{
			name:         "date",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"date": {Type: TypeDate}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "date"}: All{{Date{
						Value:     time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
						Precision: PrecisionDay,
					}}},
				},
			},
			want: "date=2024-12-25",
		},
		{
			name:         "string",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"string": {Type: TypeString}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "string"}: All{{String("example")}},
				},
			},
			want: "string=example",
		},
		{
			name:         "token",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"token": {Type: TypeToken}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "token"}: All{
						{Token{Code: "value"}},
					},
				},
			},
			want: "token=value",
		},
		{
			name:         "token parameter with system",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"token": {Type: TypeToken}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "token"}: All{{Token{System: &url.URL{Scheme: "scheme", Host: "system"}, Code: "value"}}},
				},
			},
			want: "token=scheme://system|value",
		},
		{
			name:         "local reference",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"ref": {Type: TypeReference}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "ref"}: All{{Reference{Type: "Patient", Id: "123"}}},
				},
			},
			want: "ref=Patient/123",
		},
		{
			name:         "local reference with version",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"ref": {Type: TypeReference}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "ref"}: All{{Reference{Type: "Patient", Id: "123", Version: "456"}}},
				},
			},
			want: "ref=Patient/123/_history/456",
		},
		{
			name:         "url reference",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"ref": {Type: TypeReference}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "ref"}: All{{Reference{Url: &url.URL{Scheme: "scheme", Host: "host"}}}},
				},
			},
			want: "ref=scheme://host",
		},
		{
			name:         "url reference with version",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"ref": {Type: TypeReference}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "ref"}: All{{Reference{Url: &url.URL{Scheme: "scheme", Host: "host"}, Version: "456"}}},
				},
			},
			want: "ref=scheme://host|456",
		},
		{
			name:         "reference identifier modifier (treated as token)",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"ref": {Type: TypeReference, Modifiers: []Modifier{ModifierIdentifier}}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "ref", Modifier: ModifierIdentifier}: All{{Token{System: &url.URL{Scheme: "scheme", Host: "system"}, Code: "value"}}},
				},
			},
			want: "ref:identifier=scheme://system|value",
		},
		{
			name:         "composite",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"composite": {Type: TypeComposite}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "composite"}: All{{Composite{"a", "b"}}},
				},
			},
			want: "composite=a$b",
		},
		{
			name:         "quantity",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"quantity": {Type: TypeQuantity}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "quantity"}: All{{Quantity{Prefix: PrefixGreaterOrEqual, Value: apd.New(100, -3), System: &url.URL{Scheme: "scheme", Host: "host"}, Code: "code"}}},
				},
			},
			want: "quantity=ge0.100|scheme://host|code",
		},
		{
			name:         "uri",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"uri": {Type: TypeUri}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "uri"}: All{{Uri{&url.URL{Scheme: "urn", Opaque: "oid:1.2.3.4.5"}}}},
				},
			},
			want: "uri=urn:oid:1.2.3.4.5",
		},
		{
			name:         "special",
			capabilities: Capabilities{Parameters: ParameterDescriptions{"special": {Type: TypeSpecial}}},
			options: Options{
				Parameters: Parameters{
					ParameterKey{Name: "special"}: All{{Special("abc")}},
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
		parameter Parameters
		expected  string
	}{
		{
			name:      "No Modifier",
			parameter: Parameters{ParameterKey{Name: "exampleName"}: All{{Number{Value: apd.New(100, -3)}}}},
			expected:  `{"exampleName":[[{"Prefix":"","Value":"0.100"}]]}`},
		{
			name:      "Modifier",
			parameter: Parameters{ParameterKey{Name: "exampleName", Modifier: ModifierExact}: All{{Number{Value: apd.New(100, -3)}}}},
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
