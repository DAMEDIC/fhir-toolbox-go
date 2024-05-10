package main

import (
	"context"
	"fmt"
	"os"

	"fhir-toolbox/backend/fake"
)

func main() {
	var backendUrl = os.Args[1]
	fmt.Println("FHIR Server:", backendUrl)

	f := fake.Fake{}

	fmt.Println(f.ReadPatient(context.TODO(), "1234"))
	fmt.Println(f.ReadObservation(context.TODO(), "1234"))
}
