package main

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	r4 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
)

func main() {
	patient := r4.Patient{
		Name: []r4.HumanName{{
			Given:  []r4.String{{Value: utils.Ptr("Donald")}},
			Family: &r4.String{Value: utils.Ptr("Duck")},
		}},
		MultipleBirth: r4.Integer{Value: utils.Ptr[int32](3)},
	}

	fmt.Println(fhirpath.Evaluate(r4.Context(), patient, fhirpath.MustParse("Patient.name.given")))
	fmt.Println(fhirpath.Evaluate(r4.Context(), patient, fhirpath.MustParse("Patient.multipleBirth")))
}
