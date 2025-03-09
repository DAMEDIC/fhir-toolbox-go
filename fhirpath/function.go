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
	parameters []Expression,
	evalParameter EvalParamFunc,
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

type EvalParamFunc = func(expr Expression) (Collection, error)

var defaultFunctions = Functions{
	"is": func(
		ctx context.Context,
		root, target Element,
		parameters []Expression,
		evalParameter EvalParamFunc,
	) (Collection, error) {
		if len(parameters) != 1 {
			return nil, fmt.Errorf("expected single type argument")
		}
		typeSpec := ParseTypeSpecifier(parameters[0].String())

		r, err := isType(ctx, target, typeSpec)
		if err != nil {
			return nil, err
		}
		return Collection{r}, nil
	},
	"as": func(
		ctx context.Context,
		root, target Element,
		parameters []Expression,
		evalParameter EvalParamFunc,
	) (Collection, error) {
		if len(parameters) != 1 {
			return nil, fmt.Errorf("expected single type specifier parameter")
		}
		typeSpec := ParseTypeSpecifier(parameters[0].String())

		c, err := asType(ctx, target, typeSpec)
		if err != nil {
			return nil, err
		}
		return c, nil
	},
	"not": func(
		ctx context.Context,
		root, target Element,
		parameters []Expression,
		evalParameter EvalParamFunc,
	) (Collection, error) {
		if len(parameters) != 0 {
			return nil, fmt.Errorf("expected no parameter")
		}
		b, err := target.ToBoolean(false)
		if err != nil {
			return nil, err
		}
		if b == nil {
			return nil, nil
		}
		return Collection{!*b}, nil
	},
}
