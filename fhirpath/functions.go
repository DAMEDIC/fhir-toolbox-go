package fhirpath

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"maps"
	"math"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/cockroachdb/apd/v3"
)

// Tracer defines the interface for logging trace messages
type Tracer interface {
	// Log logs a trace message with the given name and collection
	Log(name string, collection Collection) error
}

// StdoutTracer writes traces to io.Stdout.
type StdoutTracer struct{}

func (w StdoutTracer) Log(name string, collection Collection) error {
	_, err := fmt.Printf("%s: %v\n", name, collection)
	return err
}

type tracerKey struct{}

// WithTracer installs the given trace logger into the context.
//
// By default, traces are logged to stdout.
// To redirect trace logs to a custom output, use:
//
//	ctx = fhirpath.WithTracer(ctx, MyCustomTraceLogger(true, file))
func WithTracer(ctx context.Context, logger Tracer) context.Context {
	return context.WithValue(ctx, tracerKey{}, logger)
}

// GetTraceLogger gets the trace logger from the context
// If no trace logger is found, a NoOpTraceLogger is returned
func tracer(ctx context.Context) (Tracer, error) {
	logger, ok := ctx.Value(tracerKey{}).(Tracer)
	if !ok {
		return StdoutTracer{}, nil
	}
	if logger == nil {
		return StdoutTracer{}, fmt.Errorf("no trace logger provided")
	}
	return logger, nil
}

type Functions map[string]Function
type Function = func(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	parameters []Expression,
	evaluate EvaluateFunc,
) (result Collection, resultOrdered bool, err error)

type EvaluateFunc = func(
	ctx context.Context,
	target Element,
	expr Expression,
	fnScope ...FunctionScope,
) (result Collection, resultOrdered bool, err error)

type functionCtxKey struct{}

type FunctionScope struct {
	index int
	total Collection
}

type functionScope struct {
	this      Element
	index     int
	aggregate bool
	total     Collection
}

func withFunctionScope(
	ctx context.Context,
	fnScope functionScope,
) context.Context {
	return context.WithValue(
		ctx,
		functionCtxKey{},
		fnScope,
	)
}

func getFunctionScope(ctx context.Context) (functionScope, error) {
	fnCtx, ok := ctx.Value(functionCtxKey{}).(functionScope)
	if !ok {
		return functionScope{}, fmt.Errorf("not in function context")
	}
	return fnCtx, nil
}

type functionsKey struct{}

// WithFunctions installs the given functions into the context.
func WithFunctions(
	ctx context.Context,
	functions Functions,
) context.Context {
	allFns := getFunctions(ctx)

	maps.Copy(allFns, functions)

	return context.WithValue(ctx, functionsKey{}, allFns)
}

func getFunctions(ctx context.Context) Functions {
	fns, ok := ctx.Value(functionsKey{}).(Functions)
	if !ok {
		return maps.Clone(defaultFunctions)
	}
	return fns
}

func getFunction(ctx context.Context, name string) (Function, bool) {
	fns := getFunctions(ctx)
	fn, ok := fns[name]
	return fn, ok
}

// global variable for mocking time in tests
var now = time.Now

var defaultFunctions = Functions{
	// Type functions
	"type": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		result = make(Collection, 0, len(target))
		for _, elem := range target {
			result = append(result, elem.TypeInfo())
		}

		return result, inputOrdered, nil
	},
	"is": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) != 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single type argument")
		}
		typeSpec := ParseTypeSpecifier(parameters[0].String())

		r, err := isType(ctx, target[0], typeSpec)
		if err != nil {
			return nil, false, err
		}
		return Collection{r}, true, nil
	},
	"as": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) != 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single type specifier parameter")
		}
		typeSpec := ParseTypeSpecifier(parameters[0].String())

		c, err := asType(ctx, target[0], typeSpec)
		if err != nil {
			return nil, false, err
		}
		return c, true, nil
	},
	"ofType": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single type specifier parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		typeSpec := ParseTypeSpecifier(parameters[0].String())

		for _, elem := range target {
			isOfType, err := isType(ctx, elem, typeSpec)
			if err != nil {
				return nil, false, err
			}

			// Check if isOfType is a Boolean with value true
			if boolVal, ok := isOfType.(Boolean); ok && bool(boolVal) {
				result = append(result, elem)
			}
		}

		return result, inputOrdered, nil
	},

	// Boolean functions
	"not": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameter")
		}
		b, ok, err := Singleton[Boolean](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}
		return Collection{!b}, true, nil
	},
	"empty": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}
		return Collection{Boolean(len(target) == 0)}, true, nil
	},
	"exists": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) > 1 {
			return nil, false, fmt.Errorf("expected at most one criteria parameter")
		}

		if len(parameters) == 0 {
			return Collection{Boolean(len(target) > 0)}, true, nil
		}

		// With criteria, equivalent to where(criteria).exists()
		for i, elem := range target {
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i})
			if err != nil {
				return nil, false, err
			}

			b, ok, err := Singleton[Boolean](criteria)
			if err != nil {
				return nil, false, err
			}
			if ok && bool(b) {
				// Found at least one element that matches the criteria
				return Collection{Boolean(true)}, true, nil
			}
		}

		return Collection{Boolean(false)}, true, nil
	},
	"all": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single criteria parameter")
		}

		// If the input collection is empty, the result is true
		if len(target) == 0 {
			return Collection{Boolean(true)}, true, nil
		}

		for i, elem := range target {
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i})
			if err != nil {
				return nil, false, err
			}

			b, ok, err := Singleton[Boolean](criteria)
			if err != nil {
				return nil, false, err
			}
			if !ok || !bool(b) {
				// Found at least one element that doesn't match the criteria
				return Collection{Boolean(false)}, true, nil
			}
		}

		return Collection{Boolean(true)}, true, nil
	},
	"allTrue": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is true
		if len(target) == 0 {
			return Collection{Boolean(true)}, true, nil
		}

		for _, elem := range target {
			b, ok, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if !ok || !bool(b) {
				return Collection{Boolean(false)}, true, nil
			}
		}
		return Collection{Boolean(true)}, true, nil
	},
	"anyTrue": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is false
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		}

		for _, elem := range target {
			b, ok, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if ok && bool(b) {
				return Collection{Boolean(true)}, true, nil
			}
		}
		return Collection{Boolean(false)}, true, nil
	},
	"allFalse": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is true
		if len(target) == 0 {
			return Collection{Boolean(true)}, true, nil
		}

		for _, elem := range target {
			b, ok, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if !ok || bool(b) {
				return Collection{Boolean(false)}, true, nil
			}
		}
		return Collection{Boolean(true)}, true, nil
	},
	"anyFalse": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is false
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		}

		for _, elem := range target {
			b, ok, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if ok && !bool(b) {
				return Collection{Boolean(true)}, true, nil
			}
		}
		return Collection{Boolean(false)}, true, nil
	},

	// Collection functions
	"subsetOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single collection parameter")
		}

		// If the input collection is empty, the result is true
		if len(target) == 0 {
			return Collection{Boolean(true)}, true, nil
		}

		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the other collection is empty, the result is false
		if len(other) == 0 {
			return Collection{Boolean(false)}, true, nil
		}

		for _, elem := range target {
			found := false
			for _, otherElem := range other {
				eq, ok := elem.Equal(otherElem)
				if ok && eq {
					found = true
					break
				}
			}
			if !found {
				return Collection{Boolean(false)}, true, nil
			}
		}
		return Collection{Boolean(true)}, true, nil
	},
	"supersetOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single collection parameter")
		}

		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the other collection is empty, the result is true
		if len(other) == 0 {
			return Collection{Boolean(true)}, true, nil
		}

		// If the input collection is empty, the result is false
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		}

		for _, otherElem := range other {
			found := false
			for _, elem := range target {
				eq, ok := otherElem.Equal(elem)
				if ok && eq {
					found = true
					break
				}
			}
			if !found {
				return Collection{Boolean(false)}, true, nil
			}
		}
		return Collection{Boolean(true)}, true, nil
	},
	"count": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}
		return Collection{Integer(len(target))}, true, nil
	},
	"distinct": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		for _, elem := range target {
			found := false
			for _, resultElem := range result {
				eq, ok := elem.Equal(resultElem)
				if ok && eq {
					found = true
					break
				}
			}
			if !found {
				result = append(result, elem)
			}
		}
		return result, false, nil
	},
	"isDistinct": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is true
		if len(target) == 0 {
			return Collection{Boolean(true)}, true, nil
		}

		// Check if all elements are distinct
		for i := 0; i < len(target); i++ {
			for j := i + 1; j < len(target); j++ {
				eq, ok := target[i].Equal(target[j])
				if ok && eq {
					return Collection{Boolean(false)}, true, nil
				}
			}
		}
		return Collection{Boolean(true)}, true, nil
	},
	"where": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single criteria parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		for i, elem := range target {
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i})
			if err != nil {
				return nil, false, err
			}

			b, ok, err := Singleton[Boolean](criteria)
			if err != nil {
				return nil, false, err
			}
			if ok && bool(b) {
				// Element matches the criteria, add it to the result
				result = append(result, elem)
			}
		}

		return result, true, nil
	},
	"select": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single projection parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		resultOrdered = inputOrdered
		for i, elem := range target {
			projection, ordered, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i})
			if err != nil {
				return nil, false, err
			}

			// Add all items from the projection to the result (flatten)
			result = append(result, projection...)

			if !ordered {
				resultOrdered = false
			}
		}

		return result, resultOrdered, nil
	},
	// sort implements FHIRPath v3.0.0 §4.1.26 sort()
	"sort": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, true, nil
		}

		type sortKeyValue struct {
			empty bool
			value Element
		}
		type sortItem struct {
			elem Element
			keys []sortKeyValue
		}

		items := make([]sortItem, len(target))
		for i, elem := range target {
			items[i].elem = elem
			if len(parameters) == 0 {
				continue
			}
			items[i].keys = make([]sortKeyValue, len(parameters))
			for j, param := range parameters {
				keyResult, _, err := evaluate(ctx, elem, param, FunctionScope{index: i})
				if err != nil {
					return nil, false, err
				}

				switch len(keyResult) {
				case 0:
					items[i].keys[j] = sortKeyValue{empty: true}
				case 1:
					items[i].keys[j] = sortKeyValue{value: keyResult[0]}
				default:
					return nil, false, fmt.Errorf(
						"sort key %d evaluated to %d items (expected 0 or 1)",
						j+1, len(keyResult),
					)
				}
			}
		}

		var sortErr error
		slices.SortStableFunc(items, func(a, b sortItem) int {
			if sortErr != nil {
				return 0
			}

			if len(parameters) == 0 {
				cmp, err := compareElementsForSort(a.elem, b.elem)
				if err != nil {
					sortErr = err
					return 0
				}
				return cmp
			}

			for idx, param := range parameters {
				dir := param.sortDirection
				if dir == sortDirectionNone {
					dir = sortDirectionAsc
				}

				av := a.keys[idx]
				bv := b.keys[idx]

				if av.empty && bv.empty {
					continue
				}
				if av.empty {
					return -1
				}
				if bv.empty {
					return 1
				}

				cmp, err := compareElementsForSort(av.value, bv.value)
				if err != nil {
					sortErr = err
					return 0
				}
				if cmp != 0 {
					if dir == sortDirectionDesc {
						cmp = -cmp
					}
					return cmp
				}
			}

			return 0
		})
		if sortErr != nil {
			return nil, false, sortErr
		}

		result = make(Collection, len(items))
		for i, item := range items {
			result[i] = item.elem
		}
		return result, true, nil
	},
	"repeat": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single projection parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		var current = target
		var newItems Collection

		// Keep repeating the projection until no new items are found
		for {
			newItems = nil
			for i, elem := range current {
				projection, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i})
				if err != nil {
					return nil, false, err
				}

				// FHIRPath 3.0: Check for new items against both result AND newItems
				// Items should only be added if not already in the output collection
				for _, item := range projection {
					add := true
					// Check against already accumulated results
					for _, seen := range result {
						eq, ok := seen.Equal(item)
						if ok && eq {
							add = false
							break
						}
					}
					// Also check against items found in this iteration
					if add {
						for _, seen := range newItems {
							eq, ok := seen.Equal(item)
							if ok && eq {
								add = false
								break
							}
						}
					}
					if add {
						newItems = append(newItems, item)
					}
				}
			}

			// If no new items were found, we're done
			if len(newItems) == 0 {
				break
			}

			// Add new items to the result and set them as the current items for the next iteration
			result = append(result, newItems...)
			current = newItems
		}

		return result, false, nil
	},
	// repeatAll implements FHIRPath v3.0.0 §4.1.26 repeatAll()
	"repeatAll": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single projection parameter")
		}
		if len(target) == 0 {
			return nil, true, nil
		}

		queue := slices.Clone(target)

		for len(queue) > 0 {
			var next Collection
			for i, elem := range queue {
				projection, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i})
				if err != nil {
					return nil, false, err
				}
				if len(projection) == 0 {
					continue
				}
				result = append(result, projection...)
				next = append(next, projection...)
			}
			queue = next
		}

		return result, false, nil
	},
	"single": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// If there are multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Return the single item
		return Collection{target[0]}, true, nil
	},
	"first": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Return the first item
		return Collection{target[0]}, true, nil
	},
	"last": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Return the last item
		return Collection{target[len(target)-1]}, true, nil
	},
	"tail": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection has no items or only one item, the result is empty
		if len(target) <= 1 {
			return nil, true, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Return all but the first item
		return target[1:], inputOrdered, nil
	},
	"skip": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single num parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Evaluate the num parameter
		numCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert to integer
		num, ok, err := Singleton[Integer](numCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected integer parameter")
		}

		// If num is less than or equal to zero, return the input collection
		if num <= 0 {
			return target, inputOrdered, nil
		}

		// If num is greater than or equal to the length of the collection, return empty
		if int(num) >= len(target) {
			return nil, true, nil
		}

		// Return all but the first num items
		return target[num:], inputOrdered, nil
	},
	"take": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single num parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Evaluate the num parameter
		numCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert to integer
		num, ok, err := Singleton[Integer](numCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected integer parameter")
		}

		// If num is less than or equal to zero, return empty
		if num <= 0 {
			return nil, true, nil
		}

		// If num is greater than the length of the collection, return the whole collection
		if int(num) >= len(target) {
			return target, inputOrdered, nil
		}

		// Return the first num items
		return target[:num], inputOrdered, nil
	},
	"intersect": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single collection parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Evaluate the other collection parameter
		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the other collection is empty, the result is empty
		if len(other) == 0 {
			return nil, true, nil
		}

		// Find the intersection of the two collections
		for _, elem := range target {
			// Check if the element is in the other collection
			for _, otherElem := range other {
				eq, ok := elem.Equal(otherElem)
				if ok && eq {
					// Check if the element is already in the result (eliminate duplicates)
					found := false
					for _, resultElem := range result {
						eq, ok := elem.Equal(resultElem)
						if ok && eq {
							found = true
							break
						}
					}
					if !found {
						result = append(result, elem)
					}
					break
				}
			}
		}

		return result, false, nil
	},
	"exclude": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single collection parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Evaluate the other collection parameter
		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the other collection is empty, the result is the input collection
		if len(other) == 0 {
			return target, inputOrdered, nil
		}

		// Find the elements in the input collection that are not in the other collection
		for _, elem := range target {
			// Check if the element is in the other collection
			found := false
			for _, otherElem := range other {
				eq, ok := elem.Equal(otherElem)
				if ok && eq {
					found = true
					break
				}
			}
			if !found {
				// Element is not in the other collection, add it to the result
				result = append(result, elem)
			}
		}

		return result, inputOrdered, nil
	},
	"union": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single collection parameter")
		}

		// Evaluate the other collection parameter
		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Use the Union method to merge the collections
		return target.Union(other), false, nil
	},
	"combine": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single collection parameter")
		}

		// Evaluate the other collection parameter
		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Use the Combine method to merge the collections
		return target.Combine(other), false, nil
	},
	// coalesce implements FHIRPath v3.0.0 §4.1.26 coalesce()
	"coalesce": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) == 0 {
			return nil, false, fmt.Errorf("expected at least one collection parameter")
		}

		for _, param := range parameters {
			value, ordered, err := evaluate(ctx, nil, param)
			if err != nil {
				return nil, false, err
			}
			if len(value) > 0 {
				return value, ordered, nil
			}
		}

		return nil, true, nil
	},

	// String functions
	"indexOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single substring parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the substring parameter
		substringCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert substring to string
		substring, ok, err := Singleton[String](substringCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			// If substring is empty/null, return empty collection
			return nil, true, nil
		}

		// If substring is an empty string (''), the function returns 0
		if substring == "" {
			return Collection{Integer(0)}, true, nil
		}

		// Return the index of the substring in the string
		index := strings.Index(string(s), string(substring))
		return Collection{Integer(index)}, true, nil
	},
	"lastIndexOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single substring parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the substring parameter
		substringCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert substring to string
		substring, ok, err := Singleton[String](substringCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			// If substring is empty/null, return empty collection
			return nil, true, nil
		}

		// If substring is an empty string (''), the function returns the length of the input string
		if substring == "" {
			return Collection{Integer(len([]rune(s)))}, true, nil
		}

		// Return the index of the last occurrence of the substring in the string
		index := strings.LastIndex(string(s), string(substring))
		return Collection{Integer(index)}, true, nil
	},
	"substring": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) < 1 || len(parameters) > 2 {
			return nil, false, fmt.Errorf("expected one or two parameters (start, [length])")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the start parameter
		startCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert start to integer
		start, ok, err := Singleton[Integer](startCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected integer start parameter")
		}

		// If start is negative or greater than or equal to the length of the string, return empty
		if start < 0 || int(start) >= len(s) {
			return nil, true, nil
		}

		// If length parameter is provided
		if len(parameters) == 2 {
			// Evaluate the length parameter
			lengthCollection, _, err := evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}

			// Convert length to integer
			length, ok, err := Singleton[Integer](lengthCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected integer length parameter")
			}

			// If length is negative, treat it as 0
			if length < 0 {
				length = 0
			}

			// Calculate end index (start + length)
			end := int(start) + int(length)
			if end > len(s) {
				end = len(s)
			}

			result := String(s[start:end])
			return Collection{result}, true, nil
		}

		// If length parameter is not provided, return the rest of the string
		return Collection{String(s[start:])}, true, nil
	},
	"startsWith": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single prefix parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the prefix parameter
		prefixCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert prefix to string
		prefix, ok, err := Singleton[String](prefixCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string prefix parameter")
		}

		// If prefix is an empty string (''), the result is true
		if prefix == "" {
			return Collection{Boolean(true)}, true, nil
		}

		// Check if the string starts with the prefix
		startsWith := strings.HasPrefix(string(s), string(prefix))
		return Collection{Boolean(startsWith)}, true, nil
	},
	"endsWith": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single suffix parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the suffix parameter
		suffixCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert suffix to string
		suffix, ok, err := Singleton[String](suffixCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string suffix parameter")
		}

		// If suffix is an empty string (''), the result is true
		if suffix == "" {
			return Collection{Boolean(true)}, true, nil
		}

		// Check if the string ends with the suffix
		endsWith := strings.HasSuffix(string(s), string(suffix))
		return Collection{Boolean(endsWith)}, true, nil
	},
	"contains": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single substring parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the substring parameter
		substringCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert substring to string
		substring, ok, err := Singleton[String](substringCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string substring parameter")
		}

		// If substring is an empty string (''), the result is true
		if substring == "" {
			return Collection{Boolean(true)}, true, nil
		}

		// Check if the string contains the substring
		contains := strings.Contains(string(s), string(substring))
		return Collection{Boolean(contains)}, true, nil
	},
	"upper": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Convert the string to upper case
		return Collection{String(strings.ToUpper(string(s)))}, true, nil
	},
	"lower": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Convert the string to lower case
		return Collection{String(strings.ToLower(string(s)))}, true, nil
	},
	"replace": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 2 {
			return nil, false, fmt.Errorf("expected two parameters (pattern, substitution)")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate pattern and substitution parameters against the input string
		var evalTarget Element
		if len(target) > 0 {
			evalTarget = target[0]
		}

		// Evaluate the pattern parameter
		patternCollection, _, err := evaluate(ctx, evalTarget, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert pattern to string
		pattern, ok, err := Singleton[String](patternCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string pattern parameter")
		}

		// Evaluate the substitution parameter
		substitutionCollection, _, err := evaluate(ctx, evalTarget, parameters[1])
		if err != nil {
			return nil, false, err
		}

		// Convert substitution to string
		substitution, ok, err := Singleton[String](substitutionCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string substitution parameter")
		}

		// If pattern is an empty string (''), every character in the input string is surrounded by the substitution
		if pattern == "" {
			var result strings.Builder
			result.WriteString(string(substitution))
			for _, c := range s {
				result.WriteRune(c)
				result.WriteString(string(substitution))
			}
			return Collection{String(result.String())}, true, nil
		}

		// Replace all instances of pattern with substitution
		return Collection{String(strings.ReplaceAll(string(s), string(pattern), string(substitution)))}, true, nil
	},
	"matches": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) < 1 || len(parameters) > 2 {
			return nil, false, fmt.Errorf("expected regex parameter and optional flags parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the regex parameter
		regexCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If regex is empty collection, return empty
		if len(regexCollection) == 0 {
			return nil, true, nil
		}

		// Convert regex to string
		regexStr, ok, err := Singleton[String](regexCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate optional flags parameter
		var flags string
		if len(parameters) == 2 {
			flagsCollection, _, err := evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}

			flagsStr, ok, err := Singleton[String](flagsCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected string flags parameter")
			}
			flags = string(flagsStr)
		}

		// Apply flags to the regular expression
		// Per FHIRPath spec, regex operates in single-line mode by default (. matches newlines)
		pattern := "(?s)" + string(regexStr)
		for _, flag := range flags {
			switch flag {
			case 'i':
				// Case-insensitive: wrap pattern with (?i)
				pattern = "(?i)" + pattern
			case 'm':
				// Multiline mode: wrap pattern with (?m)
				pattern = "(?m)" + pattern
			default:
				return nil, false, fmt.Errorf("unsupported regex flag: %c", flag)
			}
		}

		// Compile the regular expression
		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, false, fmt.Errorf("invalid regular expression: %v", err)
		}

		// Check if the string matches the regular expression
		return Collection{Boolean(regex.MatchString(string(s)))}, true, nil
	},
	"replaceMatches": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) < 2 || len(parameters) > 3 {
			return nil, false, fmt.Errorf("expected regex, substitution, and optional flags parameters")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the regex parameter
		regexCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If regex is empty collection, return empty
		if len(regexCollection) == 0 {
			return nil, true, nil
		}

		// Convert regex to string
		regexStr, ok, err := Singleton[String](regexCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// If regex is an empty string, return input unchanged per spec
		if regexStr == "" {
			return Collection{s}, true, nil
		}

		// Evaluate the substitution parameter
		substitutionCollection, _, err := evaluate(ctx, nil, parameters[1])
		if err != nil {
			return nil, false, err
		}

		// If substitution is empty collection, return empty
		if len(substitutionCollection) == 0 {
			return nil, true, nil
		}

		// Convert substitution to string
		substitution, ok, err := Singleton[String](substitutionCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate optional flags parameter
		var flags string
		if len(parameters) == 3 {
			flagsCollection, _, err := evaluate(ctx, nil, parameters[2])
			if err != nil {
				return nil, false, err
			}

			flagsStr, ok, err := Singleton[String](flagsCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected string flags parameter")
			}
			flags = string(flagsStr)
		}

		// Apply flags to the regular expression
		// Per FHIRPath spec, regex operates in single-line mode by default (. matches newlines)
		pattern := "(?s)" + string(regexStr)
		for _, flag := range flags {
			switch flag {
			case 'i':
				// Case-insensitive: wrap pattern with (?i)
				pattern = "(?i)" + pattern
			case 'm':
				// Multiline mode: wrap pattern with (?m)
				pattern = "(?m)" + pattern
			default:
				return nil, false, fmt.Errorf("unsupported regex flag: %c", flag)
			}
		}

		// Compile the regular expression
		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, false, fmt.Errorf("invalid regular expression: %v", err)
		}

		// Replace all matches of the regular expression with the substitution
		return Collection{String(regex.ReplaceAllString(string(s), string(substitution)))}, true, nil
	},
	"length": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Return the length of the string
		return Collection{Integer(len(s))}, true, nil
	},
	"toChars": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Convert the string to a collection of characters
		chars := make(Collection, len(s))
		for i, c := range string(s) {
			chars[i] = String(c)
		}
		return chars, true, nil
	},
	"matchesFull": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single regex parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the regex parameter
		regexCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert regex to string
		regexStr, ok, err := Singleton[String](regexCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string regex parameter")
		}

		// Compile the regular expression
		regex, err := regexp.Compile("^" + string(regexStr) + "$")
		if err != nil {
			return nil, false, fmt.Errorf("invalid regular expression: %v", err)
		}

		// Check if the string matches the regular expression exactly
		return Collection{Boolean(regex.MatchString(string(s)))}, true, nil
	},
	"trim": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Trim whitespace from both ends
		return Collection{String(strings.TrimSpace(string(s)))}, true, nil
	},
	"split": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single separator parameter")
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the separator parameter
		separatorCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert separator to string
		separator, ok, err := Singleton[String](separatorCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string separator parameter")
		}

		// Split the string by the separator
		parts := strings.Split(string(s), string(separator))
		result = make(Collection, len(parts))
		for i, part := range parts {
			result[i] = String(part)
		}
		return result, true, nil
	},
	"join": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) > 1 {
			return nil, false, fmt.Errorf("expected at most one separator parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Default separator is empty string
		separator := String("")
		if len(parameters) == 1 {
			// Evaluate the separator parameter
			separatorCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			// Convert separator to string
			sep, ok, err := Singleton[String](separatorCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected string separator parameter")
			}
			separator = sep
		}

		// Convert all elements to strings
		parts := make([]string, 0, len(target))
		for _, elem := range target {
			s, ok, err := elementTo[String](elem, true)
			if err != nil || !ok {
				// Skip elements that can't be converted to strings
				continue
			}
			parts = append(parts, string(s))
		}

		// If no elements could be converted to strings, return empty
		if len(parts) == 0 {
			return nil, true, nil
		}

		// Join the strings with the separator
		return Collection{String(strings.Join(parts, string(separator)))}, true, nil
	},
	"encode": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single format parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the format parameter
		formatCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert format to string
		format, ok, err := Singleton[String](formatCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string format parameter")
		}

		// Encode according to format
		switch string(format) {
		case "hex":
			// Convert to hex using encoding/hex
			hex := hex.EncodeToString([]byte(s))
			return Collection{String(hex)}, true, nil
		case "base64":
			// Convert to base64
			encoded := base64.StdEncoding.EncodeToString([]byte(s))
			return Collection{String(encoded)}, true, nil
		case "urlbase64":
			// Convert to URL-safe base64
			encoded := base64.URLEncoding.EncodeToString([]byte(s))
			return Collection{String(encoded)}, true, nil
		default:
			return nil, false, fmt.Errorf("unsupported encoding format: %s", format)
		}
	},
	"decode": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single format parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the format parameter
		formatCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert format to string
		format, ok, err := Singleton[String](formatCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string format parameter")
		}

		// Decode according to format
		switch string(format) {
		case "hex":
			// Convert from hex
			decoded, err := hex.DecodeString(string(s))
			if err != nil {
				return nil, false, fmt.Errorf("invalid hex string: %v", err)
			}
			return Collection{String(decoded)}, true, nil
		case "base64":
			// Convert from base64
			decoded, err := base64.StdEncoding.DecodeString(string(s))
			if err != nil {
				return nil, false, fmt.Errorf("invalid base64 string: %v", err)
			}
			return Collection{String(decoded)}, true, nil
		case "urlbase64":
			// Convert from URL-safe base64
			decoded, err := base64.URLEncoding.DecodeString(string(s))
			if err != nil {
				return nil, false, fmt.Errorf("invalid URL-safe base64 string: %v", err)
			}
			return Collection{String(decoded)}, true, nil
		default:
			return nil, false, fmt.Errorf("unsupported encoding format: %s", format)
		}
	},
	"escape": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single target parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the target parameter
		targetCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert target to string
		targetStr, ok, err := Singleton[String](targetCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string target parameter")
		}

		// Escape according to target
		switch string(targetStr) {
		case "html":
			// Per FHIRPath spec: escape <, &, " and ideally anything with character encoding above 127
			var result strings.Builder
			result.Grow(len(s))
			for _, r := range string(s) {
				switch r {
				case '<':
					result.WriteString("&lt;")
				case '>':
					result.WriteString("&gt;")
				case '&':
					result.WriteString("&amp;")
				case '"':
					result.WriteString("&quot;")
				case '\'':
					result.WriteString("&#39;")
				default:
					// Escape high Unicode characters (above 127)
					if r > 127 {
						result.WriteString(fmt.Sprintf("&#%d;", r))
					} else {
						result.WriteRune(r)
					}
				}
			}
			return Collection{String(result.String())}, true, nil
		case "json":
			// Escape JSON special characters per FHIRPath spec
			// We need to escape quotes and backslashes, but NOT < > & (unlike Go's default json.Marshal)
			var result strings.Builder
			for _, r := range string(s) {
				switch r {
				case '"':
					result.WriteString(`\"`)
				case '\\':
					result.WriteString(`\\`)
				case '\n':
					result.WriteString(`\n`)
				case '\r':
					result.WriteString(`\r`)
				case '\t':
					result.WriteString(`\t`)
				case '\b':
					result.WriteString(`\b`)
				case '\f':
					result.WriteString(`\f`)
				default:
					result.WriteRune(r)
				}
			}
			return Collection{String(result.String())}, true, nil
		default:
			return nil, false, fmt.Errorf("unsupported escape target: %s", targetStr)
		}
	},
	"unescape": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single target parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// Convert input to string
		s, ok, err := Singleton[String](target)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, true, nil
		}

		// Evaluate the target parameter
		targetCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert target to string
		targetStr, ok, err := Singleton[String](targetCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string target parameter")
		}

		// Unescape according to target
		switch string(targetStr) {
		case "html":
			// Unescape HTML entities
			unescaped := html.UnescapeString(string(s))
			return Collection{String(unescaped)}, true, nil
		case "json":
			// Unescape JSON string
			// The input string may contain JSON escape sequences like \", \\, \n, \t, \r, \b, \f, \/
			// We need to interpret these escape sequences manually
			var result strings.Builder
			input := string(s)
			for i := 0; i < len(input); i++ {
				if input[i] == '\\' && i+1 < len(input) {
					// Process escape sequence
					switch input[i+1] {
					case '"':
						result.WriteByte('"')
						i++ // Skip next char
					case '\\':
						result.WriteByte('\\')
						i++
					case '/':
						result.WriteByte('/')
						i++
					case 'b':
						result.WriteByte('\b')
						i++
					case 'f':
						result.WriteByte('\f')
						i++
					case 'n':
						result.WriteByte('\n')
						i++
					case 'r':
						result.WriteByte('\r')
						i++
					case 't':
						result.WriteByte('\t')
						i++
					case 'u':
						// Unicode escape \uXXXX
						if i+5 < len(input) {
							// Parse hex digits
							hexStr := input[i+2 : i+6]
							var codePoint int
							_, err := fmt.Sscanf(hexStr, "%x", &codePoint)
							if err == nil {
								result.WriteRune(rune(codePoint))
								i += 5 // Skip u and 4 hex digits
							} else {
								// Invalid unicode escape, keep as-is
								result.WriteByte(input[i])
							}
						} else {
							// Not enough characters for unicode escape
							result.WriteByte(input[i])
						}
					default:
						// Unknown escape sequence, keep the backslash
						result.WriteByte(input[i])
					}
				} else {
					result.WriteByte(input[i])
				}
			}
			return Collection{String(result.String())}, true, nil
		default:
			return nil, false, fmt.Errorf("unsupported unescape target: %s", targetStr)
		}
	},
	"lowBoundary": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}

		// Get precision parameter if provided
		var precision Integer
		if len(parameters) == 1 {
			precisionCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			prec, ok, err := Singleton[Integer](precisionCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected integer precision parameter")
			}
			precision = prec
		}

		// Handle Decimal type
		if value, ok, err := Singleton[Decimal](target); err == nil && ok {
			// If no precision specified, use at least 8
			if len(parameters) == 0 {
				precision = 8
			}

			// If precision is greater than maximum possible, return empty
			if precision > 8 {
				return nil, true, nil
			}

			// Calculate lower boundary for decimal
			var boundary apd.Decimal
			_, err = apdContext(ctx).Floor(&boundary, value.Value)
			if err != nil {
				return nil, false, err
			}

			// Adjust precision
			if precision < 0 {
				// For negative precision, round to nearest multiple of 10^|precision|
				var factor apd.Decimal
				_, err = apdContext(ctx).Pow(&factor, apd.New(10, 0), apd.New(int64(-precision), 0))
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Quo(&boundary, &boundary, &factor)
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Floor(&boundary, &boundary)
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Mul(&boundary, &boundary, &factor)
				if err != nil {
					return nil, false, err
				}
			} else {
				// For non-negative precision, round to specified decimal places
				var factor apd.Decimal
				_, err = apdContext(ctx).Pow(&factor, apd.New(10, 0), apd.New(int64(precision), 0))
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Mul(&boundary, &boundary, &factor)
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Floor(&boundary, &boundary)
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Quo(&boundary, &boundary, &factor)
				if err != nil {
					return nil, false, err
				}
			}

			return Collection{Decimal{Value: &boundary}}, true, nil
		}

		// Handle Date type
		if value, ok, err := Singleton[Date](target); err == nil && ok {
			// If no precision specified, use at least 4
			if len(parameters) == 0 {
				precision = 4
			}

			// If precision is greater than maximum possible, return empty
			if precision > 4 {
				return nil, true, nil
			}

			// Adjust precision by truncating to the appropriate level
			result := value
			switch precision {
			case 1: // Year
				result.Precision = DatePrecisionYear
			case 2: // Month
				result.Precision = DatePrecisionMonth
			default: // Full precision
				result.Precision = DatePrecisionFull
			}
			return Collection{result}, true, nil
		}

		// Handle DateTime type
		if value, ok, err := Singleton[DateTime](target); err == nil && ok {
			// If no precision specified, use at least 6 (up to minute)
			if len(parameters) == 0 {
				precision = 6
			}

			// If precision is greater than maximum possible, return empty
			if precision > 6 {
				return nil, true, nil
			}

			// Adjust precision by truncating to the appropriate level
			result := value
			switch precision {
			case 1: // Year
				result.Precision = DateTimePrecisionYear
			case 2: // Month
				result.Precision = DateTimePrecisionMonth
			case 3: // Day
				result.Precision = DateTimePrecisionDay
			case 4: // Hour
				result.Precision = DateTimePrecisionHour
			case 5: // Minute
				result.Precision = DateTimePrecisionMinute
			default: // Full precision
				result.Precision = DateTimePrecisionFull
			}
			return Collection{result}, true, nil
		}

		// Handle Time type
		if value, ok, err := Singleton[Time](target); err == nil && ok {
			// If no precision specified, use at least 2 (up to minute)
			if len(parameters) == 0 {
				precision = 2
			}

			// If precision is greater than maximum possible, return empty
			if precision > 2 {
				return nil, true, nil
			}

			// Adjust precision by truncating to the appropriate level
			result := value
			switch precision {
			case 1: // Hour
				result.Precision = TimePrecisionHour
			case 2: // Minute
				result.Precision = TimePrecisionMinute
			default: // Full precision
				result.Precision = TimePrecisionFull
			}
			return Collection{result}, true, nil
		}

		return nil, false, fmt.Errorf("expected Decimal, Date, DateTime, or Time but got %T", target[0])
	},
	"highBoundary": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}

		// Get precision parameter if provided
		var precision Integer
		if len(parameters) == 1 {
			precisionCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			prec, ok, err := Singleton[Integer](precisionCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected integer precision parameter")
			}
			precision = prec
		}

		// Handle Decimal type
		if value, ok, err := Singleton[Decimal](target); err == nil && ok {
			// If no precision specified, use at least 8
			if len(parameters) == 0 {
				precision = 8
			}

			// If precision is greater than maximum possible, return empty
			if precision > 8 {
				return nil, true, nil
			}

			// Calculate high boundary for decimal
			var boundary apd.Decimal
			boundary.Set(value.Value)
			if precision < 0 {
				// For negative precision, round up to nearest multiple of 10^|precision| - 1
				var factor apd.Decimal
				_, err = apdContext(ctx).Pow(&factor, apd.New(10, 0), apd.New(int64(-precision), 0))
				if err != nil {
					return nil, false, err
				}
				var tmp apd.Decimal
				_, err = apdContext(ctx).Quo(&tmp, &boundary, &factor)
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Ceil(&tmp, &tmp)
				if err != nil {
					return nil, false, err
				}
				_, err = apdContext(ctx).Mul(&boundary, &tmp, &factor)
				if err != nil {
					return nil, false, err
				}
				// Subtract 1 to get the highest value in the range
				_, err = apdContext(ctx).Sub(&boundary, &boundary, apd.New(1, 0))
				if err != nil {
					return nil, false, err
				}
			} else {
				// For non-negative precision, set all digits after precision to 9
				str := boundary.Text('f')
				if dot := strings.Index(str, "."); dot != -1 {
					intPart := str[:dot]
					fracPart := str[dot+1:]
					if len(fracPart) < int(precision) {
						fracPart += strings.Repeat("0", int(precision)-len(fracPart))
					}
					fracPart = fracPart[:int(precision)]
					fracPart += strings.Repeat("9", 8-int(precision))
					str = intPart + "." + fracPart
				} else {
					str += "." + strings.Repeat("9", 8)
				}
				newVal, _, err := apd.NewFromString(str)
				if err != nil {
					return nil, false, err
				}
				boundary.Set(newVal)
			}
			return Collection{Decimal{Value: &boundary}}, true, nil
		}

		// Handle Date type
		if value, ok, err := Singleton[Date](target); err == nil && ok {
			if len(parameters) == 0 {
				precision = 4
			}
			if precision > 4 {
				return nil, true, nil
			}
			result := value
			switch precision {
			case 1: // Year
				result.Precision = DatePrecisionYear
				result.Value = time.Date(result.Value.Year(), 12, 31, 0, 0, 0, 0, result.Value.Location())
			case 2: // Month
				result.Precision = DatePrecisionMonth
				lastDay := time.Date(result.Value.Year(), result.Value.Month()+1, 0, 0, 0, 0, 0, result.Value.Location()).Day()
				result.Value = time.Date(result.Value.Year(), result.Value.Month(), lastDay, 0, 0, 0, 0, result.Value.Location())
			default: // Full precision
				result.Precision = DatePrecisionFull
			}
			return Collection{result}, true, nil
		}

		// Handle DateTime type
		if value, ok, err := Singleton[DateTime](target); err == nil && ok {
			if len(parameters) == 0 {
				precision = 6
			}
			if precision > 6 {
				return nil, true, nil
			}
			result := value
			switch precision {
			case 1: // Year
				result.Precision = DateTimePrecisionYear
				result.Value = time.Date(result.Value.Year(), 12, 31, 23, 59, 59, int(time.Millisecond*999), result.Value.Location())
			case 2: // Month
				result.Precision = DateTimePrecisionMonth
				lastDay := time.Date(result.Value.Year(), result.Value.Month()+1, 0, 0, 0, 0, 0, result.Value.Location()).Day()
				result.Value = time.Date(result.Value.Year(), result.Value.Month(), lastDay, 23, 59, 59, int(time.Millisecond*999), result.Value.Location())
			case 3: // Day
				result.Precision = DateTimePrecisionDay
				result.Value = time.Date(result.Value.Year(), result.Value.Month(), result.Value.Day(), 23, 59, 59, int(time.Millisecond*999), result.Value.Location())
			case 4: // Hour
				result.Precision = DateTimePrecisionHour
				result.Value = time.Date(result.Value.Year(), result.Value.Month(), result.Value.Day(), result.Value.Hour(), 59, 59, int(time.Millisecond*999), result.Value.Location())
			case 5: // Minute
				result.Precision = DateTimePrecisionMinute
				result.Value = time.Date(result.Value.Year(), result.Value.Month(), result.Value.Day(), result.Value.Hour(), result.Value.Minute(), 59, int(time.Millisecond*999), result.Value.Location())
			default: // Full precision
				result.Precision = DateTimePrecisionFull
			}
			return Collection{result}, true, nil
		}

		// Handle Time type
		if value, ok, err := Singleton[Time](target); err == nil && ok {
			if len(parameters) == 0 {
				precision = 2
			}
			if precision > 2 {
				return nil, true, nil
			}
			result := value
			switch precision {
			case 1: // Hour
				result.Precision = TimePrecisionHour
				result.Value = time.Date(0, 1, 1, result.Value.Hour(), 59, 59, int(time.Millisecond*999), result.Value.Location())
			case 2: // Minute
				result.Precision = TimePrecisionMinute
				result.Value = time.Date(0, 1, 1, result.Value.Hour(), result.Value.Minute(), 59, int(time.Millisecond*999), result.Value.Location())
			default: // Full precision
				result.Precision = TimePrecisionFull
			}
			return Collection{result}, true, nil
		}

		return nil, false, fmt.Errorf("expected Decimal, Date, DateTime, or Time but got %T", target[0])
	},
	"precision": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}

		// Handle Decimal type
		if value, ok, err := Singleton[Decimal](target); err == nil && ok {
			// For Decimal, return the number of digits after the decimal point
			str := value.Value.Text('f')
			if dot := strings.Index(str, "."); dot != -1 {
				// Count trailing zeros after decimal point
				fracPart := str[dot+1:]
				precision := len(fracPart)
				// Remove trailing zeros
				for precision > 0 && fracPart[precision-1] == '0' {
					precision--
				}
				return Collection{Integer(precision)}, true, nil
			}
			return Collection{Integer(0)}, true, nil
		}

		// Handle Date type
		if value, ok, err := Singleton[Date](target); err == nil && ok {
			// For Date, return the number of digits in the date
			switch value.Precision {
			case "year":
				return Collection{Integer(4)}, true, nil
			case "month":
				return Collection{Integer(6)}, true, nil
			case "day":
				return Collection{Integer(8)}, true, nil
			default:
				return nil, false, fmt.Errorf("invalid date precision")
			}
		}

		// Handle DateTime type
		if value, ok, err := Singleton[DateTime](target); err == nil && ok {
			// For DateTime, return the number of digits in the datetime
			switch value.Precision {
			case "year":
				return Collection{Integer(4)}, true, nil
			case "month":
				return Collection{Integer(6)}, true, nil
			case "day":
				return Collection{Integer(8)}, true, nil
			case "hour":
				return Collection{Integer(10)}, true, nil
			case "minute":
				return Collection{Integer(12)}, true, nil
			case "second":
				return Collection{Integer(14)}, true, nil
			case "millisecond":
				return Collection{Integer(17)}, true, nil
			default:
				return nil, false, fmt.Errorf("invalid datetime precision")
			}
		}

		// Handle Time type
		if value, ok, err := Singleton[Time](target); err == nil && ok {
			// For Time, return the number of digits in the time
			switch value.Precision {
			case "hour":
				return Collection{Integer(2)}, true, nil
			case "minute":
				return Collection{Integer(4)}, true, nil
			case "second":
				return Collection{Integer(6)}, true, nil
			case "millisecond":
				return Collection{Integer(9)}, true, nil
			default:
				return nil, false, fmt.Errorf("invalid time precision")
			}
		}

		return nil, false, fmt.Errorf("expected Decimal, Date, DateTime, or Time but got %T", target[0])
	},
	"defineVariable": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 && len(parameters) != 2 {
			return nil, false, fmt.Errorf("expected one or two parameters (name [, value])")
		}

		nameCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert name to string
		name, ok, err := Singleton[String](nameCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected string name parameter")
		}

		// Protect system variables from being overwritten
		if _, isSystem := systemVariables[string(name)]; isSystem {
			return nil, false, fmt.Errorf("cannot redefine system variable '%s'", name)
		}

		// Check if variable already defined in current scope
		if frame, ok := envStackFrame(ctx); ok {
			if _, exists := frame[string(name)]; exists {
				return nil, false, fmt.Errorf("variable %%%s already defined", name)
			}
		}

		// Determine the value to store
		// Variables in FHIRPath store the entire evaluated result (which can be a collection)
		var value Collection
		if len(parameters) == 2 {
			// Evaluate for each item in the input collection and aggregate results
			for _, item := range target {
				itemCollection, _, err := evaluate(ctx, item, parameters[1])
				if err != nil {
					return nil, false, err
				}
				value = append(value, itemCollection...)
			}
		} else {
			// Use the input collection as the value
			value = target
		}

		// Store the collection as the variable value in the parent context
		ctx = WithEnv(ctx, string(name), value)

		// Return the input collection (does not change input)
		return target, inputOrdered, nil
	},

	// Math functions
	"abs": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		i, ok, err := Singleton[Integer](target)
		if err == nil && ok {
			if i < 0 {
				return Collection{-i}, true, nil
			}
			return Collection{i}, true, nil
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Create a new Decimal with the absolute value
			var absValue apd.Decimal
			absValue.Abs(d.Value)
			return Collection{Decimal{Value: &absValue}}, true, nil
		}

		q, ok, err := Singleton[Quantity](target)
		if err == nil && ok {
			// Create a new Quantity with the absolute value of the value
			var absValue apd.Decimal
			absValue.Abs(q.Value.Value)
			return Collection{Quantity{Value: Decimal{Value: &absValue}, Unit: q.Unit}}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer, Decimal, or Quantity but got %T", target[0])
	},
	"ceiling": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		i, ok, err := Singleton[Integer](target)
		if err == nil && ok {
			// Integer is already a whole number, so ceiling is the same
			return Collection{i}, true, nil
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Get the integer part
			var intPart apd.Decimal
			_, err = apdContext(ctx).Ceil(&intPart, d.Value)
			if err != nil {
				return nil, false, err
			}

			intVal, err := intPart.Int64()
			if err != nil {
				return nil, false, err
			}

			return Collection{Integer(intVal)}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"floor": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		i, ok, err := Singleton[Integer](target)
		if err == nil && ok {
			// Integer is already a whole number, so floor is the same
			return Collection{i}, true, nil
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Get the integer part
			var intPart apd.Decimal
			_, err = apdContext(ctx).Floor(&intPart, d.Value)
			if err != nil {
				return nil, false, err
			}

			// Convert to Integer
			intVal, err := intPart.Int64()
			if err != nil {
				return nil, false, err
			}

			return Collection{Integer(intVal)}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"truncate": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		i, ok, err := Singleton[Integer](target)
		if err == nil && ok {
			// Integer is already a whole number, so truncate is the same
			return Collection{i}, true, nil
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Get the integer part
			var intPart apd.Decimal

			// Use Floor for positive numbers and Ceil for negative numbers
			if d.Value.Negative {
				_, err = apdContext(ctx).Ceil(&intPart, d.Value)
			} else {
				_, err = apdContext(ctx).Floor(&intPart, d.Value)
			}

			if err != nil {
				return nil, false, err
			}

			// Convert to Integer
			intVal, err := intPart.Int64()
			if err != nil {
				return nil, false, err
			}

			return Collection{Integer(intVal)}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"round": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) > 1 {
			return nil, false, fmt.Errorf("expected at most one precision parameter")
		}

		if len(target) == 0 {
			return nil, true, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}

		decimalPlaces := int64(0)
		if len(parameters) == 1 {
			c, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			decimalPlacesValue, ok, err := Singleton[Integer](c)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected integer precision parameter")
			}

			if decimalPlacesValue < 0 {
				return nil, false, fmt.Errorf("precision must be >= 0")
			}

			decimalPlaces = int64(decimalPlacesValue)
		}

		var dec *apd.Decimal
		// Convert Integer to Decimal if needed
		if i, ok, _ := Singleton[Integer](target); ok {
			dec = apd.New(int64(i), 0)
		} else if d, ok, _ := Singleton[Decimal](target); ok {
			dec = d.Value
		} else {
			return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
		}

		// Set precision for rounding
		apdCtx := apdContext(ctx).WithPrecision(uint32(dec.NumDigits() + decimalPlaces))
		var rounded apd.Decimal
		// Use Quantize to round to the specified number of decimal places
		_, err = apdCtx.Quantize(&rounded, dec, int32(-decimalPlaces))
		if err != nil {
			return nil, false, err
		}

		return Collection{Decimal{Value: &rounded}}, true, nil
	},
	"exp": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Calculate e^x
			var result apd.Decimal
			_, err = apdContext(ctx).Exp(&result, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"ln": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Calculate ln(x)
			var result apd.Decimal
			_, err = apdContext(ctx).Ln(&result, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"log": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected one base parameter")
		}
		baseCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		baseDecimal, ok, err := Singleton[Decimal](baseCollection)
		if err != nil || !ok {
			// Try to convert Integer to Decimal
			baseInt, ok, err := Singleton[Integer](baseCollection)
			if err != nil || !ok {
				return nil, false, fmt.Errorf("expected Integer or Decimal base parameter but got %T", baseCollection[0])
			}
			baseDecimal = Decimal{Value: apd.New(int64(baseInt), 0)}
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Calculate log_base(x) = ln(x) / ln(base)
			var lnX, lnBase, result apd.Decimal
			_, err = apdContext(ctx).Ln(&lnX, d.Value)
			if err != nil {
				return nil, false, err
			}

			_, err = apdContext(ctx).Ln(&lnBase, baseDecimal.Value)
			if err != nil {
				return nil, false, err
			}

			_, err = apdContext(ctx).Quo(&result, &lnX, &lnBase)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"power": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected one exponent parameter")
		}

		exponentCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		exponentInt, ok, err := Singleton[Integer](exponentCollection)
		if err == nil && ok {
			// Handle Integer base
			i, ok, err := Singleton[Integer](target)
			if err == nil && ok {
				// For Integer base and Integer exponent, return Integer if possible
				result := int64(math.Pow(float64(i), float64(exponentInt)))

				// Check if the result can be represented as an Integer
				if math.Pow(float64(i), float64(exponentInt)) == float64(result) {
					return Collection{Integer(result)}, true, nil
				}

				// Otherwise, return as Decimal
				resultDecimal := apd.New(0, 0)
				_, err := resultDecimal.SetFloat64(math.Pow(float64(i), float64(exponentInt)))
				if err != nil {
					return nil, false, err
				}
				return Collection{Decimal{Value: resultDecimal}}, true, nil
			}
		}

		// Get the exponent as a Decimal
		exponentDecimal, ok, err := Singleton[Decimal](exponentCollection)
		if err != nil || !ok {
			// Try to convert Integer to Decimal
			exponentInt, ok, err := Singleton[Integer](exponentCollection)
			if err != nil || !ok {
				return nil, false, fmt.Errorf("expected Integer or Decimal exponent parameter but got %T", exponentCollection[0])
			}
			exponentDecimal = Decimal{Value: apd.New(int64(exponentInt), 0)}
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			_, err := exponentDecimal.Value.Int64()
			// For negative base, the result is empty
			if d.Value.Negative {
				return nil, true, nil
			}

			// Calculate x^y
			var result apd.Decimal
			_, err = apdContext(ctx).Pow(&result, d.Value, exponentDecimal.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
	"sqrt": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			if d.Value.Negative {
				return nil, true, nil
			}

			var result apd.Decimal
			_, err = apdContext(ctx).Sqrt(&result, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, true, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},

	// Type conversion functions
	"toBoolean": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to boolean
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to boolean: collection contains > 1 values")
		}

		b, ok, err := elementTo[Boolean](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{b}, true, nil
	},
	"convertsToBoolean": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to boolean
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to boolean: collection contains > 1 values")
		}

		_, ok, err := elementTo[Boolean](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toInteger": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to integer
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to integer: collection contains > 1 values")
		}

		i, ok, err := elementTo[Integer](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{i}, true, nil
	},
	"convertsToInteger": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to integer
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to integer: collection contains > 1 values")
		}

		_, ok, err := elementTo[Integer](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toLong": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to long: collection contains > 1 values")
		}

		l, ok, err := target[0].ToLong(true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{l}, true, nil
	},
	"convertsToLong": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to long: collection contains > 1 values")
		}

		_, ok, err := target[0].ToLong(true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toDate": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to date
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to date: collection contains > 1 values")
		}

		d, ok, err := elementTo[Date](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{d}, true, nil
	},
	"convertsToDate": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to date
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to date: collection contains > 1 values")
		}

		_, ok, err := elementTo[Date](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toDateTime": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to datetime
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to datetime: collection contains > 1 values")
		}

		dt, ok, err := elementTo[DateTime](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{dt}, true, nil
	},
	"convertsToDateTime": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to datetime
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to datetime: collection contains > 1 values")
		}

		_, ok, err := elementTo[DateTime](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toTime": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to time
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to time: collection contains > 1 values")
		}

		t, ok, err := elementTo[Time](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{t}, true, nil
	},
	"convertsToTime": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to time
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to time: collection contains > 1 values")
		}

		_, ok, err := elementTo[Time](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toDecimal": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to decimal
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to decimal: collection contains > 1 values")
		}

		d, ok, err := elementTo[Decimal](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{d}, true, nil
	},
	"convertsToDecimal": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to decimal
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to decimal: collection contains > 1 values")
		}

		_, ok, err := elementTo[Decimal](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toQuantity": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		// Check parameters: optional unit
		if len(parameters) > 1 {
			return nil, false, fmt.Errorf("expected at most one unit parameter")
		}

		// Convert to quantity
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to quantity: collection contains > 1 values")
		}

		q, ok, err := elementTo[Quantity](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		// If unit parameter is provided, check if the quantity can be converted to the given unit
		if len(parameters) == 1 {
			// Evaluate the unit parameter
			unitCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			unitStr, ok, err := Singleton[String](unitCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected string unit parameter")
			}

			q.Unit = unitStr
		}

		return Collection{q}, true, nil
	},
	"convertsToQuantity": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		// Check parameters: optional unit
		if len(parameters) > 1 {
			return nil, false, fmt.Errorf("expected at most one unit parameter")
		}

		// Check if convertible to quantity
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to quantity: collection contains > 1 values")
		}

		_, ok, err := elementTo[Quantity](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		// If unit parameter is provided, check if the quantity can be converted to the given unit
		if len(parameters) == 1 {
			// Evaluate the unit parameter
			unitCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			// Convert to string
			if len(unitCollection) == 0 {
				return nil, false, fmt.Errorf("expected string unit parameter")
			} else if len(unitCollection) > 1 {
				return nil, false, fmt.Errorf("expected single string unit parameter")
			}

			_, ok, err := Singleton[String](unitCollection)
			if err != nil {
				return nil, false, err
			}
			if !ok {
				return nil, false, fmt.Errorf("expected string unit parameter")
			}
		}

		return Collection{Boolean(true)}, true, nil
	},
	"toString": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Convert to string
		if len(target) == 0 {
			return nil, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to string: collection contains > 1 values")
		}

		s, ok, err := elementTo[String](target[0], true)
		if err != nil || !ok {
			return nil, true, nil
		}

		return Collection{s}, true, nil
	},
	"convertsToString": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Check if convertible to string
		if len(target) == 0 {
			return Collection{Boolean(false)}, true, nil
		} else if len(target) > 1 {
			return nil, false, fmt.Errorf("cannot convert to string: collection contains > 1 values")
		}

		_, ok, err := elementTo[String](target[0], true)
		if err != nil || !ok {
			return Collection{Boolean(false)}, true, nil
		}

		return Collection{Boolean(true)}, true, nil
	},

	// Tree navigation functions
	"children": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		for _, elem := range target {
			// Get all immediate child nodes
			children := elem.Children()
			result = append(result, children...)
		}

		return result, false, nil
	},
	"descendants": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		// descendants() is a shorthand for repeat(children())
		// Manually implement the logic of repeat(children())

		var current = target
		var newItems Collection

		// Keep repeating the children() function until no new items are found
		for {
			newItems = nil

			// Get all children of the current elements
			for _, elem := range current {
				children := elem.Children()

				// Check for new items
				for _, child := range children {
					isNew := true
					for _, seen := range result {
						eq, ok := seen.Equal(child)
						if ok && eq {
							isNew = false
							break
						}
					}
					if isNew {
						newItems = append(newItems, child)
					}
				}
			}

			// If no new items were found, we're done
			if len(newItems) == 0 {
				break
			}

			// Add new items to the result and set them as the current items for the next iteration
			result = append(result, newItems...)
			current = newItems
		}

		return result, false, nil
	},

	// Utility functions
	"trace": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) == 0 || len(parameters) > 2 {
			return nil, false, fmt.Errorf("expected one or two parameters")
		}

		logger, err := tracer(ctx)
		if err != nil {
			return nil, false, err
		}

		nameParam, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}
		if len(nameParam) != 1 {
			return nil, false, fmt.Errorf("expected single name parameter")
		}
		nameStr, ok, err := Singleton[String](nameParam)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("name parameter cannot be null")
		}

		var logCollection Collection
		if len(parameters) == 2 {
			for i, elem := range target {
				projection, _, err := evaluate(ctx, elem, parameters[1], FunctionScope{index: i})
				if err != nil {
					return nil, false, err
				}
				logCollection = append(logCollection, projection...)
			}
		} else {
			logCollection = target
		}

		err = logger.Log(string(nameStr), logCollection)
		if err != nil {
			return nil, false, err
		}

		return target, inputOrdered, nil
	},
	"aggregate": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) == 0 || len(parameters) > 2 {
			return nil, false, fmt.Errorf("expected one or two parameters")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, true, nil
		}

		total := Collection{}
		totalOrdered := inputOrdered

		if len(parameters) == 2 {
			// If init value is provided, evaluate it
			var err error
			total, totalOrdered, err = evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}
		}

		// Iterate over the target collection
		for i, elem := range target {
			var ordered bool
			total, ordered, err = evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: total})
			if err != nil {
				return nil, false, err
			}
			if !ordered {
				resultOrdered = false
			}
		}

		return total, totalOrdered, nil
	},
	"now": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		now := now()
		dt := DateTime{Value: now, Precision: DateTimePrecisionFull}

		return Collection{dt}, inputOrdered, nil
	},
	"timeOfDay": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		// Get the current time
		now := now()
		t := Time{Value: now, Precision: TimePrecisionFull}

		return Collection{t}, inputOrdered, nil
	},
	"today": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameters")
		}

		now := now()
		d := Date{Value: now, Precision: DatePrecisionFull}

		return Collection{d}, inputOrdered, nil
	},
	"iif": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		// Check parameters: criterion, true-result, and optional otherwise-result
		if len(parameters) < 2 || len(parameters) > 3 {
			return nil, false, fmt.Errorf("expected 2 or 3 parameters (criterion, true-result, [otherwise-result])")
		}

		// FHIRPath 3.0: iif() requires 0 or 1 items in the input collection
		if len(target) > 1 {
			return nil, false, fmt.Errorf("iif() requires an input collection with 0 or 1 items, got %d items", len(target))
		}

		// Determine the evaluation target for $this context
		// If target has one item, use it; otherwise use nil
		var evalTarget Element
		var fnScope []FunctionScope
		if len(target) == 1 {
			evalTarget = target[0]
			// Preserve the parent function scope's index if it exists
			parentScope, err := getFunctionScope(ctx)
			if err == nil {
				// Use parent scope's index
				fnScope = []FunctionScope{{index: parentScope.index, total: target}}
			} else {
				// No parent scope, set index to 0
				fnScope = []FunctionScope{{index: 0, total: target}}
			}
		}

		// Evaluate the criterion expression with $this context
		criterion, _, err := evaluate(ctx, evalTarget, parameters[0], fnScope...)
		if err != nil {
			return nil, false, err
		}

		// FHIRPath 3.0: criterion must be a boolean value, not just convertible to boolean
		// Check if the criterion is actually a Boolean type
		if len(criterion) > 0 {
			if _, isBool := criterion[0].(Boolean); !isBool {
				return nil, false, fmt.Errorf("iif() criterion must evaluate to a boolean value, got %T", criterion[0])
			}
		}

		// Convert criterion to boolean
		criterionBool, ok, err := Singleton[Boolean](criterion)
		if err != nil {
			return nil, false, err
		}

		// Short-circuit evaluation: only evaluate the taken branch
		// If criterion is true, return the value of the true-result argument
		if ok && bool(criterionBool) {
			trueResult, ordered, err := evaluate(ctx, evalTarget, parameters[1], fnScope...)
			if err != nil {
				return nil, false, err
			}
			return trueResult, ordered, nil
		}

		// If criterion is false or an empty collection, return otherwise-result
		// If otherwise-result is not given, return an empty collection
		if len(parameters) == 3 {
			otherwiseResult, ordered, err := evaluate(ctx, evalTarget, parameters[2], fnScope...)
			if err != nil {
				return nil, false, err
			}
			return otherwiseResult, ordered, nil
		}

		// No otherwise-result, return empty collection
		return nil, true, nil
	},
	"yearOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date or DateTime, got %d items", len(target))
		}

		dt, ok, err := target[0].ToDateTime(false)
		if err != nil || !ok {
			d, ok, err := target[0].ToDate(false)
			if err != nil || !ok {
				return nil, false, fmt.Errorf("expected Date or DateTime, got %T", target[0])
			}
			dt = DateTime{Value: d.Value, Precision: DateTimePrecision(d.Precision)}
		}

		if dt.Precision == DateTimePrecisionYear {
			return Collection{Integer(dt.Value.Year())}, inputOrdered, nil
		}
		return nil, false, nil
	},
	"monthOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date or DateTime, got %d items", len(target))
		}

		dt, ok, err := target[0].ToDateTime(false)
		if err != nil || !ok {
			d, ok, err := target[0].ToDate(false)
			if err != nil || !ok {
				return nil, false, fmt.Errorf("expected Date or DateTime, got %T", target[0])
			}
			dt = DateTime{Value: d.Value, Precision: DateTimePrecision(d.Precision)}
		}

		if dt.Precision == DateTimePrecisionYear {
			return nil, inputOrdered, nil
		}
		return Collection{Integer(dt.Value.Month())}, inputOrdered, nil
	},
	"dayOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date or DateTime, got %d items", len(target))
		}

		dt, ok, err := target[0].ToDateTime(false)
		if err != nil || !ok {
			d, ok, err := target[0].ToDate(false)
			if err != nil || !ok {
				return nil, false, fmt.Errorf("expected Date or DateTime, got %T", target[0])
			}
			dt = DateTime{Value: d.Value, Precision: DateTimePrecision(d.Precision)}
		}

		if dt.Precision == DateTimePrecisionYear || dt.Precision == DateTimePrecisionMonth {
			return nil, inputOrdered, nil
		}
		return Collection{Integer(dt.Value.Day())}, inputOrdered, nil
	},
	"hourOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date, DateTime or Time, got %d items", len(target))
		}

		var t time.Time
		var precision string

		dt, ok, err := target[0].ToDateTime(false)
		if err == nil && ok {
			t = dt.Value
			precision = string(dt.Precision)
		} else {
			d, ok, err := target[0].ToDate(false)
			if err == nil && ok {
				t = d.Value
				precision = string(d.Precision)
			} else {
				time, ok, err := target[0].ToTime(false)
				if err != nil || !ok {
					return nil, false, fmt.Errorf("expected Date, DateTime or Time, got %T", target[0])
				}
				t = time.Value
				precision = string(time.Precision)
			}
		}

		if precision == DateTimePrecisionYear || precision == DateTimePrecisionMonth || precision == DateTimePrecisionDay {
			return nil, inputOrdered, nil
		}
		return Collection{Integer(t.Hour())}, inputOrdered, nil
	},
	"minuteOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date, DateTime or Time, got %d items", len(target))
		}

		var t time.Time
		var precision string

		dt, ok, err := target[0].ToDateTime(false)
		if err == nil && ok {
			t = dt.Value
			precision = string(dt.Precision)
		} else {
			d, ok, err := target[0].ToDate(false)
			if err == nil && ok {
				t = d.Value
				precision = string(d.Precision)
			} else {
				time, ok, err := target[0].ToTime(false)
				if err != nil || !ok {
					return nil, false, fmt.Errorf("expected Date, DateTime or Time, got %T", target[0])
				}
				t = time.Value
				precision = string(time.Precision)
			}
		}

		if precision == DateTimePrecisionYear || precision == DateTimePrecisionMonth || precision == DateTimePrecisionDay || precision == DateTimePrecisionHour {
			return nil, inputOrdered, nil
		}
		return Collection{Integer(t.Minute())}, inputOrdered, nil
	},
	"secondOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date, DateTime or Time, got %d items", len(target))
		}

		var t time.Time
		var precision string

		dt, ok, err := target[0].ToDateTime(false)
		if err == nil && ok {
			t = dt.Value
			precision = string(dt.Precision)
		} else {
			d, ok, err := target[0].ToDate(false)
			if err == nil && ok {
				t = d.Value
				precision = string(d.Precision)
			} else {
				time, ok, err := target[0].ToTime(false)
				if err != nil || !ok {
					return nil, false, fmt.Errorf("expected Date, DateTime or Time, got %T", target[0])
				}
				t = time.Value
				precision = string(time.Precision)
			}
		}

		if precision == DateTimePrecisionYear || precision == DateTimePrecisionMonth || precision == DateTimePrecisionDay || precision == DateTimePrecisionHour || precision == DateTimePrecisionMinute {
			return nil, inputOrdered, nil
		}
		return Collection{Integer(t.Second())}, inputOrdered, nil
	},
	"millisecondOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date, DateTime or Time, got %d items", len(target))
		}

		var t time.Time
		var precision string

		dt, ok, err := target[0].ToDateTime(false)
		if err == nil && ok {
			t = dt.Value
			precision = string(dt.Precision)
		} else {
			d, ok, err := target[0].ToDate(false)
			if err == nil && ok {
				t = d.Value
				precision = string(d.Precision)
			} else {
				time, ok, err := target[0].ToTime(false)
				if err != nil || !ok {
					return nil, false, fmt.Errorf("expected Date, DateTime or Time, got %T", target[0])
				}
				t = time.Value
				precision = string(time.Precision)
			}
		}

		if precision == DateTimePrecisionYear || precision == DateTimePrecisionMonth || precision == DateTimePrecisionDay || precision == DateTimePrecisionHour || precision == DateTimePrecisionMinute {
			return nil, inputOrdered, nil
		}
		return Collection{Integer(t.Nanosecond() / 1000000)}, inputOrdered, nil
	},
	"timezoneOffsetOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single DateTime, got %d items", len(target))
		}

		dt, ok, err := target[0].ToDateTime(false)
		if err != nil || !ok {
			return nil, false, fmt.Errorf("expected DateTime, got %T", target[0])
		}

		_, offset := dt.Value.Zone()
		hours := float64(offset) / 3600.0
		dec := apd.New(0, 0)
		dec.SetFloat64(hours)
		return Collection{Decimal{Value: dec}}, inputOrdered, nil
	},
	"dateOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single Date or DateTime, got %d items", len(target))
		}

		dt, ok, err := target[0].ToDateTime(false)
		if err != nil || !ok {
			d, ok, err := target[0].ToDate(false)
			if err != nil || !ok {
				return nil, false, fmt.Errorf("expected Date or DateTime, got %T", target[0])
			}
			return Collection{d}, inputOrdered, nil
		}

		var precision DatePrecision
		switch dt.Precision {
		case DateTimePrecisionYear:
			precision = DatePrecisionYear
		case DateTimePrecisionMonth:
			precision = DatePrecisionMonth
		default:
			precision = DatePrecisionFull
		}

		return Collection{Date{Value: dt.Value, Precision: precision}}, inputOrdered, nil
	},
	"timeOf": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) == 0 {
			return nil, inputOrdered, nil
		}
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single DateTime, got %d items", len(target))
		}

		dt, ok, err := target[0].ToDateTime(false)
		if err != nil || !ok {
			return nil, false, fmt.Errorf("expected DateTime, got %T", target[0])
		}

		if dt.Precision == DateTimePrecisionYear || dt.Precision == DateTimePrecisionMonth || dt.Precision == DateTimePrecisionDay {
			return nil, inputOrdered, nil
		}

		var precision TimePrecision
		switch dt.Precision {
		case DateTimePrecisionHour:
			precision = TimePrecisionHour
		case DateTimePrecisionMinute:
			precision = TimePrecisionMinute
		default:
			precision = TimePrecisionFull
		}

		return Collection{Time{Value: dt.Value, Precision: precision}}, inputOrdered, nil
	},
	"extension": FHIRFunctions["extension"],
}

var FHIRFunctions = Functions{
	"extension": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected a single extension parameter")
		}

		extCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		extUrl, ok, err := Singleton[String](extCollection)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			return nil, false, fmt.Errorf("expected extension parameter string")
		}

		var foundExtensions Collection
		for _, v := range target {
			for _, e := range v.Children("extension") {
				url, ok, err := Singleton[String](e.Children("url"))
				if err == nil && ok && url == extUrl {
					foundExtensions = append(foundExtensions, e)
					break
				}
			}
		}
		return foundExtensions, inputOrdered, nil
	},
}

func compareElementsForSort(a, b Element) (int, error) {
	cmp, ok, err := Collection{a}.Cmp(Collection{b})
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("elements %T and %T are not comparable", a, b)
	}
	return cmp, nil
}
