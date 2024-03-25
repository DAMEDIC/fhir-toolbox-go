package gen

import (
	"fhir-toolbox/generate/ir"
	"log"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
)

func GenerateMarshalHelpers(resources []ir.Struct, genTarget, release string) {
	dir := genDir(genTarget, release)

	f := NewFilePath(dir)

	f.Type().Id("primitiveElement").Struct(
		Id("Id").Id("*string").Tag(map[string]string{"json": "id,omitempty"}),
		Id("Extension").Index().Id("Extension").Tag(map[string]string{"json": "extension,omitempty"}),
	)

	f.Type().Id("containedResource").Struct(
		Id("resource").Id("Resource"),
	)
	implementMarshalJSONContained(f)
	implementUnmarshalJSONContained(resources, f)

	err := f.Save(filepath.Join(dir, "marshal_helpers.go"))
	if err != nil {
		log.Panic(err)
	}
}

func implementMarshalJSONContained(f *File) {
	f.Func().Params(Id("r").Id("containedResource")).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).BlockFunc(func(g *Group) {
		g.Return(Qual("encoding/json", "Marshal").Params(Id("r.resource")))
	})
}

func implementUnmarshalJSONContained(resources []ir.Struct, f *File) {
	f.Func().Params(Id("cr").Op("*").Id("containedResource")).Id("UnmarshalJSON").Params(Id("b").Index().Byte()).Params(Error()).Block(
		Var().Id("t").Struct(
			Id("ResourceType").String().Tag(map[string]string{"json": "resourceType"}),
		),
		If(Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("t")).Op("!=").Nil()).Block(
			Return(Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("t"))),
		),
		Switch(Id("t.ResourceType")).BlockFunc(func(g *Group) {
			for _, r := range resources {
				g.Case(Lit(r.Name)).Block(
					Var().Id("r").Id(r.Name),
					If(Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("r")).Op("!=").Nil()).Block(
						Return(Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("r"))),
					),
					Op("*").Id("cr").Op("=").Id("containedResource").Values(Dict{
						Id("resource"): Id("r"),
					}),
				)
			}

			g.Default().Block(
				Return(Qual("fmt", "Errorf").Call(Lit("unknown resource type: %s"), Id("t.ResourceType"))),
			)
		}),
		Return(Nil()),
	)
}
