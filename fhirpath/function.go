package fhirpath

import (
	"context"
	"fmt"
	"maps"
	"sync"
)

type Functions map[string]Function
type Function = func(
	ctx context.Context,
	root Element, target Collection,
	parameters []Expression,
	evaluate EvaluateFunc,
) (Collection, error)

type EvaluateFunc = func(
	ctx context.Context,
	target Element,
	expr Expression,
	fnScope ...FunctionScope,
) (Collection, error)

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
		return maps.Clone(defaultFunctions())
	}
	return fns
}

func getFunction(ctx context.Context, name string) (Function, bool) {
	fns := getFunctions(ctx)
	fn, ok := fns[name]
	return fn, ok
}

var defaultFunctions = sync.OnceValue(func() Functions {
	return Functions{
		"is": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(target) != 1 {
				return nil, fmt.Errorf("expected single input element")
			}
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single type argument")
			}
			typeSpec := ParseTypeSpecifier(parameters[0].String())

			r, err := isType(ctx, target[0], typeSpec)
			if err != nil {
				return nil, err
			}
			return Collection{r}, nil
		},
		"as": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(target) != 1 {
				return nil, fmt.Errorf("expected single input element")
			}
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single type specifier parameter")
			}
			typeSpec := ParseTypeSpecifier(parameters[0].String())

			c, err := asType(ctx, target[0], typeSpec)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		"not": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(target) != 1 {
				return nil, fmt.Errorf("expected single input element")
			}
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameter")
			}
			b, err := target[0].ToBoolean(false)
			if err != nil {
				return nil, err
			}
			if b == nil {
				return nil, nil
			}
			return Collection{!*b}, nil
		},
		"empty": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}
			return Collection{Boolean(len(target) == 0)}, nil
		},
		"exists": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) > 1 {
				return nil, fmt.Errorf("expected at most one criteria parameter")
			}

			if len(parameters) == 0 {
				return Collection{Boolean(len(target) > 0)}, nil
			}

			// With criteria, equivalent to where(criteria).exists()
			for i, elem := range target {
				criteria, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
				if err != nil {
					return nil, err
				}

				if len(criteria) == 1 {
					b, err := criteria[0].ToBoolean(false)
					if err != nil {
						return nil, err
					}
					if b != nil && *b {
						// Found at least one element that matches the criteria
						return Collection{Boolean(true)}, nil
					}
				}
			}

			return Collection{Boolean(false)}, nil
		},
		"all": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single criteria parameter")
			}

			// If the input collection is empty, the result is true
			if len(target) == 0 {
				return Collection{Boolean(true)}, nil
			}

			for i, elem := range target {
				criteria, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
				if err != nil {
					return nil, err
				}

				if len(criteria) == 1 {
					b, err := criteria[0].ToBoolean(false)
					if err != nil {
						return nil, err
					}
					if b == nil || !*b {
						// Found at least one element that doesn't match the criteria
						return Collection{Boolean(false)}, nil
					}
				} else {
					return Collection{Boolean(false)}, nil
				}
			}

			return Collection{Boolean(true)}, nil
		},
		"allTrue": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}

			// If the input collection is empty, the result is true
			if len(target) == 0 {
				return Collection{Boolean(true)}, nil
			}

			for _, elem := range target {
				b, err := elem.ToBoolean(false)
				if err != nil {
					return nil, err
				}
				if b == nil || !*b {
					return Collection{Boolean(false)}, nil
				}
			}
			return Collection{Boolean(true)}, nil
		},
		"anyTrue": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}

			// If the input collection is empty, the result is false
			if len(target) == 0 {
				return Collection{Boolean(false)}, nil
			}

			for _, elem := range target {
				b, err := elem.ToBoolean(false)
				if err != nil {
					return nil, err
				}
				if b != nil && *b {
					return Collection{Boolean(true)}, nil
				}
			}
			return Collection{Boolean(false)}, nil
		},
		"allFalse": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}

			// If the input collection is empty, the result is true
			if len(target) == 0 {
				return Collection{Boolean(true)}, nil
			}

			for _, elem := range target {
				b, err := elem.ToBoolean(false)
				if err != nil {
					return nil, err
				}
				if b == nil || *b {
					return Collection{Boolean(false)}, nil
				}
			}
			return Collection{Boolean(true)}, nil
		},
		"anyFalse": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}

			// If the input collection is empty, the result is false
			if len(target) == 0 {
				return Collection{Boolean(false)}, nil
			}

			for _, elem := range target {
				b, err := elem.ToBoolean(false)
				if err != nil {
					return nil, err
				}
				if b != nil && !*b {
					return Collection{Boolean(true)}, nil
				}
			}
			return Collection{Boolean(false)}, nil
		},
		"subsetOf": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single collection parameter")
			}

			// If the input collection is empty, the result is true
			if len(target) == 0 {
				return Collection{Boolean(true)}, nil
			}

			other, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, err
			}

			// If the other collection is empty, the result is false
			if len(other) == 0 {
				return Collection{Boolean(false)}, nil
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
					return Collection{Boolean(false)}, nil
				}
			}
			return Collection{Boolean(true)}, nil
		},
		"supersetOf": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single collection parameter")
			}

			other, err := evaluate(ctx, nil, parameters[0])
			if err != nil {
				return nil, err
			}

			// If the other collection is empty, the result is true
			if len(other) == 0 {
				return Collection{Boolean(true)}, nil
			}

			// If the input collection is empty, the result is false
			if len(target) == 0 {
				return Collection{Boolean(false)}, nil
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
					return Collection{Boolean(false)}, nil
				}
			}
			return Collection{Boolean(true)}, nil
		},
		"count": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}
			return Collection{Integer(len(target))}, nil
		},
		"distinct": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}

			// If the input collection is empty, the result is empty
			if len(target) == 0 {
				return nil, nil
			}

			var result Collection
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
			return result, nil
		},
		"isDistinct": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 0 {
				return nil, fmt.Errorf("expected no parameters")
			}

			// If the input collection is empty, the result is true
			if len(target) == 0 {
				return Collection{Boolean(true)}, nil
			}

			// Check if all elements are distinct
			for i := 0; i < len(target); i++ {
				for j := i + 1; j < len(target); j++ {
					if target[i].Equal(target[j]) {
						return Collection{Boolean(false)}, nil
					}
				}
			}
			return Collection{Boolean(true)}, nil
		},
		"where": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single criteria parameter")
			}

			// If the input collection is empty, the result is empty
			if len(target) == 0 {
				return nil, nil
			}

			var result Collection
			for i, elem := range target {
				criteria, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
				if err != nil {
					return nil, err
				}

				if len(criteria) == 1 {
					b, err := criteria[0].ToBoolean(false)
					if err != nil {
						return nil, err
					}
					if b != nil && *b {
						// Element matches the criteria, add it to the result
						result = append(result, elem)
					}
				}
			}

			return result, nil
		},
		"select": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single projection parameter")
			}

			// If the input collection is empty, the result is empty
			if len(target) == 0 {
				return nil, nil
			}

			var result Collection
			for i, elem := range target {
				projection, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(target)})
				if err != nil {
					return nil, err
				}

				// Add all items from the projection to the result (flatten)
				result = append(result, projection...)
			}

			return result, nil
		},
		"repeat": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single projection parameter")
			}

			// If the input collection is empty, the result is empty
			if len(target) == 0 {
				return nil, nil
			}

			var result Collection
			var current = target
			var newItems Collection

			for _, elem := range current {
				result = append(result, elem)
			}

			// Keep repeating the projection until no new items are found
			for {
				newItems = nil
				for i, elem := range current {
					projection, err := evaluate(ctx, elem, parameters[0], FunctionScope{index: i, total: len(current)})
					if err != nil {
						return nil, err
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

			return result, nil
		},
		"ofType": func(
			ctx context.Context,
			root Element, target Collection,
			parameters []Expression,
			evaluate EvaluateFunc,
		) (Collection, error) {
			if len(parameters) != 1 {
				return nil, fmt.Errorf("expected single type specifier parameter")
			}

			// If the input collection is empty, the result is empty
			if len(target) == 0 {
				return nil, nil
			}

			typeSpec := ParseTypeSpecifier(parameters[0].String())

			var result Collection
			for _, elem := range target {
				isOfType, err := isType(ctx, elem, typeSpec)
				if err != nil {
					return nil, err
				}

				// Check if isOfType is a Boolean with value true
				if boolVal, ok := isOfType.(Boolean); ok && bool(boolVal) {
					result = append(result, elem)
				}
			}

			return result, nil
		},
	}
})
