package fhirpath

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	. "github.com/dave/jennifer/jen"
	"strings"
)

const fhirpathModuleName = "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"

type FHIRPathGenerator struct{}

func (g FHIRPathGenerator) GenerateType(f *File, rt ir.ResourceOrType) bool {
	for _, s := range rt.Structs {
		generateTypeRecv(f, s)
	}

	return true
}

func generateTypeRecv(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("Type").Params().BlockFunc(func(g *Group) {
		g.Return()
		generateType(g.)
	})
}

func (g FHIRPathGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	contextFile := f("fhirpath", strings.ToLower(release))

	generateContextFunc(contextFile)
	generateWithContext(contextFile)
	generateTypes(contextFile, rt)
	generateFunctions(contextFile)
}

func generateContextFunc(f *File) *Statement {
	return f.Func().Id("Context").Params().Qual("context", "Context").Block(
		Return(
			Id("WithContext").Call(Qual("context", "Background").Call()),
		),
	)
}

func generateWithContext(contextFile *File) *Statement {
	return contextFile.Func().Id("WithContext").Params(
		Id("ctx").Qual("context", "Context"),
	).Qual("context", "Context").Block(
		Id("ctx").Op("=").Qual(fhirpathModuleName, "WithNamespace").Call(
			Id("ctx"),
			Lit("FHIR"),
		),
		Id("ctx").Op("=").Qual(fhirpathModuleName, "WithTypes").Call(
			Id("ctx"),
			Id("allFHIRPathTypes"),
		),
		Id("ctx").Op("=").Qual(fhirpathModuleName, "WithFunctions").Call(
			Id("ctx"),
			Id("fhirFunctions"),
		),
		Return(Id("ctx")),
	)
}

func generateFunctions(f *File) *Statement {
	return f.Var().Id("fhirFunctions").Op("=").Qual(fhirpathModuleName, "Functions").Values()
}
