package main

import (
	"fhir-toolbox/fhirpath"
	"fmt"
)

func main() {
	fmt.Println(fhirpath.Evaluate(object{"Patient"}, fhirpath.MustParse("Patient.name.given")))
}

type object struct {
	name string
}

func (object) Member(s string) fhirpath.Collection {
	fmt.Println("get member", s)
	return fhirpath.Collection{object{name: s}}
}
