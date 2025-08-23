// Short demo that serves a mock Observation resource (only read interaction).
package main

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"github.com/cockroachdb/apd/v3"
	"log"
	"net/http"
	"time"

	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
)

// 1. Define our backend
type demoBackend struct{}

// 2. Implement CapabilityBase
func (b *demoBackend) CapabilityBase(ctx context.Context) (basic.CapabilityStatement, error) {
	return basic.CapabilityStatement{
		Status:      basic.Code{Value: ptr.To("active")},
		Date:        basic.DateTime{Value: ptr.To(time.Now().Format(time.RFC3339))},
		Kind:        basic.Code{Value: ptr.To("instance")},
		FhirVersion: basic.Code{Value: ptr.To("5.0")},
		Format: []basic.Code{
			{Value: ptr.To("xml")},
			{Value: ptr.To("json")},
		},
		Software: &basic.CapabilityStatementSoftware{
			Name:    basic.String{Value: ptr.To("fhir-toolbox-go demo")},
			Version: &basic.String{Value: ptr.To("1.0.0")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: ptr.To("Demo FHIR server built with fhir-toolbox-go")},
			Url:         &basic.Url{Value: ptr.To("http://localhost")},
		},
	}, nil
}

// 3. Implement your desired capabilities (interactions)
func (b *demoBackend) ReadObservation(ctx context.Context, id string) (r5.Observation, error) {
	return r5.Observation{
		Id:     &r5.Id{Value: &id},
		Status: r5.Code{Value: ptr.To("final")},
		Code: r5.CodeableConcept{
			Coding: []r5.Coding{{
				System:  &r5.Uri{Value: ptr.To("http://loinc.org")},
				Code:    &r5.Code{Value: ptr.To("8480-6")},
				Display: &r5.String{Value: ptr.To("Systolic blood pressure")},
			}},
		},
		Effective: &r5.DateTime{Value: ptr.To(time.Now().Format(time.RFC3339))},
		Value:     &r5.Quantity{Value: &r5.Decimal{Value: apd.New(120, 0)}, Unit: &r5.String{Value: ptr.To("mmHg")}},
	}, nil
}

func main() {
	// 4. Instantiate your backend
	backend := demoBackend{}

	// 5. Start your server!
	server := &rest.Server[model.R5]{
		Backend: &backend,
	}

	log.Println("listening on http://localhost")
	log.Fatal(http.ListenAndServe(":80", server))

	// 6. Visit http://localhost/Observation/1234
}
