package fhirpath

import (
	"context"
	"fmt"
	"strings"

	parser "github.com/DAMEDIC/fhir-toolbox-go/fhirpath/parser/gen"
)

func evalInvocation(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	tree parser.IInvocationContext,
	isRoot bool,
) (Collection, bool, error) {
	switch t := tree.(type) {
	case *parser.MemberInvocationContext:
		ident, err := evalIdentifier(t.Identifier())
		if err != nil {
			return nil, false, err
		}

		if isRoot {
			expectedType, ok := resolveType(ctx, TypeSpecifier{Name: ident})
			if ok {
				rootType := root.TypeInfo()
				if !subTypeOf(ctx, rootType, expectedType) {
					return nil, false, fmt.Errorf("expected element of type %s, got %s", expectedType, rootType)
				}
				return Collection{root}, inputOrdered, nil
			}
		}

		var members Collection
		for _, e := range target {
			members = append(members, e.Children(ident)...)
		}
		return members, inputOrdered, nil
	case *parser.FunctionInvocationContext:
		return evalFunc(ctx, root, target, inputOrdered, t.Function())
	case *parser.ThisInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err == nil {
			return Collection{scope.this}, true, nil
		}
		return Collection{root}, true, nil
	case *parser.IndexInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, false, err
		}
		return Collection{Integer(scope.index)}, true, nil
	case *parser.TotalInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, false, err
		}
		if !scope.aggregate {
			return nil, false, fmt.Errorf("$total not defined (only in aggregate)")
		}
		return scope.total, true, nil
	default:
		return nil, false, fmt.Errorf("unexpected invocation %T", tree)
	}
}

func evalQualifiedIdentifier(tree parser.IQualifiedIdentifierContext) (TypeSpecifier, error) {
	var idents []string
	for _, i := range tree.AllIdentifier() {
		ident, err := evalIdentifier(i)
		if err != nil {
			return TypeSpecifier{}, err
		}
		idents = append(idents, ident)
	}

	return TypeSpecifier{
		Namespace: strings.Join(idents[:len(idents)-1], "."),
		Name:      idents[len(idents)-1],
	}, nil
}

func evalIdentifier(tree parser.IIdentifierContext) (string, error) {
	s := tree.GetText()

	if tree.DELIMITEDIDENTIFIER() != nil {
		return unescape(s[1 : len(s)-1])
	}

	return s, nil
}

func evalFunc(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	tree parser.IFunctionContext,
) (Collection, bool, error) {
	ident, err := evalIdentifier(tree.Identifier())
	if err != nil {
		return nil, false, err
	}

	paramList := tree.ParamList()
	if paramList == nil {
		return callFunc(ctx, root, target, inputOrdered, ident)
	}

	return callFunc(ctx, root, target, inputOrdered, ident, paramList.AllExpression()...)
}

func callFunc(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	ident string, paramTerms ...parser.IExpressionContext,
) (Collection, bool, error) {
	fn, ok := getFunction(ctx, ident)
	if !ok {
		return nil, false, fmt.Errorf("function \"%s\" not found", ident)
	}

	paramExprs := make([]Expression, 0, len(paramTerms))
	for _, param := range paramTerms {
		paramExprs = append(paramExprs, Expression{param})
	}

	result, ordered, err := fn(
		ctx,
		root, target,
		inputOrdered,
		paramExprs,
		func(
			ctx context.Context,
			target Element,
			expr Expression,
			fnScope ...FunctionScope,
		) (result Collection, resultOrdered bool, err error) {
			if len(fnScope) > 0 {
				scope := functionScope{
					this:  target,
					index: fnScope[0].index,
				}
				if ident == "aggregate" {
					scope.aggregate = true
					scope.total = fnScope[0].total
				}
				ctx = withFunctionScope(ctx, scope)
			}
			var targetCollection Collection
			if target != nil {
				targetCollection = Collection{target}
			}
			return evalExpression(ctx,
				root, targetCollection,
				true,
				expr.tree, true,
			)
		},
	)
	if err != nil {
		return nil, false, err
	}
	return result, ordered, nil
}
