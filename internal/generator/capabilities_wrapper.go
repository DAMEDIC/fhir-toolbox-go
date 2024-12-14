package generator

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generator/ir"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

const genericWrapperName = "Generic"
const concreteWrapperName = "Concrete"

type CapabilitiesWrapperGenerator struct {
	NoOpGenerator
}

func (g CapabilitiesWrapperGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	generateGenericWrapperStruct(f("generic", "capabilities"+release))
	generateWrapperAllCapabilities(f("generic", "capabilities"+release), release)
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "read", readParams, anyReadReturn)
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "search", searchParams, searchReturn)

	generateConcreteIface(f("concrete", "capabilities"+release), ir.FilterResources(rt))
	generateConcreteWrapperStruct(f("concrete", "capabilities"+release))
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "read", readParams, anyReadReturn)
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "search", searchParams, searchReturn)
}

func generateGenericWrapperStruct(f *File) {
	f.Type().Id(genericWrapperName).Struct(
		Id("Concrete").Any(),
	)
}

func anyReadReturn(_, _ string) *Statement {
	return Qual(moduleName+"/model", "Resource")
}

func generateGeneric(f *File, release string, resources []ir.ResourceOrType, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)

	allParams := []Code{Id("ctx").Qual("context", "Context"), Id("resourceType").String()}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	passParams := []Code{Id("ctx")}
	for k := range params {
		passParams = append(passParams, k)
	}

	if interaction == "search" {
		f.Add(generateGenericSearchCapabilities(release, resources))
	}

	f.Func().Params(Id("w").Id(genericWrapperName)).Id(interactionName).
		Params(allParams...).
		Params(returnFunc("", ""), Qual(moduleName+"/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
						g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))

						if interaction == "search" {
							g.If(Op("!").Id("ok")).Block(returnNotImplementedError(interactionName, r.Name, returnFunc("", "").Block()))
						} else {
							g.If(Op("!").Id("ok")).Block(returnNotImplementedError(interactionName, r.Name, Nil()))
						}

						g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
					})
				}

				if interaction == "search" {
					g.Default().Block(returnUnknownError("resourceType", returnFunc("", "").Block()))
				} else {
					g.Default().Block(returnUnknownError("resourceType", Nil()))
				}
			}),
		)
}

func generateGenericSearchCapabilities(release string, resources []ir.ResourceOrType) Code {
	return Func().Params(Id("w").Id(genericWrapperName)).Id("SearchCapabilities").
		Params(Id("resourceType").String()).
		Params(searchCapabilitiesReturn, Qual(moduleName+"/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name+"Search")),
						If(Op("!").Id("ok")).Block(
							returnNotImplementedError("search", r.Name, searchCapabilitiesReturn.Clone().Block()),
						),
						Return(Id("impl.SearchCapabilities"+r.Name).Call(), Nil()),
					)
				}

				g.Default().Block(returnUnknownError("resourceType", searchCapabilitiesReturn.Clone().Block()))
			}),
		)
}

func generateWrapperAllCapabilities(f *File, release string) {
	f.Func().Params(Id("w").Id(genericWrapperName)).Id("AllCapabilities").Params().
		Params(Qual(moduleName+"/capabilities", "Capabilities")).
		Block(
			Return(Id("AllCapabilities").
				Call(Id("w").Dot("Concrete"))),
		)
}

func generateConcreteIface(f *File, resources []ir.ResourceOrType) {
	f.Type().Id("ConcreteAPI").InterfaceFunc(func(g *Group) {
		for _, r := range resources {
			g.Id(r.Name + "Read")
			g.Id(r.Name + "Search")
		}
	})
}

func generateConcreteWrapperStruct(f *File) {
	f.Type().Id(concreteWrapperName).Struct(
		Id("Generic").Qual(moduleName+"/capabilities", "GenericAPI"),
	)
}

func generateConcrete(f *File, release string, resources []ir.ResourceOrType, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)

	allParams := []Code{Id("ctx").Qual("context", "Context")}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	for _, r := range resources {
		passParams := []Code{Id("ctx"), Lit(r.Name)}
		for k := range params {
			passParams = append(passParams, k)
		}

		var returnTypeId *Statement
		if interaction == "search" {
			returnTypeId = Qual(moduleName+"/capabilities/search", "Result")
		} else {
			returnTypeId = Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)
		}

		if interaction == "search" {
			f.Add(generateConcreteSearchCapabilities(r))
		}

		f.Func().Params(Id("w").Id(concreteWrapperName)).Id(interactionName+r.Name).
			Params(allParams...).
			Params(returnTypeId, Qual(moduleName+"/capabilities", "FHIRError")).
			BlockFunc(func(g *Group) {
				g.List(Id("v"), Id("err")).Op(":=").Id("w.Generic." + interactionName).Params(passParams...)
				g.If(Id("err").Op("!=").Nil()).Block(Return(returnTypeId.Clone().Block(), Id("err")))

				if interaction == "read" {
					g.List(Id("contained"), Id("ok")).Op(":=").Id("v").Assert(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), "ContainedResource"),
					)
					g.If(Id("ok")).Block(
						Id("v").Op("=").Id("contained.Resource"),
					)
					g.List(Id("impl"), Id("ok")).Op(":=").Id("v").Assert(returnTypeId)
					g.If(Id("!ok")).Block(
						returnInvalidError(r.Name, returnTypeId.Clone().Block()),
					)
					g.Return(Id("impl"), Nil())
				} else {
					g.Return(Id("v"), Nil())
				}
			})
	}
}

func generateConcreteSearchCapabilities(r ir.ResourceOrType) Code {
	returnId := Qual(moduleName+"/capabilities/search", "Capabilities")
	return Func().Params(Id("w").Id(concreteWrapperName)).Id("SearchCapabilities" + r.Name).
		Params().
		Params(returnId).
		BlockFunc(func(g *Group) {
			g.List(Id("c"), Id("err")).Op(":=").Id("w.Generic.SearchCapabilities").Params(Lit(r.Name))
			g.If(Id("err").Op("!=").Nil()).Block(Return(returnId.Clone().Block()))
			g.Return(Id("c"))
		})
}

func returnNotImplementedError(interaction, resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual(moduleName+"/capabilities", "NotImplementedError").Values(
		Id("Interaction").Op(":").Lit(strings.ToLower(interaction)),
		Id("ResourceType").Op(":").Lit(resourceType)),
	)
}

func returnUnknownError(resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual(moduleName+"/capabilities", "UnknownResourceError").Values(
		Id("ResourceType").Op(":").Id(resourceType)),
	)
}

func returnInvalidError(resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual(moduleName+"/capabilities", "InvalidResourceError").Values(
		Id("ResourceType").Op(":").Lit(resourceType)),
	)
}
