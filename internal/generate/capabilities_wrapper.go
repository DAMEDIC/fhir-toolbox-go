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
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "create")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "update")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "search")

	generateConcreteWrapperStruct(f("concrete", "capabilities"+release))
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "create")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "update")
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
	shortcutParams := []Code{Id("ctx")}
	passParams := []Code{Id("ctx")}
	switchType := Id("resourceType")
	var returnType *Statement
	switch interaction {
	case "create":
		params = append(params, Id("resource").Qual(moduleName+"/model", "Resource"))
		shortcutParams = append(shortcutParams, Id("resource"))
		passParams = append(passParams, Id("r"))
		switchType = Id("r").Op(":=").Id("resource").Assert(Type())
		returnType = Qual(moduleName+"/model", "Resource")
	case "read":
		params = append(params, Id("resourceType").String(), Id("id").String())
		shortcutParams = append(shortcutParams, Id("resourceType"), Id("id"))
		passParams = append(passParams, Id("id"))
		returnType = Qual(moduleName+"/model", "Resource")
	case "update":
		params = append(params, Id("resource").Qual(moduleName+"/model", "Resource"))
		shortcutParams = append(shortcutParams, Id("resource"))
		passParams = append(passParams, Id("r"))
		switchType = Id("r").Op(":=").Id("resource").Assert(Type())
		returnType = Qual(moduleName+"/capabilities", "UpdateResult").Index(Qual(moduleName+"/model", "Resource"))
	case "search":
		params = append(params, Id("resourceType").String(), Id("options").Qual(moduleName+"/capabilities/search", "Options"))
		shortcutParams = append(shortcutParams, Id("resourceType"), Id("options"))
		passParams = append(passParams, Id("options"))
		returnType = Qual(moduleName+"/capabilities/search", "Result")
		f.Add(generateGenericSearchCapabilities(resources))
	}

	f.Func().Params(Id("w").Id(genericWrapperName)).Id(interactionName).
		Params(params...).
		Params(returnType, Qual(moduleName+"/capabilities", "FHIRError")).
		Block(
			List(Id("g"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Qual(moduleName+"/capabilities", "Generic"+interactionName)),
			If(Id("ok")).Block(
				Comment("// shortcut for the case that the underlying implementation already implements the generic API"),
				Return(Id("g."+interactionName).Params(shortcutParams...)),
			),

			Switch(switchType).BlockFunc(func(g *Group) {
				switch interaction {
				case "create":
					for _, r := range resources {
						g.Case(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(Return(Nil(), notImplementedError(interactionName, r.Name)))
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						})
					}

					g.Default().Block(Return(Nil(), unknownError(Id("resource").Dot("ResourceType").Call())))
				case "read":
					for _, r := range resources {
						g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(Return(Nil(), notImplementedError(interactionName, r.Name)))
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						})
					}

					g.Default().Block(Return(Nil(), unknownError(Id("resourceType"))))
				case "update":
					for _, r := range resources {
						g.Case(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(
								Return(returnType.Clone().Block(), notImplementedError(interactionName, r.Name)),
							)
							g.List(Id("result"), Id("err")).Op(":=").Id("impl." + interactionName + r.Name).Call(passParams...)
							g.If(Id("err").Op("!=").Nil()).Block(
								Return(returnType.Clone().Block(), Id("err")),
							)
							g.Return(
								returnType.Clone().Block(Dict{
									Id("Resource"): Id("result").Dot("Resource"),
									Id("Created"):  Id("result").Dot("Created"),
								}),
								Nil(),
							)
						})
					}

					g.Default().Block(Return(returnType.Clone().Block(), unknownError(Id("resource").Dot("ResourceType").Call())))
				case "search":
					for _, r := range resources {
						g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(Return(returnType.Clone().Block(), notImplementedError(interactionName, r.Name)))
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						})
					}

					g.Default().Block(Return(
						Qual(moduleName+"/capabilities/search", "Result").Block(),
						unknownError(Id("resourceType")),
					))
				}

			}),
		)
}

func generateGenericSearchCapabilities(resources []ir.ResourceOrType) Code {
	return Func().Params(Id("w").Id(genericWrapperName)).Id("SearchCapabilities").
		Params(Id("ctx").Qual("context", "Context"), Id("resourceType").String()).
		Params(Qual(moduleName+"/capabilities/search", "Capabilities"), Qual(moduleName+"/capabilities", "FHIRError")).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name+"Search")),
						If(Op("!").Id("ok")).Block(
							Return(
								Qual(moduleName+"/capabilities/search", "Capabilities").Block(),
								notImplementedError("search", r.Name),
							),
						),
						Return(Id("impl.SearchCapabilities"+r.Name).Call(Id("ctx"))),
					)
				}

				g.Default().Block(Return(Qual(moduleName+"/capabilities/search", "Capabilities").Block(), unknownError(Id("resourceType"))))
			}),
		)
}

func generateWrapperAllCapabilities(f *File) {
	f.Func().Params(Id("w").Id(genericWrapperName)).Id("AllCapabilities").
		Params(Id("ctx").Qual("context", "Context")).
		Params(
			Qual(moduleName+"/capabilities", "Capabilities"),
			Qual(moduleName+"/capabilities", "FHIRError"),
		).
		Block(
			List(Id("g"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Qual(moduleName+"/capabilities", "GenericCapabilities")),
			If(Id("ok")).Block(
				Comment("// shortcut for the case that the underlying implementation already implements the generic API"),
				Return(Id("g.AllCapabilities").Params(Id("ctx"))),
			),
			Return(Id("AllCapabilities").
				Call(Id("ctx"), Id("w").Dot("Concrete"))),
		)
}

func generateConcreteWrapperStruct(f *File) {
	f.Type().Id(concreteWrapperName).Struct(
		Id("Generic").Qual(moduleName+"/capabilities", "GenericCapabilities"),
	)
}

func generateConcrete(f *File, release string, resources []ir.ResourceOrType, interaction string) {
	interactionName := strcase.ToCamel(interaction)

	for _, r := range resources {
		params := []Code{Id("ctx").Qual("context", "Context")}
		passParams := []Code{Id("ctx")}
		var returnType *Statement
		switch interaction {
		case "create":
			params = append(params, Id("resource").Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name))
			passParams = append(passParams, Id("resource"))
			returnType = Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)
		case "read":
			params = append(params, Id("id").String())
			passParams = append(passParams, Lit(r.Name), Id("id"))
			returnType = Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)
		case "update":
			params = append(params, Id("resource").Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name))
			passParams = append(passParams, Id("resource"))
			returnType = Qual(moduleName+"/capabilities", "UpdateResult").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name))
		case "search":
			params = append(params, Id("options").Qual(moduleName+"/capabilities/search", "Options"))
			passParams = append(passParams, Lit(r.Name), Id("options"))
			returnType = Qual(moduleName+"/capabilities/search", "Result")
			f.Add(generateConcreteSearchCapabilities(r))
		}

		f.Func().Params(Id("w").Id(concreteWrapperName)).Id(interactionName+r.Name).
			Params(params...).
			Params(returnType, Qual(moduleName+"/capabilities", "FHIRError")).
			BlockFunc(func(g *Group) {
				g.List(Id("g"), Id("ok")).Op(":=").Id("w.Generic").Assert(Qual(moduleName+"/capabilities", "Generic"+interactionName))
				g.If(Id("!ok")).Block(
					Return(returnType.Clone().Block(), notImplementedError(interactionName, r.Name)),
				)

				switch interaction {
				case "create", "read":
					g.List(Id("v"), Id("err")).Op(":=").Id("g." + interactionName).Params(passParams...)
					g.If(Id("err").Op("!=").Nil()).Block(Return(returnType.Clone().Block(), Id("err")))

					g.List(Id("contained"), Id("ok")).Op(":=").Id("v").Assert(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), "ContainedResource"),
					)
					g.If(Id("ok")).Block(
						Id("v").Op("=").Id("contained.Resource"),
					)
					g.List(Id("r"), Id("ok")).Op(":=").Id("v").Assert(returnType)
					g.If(Id("!ok")).Block(
						Return(returnType.Clone().Block(), unexpectedError(Lit(r.Name), Id("v").Dot("ResourceType").Call())),
					)
					g.Return(Id("r"), Nil())
				case "update":
					g.List(Id("result"), Id("err")).Op(":=").Id("g." + interactionName).Params(passParams...)
					g.If(Id("err").Op("!=").Nil()).Block(Return(returnType.Clone().Block(), Id("err")))

					g.Id("v").Op(":=").Id("result.Resource")
					g.List(Id("contained"), Id("ok")).Op(":=").Id("v").Assert(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), "ContainedResource"),
					)
					g.If(Id("ok")).Block(
						Id("v").Op("=").Id("contained.Resource"),
					)
					g.List(Id("r"), Id("ok")).Op(":=").Id("v").Assert(
						Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name),
					)
					g.If(Id("!ok")).Block(
						Return(returnType.Clone().Block(), unexpectedError(Lit(r.Name), Id("v").Dot("ResourceType").Call())),
					)
					g.Return(
						returnType.Clone().Block(Dict{
							Id("Resource"): Id("r"),
							Id("Created"):  Id("result").Dot("Created"),
						}),
						Nil(),
					)
				case "search":
					g.List(Id("v"), Id("err")).Op(":=").Id("g." + interactionName).Params(passParams...)
					g.If(Id("err").Op("!=").Nil()).Block(Return(returnType.Clone().Block(), Id("err")))

					g.Return(Id("v"), Nil())
				}
			})
	}
}

func generateConcreteSearchCapabilities(r ir.ResourceOrType) Code {
	returnType := Qual(moduleName+"/capabilities/search", "Capabilities")
	return Func().Params(Id("w").Id(concreteWrapperName)).Id("SearchCapabilities"+r.Name).
		Params(Id("ctx").Qual("context", "Context")).
		Params(returnType, Qual(moduleName+"/capabilities", "FHIRError")).
		BlockFunc(func(g *Group) {
			g.List(Id("allCapabilities"), Id("err")).Op(":=").Id("w.Generic.AllCapabilities").Call(Id("ctx"))
			g.If(Id("err").Op("!=").Nil()).Block(
				Return(returnType.Clone().Block(), Id("err")),
			)
			g.Return(Id("allCapabilities.SearchCapabilities").Index(Lit(r.Name)), Id("err"))
		})
}

func notImplementedError(interaction, resourceType string) Code {
	return Qual(moduleName+"/capabilities", "NotImplementedError").Values(
		Id("Interaction").Op(":").Lit(strings.ToLower(interaction)),
		Id("ResourceType").Op(":").Lit(resourceType),
	)
}

func unknownError(resourceType Code) Code {
	return Qual(moduleName+"/capabilities", "InvalidResourceError").Values(
		Id("ResourceType").Op(":").Add(resourceType),
	)
}

func unexpectedError(expectedType, gotType Code) Code {
	return Qual(moduleName+"/capabilities", "UnexpectedResourceError").Values(
		Id("ExpectedType").Op(":").Add(expectedType),
		Id("GotType").Op(":").Add(gotType),
	)
}
