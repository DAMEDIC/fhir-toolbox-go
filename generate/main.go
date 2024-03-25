package main

import (
	"fhir-toolbox/generate/gen"
	"fhir-toolbox/generate/ir"
	"fmt"
	"log"
)

var (
	buildReleases        = []string{"R4"}
	definitionsURLFmtStr = "http://hl7.org/fhir/%s/definitions.json.zip"
	modelGenTarget       = "model/gen"
)

func main() {
	fmt.Println("Running code generation...")
	gen.CleanGeneratedFiles(modelGenTarget)

	for _, r := range buildReleases {
		log.Printf("Generating for FHIR %v ...\n", r)

		zipPath := downloadDefinitions(r)
		bundles := readJSONFromZIP(zipPath)

		typesIR := ir.Parse(&bundles.types)
		resourcesIR := ir.Parse(&bundles.resources)

		gen.GenerateInterfaces(modelGenTarget, r)
		gen.GenerateMarshalHelpers(ir.FilterResources(resourcesIR), modelGenTarget, r)
		gen.GenerateTypes(typesIR, modelGenTarget, r)
		gen.GenerateResources(resourcesIR, modelGenTarget, r)
	}

	log.Println("Code generation done.")
}
