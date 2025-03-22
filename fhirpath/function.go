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
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
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
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
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
			criteria, _, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
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
			return nil, true, nil
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
						eq, ok := seen.Equal(item)
						if !ok || !eq {
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
		criterionBool, ok, err := Singleton[Boolean](criterion)
		if err != nil {
			return nil, false, err
		}

		// If criterion is true, return the value of the true-result argument
		if ok && bool(criterionBool) {
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
		return nil, true, nil
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
			return nil, false, fmt.Errorf("expected string substring parameter")
		}

		// If substring is an empty string (''), the function returns 0
		if substring == "" {
			return Collection{Integer(0)}, true, nil
		}

		// Return the index of the substring in the string
		index := strings.Index(string(s), string(substring))
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

		// Evaluate the pattern parameter
		patternCollection, _, err := evaluate(ctx, nil, parameters[0])
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
		substitutionCollection, _, err := evaluate(ctx, nil, parameters[1])
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
		regex, err := regexp.Compile(string(regexStr))
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
		if len(parameters) != 2 {
			return nil, false, fmt.Errorf("expected two parameters (regex, substitution)")
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

		// Evaluate the substitution parameter
		substitutionCollection, _, err := evaluate(ctx, nil, parameters[1])
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

		// Compile the regular expression
		regex, err := regexp.Compile(string(regexStr))
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

		precision := 0
		if len(parameters) == 1 {
			precisionCollection, _, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, false, err
			}

			// If the precision collection is empty, use default precision
			if len(precisionCollection) > 0 {
				precisionInt, ok, err := Singleton[Integer](precisionCollection)
				if err != nil {
					return nil, false, err
				}
				if !ok {
					return nil, false, fmt.Errorf("expected integer precision parameter")
				}

				// Precision must be >= 0
				if precisionInt < 0 {
					return nil, false, fmt.Errorf("precision must be >= 0")
				}

				precision = int(precisionInt)
			}
		}

		d, ok, err := Singleton[Decimal](target)
		if err == nil && ok {
			// Create a context with the specified precision
			apdCtx := apdContext(ctx).WithPrecision(uint32(precision + 1))

			// Round the decimal to the specified precision
			var rounded apd.Decimal
			_, err = apdCtx.Round(&rounded, d.Value)
			if err != nil {
				return nil, false, err
			}

			return Collection{Decimal{Value: &rounded}}, true, nil
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

		logger, err := traceLogger(ctx)
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
				projection, _, err := evaluate(ctx, elem, parameters[1], FunctionScope{index: i, total: len(target)})
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

		if len(parameters) == 2 {
			// If init value is provided, evaluate it
			var err error
			result, _, err = evaluate(ctx, nil, parameters[1])
			if err != nil {
				return nil, false, err
			}
		}

		total := Collection{}
		totalOrdered := inputOrdered

		// Iterate over the target collection
		for i, elem := range target {
			var ordered bool
			total, ordered, err = evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
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
