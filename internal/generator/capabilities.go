package generator

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generator/ir"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

type CapabilitiesGenerator struct {
	NoOpGenerator
}

func (g CapabilitiesGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	generateCapability(f("read", "capabilities"+release), ir.FilterResources(rt), release, "read", readParams, readReturn)
	generateCapability(f("search", "capabilities"+release), ir.FilterResources(rt), release, "search", searchParams, searchReturn)

	generateAllCapabilitiesFn(f("capabilities", "capabilities"+release), release, ir.FilterResources(rt))
}

var (
	readParams = map[Code]Code{
		Id("id"): String(),
	}
	searchParams = map[Code]Code{
		Id("options"): Qual(moduleName+"/capabilities/search", "Options"),
	}
	searchCapabilitiesReturn = Qual(moduleName+"/capabilities/search", "Capabilities")
)

type returnTypeFunc = func(typeName, releaseLower string) *Statement

func readReturn(typeName, release string) *Statement {
	return Qual(moduleName+"/model/gen/"+strings.ToLower(release), typeName)
}

func searchReturn(_, _ string) *Statement {
	return Qual(moduleName+"/capabilities/search", "Result")
}

func generateCapability(f *File, resources []ir.ResourceOrType, release, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)

	allParams := []Code{Id("ctx").Qual("context", "Context")}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	for _, r := range resources {
		f.Type().Id(r.Name + interactionName).InterfaceFunc(func(g *Group) {
			if interaction == "search" {
				g.Id("SearchCapabilities" + r.Name).Params().Params(searchCapabilitiesReturn)
			}

			g.Id(interactionName+r.Name).Params(allParams...).Params(returnFunc(r.Name, release), Qual(moduleName+"/capabilities", "FHIRError"))
		})
	}
}

func generateAllCapabilitiesFn(f *File, release string, resources []ir.ResourceOrType) {
	f.Func().Id("AllCapabilities").Params(Id("api").Any()).Params(Qual(moduleName+"/capabilities", "Capabilities")).BlockFunc(func(g *Group) {
		g.Id("read").Op(":=").Index().String().Values()
		g.Id("search").Op(":=").Map(String()).Qual(moduleName+"/capabilities/search", "Capabilities").Values()

		for _, r := range resources {
			g.If(
				List(Id("_"), Id("ok")).Op(":=").Id("api").Dot("").Call(
					Id(r.Name+"Read"),
				),
				Id("ok"),
			).Block(
				Id("read").Op("=").Append(List(Id("read"), Lit(r.Name))),
			)

			g.If(
				List(Id("c"), Id("ok")).Op(":=").Id("api").Dot("").Call(
					Id(r.Name+"Search"),
				),
				Id("ok"),
			).Block(
				Id("search").Index(Lit(r.Name)).Op("=").Id("c").Dot("SearchCapabilities" + r.Name).Call(),
			)
		}

		g.Return(Qual(moduleName+"/capabilities", "Capabilities").Values(Dict{
			Id("ReadInteractions"):   Id("read"),
			Id("SearchCapabilities"): Id("search"),
		}))
	})
}
