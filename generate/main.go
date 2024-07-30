package main

import (
	"fhir-toolbox/generate/gen"
	"fhir-toolbox/generate/ir"
	"fmt"
	"log"
	"os"
)

var (
	buildReleases         = []string{"R4"}
	definitionsURLFmtStr  = "http://hl7.org/fhir/%s/definitions.json.zip"
	modelGenTarget        = "model/gen"
	capabilitiesGenTarget = "capabilities/gen"
	genericGenTarget      = "generic/gen"
)

func main() {
	fmt.Println("Running code generation...")
	os.RemoveAll(modelGenTarget)
	os.RemoveAll(capabilitiesGenTarget)
	os.RemoveAll(genericGenTarget)

	for _, r := range buildReleases {
		log.Printf("Generating for FHIR %v ...\n", r)

		zipPath := downloadDefinitions(r)
		bundles := readJSONFromZIP(zipPath)

		typesSourceFiles := ir.Parse(&bundles.types)
		resourcesSourceFiles := ir.Parse(&bundles.resources)
		allResourcesStructs := ir.FilterResources(resourcesSourceFiles)

		gen.GenerateTypes(typesSourceFiles, modelGenTarget, r)
		gen.GenerateResources(resourcesSourceFiles, modelGenTarget, r)
		gen.GenerateMarshalHelpers(allResourcesStructs, modelGenTarget, r)

		gen.GenerateCapabilityInterfaces(allResourcesStructs, capabilitiesGenTarget, r)
		gen.GenerateGeneric(allResourcesStructs, genericGenTarget, r)
	}

	log.Println("Code generation done.")
}
