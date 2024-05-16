package gen

import (
	"fhir-toolbox/generate/ir"
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func genDir(genTarget, release string) string {
	dir := filepath.Join(genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	return dir
}

func GenerateTypes(files []ir.SourceFile, genTarget, release string) {
	generateTypes(files, genTarget, release, "types")
}

func GenerateResources(files []ir.SourceFile, genTarget, release string) {
	generateTypes(files, genTarget, release, "resources")
}

func generateTypes(files []ir.SourceFile, genTarget, release, pkg string) {
	log.Printf("Generating %s...\n", pkg)

	dir := genDir(genTarget, release)

	for _, sf := range files {
		f := NewFilePath(dir)

		for _, s := range sf.Structs {
			generateStruct(f, *s)
			implementResource(f, *s)
			generatePrimitiveEnums(f, *s)

			generateMarshalJSONStruct(f, *s)

			implementMarshalJSON(f, *s)
			implementUnmarshalJSON(f, *s)

			if !s.IsPrimitive {
				implementString(f, *s)
			}
		}

		err := f.Save(filepath.Join(dir, sf.Name+".go"))
		if err != nil {
			log.Panic(err)
		}
	}
}

func generateStruct(f *File, s ir.Struct) {
	for _, line := range strings.Split(s.DocComment, "\n") {
		f.Comment(line)
	}

	f.Type().Id(s.Name).StructFunc(func(g *Group) {
		for _, f := range s.Fields {
			for _, line := range strings.Split(f.DocComment, "\n") {
				g.Comment(line)
			}

			stmt := g.Id(f.Name)
			if f.Multiple {
				stmt.Index()
			} else if f.Optional && !f.Polymorph {
				stmt.Op("*")
			}

			if f.Polymorph {
				stmt.Id("is" + s.Name + f.Name)
			} else {
				t := f.PossibleTypes[0]

				if t.IsNestedResource {
					stmt.Qual("fhir-toolbox/model", "Resource")
				} else {
					stmt.Id(t.Name)
				}
			}
		}
	})
}

func implementResource(f *File, s ir.Struct) {
	if s.IsResource {
		f.Func().Params(Id("r").Id(s.Name)).Id("ResourceType").Params().String().Block(
			Return(Lit(s.Name)),
		)
		f.Func().Params(Id("r").Id(s.Name)).Id("ResourceId").Params().Params(String(), Bool()).Block(
			If(Id("r.Id").Op("==").Nil()).Block(
				Return(Lit(""), False()),
			),
			If(Id("r.Id.Id").Op("==").Nil()).Block(
				Return(Lit(""), False()),
			),
			Return(Id("*r.Id.Id"), True()),
		)
	}
}

func generatePrimitiveEnums(f *File, s ir.Struct) {
	for _, sf := range s.Fields {
		if !sf.Polymorph {
			continue
		}

		f.Type().Id("is" + s.Name + sf.Name).Interface(
			Id("is" + s.Name + sf.Name).Params(),
		)

		for _, t := range sf.PossibleTypes {
			f.Func().Params(Id("r").Id(t.Name)).Id("is" + s.Name + sf.Name).Params().Block()
		}
	}
}

func implementString(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("String").Params().String().Block(
		List(Id("buf"), Id("err")).Op(":=").Qual("encoding/json", "MarshalIndent").Params(Id("r"), Lit(""), Lit("  ")),
		If(Id("err").Op("!=").Nil()).Block(
			Panic(Id("err")),
		),
		Return(Id("string").Params(Id("buf"))),
	)
}
