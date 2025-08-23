package generate

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type ClientGenerator struct {
	NoOpGenerator
}

func (g ClientGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	resources := ir.FilterResources(rt)

	// Generate unified client that implements both generic and concrete APIs
	GenerateUnifiedClient(f("gen_client_"+strings.ToLower(release), "rest"), resources, release)
}

func GenerateUnifiedClient(f *File, resources []ir.ResourceOrType, release string) {
	releaseType := strings.ToUpper(release)
	lowerRelease := strings.ToLower(release)

	// Add build tag
	var buildTag string
	switch lowerRelease {
	case "r4":
		buildTag = "//go:build r4 || !(r4 || r4b || r5)"
	case "r4b":
		buildTag = "//go:build r4b || !(r4 || r4b || r5)"
	case "r5":
		buildTag = "//go:build r5 || !(r4 || r4b || r5)"
	}
	f.HeaderComment(buildTag)
	f.Line()

	// Add imports
	f.ImportName("context", "context")
	f.ImportName("fmt", "fmt")
	f.ImportName("net/http", "http")
	f.ImportName("net/url", "url")
	f.ImportName(moduleName+"/capabilities/search", "search")
	f.ImportName(moduleName+"/capabilities/update", "update")
	f.ImportName(moduleName+"/model", "model")
	f.ImportName(moduleName+"/model/gen/basic", "basic")
	f.ImportName(moduleName+"/model/gen/"+lowerRelease, lowerRelease)

	// Generate client struct
	clientName := "Client" + releaseType
	f.Comment("// " + clientName + " provides both generic and resource-specific FHIR client capabilities.")
	f.Type().Id(clientName).Struct(
		Id("BaseURL").Op("*").Qual("net/url", "URL"),
		Id("Client").Op("*").Qual("net/http", "Client"),
	)

	// Add HTTP client helper method
	f.Comment("// httpClient returns the HTTP client, using http.DefaultClient if none is set.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("httpClient").Params().Op("*").Qual("net/http", "Client").Block(
		If(Id("c").Dot("Client").Op("!=").Nil()).Block(
			Return(Id("c").Dot("Client")),
		),
		Return(Qual("net/http", "DefaultClient")),
	)

	// Add constructor function
	constructorName := "NewClient" + releaseType
	f.Comment("// " + constructorName + " creates a new " + releaseType + " FHIR client with the given base URL and HTTP client.")
	f.Comment("// If httpClient is nil, http.DefaultClient will be used.")
	f.Func().Id(constructorName).Params(
		Id("baseURL").String(),
		Id("httpClient").Op("*").Qual("net/http", "Client"),
	).Params(
		Op("*").Id(clientName),
		Error(),
	).Block(
		List(Id("u"), Id("err")).Op(":=").Qual("net/url", "Parse").Call(Id("baseURL")),
		If(Id("err").Op("!=").Nil()).Block(
			Return(Nil(), Id("err")),
		),
		Return(Op("&").Id(clientName).Values(
			Id("BaseURL").Op(":").Id("u"),
			Id("Client").Op(":").Id("httpClient"),
		), Nil()),
	)

	// Generate generic methods
	generateUnifiedGenericMethods(f, clientName, releaseType)

	// Generate concrete methods for each resource
	for _, resource := range resources {
		generateUnifiedConcreteResourceMethods(f, resource, release, releaseType, clientName)
	}
}

func generateUnifiedGenericMethods(f *File, clientName, releaseType string) {
	// CapabilityStatement method
	f.Comment("// CapabilityStatement retrieves the server's CapabilityStatement.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("CapabilityStatement").Params(
		Id("ctx").Qual("context", "Context"),
	).Params(
		Qual(moduleName+"/model/gen/basic", "CapabilityStatement"),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("CapabilityStatement").Call(Id("ctx"))),
	)

	// Generic Create method
	f.Comment("// Create creates a new resource.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Create").Params(
		Id("ctx").Qual("context", "Context"),
		Id("resource").Qual(moduleName+"/model", "Resource"),
	).Params(
		Qual(moduleName+"/model", "Resource"),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("Create").Call(Id("ctx"), Id("resource"))),
	)

	// Generic Read method
	f.Comment("// Read retrieves a resource by type and ID.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Read").Params(
		Id("ctx").Qual("context", "Context"),
		Id("resourceType").String(),
		Id("id").String(),
	).Params(
		Qual(moduleName+"/model", "Resource"),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("Read").Call(Id("ctx"), Id("resourceType"), Id("id"))),
	)

	// Generic Update method
	f.Comment("// Update updates an existing resource.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Update").Params(
		Id("ctx").Qual("context", "Context"),
		Id("resource").Qual(moduleName+"/model", "Resource"),
	).Params(
		Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model", "Resource")),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("Update").Call(Id("ctx"), Id("resource"))),
	)

	// Generic Delete method
	f.Comment("// Delete deletes a resource by type and ID.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Delete").Params(
		Id("ctx").Qual("context", "Context"),
		Id("resourceType").String(),
		Id("id").String(),
	).Params(
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("Delete").Call(Id("ctx"), Id("resourceType"), Id("id"))),
	)

	// Generic Search method
	f.Comment("// Search performs a search operation for the given resource type.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Search").Params(
		Id("ctx").Qual("context", "Context"),
		Id("resourceType").String(),
		Id("parameters").Qual(moduleName+"/capabilities/search", "Parameters"),
		Id("options").Qual(moduleName+"/capabilities/search", "Options"),
	).Params(
		Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model", "Resource")),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("Search").Call(Id("ctx"), Id("resourceType"), Id("parameters"), Id("options"))),
	)
}

func generateUnifiedConcreteResourceMethods(f *File, resource ir.ResourceOrType, release, releaseType, clientName string) {
	resourceName := resource.Name
	lowerRelease := strings.ToLower(release)

	// Create method
	f.Comment("// Create" + resourceName + " creates a new " + resourceName + " resource.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Create"+resourceName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("resource").Qual(moduleName+"/model/gen/"+lowerRelease, resourceName),
	).Params(
		Qual(moduleName+"/model/gen/"+lowerRelease, resourceName),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		List(Id("result"), Id("err")).Op(":=").Id("client").Dot("Create").Call(Id("ctx"), Id("resource")),
		If(Id("err").Op("!=").Nil()).Block(
			Return(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName).Values(), Id("err")),
		),
		List(Id("typed"), Id("ok")).Op(":=").Id("result").Assert(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)),
		If(Op("!").Id("ok")).Block(
			Return(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName).Values(), Qual("fmt", "Errorf").Call(Lit("unexpected resource type: %T"), Id("result"))),
		),
		Return(Id("typed"), Nil()),
	)

	// Read method
	f.Comment("// Read" + resourceName + " retrieves a " + resourceName + " resource by ID.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Read"+resourceName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("id").String(),
	).Params(
		Qual(moduleName+"/model/gen/"+lowerRelease, resourceName),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		List(Id("result"), Id("err")).Op(":=").Id("client").Dot("Read").Call(Id("ctx"), Lit(resourceName), Id("id")),
		If(Id("err").Op("!=").Nil()).Block(
			Return(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName).Values(), Id("err")),
		),
		List(Id("typed"), Id("ok")).Op(":=").Id("result").Assert(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)),
		If(Op("!").Id("ok")).Block(
			Return(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName).Values(), Qual("fmt", "Errorf").Call(Lit("unexpected resource type: %T"), Id("result"))),
		),
		Return(Id("typed"), Nil()),
	)

	// Update method
	f.Comment("// Update" + resourceName + " updates an existing " + resourceName + " resource.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Update"+resourceName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("resource").Qual(moduleName+"/model/gen/"+lowerRelease, resourceName),
	).Params(
		Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		List(Id("result"), Id("err")).Op(":=").Id("client").Dot("Update").Call(Id("ctx"), Id("resource")),
		If(Id("err").Op("!=").Nil()).Block(
			Return(Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)).Values(), Id("err")),
		),
		List(Id("typed"), Id("ok")).Op(":=").Id("result").Dot("Resource").Assert(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)),
		If(Op("!").Id("ok")).Block(
			Return(Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)).Values(), Qual("fmt", "Errorf").Call(Lit("unexpected resource type: %T"), Id("result").Dot("Resource"))),
		),
		Return(Qual(moduleName+"/capabilities/update", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)).Values(
			Id("Resource").Op(":").Id("typed"),
			Id("Created").Op(":").Id("result").Dot("Created"),
		), Nil()),
	)

	// Delete method
	f.Comment("// Delete" + resourceName + " deletes a " + resourceName + " resource by ID.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Delete"+resourceName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("id").String(),
	).Params(
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		Return(Id("client").Dot("Delete").Call(Id("ctx"), Lit(resourceName), Id("id"))),
	)

	// Search method
	f.Comment("// Search" + resourceName + " performs a search for " + resourceName + " resources.")
	f.Func().Params(Id("c").Op("*").Id(clientName)).Id("Search"+resourceName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("parameters").Qual(moduleName+"/capabilities/search", "Parameters"),
		Id("options").Qual(moduleName+"/capabilities/search", "Options"),
	).Params(
		Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)),
		Error(),
	).Block(
		Id("client").Op(":=").Op("&").Id("internalClient").Index(Qual(moduleName+"/model", releaseType)).Values(
			Id("baseURL").Op(":").Id("c").Dot("BaseURL"),
			Id("client").Op(":").Id("c").Dot("httpClient").Call(),
		),
		List(Id("result"), Id("err")).Op(":=").Id("client").Dot("Search").Call(Id("ctx"), Lit(resourceName), Id("parameters"), Id("options")),
		If(Id("err").Op("!=").Nil()).Block(
			Return(Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)).Values(), Id("err")),
		),
		Comment("// Convert generic resources to typed resources"),
		Id("typedResources").Op(":=").Make(Index().Qual(moduleName+"/model/gen/"+lowerRelease, resourceName), Len(Id("result").Dot("Resources"))),
		For(List(Id("i"), Id("resource")).Op(":=").Range().Id("result").Dot("Resources")).Block(
			List(Id("typed"), Id("ok")).Op(":=").Id("resource").Assert(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)),
			If(Op("!").Id("ok")).Block(
				Return(Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)).Values(), Qual("fmt", "Errorf").Call(Lit("unexpected resource type in results: %T"), Id("resource"))),
			),
			Id("typedResources").Index(Id("i")).Op("=").Id("typed"),
		),
		Return(Qual(moduleName+"/capabilities/search", "Result").Index(Qual(moduleName+"/model/gen/"+lowerRelease, resourceName)).Values(
			Id("Resources").Op(":").Id("typedResources"),
			Id("Included").Op(":").Id("result").Dot("Included"),
			Id("Next").Op(":").Id("result").Dot("Next"),
		), Nil()),
	)
}
