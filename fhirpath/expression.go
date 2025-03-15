package fhirpath

import (
	"context"
	"errors"
	"fmt"
	parser "github.com/DAMEDIC/fhir-toolbox-go/fhirpath/parser/gen"
	"github.com/antlr4-go/antlr/v4"
	"github.com/cockroachdb/apd/v3"
	"strconv"
	"strings"
)

type Expression struct {
	tree parser.IExpressionContext
}

func (e Expression) String() string {
	return e.tree.GetText()
}

func Parse(expr string) (Expression, error) {
	parser, err := parse(expr)
	if err != nil {
		return Expression{}, err
	}
	entireExpr := parser.EntireExpression()
	if entireExpr.EOF() == nil {
		return Expression{}, fmt.Errorf(
			"can not parse expression at index %v: %v",
			len(entireExpr.Expression().GetText()), entireExpr.GetText(),
		)
	}
	return Expression{entireExpr.Expression()}, nil
}

func MustParse(path string) Expression {
	expr, err := Parse(path)
	if err != nil {
		panic(err)
	}
	return expr
}

type SyntaxError struct {
	line, column int
	msg          string
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: %s", s.line, s.column, s.msg)
}

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (c *SyntaxErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line, column int,
	msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, SyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func parse(expr string) (*parser.FHIRPathParser, error) {
	errListener := SyntaxErrorListener{}
	inputStream := antlr.NewInputStream(expr)

	lexer := parser.NewFHIRPathLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewFHIRPathParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(&errListener)

	return parser, errors.Join(errListener.Errors...)
}

func Evaluate(ctx context.Context, target Element, expr Expression) (Collection, error) {
	ctx = WithEnv(ctx, "context", target)
	ctx = WithEnv(ctx, "ucum", String("http://unitsofmeasure.org"))

	result, _, err := evalExpression(
		ctx,
		target, Collection{target},
		true,
		expr.tree,
		true,
	)
	return result, err
}

func evalExpression(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	tree parser.IExpressionContext,
	isRoot bool,
) (result Collection, resultOrdered bool, err error) {
	switch t := tree.(type) {
	case *parser.TermExpressionContext:
		return evalTerm(ctx, root, target, inputOrdered, t.Term(), isRoot)
	case *parser.InvocationExpressionContext:
		expr, ordered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(), isRoot)
		if err != nil {
			return nil, false, err
		}
		return evalInvocation(ctx, root, expr, ordered, t.Invocation(), false)
	case *parser.IndexerExpressionContext:
		expr, ordered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		if !ordered {
			return nil, false, errors.New("can not index into unordered collection")
		}
		indexCollection, _, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), false)
		if err != nil {
			return nil, false, err
		}
		index, err := Singleton[Integer](indexCollection)
		if err != nil {
			return nil, false, err
		}
		if index == nil {
			return nil, false, fmt.Errorf("can not index with null index")
		}
		i := int(*index)
		if i >= len(target) {
			return nil, false, nil
		} else {
			return Collection{expr[i]}, true, nil
		}
	case *parser.PolarityExpressionContext:
		expr, ordered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(0).(antlr.ParseTree).GetText()

		switch op {
		case "+":
			// noop
			return expr, ordered, nil
		case "-":
			result, err = expr.Multiply(ctx, Collection{Integer(-1)})
			return result, ordered, err

		}
		return nil, false, nil
	case *parser.MultiplicativeExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "*":
			result, err = left.Multiply(ctx, right)
		case "/":
			result, err = left.Divide(ctx, right)
		case "div":
			result, err = left.Div(ctx, right)
		case "mod":
			result, err = left.Mod(ctx, right)
		}
		return result, leftOrdered && rightOrdered, err
	case *parser.AdditiveExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "+":
			result, err = left.Add(ctx, right)
		case "-":
			result, err = left.Subtract(ctx, right)
		case "&":
			result, err = left.Concat(ctx, right)
		}
		return result, leftOrdered && rightOrdered, err
	case *parser.TypeExpressionContext:
		expr, ordered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()
		spec, err := evalQualifiedIdentifier(t.TypeSpecifier().QualifiedIdentifier())
		if err != nil {
			return nil, false, err
		}

		if len(expr) != 1 {
			return nil, false, fmt.Errorf("expected single input element")
		}

		switch op {
		case "is":
			r, err := isType(ctx, expr[0], spec)
			if err != nil {
				return nil, false, err
			}
			return Collection{r}, ordered, nil
		case "as":
			c, err := asType(ctx, expr[0], spec)
			if err != nil {
				return nil, false, err
			}
			if c != nil {
				return c, ordered, nil
			}
			return nil, false, nil

		}

		return nil, false, nil

	case *parser.UnionExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}

		return left.Union(right), leftOrdered && rightOrdered, nil

	case *parser.InequalityExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		cmp, err := left.Cmp(right)
		if err != nil {
			return nil, false, err
		}
		if cmp == nil {
			return nil, false, nil
		}

		result = Collection{Boolean(false)}
		switch op {
		case "<=":
			if *cmp <= 0 {
				result, err = Collection{Boolean(true)}, nil
			}
		case "<":
			if *cmp < 0 {
				result, err = Collection{Boolean(true)}, nil
			}
		case ">":
			if *cmp > 0 {
				result, err = Collection{Boolean(true)}, nil
			}
		case ">=":
			if *cmp >= 0 {
				result, err = Collection{Boolean(true)}, nil
			}
		}
		return result, leftOrdered && rightOrdered, err

	case *parser.EqualityExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		if !leftOrdered || !rightOrdered {
			return nil, false, fmt.Errorf("expected ordered inputs for equality expression")
		}

		switch op {
		case "=":
			eq := left.Equal(right)
			if eq != nil {
				result, err = Collection{Boolean(*eq)}, nil
			}
		case "~":
			eq := left.Equivalent(right)
			if eq != nil {
				result, err = Collection{Boolean(*eq)}, nil
			}
		case "!=":
			eq := left.Equal(right)
			if eq != nil {
				result, err = Collection{Boolean(!*eq)}, nil
			}
		case "!~":
			eq := left.Equivalent(right)
			if eq != nil {
				result, err = Collection{Boolean(!*eq)}, nil
			}
		}
		return result, leftOrdered && rightOrdered, err
	case *parser.MembershipExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "in":
			if len(left) == 0 {
				return nil, false, nil
			} else if len(left) > 1 {
				return nil, false, fmt.Errorf("left operand of \"in\" (membership) has more than 1 value")
			}
			result, err = Collection{Boolean(right.Contains(left[0]))}, nil
		case "contains":
			if len(right) == 0 {
				return nil, false, nil
			} else if len(right) > 1 {
				return nil, false, fmt.Errorf("left operand of \"contains\" (membership) has more than 1 value")
			}
			result, err = Collection{Boolean(left.Contains(right[0]))}, nil
		}
		return result, leftOrdered && rightOrdered, err

	case *parser.AndExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}

		leftSingle, err := Singleton[Boolean](left)
		if err != nil {
			return nil, false, err
		}
		rightSingle, err := Singleton[Boolean](right)
		if err != nil {
			return nil, false, err
		}

		if leftSingle != nil && *leftSingle == true &&
			rightSingle != nil && *rightSingle == true {
			result, err = Collection{Boolean(true)}, nil
		} else if leftSingle != nil && *leftSingle == false {
			result, err = Collection{Boolean(false)}, nil
		} else if rightSingle != nil && *rightSingle == false {
			result, err = Collection{Boolean(false)}, nil
		}
		return result, leftOrdered && rightOrdered, err

	case *parser.OrExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		leftSingle, err := Singleton[Boolean](left)
		if err != nil {
			return nil, false, err
		}
		rightSingle, err := Singleton[Boolean](right)
		if err != nil {
			return nil, false, err
		}

		switch op {
		case "or":
			if leftSingle != nil && *leftSingle == false &&
				rightSingle != nil && *rightSingle == false {
				result, err = Collection{Boolean(false)}, nil
			} else if leftSingle != nil && *leftSingle == true {
				result, err = Collection{Boolean(true)}, nil
			} else if rightSingle != nil && *rightSingle == true {
				result, err = Collection{Boolean(true)}, nil
			}
		case "xor":
			if (leftSingle != nil && *leftSingle == true && rightSingle != nil && *rightSingle == false) ||
				(leftSingle != nil && *leftSingle == false && rightSingle != nil && *rightSingle == true) {
				result, err = Collection{Boolean(true)}, nil
			} else if leftSingle != nil && rightSingle != nil &&
				*rightSingle == *leftSingle {
				result, err = Collection{Boolean(false)}, nil
			}
		}
		return result, leftOrdered && rightOrdered, err

	case *parser.ImpliesExpressionContext:
		left, leftOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(0), isRoot)
		if err != nil {
			return nil, false, err
		}
		right, rightOrdered, err := evalExpression(ctx, root, target, inputOrdered, t.Expression(1), isRoot)
		if err != nil {
			return nil, false, err
		}

		leftSingle, err := Singleton[Boolean](left)
		if err != nil {
			return nil, false, err
		}
		rightSingle, err := Singleton[Boolean](right)
		if err != nil {
			return nil, false, err
		}

		if leftSingle != nil && *leftSingle == true && rightSingle != nil {
			result, err = Collection{*rightSingle}, nil
		} else if leftSingle != nil && *leftSingle == false {
			result, err = Collection{Boolean(true)}, nil
		} else if leftSingle == nil && rightSingle != nil && *rightSingle == true {
			result, err = Collection{Boolean(true)}, nil
		}
		return result, leftOrdered && rightOrdered, err

	default:
		panic(fmt.Sprintf("unexpected expression %T", tree))
	}
}

func evalTerm(
	ctx context.Context,
	root Element, target Collection,
	inputOrdered bool,
	tree parser.ITermContext,
	isRoot bool,
) (result Collection, resultOrdered bool, err error) {
	switch t := tree.(type) {
	case *parser.InvocationTermContext:
		return evalInvocation(ctx, root, target, inputOrdered, t.Invocation(), isRoot)
	case *parser.LiteralTermContext:
		return evalLiteral(t.Literal())
	case *parser.ExternalConstantTermContext:
		return evalExternalConstant(ctx, t.ExternalConstant())
	case *parser.ParenthesizedTermContext:
		return evalExpression(ctx, root, target, inputOrdered, t.Expression(), isRoot)
	default:
		return nil, false, fmt.Errorf("unexpected term %T", tree)
	}
}

func evalLiteral(
	tree parser.ILiteralContext,
) (result Collection, resultOrdered bool, err error) {
	s := tree.GetText()

	switch tt := tree.(type) {
	case *parser.NullLiteralContext:
		return nil, true, nil
	case *parser.BooleanLiteralContext:
		if s == "true" {
			return Collection{Boolean(true)}, true, nil
		} else if s == "false" {
			return Collection{Boolean(false)}, true, nil
		} else {
			return nil, false, fmt.Errorf("expected boolean literal, got %s", s)
		}
	case *parser.StringLiteralContext:
		unescaped, err := unescape(s[1 : len(s)-1])
		return Collection{String(unescaped)}, true, err
	case *parser.NumberLiteralContext:
		if strings.Contains(s, ".") {
			d, _, err := apd.NewFromString(s)
			return Collection{Decimal{Value: d}}, true, err
		}

		i, err := strconv.Atoi(s)
		return Collection{Integer(i)}, true, err
	case *parser.DateLiteralContext:
		d, err := ParseDate(s)
		return Collection{d}, true, err
	case *parser.TimeLiteralContext:
		t, err := ParseTime(s)
		return Collection{t}, true, err
	case *parser.DateTimeLiteralContext:
		dt, err := ParseDateTime(s)
		return Collection{dt}, true, err
	case *parser.QuantityLiteralContext:
		q := tt.Quantity()
		v, _, err := apd.NewFromString(q.NUMBER().GetText())
		u := q.Unit().GetText()
		return Collection{Quantity{Value: Decimal{Value: v}, Unit: String(u)}}, true, err
	default:
		return nil, false, fmt.Errorf("unexpected term %T: %v", tree, tree)
	}
}

type envKey string

func WithEnv(ctx context.Context, name string, value Element) context.Context {
	return context.WithValue(ctx, envKey(name), value)
}

func envValue(ctx context.Context, name string) (Element, bool) {
	val, ok := ctx.Value(envKey(name)).(Element)
	return val, ok
}

func evalExternalConstant(
	ctx context.Context,
	tree parser.IExternalConstantContext,
) (result Collection, resultOrdered bool, err error) {
	name := strings.TrimLeft(tree.GetText(), "%")
	value, ok := envValue(ctx, name)
	if !ok {
		return nil, false, fmt.Errorf("environment variable %q undefined", name)
	}
	return Collection{value}, true, nil
}

func Singleton[T Element](c Collection) (*T, error) {
	if len(c) == 0 {
		return nil, nil
	} else if len(c) > 1 {
		return nil, fmt.Errorf("can not convert to singleton: collection contains > 1 values")
	}

	// convert to input type
	v, err := elementTo[T](c[0], false)

	// if not convertible but contains a single value, evaluate to true
	if _, wantBool := any(v).(Boolean); err != nil && wantBool {
		b := Boolean(true)
		return any(&b).(*T), nil
	}

	return v, err
}
