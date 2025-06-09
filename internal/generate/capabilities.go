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
	generateCapability(f("delete", "capabilities"+release), ir.FilterResources(rt), release, "delete")
	generateCapability(f("search", "capabilities"+release), ir.FilterResources(rt), release, "search")

	generateAllCapabilitiesFn(f("capabilities", "capabilities"+release), release, ir.FilterResources(rt))
}

func generateCapability(f *File, resources []ir.ResourceOrType, release, interaction string) {
	interactionName := strcase.ToCamel(interaction)

	for _, r := range resources {
		f.Commentf("// %s%s needs to be implemented to support the %s interaction.", r.Name, interactionName, interaction)
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
						Error(),
					)
			case "read":
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("id").String(),
					).
					Params(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
						Error(),
					)
			case "update":
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("resource").Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
					).
					Params(
						Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)),
						Error(),
					)
			case "delete":
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("id").String(),
					).
					Params(
						Error(),
					)
			case "search":
				g.Id("SearchCapabilities"+r.Name).Params(Id("ctx").Qual("context", "Context")).Params(
					searchCapabilitiesType(release),
					Error(),
				)
				g.Id(interactionName+r.Name).
					Params(
						Id("ctx").Qual("context", "Context"),
						Id("options").Qual(moduleName+"/capabilities/search", "Options"),
					).
					Params(
						Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)),
						Error(),
					)
			}
		})

		if interaction == "update" {
			f.Commentf("// %sUpdateCapabilities is optional and only needs to be implemented if the backend deviates from the default behavior.", r.Name)
			f.Type().Id(r.Name + "UpdateCapabilities").Interface(
				Id("UpdateCapabilities"+r.Name).Params(Id("ctx").Qual("context", "Context")).Params(
					Qual(moduleName+"/capabilities/update", "Capabilities"),
					Error(),
				),
			)
		}
	}
}

func generateAllCapabilitiesFn(f *File, release string, resources []ir.ResourceOrType) {
	f.Func().Id("AllCapabilities").
		Params(
			Id("ctx").Qual("context", "Context"),
			Id("api").Any(),
		).
		Params(
			Qual(moduleName+"/capabilities", "Capabilities"),
			Error(),
		).BlockFunc(func(g *Group) {
		g.Id("allCapabilities").Op(":=").Qual(moduleName+"/capabilities", "Capabilities").Values(Dict{
			Id("Create"): Make(Map(String()).Qual(moduleName+"/capabilities/create", "Capabilities")),
			Id("Read"):   Make(Map(String()).Qual(moduleName+"/capabilities/read", "Capabilities")),
			Id("Update"): Make(Map(String()).Qual(moduleName+"/capabilities/update", "Capabilities")),
			Id("Delete"): Make(Map(String()).Qual(moduleName+"/capabilities/deletion", "Capabilities")),
			Id("Search"): Make(Map(String()).Qual(moduleName+"/capabilities/search", "Capabilities").Index(Qual(moduleName+"/capabilities/search", "Parameter"))),
		})
		g.Var().Id("errs").Index().Error()

		for _, r := range resources {
			for _, interaction := range []string{"create", "read", "delete"} {
				interactionName := strcase.ToCamel(interaction)
				capPkg := interaction
				if interaction == "delete" {
					capPkg = "deletion"
				}

				g.If(
					List(Id("_"), Id("ok")).Op(":=").Id("api").Dot("").Call(
						Id(r.Name+interactionName),
					),
					Id("ok"),
				).Block(
					Id("allCapabilities."+interactionName).Index(Lit(r.Name)).Op("=").Qual(moduleName+"/capabilities/"+capPkg, "Capabilities").Values(),
				)
			}

			// Update
			g.If(
				List(Id("c"), Id("ok")).Op(":=").Id("api").Dot("").Call(
					Id(r.Name+"Update"),
				),
				Id("ok"),
			).Block(
				Id("allCapabilities.Update").Index(Lit(r.Name)).Op("=").Qual(moduleName+"/capabilities/update", "Capabilities").Values(),
				List(Id("c"), Id("ok")).Op(":=").Id("c").Dot("").Call(Id(r.Name+"UpdateCapabilities")),
				If(Id("ok")).Block(
					List(Id("c"), Err()).Op(":=").Id("c").Dot("UpdateCapabilities"+r.Name).Call(Id("ctx")),
					If(Err().Op("!=").Nil()).Block(
						Id("errs").Op("=").Append(Id("errs"), Err()),
					).Else().Block(
						Id("allCapabilities.Update").Index(Lit(r.Name)).Op("=").Id("c"),
					),
				),
			)

			// Search
			g.If(
				List(Id("c"), Id("ok")).Op(":=").Id("api").Dot("").Call(
					Id(r.Name+"Search"),
				),
				Id("ok"),
			).Block(
				List(Id("c"), Err()).Op(":=").Id("c").Dot("SearchCapabilities"+r.Name).Call(Id("ctx")),
				If(Err().Op("!=").Nil()).Block(
					Id("errs").Op("=").Append(Id("errs"), Err()),
				).Else().Block(
					// Assign a new capabilities struct, initializing the Parameters map
					Id("allCapabilities").Dot("Search").Index(Lit(r.Name)).Op("=").Qual(
						moduleName+"/capabilities/search", "Capabilities",
					).Index(
						Qual(moduleName+"/capabilities/search", "Parameter"),
					).Values(Dict{
						Id("Includes"): Id("c").Dot("Includes"),
						Id("Parameters"): Make(
							Map(String()).Qual(moduleName+"/capabilities/search", "Parameter"),
						),
					}),
					// Iterate over the source parameters and copy them
					For(
						List(Id("n"), Id("p")).Op(":=").Range().Id("c").Dot("Parameters"),
					).Block(
						Id("allCapabilities").Dot("Search").Index(Lit(r.Name)).Dot("Parameters").Index(Id("n")).Op("=").Id("p"),
					),
				),
			)
		}

		g.Return(Id("allCapabilities"), Qual("errors", "Join").Call(Id("errs").Op("...")))
	})
}
