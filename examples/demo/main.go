// Short demo that serves a mock Observation resource (only read interaction).
package main

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"github.com/cockroachdb/apd/v3"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
)

// 1. Define our backend
type demoBackend struct{}

// 2. Implement your desired capabilities (interactions)
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
	// 3. Instantiate your backend
	backend := demoBackend{}

	// 4. Start your server!
	cfg := rest.DefaultConfig
	cfg.Base = &url.URL{Scheme: "http", Host: "localhost"}
	server, err := rest.NewServer[model.R5](&backend, cfg)
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	log.Println("listening on http://localhost")
	log.Fatal(http.ListenAndServe(":80", server))

	// 5. Visit http://localhost/Observation/1234
}
