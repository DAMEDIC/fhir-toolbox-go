package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func zipFilePath(release string) string {
	dir := filepath.Join("build", "fhir", strings.ToLower(release))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	return filepath.Join(dir, "definitions.json.zip")
}

func formatDownloadURL(release string) string {
	return fmt.Sprintf(definitionsURLFmtStr, release)
}

func downloadDefinitions(release string) string {
	zip := zipFilePath(release)

	if _, err := os.Stat(zip); err == nil {
		log.Printf("Download skipped! %s FHIR definitions already present.\n", release)
		return zip
	}

	url := formatDownloadURL(release)

	log.Printf("Downloading %s FHIR definitions...\n", release)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error downloading %s definitions", release)
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
