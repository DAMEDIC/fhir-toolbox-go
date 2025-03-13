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
	tree parser.IInvocationContext,
	isRoot bool,
) (Collection, error) {
	switch t := tree.(type) {
	case *parser.MemberInvocationContext:
		ident, err := evalIdentifier(t.Identifier())
		if err != nil {
			return nil, err
		}

		if isRoot {
			expectedType, ok := resolveType(ctx, TypeSpecifier{Name: ident})
			if ok {
				rootType := root.TypeInfo()
				if !subTypeOf(ctx, rootType, expectedType) {
					return nil, fmt.Errorf("expected element of type %s, got %s", expectedType, rootType)
				}
				return target, nil
			}
		}

		var members Collection
		for _, e := range target {
			members = append(members, e.Children(ident)...)
		}
		return members, nil
	case *parser.FunctionInvocationContext:
		return evalFunc(ctx, root, target, t.Function())
	case *parser.ThisInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, err
		}
		return Collection{scope.this}, nil
	case *parser.IndexInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, err
		}
		return Collection{Integer(scope.index)}, nil
	case *parser.TotalInvocationContext:
		scope, err := getFunctionScope(ctx)
		if err != nil {
			return nil, err
		}
		return Collection{Integer(scope.total)}, nil
	default:
		return nil, fmt.Errorf("unexpected invocation %T", tree)
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
	tree parser.IFunctionContext,
) (Collection, error) {
	ident, err := evalIdentifier(tree.Identifier())
	if err != nil {
		return nil, err
	}

	paramList := tree.ParamList()
	if paramList == nil {
		return callFunc(ctx, root, target, ident)
	}

	return callFunc(ctx, root, target, ident, paramList.AllExpression()...)
}

func callFunc(
	ctx context.Context,
	root Element, target Collection,
	ident string, paramTerms ...parser.IExpressionContext,
) (Collection, error) {
	fn, ok := getFunction(ctx, ident)
	if !ok {
		return nil, fmt.Errorf("function \"%s\" not found", ident)
	}

	paramExprs := make([]Expression, 0, len(paramTerms))
	for _, param := range paramTerms {
		paramExprs = append(paramExprs, Expression{param})
	}

	return fn(
		ctx,
		root, target,
		paramExprs,
		func(
			ctx context.Context,
			target Element,
			expr Expression,
			fnScope ...FunctionScope,
		) (Collection, error) {
			if len(fnScope) > 0 {
				ctx = withFunctionScope(ctx, functionScope{
					this:  target,
					index: fnScope[0].index,
					total: fnScope[0].total,
				})
			}
			return evalExpression(ctx,
				root, Collection{target},
				expr.tree, false,
			)
		},
	)
}
