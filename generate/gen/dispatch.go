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
	generateDispatch(dir, release, resources, "search", searchParams, anySearchReturn)

}

func anyReadReturn(_, _ string) Code {
	return Qual("fhir-toolbox/model", "Resource")
}

func anySearchReturn(_, _ string) Code {
	return Index().Qual("fhir-toolbox/model", "Resource")
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
						g.If(Op("!").Id("ok")).Block(returnNotImplementedError(interactionName, r.Name, Nil()))

						if interaction == "search" {
							g.List(Id("v"), Id("err")).Op(":=").Id("impl." + interactionName + r.Name).Call(passParams...)
							g.If(Id("err").Op("!=").Nil()).Block(Return(Nil(), Id("err")))
							g.Id("r").Op(":=").Make(Index().Qual("fhir-toolbox/model", "Resource"), Lit(0), Len(Id("v")))
							g.For(List(Id("_"), Id("e")).Op(":=").Range().Id("v")).Block(
								Id("r").Op("=").Append(Id("r"), Id("e")),
							)
							g.Return(Id("r"), Nil())
						} else {
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						}
					})
				}

				g.Default().Block(returnUnknownError("resourceType", Nil()))
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
							If(Op("!").Id("ok")).Block(returnNotImplementedError("SearchParams", r.Name, searchCapabilitiesReturn.Clone().Block())),
						),
						Return(Id("impl.SearchCapabilities"+r.Name).Call(), Nil()),
					)
				}

				g.Default().Block(returnUnknownError("resourceType", searchCapabilitiesReturn.Clone().Block()))
			}),
		)
}

func returnNotImplementedError(interaction, resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual("fhir-toolbox/dispatch/errors", "NotImplementedError").Values(
		Id("Interaction").Op(":").Lit(interaction),
		Id("ResourceType").Op(":").Lit(resourceType)),
	)
}

func returnUnknownError(resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual("fhir-toolbox/dispatch/errors", "UnknownResourceError").Values(
		Id("ResourceType").Op(":").Id(resourceType)),
	)
}
