package fhirpath

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	. "github.com/dave/jennifer/jen"
)

func generateTypes(f *File, rt []ir.ResourceOrType) {
	f.Var().Id("allFHIRPathTypes").Op("=").Index().Qual(fhirpathModuleName, "TypeInfo").
		ValuesFunc(func(g *Group) {
			for _, t := range baseTypes {
				g.Add(t)
			}

			for _, t := range rt {
				generateType(g, t.Structs[0])
			}
		})
}

func generateType(g *Group, s ir.Struct) {
	var base *Statement
	if s.IsResource {
		base = Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("DomainResource"),
		})
	} else if s.IsPrimitive {
		base = Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("PrimitiveType"),
		})
	} else {
		base = Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("DataType"),
		})
	}

	elements := Index().Qual(fhirpathModuleName, "ClassInfoElement").ValuesFunc(func(g *Group) {
		for _, f := range s.Fields {
			if s.IsPrimitive && f.Name == "Value" {
				continue
			}

			var t string
			if f.Polymorph {
				t = "FHIR.PrimitiveElement"
			} else {
				t = "FHIR." + f.PossibleTypes[0].Name
			}
			if f.Multiple {
				t = fmt.Sprintf("List<%s>", t)
			}
			g.Values(Dict{
				Id("Name"): Lit(f.Name),
				Id("Type"): Lit(t),
			})
		}
	})

	g.Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit(s.Name),
			Id("BaseType"):  base,
		}),
		Id("Element"): elements,
	})
}

var baseTypes = []Code{
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("Base"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("System"),
				Id("Name"):      Lit("Any"),
			}),
		}),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("Element"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("Base"),
			}),
		}),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("DataType"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("Element"),
			}),
		}),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("PrimitiveType"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("DataType"),
			}),
		}),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("BackboneElement"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("Element"),
			}),
		}),
		Id("Element"): Index().Qual(fhirpathModuleName, "ClassInfoElement").Values(Values(Dict{
			Id("Name"): Lit("modifierExtension"),
			Id("Type"): Lit("List<FHIR.Extension>"),
		})),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("BackboneType"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("DataType"),
			}),
		}),
		Id("Element"): Index().Qual(fhirpathModuleName, "ClassInfoElement").Values(Values(Dict{
			Id("Name"): Lit("modifierExtension"),
			Id("Type"): Lit("List<FHIR.Extension>"),
		})),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("Resource"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("Base"),
			}),
		}),
		Id("Element"): Index().Qual(fhirpathModuleName, "ClassInfoElement").Values(
			Values(Dict{
				Id("Name"): Lit("id"),
				Id("Type"): Lit("FHIR.id"),
			}),
			Values(Dict{
				Id("Name"): Lit("meta"),
				Id("Type"): Lit("FHIR.Meta"),
			}),
			Values(Dict{
				Id("Name"): Lit("implicitRules"),
				Id("Type"): Lit("FHIR.uri"),
			}),
			Values(Dict{
				Id("Name"): Lit("language"),
				Id("Type"): Lit("FHIR.code"),
			}),
		),
	}),
	Qual(fhirpathModuleName, "ClassInfo").Values(Dict{
		Id("SimpleTypeInfo"): Qual(fhirpathModuleName, "SimpleTypeInfo").Values(Dict{
			Id("Namespace"): Lit("FHIR"),
			Id("Name"):      Lit("DomainResource"),
			Id("BaseType"): Qual(fhirpathModuleName, "TypeSpecifier").Values(Dict{
				Id("Namespace"): Lit("FHIR"),
				Id("Name"):      Lit("Resource"),
			}),
		}),
		Id("Element"): Index().Qual(fhirpathModuleName, "ClassInfoElement").Values(
			Values(Dict{
				Id("Name"): Lit("id"),
				Id("Type"): Lit("FHIR.id"),
			}),
			Values(Dict{
				Id("Name"): Lit("meta"),
				Id("Type"): Lit("FHIR.Meta"),
			}),
			Values(Dict{
				Id("Name"): Lit("implicitRules"),
				Id("Type"): Lit("FHIR.uri"),
			}),
			Values(Dict{
				Id("Name"): Lit("language"),
				Id("Type"): Lit("FHIR.code"),
			}),
			Values(Dict{
				Id("Name"): Lit("text"),
				Id("Type"): Lit("FHIR.Narrative"),
			}),
			Values(Dict{
				Id("Name"): Lit("contained"),
				Id("Type"): Lit("List<FHIR.Resource>"),
			}),
			Values(Dict{
				Id("Name"): Lit("extension"),
				Id("Type"): Lit("List<FHIR.Extension>"),
			}),
			Values(Dict{
				Id("Name"): Lit("modifierExtension"),
				Id("Type"): Lit("List<FHIR.Extension>"),
			}),
		),
	}),
}
