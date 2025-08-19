package search

import (
	"encoding/json"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"github.com/cockroachdb/apd/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestParseAndToString(t *testing.T) {
	tests := []struct {
		name         string
		capabilities r4.SearchCapabilities
		options      Options
		want         string
	}{
		{
			name: "number",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{ParameterKey{Name: "number"}: MatchAll{{Number{Value: apd.New(100, -3)}}}},
			},
			want: "number=0.100",
		},
		{
			name: "number with prefix",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{ParameterKey{Name: "number"}: MatchAll{{Number{Prefix: PrefixGreaterOrEqual, Value: apd.New(100, -3)}}}},
			},
			want: "number=ge0.100",
		},
		{
			name: "number with modifer",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type:     r4.Code{Value: ptr.To(TypeNumber)},
						Modifier: []r4.Code{{Value: ptr.To(ModifierMissing)}},
					},
				},
			},
			options: Options{
				Parameters: internalParams{ParameterKey{Name: "number", Modifier: ModifierMissing}: MatchAll{{Number{Value: apd.New(100, -3)}}}},
			},
			want: "number:missing=0.100",
		},
		{
			name: "number with count",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{ParameterKey{Name: "number"}: MatchAll{{Number{Value: apd.New(100, -3)}}}},
				Count:      100,
			},
			want: "number=0.100&_count=100",
		},
		{
			name: "number with max count",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"number": {
						Type: r4.Code{Value: ptr.To(TypeNumber)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{ParameterKey{Name: "number"}: MatchAll{{Number{Value: apd.New(100, -3)}}}},
				Count:      1000,
			},
			want: "number=0.100&_count=500",
		},
		{
			name: "date",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"date": {
						Type: r4.Code{Value: ptr.To(TypeDate)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "date"}: MatchAll{{Date{
						Value:     time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
						Precision: PrecisionDay,
					}}},
				},
			},
			want: "date=2024-12-25",
		},
		{
			name: "string",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"string": {
						Type: r4.Code{Value: ptr.To(TypeString)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "string"}: MatchAll{{String("example")}},
				},
			},
			want: "string=example",
		},
		{
			name: "token",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"token": {
						Type: r4.Code{Value: ptr.To(TypeToken)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "token"}: MatchAll{
						{Token{Code: "value"}},
					},
				},
			},
			want: "token=value",
		},
		{
			name: "token parameter with system",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"token": {
						Type: r4.Code{Value: ptr.To(TypeToken)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "token"}: MatchAll{{Token{System: &url.URL{Scheme: "scheme", Host: "system"}, Code: "value"}}},
				},
			},
			want: "token=scheme://system|value",
		},
		{
			name: "local reference",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "ref"}: MatchAll{{Reference{Type: "Patient", Id: "123"}}},
				},
			},
			want: "ref=Patient/123",
		},
		{
			name: "local reference with version",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "ref"}: MatchAll{{Reference{Type: "Patient", Id: "123", Version: "456"}}},
				},
			},
			want: "ref=Patient/123/_history/456",
		},
		{
			name: "url reference",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "ref"}: MatchAll{{Reference{URL: &url.URL{Scheme: "scheme", Host: "host"}}}},
				},
			},
			want: "ref=scheme://host",
		},
		{
			name: "url reference with version",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type: r4.Code{Value: ptr.To(TypeReference)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "ref"}: MatchAll{{Reference{URL: &url.URL{Scheme: "scheme", Host: "host"}, Version: "456"}}},
				},
			},
			want: "ref=scheme://host|456",
		},
		{
			name: "reference identifier modifier (treated as token)",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"ref": {
						Type:     r4.Code{Value: ptr.To(TypeReference)},
						Modifier: []r4.Code{{Value: ptr.To(ModifierIdentifier)}},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "ref", Modifier: ModifierIdentifier}: MatchAll{{Token{System: &url.URL{Scheme: "scheme", Host: "system"}, Code: "value"}}},
				},
			},
			want: "ref:identifier=scheme://system|value",
		},
		{
			name: "composite",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"composite": {
						Type: r4.Code{Value: ptr.To(TypeComposite)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "composite"}: MatchAll{{Composite{"a", "b"}}},
				},
			},
			want: "composite=a$b",
		},
		{
			name: "quantity",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"quantity": {
						Type: r4.Code{Value: ptr.To(TypeQuantity)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "quantity"}: MatchAll{{Quantity{Prefix: PrefixGreaterOrEqual, Value: apd.New(100, -3), System: &url.URL{Scheme: "scheme", Host: "host"}, Code: "code"}}},
				},
			},
			want: "quantity=ge0.100|scheme://host|code",
		},
		{
			name: "uri",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"uri": {
						Type: r4.Code{Value: ptr.To(TypeUri)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "uri"}: MatchAll{{Uri{&url.URL{Scheme: "urn", Opaque: "oid:1.2.3.4.5"}}}},
				},
			},
			want: "uri=urn:oid:1.2.3.4.5",
		},
		{
			name: "special",
			capabilities: r4.SearchCapabilities{
				Parameters: map[string]r4.SearchParameter{
					"special": {
						Type: r4.Code{Value: ptr.To(TypeSpecial)},
					},
				},
			},
			options: Options{
				Parameters: internalParams{
					ParameterKey{Name: "special"}: MatchAll{{Special("abc")}},
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

			// Convert old test format to new format
			// Create a mock CapabilityStatement with SearchParam definitions
			capabilityStatement := basic.CapabilityStatement{
				Rest: []basic.CapabilityStatementRest{
					{
						Resource: []basic.CapabilityStatementRestResource{
							{
								Type:        basic.Code{Value: ptr.To("TestResource")},
								SearchParam: []basic.CapabilityStatementRestResourceSearchParam{},
							},
						},
					},
				},
			}

			// Add SearchParam definitions and create resolver map
			searchParamMap := make(map[string]r4.SearchParameter)
			for name, param := range tt.capabilities.Parameters {
				canonical := "http://example.com/SearchParameter/" + name
				capabilityStatement.Rest[0].Resource[0].SearchParam = append(
					capabilityStatement.Rest[0].Resource[0].SearchParam,
					basic.CapabilityStatementRestResourceSearchParam{
						Name:       basic.String{Value: ptr.To(name)},
						Definition: &basic.Canonical{Value: ptr.To(canonical)},
					},
				)
				searchParamMap[canonical] = param
			}

			// Create resolver function
			resolveSearchParameter := func(canonical string) (model.Element, error) {
				if param, ok := searchParamMap[canonical]; ok {
					return param, nil
				}
				return nil, fmt.Errorf("SearchParameter not found: %s", canonical)
			}

			// test parse
			parsedOpts, err := ParseOptions(capabilityStatement, "TestResource", resolveSearchParameter, wantValues, time.UTC, 500, tt.options.Count, false)
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

func TestParseOptionsStrict(t *testing.T) {
	capabilityStatement := basic.CapabilityStatement{
		Rest: []basic.CapabilityStatementRest{
			{
				Resource: []basic.CapabilityStatementRestResource{
					{
						Type: basic.Code{Value: ptr.To("TestResource")},
						SearchParam: []basic.CapabilityStatementRestResourceSearchParam{
							{
								Name:       basic.String{Value: ptr.To("supported-param")},
								Definition: &basic.Canonical{Value: ptr.To("http://example.com/SearchParameter/supported")},
								Type:       basic.Code{Value: ptr.To("string")},
							},
						},
					},
				},
			},
		},
	}

	resolveSearchParameter := func(canonical string) (model.Element, error) {
		return &r4.SearchParameter{
			Type: r4.Code{Value: ptr.To("string")},
		}, nil
	}

	testCases := []struct {
		name          string
		strict        bool
		queryParams   url.Values
		expectError   bool
		errorContains string
	}{
		{
			name:   "strict_mode_with_supported_param",
			strict: true,
			queryParams: url.Values{
				"supported-param": []string{"value"},
			},
			expectError: false,
		},
		{
			name:   "strict_mode_with_unsupported_param",
			strict: true,
			queryParams: url.Values{
				"unsupported-param": []string{"value"},
			},
			expectError:   true,
			errorContains: "unsupported search parameter: unsupported-param",
		},
		{
			name:   "non_strict_mode_with_unsupported_param",
			strict: false,
			queryParams: url.Values{
				"unsupported-param": []string{"value"},
			},
			expectError: false,
		},
		{
			name:   "strict_mode_with_result_modifying_params",
			strict: true,
			queryParams: url.Values{
				"_count":   []string{"10"},
				"_include": []string{"Patient:organization"},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseOptions(capabilityStatement, "TestResource", resolveSearchParameter, tc.queryParams, time.UTC, 500, 50, tc.strict)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("Expected error to contain '%s' but got: %s", tc.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestParametersMarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		parameter internalParams
		expected  string
	}{
		{
			name:      "No Modifier",
			parameter: internalParams{ParameterKey{Name: "exampleName"}: MatchAll{{Number{Value: apd.New(100, -3)}}}},
			expected:  `{"exampleName":[[{"Prefix":"","Value":"0.100"}]]}`},
		{
			name:      "Modifier",
			parameter: internalParams{ParameterKey{Name: "exampleName", Modifier: ModifierExact}: MatchAll{{Number{Value: apd.New(100, -3)}}}},
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
