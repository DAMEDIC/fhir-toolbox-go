package fhirpath

import (
	parser "fhir-toolbox/fhirpath/parser/gen"
	"fmt"
)

func evalInvocation(target Collection, tree parser.IInvocationContext) (Collection, error) {
	switch t := tree.(type) {
	case *parser.MemberInvocationContext:
		ident, err := evalIdentifier(t.Identifier())
		if err != nil {
			return nil, err
		}
		return target.Member(ident), nil
	default:
		panic(fmt.Sprintf("unexpected invocation %T", tree))
	}
}
