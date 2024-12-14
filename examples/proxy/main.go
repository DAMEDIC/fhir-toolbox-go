// A FHIR server that serves resources from another FHIR server.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
)

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
	// A full-featured client should be in scope of this package and might be added eventually.
	// The example client only supports the read and search operations and exposes them using the generic API.
	genericClient := Client{
		url: strings.TrimRight(backendUrl, "/"),
	}

	// You can provide the concrete API by wrapping the generic API
	// (uncomment the following lines to try it out):
	//concreteApi := wrap.Concreter5(&genericClient)
	//somePatient, fhirErr := concreteApi.ReadPatient(context.Background(), "547")
	//if fhirErr != nil {
	//	log.Fatalf("error reading some Patient %v", fhirErr)
	//}
	//fmt.Printf("some Patient: %v\n", somePatient)

	// Create the REST server.
	// You can plug in any backend you want here.
	// Note: it is important to pass a references, as the methods below also implemented on a pointer as receiver.
	server, err := rest.NewServer[model.R5](&genericClient, rest.DefaultConfig)
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

func AllCapabilities() capabilities.Capabilities {
	// TODO: this should call the remote service for its capabilities
	return capabilities.Capabilities{}
}

func (c *Client) Read(ctx context.Context, resourceType string, id string) (model.Resource, capabilities.FHIRError) {
	// ContainedResource is a concrete representation of any resource
	// internally this e.g. uses in bundles
	var resource r5.ContainedResource

	url := fmt.Sprintf("%s/%s/%s", c.url, resourceType, id)
	log.Printf("forwarding GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		// TODO: return a proper error with operation outcome
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&resource)
	if err != nil {
		// TODO: return a proper error with operation outcome
		panic(err)
	}

	return resource, nil
}

func (c *Client) SearchCapabilities(resourceType string) (search.Capabilities, capabilities.FHIRError) {
	// TODO: These should be read from the remote servers CapabilityStatement.
	return search.Capabilities{
		Params: map[string]search.ParamDesc{
			"_id": {Type: search.String},
		},
	}, nil
}

func (c *Client) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, capabilities.FHIRError) {
	params := url.Values{}
	for key, ands := range options.Params {
		for _, and := range ands {
			orStrings := make([]string, 0, len(and))
			for _, or := range and {
				orStrings = append(orStrings, or.String())
			}
			params.Add(key, strings.Join(orStrings, ","))
		}
	}

	url := fmt.Sprintf("%s/%s?%s", c.url, resourceType, params.Encode())
	log.Printf("forwarding GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		// TODO: return a proper error with operation outcome
		panic(err)
	}

	var bundle r5.Bundle
	err = json.NewDecoder(resp.Body).Decode(&bundle)
	if err != nil {
		// TODO: return a proper error with operation outcome
		panic(err)
	}

	resources := make([]model.Resource, 0, len(bundle.Entry))
	for _, entry := range bundle.Entry {
		resources = append(resources, *entry.Resource)
	}

	return search.Result{
		Resources: resources,
	}, nil
}
