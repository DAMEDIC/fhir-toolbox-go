package gen

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func GenerateInterfaces(genTarget, release string) {
	dir := filepath.Join(genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	f := NewFilePath(dir)

	f.Type().Id("Resource").Interface()
	f.Type().Id("Element").Interface()

	err = f.Save(filepath.Join(dir, "interfaces.go"))
	if err != nil {
		log.Panic(err)
	}
}
