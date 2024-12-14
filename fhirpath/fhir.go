package fhirpath

import "context"

// todo: move this to generated model packages

func FHIRContext() context.Context {
	return WithFHIR(context.Background())
}

func WithFHIR(
	ctx context.Context,
) context.Context {
	ctx = WithNamespace(ctx, "FHIR")
	ctx = WithTypes(ctx, []TypeInfo{
		ClassInfo{
			SimpleTypeInfo: SimpleTypeInfo{Namespace: "FHIR", Name: "Patient", BaseType: TypeSpecifier{"FHIR", "DomainResource"}},
			element:        nil,
		},
	})
	ctx = WithFunctions(ctx, fhirFunctions)
	return ctx
}

var fhirFunctions = Functions{}
