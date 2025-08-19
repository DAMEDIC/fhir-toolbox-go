package main

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"log"
)

func main() {
	client, err := rest.NewClientR4("https://server.fire.ly", nil)
	if err != nil {
		log.Fatal(err)
	}

	patient, err := client.ReadPatient(context.Background(), "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read patient:\n%s\n", patient)

	// Search for patients
	result, err := client.SearchPatient(context.Background(),
		search.Params{
			"birthdate": search.String("ge2000-01-01"),
		},
		search.Options{
			Count: 5,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d patients\n%s\n", len(result.Resources), result)

	// If there's a next page, show the cursor
	if result.Next != "" {
		fmt.Printf("Next cursor: %s\n", result.Next)
	}
}
