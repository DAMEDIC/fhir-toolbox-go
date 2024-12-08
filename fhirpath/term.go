package fhirpath

import (
	parser "fhir-toolbox/fhirpath/parser/gen"
	"fmt"
)

func evalTerm(target Collection, tree parser.ITermContext) (Collection, error) {
	switch t := tree.(type) {
	case *parser.InvocationTermContext:
		return evalInvocation(target, t.Invocation())
	default:
		panic(fmt.Sprintf("unexpected term %T", tree))
	}
}
