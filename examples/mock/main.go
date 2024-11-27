// This is a simple examples of a FHIR server that reads resources from another FHIR server.
package main

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/utils"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"fhir-toolbox/model"
	"fhir-toolbox/rest"
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
	server, err := rest.NewServer[model.R4](&backend, rest.DefaultConfig)
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	// Start the server and listen on port 80.
	log.Println("listening on http://localhost")
	log.Fatal(http.ListenAndServe(":80", server))
}

type mockBackend struct{}

func (b *mockBackend) ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError) {
	// forward single resource read to a search for the specific id
	result, err := b.SearchObservation(ctx, search.Options{
		Params: search.Params{
			"_id": search.AndList{search.OrList{{Value: id}}},
		},
		Count: 1,
	})
	if err != nil {
		return r4.Observation{}, err
	}

	if len(result.Resources) == 0 {
		return r4.Observation{}, capabilities.NotFoundError{ResourceType: "Observation", ID: id}
	}

	return *result.Resources[0].(*r4.Observation), nil
}

// SearchCapabilitiesObservation describes the search capabilities on the Observation resource.
func (b *mockBackend) SearchCapabilitiesObservation() search.Capabilities {
	return search.Capabilities{
		Params: map[string]search.ParamDesc{
			"_id": {Type: search.Token},
		},
	}
}

func (b *mockBackend) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	return search.Result{
		Resources: []model.Resource{
			r4.Observation{
				Id: &r4.Id{Value: utils.Ptr("123")},
				Meta: &r4.Meta{
					LastUpdated: &r4.Instant{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
				},
				Category: []r4.CodeableConcept{{
					Coding: []r4.Coding{{
						System:  &r4.Uri{Value: utils.Ptr("http://terminology.hl7.org/CodeSystem/observation-category")},
						Code:    &r4.Code{Value: utils.Ptr("vital-signs")},
						Display: &r4.String{Value: utils.Ptr("Vital Signs")},
					}},
				}},
				Code: r4.CodeableConcept{
					Coding: []r4.Coding{{
						System:  &r4.Uri{Value: utils.Ptr("http://loinc.org")},
						Code:    &r4.Code{Value: utils.Ptr("85354-9")},
						Display: &r4.String{Value: utils.Ptr("Blood pressure panel with all children optional")},
					}},
					Text: &r4.String{Value: utils.Ptr("Blood pressure systolic & diastolic")},
				},
				Effective: &r4.DateTime{Value: utils.Ptr(time.Now().AddDate(0, 0, -1).Format(time.RFC3339))},
				Issued:    &r4.Instant{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
				Status:    r4.Code{Value: utils.Ptr("final")},
				Component: []r4.ObservationComponent{
					{
						Code: r4.CodeableConcept{
							Coding: []r4.Coding{{
								System:  &r4.Uri{Value: utils.Ptr("http://loinc.org")},
								Code:    &r4.Code{Value: utils.Ptr("8480-6")},
								Display: &r4.String{Value: utils.Ptr("Systolic blood pressure")},
							}},
						},
						Value: &r4.Quantity{
							Value:  &r4.Decimal{Value: utils.Ptr("120.0")},
							Unit:   &r4.String{Value: utils.Ptr("mmHg")},
							System: &r4.Uri{Value: utils.Ptr("http://unitsofmeasure.org")},
							Code:   &r4.Code{Value: utils.Ptr("mm[Hg]")},
						},
					},
					{
						Code: r4.CodeableConcept{
							Coding: []r4.Coding{{
								System:  &r4.Uri{Value: utils.Ptr("http://loinc.org")},
								Code:    &r4.Code{Value: utils.Ptr("8462-4")},
								Display: &r4.String{Value: utils.Ptr("Diastolic blood pressure")},
							}},
						},
						Value: &r4.Quantity{
							Value:  &r4.Decimal{Value: utils.Ptr("600")},
							Unit:   &r4.String{Value: utils.Ptr("mmHg")},
							System: &r4.Uri{Value: utils.Ptr("http://unitsofmeasure.org")},
							Code:   &r4.Code{Value: utils.Ptr("mm[Hg]")},
						},
					},
				},
			},
		},
	}, nil
}

func (b *mockBackend) ReadComposition(ctx context.Context, id string) (r4.Composition, capabilities.FHIRError) {
	result, err := b.SearchComposition(ctx, search.Options{
		Params: search.Params{
			"_id": search.AndList{search.OrList{{Value: id}}},
		},
		Count: 1,
	})
	if err != nil {
		return r4.Composition{}, err
	}

	if len(result.Resources) == 0 {
		return r4.Composition{}, capabilities.NotFoundError{ResourceType: "Composition", ID: id}
	}

	return *result.Resources[0].(*r4.Composition), nil
}

// SearchCapabilitiesComposition describes the search capabilities on the Composition resource.
func (b *mockBackend) SearchCapabilitiesComposition() search.Capabilities {
	return search.Capabilities{
		Params: map[string]search.ParamDesc{
			"_id": {Type: search.Token},
		},
	}
}

func (b *mockBackend) SearchComposition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	return search.Result{
		Resources: []model.Resource{
			r4.Composition{
				Id: &r4.Id{Value: utils.Ptr("123")},
				Meta: &r4.Meta{
					LastUpdated: &r4.Instant{Value: utils.Ptr(time.Now().Format(time.RFC3339))},
				},
				Type: r4.CodeableConcept{
					Coding: []r4.Coding{{
						System:  &r4.Uri{Value: utils.Ptr("http://loinc.org")},
						Code:    &r4.Code{Value: utils.Ptr("11503-0")},
						Display: &r4.String{Value: utils.Ptr("Medical records")},
					}},
				},
				Title:  r4.String{Value: utils.Ptr("Test Composition")},
				Status: r4.Code{Value: utils.Ptr("final")},
				Date:   r4.DateTime{Value: utils.Ptr(time.Now().AddDate(0, 0, -1).Format(time.RFC3339))},
			},
		},
	}, nil
}
