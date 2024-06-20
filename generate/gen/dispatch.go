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

func GenerateServerDispatch(resources []ir.Struct, genTarget, release string) {
	dir := filepath.Join(genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	generateDispatch(dir, release, resources, "read", readParams, anyReadReturn)
	generateDispatch(dir, release, resources, "search", searchParams, searchReturn)

}

func anyReadReturn(_, _ string) *Statement {
	return Qual("fhir-toolbox/model", "Resource")
}

func generateDispatch(genDir, release string, resources []ir.Struct, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)
	fileName := strings.ToLower(interaction) + ".go"

	allParams := []Code{Id("ctx").Qual("context", "Context"), Id("api").Id("any"), Id("resourceType").String()}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	passParams := []Code{Id("ctx")}
	for k := range params {
		passParams = append(passParams, k)
	}

	f := NewFilePathName(genDir, "dispatch"+strings.ToUpper(release))

	if interaction == "search" {
		f.Add(generateSearchCapabilities(interactionName, release, resources))
	}

	f.Func().Id(interactionName).
		Params(allParams...).
		Params(returnFunc("", ""), Qual("fhir-toolbox/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
						g.List(Id("impl"), Id("ok")).Op(":=").Id("api").Assert(Qual("fhir-toolbox/capabilities/gen/"+strings.ToLower(release), r.Name+interactionName))

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

	err := f.Save(filepath.Join(genDir, fileName))
	if err != nil {
		log.Panic(err)
	}
}

func generateSearchCapabilities(interactionName, release string, resources []ir.Struct) Code {
	return Func().Id("SearchCapabilities").
		Params(Id("api").Id("any"), Id("resourceType").String()).
		Params(searchCapabilitiesReturn, Qual("fhir-toolbox/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("api").Assert(Qual("fhir-toolbox/capabilities/gen/"+strings.ToLower(release), r.Name+interactionName)),
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

func returnNotImplementedError(interaction, resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual("fhir-toolbox/capabilities", "NotImplementedError").Values(
		Id("Interaction").Op(":").Lit(strings.ToLower(interaction)),
		Id("ResourceType").Op(":").Lit(resourceType)),
	)
}

func returnUnknownError(resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual("fhir-toolbox/capabilities", "UnknownResourceError").Values(
		Id("ResourceType").Op(":").Id(resourceType)),
	)
}
