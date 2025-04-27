package assert

import (
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assert/internal/diff"
	"testing"
)

func FHIRPathEqual(t *testing.T, expected, actual fhirpath.Collection) {
	// use equivalence to have empty results { } ~ { } result in true
	if !expected.Equivalent(actual) {
		t.Error(string(diff.Diff("expected", []byte(expected.String()), "actual", []byte(actual.String()))))
	}
}
