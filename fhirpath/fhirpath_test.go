package fhirpath_test

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFHIRPathTestSuiteR4(t *testing.T) {
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
					expr, err := fhirpath.Parse(test.Expression.Expression)
					if err != nil && test.Expression.Invalid != "" {
						return
					}

					result, err := fhirpath.Evaluate(
						r4.Context(),
						test.InputResource,
						expr,
					)
					assert.NoError(t, err)
					assert.Equal(t, test.OutputCollection(), result)
				})
			}
		})
	}
}
