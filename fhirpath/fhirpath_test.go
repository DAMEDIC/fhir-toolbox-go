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

				t.Run(name, func(t *testing.T) {
					defer func() {
						if err := recover(); err != nil {
							t.Fatal(err)
						}
					}()

					expr, err := fhirpath.Parse(test.Expression.Expression)
					if err != nil && test.Expression.Invalid != "" {
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

					expected := test.OutputCollection()
					diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
						A:        difflib.SplitLines(expected.String()),
						B:        difflib.SplitLines(result.String()),
						FromFile: "Expected",
						ToFile:   "Actual",
						Context:  1,
					})
					// use equivalence to have empty results { } ~ { } result in true
					assert.True(t, *expected.Equivalent(result), diff)
				})
			}
		})
	}
}
