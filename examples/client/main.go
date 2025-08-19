package main

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"io"
	"log"
)

func main() {
	client, err := rest.NewClientR4("https://server.fire.ly", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Read patient
	patient, err := client.ReadPatient(context.Background(), "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read patient:\n%s\n", patient)

	// Search for patients using typed search parameters
	result, err := client.SearchPatient(context.Background(),
		r4.PatientParams{
			Birthdate: search.String("ge2000-01-01"), // String-based (flexible)
			Gender:    search.Token{Code: "female"},  // Strongly-typed (safe)
		},
		search.Options{
			Count: 5,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d patients:\n%s\n", len(result.Resources), result)

	// Search for patient with pagination
	initialResult, err := client.SearchPatient(context.Background(),
		r4.PatientParams{
			Birthdate: search.String("ge2000-01-01"),
		},
		search.Options{
			Count: 5,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	iter := rest.Iterator(context.Background(), client, initialResult)
	pageNo := 0

	// Get 5 pages
	for pageNo < 5 {
		page, err := iter.Next()
		if err != nil {
			// io.EOF signals that there are no more pages.
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to fetch next page: %v", err)
		}

		// handle page
		fmt.Printf("Page %d:\n%s\n\n", pageNo, page)

		pageNo++
	}
}
