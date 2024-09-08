// This is a simple examples of a FHIR server that reads resources from another FHIR server.
package main

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/utils"
	"log"
	"net/http"
	"time"

	"fhir-toolbox/model"
	"fhir-toolbox/rest"
)

// 1. Define our backend
type demoBackend struct{}

// 2. Implement your desired capabilities (interactions)
func (b *demoBackend) ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError) {
	return r4.Observation{
		Id:     &r4.Id{Value: &id},
		Status: r4.Code{Value: utils.Ptr("final")},
		Code: r4.CodeableConcept{
			Coding: []r4.Coding{{
				System:  &r4.Uri{Value: utils.Ptr("http://loinc.org")},
				Code:    &r4.Code{Value: utils.Ptr("8480-6")},
				Display: &r4.String{Value: utils.Ptr("Systolic blood pressure")},
			}},
		},
		Effective: &r4.DateTime{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
		Value:     &r4.Quantity{Value: &r4.Decimal{Value: utils.Ptr("120.0")}, Unit: &r4.String{Value: utils.Ptr("mmHg")}},
	}, nil
}

func main() {
	// 3. Instantiate your backend
	backend := demoBackend{}

	// 4. Start your server!
	server, err := rest.NewServer[model.R4](&backend, rest.DefaultConfig)
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	log.Println("listening on http://localhost")
	log.Fatal(http.ListenAndServe(":80", server))

	// 5. Visit http://localhost/Observation/1234
}
