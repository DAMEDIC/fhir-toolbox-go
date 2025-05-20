package fhirpath

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cockroachdb/apd/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func init() {
	// mock time.Now for testing
	now = func() time.Time { return time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC) }
}

// testElement is a helper type for testing FHIRPath functions
type testElement struct {
	value any
}

func (e testElement) Children(name ...string) Collection {
	if len(name) > 0 {
		return Collection{}
	}

	switch v := e.value.(type) {
	case []any:
		var result Collection
		for _, item := range v {
			if el, ok := item.(Element); ok {
				result = append(result, el)
			} else {
				// Convert primitives to FHIRPath types
				switch val := item.(type) {
				case int:
					result = append(result, Integer(val))
				case string:
					result = append(result, String(val))
				case bool:
					result = append(result, Boolean(val))
				case float64:
					result = append(result, Decimal{Value: apd.New(int64(val), 0)})
				default:
					// For other types, wrap in testElement
					result = append(result, testElement{value: val})
				}
			}
		}
		return result
	case []int:
		var result Collection
		for _, item := range v {
			result = append(result, Integer(item))
		}
		return result
	case []string:
		var result Collection
		for _, item := range v {
			result = append(result, String(item))
		}
		return result
	case []bool:
		var result Collection
		for _, item := range v {
			result = append(result, Boolean(item))
		}
		return result
	default:
		// For non-array values, return the value itself
		switch v := e.value.(type) {
		case int:
			return Collection{Integer(v)}
		case string:
			return Collection{String(v)}
		case bool:
			return Collection{Boolean(v)}
		case float64:
			return Collection{Decimal{Value: apd.New(int64(v), 0)}}
		default:
			return Collection{e}
		}
	}
}

func (e testElement) ToBoolean(explicit bool) (Boolean, bool, error) {
	switch v := e.value.(type) {
	case bool:
		return Boolean(v), true, nil
	case []bool:
		if len(v) == 0 {
			return false, false, nil
		}
		return Boolean(v[0]), true, nil
	default:
		return false, false, nil
	}
}

func (e testElement) ToString(explicit bool) (String, bool, error) {
	switch v := e.value.(type) {
	case string:
		return String(v), true, nil
	case []string:
		if len(v) == 0 {
			return "", false, nil
		}
		return String(v[0]), true, nil
	default:
		return "", false, nil
	}
}

func (e testElement) ToInteger(explicit bool) (Integer, bool, error) {
	switch v := e.value.(type) {
	case int:
		return Integer(v), true, nil
	case []int:
		if len(v) == 0 {
			return 0, false, nil
		}
		return Integer(v[0]), true, nil
	default:
		return 0, false, nil
	}
}

func (e testElement) ToDecimal(explicit bool) (Decimal, bool, error) {
	switch v := e.value.(type) {
	case Decimal:
		return v, true, nil
	case []Decimal:
		if len(v) == 0 {
			return Decimal{}, false, nil
		}
		return v[0], true, nil
	default:
		return Decimal{}, false, nil
	}
}

func (e testElement) ToDate(explicit bool) (Date, bool, error) {
	switch v := e.value.(type) {
	case Date:
		return v, true, nil
	case []Date:
		if len(v) == 0 {
			return Date{}, false, nil
		}
		return v[0], true, nil
	default:
		return Date{}, false, nil
	}
}

func (e testElement) ToTime(explicit bool) (Time, bool, error) {
	switch v := e.value.(type) {
	case Time:
		return v, true, nil
	case []Time:
		if len(v) == 0 {
			return Time{}, false, nil
		}
		return v[0], true, nil
	default:
		return Time{}, false, nil
	}
}

func (e testElement) ToDateTime(explicit bool) (DateTime, bool, error) {
	switch v := e.value.(type) {
	case DateTime:
		return v, true, nil
	case []DateTime:
		if len(v) == 0 {
			return DateTime{}, false, nil
		}
		return v[0], true, nil
	default:
		return DateTime{}, false, nil
	}
}

func (e testElement) ToQuantity(explicit bool) (Quantity, bool, error) {
	switch v := e.value.(type) {
	case Quantity:
		return v, true, nil
	case []Quantity:
		if len(v) == 0 {
			return Quantity{}, false, nil
		}
		return v[0], true, nil
	default:
		return Quantity{}, false, nil
	}
}

func (e testElement) Equal(other Element, _noReverseTypeConversion ...bool) (bool, bool) {
	o, ok := other.(testElement)
	if !ok {
		return false, false
	}
	return e.value == o.value, true
}

func (e testElement) Equivalent(other Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := e.Equal(other)
	return eq
}

func (e testElement) TypeInfo() TypeInfo {
	return ClassInfo{
		Namespace: "System",
		Name:      "testElement",
		BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
		Element: []ClassInfoElement{
			{Name: "value", Type: TypeSpecifier{Namespace: "System", Name: "Any"}},
		},
	}
}

func (e testElement) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (e testElement) String() string {
	return ""
}

// testFunction is a helper to test FHIRPath functions directly
func testFunction(t *testing.T, fn Function, target Collection, params []Expression, expected Collection, expectedOrdered bool, mockTime bool) {
	t.Helper()
	ctx := context.Background()
	// Set APD context with precision
	apdCtx := apd.BaseContext.WithPrecision(20)
	ctx = WithAPDContext(ctx, apdCtx)
	root := testElement{value: nil}

	// Mock evaluate function that can handle simple expressions
	mockEvaluate := func(ctx context.Context, target Element, expr Expression, fnScope ...FunctionScope) (Collection, bool, error) {
		if expr.tree == nil {
			return Collection{}, false, fmt.Errorf("unexpected expression <nil>")
		}

		// Handle simple expressions
		switch expr.String() {
		case "true":
			return Collection{Boolean(true)}, true, nil
		case "false":
			return Collection{Boolean(false)}, true, nil
		case "$this > 0", "$this>0":
			switch v := target.(type) {
			case testElement:
				if val, ok := v.value.(int); ok {
					return Collection{Boolean(val > 0)}, true, nil
				}
			case Integer:
				return Collection{Boolean(int(v) > 0)}, true, nil
			}
			return Collection{Boolean(false)}, true, nil
		case "$this > 2", "$this>2":
			switch v := target.(type) {
			case testElement:
				if val, ok := v.value.(int); ok {
					return Collection{Boolean(val > 2)}, true, nil
				}
			case Integer:
				return Collection{Boolean(int(v) > 2)}, true, nil
			}
			return Collection{Boolean(false)}, true, nil
		case "$this > 5", "$this>5":
			switch v := target.(type) {
			case testElement:
				if val, ok := v.value.(int); ok {
					return Collection{Boolean(val > 5)}, true, nil
				}
			case Integer:
				return Collection{Boolean(int(v) > 5)}, true, nil
			}
			return Collection{Boolean(false)}, true, nil
		case "$this * 2", "$this*2":
			switch v := target.(type) {
			case testElement:
				if val, ok := v.value.(int); ok {
					return Collection{Integer(val * 2)}, true, nil
				}
			case Integer:
				return Collection{Integer(int(v) * 2)}, true, nil
			}
			return Collection{}, true, nil
		case "1":
			return Collection{Integer(1)}, true, nil
		case "2":
			return Collection{Integer(2)}, true, nil
		case "3":
			return Collection{Integer(3)}, true, nil
		case "4":
			return Collection{Integer(4)}, true, nil
		case "children()":
			return Collection{}, true, nil
		case "'test'":
			return Collection{String("test")}, true, nil
		case "'hello'":
			return Collection{String("hello")}, true, nil
		case "'world'":
			return Collection{String("world")}, true, nil
		case "'hi'":
			return Collection{String("hi")}, true, nil
		case "'^hello.*'":
			return Collection{String("^hello.*")}, true, nil
		case "','":
			return Collection{String(",")}, true, nil
		case "{}":
			return Collection{}, true, nil
		case "{2, 3}":
			return Collection{Integer(2), Integer(3)}, true, nil
		case "{2, 3, 4}":
			return Collection{Integer(2), Integer(3), Integer(4)}, true, nil
		case "{3, 4, 5}":
			return Collection{Integer(3), Integer(4), Integer(5)}, true, nil
		case "2 | 3", "2|3":
			return Collection{Integer(2), Integer(3)}, true, nil
		case "2 | 3 | 4", "2|3|4":
			return Collection{Integer(2), Integer(3), Integer(4)}, true, nil
		case "3|4|5", "3 | 4 | 5":
			return Collection{Integer(3), Integer(4), Integer(5)}, true, nil
		default:
			// Try to parse as integer literal
			var intVal int
			if _, err := fmt.Sscanf(expr.String(), "%d", &intVal); err == nil {
				return Collection{Integer(intVal)}, true, nil
			}
			// Try to parse as boolean literal
			if expr.String() == "true" {
				return Collection{Boolean(true)}, true, nil
			}
			if expr.String() == "false" {
				return Collection{Boolean(false)}, true, nil
			}
			return Collection{}, false, fmt.Errorf("evaluate not implemented for expression: %s", expr.String())
		}
	}

	if fn == nil {
		t.Fatal("Function is nil")
	}

	result, ordered, err := fn(ctx, root, target, true, params, mockEvaluate)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if ordered != expectedOrdered {
		t.Errorf("Expected ordered=%v, got ordered=%v", expectedOrdered, ordered)
	}

	eq := result.Equivalent(expected)

	if !eq {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestExistenceFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       Function
		target   Collection
		params   []Expression
		expected Collection
	}{
		{
			name:     "empty()",
			fn:       defaultFunctions["empty"],
			target:   Collection{},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "empty() with non-empty collection",
			fn:       defaultFunctions["empty"],
			target:   Collection{Integer(1)},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "exists()",
			fn:       defaultFunctions["exists"],
			target:   Collection{Integer(1)},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "exists() with empty collection",
			fn:       defaultFunctions["exists"],
			target:   Collection{},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "all() with empty collection",
			fn:       defaultFunctions["allTrue"],
			target:   Collection{},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "all() with all true",
			fn:       defaultFunctions["allTrue"],
			target:   Collection{Boolean(true), Boolean(true)},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "all() with some false",
			fn:       defaultFunctions["allTrue"],
			target:   Collection{Boolean(true), Boolean(false)},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "any() with empty collection",
			fn:       defaultFunctions["anyTrue"],
			target:   Collection{},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "any() with some true",
			fn:       defaultFunctions["anyTrue"],
			target:   Collection{Boolean(false), Boolean(true)},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "any() with all false",
			fn:       defaultFunctions["anyTrue"],
			target:   Collection{Boolean(false), Boolean(false)},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "allFalse() with empty collection",
			fn:       defaultFunctions["allFalse"],
			target:   Collection{},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "allFalse() with all false",
			fn:       defaultFunctions["allFalse"],
			target:   Collection{Boolean(false), Boolean(false)},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "allFalse() with some true",
			fn:       defaultFunctions["allFalse"],
			target:   Collection{Boolean(false), Boolean(true)},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "anyFalse() with empty collection",
			fn:       defaultFunctions["anyFalse"],
			target:   Collection{},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
		{
			name:     "anyFalse() with some false",
			fn:       defaultFunctions["anyFalse"],
			target:   Collection{Boolean(true), Boolean(false)},
			params:   nil,
			expected: Collection{Boolean(true)},
		},
		{
			name:     "anyFalse() with all true",
			fn:       defaultFunctions["anyFalse"],
			target:   Collection{Boolean(true), Boolean(true)},
			params:   nil,
			expected: Collection{Boolean(false)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, true, false)
		})
	}
}

func TestFilteringAndProjectionFunctions(t *testing.T) {
	tests := []struct {
		name            string
		fn              Function
		target          Collection
		params          []Expression
		expected        Collection
		expectedOrdered bool
	}{
		{
			name:            "where()",
			fn:              defaultFunctions["where"],
			target:          Collection{Integer(1), Integer(2), Integer(3), Integer(4)},
			params:          []Expression{MustParse("$this > 2")},
			expected:        Collection{Integer(3), Integer(4)},
			expectedOrdered: true,
		},
		{
			name:            "select()",
			fn:              defaultFunctions["select"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          []Expression{MustParse("$this * 2")},
			expected:        Collection{Integer(2), Integer(4), Integer(6)},
			expectedOrdered: true,
		},
		{
			name:            "repeat()",
			fn:              defaultFunctions["repeat"],
			target:          Collection{String("test")},
			params:          []Expression{MustParse("children()")},
			expected:        Collection{},
			expectedOrdered: false,
		},
		{
			name:            "ofType()",
			fn:              defaultFunctions["ofType"],
			target:          Collection{String("test"), Integer(1)},
			params:          []Expression{MustParse("System.String")},
			expected:        Collection{String("test")},
			expectedOrdered: true,
		},
		{
			name:            "where() with empty result",
			fn:              defaultFunctions["where"],
			target:          Collection{Integer(1), Integer(2)},
			params:          []Expression{MustParse("$this > 5")},
			expected:        Collection{},
			expectedOrdered: true,
		},
		{
			name:            "select() with empty collection",
			fn:              defaultFunctions["select"],
			target:          Collection{},
			params:          []Expression{MustParse("$this * 2")},
			expected:        Collection{},
			expectedOrdered: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, tt.expectedOrdered, false)
		})
	}
}

func TestSubsettingFunctions(t *testing.T) {
	tests := []struct {
		name            string
		fn              Function
		target          Collection
		params          []Expression
		expected        Collection
		expectedOrdered bool
	}{
		{
			name:            "first()",
			fn:              defaultFunctions["first"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          nil,
			expected:        Collection{Integer(1)},
			expectedOrdered: true,
		},
		{
			name:            "last()",
			fn:              defaultFunctions["last"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          nil,
			expected:        Collection{Integer(3)},
			expectedOrdered: true,
		},
		{
			name:            "tail()",
			fn:              defaultFunctions["tail"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          nil,
			expected:        Collection{Integer(2), Integer(3)},
			expectedOrdered: true,
		},
		{
			name:            "skip()",
			fn:              defaultFunctions["skip"],
			target:          Collection{Integer(1), Integer(2), Integer(3), Integer(4)},
			params:          []Expression{MustParse("2")},
			expected:        Collection{Integer(3), Integer(4)},
			expectedOrdered: true,
		},
		{
			name:            "take()",
			fn:              defaultFunctions["take"],
			target:          Collection{Integer(1), Integer(2), Integer(3), Integer(4)},
			params:          []Expression{MustParse("2")},
			expected:        Collection{Integer(1), Integer(2)},
			expectedOrdered: true,
		},
		{
			name:            "intersect()",
			fn:              defaultFunctions["intersect"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          []Expression{MustParse("2 | 3 | 4")},
			expected:        Collection{Integer(2), Integer(3)},
			expectedOrdered: false,
		},
		{
			name:            "exclude()",
			fn:              defaultFunctions["exclude"],
			target:          Collection{Integer(1), Integer(2), Integer(3), Integer(4)},
			params:          []Expression{MustParse("2 | 3")},
			expected:        Collection{Integer(1), Integer(4)},
			expectedOrdered: true,
		},
		{
			name:            "first() with empty collection",
			fn:              defaultFunctions["first"],
			target:          Collection{},
			params:          nil,
			expected:        Collection{},
			expectedOrdered: true,
		},
		{
			name:            "last() with empty collection",
			fn:              defaultFunctions["last"],
			target:          Collection{},
			params:          nil,
			expected:        Collection{},
			expectedOrdered: true,
		},
		{
			name:            "tail() with empty collection",
			fn:              defaultFunctions["tail"],
			target:          Collection{},
			params:          nil,
			expected:        Collection{},
			expectedOrdered: true,
		},
		{
			name:            "skip() with empty collection",
			fn:              defaultFunctions["skip"],
			target:          Collection{},
			params:          []Expression{MustParse("2")},
			expected:        Collection{},
			expectedOrdered: true,
		},
		{
			name:            "take() with empty collection",
			fn:              defaultFunctions["take"],
			target:          Collection{},
			params:          []Expression{MustParse("2")},
			expected:        Collection{},
			expectedOrdered: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, tt.expectedOrdered, false)
		})
	}
}

func TestCombiningFunctions(t *testing.T) {
	tests := []struct {
		name            string
		fn              Function
		target          Collection
		params          []Expression
		expected        Collection
		expectedOrdered bool
	}{
		{
			name:            "union()",
			fn:              defaultFunctions["union"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          []Expression{MustParse("3|4|5")},
			expected:        Collection{Integer(1), Integer(2), Integer(3), Integer(4), Integer(5)},
			expectedOrdered: false,
		},
		{
			name:            "combine()",
			fn:              defaultFunctions["combine"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          []Expression{MustParse("3|4|5")},
			expected:        Collection{Integer(1), Integer(2), Integer(3), Integer(3), Integer(4), Integer(5)},
			expectedOrdered: false,
		},
		{
			name:            "union() with empty collections",
			fn:              defaultFunctions["union"],
			target:          Collection{},
			params:          []Expression{MustParse("{}")},
			expected:        Collection{},
			expectedOrdered: false,
		},
		{
			name:            "combine() with empty collections",
			fn:              defaultFunctions["combine"],
			target:          Collection{},
			params:          []Expression{MustParse("{}")},
			expected:        Collection{},
			expectedOrdered: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, tt.expectedOrdered, false)
		})
	}
}

func TestStringFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       Function
		target   Collection
		params   []Expression
		expected Collection
	}{
		{
			name:     "indexOf()",
			fn:       defaultFunctions["indexOf"],
			target:   Collection{String("hello test world")},
			params:   []Expression{MustParse("'test'")},
			expected: Collection{Integer(6)},
		},
		{
			name:     "substring()",
			fn:       defaultFunctions["substring"],
			target:   Collection{String("hello test world")},
			params:   []Expression{MustParse("6"), MustParse("4")},
			expected: Collection{String("test")},
		},
		{
			name:     "startsWith()",
			fn:       defaultFunctions["startsWith"],
			target:   Collection{String("hello world")},
			params:   []Expression{MustParse("'hello'")},
			expected: Collection{Boolean(true)},
		},
		{
			name:     "endsWith()",
			fn:       defaultFunctions["endsWith"],
			target:   Collection{String("hello world")},
			params:   []Expression{MustParse("'world'")},
			expected: Collection{Boolean(true)},
		},
		{
			name:     "contains()",
			fn:       defaultFunctions["contains"],
			target:   Collection{String("hello test world")},
			params:   []Expression{MustParse("'test'")},
			expected: Collection{Boolean(true)},
		},
		{
			name:     "upper()",
			fn:       defaultFunctions["upper"],
			target:   Collection{String("hello")},
			params:   nil,
			expected: Collection{String("HELLO")},
		},
		{
			name:     "lower()",
			fn:       defaultFunctions["lower"],
			target:   Collection{String("HELLO")},
			params:   nil,
			expected: Collection{String("hello")},
		},
		{
			name:     "replace()",
			fn:       defaultFunctions["replace"],
			target:   Collection{String("hello world")},
			params:   []Expression{MustParse("'hello'"), MustParse("'hi'")},
			expected: Collection{String("hi world")},
		},
		{
			name:     "matches()",
			fn:       defaultFunctions["matches"],
			target:   Collection{String("hello world")},
			params:   []Expression{MustParse("'^hello.*'")},
			expected: Collection{Boolean(true)},
		},
		{
			name:     "length()",
			fn:       defaultFunctions["length"],
			target:   Collection{String("hello")},
			params:   nil,
			expected: Collection{Integer(5)},
		},
		{
			name:     "toChars()",
			fn:       defaultFunctions["toChars"],
			target:   Collection{String("hi")},
			params:   nil,
			expected: Collection{String("h"), String("i")},
		},
		{
			name:     "trim()",
			fn:       defaultFunctions["trim"],
			target:   Collection{String("  hello  ")},
			params:   nil,
			expected: Collection{String("hello")},
		},
		{
			name:     "split()",
			fn:       defaultFunctions["split"],
			target:   Collection{String("hello,world")},
			params:   []Expression{MustParse("','")},
			expected: Collection{String("hello"), String("world")},
		},
		{
			name:     "join()",
			fn:       defaultFunctions["join"],
			target:   Collection{String("hello"), String("world")},
			params:   []Expression{MustParse("','")},
			expected: Collection{String("hello,world")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, true, false)
		})
	}
}

func TestMathFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       Function
		target   Collection
		params   []Expression
		expected Collection
	}{
		{
			name:     "abs()",
			fn:       defaultFunctions["abs"],
			target:   Collection{Integer(-5)},
			params:   nil,
			expected: Collection{Integer(5)},
		},
		{
			name:     "ceiling()",
			fn:       defaultFunctions["ceiling"],
			target:   Collection{Decimal{Value: apd.New(5, -1)}}, // 0.5
			params:   nil,
			expected: Collection{Integer(1)},
		},
		{
			name:     "exp()",
			fn:       defaultFunctions["exp"],
			target:   Collection{Decimal{Value: apd.New(1, 0)}}, // 1
			params:   nil,
			expected: Collection{Decimal{Value: apd.New(2718281828, -9)}}, // e ≈ 2.718281828
		},
		{
			name:     "floor()",
			fn:       defaultFunctions["floor"],
			target:   Collection{Decimal{Value: apd.New(5, -1)}}, // 0.5
			params:   nil,
			expected: Collection{Integer(0)},
		},
		{
			name:     "ln()",
			fn:       defaultFunctions["ln"],
			target:   Collection{Decimal{Value: apd.New(1, 0)}}, // 1
			params:   nil,
			expected: Collection{Decimal{Value: apd.New(0, 0)}}, // ln(1) = 0
		},
		{
			name:     "log()",
			fn:       defaultFunctions["log"],
			target:   Collection{Decimal{Value: apd.New(100, 0)}}, // 100
			params:   []Expression{MustParse("10")},
			expected: Collection{Decimal{Value: apd.New(2, 0)}}, // log10(100) = 2
		},
		{
			name:     "power()",
			fn:       defaultFunctions["power"],
			target:   Collection{Decimal{Value: apd.New(3, 0)}}, // 3
			params:   []Expression{MustParse("2")},
			expected: Collection{Decimal{Value: apd.New(9, 0)}}, // 3^2 = 9
		},
		{
			name:     "round()",
			fn:       defaultFunctions["round"],
			target:   Collection{Decimal{Value: apd.New(5, -1)}}, // 0.5
			params:   nil,
			expected: Collection{Decimal{Value: apd.New(1, 0)}},
		},
		{
			name:     "sqrt()",
			fn:       defaultFunctions["sqrt"],
			target:   Collection{Decimal{Value: apd.New(4, 0)}}, // 4
			params:   nil,
			expected: Collection{Decimal{Value: apd.New(2, 0)}}, // √4 = 2
		},
		{
			name:     "truncate()",
			fn:       defaultFunctions["truncate"],
			target:   Collection{Decimal{Value: apd.New(5, -1)}}, // 0.5
			params:   nil,
			expected: Collection{Integer(0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, true, false)
		})
	}
}

func TestTreeNavigationFunctions(t *testing.T) {
	tests := []struct {
		name            string
		fn              Function
		target          Collection
		params          []Expression
		expected        Collection
		expectedOrdered bool
	}{
		{
			name:            "children()",
			fn:              defaultFunctions["children"],
			target:          Collection{testElement{value: []any{}}},
			params:          nil,
			expected:        Collection{},
			expectedOrdered: false,
		},
		{
			name:            "descendants()",
			fn:              defaultFunctions["descendants"],
			target:          Collection{testElement{value: []any{}}},
			params:          nil,
			expected:        Collection{},
			expectedOrdered: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, tt.expectedOrdered, false)
		})
	}
}

func TestUtilityFunctions(t *testing.T) {
	tests := []struct {
		name            string
		fn              Function
		target          Collection
		params          []Expression
		expected        Collection
		expectedOrdered bool
	}{
		{
			name:            "trace()",
			fn:              defaultFunctions["trace"],
			target:          Collection{String("value")},
			params:          []Expression{MustParse("'test'")},
			expected:        Collection{String("value")},
			expectedOrdered: true,
		},
		{
			name:            "now()",
			fn:              defaultFunctions["now"],
			target:          Collection{},
			params:          nil,
			expected:        Collection{DateTime{Value: time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC), Precision: DateTimePrecisionFull}},
			expectedOrdered: true,
		},
		{
			name:            "today()",
			fn:              defaultFunctions["today"],
			target:          Collection{},
			params:          nil,
			expected:        Collection{Date{Value: time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC), Precision: DatePrecisionFull}},
			expectedOrdered: true,
		},
		{
			name:            "extension()",
			fn:              defaultFunctions["extension"],
			target:          Collection{String("test")},
			params:          []Expression{MustParse("'test'")},
			expected:        Collection{},
			expectedOrdered: true,
		},
		{
			name:            "isDistinct()",
			fn:              defaultFunctions["isDistinct"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          nil,
			expected:        Collection{Boolean(true)},
			expectedOrdered: true,
		},
		{
			name:            "distinct()",
			fn:              defaultFunctions["distinct"],
			target:          Collection{Integer(1), Integer(2), Integer(2), Integer(3)},
			params:          nil,
			expected:        Collection{Integer(1), Integer(2), Integer(3)},
			expectedOrdered: false,
		},
		{
			name:            "count()",
			fn:              defaultFunctions["count"],
			target:          Collection{Integer(1), Integer(2), Integer(3)},
			params:          nil,
			expected:        Collection{Integer(3)},
			expectedOrdered: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, tt.expectedOrdered, false)
		})
	}
}

func TestDefineVariable(t *testing.T) {
	tests := []struct {
		name        string
		target      Collection
		params      []Expression
		expected    Collection
		expectError bool
		varName     string
		varValue    Element
	}{
		{
			name:        "define_variable_with_value",
			target:      Collection{},
			params:      []Expression{MustParse("'test'"), MustParse("'value'")},
			expected:    Collection{},
			expectError: false,
			varName:     "test",
			varValue:    String("value"),
		},
		{
			name:        "define_variable_using_input_collection",
			target:      Collection{String("inputValue")},
			params:      []Expression{MustParse("'myVar'")},
			expected:    Collection{String("inputValue")},
			expectError: false,
			varName:     "myVar",
			varValue:    String("inputValue"),
		},
		{
			name:        "invalid_number_of_parameters",
			target:      Collection{},
			params:      []Expression{MustParse("'myVar'"), MustParse("'test'"), MustParse("'extra'")},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "non_string_name_parameter",
			target:      Collection{},
			params:      []Expression{MustParse("123"), MustParse("'test'")},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "multiple_values_in_input_collection",
			target:      Collection{String("test1"), String("test2")},
			params:      []Expression{MustParse("'myVar'")},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockEvaluate := func(ctx context.Context, target Element, expr Expression, fnScope ...FunctionScope) (Collection, bool, error) {
				if expr.tree == nil {
					return Collection{}, false, fmt.Errorf("unexpected expression <nil>")
				}
				switch expr.String() {
				case "'test'":
					return Collection{String("test")}, true, nil
				case "'myVar'":
					return Collection{String("myVar")}, true, nil
				case "'testVar'":
					return Collection{String("testVar")}, true, nil
				case "'value'":
					return Collection{String("value")}, true, nil
				case "'invalid'":
					return Collection{String("invalid")}, true, nil
				case "'multiple'":
					return Collection{String("multiple")}, true, nil
				default:
					return Collection{}, false, fmt.Errorf("evaluate not implemented for expression: %s", expr.String())
				}
			}
			ctx := context.Background()
			ctx, _ = withNewEnvStackFrame(ctx)
			result, _, err := defaultFunctions["defineVariable"](ctx, testElement{}, tt.target, true, tt.params, mockEvaluate)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if !cmp.Equal(result, tt.expected, cmpopts.IgnoreUnexported(testElement{})) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}

			// Verify variable was set in environment
			if !tt.expectError {
				val, exists := envValue(ctx, tt.varName)
				if !exists {
					t.Errorf("Variable %s was not set in environment", tt.varName)
				} else if val != tt.varValue {
					t.Errorf("Expected variable value '%s', got '%s'", tt.varValue, val.String())
				}
			}
		})
	}
}

func TestTypeFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       Function
		target   Collection
		params   []Expression
		expected Collection
	}{
		{
			name:     "is()",
			fn:       defaultFunctions["is"],
			target:   Collection{String("test")},
			params:   []Expression{MustParse("System.String")},
			expected: Collection{Boolean(true)},
		},
		{
			name:     "as()",
			fn:       defaultFunctions["as"],
			target:   Collection{String("test")},
			params:   []Expression{MustParse("System.String")},
			expected: Collection{String("test")},
		},
		{
			name:     "ofType()",
			fn:       defaultFunctions["ofType"],
			target:   Collection{String("test"), Integer(1)},
			params:   []Expression{MustParse("System.String")},
			expected: Collection{String("test")},
		},
		{
			name:   "type()",
			fn:     defaultFunctions["type"],
			target: Collection{String("test")},
			params: nil,
			expected: Collection{SimpleTypeInfo{
				Namespace: "System",
				Name:      "String",
				BaseType:  TypeSpecifier{Namespace: "System", Name: "Any"},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFunction(t, tt.fn, tt.target, tt.params, tt.expected, true, false)
		})
	}
}
