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
	generateCapability(f("create", "capabilities"+release), ir.FilterResources(rt), release, "create")
	generateCapability(f("read", "capabilities"+release), ir.FilterResources(rt), release, "read")
	generateCapability(f("update", "capabilities"+release), ir.FilterResources(rt), release, "update")
	generateCapability(f("search", "capabilities"+release), ir.FilterResources(rt), release, "search")

	generateAllCapabilitiesFn(f("capabilities", "capabilities"+release), release, ir.FilterResources(rt))
}

func generateCapability(f *File, resources []ir.ResourceOrType, release, interaction string) {
	interactionName := strcase.ToCamel(interaction)

	for _, r := range resources {
		f.Type().Id(r.Name + interactionName).InterfaceFunc(func(g *Group) {
			switch interaction {
			case "create":
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("resource").Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
					).
					Params(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
						Qual(moduleName+"/capabilities", "FHIRError"),
					)
			case "read":
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("id").String(),
					).
					Params(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
						Qual(moduleName+"/capabilities", "FHIRError"),
					)
			case "update":
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("resource").Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
					).
					Params(
						Qual(moduleName+"/capabilities", "UpdateResult").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)),
						Qual(moduleName+"/capabilities", "FHIRError"),
					)
			case "search":
				g.Id("SearchCapabilities"+r.Name).Params(Id("ctx").Qual("context", "Context")).Params(
					Qual(moduleName+"/capabilities/search", "Capabilities"),
					Qual(moduleName+"/capabilities", "FHIRError"),
				)
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("options").Qual(moduleName+"/capabilities/search", "Options"),
					).
					Params(
						Qual(moduleName+"/capabilities/search", "Result"),
						Qual(moduleName+"/capabilities", "FHIRError"),
					)
			}
		})
	}
}

func generateAllCapabilitiesFn(f *File, release string, resources []ir.ResourceOrType) {
	f.Func().Id("AllCapabilities").Params(Id("ctx").Qual("context", "Context"), Id("api").Any()).Params(
		Qual(moduleName+"/capabilities", "Capabilities"),
		Qual(moduleName+"/capabilities", "FHIRError"),
	).BlockFunc(func(g *Group) {
		g.Var().Id("errs").Index().Qual(moduleName+"/capabilities", "FHIRError")
		g.Id("create").Op(":=").Index().String().Values()
		g.Id("read").Op(":=").Index().String().Values()
		g.Id("search").Op(":=").Map(String()).Qual(moduleName+"/capabilities/search", "Capabilities").Values()

		for _, r := range resources {
			for _, interaction := range []string{"create", "read"} {
				interactionName := strcase.ToCamel(interaction)

				g.If(
					List(Id("_"), Id("ok")).Op(":=").Id("api").Dot("").Call(
						Id(r.Name+interactionName),
					),
					Id("ok"),
				).Block(
					Id(interaction).Op("=").Append(List(Id(interaction), Lit(r.Name))),
				)
			}

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
