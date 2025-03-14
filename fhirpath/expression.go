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

func Parse(path string) (Expression, error) {
	parser, err := parse(path)
	if err != nil {
		return Expression{}, err
	}

	return Expression{tree: parser.Expression()}, nil
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

func parse(path string) (*parser.FHIRPathParser, error) {
	errListener := SyntaxErrorListener{}
	inputStream := antlr.NewInputStream(path)

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

	return evalExpression(
		ctx,
		target, Collection{target},
		expr.tree,
		true,
	)
}

func evalExpression(
	ctx context.Context,
	root Element, target Collection,
	tree parser.IExpressionContext,
	isRoot bool,
) (Collection, error) {
	switch t := tree.(type) {
	case *parser.TermExpressionContext:
		return evalTerm(ctx, root, target, t.Term(), isRoot)
	case *parser.InvocationExpressionContext:
		expr, err := evalExpression(ctx, root, target, t.Expression(), isRoot)
		if err != nil {
			return nil, err
		}
		return evalInvocation(ctx, root, expr, t.Invocation(), false)
	case *parser.IndexerExpressionContext:
		expr, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		indexCollection, err := evalExpression(ctx, root, target, t.Expression(1), false)
		if err != nil {
			return nil, err
		}
		index, err := Singleton[Integer](indexCollection)
		if err != nil {
			return nil, err
		}
		if index == nil {
			return nil, fmt.Errorf("can not index with null index")
		}
		i := int(*index)
		if i >= len(target) {
			return nil, nil
		} else {
			return Collection{expr[i]}, nil
		}
	case *parser.PolarityExpressionContext:
		expr, err := evalExpression(ctx, root, target, t.Expression(), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "+":
			// noop
			return expr, nil
		case "-":
			return expr.Multiply(ctx, Collection{Integer(-1)})

		}
		return nil, nil
	case *parser.MultiplicativeExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "*":
			return left.Multiply(ctx, right)
		case "/":
			return left.Divide(ctx, right)
		case "div":
			return left.Div(ctx, right)
		case "mod":
			return left.Mod(ctx, right)
		}
		return nil, nil
	case *parser.AdditiveExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "+":
			return left.Add(ctx, right)
		case "-":
			return left.Subtract(ctx, right)
		case "&":
			return left.Concat(ctx, right)
		}
		return nil, nil
	case *parser.TypeExpressionContext:
		expr, err := evalExpression(ctx, root, target, t.Expression(), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()
		spec, err := evalQualifiedIdentifier(t.TypeSpecifier().QualifiedIdentifier())
		if err != nil {
			return nil, err
		}

		if len(expr) == 1 {
			switch op {
			case "is":
				r, err := isType(ctx, expr[0], spec)
				if err != nil {
					return nil, err
				}
				return Collection{r}, nil
			case "as":
				c, err := asType(ctx, expr[0], spec)
				if err != nil {
					return nil, err
				}
				if c != nil {
					return c, nil
				}
				return nil, nil

			}
		}
		return nil, nil

	case *parser.UnionExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}

		return left.Union(right), nil

	case *parser.InequalityExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		cmp, err := left.Cmp(right)
		if err != nil {
			return nil, err
		}
		if cmp == nil {
			return nil, nil
		}
		switch op {
		case "<=":
			if *cmp <= 0 {
				return Collection{Boolean(true)}, nil
			}
		case "<":
			if *cmp < 0 {
				return Collection{Boolean(true)}, nil
			}
		case ">":
			if *cmp > 0 {
				return Collection{Boolean(true)}, nil
			}
		case ">=":
			if *cmp >= 0 {
				return Collection{Boolean(true)}, nil
			}
		}
		return Collection{Boolean(false)}, nil

	case *parser.EqualityExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "=":
			eq := left.Equal(right)
			if eq != nil {
				return Collection{Boolean(*eq)}, nil
			}
		case "~":
			eq := left.Equivalent(right)
			if eq != nil {
				return Collection{Boolean(*eq)}, nil
			}
		case "!=":
			eq := left.Equal(right)
			if eq != nil {
				return Collection{Boolean(!*eq)}, nil
			}
		case "!~":
			eq := left.Equivalent(right)
			if eq != nil {
				return Collection{Boolean(!*eq)}, nil
			}
		}
		return nil, nil

	case *parser.MembershipExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		switch op {
		case "in":
			if len(left) == 0 {
				return nil, nil
			} else if len(left) > 1 {
				return nil, fmt.Errorf("left operand of \"in\" (membership) has more than 1 value")
			}
			return Collection{Boolean(right.Contains(left[0]))}, nil
		case "contains":
			if len(right) == 0 {
				return nil, nil
			} else if len(right) > 1 {
				return nil, fmt.Errorf("left operand of \"contains\" (membership) has more than 1 value")
			}
			return Collection{Boolean(left.Contains(right[0]))}, nil
		}
		return nil, nil

	case *parser.AndExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}

		leftSingle, err := Singleton[Boolean](left)
		if err != nil {
			return nil, err
		}
		rightSingle, err := Singleton[Boolean](right)
		if err != nil {
			return nil, err
		}

		if leftSingle != nil && *leftSingle == true &&
			rightSingle != nil && *rightSingle == true {
			return Collection{Boolean(true)}, nil
		} else if leftSingle != nil && *leftSingle == false {
			return Collection{Boolean(false)}, nil
		} else if rightSingle != nil && *rightSingle == false {
			return Collection{Boolean(false)}, nil
		}
		return nil, nil

	case *parser.OrExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}
		op := t.GetChild(1).(antlr.ParseTree).GetText()

		leftSingle, err := Singleton[Boolean](left)
		if err != nil {
			return nil, err
		}
		rightSingle, err := Singleton[Boolean](right)
		if err != nil {
			return nil, err
		}

		switch op {
		case "or":
			if leftSingle != nil && *leftSingle == false &&
				rightSingle != nil && *rightSingle == false {
				return Collection{Boolean(false)}, nil
			} else if leftSingle != nil && *leftSingle == true {
				return Collection{Boolean(true)}, nil
			} else if rightSingle != nil && *rightSingle == true {
				return Collection{Boolean(true)}, nil
			}
		case "xor":
			if (leftSingle != nil && *leftSingle == true && rightSingle != nil && *rightSingle == false) ||
				(leftSingle != nil && *leftSingle == false && rightSingle != nil && *rightSingle == true) {
				return Collection{Boolean(true)}, nil
			} else if leftSingle != nil && rightSingle != nil &&
				*rightSingle == *leftSingle {
				return Collection{Boolean(false)}, nil
			}
		}
		return nil, nil

	case *parser.ImpliesExpressionContext:
		left, err := evalExpression(ctx, root, target, t.Expression(0), isRoot)
		if err != nil {
			return nil, err
		}
		right, err := evalExpression(ctx, root, target, t.Expression(1), isRoot)
		if err != nil {
			return nil, err
		}

		leftSingle, err := Singleton[Boolean](left)
		if err != nil {
			return nil, err
		}
		rightSingle, err := Singleton[Boolean](right)
		if err != nil {
			return nil, err
		}

		if leftSingle != nil && *leftSingle == true && rightSingle != nil {
			return Collection{*rightSingle}, nil
		} else if leftSingle != nil && *leftSingle == false {
			return Collection{Boolean(true)}, nil
		} else if leftSingle == nil && rightSingle != nil && *rightSingle == true {
			return Collection{Boolean(true)}, nil
		}
		return nil, nil

	default:
		panic(fmt.Sprintf("unexpected expression %T", tree))
	}
}

func evalTerm(
	ctx context.Context,
	root Element, target Collection,
	tree parser.ITermContext,
	isRoot bool,
) (Collection, error) {
	switch t := tree.(type) {
	case *parser.InvocationTermContext:
		return evalInvocation(ctx, root, target, t.Invocation(), isRoot)
	case *parser.LiteralTermContext:
		return evalLiteral(t.Literal())
	case *parser.ExternalConstantTermContext:
		return evalExternalConstant(ctx, t.ExternalConstant())
	case *parser.ParenthesizedTermContext:
		return evalExpression(ctx, root, target, t.Expression(), isRoot)
	default:
		return nil, fmt.Errorf("unexpected term %T", tree)
	}
}

func evalLiteral(
	tree parser.ILiteralContext,
) (Collection, error) {
	s := tree.GetText()

	switch tt := tree.(type) {
	case *parser.NullLiteralContext:
		return nil, nil
	case *parser.BooleanLiteralContext:
		if s == "true" {
			return Collection{Boolean(true)}, nil
		} else if s == "false" {
			return Collection{Boolean(false)}, nil
		} else {
			return nil, fmt.Errorf("expected boolean literal, got %s", s)
		}
	case *parser.StringLiteralContext:
		unescaped, err := unescape(s)
		return Collection{String(unescaped)}, err
	case *parser.NumberLiteralContext:
		if strings.Contains(s, ".") {
			d, _, err := apd.NewFromString(s)
			return Collection{Decimal{Value: d}}, err
		}

		i, err := strconv.Atoi(s)
		return Collection{Integer(i)}, err
	case *parser.DateLiteralContext:
		d, err := ParseDate(s)
		return Collection{d}, err
	case *parser.TimeLiteralContext:
		t, err := ParseTime(s)
		return Collection{t}, err
	case *parser.DateTimeLiteralContext:
		dt, err := ParseDateTime(s)
		return Collection{dt}, err
	case *parser.QuantityLiteralContext:
		q := tt.Quantity()
		v, _, err := apd.NewFromString(q.NUMBER().GetText())
		u := q.Unit().GetText()
		return Collection{Quantity{Value: Decimal{Value: v}, Unit: String(u)}}, err
	default:
		return nil, fmt.Errorf("unexpected term %T: %v", tree, tree)
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
) (Collection, error) {
	name := strings.TrimLeft(tree.GetText(), "%")
	value, ok := envValue(ctx, name)
	if !ok {
		return nil, fmt.Errorf("environment variable %q undefined", name)
	}
	return Collection{value}, nil
}

func Singleton[T Element](c Collection) (*T, error) {
	if len(c) == 0 {
		return nil, nil
	} else if len(c) > 1 {
		return nil, fmt.Errorf("can not convert to singleton: collection contains > 1 values")
	}

	// convert to input type
	v, err := ElementTo[T](c[0], false)

	// if not convertible but contains a single value, evaluate to true
	if _, wantBool := any(v).(Boolean); err != nil && wantBool {
		b := Boolean(true)
		return any(&b).(*T), nil
	}

	return v, err
}
