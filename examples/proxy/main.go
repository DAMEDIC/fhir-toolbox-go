// This is a simple examples of a FHIR server that reads resources from another FHIR server.
package main

import (
	"context"
	"encoding/json"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model/basic"
	"fhir-toolbox/model/raw"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fhir-toolbox/rest"
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

	// Create the client to the backing server.
	// The client is implemented below and is only an examples.
	// It currently only supports the read and search operations.
	client := Client{
		url: strings.TrimRight(backendUrl, "/"),
	}

	// Create the REST server.
	// You can plug in any backend you want here.
	server, err := rest.NewServer[model.R4](&client, rest.DefaultConfig)
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	// Start the server and listen on port 80.
	log.Fatal(http.ListenAndServe(":80", server))
}

type Client struct {
	url string
}

func (c *Client) Read(ctx context.Context, resourceType string, id string) (model.Resource, capabilities.FHIRError) {
	var resource raw.Resource

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
		Parameters: map[string]search.Desc{
			"_id": {Type: search.String},
		},
	}, nil
}

func (c *Client) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, capabilities.FHIRError) {
	params := url.Values{}
	for key, ands := range options.Parameters {
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

	var bundle basic.Bundle
	err = json.NewDecoder(resp.Body).Decode(&bundle)
	if err != nil {
		// TODO: return a proper error with operation outcome
		panic(err)
	}

	resources := make([]model.Resource, 0, len(bundle.Entry))
	for _, entry := range bundle.Entry {
		resources = append(resources, entry.Resource)
	}

	return search.Result{
		Resources: resources,
	}, nil
}
