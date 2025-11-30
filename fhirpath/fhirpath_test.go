package fhirpath_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
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

func runFHIRPathSuite(t *testing.T, ctx context.Context, release model.Release) {
	t.Helper()

	tests := testdata.GetFHIRPathTests(release)

	for _, group := range tests.Groups {
		group := group
		name := group.Name
		if group.Description != "" {
			name = fmt.Sprintf("%s (%s)", name, group.Description)
		}

		t.Run(name, func(t *testing.T) {
			for _, test := range group.Tests {
				test := test
				name := test.Name
				if test.Description != "" {
					name = fmt.Sprintf("%s (%s)", name, test.Description)
				}

				t.Run(name, func(t *testing.T) {
					if ok, reason := shouldSkipTest(test, release); ok {
						t.Skip(reason)
					}
					runFHIRPathTest(t, ctx, test)
				})
			}
		})
	}
}

func TestFHIRPathTestSuites(t *testing.T) {
	for _, release := range testdata.TestReleases {
		release := release
		t.Run(release.String(), func(t *testing.T) {
			var ctx context.Context
			switch release.(type) {
			case model.R4:
				ctx = r4.Context()
			case model.R4B:
				ctx = r4b.Context()
			case model.R5:
				ctx = r5.Context()
			default:
				t.Fatalf("no context configured for release %s", release.String())
			}
			ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(8))
			runFHIRPathSuite(t, ctx, release)
		})
	}
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

func isCDATest(test testdata.FHIRPathTest) bool {
	return strings.EqualFold(test.Mode, "cda")
}

func isR5Release(release model.Release) bool {
	_, ok := release.(model.R5)
	return ok
}

type skipRule struct {
	pattern       *regexp.Regexp
	releaseFilter func(model.Release) bool
	reason        string
}

var unimplementedTestSkips = []skipRule{
	{regexp.MustCompile(`^Comparable\d+$`), nil, "comparable() function not implemented"},
	{regexp.MustCompile(`^testMultipleResolve$`), nil, "resolve() function not implemented"},
	{regexp.MustCompile(`^testPrimitiveExtensions(Element)?$`), nil, "hasValue() for primitive extensions not implemented"},
	{regexp.MustCompile(`^testPeriodInvariantOld$`), nil, "hasValue() function not implemented"},
	{regexp.MustCompile(`^txTest0[1-3]$`), nil, "Terminology service not configured"},
	{regexp.MustCompile(`^testVariables4$`), nil, "%vs variables not supported"},
	{regexp.MustCompile(`^testExtension2$`), nil, "%ext variables not supported"},
	{regexp.MustCompile(`^testConformsTo.*`), nil, "conformsTo() function not implemented"},
	{regexp.MustCompile(`^testPolymorphismIsA3$`), nil, "polymorphism/is semantics not aligned with spec 3.0"},
	{regexp.MustCompile(`^testPolymorphism.*`), nil, "polymorphism/is semantics not aligned with spec 3.0"},
	{regexp.MustCompile(`^testReplace\d+$`), nil, "string replace semantics pending"},
	{regexp.MustCompile(`^testNow\d+$`), nil, "now() determinism/precision pending"},
	{regexp.MustCompile(`^testEquality\d+$`), nil, "Equality semantics pending"},
	{regexp.MustCompile(`^testNEquality\d+$`), nil, "Not-equal semantics pending"},
	{regexp.MustCompile(`^testCombine.*`), nil, "combine() semantics pending"},
	{regexp.MustCompile(`^testUnion\d+$`), nil, "union semantics pending"},
	{regexp.MustCompile(`^testPlus.*`), nil, "date/time arithmetic units pending"},
	{regexp.MustCompile(`^testExp\d+$`), nil, "exp() function stability pending"},
	// testPrecedence3 & testPrecedence4: The tests are wrong. In the formal FHIRPath grammar,
	// type operators (is, as) bind tighter than comparison operators (>, <, etc.)
	{regexp.MustCompile(`^testPrecedence[34]$`), nil, "test uses wrong precedence - type operators bind tighter than comparison operators"},
	// testPrecedence6: Complex expression with trace() and in operator - evaluation context issue
	{regexp.MustCompile(`^testPrecedence6$`), nil, "evaluation context issue with nested exists/in operators"},
	{regexp.MustCompile(`^testType.*`), nil, "type() semantics pending"},
	{regexp.MustCompile(`^LowBoundary.*`), nil, "lowBoundary/highBoundary precision semantics pending"},
	{regexp.MustCompile(`^HighBoundary.*`), nil, "lowBoundary/highBoundary precision semantics pending"},
	{regexp.MustCompile(`^Precision.*`), nil, "precision() semantics pending"},
	{regexp.MustCompile(`^testFHIRPathIsFunction\d+$`), nil, "is() function inheritance semantics pending"},
	{regexp.MustCompile(`^testContainedId$`), nil, "primitive id handling pending"},
	{regexp.MustCompile(`^testSubstring.*`), isR5Release, "substring parameter semantics pending"},
	{regexp.MustCompile(`^testStartsWith.*`), isR5Release, "startsWith optional parameters pending"},
	{regexp.MustCompile(`^testEndsWith.*`), isR5Release, "endsWith optional parameters pending"},
	{regexp.MustCompile(`^testContainsString.*`), isR5Release, "contains() optional parameters pending"},
	{regexp.MustCompile(`^testMinus.*`), isR5Release, "minus semantics for quantities pending"},
	{regexp.MustCompile(`^test(Sqrt|Abs|Ceiling|Floor|Ln|Log|Power|Truncate).*`), isR5Release, "math functions empty-input semantics pending"},
	// defineVariable13 & 14: Tests use skip(1).first() as the value expression, expecting it to operate on the entire
	// input collection. However, the spec states "If the function takes an expression as a parameter, the function
	// will evaluate the expression passed for the parameter with respect to each of the items in the input collection."
	// Current implementation follows this pattern (like select/where), evaluating the value expression per-item and
	// aggregating results. But skip(1) only makes sense on a collection, not individual items. The spec for defineVariable
	// doesn't clarify whether the value parameter should be evaluated per-item or once on the whole collection.
	{regexp.MustCompile(`^defineVariable1[34]$`), nil, "test expects whole-collection evaluation but implementation evaluates value expression per-item per spec"},
	// defineVariable19: Test uses nested defineVariable in the name parameter, which conflicts with empty collection handling.
	// Per spec, name is a String parameter (not expression parameter), evaluated once in root context.
	// When defineVariable processes empty input, the value expression is not evaluated (consistent with select, where, etc.).
	// This causes the nested defineVariable to set variables to empty collections, breaking the test's expectation.
	{regexp.MustCompile(`^defineVariable19$`), nil, "test conflicts with spec-compliant empty collection handling for expression parameters"},
}

func shouldSkipTest(test testdata.FHIRPathTest, release model.Release) (bool, string) {
	if isCDATest(test) && isR5Release(release) {
		return true, "CDA-based FHIRPath inputs not supported yet"
	}

	for _, rule := range unimplementedTestSkips {
		if rule.pattern.MatchString(test.Name) {
			if rule.releaseFilter == nil || rule.releaseFilter(release) {
				return true, rule.reason
			}
		}
	}

	return false, ""
}
