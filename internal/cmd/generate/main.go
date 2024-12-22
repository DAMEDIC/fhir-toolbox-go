package main

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/json"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/xml"
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

		generate.GenerateAll(all, genDir(modelGenTarget, r), r,
			generate.ModelPkgDocGenerator{},
			generate.TypesGenerator{},
			generate.ImplResourceGenerator{},
			generate.ImplElementGenerator{},
			generate.StringerGenerator{},
			json.MarshalGenerator{},
			json.UnmarshalGenerator{},
			xml.MarshalGenerator{},
			xml.UnmarshalGenerator{},
		)

		generate.GenerateAll(all, genDir(capabilitiesGenTarget, r), r,
			generate.CapabilityPkgDocGenerator{},
			generate.CapabilitiesGenerator{},
			generate.CapabilitiesWrapperGenerator{},
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
