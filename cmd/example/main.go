// This is a simple example of a FHIR server that reads resources from another FHIR server.
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

	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/rest"
)

func main() {
	var backendUrl = os.Args[1]

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	// Add some addtionnal context to the log (like request id).
	requestContextHandler := rest.NewRequestContextSlogHandler(textHandler)
	slog.SetDefault(slog.New(requestContextHandler))

	log.Printf("FHIR Server: %s", backendUrl)

	// Create the client to the backing server.
	// The client is implemented below and is only an example.
	// It currenlty only supports the read operation.
	client := Client{
		url: strings.TrimRight(backendUrl, "/"),
	}

	// Create the REST server.
	// You can plug in any backend you want here.
	server := rest.NewServer[model.R4](&client, rest.DefaultConfig)

	// Start the server and listen on port 80.
	log.Fatal(http.ListenAndServe(":80", server))
}

type Client struct {
	url string
}

func Read[R model.Resource](ctx context.Context, client *Client, id string) (R, capabilities.FHIRError) {
	var resource R
	resourceType := resource.ResourceType()

	url := fmt.Sprintf("%s/%s/%s", client.url, resourceType, id)
	log.Printf("forwarding GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		// TODO: retrun a proper error with operation outcome
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&resource)
	if err != nil {
		// TODO: retrun a proper error with operation outcome
		panic(err)
	}

	return resource, nil
}

// In a real implementation, these methods would be auto-generated for all resource types.
func (c *Client) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	return Read[r4.Patient](ctx, c, id)
}

func (c *Client) ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError) {
	return Read[r4.Observation](ctx, c, id)
}
