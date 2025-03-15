package fhirpath

import (
	"context"
	"fmt"
	parser "github.com/DAMEDIC/fhir-toolbox-go/fhirpath/parser/gen"
	"strings"
)

func evalInvocation(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	tree parser.IInvocationContext,
	isRoot bool,
) (result Collection, resultOrdered bool, err error) {
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
				return target, inputOrdered, nil
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
		if err != nil {
			return nil, false, err
		}
		return Collection{scope.this}, inputOrdered, nil
	case *parser.IndexInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, false, err
		}
		return Collection{Integer(scope.index)}, inputOrdered, nil
	case *parser.TotalInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, false, err
		}
		if scope.total == 0 {
			return nil, false, fmt.Errorf("$total not defined (only in aggregate)")
		}
		return Collection{Integer(scope.total)}, inputOrdered, nil
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
) (result Collection, resultOrdered bool, err error) {
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
) (result Collection, resultOrdered bool, err error) {
	fn, ok := getFunction(ctx, ident)
	if !ok {
		return nil, false, fmt.Errorf("function \"%s\" not found", ident)
	}

	paramExprs := make([]Expression, 0, len(paramTerms))
	for _, param := range paramTerms {
		paramExprs = append(paramExprs, Expression{param})
	}

	return fn(
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
				expr.tree, false,
			)
		},
	)
}
