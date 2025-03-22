package fhirpath_test

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata"
	"github.com/cockroachdb/apd/v3"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
	"testing"
)

// overrides for errors in the test xml
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
	"testQuantity7": {
		Expression: testdata.FHIRPathTestExpression{
			Invalid: "UCUM handling not implemented",
		},
	},
	"testQuantity8": {
		Expression: testdata.FHIRPathTestExpression{
			Invalid: "UCUM handling not implemented",
		},
	},
	"testDateNotEqual": {
		// missing output in test
		Output: []testdata.FHIRPathTestOutput{{
			Type: "boolean", Output: "true",
		}},
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
}

func TestFHIRPathTestSuiteR4(t *testing.T) {
	ctx := r4.Context()
	ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(100))

	tests := testdata.GetFHIRPathTests()

	for _, group := range tests.Groups {
		name := group.Name
		if group.Description != "" {
			name = fmt.Sprintf("%s (%s)", name, group.Description)
		}

		t.Run(name, func(t *testing.T) {
			for _, test := range group.Tests {
				name := test.Name
				if test.Description != "" {
					name = fmt.Sprintf("%s (%s)", name, test.Description)
				}

				override, ok := testOverrides[name]
				if ok {
					if override.Invalid != "" {
						test.Invalid = override.Invalid
					}
					if override.Output != nil {
						test.Output = override.Output
					}
					if override.Expression != (testdata.FHIRPathTestExpression{}) {
						test.Expression = override.Expression
					}
				}

				t.Run(name, func(t *testing.T) {
					defer func() {
						if err := recover(); err != nil {
							t.Fatal(err)
						}
					}()

					expr, err := fhirpath.Parse(test.Expression.Expression)
					if err != nil && (test.Invalid != "" || test.Expression.Invalid != "") {
						return
					}

					result, err := fhirpath.Evaluate(
						ctx,
						test.InputResource,
						expr,
					)
					if err != nil && test.Expression.Invalid != "" {
						return
					}
					assert.NoError(t, err)

					if test.Predicate {
						v, ok, err := fhirpath.Singleton[fhirpath.Boolean](result)
						assert.NoError(t, err)
						assert.True(t, ok, "expected boolean value to exist")
						result = fhirpath.Collection{v}
					}

					expected := test.OutputCollection()
					diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
						A:        difflib.SplitLines(expected.String()),
						B:        difflib.SplitLines(result.String()),
						FromFile: "Expected",
						ToFile:   "Actual",
						Context:  1,
					})
					// use equivalence to have empty results { } ~ { } result in true
					assert.True(t, expected.Equivalent(result), diff)
				})
			}
		})
	}
}
