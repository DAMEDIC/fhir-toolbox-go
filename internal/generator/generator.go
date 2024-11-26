// Package generator contains the code generation for this module.
//
// By implementing the visitor pattern, the api allows different generators to
// emit source code into the same source files.
package generator

import (
	"fhir-toolbox/internal/generator/ir"
	"github.com/dave/jennifer/jen"
	"log"
	"os"
	"path/filepath"
)

// The Generator can be used to implement code generation.
type Generator interface {
	// GenerateType is visited for each resource or data type.
	GenerateType(f *jen.File, rt ir.ResourceOrType) bool
	// GenerateAdditional allows the generation of additional
	//source code that needs access to multiple or all resources or types.
	GenerateAdditional(f func(fileName string, pkgName string) *jen.File, release string, rt []ir.ResourceOrType)
}

// NoOpGenerator can be embedded when only some visitor functions need to be implemented.
type NoOpGenerator struct{}

func (g NoOpGenerator) GenerateType(f *jen.File, rt ir.ResourceOrType) bool {
	return false
}
func (g NoOpGenerator) GenerateAdditional(f func(fileName string, pkgName string) *jen.File, release string, rt []ir.ResourceOrType) {
}

// GenerateAll into specified target dir using given generators.
func GenerateAll(types []ir.ResourceOrType, targetDir, release string, generators ...Generator) {
	for _, t := range types {
		f := jen.NewFilePath(targetDir)
		writeFile := false

		for _, g := range generators {
			if g.GenerateType(f, t) {
				writeFile = true
			}
		}

		if writeFile {
			err := f.Save(filepath.Join(targetDir, t.FileName+".go"))
			if err != nil {
				log.Panic(err)
			}
		}
	}

	files := map[string]*jen.File{}

	fileFn := func(fileName, pkgName string) *jen.File {
		_, ok := files[fileName]
		if !ok {
			files[fileName] = jen.NewFilePathName(targetDir, pkgName)
		}
		return files[fileName]
	}

	for _, g := range generators {
		g.GenerateAdditional(fileFn, release, types)
	}

	for k, f := range files {
		file := filepath.Join(targetDir, k+"go")
		dir := filepath.Dir(file)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
		err = f.Save(filepath.Join(targetDir, k+".go"))
		if err != nil {
			log.Panic(err)
		}
	}
}
