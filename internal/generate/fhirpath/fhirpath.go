package fhirpath

import (
	"fmt"
	"slices"
	"strings"

	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	. "github.com/dave/jennifer/jen"
)

const (
	fhirpathModuleName = "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	ucumSystem         = "http://unitsofmeasure.org"
)

var (
	stringTypes   = []string{"String", "Uri", "Code", "Oid", "Id", "Uuid", "Markdown", "Base64Binary"}
	intTypes      = []string{"Integer", "UnsignedInt", "PositiveInt"}
	dateTimeTypes = []string{"Date", "DateTime", "Instant"}
)

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
		generateEqualFunc(f, s)
		generateEquivalentFunc(f, s)
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
	f.Func().Params(Id("r").Id(s.Name)).Id("ToBoolean").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "Boolean"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if s.Name == "Boolean" {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				Id("v").Op(":=").Qual(fhirpathModuleName, "Boolean").Call(Op("*").Id("r").Dot("Value")),
				Return(
					Id("v"),
					True(),
					Nil(),
				),
			).Else().Block(
				Return(
					False(),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				False(),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Boolean", s.Name))),
			)
		}
	})
}

func generateToStringFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToString").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "String"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if slices.Contains(stringTypes, s.Name) {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				Id("v").Op(":=").Qual(fhirpathModuleName, "String").Call(Op("*").Id("r").Dot("Value")),
				Return(
					Id("v"),
					True(),
					Nil(),
				),
			).Else().Block(
				Return(
					Lit(""),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				Lit(""),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to String", s.Name))),
			)
		}
	})
}

func generateToIntegerFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToInteger").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "Integer"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if slices.Contains(intTypes, s.Name) {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				Id("v").Op(":=").Qual(fhirpathModuleName, "Integer").Call(Op("*").Id("r").Dot("Value")),
				Return(
					Id("v"),
					True(),
					Nil(),
				),
			).Else().Block(
				Return(
					Lit(0),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				Lit(0),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Integer", s.Name))),
			)
		}
	})
}

func generateToDecimalFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToDecimal").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "Decimal"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if s.Name == "Decimal" {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				Id("v").Op(":=").Qual(fhirpathModuleName, "Decimal").Values(Dict{
					Id("Value"): Id("r").Dot("Value"),
				}),
				Return(
					Id("v"),
					True(),
					Nil(),
				),
			).Else().Block(
				Return(
					Qual(fhirpathModuleName, "Decimal").Block(),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				Qual(fhirpathModuleName, "Decimal").Block(),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Decimal", s.Name))),
			)
		}
	})
}

func generateToDateFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToDate").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "Date"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if s.Name == "Date" {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				List(Id("v"), Err()).Op(":=").Qual(fhirpathModuleName, "ParseDate").Call(Op("*").Id("r").Dot("Value")),
				Return(
					Id("v"),
					True(),
					Err(),
				),
			).Else().Block(
				Return(
					Qual(fhirpathModuleName, "Date").Block(),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				Qual(fhirpathModuleName, "Date").Block(),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Date", s.Name))),
			)
		}
	})
}

func generateToTimeFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToTime").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "Time"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if s.Name == "Time" {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				List(Id("v"), Err()).Op(":=").Qual(fhirpathModuleName, "ParseTime").Call(Op("*").Id("r").Dot("Value")),
				Return(
					Id("v"),
					True(),
					Err(),
				),
			).Else().Block(
				Return(
					Qual(fhirpathModuleName, "Time").Block(),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				Qual(fhirpathModuleName, "Time").Block(),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Time", s.Name))),
			)
		}
	})
}

func generateToDateTimeFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToDateTime").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "DateTime"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if slices.Contains(dateTimeTypes, s.Name) {
			g.If(Id("r").Dot("Value").Op("!=").Nil()).Block(
				List(Id("v"), Err()).Op(":=").Qual(fhirpathModuleName, "ParseDateTime").Call(Op("*").Id("r").Dot("Value")),
				Return(
					Id("v"),
					True(),
					Err(),
				),
			).Else().Block(
				Return(
					Qual(fhirpathModuleName, "DateTime").Block(),
					False(),
					Nil(),
				),
			)
		} else {
			g.Return(
				Qual(fhirpathModuleName, "DateTime").Block(),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to DateTime", s.Name))),
			)
		}
	})
}

func generateToQuantityFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("ToQuantity").Params(Id("explicit").Bool()).Params(
		Qual(fhirpathModuleName, "Quantity"),
		Bool(),
		Error(),
	).BlockFunc(func(g *Group) {
		if s.Name == "Quantity" {
			g.If(Id("r").Dot("System").Op("==").Nil().Op("||").
				Id("r").Dot("System").Dot("Value").Op("==").Nil().Op("||").
				Op("*").Id("r").Dot("System").Dot("Value").Op("!=").Lit(ucumSystem)).Block(
				Return(
					Qual(fhirpathModuleName, "Quantity").Block(),
					False(),
					Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Quantity, no UCUM system", s.Name))),
				),
			).Else().If(Id("r").Dot("Value").Op("==").Nil()).Block(
				Return(
					Qual(fhirpathModuleName, "Quantity").Block(),
					False(),
					Nil(),
				),
			)

			g.Var().Id("unit").String()
			g.If(Id("r").Dot("Unit").Op("!=").Nil().Op("&&").
				Id("r").Dot("Unit").Dot("Value").Op("!=").Nil()).Block(
				Switch(Op("*").Id("r").Dot("Unit").Dot("Value")).Block(
					Case(Lit("a")).Id("unit").Op("=").Lit("year"),
					Case(Lit("mo")).Id("unit").Op("=").Lit("month"),
					Case(Lit("d")).Id("unit").Op("=").Lit("day"),
					Case(Lit("h")).Id("unit").Op("=").Lit("hour"),
					Case(Lit("min")).Id("unit").Op("=").Lit("minute"),
					Case(Lit("s")).Id("unit").Op("=").Lit("second"),
				),
			)

			g.Return(
				Qual(fhirpathModuleName, "Quantity").Values(Dict{
					Id("Value"): Qual(fhirpathModuleName, "Decimal").Values(Dict{
						Id("Value"): Id("r").Dot("Value").Dot("Value"),
					}),
					Id("Unit"): Qual(fhirpathModuleName, "String").Call(Id("unit")),
				}),
				True(),
				Nil(),
			)
		} else {
			g.Return(
				Qual(fhirpathModuleName, "Quantity").Block(),
				False(),
				Qual("errors", "New").Call(Lit(fmt.Sprintf("can not convert %s to Quantity", s.Name))),
			)
		}
	})
}

func generateTypeInfoFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("TypeInfo").Params().Qual(fhirpathModuleName, "TypeInfo").
		Block(ReturnFunc(func(g *Group) {
			generateType(g, s)
		}))
}

func generateEqualFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("Equal").Params(
		Id("other").Qual(fhirpathModuleName, "Element"),
		Id("_noReverseTypeConversion").Op("...").Bool(),
	).Params(
		Bool(),
		Bool(),
	).BlockFunc(func(g *Group) {
		if s.IsPrimitive {
			if s.Name == "Boolean" {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToBoolean").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToBoolean").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if slices.Contains(stringTypes, s.Name) {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToString").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToString").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if slices.Contains(intTypes, s.Name) {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToInteger").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToInteger").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if s.Name == "Decimal" {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToDecimal").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToDecimal").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if s.Name == "Time" {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToTime").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToTime").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if slices.Contains(dateTimeTypes, s.Name) {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToDateTime").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToDateTime").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if s.Name == "Quantity" {
				g.List(Id("a"), Id("ok"), Err()).Op(":=").Id("r").Dot("ToQuantity").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
				g.List(Id("b"), Id("ok"), Err()).Op(":=").Id("other").Dot("ToQuantity").Call(False())
				g.If(Err().Op("!=").Nil().Op("||").Op("!").Id("ok")).Block(Return(False(), True()))
			} else if s.Name == "Xhtml" {
				g.List(Id("o"), Id("ok")).Op(":=").Id("other").Dot("").Call(Id(s.Name))
				g.If(Op("!").Id("ok")).Block(Return(False(), True()))
				g.Return(Id("r.Value").Op("==").Id("o.Value"), True())
				return
			} else {
				g.List(Id("o"), Id("ok")).Op(":=").Id("other").Dot("").Call(Id(s.Name))
				g.If(Op("!").Id("ok")).Block(Return(False(), True()))
				g.If(Id("r.Value").Op("==").Nil().Op("||").Id("o.Value").Op("==").Nil()).
					Block(Return(False(), True()))

				g.Return(Id("*r.Value").Op("==").Id("*o.Value"), True())
				return
			}

			g.Return().Id("a").Dot("Equal").Call(Id("b"))
		} else {
			g.Var().Id("o").Op("*").Id(s.Name)
			g.Switch(Id("other").Op(":=").Id("other").Dot("(type)")).Block(
				Case(Id(s.Name)).Block(
					Id("o").Op("=").Op("&").Id("other"),
				),
				Case(Op("*").Id(s.Name)).Block(
					Id("o").Op("=").Id("other"),
				),
				Default().Block(
					Return(False(), True()),
				),
			)
			g.If(Id("o").Op("==").Nil()).Block(
				Return(False(), True()),
			)
			g.List(Id("eq"), Id("ok")).Op(":=").Id("r").Dot("Children").Call().Dot("Equal").Call(
				Id("o").Dot("Children").Call(),
			)
			g.Return(Id("eq").Op("&&").Id("ok"), True())
		}
	})
}

func generateEquivalentFunc(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("Equivalent").Params(
		Id("other").Qual(fhirpathModuleName, "Element"),
		Id("_noReverseTypeConversion").Op("...").Bool(),
	).Params(Bool()).Block(
		List(Id("eq"), Id("ok")).Op(":=").Id("r").Dot("Equal").Call(Id("other")),
		Return(Id("eq").Op("&&").Id("ok")),
	)
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
