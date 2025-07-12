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
	generateGenericWrapperStruct(f("generic", "capabilities"+release), release)
	generateWrapperCapabilityStatement(f("generic", "capabilities"+release), release, ir.FilterResources(rt))
	generateSearchParametersFn(f("generic", "capabilities"+release), release, ir.FilterResources(rt))
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "create")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "update")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "delete")
	generateGeneric(f("generic", "capabilities"+release), release, ir.FilterResources(rt), "search")

	generateConcreteWrapperStruct(f("concrete", "capabilities"+release), release)
	generateConcreteCapabilityBase(f("concrete", "capabilities"+release), release)
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "create")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "read")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "update")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "delete")
	generateConcrete(f("concrete", "capabilities"+release), release, ir.FilterResources(rt), "search")
}

func generateGenericWrapperStruct(f *File, release string) {
	f.Type().Id(genericWrapperName).Struct(
		Id("Concrete").Qual(moduleName+"/capabilities", "ConcreteCapabilities"),
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
		BlockFunc(func(g *Group) {
			g.List(Id("g"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Qual(moduleName+"/capabilities", "Generic"+interactionName))
			g.If(Id("ok")).Block(
				Comment("// shortcut for the case that the underlying implementation already implements the generic API"),
				Return(Id("g."+interactionName).Params(shortcutParams...)),
			)

			g.Switch(switchType).BlockFunc(func(g *Group) {
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
							if r.Name == "SearchParameter" {
								// Special handling for SearchParameter - fallback to SearchCapabilities methods
								g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
								g.If(Id("ok")).Block(
									Return(Id("impl." + interactionName + r.Name).Call(passParams...)),
								)

								// Fallback: gather SearchParameter from SearchCapabilities methods
								g.Comment("// Fallback: gather SearchParameter from SearchCapabilities methods if ReadSearchParameter not implemented")
								g.List(Id("searchParameters"), Id("err")).Op(":=").Id("searchParameters").Call(Id("ctx"), Id("w").Dot("Concrete"))
								g.If(Id("err").Op("!=").Nil()).Block(
									Return(Nil(), Id("err")),
								)

								g.List(Id("searchParam"), Id("exists")).Op(":=").Id("searchParameters").Index(Id("id"))
								g.If(Id("exists")).Block(
									Return(Id("searchParam"), Nil()),
								)

								g.Return(Nil(), notImplementedError(release, interaction, r.Name))
							} else {
								g.List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name + interactionName))
								g.If(Op("!").Id("ok")).Block(Return(Nil(), notImplementedError(release, interaction, r.Name)))
								g.Return(Id("impl." + interactionName + r.Name).Call(passParams...))
							}
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

			})
		})
}

func generateGenericSearchCapabilities(release string, resources []ir.ResourceOrType) Code {
	return Func().Params(Id("w").Id(genericWrapperName)).Id("SearchCapabilities").
		Params(Id("ctx").Qual("context", "Context"), Id("resourceType").String()).
		Params(searchCapabilitiesType(release), Error()).
		Block(
			Switch(Id("resourceType")).BlockFunc(func(g *Group) {
				for _, r := range resources {
					g.Case(Lit(r.Name)).Block(
						List(Id("impl"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Id(r.Name+"Search")),
						If(Op("!").Id("ok")).Block(
							Return(
								searchCapabilitiesType(release).Block(),
								notImplementedError(release, "search", r.Name),
							),
						),
						Return(Id("impl.SearchCapabilities"+r.Name).Call(Id("ctx"))),
					)
				}

				g.Default().Block(Return(searchCapabilitiesType(release).Block(), invalidResourceTypeError(release, Id("resourceType"))))
			}),
		)
}

func generateWrapperCapabilityStatement(f *File, release string, resources []ir.ResourceOrType) {
	f.Func().Params(Id("w").Id(genericWrapperName)).Id("CapabilityStatement").
		Params(
			Id("ctx").Qual("context", "Context"),
		).
		Params(
			Qual(moduleName+"/model/gen/basic", "CapabilityStatement"),
			Error(),
		).
		BlockFunc(func(g *Group) {
			// Check if concrete implementation also implements GenericCapabilities for shortcut
			g.List(Id("gen"), Id("ok")).Op(":=").Id("w.Concrete").Assert(Qual(moduleName+"/capabilities", "GenericCapabilities"))
			g.If(Id("ok")).Block(
				Comment("// shortcut for the case that the underlying implementation already implements the generic API"),
				Return(Id("gen.CapabilityStatement").Params(Id("ctx"))),
			)

			// Generate CapabilityStatement from ConcreteCapabilities
			g.Comment("// Generate CapabilityStatement from concrete implementation")
			g.List(Id("baseCapabilityStatement"), Id("err")).Op(":=").Id("w.Concrete.CapabilityBase").Call(Id("ctx"))
			g.If(Id("err").Op("!=").Nil()).Block(
				Return(Qual(moduleName+"/model/gen/basic", "CapabilityStatement").Values(), Id("err")),
			)

			// Extract and validate base URL for canonical references
			g.Var().Id("baseUrl").String()
			g.If(Id("baseCapabilityStatement.Implementation").Op("==").Nil().Op("||").Id("baseCapabilityStatement.Implementation.Url").Op("==").Nil().Op("||").Id("baseCapabilityStatement.Implementation.Url.Value").Op("==").Nil()).Block(
				Return(Qual(moduleName+"/model/gen/basic", "CapabilityStatement").Values(), Qual("fmt", "Errorf").Call(Lit("base CapabilityStatement must have implementation.url set for canonical SearchParameter references"))),
			)
			g.Id("baseUrl").Op("=").Op("*").Id("baseCapabilityStatement.Implementation.Url.Value")

			// Initialize resourcesMap from base CapabilityStatement
			g.Id("resourcesMap").Op(":=").Make(Map(String()).Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResource"))
			g.For(List(Id("_"), Id("rest")).Op(":=").Range().Id("baseCapabilityStatement.Rest")).Block(
				For(List(Id("_"), Id("resource")).Op(":=").Range().Id("rest.Resource")).Block(
					If(Id("resource.Type.Value").Op("!=").Nil()).Block(
						Id("resourcesMap").Index(Op("*").Id("resource.Type.Value")).Op("=").Id("resource"),
					),
				),
			)

			g.Var().Id("errs").Index().Error()

			// Helper function to add interactions
			g.Id("addInteraction").Op(":=").Func().Params(
				Id("name").String(),
				Id("interactionCode").String(),
			).Params(
				Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResource"),
			).Block(
				List(Id("r"), Id("ok")).Op(":=").Id("resourcesMap").Index(Id("name")),
				If(Op("!").Id("ok")).Block(
					Id("r").Op("=").Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResource").Values(Dict{
						Id("Type"): Qual(moduleName+"/model/gen/basic", "Code").Values(Dict{Id("Value"): Op("&").Id("name")}),
					}),
				),
				Id("r").Dot("Interaction").Op("=").Append(
					Id("r").Dot("Interaction"),
					Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResourceInteraction").Values(Dict{
						Id("Code"): Qual(moduleName+"/model/gen/basic", "Code").Values(Dict{Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Id("interactionCode"))}),
					}),
				),
				Return(Id("r")),
			)

			for _, r := range resources {
				for _, interaction := range []string{"create", "read", "delete"} {
					interactionName := strcase.ToCamel(interaction)

					g.If(
						List(Id("_"), Id("ok")).Op(":=").Id("w.Concrete").Dot("").Call(
							Id(r.Name+interactionName),
						),
						Id("ok"),
					).Block(
						Id("resourcesMap").Index(Lit(r.Name)).Op("=").Id("addInteraction").Call(Lit(r.Name), Lit(interaction)),
					)
				}

				// Update
				g.If(
					List(Id("c"), Id("ok")).Op(":=").Id("w.Concrete").Dot("").Call(
						Id(r.Name+"Update"),
					),
					Id("ok"),
				).Block(
					Id("r").Op(":=").Id("addInteraction").Call(Lit(r.Name), Lit("update")),
					List(Id("c"), Id("ok")).Op(":=").Id("c").Dot("").Call(Id(r.Name+"UpdateCapabilities")),
					If(Id("ok")).Block(
						List(Id("c"), Err()).Op(":=").Id("c").Dot("UpdateCapabilities"+r.Name).Call(Id("ctx")),
						If(Err().Op("!=").Nil()).Block(
							Id("errs").Op("=").Append(Id("errs"), Err()),
						).Else().Block(
							Id("r").Dot("UpdateCreate").Op("=").Op("&").Qual(moduleName+"/model/gen/basic", "Boolean").Values(Dict{
								Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Id("c").Dot("UpdateCreate")),
							}),
						),
					),
					Id("resourcesMap").Index(Lit(r.Name)).Op("=").Id("r"),
				)

				// Search
				g.If(
					List(Id("c"), Id("ok")).Op(":=").Id("w.Concrete").Dot("").Call(
						Id(r.Name+"Search"),
					),
					Id("ok"),
				).Block(
					List(Id("c"), Err()).Op(":=").Id("c").Dot("SearchCapabilities"+r.Name).Call(Id("ctx")),
					If(Err().Op("!=").Nil()).Block(
						Id("errs").Op("=").Append(Id("errs"), Err()),
					).Else().Block(
						Id("r").Op(":=").Id("addInteraction").Call(Lit(r.Name), Lit("search-type")),

						// Add includes
						For(List(Id("_"), Id("include")).Op(":=").Range().Id("c").Dot("Includes")).Block(
							Id("r").Dot("SearchInclude").Op("=").Append(
								Id("r").Dot("SearchInclude"),
								Qual(moduleName+"/model/gen/basic", "String").Values(Dict{Id("Value"): Op("&").Id("include")}),
							),
						),

						// Add search parameters
						For(List(Id("n"), Id("p")).Op(":=").Range().Id("c").Dot("Parameters")).Block(
							List(Id("fhirpathType"), Id("ok"), Id("err")).Op(":=").Qual(moduleName+"/fhirpath", "Singleton").Index(Qual(moduleName+"/fhirpath", "String")).Call(Id("p").Dot("Children").Call(Lit("type"))),
							If(Op("!").Id("ok").Op("||").Id("err").Op("!=").Nil()).Block(
								Continue(),
							),
							Id("resolvedType").Op(":=").String().Call(Id("fhirpathType")),

							// Extract SearchParameter ID for canonical reference
							Var().Id("definition").Op("*").Qual(moduleName+"/model/gen/basic", "Canonical"),
							If(Id("baseUrl").Op("!=").Lit("")).Block(
								Id("searchParameterId").Op(":=").Lit(""),
								List(Id("fhirpathId"), Id("idOk"), Id("idErr")).Op(":=").Qual(moduleName+"/fhirpath", "Singleton").Index(Qual(moduleName+"/fhirpath", "String")).Call(Id("p").Dot("Children").Call(Lit("id"))),
								If(Id("idOk").Op("&&").Id("idErr").Op("==").Nil()).Block(
									Id("searchParameterId").Op("=").String().Call(Id("fhirpathId")),
								).Else().Block(
									Comment("// If no explicit ID is set, create one of pattern {resourceType}-{name}"),
									Id("searchParameterId").Op("=").Lit(r.Name).Op("+").Lit("-").Op("+").Id("n"),
								),
								Id("canonicalUrl").Op(":=").Id("baseUrl").Op("+").Lit("/SearchParameter/").Op("+").Id("searchParameterId"),
								Id("definition").Op("=").Op("&").Qual(moduleName+"/model/gen/basic", "Canonical").Values(Dict{Id("Value"): Op("&").Id("canonicalUrl")}),
							),

							Id("r").Dot("SearchParam").Op("=").Append(
								Id("r").Dot("SearchParam"),
								Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResourceSearchParam").Values(Dict{
									Id("Name"):       Qual(moduleName+"/model/gen/basic", "String").Values(Dict{Id("Value"): Op("&").Id("n")}),
									Id("Type"):       Qual(moduleName+"/model/gen/basic", "Code").Values(Dict{Id("Value"): Op("&").Id("resolvedType")}),
									Id("Definition"): Id("definition"),
								}),
							),
						),

						Id("resourcesMap").Index(Lit(r.Name)).Op("=").Id("r"),
					),
				)
			}

			// Always add SearchParameter read capability
			g.Id("resourcesMap").Index(Lit("SearchParameter")).Op("=").Id("addInteraction").Call(Lit("SearchParameter"), Lit("read"))

			// Check for errors before proceeding
			g.If(Len(Id("errs")).Op(">").Lit(0)).Block(
				Return(Qual(moduleName+"/model/gen/basic", "CapabilityStatement").Values(), Qual("errors", "Join").Call(Id("errs").Op("..."))),
			)

			// Convert map to slice and sort
			g.Id("resourcesList").Op(":=").Make(Index().Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResource"), Lit(0), Len(Id("resourcesMap")))

			g.For(List(Id("_"), Id("r")).Op(":=").Range().Id("resourcesMap")).Block(
				// Sort search params by name
				Qual("slices", "SortStableFunc").Call(
					Id("r").Dot("SearchParam"),
					Func().Params(
						Id("a"),
						Id("b").Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResourceSearchParam"),
					).Int().Block(
						Return(Qual("cmp", "Compare").Call(Op("*").Id("a").Dot("Name").Dot("Value"), Op("*").Id("b").Dot("Name").Dot("Value"))),
					),
				),
				// Sort interactions in standard order: create, read, update, delete, search-type
				Qual("slices", "SortStableFunc").Call(
					Id("r").Dot("Interaction"),
					Func().Params(
						Id("a"),
						Id("b").Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResourceInteraction"),
					).Int().Block(
						Id("order").Op(":=").Map(String()).Int().Values(Dict{
							Lit("create"):      Lit(1),
							Lit("read"):        Lit(2),
							Lit("update"):      Lit(3),
							Lit("delete"):      Lit(4),
							Lit("search-type"): Lit(5),
						}),
						Id("aCode").Op(":=").Lit(""),
						If(Id("a").Dot("Code").Dot("Value").Op("!=").Nil()).Block(
							Id("aCode").Op("=").Op("*").Id("a").Dot("Code").Dot("Value"),
						),
						Id("bCode").Op(":=").Lit(""),
						If(Id("b").Dot("Code").Dot("Value").Op("!=").Nil()).Block(
							Id("bCode").Op("=").Op("*").Id("b").Dot("Code").Dot("Value"),
						),
						Return(Qual("cmp", "Compare").Call(Id("order").Index(Id("aCode")), Id("order").Index(Id("bCode")))),
					),
				),
				Id("resourcesList").Op("=").Append(Id("resourcesList"), Id("r")),
			)

			// Sort resources by type
			g.Qual("slices", "SortFunc").Call(
				Id("resourcesList"),
				Func().Params(
					Id("a"),
					Id("b").Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResource"),
				).Int().Block(
					Return(Qual("cmp", "Compare").Call(Op("*").Id("a").Dot("Type").Dot("Value"), Op("*").Id("b").Dot("Type").Dot("Value"))),
				),
			)

			// Update CapabilityStatement with detected capabilities
			g.Id("capabilityStatement").Op(":=").Id("baseCapabilityStatement")

			// Set FHIR version if not explicitly set based on concrete implementation
			var fhirVersion string
			switch release {
			case "R4":
				fhirVersion = "4.0"
			case "R4B":
				fhirVersion = "4.3"
			case "R5":
				fhirVersion = "5.0"
			default:
				fhirVersion = release // fallback to release name
			}

			g.If(Id("capabilityStatement.FhirVersion.Value").Op("==").Nil()).Block(
				Id("capabilityStatement.FhirVersion").Op("=").Qual(moduleName+"/model/gen/basic", "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit(fhirVersion)),
				}),
			)

			// Ensure REST section exists
			g.If(Len(Id("capabilityStatement.Rest")).Op("==").Lit(0)).Block(
				Id("capabilityStatement.Rest").Op("=").Index().Qual(moduleName+"/model/gen/basic", "CapabilityStatementRest").Values(
					Values(Dict{
						Id("Mode"): Qual(moduleName+"/model/gen/basic", "Code").Values(Dict{Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("server"))}),
					}),
				),
			)

			// Update the first REST entry with detected resources
			g.Id("capabilityStatement.Rest").Index(Lit(0)).Dot("Resource").Op("=").Id("resourcesList")

			g.Return(Id("capabilityStatement"), Nil())
		})
}

func generateConcreteWrapperStruct(f *File, release string) {
	f.Type().Id(concreteWrapperName).Struct(
		Id("Generic").Qual(moduleName+"/capabilities", "GenericCapabilities"),
	)
}

func generateConcreteCapabilityBase(f *File, release string) {
	f.Func().Params(Id("w").Id(concreteWrapperName)).Id("CapabilityBase").
		Params(Id("ctx").Qual("context", "Context")).
		Params(
			Qual(moduleName+"/model/gen/basic", "CapabilityStatement"),
			Error(),
		).
		Block(
			Comment("// Delegate to the generic CapabilityStatement method"),
			Return(Id("w.Generic.CapabilityStatement").Params(Id("ctx"))),
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
			f.Add(generateConcreteCapabilities(r, release, interaction))
		}
	}
}

func generateConcreteCapabilities(r ir.ResourceOrType, release, interaction string) Code {
	interactionName := strcase.ToCamel(interaction)

	// Define the return type for the function signature
	returnType := Qual(moduleName+"/capabilities/"+interaction, "Capabilities")
	if interaction == "search" {
		returnType.Add(Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter")))
	}

	return Func().Params(Id("w").Id(concreteWrapperName)).Id(interactionName+"Capabilities"+r.Name).
		Params(Id("ctx").Qual("context", "Context")).
		Params(returnType, Error()).
		BlockFunc(func(g *Group) {
			// The logic differs for "search" and "update"
			if interaction == "search" {
				// Read capabilities from the generic CapabilityStatement method
				g.List(Id("capabilityStatement"), Id("err")).Op(":=").Id("w.Generic.CapabilityStatement").Call(Id("ctx"))
				g.If(Id("err").Op("!=").Nil()).Block(
					Return(Add(returnType.Clone()).Values(Dict{
						Id("Includes"):   Index().String().Values(),
						Id("Parameters"): Make(Map(String()).Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter")),
					}), Id("err")),
				)

				// Find the resource in the CapabilityStatement
				g.Var().Id("searchParams").Index().Qual(moduleName+"/model/gen/basic", "CapabilityStatementRestResourceSearchParam")
				g.For(List(Id("_"), Id("rest")).Op(":=").Range().Id("capabilityStatement.Rest")).Block(
					For(List(Id("_"), Id("resource")).Op(":=").Range().Id("rest.Resource")).Block(
						If(Id("resource.Type.Value").Op("!=").Nil().Op("&&").Op("*").Id("resource.Type.Value").Op("==").Lit(r.Name)).Block(
							Id("searchParams").Op("=").Id("resource.SearchParam"),
							Break(),
						),
					),
				)

				// Initialize the parameters map
				g.Id("parameters").Op(":=").Make(Map(String()).Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter"))

				// Look up actual SearchParameter resources using the generic Read method
				g.For(List(Id("_"), Id("searchParam")).Op(":=").Range().Id("searchParams")).Block(
					If(Id("searchParam.Definition").Op("!=").Nil().Op("&&").Id("searchParam.Definition.Value").Op("!=").Nil()).Block(
						// Extract SearchParameter ID from canonical URL
						Id("canonicalUrl").Op(":=").Op("*").Id("searchParam.Definition.Value"),
						Id("lastSlash").Op(":=").Qual("strings", "LastIndex").Call(Id("canonicalUrl"), Lit("/")),
						If(Id("lastSlash").Op("!=").Lit(-1).Op("&&").Id("lastSlash").Op("<").Len(Id("canonicalUrl")).Op("-").Lit(1)).Block(
							Id("searchParamId").Op(":=").Id("canonicalUrl").Index(Id("lastSlash").Op("+").Lit(1).Op(":")),

							// Read the SearchParameter resource using generic Read
							List(Id("readImpl"), Id("readOk")).Op(":=").Id("w.Generic").Assert(Qual(moduleName+"/capabilities", "GenericRead")),
							If(Id("readOk")).Block(
								List(Id("searchParamResource"), Id("readErr")).Op(":=").Id("readImpl.Read").Call(Id("ctx"), Lit("SearchParameter"), Id("searchParamId")),
								If(Id("readErr").Op("==").Nil()).Block(
									// Type assert to SearchParameter
									List(Id("sp"), Id("ok")).Op(":=").Id("searchParamResource").Assert(Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter")),
									If(Id("ok")).Block(
										If(Id("sp.Code.Value").Op("!=").Nil()).Block(
											Id("parameters").Index(Op("*").Id("sp.Code.Value")).Op("=").Id("sp"),
										),
									),
								),
							),
						),
					),
				)

				// Return the capabilities with resolved SearchParameters
				g.Return(Add(returnType.Clone()).Values(Dict{
					Id("Includes"):   Index().String().Values(),
					Id("Parameters"): Id("parameters"),
				}), Nil())
			} else {
				// For update capabilities, use type assertion
				g.List(Id("updateImpl"), Id("ok")).Op(":=").Id("w.Generic").Assert(Id(r.Name + "UpdateCapabilities"))
				g.If(Id("ok")).Block(
					Return(Id("updateImpl.UpdateCapabilities" + r.Name).Call(Id("ctx"))),
				)

				// If not available, return default capabilities
				g.Return(Qual(moduleName+"/capabilities/update", "Capabilities").Values(), Nil())
			}
		})
}

func searchCapabilitiesType(release string) *Statement {
	return Qual(moduleName+"/capabilities/search", "Capabilities").Index(Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter"))
}

func notImplementedError(release, interaction, resourceType string) Code {
	r := strings.ToLower(release)
	return Qual(moduleName+"/model/gen/"+r, "OperationOutcome").Values(Dict{
		Id("Issue"): Index().Qual(moduleName+"/model/gen/"+r, "OperationOutcomeIssue").Values(
			Values(Dict{
				Id("Severity"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("fatal")),
				}),
				Id("Code"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("not-supported")),
				}),
				Id("Diagnostics"): Op("&").Qual(moduleName+"/model/gen/"+r, "String").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit(interaction + " not implemented for " + resourceType)),
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
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("fatal")),
				}),
				Id("Code"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("processing")),
				}),
				Id("Diagnostics"): Op("&").Qual(moduleName+"/model/gen/"+r, "String").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("invalid resource type: ").Op("+").Add(resourceType)),
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
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("fatal")),
				}),
				Id("Code"): Qual(moduleName+"/model/gen/"+r, "Code").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("processing")),
				}),
				Id("Diagnostics"): Op("&").Qual(moduleName+"/model/gen/"+r, "String").Values(Dict{
					Id("Value"): Qual(moduleName+"/utils/ptr", "To").Call(Lit("expected ").Op("+").Add(expectedType).Op("+").Lit(" but got ").Op("+").Add(gotType)),
				}),
			}),
		),
	})
}

func generateSearchParametersFn(f *File, release string, resources []ir.ResourceOrType) {
	f.Func().Id("searchParameters").
		Params(
			Id("ctx").Qual("context", "Context"),
			Id("api").Any(),
		).
		Params(
			Map(String()).Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter"),
			Error(),
		).BlockFunc(func(g *Group) {

		// Collections for building SearchParameters
		g.Id("searchParameters").Op(":=").Make(Map(String()).Qual(moduleName+"/model/gen/"+strings.ToLower(release), "SearchParameter"))
		g.Var().Id("errs").Index().Error()

		for _, r := range resources {
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
					// Store SearchParameter for aggregation, indexed by ID
					For(List(Id("n"), Id("p")).Op(":=").Range().Id("c").Dot("Parameters")).Block(
						// Extract ID from SearchParameter using FHIRPath
						Id("searchParameterId").Op(":=").Lit(""),
						List(Id("fhirpathId"), Id("ok"), Id("err")).Op(":=").Qual(moduleName+"/fhirpath", "Singleton").Index(Qual(moduleName+"/fhirpath", "String")).Call(Id("p").Dot("Children").Call(Lit("id"))),
						If(Id("ok").Op("&&").Id("err").Op("==").Nil()).Block(
							Id("searchParameterId").Op("=").String().Call(Id("fhirpathId")),
						).Else().Block(
							Comment("// If no explicit ID is set, create one of pattern {resourceType}-{name}"),
							Id("searchParameterId").Op("=").Lit(r.Name).Op("+").Lit("-").Op("+").Id("n"),
						),
						Id("searchParameters").Index(Id("searchParameterId")).Op("=").Id("p"),
					),
				),
			)
		}

		// Check for errors before proceeding
		g.If(Len(Id("errs")).Op(">").Lit(0)).Block(
			Return(Nil(), Qual("errors", "Join").Call(Id("errs").Op("..."))),
		)

		g.Return(Id("searchParameters"), Nil())
	})
}
