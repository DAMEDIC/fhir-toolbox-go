package fhirpath

import (
	"fmt"
	"strings"

	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	. "github.com/dave/jennifer/jen"
)

const fhirpathModuleName = "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"

type FHIRPathGenerator struct{}

func (g FHIRPathGenerator) GenerateType(f *File, rt ir.ResourceOrType) bool {
	for _, s := range rt.Structs {
		generateChildrenFunc(f, s)
		generateToBooleanFunc(f, s)
		generateToStringFunc(f, s)
		generateToIntegerFunc(f, s)
		generateToDecimalFunc(f, s)
		generateToDateFunc(f, s)
		generateToTimeFunc(f, s)
		generateToDateTimeFunc(f, s)
		generateToQuantityFunc(f, s)
		generateTypeInfoFunc(f, s)
	}

	return true
}

func generateChildrenFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("Children").Params(Id("name").Op("...").String()).
		Qual(fhirpathModuleName, "Collection").BlockFunc(func(g *Group) {
		g.Var().Id("children").Qual(fhirpathModuleName, "Collection")
		for _, f := range s.Fields {
			if s.IsPrimitive && f.Name == "Value" {
				continue
			}

			g.If(Len(Id("name")).Op("==").Lit(0).Op("||").
				Qual("slices", "Contains").Call(Id("name"), Lit(f.MarshalName))).BlockFunc(func(g *Group) {
				if f.Multiple {
					g.For(List(Id("_"), Id("v")).Op(":=").Range().Id("r").Dot(f.Name)).Block(
						Id("children").Op("=").Append(Id("children"), Id("v")),
					)
				} else if f.Optional {
					g.If(Id("r").Dot(f.Name).Op("!=").Nil()).BlockFunc(func(g *Group) {
						if !s.IsResource && f.Name == "Id" {
							g.Id("children").Op("=").Append(
								Id("children"),
								Qual(fhirpathModuleName, "String").Call(Op("*").Id("r").Dot(f.Name)),
							)
						} else if f.Polymorph || f.PossibleTypes[0].IsNestedResource {
							g.Id("children").Op("=").Append(
								Id("children"),
								Id("r").Dot(f.Name),
							)
						} else {
							g.Id("children").Op("=").Append(
								Id("children"),
								Op("*").Id("r").Dot(f.Name),
							)
						}
					})
				} else {
					if s.Name == "Extension" && f.Name == "Url" {
						g.Id("children").Op("=").Append(
							Id("children"),
							Qual(fhirpathModuleName, "String").Call(Id("r").Dot(f.Name)),
						)
					} else {
						g.Id("children").Op("=").Append(
							Id("children"),
							Id("r").Dot(f.Name),
						)
					}
				}
			})
		}
		g.Return(Id("children"))
	})
}

func generateToBooleanFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToBoolean").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "Boolean"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Boolean", s.Name))),
		),
	)
}

func generateToStringFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToString").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "String"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to String", s.Name))),
		),
	)
}

func generateToIntegerFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToInteger").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "Integer"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Integer", s.Name))),
		),
	)
}

func generateToDecimalFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToDecimal").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "Decimal"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Decimal", s.Name))),
		),
	)
}

func generateToDateFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToDate").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "Date"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Date", s.Name))),
		),
	)
}

func generateToTimeFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToTime").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "Time"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Time", s.Name))),
		),
	)
}

func generateToDateTimeFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToDateTime").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "DateTime"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to DateTime", s.Name))),
		),
	)
}

func generateToQuantityFunc(f *File, s ir.Struct) {
	// TODO
	f.Func().Params(Id("r").Id(s.Name)).Id("ToQuantity").Params(Id("explicit").Bool()).Params(
		Op("*").Qual(fhirpathModuleName, "Quantity"),
		Error(),
	).Block(
		Return(
			Nil(),
			Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Quantity", s.Name))),
		),
	)
}

func generateTypeInfoFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("TypeInfo").Params().Qual(fhirpathModuleName, "TypeInfo").
		Block(ReturnFunc(func(g *Group) {
			generateType(g, s)
		}))
}

func (g FHIRPathGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	contextFile := f("fhirpath", strings.ToLower(release))

	generateContextFunc(contextFile)
	generateWithContext(contextFile)
	generateTypes(contextFile, rt)
	generateFunctions(contextFile)
}

func generateContextFunc(f *File) *Statement {
	return f.Func().Id("Context").Params().Qual("context", "Context").Block(
		Return(
			Id("WithContext").Call(Qual("context", "Background").Call()),
		),
	)
}

func generateWithContext(contextFile *File) *Statement {
	return contextFile.Func().Id("WithContext").Params(
		Id("ctx").Qual("context", "Context"),
	).Qual("context", "Context").Block(
		Id("ctx").Op("=").Qual(fhirpathModuleName, "WithNamespace").Call(
			Id("ctx"),
			Lit("FHIR"),
		),
		Id("ctx").Op("=").Qual(fhirpathModuleName, "WithTypes").Call(
			Id("ctx"),
			Id("allFHIRPathTypes"),
		),
		Id("ctx").Op("=").Qual(fhirpathModuleName, "WithFunctions").Call(
			Id("ctx"),
			Id("fhirFunctions"),
		),
		Return(Id("ctx")),
	)
}

func generateFunctions(f *File) *Statement {
	return f.Var().Id("fhirFunctions").Op("=").Qual(fhirpathModuleName, "Functions").Values()
}
