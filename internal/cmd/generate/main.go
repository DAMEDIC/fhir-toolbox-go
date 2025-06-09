package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/ir"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/json"
	"github.com/DAMEDIC/fhir-toolbox-go/internal/generate/xml"
)

var (
	buildReleases         = []string{"R4", "R4B", "R5"}
	definitionsURLFmtStr  = "http://hl7.org/fhir/%s/definitions.json.zip"
	modelGenTarget        = "model/gen"
	capabilitiesGenTarget = "capabilities/gen"
	basicResources        = []string{"OperationOutcome", "Bundle", "CapabilityStatement"}
)

func main() {
	log.Println("Running code generation...")
	if err := os.RemoveAll(modelGenTarget); err != nil {
		panic(err)
	}
	if err := os.RemoveAll(capabilitiesGenTarget); err != nil {
		panic(err)
	}

	releaseTypes := loadTypes(buildReleases)
	basicTypes := collectBasicTypes(releaseTypes)

	for _, r := range buildReleases {
		log.Printf("Generating for FHIR %v ...\n", r)

		log.Println("Generating structs and implementations...")

		generate.GenerateAll(releaseTypes[r], genDir(modelGenTarget, r), r,
			generate.ModelPkgDocGenerator{},
			generate.TypesGenerator{},
			generate.ImplResourceGenerator{},
			generate.ImplElementGenerator{},
			generate.StringerGenerator{},
			generate.OperationOutcomeErrorGenerator{},
			json.MarshalGenerator{},
			json.UnmarshalGenerator{},
			xml.MarshalGenerator{},
			xml.UnmarshalGenerator{},
			fhirpath.FHIRPathGenerator{},
		)

		generate.GenerateAll(releaseTypes[r], genDir(capabilitiesGenTarget, r), r,
			generate.CapabilityPkgDocGenerator{},
			generate.CapabilitiesGenerator{},
			generate.CapabilitiesWrapperGenerator{},
		)
	}

	log.Println("Generating basic types...")

	generate.GenerateAll(basicTypes, genDir(modelGenTarget, "basic"), "basic",
		generate.BasicDocGenerator{},
		generate.TypesGenerator{},
		generate.ImplResourceGenerator{},
		generate.ImplElementGenerator{},
		generate.StringerGenerator{},
		generate.OperationOutcomeErrorGenerator{},
		json.MarshalGenerator{NotUseContainedResource: true},
		xml.MarshalGenerator{},
		fhirpath.FHIRPathGenerator{},
	)

	log.Println("Code generation done.")
}

func loadTypes(releases []string) map[string][]ir.ResourceOrType {
	types := map[string][]ir.ResourceOrType{}
	for _, r := range releases {
		zipPath := downloadDefinitions(r)
		bundles := readJSONFromZIP(zipPath)

		types[r] = append(ir.Parse(&bundles.types), ir.Parse(&bundles.resources)...)
	}
	return types
}

func collectBasicTypes(releaseTypes map[string][]ir.ResourceOrType) []ir.ResourceOrType {
	var basicTypes []ir.ResourceOrType
	for _, r4t := range releaseTypes["R4"] {
		if !slices.Contains(basicResources, r4t.Name) && r4t.IsResource {
			continue
		}

		var other []ir.ResourceOrType
		for k, v := range releaseTypes {
			if k == "R4" {
				continue
			}

			for _, ot := range v {
				if ot.Name == r4t.Name {
					other = append(other, ot)
				}
			}
		}
		basicTypes = append(basicTypes, ir.CollectBasic(r4t, other))
	}

	return basicTypes
}

func genDir(genTarget, release string) string {
	dir := filepath.Join(genTarget, strings.ToLower(release))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	return dir
}
