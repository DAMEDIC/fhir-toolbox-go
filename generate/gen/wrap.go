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

const wrapperName = "InternalWrapper"

func GenerateWrapper(resources []ir.Struct, genTarget, release string) {
	genericDir := filepath.Join(genTarget, strings.ToLower(release), "generic")
	concreteDir := filepath.Join(genTarget, strings.ToLower(release), "concrete")

	err := os.MkdirAll(genericDir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	err = os.MkdirAll(concreteDir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	generateGenericWrapperStruct(genericDir, release)
	generateGeneric(genericDir, release, resources, "read", readParams, anyReadReturn)
	generateGeneric(genericDir, release, resources, "search", searchParams, searchReturn)
	generateConcreteWrapperStruct(concreteDir, release)
	generateConcrete(concreteDir, release, resources, "read", readParams, anyReadReturn)
	generateConcrete(concreteDir, release, resources, "search", searchParams, searchReturn)
}

func generateGenericWrapperStruct(genDir, release string) {
	f := NewFilePathName(genDir, "generic"+strings.ToUpper(release))
	f.Type().Id(wrapperName).Struct(
		Id("API").Any(),
	)

	err := f.Save(filepath.Join(genDir, "wrapper.go"))
	if err != nil {
		log.Panic(err)
	}
}

func anyReadReturn(_, _ string) *Statement {
	return Qual("fhir-toolbox/model", "Resource")
}

func generateGeneric(genDir, release string, resources []ir.Struct, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)
	fileName := strings.ToLower(interaction) + ".go"

	allParams := []Code{Id("ctx").Qual("context", "Context"), Id("resourceType").String()}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	passParams := []Code{Id("ctx")}
	for k := range params {
		passParams = append(passParams, k)
	}

	f := NewFilePathName(genDir, "generic"+strings.ToUpper(release))

	if interaction == "search" {
		f.Add(generateGenericSearchCapabilities(release, resources))
	}

	f.Func().Params(Id("w").Id(wrapperName)).Id(interactionName).
		Params(allParams...).
		Params(returnFunc("", ""), Qual("fhir-toolbox/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
						g.List(Id("impl"), Id("ok")).Op(":=").Id("w.API").Assert(Qual("fhir-toolbox/capabilities/gen/"+strings.ToLower(release), r.Name+interactionName))

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

func generateGenericSearchCapabilities(release string, resources []ir.Struct) Code {
	return Func().Params(Id("w").Id(wrapperName)).Id("SearchCapabilities").
		Params(Id("resourceType").String()).
		Params(searchCapabilitiesReturn, Qual("fhir-toolbox/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("w.API").Assert(Qual("fhir-toolbox/capabilities/gen/"+strings.ToLower(release), r.Name+"Search")),
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

func generateConcreteWrapperStruct(genDir, release string) {
	f := NewFilePathName(genDir, "concrete"+strings.ToUpper(release))
	f.Type().Id(wrapperName).Struct(
		Qual("fhir-toolbox/capabilities", "GenericAPI"),
	)

	err := f.Save(filepath.Join(genDir, "wrapper.go"))
	if err != nil {
		log.Panic(err)
	}
}

func generateConcrete(genDir, release string, resources []ir.Struct, interaction string, params map[Code]Code, returnFunc returnTypeFunc) {
	interactionName := strcase.ToCamel(interaction)
	fileName := strings.ToLower(interaction) + ".go"

	allParams := []Code{Id("ctx").Qual("context", "Context")}
	for k, v := range params {
		allParams = append(allParams, &Statement{k, v})
	}

	f := NewFilePathName(genDir, "concrete"+strings.ToUpper(release))

	for _, r := range resources {
		passParams := []Code{Id("ctx"), Lit(r.Name)}
		for k := range params {
			passParams = append(passParams, k)
		}

		var returnTypeId *Statement
		if interaction == "search" {
			returnTypeId = Qual("fhir-toolbox/capabilities/search", "Result")
		} else {
			returnTypeId = Qual("fhir-toolbox/model/gen/"+strings.ToLower(release), r.Name)
		}

		if interaction == "search" {
			f.Add(generateConcreteSearchCapabilities(release, r))
		}

		f.Func().Params(Id("w").Id(wrapperName)).Id(interactionName+r.Name).
			Params(allParams...).
			Params(returnTypeId, Qual("fhir-toolbox/capabilities", "FHIRError")).
			BlockFunc(func(g *Group) {
				g.List(Id("v"), Id("err")).Op(":=").Id("w.GenericAPI." + interactionName).Params(passParams...)
				g.If(Id("err").Op("!=").Nil()).Block(Return(returnTypeId.Clone().Block(), Id("err")))

				if interaction == "read" {
					g.List(Id("contained"), Id("ok")).Op(":=").Id("v").Assert(
						Qual("fhir-toolbox/model/gen/"+strings.ToLower(release), "ContainedResource"),
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

	err := f.Save(filepath.Join(genDir, fileName))
	if err != nil {
		log.Panic(err)
	}
}

func generateConcreteSearchCapabilities(release string, r ir.Struct) Code {
	returnId := Qual("fhir-toolbox/capabilities/search", "Capabilities")
	return Func().Params(Id("w").Id(wrapperName)).Id("SearchCapabilities" + r.Name).
		Params().
		Params(returnId).
		BlockFunc(func(g *Group) {
			g.List(Id("c"), Id("err")).Op(":=").Id("w.GenericAPI.SearchCapabilities").Params(Lit(r.Name))
			g.If(Id("err").Op("!=").Nil()).Block(Return(returnId.Clone().Block()))
			g.Return(Id("c"))
		})
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

func returnInvalidError(resourceType string, defaultReturn Code) Code {
	return Return(defaultReturn, Qual("fhir-toolbox/capabilities", "InvalidResourceError").Values(
		Id("ResourceType").Op(":").Lit(resourceType)),
	)
}
