package generate

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"slices"
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
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "delete")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "search")

	generateConcreteWrapperStruct(f("concrete", "capabilities"+release))
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "create")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "update")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "delete")
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
		returnType = Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model", "Resource"))
	case "delete":
		params = append(params, Id("resourceType").String(), Id("id").String())
		shortcutParams = append(shortcutParams, Id("resourceType"), Id("id"))
		passParams = append(passParams, Id("id"))
	case "search":
		params = append(params, Id("resourceType").String(), Id("options").Qual(moduleName+"/capabilities/search", "Options"))
		shortcutParams = append(shortcutParams, Id("resourceType"), Id("options"))
		passParams = append(passParams, Id("options"))
		returnType = Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model", "Resource"))
		f.Add(generateGenericSearchCapabilities(release, resources))
	}

	var returns []Code
	if returnType != nil {
		returns = append(returns, returnType)
	}
	returns = append(returns, Error())

	f.Func().Params(Id("w").Id(genericWrapperName)).Id(interactionName).
		Params(params...).
		Params(returns...).
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
							g.If(Op("!").Id("ok")).Block(Return(Nil(), notImplementedError(release, interaction, r.Name)))
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						})
					}

					g.Default().Block(Return(Nil(), invalidResourceTypeError(release, Id("resource").Dot("ResourceType").Call())))
				case "read":
					for _, r := range resources {
						g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(Return(Nil(), notImplementedError(release, interaction, r.Name)))
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						})
					}

					g.Default().Block(Return(Nil(), invalidResourceTypeError(release, Id("resourceType"))))
				case "update":
					for _, r := range resources {
						g.Case(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(
								Return(returnType.Clone().Block(), notImplementedError(release, interaction, r.Name)),
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

					g.Default().Block(Return(returnType.Clone().Block(), invalidResourceTypeError(release, Id("resource").Dot("ResourceType").Call())))
				case "delete":
					for _, r := range resources {
						g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(Return(notImplementedError(release, interaction, r.Name)))
							g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
						})
					}

					g.Default().Block(Return(invalidResourceTypeError(release, Id("resourceType"))))
				case "search":
					for _, r := range resources {
						g.Case(Lit(r.Name)).BlockFunc(func(g *Group) {
							g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
							g.If(Op("!").Id("ok")).Block(Return(returnType.Clone().Block(), notImplementedError(release, interaction, r.Name)))
							g.List(Id("result"), Id("err")).Op(":=").Id("impl." + interactionName + r.Name).Call(passParams...)
							g.If(Id("err").Op("!=").Nil()).Block(Return(returnType.Clone().Block(), Id("err")))

							g.Id("genericResources").Op(":=").Make(Index().Qual(moduleName+"/model", "Resource"), Len(Id("result.Resources")))
							g.For(List(Id("i"), Id("r")).Op(":=").Range().Id("result.Resources")).Block(
								Id("genericResources").Index(Id("i")).Op("=").Id("r"),
							)

							g.Return(returnType.Clone().Block(Dict{
								Id("Resources"): Id("genericResources"),
								Id("Included"):  Id("result").Dot("Included"),
								Id("Next"):      Id("result").Dot("Next"),
							}), Nil())
						})
					}

					g.Default().Block(Return(
						returnType.Clone().Block(),
						invalidResourceTypeError(release, Id("resourceType")),
					))
				}

			}),
		)
}

func generateGenericSearchCapabilities(release string, resources []ir.ResourceOrType) Code {
	return Func().Params(Id("w").Id(genericWrapperName)).Id("SearchCapabilities").
		Params(Id("ctx").Qual("context", "Context"), Id("resourceType").String()).
		Params(Qual(moduleName+"/capabilities/search", "Capabilities"), Error()).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name+"Search")),
						If(Op("!").Id("ok")).Block(
							Return(
								Qual(moduleName+"/capabilities/search", "Capabilities").Block(),
								notImplementedError(release, "search", r.Name),
							),
						),
						Return(Id("impl.SearchCapabilities"+r.Name).Call(Id("ctx"))),
					)
				}

				g.Default().Block(Return(Qual(moduleName+"/capabilities/search", "Capabilities").Block(), invalidResourceTypeError(release, Id("resourceType"))))
			}),
		)
}

func generateWrapperAllCapabilities(f *File) {
	f.Func().Params(Id("w").Id(genericWrapperName)).Id("AllCapabilities").
		Params(Id("ctx").Qual("context", "Context")).
		Params(
			Qual(moduleName+"/capabilities", "Capabilities"),
			Error(),
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
			returnType = Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name))
		case "delete":
			params = append(params, Id("id").String())
			passParams = append(passParams, Lit(r.Name), Id("id"))
		case "search":
			params = append(params, Id("options").Qual(moduleName+"/capabilities/search", "Options"))
			passParams = append(passParams, Lit(r.Name), Id("options"))
			returnType = Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name))
		}

		var returns []Code
		if returnType != nil {
			returns = append(returns, returnType)
		}
		returns = append(returns, Error())

		f.Func().Params(Id("w").Id(concreteWrapperName)).Id(interactionName + r.Name).
			Params(params...).
			Params(returns...).
			BlockFunc(func(g *Group) {
				g.List(Id("g"), Id("ok")).Op(":=").Id("w.Generic").Assert(Qual(moduleName+"/capabilities", "Generic"+interactionName))
				g.If(Id("!ok")).BlockFunc(func(g *Group) {
					if interaction == "delete" {
						g.Return(notImplementedError(release, interactionName, r.Name))
					} else {
						g.Return(returnType.Clone().Block(), notImplementedError(release, interactionName, r.Name))
					}
				})

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
						Return(returnType.Clone().Block(), unexpectedResourceTypeError(release, Lit(r.Name), Id("v").Dot("ResourceType").Call())),
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
						Return(returnType.Clone().Block(), unexpectedResourceTypeError(release, Lit(r.Name), Id("v").Dot("ResourceType").Call())),
					)
					g.Return(
						returnType.Clone().Block(Dict{
							Id("Resource"): Id("r"),
							Id("Created"):  Id("result").Dot("Created"),
						}),
						Nil(),
					)
				case "delete":
					g.Return().Id("g." + interactionName).Params(passParams...)
				case "search":
					g.List(Id("result"), Id("err")).Op(":=").Id("g." + interactionName).Params(passParams...)
					g.If(Id("err").Op("!=").Nil()).Block(Return(returnType.Clone().Block(), Id("err")))

					g.Id("resources").Op(":=").Make(Index().Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name), Len(Id("result").Dot("Resources")))
					g.For(List(Id("i"), Id("v")).Op(":=").Range().Id("result").Dot("Resources")).BlockFunc(func(g *Group) {
						g.List(Id("contained"), Id("ok")).Op(":=").Id("v").Assert(Qual(moduleName+"/model/gen/"+strings.ToLower(release), "ContainedResource"))
						g.If(Id("ok")).Block(Id("v").Op("=").Id("contained").Dot("Resource"))
						g.List(Id("r"), Id("ok")).Op(":=").Id("v").Assert(Qual(moduleName+"/model/gen/"+strings.ToLower(release), r.Name))
						g.If(Op("!").Id("ok")).Block(Return(returnType.Clone().Block(), unexpectedResourceTypeError(release, Lit(r.Name), Id("v").Dot("ResourceType").Call())))
						g.Id("resources").Index(Id("i")).Op("=").Id("r")
					})
					g.Return(returnType.Clone().Block(Dict{
						Id("Resources"): Id("resources"),
						Id("Included"):  Id("result").Dot("Included"),
						Id("Next"):      Id("result").Dot("Next"),
					}), Nil())
				}
			})

		if slices.Contains([]string{"search", "update"}, interaction) {
			f.Add(generateConcreteCapabilities(r, interaction))
		}
	}
}

func generateConcreteCapabilities(r ir.ResourceOrType, interaction string) Code {
	interactionName := strcase.ToCamel(interaction)

	returnType := Qual(moduleName+"/capabilities/"+interaction, "Capabilities")
	return Func().Params(Id("w").Id(concreteWrapperName)).Id(interactionName+"Capabilities"+r.Name).
		Params(Id("ctx").Qual("context", "Context")).
		Params(returnType, Error()).
		BlockFunc(func(g *Group) {
			g.List(Id("allCapabilities"), Id("err")).Op(":=").Id("w.Generic.AllCapabilities").Call(Id("ctx"))
			g.If(Id("err").Op("!=").Nil()).Block(
				Return(returnType.Clone().Block(), Id("err")),
			)
			g.Return(Id("allCapabilities."+interactionName).Index(Lit(r.Name)), Id("err"))
		})
}

func notImplementedError(release, interaction, resourceType string) Code {
	r := strings.ToLower(release)
	return Qual(moduleName+"/model/gen/"+r, "OperationOutcome").Values(Dict{
		Id("Issue"): Index().Qual(moduleName+"/model/gen/"+r, "OperationOutcomeIssue").Values(
			Values(Dict{
				Id("Severity"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("fatal")),
				}),
				Id("Code"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("not-supported")),
				}),
				Id("Diagnostics"): Op("&").Qual(moduleName+"/model/gen/"+r, "String").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit(interaction + " not implemented for " + resourceType)),
				}),
			}),
		),
	})
}

func invalidResourceTypeError(release string, resourceType Code) Code {
	r := strings.ToLower(release)
	return Qual(moduleName+"/model/gen/"+r, "OperationOutcome").Values(Dict{
		Id("Issue"): Index().Qual(moduleName+"/model/gen/"+r, "OperationOutcomeIssue").Values(
			Values(Dict{
				Id("Severity"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("fatal")),
				}),
				Id("Code"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("processing")),
				}),
				Id("Diagnostics"): Op("&").Qual(moduleName+"/model/gen/"+r, "String").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("invalid resource type: ").Op("+").Add(resourceType)),
				}),
			}),
		),
	})
}

func unexpectedResourceTypeError(release string, expectedType, gotType Code) Code {
	r := strings.ToLower(release)
	return Qual(moduleName+"/model/gen/"+r, "OperationOutcome").Values(Dict{
		Id("Issue"): Index().Qual(moduleName+"/model/gen/"+r, "OperationOutcomeIssue").Values(
			Values(Dict{
				Id("Severity"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("fatal")),
				}),
				Id("Code"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("processing")),
				}),
				Id("Diagnostics"): Op("&").Qual(moduleName+"/model/gen/"+r, "String").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils", "Ptr").Call(Lit("expected ").Op("+").Add(expectedType).Op("+").Lit(" but got ").Op("+").Add(gotType)),
				}),
			}),
		),
	})
}
