package fhirpath

import (
	"errors"
	parser "fhir-toolbox/fhirpath/parser/gen"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type Expression struct {
	tree parser.IExpressionContext
}

func (e Expression) String() string {
	return e.tree.GetText()
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

func Parse(path string) (Expression, error) {
	errListener := SyntaxErrorListener{}
	inputStream := antlr.NewInputStream(path)

	lexer := parser.NewFHIRPathLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewFHIRPathParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(&errListener)

	return Expression{tree: parser.Expression()}, errors.Join(errListener.Errors...)
}

func MustParse(path string) Expression {
	expr, err := Parse(path)
	if err != nil {
		panic(err)
	}
	return expr
}

func Evaluate(target Element, path Expression) (Collection, error) {
	return evalExpression(Collection{target}, path.tree)
}

func evalExpression(target Collection, tree parser.IExpressionContext) (Collection, error) {
	switch t := tree.(type) {
	case *parser.InvocationExpressionContext:
		expr, err := evalExpression(target, t.Expression())
		if err != nil {
			return nil, err
		}
		return evalInvocation(expr, t.Invocation())
	case *parser.TermExpressionContext:
		return evalTerm(target, t.Term())

	// todo: all the other cases...

	default:
		panic(fmt.Sprintf("unexpected expression %T", tree))
	}
}
