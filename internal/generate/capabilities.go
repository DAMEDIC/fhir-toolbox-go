package generate

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

type CapabilitiesGenerator struct {
	NoOpGenerator
}

func (g CapabilitiesGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	generateCapability(f("read", "capabilities"+release), ir.FilterResources(rt), release, "read")
	generateCapability(f("search", "capabilities"+release), ir.FilterResources(rt), release, "search")

	generateAllCapabilitiesFn(f("capabilities", "capabilities"+release), release, ir.FilterResources(rt))
}

func generateCapability(f *File, resources []ir.ResourceOrType, release, interaction string) {
	interactionName := strcase.ToCamel(interaction)

	params := []Code{Id("ctx").Qual("context", "Context")}
	switch interaction {
	case "read":
		params = append(params, Id("id").String())
	case "search":
		params = append(params, Id("options").Qual(moduleName+"/capabilities/search", "Options"))
	}

	for _, r := range resources {
		f.Type().Id(r.Name + interactionName).InterfaceFunc(func(g *Group) {
			var returnType Code
			switch interaction {
			case "read":
				returnType = Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)
			case "search":
				returnType = Qual(moduleName+"/capabilities/search", "Result")
				g.Id("SearchCapabilities"+r.Name).Params(Id("ctx").Qual("context", "Context")).Params(
					Qual(moduleName+"/capabilities/search", "Capabilities"),
					Qual(moduleName+"/capabilities", "FHIRError"),
				)
			}

			g.Id(interactionName+r.Name).Params(params...).Params(returnType, Qual(moduleName+"/capabilities", "FHIRError"))
		})
	}
}

func generateAllCapabilitiesFn(f *File, release string, resources []ir.ResourceOrType) {
	f.Func().Id("AllCapabilities").Params(Id("ctx").Qual("context", "Context"), Id("api").Any()).Params(
		Qual(moduleName+"/capabilities", "Capabilities"),
		Qual(moduleName+"/capabilities", "FHIRError"),
	).BlockFunc(func(g *Group) {
		g.Var().Id("errs").Index().Qual(moduleName+"/capabilities", "FHIRError")
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
				List(Id("capability"), Err()).Op(":=").Id("c").Dot("SearchCapabilities"+r.Name).Call(Id("ctx")),
				If(Err().Op("!=").Nil()).Block(
					Id("errs").Op("=").Append(List(Id("errs"), Err())),
				).Else().Block(
					Id("search").Index(Lit(r.Name)).Op("=").Id("capability"),
				),
			)
		}

		g.Return(
			Qual(moduleName+"/capabilities", "Capabilities").Values(Dict{
				Id("ReadInteractions"):   Id("read"),
				Id("SearchCapabilities"): Id("search"),
			}),
			Qual(moduleName+"/capabilities", "JoinErrors").Call(Id("errs")),
		)
	})
}
