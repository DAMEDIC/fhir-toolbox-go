// A FHIR server that serves resources from another FHIR server.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
)

// For implementing generic APIs, it is good practice to define a common interface for the client.
// This groups all the capabilities that are supported.
type myClient interface {
	capabilities.GenericCapabilities
	capabilities.GenericRead
	capabilities.GenericSearch
}

func main() {
	var backendUrl = os.Args[1]

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	// Add some additional context to the log (like request id).
	requestContextHandler := rest.NewRequestContextSlogHandler(textHandler)
	slog.SetDefault(slog.New(requestContextHandler))

	log.Printf("FHIR Server: %s", backendUrl)

	// Create a client to the backing server.
	//
	// The client is implemented below and serves only as an example.
	// A full-featured client might be in the scope of the fhir-toolbox-go module and will be added eventually.
	// The example client only supports the read and search operations and exposes them using the generic API.
	var genericClient myClient = &Client{
		url: strings.TrimRight(backendUrl, "/"),
	}

	// Create the REST server.
	// You can plug in any backend you want here.
	cfg := rest.DefaultConfig
	server, err := rest.NewServer[model.R5](genericClient, cfg)
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	// Start the server and listen on port 80.
	log.Println("listening on http://localhost")
	log.Fatal(http.ListenAndServe(":80", server))
}

type Client struct {
	url string
}

func (c *Client) CapabilityStatement(ctx context.Context) (basic.CapabilityStatement, error) {
	// TODO: this should call the remote service for its CapabilityStatement
	// For now, return a basic CapabilityStatement with base URL for canonical references

	return basic.CapabilityStatement{
		Status: basic.Code{Value: ptr.To("active")},
		Date:   basic.DateTime{Value: ptr.To(time.Now().Format(time.RFC3339))},
		Kind:   basic.Code{Value: ptr.To("instance")},
		Software: &basic.CapabilityStatementSoftware{
			Name: basic.String{Value: ptr.To("fhir-toolbox-go-proxy")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: ptr.To("a FHIR proxy built with fhir-toolbox-go")},
			Url:         &basic.Url{Value: ptr.To("http://localhost")},
		},
		FhirVersion: basic.Code{Value: ptr.To("5.0.0")},
		Format:      []basic.Code{{Value: ptr.To("xml")}, {Value: ptr.To("json")}},
		Rest:        []basic.CapabilityStatementRest{},
	}, nil
}

func (c *Client) Read(ctx context.Context, resourceType string, id string) (model.Resource, error) {
	// ContainedResource is a concrete representation of any resource
	// internally this e.g. uses in bundles
	var resource r5.ContainedResource

	url := fmt.Sprintf("%s/%s/%s", c.url, resourceType, id)
	log.Printf("forwarding GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		// OperationOutcome implement the Go error interface.
		// The server will return an appropriate HTTP status code and the OperationOutcome as the response body.
		return nil, r5.OperationOutcome{
			Issue: []r5.OperationOutcomeIssue{
				{
					Severity:    r5.Code{Value: ptr.To("error")},
					Code:        r5.Code{Value: ptr.To("exception")},
					Diagnostics: &r5.String{Value: ptr.To(fmt.Sprintf("error executing backend request: %s", err))},
				},
			},
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&resource)
	if err != nil {
		// TODO: return a proper error with operation outcome
		panic(err)
	}

	return resource, nil
}

func (c *Client) Search(ctx context.Context, resourceType string, parameters search.Parameters, options search.Options) (search.Result[model.Resource], error) {
	opts := options

	queryString := search.BuildQuery(parameters, opts)
	url := fmt.Sprintf("%s/%s?%s", c.url, resourceType, queryString)
	log.Printf("forwarding GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return search.Result[model.Resource]{}, r5.OperationOutcome{
			Issue: []r5.OperationOutcomeIssue{
				{
					Severity:    r5.Code{Value: ptr.To("error")},
					Code:        r5.Code{Value: ptr.To("exception")},
					Diagnostics: &r5.String{Value: ptr.To(fmt.Sprintf("error executing backend request: %s", err))},
				},
			},
		}
	}

	// TODO: proper error handling based on the response status code

	var bundle r5.Bundle
	err = json.NewDecoder(resp.Body).Decode(&bundle)
	if err != nil {
		return search.Result[model.Resource]{}, r5.OperationOutcome{
			Issue: []r5.OperationOutcomeIssue{
				{
					Severity:    r5.Code{Value: ptr.To("error")},
					Code:        r5.Code{Value: ptr.To("exception")},
					Diagnostics: &r5.String{Value: ptr.To(fmt.Sprintf("got invalid JSON from backend: %s", err))},
				},
			},
		}
	}

	resources := make([]model.Resource, 0, len(bundle.Entry))
	for _, entry := range bundle.Entry {
		resources = append(resources, entry.Resource)
	}

	return search.Result[model.Resource]{
		Resources: resources,
	}, nil
}
