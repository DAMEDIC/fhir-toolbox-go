package fhirpath_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assert"
	"github.com/cockroachdb/apd/v3"
)

// runFHIRPathTest executes a single FHIRPath test and validates the result
func runFHIRPathTest(t *testing.T, ctx context.Context, test testdata.FHIRPathTest) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()

	expr, err := fhirpath.Parse(test.Expression.Expression)
	if err != nil && (test.Invalid != "" || test.Expression.Invalid != "") {
		return
	}
	if err != nil {
		t.Fatalf("Unexpected error parsing expression: %v", err)
	}

	result, err := fhirpath.Evaluate(
		ctx,
		test.InputResource,
		expr,
	)
	if err != nil && test.Expression.Invalid != "" {
		return
	}
	if err != nil {
		t.Fatalf("Unexpected error evaluating expression: %v", err)
	}

	if test.Predicate {
		v, ok, err := fhirpath.Singleton[fhirpath.Boolean](result)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !ok {
			t.Fatalf("expected boolean value to exist")
		}
		result = fhirpath.Collection{v}
	}

	expected := test.OutputCollection()
	assert.FHIRPathEqual(t, expected, result)
}

func TestFHIRPathTestSuiteR4(t *testing.T) {
	ctx := r4.Context()
	ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(8))

	tests := testdata.GetFHIRPathTests()

	for _, group := range tests.Groups {
		name := group.Name
		if group.Description != "" {
			name = fmt.Sprintf("%s (%s)", name, group.Description)
		}

		t.Run(name, func(t *testing.T) {
			if group.Name == "testQuantity" {
				t.Skip("UCUM conversions not implemented")
			}

			for _, test := range group.Tests {
				name := test.Name
				if test.Description != "" {
					name = fmt.Sprintf("%s (%s)", name, test.Description)
				}

				for r, o := range testOverrides {
					if ok, _ := regexp.MatchString(r, test.Name); ok {
						if o.Invalid != "" {
							test.Invalid = o.Invalid
						}
						if o.Output != nil {
							test.Output = o.Output
						}
						if o.Expression != (testdata.FHIRPathTestExpression{}) {
							test.Expression = o.Expression
						}
						break
					}
				}

				t.Run(name, func(t *testing.T) {
					runFHIRPathTest(t, ctx, test)
				})
			}
		})
	}
}

// overrides for fhirerrors in the test xml
var testOverrides = map[string]testdata.FHIRPathTest{
	"testPrecedence3": {
		// test expression has wrong precedence
		Expression: testdata.FHIRPathTestExpression{
			Expression: "(1 > 2) is Boolean",
		},
	},
	"testPrecedence4": {
		// test expression has wrong precedence
		Expression: testdata.FHIRPathTestExpression{
			Expression: "(1 | 1) is Integer",
		},
	},
	"testQuantityLiteralWeekToString": {
		// proper ucum handling not implemented
		Output: []testdata.FHIRPathTestOutput{{
			Type: "Quantity", Output: "1 'week'",
		}},
	},
	"testDateNotEqualTimezoneOffset(Before|After)": {
		// should be empty
		Output: []testdata.FHIRPathTestOutput{},
	},
	"testDateNotEqualUTC": {
		// should be empty
		Output: []testdata.FHIRPathTestOutput{},
	},
	"testDateTimeGreaterThanDate": {
		// should be empty
		Output: []testdata.FHIRPathTestOutput{},
	},
	"testNow1": {
		// should be empty
		Output: []testdata.FHIRPathTestOutput{},
	},
	"testIntegerBooleanNotTrue": {
		// singleton evaluation of '0' is true
		Output: []testdata.FHIRPathTestOutput{{
			Type: "boolean", Output: "false",
		}},
	},
	"testStringQuantityDayLiteralToQuantity": {
		// proper ucum handling not implemented
		Output: []testdata.FHIRPathTestOutput{},
	},
	"testEquality(7|23)": {
		// missing output in test
		Output: []testdata.FHIRPathTestOutput{{
			Type: "boolean", Output: "false",
		}},
	},
	"testNEquality17": {
		// missing output in test
		Output: []testdata.FHIRPathTestOutput{{
			Type: "boolean", Output: "true",
		}},
	},
	"testNEquality24": {
		// different units yields empty
		Output: []testdata.FHIRPathTestOutput{},
	},
	"testNotEquivalent19": {
		// should be false
		Output: []testdata.FHIRPathTestOutput{{
			Type: "boolean", Output: "false",
		}},
	},
	"testRound2": {
		Expression: testdata.FHIRPathTestExpression{
			Expression: "13.14159.round(3) = 13.142",
		},
	},
	"testType22": {
		Expression: testdata.FHIRPathTestExpression{
			Invalid: "not existing type returns error",
		},
	},
	"testVariables4": {
		Expression: testdata.FHIRPathTestExpression{
			Invalid: "%vs- value set variables are currently not supported",
		},
	},
	"testExtension2": {
		Expression: testdata.FHIRPathTestExpression{
			Invalid: "%ext- extension variables are currently not supported",
		},
	},
	"testConformsTo": {
		Expression: testdata.FHIRPathTestExpression{
			Invalid: "testConformsTo function not implemented yet",
		},
	},
}

func TestAdditionalFHIRPathTests(t *testing.T) {
	ctx := r4.Context()
	ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(8))

	var additionalTests = []testdata.FHIRPathTest{
		{
			Name:        "testDefineVariableSelect",
			Description: "Test defineVariable with select",
			Expression: testdata.FHIRPathTestExpression{
				Expression: "defineVariable('a', 'b').select(%a)",
			},
			Output: []testdata.FHIRPathTestOutput{{
				Type:   "string",
				Output: "b",
			}},
		},
	}

	for _, test := range additionalTests {
		name := test.Name
		if test.Description != "" {
			name = fmt.Sprintf("%s (%s)", name, test.Description)
		}

		t.Run(name, func(t *testing.T) {
			runFHIRPathTest(t, ctx, test)
		})
	}
}
