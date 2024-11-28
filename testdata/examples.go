package testdata

import (
	"archive/zip"
	"io"
	"log"
	"strings"
)

func GetExamples(release, format string) map[string][]byte {
	downloadExamples(release, format)
	path := zipFilePath(release, format)

	log.Println("opening zip archive...")
	zip, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer zip.Close()

	examples := map[string][]byte{}
	for _, file := range zip.File {
		// not a FHIR resource
		if strings.HasSuffix(file.Name, "package-min-ver.json") ||
			strings.HasPrefix(file.Name, "__MACOSX/") ||
			file.Name == "examples-json/" {
			continue
		}

		f, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		name := strings.TrimPrefix(file.Name, "examples-json/")
		examples[name], err = io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
	}

	return examples
}
