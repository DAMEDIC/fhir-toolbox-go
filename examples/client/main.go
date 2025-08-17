package main

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"log"
	"time"

	"github.com/DAMEDIC/fhir-toolbox-go/rest"
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
	date, err := time.Parse(time.DateOnly, "2000-01-01")
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.SearchPatient(context.Background(), search.Options{
		Count: 5,
		Parameters: map[search.ParameterKey]search.AllOf{
			search.ParameterKey{Name: "birthdate"}: search.AllOf{search.OneOf{search.Date{Prefix: "ge", Value: date}}},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d patients\n%s\n", len(result.Resources), result)

	// If there's a next page, show the cursor
	if result.Next != "" {
		fmt.Printf("Next cursor: %s\n", result.Next)
	}
}
