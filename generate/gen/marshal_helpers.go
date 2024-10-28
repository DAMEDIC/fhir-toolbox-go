package gen

import (
	"fhir-toolbox/generate/gen/json"
	"fhir-toolbox/generate/gen/xml"
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

	f.Type().Id("ContainedResource").Struct(
		Qual("fhir-toolbox/model", "Resource"),
	)
	implementStringContained(f)

	json.ImplementMarshalContained(f)
	json.ImplementUnmarshalContained(resources, f)

	xml.ImplementMarshalContained(f)
	xml.ImplementUnmarshalContained(resources, f)

	err := f.Save(filepath.Join(dir, "marshal_helpers.go"))
	if err != nil {
		log.Panic(err)
	}
}

func implementStringContained(f *File) {
	f.Func().Params(Id("r").Id("ContainedResource")).Id("String").Params().String().Block(
		List(Id("buf"), Id("err")).Op(":=").Qual("encoding/json", "MarshalIndent").Params(Id("r"), Lit(""), Lit("  ")),
		If(Id("err").Op("!=").Nil()).Block(
			Return(Lit("null")),
		),
		Return(Id("string").Params(Id("buf"))),
	)
}
