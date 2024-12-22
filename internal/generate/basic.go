package generate

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type BasicDocGenerator struct {
	NoOpGenerator
}

func (g BasicDocGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	file := f("doc", strings.ToLower(release))
	file.PackageComment(fmt.Sprintf("Package %s provides basic resources that are valid across FHIR versions.", strings.ToLower(release)))
}
