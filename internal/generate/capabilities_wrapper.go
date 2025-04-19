package generate

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
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
	generateWrapperAllCapabilities(f("generic", "capabilities"+release))
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "search")

	generateConcreteIface(f("concrete", "capabilities"+release), ir.FilterResources(rt))
	generateConcreteWrapperStruct(f("concrete", "capabilities"+release))
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "search")
}

func generateGenericWrapperStruct(f *File) {
	f.Type().Id(genericWrapperName).Struct(
		Id("Concrete").Any(),
	)
}

func generateGeneric(f *File, release string, resources []ir.ResourceOrType, interaction string) {
	interactionName := strcase.ToCamel(interaction)

	params := []Code{Id("ctx").Qual("context", "Context")}
	passParams := []Code{Id("ctx")}
	var returnType *Statement
	switch interaction {
	case "read":
		params = append(params, Id("resourceType").String(), Id("id").String())
		passParams = append(passParams, Id("id"))
		returnType = Qual(moduleName+"/model", "Resource")
	case "search":
		params = append(params, Id("resourceType").String(), Id("options").Qual(moduleName+"/capabilities/search", "Options"))
		passParams = append(passParams, Id("options"))
		returnType = Qual(moduleName+"/capabilities/search", "Result")
		f.Add(generateGenericSearchCapabilities(resources))
	}

	f.Func().Params(Id("w").Id(genericWrapperName)).Id(interactionName).
		Params(params...).
		Params(returnType, Qual(moduleName+"/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
						g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))

						if interaction == "search" {
							g.If(Op("!").Id("ok")).Block(returnNotImplementedError(interactionName, r.Name, returnType.Clone().Block()))
						} else {
							g.If(Op("!").Id("ok")).Block(returnNotImplementedError(interactionName, r.Name, Nil()))
						}

						g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
					})
				}

				if interaction == "search" {
					g.Default().Block(returnUnknownError("resourceType", returnType.Clone().Block()))
				} else {
					g.Default().Block(returnUnknownError("resourceType", Nil()))
				}
			}),
		)
}

func generateGenericSearchCapabilities(resources []ir.ResourceOrType) Code {
	return Func().Params(Id("w").Id(genericWrapperName)).Id("SearchCapabilities").
		Params(Id("resourceType").String()).
		Params(Qual(moduleName+"/capabilities/search", "Capabilities"), Qual(moduleName+"/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name+"Search")),
						If(Op("!").Id("ok")).Block(
							returnNotImplementedError("search", r.Name, Qual(moduleName+"/capabilities/search", "Capabilities").Block()),
						),
						Return(Id("impl.SearchCapabilities"+r.Name).Call()),
					)
				}

				g.Default().Block(returnUnknownError("resourceType", Qual(moduleName+"/capabilities/search", "Capabilities").Block()))
			}),
		)
}

func generateWrapperAllCapabilities(f *File) {
	f.Func().Params(Id("w").Id(genericWrapperName)).Id("AllCapabilities").Params().
		Params(
			Qual(moduleName+"/capabilities", "Capabilities"),
			Qual(moduleName+"/capabilities", "FHIRError"),
		).
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

func generateConcrete(f *File, release string, resources []ir.ResourceOrType, interaction string) {
	interactionName := strcase.ToCamel(interaction)

	params := []Code{Id("ctx").Qual("context", "Context")}
	switch interaction {
	case "read":
		params = append(params, Id("id").String())
	case "search":
		params = append(params, Id("options").Qual(moduleName+"/capabilities/search", "Options"))
	}

	for _, r := range resources {
		passParams := []Code{Id("ctx"), Lit(r.Name)}
		var returnType *Statement
		switch interaction {
		case "read":
			passParams = append(passParams, Id("id"))
			returnType = Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)
		case "search":
			passParams = append(passParams, Id("options"))
			returnType = Qual(moduleName+"/capabilities/search", "Result")
			f.Add(generateConcreteSearchCapabilities(r))
		}

		f.Func().Params(Id("w").Id(concreteWrapperName)).Id(interactionName+r.Name).
			Params(params...).
			Params(returnType, Qual(moduleName+"/capabilities", "FHIRError")).
			BlockFunc(func(g *Group) {
				g.List(Id("v"), Id("err")).Op(":=").Id("w.Generic." + interactionName).Params(passParams...)
				g.If(Id("err").Op("!=").Nil()).Block(Return(returnType.Clone().Block(), Id("err")))

				if interaction == "read" {
					g.List(Id("contained"), Id("ok")).Op(":=").Id("v").Assert(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), "ContainedResource"),
					)
					g.If(Id("ok")).Block(
						Id("v").Op("=").Id("contained.Resource"),
					)
					g.List(Id("impl"), Id("ok")).Op(":=").Id("v").Assert(returnType)
					g.If(Id("!ok")).Block(
						returnInvalidError(r.Name, returnType.Clone().Block()),
					)
					g.Return(Id("impl"), Nil())
				} else {
					g.Return(Id("v"), Nil())
				}
			})
	}
}

func generateConcreteSearchCapabilities(r ir.ResourceOrType) Code {
	returnType := Qual(moduleName+"/capabilities/search", "Capabilities")
	return Func().Params(Id("w").Id(concreteWrapperName)).Id("SearchCapabilities"+r.Name).
		Params().
		Params(returnType, Qual(moduleName+"/capabilities", "FHIRError")).
		BlockFunc(func(g *Group) {
			g.Return(Id("w.Generic.SearchCapabilities").Params(Lit(r.Name)))
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
