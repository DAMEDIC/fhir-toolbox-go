package fhirpath

import (
	"context"
	"fmt"
	"maps"
)

type functionCtxKey struct{}

type functionScope struct {
	this  Element
	index int
	total int
}

func withFunctionScope(
	ctx context.Context,
	this Element,
	index, total int,
) context.Context {
	return context.WithValue(
		ctx,
		functionCtxKey{},
		functionScope{
			this:  this,
			index: index,
			total: total,
		},
	)
}

func getFunctionScope(ctx context.Context) (functionScope, error) {
	fnCtx, ok := ctx.Value(functionCtxKey{}).(functionScope)
	if !ok {
		return functionScope{}, fmt.Errorf("not in function context")
	}
	return fnCtx, nil
}

type Functions map[string]Function
type Function = func(
	ctx context.Context,
	root, target Element,
	arguments ...Collection,
) (Collection, error)

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
	"hello": func(
		ctx context.Context,
		root, target Element,
		arguments ...Collection,
	) (Collection, error) {
		if len(arguments) != 0 {
			return nil, fmt.Errorf("expected zero arguments")
		}

		fmt.Println("hello", target)

		return nil, nil
	},
}
