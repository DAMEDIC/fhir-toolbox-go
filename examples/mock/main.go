// Serves some mock Observations (read & search interactions).
package main

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"github.com/cockroachdb/apd/v3"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
)

func main() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	// Add some additional context to the log (like request id).
	requestContextHandler := rest.NewRequestContextSlogHandler(textHandler)
	slog.SetDefault(slog.New(requestContextHandler))

	// Create the mock backend that just returns some dummy data.
	backend := mockBackend{}

	// Create the REST server.
	// You can plug in any backend you want here.
	cfg := rest.DefaultConfig
	cfg.Base = &url.URL{Scheme: "http", Host: "localhost"}
	server, err := rest.NewServer[model.R5](&backend, cfg)
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	// Start the server and listen on port 80.
	log.Println("listening on http://localhost")
	log.Fatal(http.ListenAndServe(":80", server))
}

type mockBackend struct{}

func (b *mockBackend) ReadObservation(ctx context.Context, id string) (r5.Observation, error) {
	// forward single resource read to a search for the specific id
	result, err := b.SearchObservation(ctx, search.Options{
		Parameters: search.Parameters{
			search.ParameterKey{Name: "_id"}: search.All{search.OneOf{search.String(id)}},
		},
		Count: 1,
	})
	if err != nil {
		return r5.Observation{}, err
	}

	if len(result.Resources) == 0 {
		return r5.Observation{}, r5.OperationOutcome{
			Issue: []r5.OperationOutcomeIssue{
				{
					Severity:    r5.Code{Value: utils.Ptr("error")},
					Code:        r5.Code{Value: utils.Ptr("not-found")},
					Diagnostics: &r5.String{Value: utils.Ptr(fmt.Sprintf("Observation with ID %s not found", id))},
				},
			},
		}
	}

	return result.Resources[0], nil
}

// SearchCapabilitiesObservation describes the search capabilities on the Observation resource.
func (b *mockBackend) SearchCapabilitiesObservation(ctx context.Context) (search.Capabilities, error) {
	return search.Capabilities{
		Parameters: map[string]search.ParameterDescription{
			"_id": {Type: search.TypeToken},
		},
	}, nil
}

func (b *mockBackend) SearchObservation(ctx context.Context, options search.Options) (search.Result[r5.Observation], error) {
	return search.Result[r5.Observation]{
		Resources: []r5.Observation{
			r5.Observation{
				Id: &r5.Id{Value: utils.Ptr("123")},
				Meta: &r5.Meta{
					LastUpdated: &r5.Instant{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
				},
				Category: []r5.CodeableConcept{{
					Coding: []r5.Coding{{
						System:  &r5.Uri{Value: utils.Ptr("http://terminology.hl7.org/CodeSystem/observation-category")},
						Code:    &r5.Code{Value: utils.Ptr("vital-signs")},
						Display: &r5.String{Value: utils.Ptr("Vital Signs")},
					}},
				}},
				Code: r5.CodeableConcept{
					Coding: []r5.Coding{{
						System:  &r5.Uri{Value: utils.Ptr("http://loinc.org")},
						Code:    &r5.Code{Value: utils.Ptr("85354-9")},
						Display: &r5.String{Value: utils.Ptr("Blood pressure panel with all children optional")},
					}},
					Text: &r5.String{Value: utils.Ptr("Blood pressure systolic & diastolic")},
				},
				Effective: &r5.DateTime{Value: utils.Ptr(time.Now().AddDate(0, 0, -1).Format(time.RFC3339))},
				Issued:    &r5.Instant{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
				Status:    r5.Code{Value: utils.Ptr("final")},
				Component: []r5.ObservationComponent{
					{
						Code: r5.CodeableConcept{
							Coding: []r5.Coding{{
								System:  &r5.Uri{Value: utils.Ptr("http://loinc.org")},
								Code:    &r5.Code{Value: utils.Ptr("8480-6")},
								Display: &r5.String{Value: utils.Ptr("Systolic blood pressure")},
							}},
						},
						Value: &r5.Quantity{
							Value:  &r5.Decimal{Value: apd.New(120, 0)},
							Unit:   &r5.String{Value: utils.Ptr("mmHg")},
							System: &r5.Uri{Value: utils.Ptr("http://unitsofmeasure.org")},
							Code:   &r5.Code{Value: utils.Ptr("mm[Hg]")},
						},
					},
					{
						Code: r5.CodeableConcept{
							Coding: []r5.Coding{{
								System:  &r5.Uri{Value: utils.Ptr("http://loinc.org")},
								Code:    &r5.Code{Value: utils.Ptr("8462-4")},
								Display: &r5.String{Value: utils.Ptr("Diastolic blood pressure")},
							}},
						},
						Value: &r5.Quantity{
							Value:  &r5.Decimal{Value: apd.New(600, 0)},
							Unit:   &r5.String{Value: utils.Ptr("mmHg")},
							System: &r5.Uri{Value: utils.Ptr("http://unitsofmeasure.org")},
							Code:   &r5.Code{Value: utils.Ptr("mm[Hg]")},
						},
					},
				},
			},
		},
	}, nil
}

func (b *mockBackend) ReadComposition(ctx context.Context, id string) (r5.Composition, error) {
	result, err := b.SearchComposition(ctx, search.Options{
		Parameters: search.Parameters{
			search.ParameterKey{Name: "_id"}: search.All{search.OneOf{search.String(id)}},
		},
		Count: 1,
	})
	if err != nil {
		return r5.Composition{}, err
	}

	if len(result.Resources) == 0 {
		return r5.Composition{}, r5.OperationOutcome{
			Issue: []r5.OperationOutcomeIssue{
				{
					Severity:    r5.Code{Value: utils.Ptr("error")},
					Code:        r5.Code{Value: utils.Ptr("not-found")},
					Diagnostics: &r5.String{Value: utils.Ptr(fmt.Sprintf("Composition with ID %s not found", id))},
				},
			},
		}
	}

	return result.Resources[0], nil
}

// SearchCapabilitiesComposition describes the search capabilities on the Composition resource.
func (b *mockBackend) SearchCapabilitiesComposition(ctx context.Context) (search.Capabilities, error) {
	return search.Capabilities{
		Parameters: map[string]search.ParameterDescription{
			"_id": {Type: search.TypeToken},
		},
	}, nil
}

func (b *mockBackend) SearchComposition(ctx context.Context, options search.Options) (search.Result[r5.Composition], error) {
	return search.Result[r5.Composition]{
		Resources: []r5.Composition{
			r5.Composition{
				Id: &r5.Id{Value: utils.Ptr("123")},
				Meta: &r5.Meta{
					LastUpdated: &r5.Instant{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
				},
				Type: r5.CodeableConcept{
					Coding: []r5.Coding{{
						System:  &r5.Uri{Value: utils.Ptr("http://loinc.org")},
						Code:    &r5.Code{Value: utils.Ptr("11503-0")},
						Display: &r5.String{Value: utils.Ptr("Medical records")},
					}},
				},
				Title:  r5.String{Value: utils.Ptr("Test Composition")},
				Status: r5.Code{Value: utils.Ptr("final")},
				Date:   r5.DateTime{Value: utils.Ptr(time.Now().AddDate(0, 0, -1).Format(time.RFC3339))},
			},
		},
	}, nil
}
