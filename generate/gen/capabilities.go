package gen

import (
	"fhir-toolbox/generate/ir"
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func GenerateCapabilityInterfaces(resources []ir.Struct, genTarget, release string) {
	dir := filepath.Join(genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	generateCapability(dir, release, resources, "read", readParams, readReturn)
	generateCapability(dir, release, resources, "search", searchParams, searchReturn)

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

func generateCapability(genDir string, release string, resources []ir.Struct, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)
	fileName := strings.ToLower(interaction)

	allParams := []Code{Id("ctx").Qual("context", "Context")}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	f := NewFilePathName(genDir, "capabilities"+strings.ToUpper(release))

	for _, r := range resources {
		f.Type().Id(r.Name + interactionName).InterfaceFunc(func(g *Group) {
			if interaction == "search" {
				g.Id("SearchCapabilities" + r.Name).Params().Params(searchCapabilitiesReturn)
			}

			g.Id(interactionName+r.Name).Params(allParams...).Params(returnFunc(r.Name, release), Qual("fhir-toolbox/capabilities", "FHIRError"))
		})
	}

	err := f.Save(filepath.Join(genDir, fileName+".go"))
	if err != nil {
		log.Panic(err)
	}
}
