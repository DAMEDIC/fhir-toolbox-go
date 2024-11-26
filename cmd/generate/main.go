package main

import (
	"fhir-toolbox/internal/generator"
	"fhir-toolbox/internal/generator/ir"
	"fhir-toolbox/internal/generator/json"
	"fhir-toolbox/internal/generator/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	buildReleases         = []string{"R4"}
	definitionsURLFmtStr  = "http://hl7.org/fhir/%s/definitions.json.zip"
	modelGenTarget        = "model/gen"
	capabilitiesGenTarget = "capabilities/gen"
	wrapperGenTarget      = "capabilities/wrap/gen"
)

func main() {
	fmt.Println("Running code generation...")
	if err := os.RemoveAll(modelGenTarget); err != nil {
		panic(err)
	}
	if err := os.RemoveAll(capabilitiesGenTarget); err != nil {
		panic(err)
	}
	if err := os.RemoveAll(wrapperGenTarget); err != nil {
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
			generator.TypesGenerator{},
			generator.ImplResourceGenerator{},
			generator.StringerGenerator{},
			json.MarshalGenerator{},
			json.UnmarshalGenerator{},
			xml.MarshalGenerator{},
			xml.UnmarshalGenerator{},
		)

		generator.GenerateAll(all, genDir(capabilitiesGenTarget, r), r,
			generator.CapabilitiesGenerator{},
		)
		generator.GenerateAll(all, genDir(wrapperGenTarget, r), r,
			generator.WrapperGenerator{},
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
