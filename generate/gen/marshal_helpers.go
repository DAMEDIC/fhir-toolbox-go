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

	implementContainedResource(f)
	implementStringContained(f)
	json.ImplementMarshalContained(resources, f)
	json.ImplementUnmarshalContained(resources, f)
	xml.ImplementMarshalContained(f)
	xml.ImplementUnmarshalContained(resources, f)

	implementPrimitiveElement(f)
	json.ImplementMarshalPrimitiveElement(f)
	json.ImplementUnmarshalPrimitiveElement(f)

	err := f.Save(filepath.Join(dir, "marshal_helpers.go"))
	if err != nil {
		log.Panic(err)
	}
}

func implementContainedResource(f *File) *Statement {
	return f.Type().Id("ContainedResource").Struct(
		Qual("fhir-toolbox/model", "Resource"),
	)
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

func implementPrimitiveElement(f *File) *Statement {
	return f.Type().Id("primitiveElement").Struct(
		Id("Id").Id("*string"),
		Id("Extension").Index().Id("Extension"),
	)
}
