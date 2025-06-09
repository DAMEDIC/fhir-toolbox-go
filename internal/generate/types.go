package generate

import (
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type TypesGenerator struct {
	NoOpGenerator
}

func (g TypesGenerator) GenerateType(f *File, rt ir.ResourceOrType) bool {
	for _, t := range rt.Structs {
		generateStruct(f, t)
		generatePrimitiveEnums(f, t)
	}
	return true
}

func (g TypesGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	implementContainedResource(f("contained_resource", strings.ToLower(release)))

}

func generateStruct(f *File, s ir.Struct) {
	if s.Name == "Decimal" {
		f.ImportName("github.com/cockroachdb/apd/v3", "apd")
	}

	for _, line := range strings.Split(s.DocComment, "\n") {
		f.Comment(line)
	}

	f.Type().Id(s.Name).StructFunc(func(g *Group) {
		for _, f := range s.Fields {
			for _, line := range strings.Split(f.DocComment, "\n") {
				g.Comment(line)
			}

			stmt := g.Id(f.Name)

			if f.Polymorph {
				stmt.Id("is" + s.Name + f.Name)
			} else {
				t := f.PossibleTypes[0]

				if f.Multiple {
					stmt.Index()
				} else if f.Optional && !f.Polymorph && !t.IsNestedResource {
					stmt.Op("*")
				}

				if t.IsNestedResource {
					stmt.Qual(moduleName+"/model", "Resource")
				} else if s.Name == "Decimal" && f.Name == "Value" {
					stmt.Qual("github.com/cockroachdb/apd/v3", "Decimal")
				} else {
					stmt.Id(t.Name)
				}
			}
		}
	})
}

func generatePrimitiveEnums(f *File, s ir.Struct) {
	for _, sf := range s.Fields {
		if !sf.Polymorph {
			continue
		}

		f.Type().Id("is"+s.Name+sf.Name).Interface(
			Qual(moduleName+"/model", "Element"),
			Id("is"+s.Name+sf.Name).Params(),
		)

		for _, t := range sf.PossibleTypes {
			f.Func().Params(Id("r").Id(t.Name)).Id("is" + s.Name + sf.Name).Params().Block()
		}
	}
}

func implementContainedResource(f *File) *Statement {
	return f.Type().Id("ContainedResource").Struct(
		Qual(moduleName+"/model", "Resource"),
	)
}
