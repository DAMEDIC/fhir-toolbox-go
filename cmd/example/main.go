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

	slog.SetLogLoggerLevel(slog.LevelDebug)
	log.Printf("FHIR Server: %s", backendUrl)

	client := Client{
		url: strings.TrimRight(backendUrl, "/"),
	}
	log.Fatal(http.ListenAndServe(":80", rest.NewServer[model.R4](&client, rest.DefaultConfig)))
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
