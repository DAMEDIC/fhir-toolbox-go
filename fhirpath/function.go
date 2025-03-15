package fhirpath

import (
	"context"
	"fmt"
	"github.com/cockroachdb/apd/v3"
	"maps"
	"math"
	"regexp"
	"strings"
	"time"
)

// TraceLogger defines the interface for logging trace messages
type TraceLogger interface {
	// Log logs a trace message with the given name and collection
	Log(name string, collection Collection) error
}

// StdoutTraceLogger writes traces to io.Stdout.
type StdoutTraceLogger struct{}

func (w StdoutTraceLogger) Log(name string, collection Collection) error {
	_, err := fmt.Printf("%s: %v\n", name, collection)
	return err
}

type traceLoggerKey struct{}

// WithTraceLogger installs the given trace logger into the context.
//
// By default, traces are logged to stdout.
// To redirect trace logs to a custom output, use:
//
//	ctx = fhirpath.WithTraceLogger(ctx, MyCustomTraceLogger(true, file))
func WithTraceLogger(ctx context.Context, logger TraceLogger) context.Context {
	return context.WithValue(ctx, traceLoggerKey{}, logger)
}

// GetTraceLogger gets the trace logger from the context
// If no trace logger is found, a NoOpTraceLogger is returned
func traceLogger(ctx context.Context) (TraceLogger, error) {
	logger, ok := ctx.Value(traceLoggerKey{}).(TraceLogger)
	if !ok {
		return StdoutTraceLogger{}, nil
	}
	if logger == nil {
		return StdoutTraceLogger{}, fmt.Errorf("no trace logger provided")
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
) (result Collection, resultresultOrdered bool, err error)

type EvaluateFunc = func(
	ctx context.Context,
	target Element,
	expr Expression,
	fnScope ...FunctionScope,
) (result Collection, resultOrdered bool, err error)

type functionCtxKey struct{}

type FunctionScope struct {
	index int
	total int
}

type functionScope struct {
	this  Element
	index int
	total int
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

var defaultFunctions = Functions{
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
		return Collection{r}, inputOrdered, nil
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
		return c, inputOrdered, nil
	},
	"not": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(target) != 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}
		if len(parameters) != 0 {
			return nil, false, fmt.Errorf("expected no parameter")
		}
		b, err := target[0].ToBoolean(false)
		if err != nil {
			return nil, false, err
		}
		if b == nil {
			return nil, false, nil
		}
		return Collection{!*b}, inputOrdered, nil
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
		return Collection{Boolean(len(target) == 0)}, inputOrdered, nil
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
			return Collection{Boolean(len(target) > 0)}, inputOrdered, nil
		}

		// With criteria, equivalent to where(criteria).exists()
		for i, elem := range target {
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
			if err != nil {
				return nil, false, err
			}

			if len(criteria) == 1 {
				b, err := criteria[0].ToBoolean(false)
				if err != nil {
					return nil, false, err
				}
				if b != nil && *b {
					// Found at least one element that matches the criteria
					return Collection{Boolean(true)}, inputOrdered, nil
				}
			}
		}

		return Collection{Boolean(false)}, inputOrdered, nil
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
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		for i, elem := range target {
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
			if err != nil {
				return nil, false, err
			}

			if len(criteria) == 1 {
				b, err := criteria[0].ToBoolean(false)
				if err != nil {
					return nil, false, err
				}
				if b == nil || !*b {
					// Found at least one element that doesn't match the criteria
					return Collection{Boolean(false)}, inputOrdered, nil
				}
			} else {
				return Collection{Boolean(false)}, inputOrdered, nil
			}
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		for _, elem := range target {
			b, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if b == nil || !*b {
				return Collection{Boolean(false)}, inputOrdered, nil
			}
		}
		return Collection{Boolean(true)}, inputOrdered, nil
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
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		for _, elem := range target {
			b, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if b != nil && *b {
				return Collection{Boolean(true)}, inputOrdered, nil
			}
		}
		return Collection{Boolean(false)}, inputOrdered, nil
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
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		for _, elem := range target {
			b, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if b == nil || *b {
				return Collection{Boolean(false)}, inputOrdered, nil
			}
		}
		return Collection{Boolean(true)}, inputOrdered, nil
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
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		for _, elem := range target {
			b, err := elem.ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
			if b != nil && !*b {
				return Collection{Boolean(true)}, inputOrdered, nil
			}
		}
		return Collection{Boolean(false)}, inputOrdered, nil
	},
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
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the other collection is empty, the result is false
		if len(other) == 0 {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		for _, elem := range target {
			found := false
			for _, otherElem := range other {
				if elem.Equal(otherElem) {
					found = true
					break
				}
			}
			if !found {
				return Collection{Boolean(false)}, inputOrdered, nil
			}
		}
		return Collection{Boolean(true)}, inputOrdered, nil
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
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		// If the input collection is empty, the result is false
		if len(target) == 0 {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		for _, otherElem := range other {
			found := false
			for _, elem := range target {
				if otherElem.Equal(elem) {
					found = true
					break
				}
			}
			if !found {
				return Collection{Boolean(false)}, inputOrdered, nil
			}
		}
		return Collection{Boolean(true)}, inputOrdered, nil
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
		return Collection{Integer(len(target))}, inputOrdered, nil
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
			return nil, false, nil
		}

		for _, elem := range target {
			found := false
			for _, resultElem := range result {
				if elem.Equal(resultElem) {
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
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		// Check if all elements are distinct
		for i := 0; i < len(target); i++ {
			for j := i + 1; j < len(target); j++ {
				if target[i].Equal(target[j]) {
					return Collection{Boolean(false)}, inputOrdered, nil
				}
			}
		}
		return Collection{Boolean(true)}, inputOrdered, nil
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
			return nil, false, nil
		}

		for i, elem := range target {
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
			if err != nil {
				return nil, false, err
			}

			if len(criteria) == 1 {
				b, err := criteria[0].ToBoolean(false)
				if err != nil {
					return nil, false, err
				}
				if b != nil && *b {
					// Element matches the criteria, add it to the result
					result = append(result, elem)
				}
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
			return nil, false, nil
		}

		resultOrdered = inputOrdered
		for i, elem := range target {
			projection, ordered, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
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
			return nil, false, nil
		}

		var current = target
		var newItems Collection

		for _, elem := range current {
			result = append(result, elem)
		}

		// Keep repeating the projection until no new items are found
		for {
			newItems = nil
			for i, elem := range current {
				projection, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(current)})
				if err != nil {
					return nil, false, err
				}

				// Check for new items
				for _, item := range projection {
					for _, seen := range result {
						if !seen.Equal(item) {
							newItems = append(newItems, item)
						}
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
			return nil, false, nil
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
			return nil, false, nil
		}

		// If there are multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Return the single item
		return Collection{target[0]}, inputOrdered, nil
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
			return nil, false, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Return the first item
		return Collection{target[0]}, inputOrdered, nil
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
			return nil, false, nil
		}

		if !inputOrdered {
			return nil, false, fmt.Errorf("expected ordered input")
		}

		// Return the last item
		return Collection{target[len(target)-1]}, inputOrdered, nil
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
			return nil, false, nil
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
			return nil, false, nil
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
		num, err := Singleton[Integer](numCollection)
		if err != nil {
			return nil, false, err
		}
		if num == nil {
			return nil, false, fmt.Errorf("expected integer parameter")
		}

		// If num is less than or equal to zero, return the input collection
		if *num <= 0 {
			return target, inputOrdered, nil
		}

		// If num is greater than or equal to the length of the collection, return empty
		if int(*num) >= len(target) {
			return nil, false, nil
		}

		// Return all but the first num items
		return target[int(*num):], inputOrdered, nil
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
			return nil, false, nil
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
		num, err := Singleton[Integer](numCollection)
		if err != nil {
			return nil, false, err
		}
		if num == nil {
			return nil, false, fmt.Errorf("expected integer parameter")
		}

		// If num is less than or equal to zero, return empty
		if *num <= 0 {
			return nil, false, nil
		}

		// If num is greater than the length of the collection, return the whole collection
		if int(*num) >= len(target) {
			return target, inputOrdered, nil
		}

		// Return the first num items
		return target[:int(*num)], inputOrdered, nil
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
			return nil, false, nil
		}

		// Evaluate the other collection parameter
		other, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the other collection is empty, the result is empty
		if len(other) == 0 {
			return nil, false, nil
		}

		// Find the intersection of the two collections
		for _, elem := range target {
			// Check if the element is in the other collection
			for _, otherElem := range other {
				if elem.Equal(otherElem) {
					// Check if the element is already in the result (eliminate duplicates)
					found := false
					for _, resultElem := range result {
						if elem.Equal(resultElem) {
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
			return nil, false, nil
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
				if elem.Equal(otherElem) {
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

		// Evaluate the criterion expression
		criterion, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// Convert criterion to boolean
		var criterionBool *Boolean
		if len(criterion) == 1 {
			criterionBool, err = criterion[0].ToBoolean(false)
			if err != nil {
				return nil, false, err
			}
		}

		// If criterion is true, return the value of the true-result argument
		if criterionBool != nil && bool(*criterionBool) {
			trueResult, ordered, err := evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}
			return trueResult, ordered, nil
		}

		// If criterion is false or an empty collection, return otherwise-result
		// If otherwise-result is not given, return an empty collection
		if len(parameters) == 3 {
			otherwiseResult, ordered, err := evaluate(ctx, nil, parameters[2])
			if err != nil {
				return nil, false, err
			}
			return otherwiseResult, ordered, nil
		}

		// No otherwise-result, return empty collection
		return nil, false, nil
	},
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to boolean
		b, err := target[0].ToBoolean(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if b == nil {
			return nil, false, nil
		}

		return Collection{*b}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to boolean
		_, err = target[0].ToBoolean(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to integer
		i, err := target[0].ToInteger(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if i == nil {
			return nil, false, nil
		}

		return Collection{*i}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to integer
		_, err = target[0].ToInteger(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to date
		d, err := target[0].ToDate(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if d == nil {
			return nil, false, nil
		}

		return Collection{*d}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to date
		_, err = target[0].ToDate(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to datetime
		dt, err := target[0].ToDateTime(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if dt == nil {
			return nil, false, nil
		}

		return Collection{*dt}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to datetime
		_, err = target[0].ToDateTime(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to time
		t, err := target[0].ToTime(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if t == nil {
			return nil, false, nil
		}

		return Collection{*t}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to time
		_, err = target[0].ToTime(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to decimal
		d, err := target[0].ToDecimal(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if d == nil {
			return nil, false, nil
		}

		return Collection{*d}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to decimal
		_, err = target[0].ToDecimal(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to quantity
		q, err := target[0].ToQuantity(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if q == nil {
			return nil, false, nil
		}

		// If unit parameter is provided, check if the quantity can be converted to the given unit
		if len(parameters) == 1 {
			// Evaluate the unit parameter
			unitCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			// Convert to string
			unitStr, err := Singleton[String](unitCollection)
			if err != nil {
				return nil, false, err
			}
			if unitStr == nil {
				return nil, false, fmt.Errorf("expected string unit parameter")
			}

			// Check if the quantity can be converted to the given unit
			// Note: The specification says implementations are not required to support a complete UCUM implementation,
			// and may return empty ({ }) when the unit argument is used and it is different than the input quantity unit.
			// For now, we'll just check if the units are the same.
			if q.Unit != *unitStr {
				return nil, false, nil // Return empty collection if units don't match
			}
		}

		return Collection{*q}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to quantity
		q, err := target[0].ToQuantity(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}
		if q == nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		// If unit parameter is provided, check if the quantity can be converted to the given unit
		if len(parameters) == 1 {
			// Evaluate the unit parameter
			unitCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			// Convert to string
			unitStr, err := Singleton[String](unitCollection)
			if err != nil {
				return nil, false, err
			}
			if unitStr == nil {
				return nil, false, fmt.Errorf("expected string unit parameter")
			}

			// Check if the quantity can be converted to the given unit
			// Note: The specification says implementations are not required to support a complete UCUM implementation,
			// and may return false when the unit argument is used and it is different than the input quantity unit.
			// For now, we'll just check if the units are the same.
			if q.Unit != *unitStr {
				return Collection{Boolean(false)}, inputOrdered, nil
			}
		}

		return Collection{Boolean(true)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, nil // Return empty collection if conversion fails
		}
		if s == nil {
			return nil, false, nil
		}

		return Collection{*s}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Check if convertible to string
		_, err = target[0].ToString(true)
		if err != nil {
			return Collection{Boolean(false)}, inputOrdered, nil
		}

		return Collection{Boolean(true)}, inputOrdered, nil
	},
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the substring parameter
		substringCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the substring collection is empty, the result is empty
		if len(substringCollection) == 0 {
			return nil, false, nil
		}

		// Convert substring to string
		substring, err := Singleton[String](substringCollection)
		if err != nil {
			return nil, false, err
		}
		if substring == nil {
			return nil, false, fmt.Errorf("expected string substring parameter")
		}

		// If substring is an empty string (''), the function returns 0
		if *substring == "" {
			return Collection{Integer(0)}, inputOrdered, nil
		}

		// Return the index of the substring in the string
		index := strings.Index(string(*s), string(*substring))
		return Collection{Integer(index)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the start parameter
		startCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the start collection is empty, the result is empty
		if len(startCollection) == 0 {
			return nil, false, nil
		}

		// Convert start to integer
		start, err := Singleton[Integer](startCollection)
		if err != nil {
			return nil, false, err
		}
		if start == nil {
			return nil, false, fmt.Errorf("expected integer start parameter")
		}

		// If start is negative or greater than or equal to the length of the string, return empty
		if *start < 0 || int(*start) >= len(*s) {
			return nil, false, nil
		}

		// If length parameter is provided
		if len(parameters) == 2 {
			// Evaluate the length parameter
			lengthCollection, _, err := evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}

			// If the length collection is empty, treat it as if length had not been provided
			if len(lengthCollection) == 0 {
				result := String(string(*s)[int(*start):])
				return Collection{result}, inputOrdered, nil
			}

			// Convert length to integer
			length, err := Singleton[Integer](lengthCollection)
			if err != nil {
				return nil, false, err
			}
			if length == nil {
				return nil, false, fmt.Errorf("expected integer length parameter")
			}

			// If length is negative, treat it as 0
			if *length < 0 {
				*length = 0
			}

			// Calculate end index (start + length)
			end := int(*start) + int(*length)
			if end > len(*s) {
				end = len(*s)
			}

			result := String(string(*s)[int(*start):end])
			return Collection{result}, inputOrdered, nil
		}

		// If length parameter is not provided, return the rest of the string
		return Collection{String(string(*s)[int(*start):])}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the prefix parameter
		prefixCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the prefix collection is empty, the result is empty
		if len(prefixCollection) == 0 {
			return nil, false, nil
		}

		// Convert prefix to string
		prefix, err := Singleton[String](prefixCollection)
		if err != nil {
			return nil, false, err
		}
		if prefix == nil {
			return nil, false, fmt.Errorf("expected string prefix parameter")
		}

		// If prefix is an empty string (''), the result is true
		if *prefix == "" {
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		// Check if the string starts with the prefix
		startsWith := strings.HasPrefix(string(*s), string(*prefix))
		return Collection{Boolean(startsWith)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the suffix parameter
		suffixCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the suffix collection is empty, the result is empty
		if len(suffixCollection) == 0 {
			return nil, false, nil
		}

		// Convert suffix to string
		suffix, err := Singleton[String](suffixCollection)
		if err != nil {
			return nil, false, err
		}
		if suffix == nil {
			return nil, false, fmt.Errorf("expected string suffix parameter")
		}

		// If suffix is an empty string (''), the result is true
		if *suffix == "" {
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		// Check if the string ends with the suffix
		endsWith := strings.HasSuffix(string(*s), string(*suffix))
		return Collection{Boolean(endsWith)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the substring parameter
		substringCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the substring collection is empty, the result is empty
		if len(substringCollection) == 0 {
			return nil, false, nil
		}

		// Convert substring to string
		substring, err := Singleton[String](substringCollection)
		if err != nil {
			return nil, false, err
		}
		if substring == nil {
			return nil, false, fmt.Errorf("expected string substring parameter")
		}

		// If substring is an empty string (''), the result is true
		if *substring == "" {
			return Collection{Boolean(true)}, inputOrdered, nil
		}

		// Check if the string contains the substring
		contains := strings.Contains(string(*s), string(*substring))
		return Collection{Boolean(contains)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Convert the string to upper case
		return Collection{String(strings.ToUpper(string(*s)))}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Convert the string to lower case
		return Collection{String(strings.ToLower(string(*s)))}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the pattern parameter
		patternCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the pattern collection is empty, the result is empty
		if len(patternCollection) == 0 {
			return nil, false, nil
		}

		// Convert pattern to string
		pattern, err := Singleton[String](patternCollection)
		if err != nil {
			return nil, false, err
		}
		if pattern == nil {
			return nil, false, fmt.Errorf("expected string pattern parameter")
		}

		// Evaluate the substitution parameter
		substitutionCollection, _, err := evaluate(ctx, nil, parameters[1])
		if err != nil {
			return nil, false, err
		}

		// If the substitution collection is empty, the result is empty
		if len(substitutionCollection) == 0 {
			return nil, false, nil
		}

		// Convert substitution to string
		substitution, err := Singleton[String](substitutionCollection)
		if err != nil {
			return nil, false, err
		}
		if substitution == nil {
			return nil, false, fmt.Errorf("expected string substitution parameter")
		}

		// If pattern is an empty string (''), every character in the input string is surrounded by the substitution
		if *pattern == "" {
			var result strings.Builder
			result.WriteString(string(*substitution))
			for _, c := range string(*s) {
				result.WriteRune(c)
				result.WriteString(string(*substitution))
			}
			return Collection{String(result.String())}, inputOrdered, nil
		}

		// Replace all instances of pattern with substitution
		return Collection{String(strings.ReplaceAll(string(*s), string(*pattern), string(*substitution)))}, inputOrdered, nil
	},
	"matches": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 1 {
			return nil, false, fmt.Errorf("expected single regex parameter")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the regex parameter
		regexCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the regex collection is empty, the result is empty
		if len(regexCollection) == 0 {
			return nil, false, nil
		}

		// Convert regex to string
		regexStr, err := Singleton[String](regexCollection)
		if err != nil {
			return nil, false, err
		}
		if regexStr == nil {
			return nil, false, fmt.Errorf("expected string regex parameter")
		}

		// Compile the regular expression
		regex, err := regexp.Compile(string(*regexStr))
		if err != nil {
			return nil, false, fmt.Errorf("invalid regular expression: %v", err)
		}

		// Check if the string matches the regular expression
		return Collection{Boolean(regex.MatchString(string(*s)))}, inputOrdered, nil
	},
	"replaceMatches": func(
		ctx context.Context,
		root Element, target Collection,
		inputOrdered bool,
		parameters []Expression,
		evaluate EvaluateFunc,
	) (result Collection, resultOrdered bool, err error) {
		if len(parameters) != 2 {
			return nil, false, fmt.Errorf("expected two parameters (regex, substitution)")
		}

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Evaluate the regex parameter
		regexCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the regex collection is empty, the result is empty
		if len(regexCollection) == 0 {
			return nil, false, nil
		}

		// Convert regex to string
		regexStr, err := Singleton[String](regexCollection)
		if err != nil {
			return nil, false, err
		}
		if regexStr == nil {
			return nil, false, fmt.Errorf("expected string regex parameter")
		}

		// Evaluate the substitution parameter
		substitutionCollection, _, err := evaluate(ctx, nil, parameters[1])
		if err != nil {
			return nil, false, err
		}

		// If the substitution collection is empty, the result is empty
		if len(substitutionCollection) == 0 {
			return nil, false, nil
		}

		// Convert substitution to string
		substitution, err := Singleton[String](substitutionCollection)
		if err != nil {
			return nil, false, err
		}
		if substitution == nil {
			return nil, false, fmt.Errorf("expected string substitution parameter")
		}

		// Compile the regular expression
		regex, err := regexp.Compile(string(*regexStr))
		if err != nil {
			return nil, false, fmt.Errorf("invalid regular expression: %v", err)
		}

		// Replace all matches of the regular expression with the substitution
		return Collection{String(regex.ReplaceAllString(string(*s), string(*substitution)))}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Return the length of the string
		return Collection{Integer(len(*s))}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Convert input to string
		s, err := target[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if s == nil {
			return nil, false, nil
		}

		// Convert the string to a collection of characters
		chars := make(Collection, len(*s))
		for i, c := range string(*s) {
			result[i] = String(c)
		}
		return chars, true, nil
	},
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Handle Integer
		i, err := target[0].ToInteger(false)
		if err == nil && i != nil {
			if *i < 0 {
				return Collection{Integer(-*i)}, inputOrdered, nil
			}
			return Collection{*i}, inputOrdered, nil
		}

		// Handle Decimal
		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			// Create a new Decimal with the absolute value
			var absValue apd.Decimal
			absValue.Abs(d.Value)
			return Collection{Decimal{Value: &absValue}}, inputOrdered, nil
		}

		// Handle Quantity
		q, err := target[0].ToQuantity(false)
		if err == nil && q != nil {
			// Create a new Quantity with the absolute value of the value
			var absValue apd.Decimal
			absValue.Abs(q.Value.Value)
			return Collection{Quantity{Value: Decimal{Value: &absValue}, Unit: q.Unit}}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Handle Integer
		i, err := target[0].ToInteger(false)
		if err == nil && i != nil {
			// Integer is already a whole number, so ceiling is the same
			return Collection{*i}, inputOrdered, nil
		}

		// Handle Decimal
		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			// Get the integer part
			var intPart apd.Decimal
			_, err = apdContext(ctx).Ceil(&intPart, d.Value)
			if err != nil {
				return nil, false, err
			}

			// Convert to Integer
			intVal, err := intPart.Int64()
			if err != nil {
				return nil, false, err
			}

			return Collection{Integer(intVal)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Handle Integer
		i, err := target[0].ToInteger(false)
		if err == nil && i != nil {
			// Integer is already a whole number, so floor is the same
			return Collection{*i}, inputOrdered, nil
		}

		// Handle Decimal
		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
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

			return Collection{Integer(intVal)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Handle Integer
		i, err := target[0].ToInteger(false)
		if err == nil && i != nil {
			// Integer is already a whole number, so truncate is the same
			return Collection{*i}, inputOrdered, nil
		}

		// Handle Decimal
		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
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

			return Collection{Integer(intVal)}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Default precision is 0
		precision := 0

		// If precision parameter is provided, evaluate it
		if len(parameters) == 1 {
			precisionCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			// If the precision collection is empty, use default precision
			if len(precisionCollection) > 0 {
				precisionInt, err := Singleton[Integer](precisionCollection)
				if err != nil {
					return nil, false, err
				}
				if precisionInt == nil {
					return nil, false, fmt.Errorf("expected integer precision parameter")
				}

				// Precision must be >= 0
				if *precisionInt < 0 {
					return nil, false, fmt.Errorf("precision must be >= 0")
				}

				precision = int(*precisionInt)
			}
		}

		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			// Create a context with the specified precision
			apdCtx := apdContext(ctx).WithPrecision(uint32(precision + 1))

			// Round the decimal to the specified precision
			var rounded apd.Decimal
			_, err = apdCtx.Round(&rounded, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &rounded}}, inputOrdered, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			// Calculate e^x
			var result apd.Decimal
			_, err = apdContext(ctx).Exp(&result, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			// Calculate ln(x)
			var result apd.Decimal
			_, err = apdContext(ctx).Ln(&result, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Evaluate the base parameter
		baseCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the base collection is empty, the result is empty
		if len(baseCollection) == 0 {
			return nil, false, nil
		}

		// Get the base as a Decimal
		baseDecimal, err := baseCollection[0].ToDecimal(false)
		if err != nil || baseDecimal == nil {
			// Try to convert Integer to Decimal
			baseInt, err := baseCollection[0].ToInteger(false)
			if err != nil || baseInt == nil {
				return nil, false, fmt.Errorf("expected Integer or Decimal base parameter but got %T", baseCollection[0])
			}
			baseDecimal = &Decimal{Value: apd.New(int64(*baseInt), 0)}
		}

		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
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

			return Collection{Decimal{Value: &result}}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		// Evaluate the exponent parameter
		exponentCollection, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}

		// If the exponent collection is empty, the result is empty
		if len(exponentCollection) == 0 {
			return nil, false, nil
		}

		// Try to get the exponent as an Integer
		exponentInt, err := exponentCollection[0].ToInteger(false)
		if err == nil && exponentInt != nil {
			// Handle Integer base
			i, err := target[0].ToInteger(false)
			if err == nil && i != nil {
				// For Integer base and Integer exponent, return Integer if possible
				result := int64(math.Pow(float64(*i), float64(*exponentInt)))

				// Check if the result can be represented as an Integer
				if math.Pow(float64(*i), float64(*exponentInt)) == float64(result) {
					return Collection{Integer(result)}, inputOrdered, nil
				}

				// Otherwise, return as Decimal
				resultDecimal := apd.New(0, 0)
				_, err := resultDecimal.SetFloat64(math.Pow(float64(*i), float64(*exponentInt)))
				if err != nil {
					return nil, false, err
				}
				return Collection{Decimal{Value: resultDecimal}}, inputOrdered, nil
			}
		}

		// Get the exponent as a Decimal
		exponentDecimal, err := exponentCollection[0].ToDecimal(false)
		if err != nil || exponentDecimal == nil {
			// Try to convert Integer to Decimal
			exponentInt, err := exponentCollection[0].ToInteger(false)
			if err != nil || exponentInt == nil {
				return nil, false, fmt.Errorf("expected Integer or Decimal exponent parameter but got %T", exponentCollection[0])
			}
			exponentDecimal = &Decimal{Value: apd.New(int64(*exponentInt), 0)}
		}

		// Handle Decimal base
		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			_, err := exponentDecimal.Value.Int64()
			// For negative base, the result is empty
			if d.Value.Negative {
				return nil, false, nil
			}

			// Calculate x^y
			var result apd.Decimal
			_, err = apdContext(ctx).Pow(&result, d.Value, exponentDecimal.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, inputOrdered, nil
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

		// If the input collection is empty, the result is empty
		if len(target) == 0 {
			return nil, false, nil
		}

		// If the input collection contains multiple items, signal an error
		if len(target) > 1 {
			return nil, false, fmt.Errorf("expected single item but got %d items", len(target))
		}

		d, err := target[0].ToDecimal(false)
		if err == nil && d != nil {
			// For negative input, the result is empty
			if d.Value.Negative {
				return nil, false, nil
			}

			// Calculate sqrt(x)
			var result apd.Decimal
			_, err = apdContext(ctx).Sqrt(&result, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &result}}, inputOrdered, nil
		}

		return nil, false, fmt.Errorf("expected Integer or Decimal but got %T", target[0])
	},
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
			return nil, false, nil
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
			return nil, false, nil
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
						if seen.Equal(child) {
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

		// Get the trace logger from the context
		logger, err := traceLogger(ctx)
		if err != nil {
			return nil, false, err
		}

		// Get the name parameter
		nameParam, _, err := evaluate(ctx, nil, parameters[0])
		if err != nil {
			return nil, false, err
		}
		if len(nameParam) != 1 {
			return nil, false, fmt.Errorf("expected single name parameter")
		}
		nameStr, err := nameParam[0].ToString(true)
		if err != nil {
			return nil, false, err
		}
		if nameStr == nil {
			return nil, false, fmt.Errorf("name parameter cannot be null")
		}

		// Determine what to log
		var logCollection Collection
		if len(parameters) == 2 {
			// If projection is provided, evaluate it on the input
			for i, elem := range target {
				projection, _, err := evaluate(ctx, elem, parameters[1], FunctionScope{index: i, total: len(target)})
				if err != nil {
					return nil, false, err
				}
				logCollection = append(logCollection, projection...)
			}
		} else {
			// Otherwise, log the input collection
			logCollection = target
		}

		// Log the collection with the given name
		err = logger.Log(string(*nameStr), logCollection)
		if err != nil {
			return nil, false, err
		}

		// Return the input collection unchanged
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
			return nil, false, nil
		}

		if len(parameters) == 2 {
			// If init value is provided, evaluate it
			var err error
			result, _, err = evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}
		}

		resultOrdered = inputOrdered

		// Iterate over the target collection
		for i, elem := range target {
			// Create a new context with the current total value
			var totalCtx context.Context
			if len(result) == 0 {
				// If total is empty, don't add it to the context
				// The $total variable will be empty in the expression
				totalCtx = ctx
			} else {
				// Otherwise, add the total value to the context
				totalCtx = WithEnv(ctx, "total", result[0])
			}

			// Evaluate the aggregator expression with the current element and total
			result, ordered, err := evaluate(totalCtx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
			if err != nil {
				return nil, false, err
			}

			// Update the total variable with the result
			if len(result) == 0 {
				result = Collection{}
			} else if len(result) == 1 {
				result = Collection{result[0]}
			} else {
				return nil, false, fmt.Errorf("aggregator expression must return a single value")
			}
			if !ordered {
				resultOrdered = false
			}
		}

		return result, resultOrdered, nil
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

		// Get the current date and time
		now := time.Now()
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
		now := time.Now()
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

		// Get the current date
		now := time.Now()
		d := Date{Value: now, Precision: DatePrecisionFull}

		return Collection{d}, inputOrdered, nil
	},
}
