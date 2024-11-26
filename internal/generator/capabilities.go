package generator

import (
	"fhir-toolbox/internal/generator/ir"
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
	generateFull(f("full", "capabilities"+release), ir.FilterResources(rt))
}

var (
	readParams = map[Code]Code{
		Id("id"): String(),
	}
	searchParams = map[Code]Code{
		Id("options"): Qual("fhir-toolbox/capabilities/search", "Options"),
	}
	searchCapabilitiesReturn = Qual("fhir-toolbox/capabilities/search", "Capabilities")
)

type returnTypeFunc = func(typeName, releaseLower string) *Statement

func readReturn(typeName, release string) *Statement {
	return Qual("fhir-toolbox/model/gen/"+strings.ToLower(release), typeName)
}

func searchReturn(_, _ string) *Statement {
	return Qual("fhir-toolbox/capabilities/search", "Result")
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

			g.Id(interactionName+r.Name).Params(allParams...).Params(returnFunc(r.Name, release), Qual("fhir-toolbox/capabilities", "FHIRError"))
		})
	}
}

func generateFull(f *File, resources []ir.ResourceOrType) {
	f.Type().Id("FullAPI").InterfaceFunc(func(g *Group) {
		g.Qual("fhir-toolbox/capabilities", "GenericAPI")
		for _, r := range resources {
			g.Id(r.Name + "Read")
			g.Id(r.Name + "Search")
		}
	})
}
