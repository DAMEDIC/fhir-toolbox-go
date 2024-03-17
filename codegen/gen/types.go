package gen

import (
	"fhir-toolbox/codegen/ir"
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func GenerateTypes(files []ir.SourceFile, genTarget, release string) {
	generate(files, genTarget, release, "types")
}

func GenerateResources(files []ir.SourceFile, genTarget, release string) {
	generate(files, genTarget, release, "resources")
}

func generate(files []ir.SourceFile, genTarget, release, pkg string) {
	log.Printf("Generating %s...\n", pkg)

	dir := filepath.Join(genTarget, strings.ToLower(release), pkg)
	importPrefix := filepath.Join("fhir-toolbox", genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	for _, sf := range files {
		f := NewFilePath(dir)

		for _, s := range sf.Structs {
			generateStruct(f, s, importPrefix, pkg)
			implementString(f, s)
		}

		err = f.Save(filepath.Join(dir, sf.Name+".go"))
		if err != nil {
			log.Panic(err)
		}
	}
}

func generateStruct(f *File, s ir.Struct, importPrefix, pkg string) {
	for _, line := range strings.Split(s.DocComment, "\n") {
		f.Comment(line)
	}

	f.Type().Id(s.Name).StructFunc(func(g *Group) {
		for _, f := range s.Fields {
			for _, line := range strings.Split(f.DocComment, "\n") {
				g.Comment(line)
			}

			s := g.Id(f.Name)
			if f.Multiple {
				s = s.Index()
			} else if f.Optional && f.Type.Name != "Element" {
				s = s.Op("*")
			}

			if f.Type.Package == "" || f.Type.Package == pkg {
				s = s.Id(f.Type.Name)
			} else if f.Type.Package == ".." {
				s = s.Qual(importPrefix, f.Type.Name)
			} else {
				s = s.Qual(importPrefix+"/"+f.Type.Package, f.Type.Name)
			}
		}
	})
}

func implementString(f *File, s ir.Struct) {
	f.Func().Params(Id("s").Id(s.Name)).Id("String").Params().String().Block(
		List(Id("buf"), Id("err")).Op(":=").Qual("encoding/json", "MarshalIndent").Params(Id("s"), Lit(""), Lit("  ")),
		If(Id("err").Op("!=").Nil()).Block(
			Panic(Id("err")),
		),
		Return(Id("string").Params(Id("buf"))),
	)
}
