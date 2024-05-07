package testdata

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const examplesURL = "http://hl7.org/fhir/%v/%v"

func formatFilename(format string) string {
	filename := "examples"
	if format == "json" {
		filename = "examples-json"
	}
	return filename + ".zip"
}

func formatDownloadURL(release, format string) string {
	return fmt.Sprintf(examplesURL, release, formatFilename(format))
}

func zipFilePath(release, format string) string {
	dir := filepath.Join("build", "fhir", strings.ToLower(release))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	return filepath.Join(dir, formatFilename(format))
}

func downloadExamples(release, format string) string {
	zip := zipFilePath(release, format)

	if _, err := os.Stat(zip); err == nil {
		log.Printf("Download skipped! %s examples already present.\n", release)
		return zip
	}

	url := formatDownloadURL(release, format)

	log.Printf("Downloading %s examples...\n", release)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error downloading %s examples", release)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error reading response body")
	}

	err = os.WriteFile(zip, bytes, 0644)
	if err != nil {
		log.Panicf("Error writing to \"%s\"", zip)
	}

	log.Println("Download successful!")

	return zip
}
