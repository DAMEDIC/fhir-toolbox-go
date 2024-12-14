package main

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generator"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generator/ir"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generator/json"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generator/xml"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	buildReleases         = []string{"R4", "R4B", "R5"}
	definitionsURLFmtStr  = "http://hl7.org/fhir/%s/definitions.json.zip"
	modelGenTarget        = "model/gen"
	capabilitiesGenTarget = "capabilities/gen"
)

func main() {
	fmt.Println("Running code generation...")
	if err := os.RemoveAll(modelGenTarget); err != nil {
		panic(err)
	}
	if err := os.RemoveAll(capabilitiesGenTarget); err != nil {
		panic(err)
	}

	for _, r := range buildReleases {
		log.Printf("Generating for FHIR %v ...\n", r)

		zipPath := downloadDefinitions(r)
		bundles := readJSONFromZIP(zipPath)

		types := ir.Parse(&bundles.types)
		resources := ir.Parse(&bundles.resources)
		all := append(types, resources...)

		log.Println("Generating structs and implementations...")

		generator.GenerateAll(all, genDir(modelGenTarget, r), r,
			generator.ModelPkgDocGenerator{},
			generator.TypesGenerator{},
			generator.ImplResourceGenerator{},
			generator.ImplElementGenerator{},
			generator.StringerGenerator{},
			json.MarshalGenerator{},
			json.UnmarshalGenerator{},
			xml.MarshalGenerator{},
			xml.UnmarshalGenerator{},
		)

		generator.GenerateAll(all, genDir(capabilitiesGenTarget, r), r,
			generator.CapabilityPkgDocGenerator{},
			generator.CapabilitiesGenerator{},
			generator.CapabilitiesWrapperGenerator{},
		)
	}

	log.Println("Code generation done.")
}

func genDir(genTarget, release string) string {
	dir := filepath.Join(genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	return dir
}
