package main

import (
	"fmt"
	"os"

	"fhir-toolbox/fake"
)

func main() {
	var backendUrl = os.Args[1]
	fmt.Println("FHIR Server:", backendUrl)

	f := fake.Fake{}

	fmt.Printf("%v\n", f.GetPatient())
	fmt.Printf("%v\n", f.GetObservtion())
}
