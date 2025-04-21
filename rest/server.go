// Package rest provides a FHIR REST API.
//
// Capabilities are detected by type assertion, a generated CapabilityStatement is served at the "/metadata" endpoint.
//
// Following interactions are currently supported:
//   - read
//   - search (parameters are passed down to the supplied backend implementation)
//
// # Base URL and routes
//
// You have to pass a base URL using the config.
// This base URL is only used for building response Bundles.
// For supported interactions, the returned http.Handler has sub-routes installed.
// These are always installed at the root of this handler.
// The base URL from the config is not used.
//
// Currently, installed patterns are:
//   - capabilities: "GET /metadata"
//   - read: "GET /{type}/{id}"
//   - search: "GET /{type}/"
//
// If you do not want your FHIR handlers installed at the root, use something like
//
//	mux.Handle("/path/", http.StripPrefix("/path", serverHandler)
//
// This allows you to implement multiple FHIR REST APIs on the same HTTP server
// (e.g. for multi-tenancy scenarios).
//
// # Pagination
//
// Cursor-based pagination can be implemented by the backend.
// Therefor the parameters "_count" and "_cursor" are passed down to the backend.
package rest

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/rest/internal/wrap"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

// NewServer returns a http.Handler that serves the supplied backend.
func NewServer[R model.Release](backend any, config Config) (http.Handler, error) {
	mux := http.NewServeMux()

	genericBackend, err := wrap.Generic[R](backend)
	if err != nil {
		return nil, err
	}

	err = registerRoutes[R](mux, genericBackend, config)
	if err != nil {
		return nil, err
	}

	var handler http.Handler = mux
	handler = withLogging(handler)
	handler = withRequestContext(handler)

	return handler, nil
}

func registerRoutes[R model.Release](
	mux *http.ServeMux,
	backend capabilities.GenericCapabilities,
	config Config,
) error {
	date, _, err := search.ParseDate(config.Date, config.Timezone)
	if err != nil {
		return fmt.Errorf("error parsing date '%s': %w", config.Date, err)
	}

	mux.Handle("GET /metadata", metadataHandler[R](backend, config, date))
	mux.Handle("POST /{type}", createHandler[R](backend, config))
	mux.Handle("GET /{type}/{id}", readHandler(backend, config))
	mux.Handle("GET /{type}", searchHandler(backend, config, config.Timezone))

	return nil
}

func metadataHandler[R model.Release](
	backend capabilities.GenericCapabilities,
	config Config,
	date time.Time,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormat := detectFormat(r, "Accept", config.DefaultFormat)

		capabilities, err := backend.AllCapabilities(r.Context())
		if err != nil {
			returnResult(w, responseFormat, err.StatusCode(), err.OperationOutcome())
			return
		}

		capabilityStatement := capabilityStatement[R](config.Base, capabilities, date)
		returnResult(w, responseFormat, http.StatusOK, capabilityStatement)
	})
}

func createHandler[R model.Release](
	anyBackend any,
	config Config,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericCreate)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestFormat := detectFormat(r, "Content-Type", config.DefaultFormat)
		responseFormat := detectFormat(r, "Accept", requestFormat)
		resourceType := r.PathValue("type")

		if !implemented {
			err := capabilities.NotImplementedError{
				Interaction: "create",
			}
			returnResult(w, responseFormat, err.StatusCode(), err.OperationOutcome())
			return
		}

		status, resource := dispatchCreate[R](r, backend, requestFormat, resourceType)
		if status != http.StatusCreated {
			slog.Error("error creating resource", "resourceType", resourceType, "outcome", resource)
		} else {
			// fall back to empty string if id is not set
			id, _ := resource.ResourceId()
			w.Header().Set("Location", config.Base.JoinPath(resourceType, id).String())
		}

		returnResult(w, responseFormat, status, resource)
	})
}

type wrongRequestBodyError struct {
	ResourceType string
}

// An UnexpectedResourceError is returned when an unexpected resource type is returned.
type unexpectedResourceClientError struct {
	ExpectedType string
	GotType      string
}

func (e unexpectedResourceClientError) Error() string {
	return fmt.Sprintf("unexpected resource from client: expected %s, got %s", e.ExpectedType, e.GotType)
}

func (e unexpectedResourceClientError) StatusCode() int {
	return 400
}

func (e unexpectedResourceClientError) OperationOutcome() basic.OperationOutcome {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{
				Severity:    basic.Code{Value: utils.Ptr("fatal")},
				Code:        basic.Code{Value: utils.Ptr("processing")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
}
func dispatchCreate[R model.Release](
	r *http.Request,
	backend capabilities.GenericCreate,
	requestFormat format,
	resourceType string,
) (int, model.Resource) {
	var resource model.Resource
	var err capabilities.FHIRError

	switch requestFormat {
	case formatJSON:
		resource, err = decodeResourceJSON[R](r)
	case formatXML:
		resource, err = decodeResourceXML[R](r)
	}

	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	if resourceType != resource.ResourceType() {
		err = unexpectedResourceClientError{
			ExpectedType: resourceType,
			GotType:      resource.ResourceType(),
		}
		return err.StatusCode(), err.OperationOutcome()
	}

	createdResource, err := backend.Create(r.Context(), resource)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	return http.StatusCreated, createdResource
}

func readHandler(
	anyBackend any,
	config Config,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericRead)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormat := detectFormat(r, "Accept", config.DefaultFormat)
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		if !implemented {
			err := capabilities.NotImplementedError{
				Interaction: "read",
			}
			returnResult(w, responseFormat, err.StatusCode(), err.OperationOutcome())
			return
		}

		status, resource := dispatchRead(r.Context(), backend, resourceType, resourceID)
		if status != http.StatusOK {
			slog.Error("error reading resource", "resourceType", resourceType, "outcome", resource)
		}

		returnResult(w, responseFormat, status, resource)
	})
}

func dispatchRead(
	ctx context.Context,
	backend capabilities.GenericRead,
	resourceType string,
	resourceID string,
) (int, model.Resource) {
	resource, err := backend.Read(ctx, resourceType, resourceID)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}
	return http.StatusOK, resource
}

func searchHandler(
	anyBackend any,
	config Config,
	tz *time.Location,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericSearch)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormat := detectFormat(r, "Accept", config.DefaultFormat)
		resourceType := r.PathValue("type")

		if !implemented {
			err := capabilities.NotImplementedError{
				Interaction: "search-type",
			}
			returnResult(w, responseFormat, err.StatusCode(), err.OperationOutcome())
			return
		}

		status, resource := dispatchSearch(
			r.Context(),
			backend,
			config,
			resourceType,
			r.URL.Query(),
			tz,
		)
		if status != http.StatusOK {
			slog.Error("error reading resource", "resourceType", resourceType, "outcome", resource)
		}

		returnResult(w, responseFormat, status, resource)
	})
}

func dispatchSearch(
	ctx context.Context,
	backend capabilities.GenericSearch,
	config Config,
	resourceType string,
	parameters url.Values,
	tz *time.Location,
) (int, model.Resource) {
	allCapabilities, err := backend.AllCapabilities(ctx)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	searchCapabilities := allCapabilities.SearchCapabilities[resourceType]

	options, err := parseSearchOptions(searchCapabilities, parameters, tz, config.MaxCount, config.DefaultCount)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	resources, err := backend.Search(ctx, resourceType, options)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	bundle, err := searchBundle(resourceType, resources, options, searchCapabilities, config.Base)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	return http.StatusOK, bundle
}

func parseSearchOptions(
	searchCapabilities search.Capabilities,
	params url.Values,
	tz *time.Location,
	maxCount, defaultCount int) (search.Options, capabilities.FHIRError) {
	options, err := search.ParseOptions(searchCapabilities, params, tz, maxCount, defaultCount)
	if err != nil {
		return search.Options{}, capabilities.NewSearchError(err)
	}
	return options, nil
}

func returnResult[T any](w http.ResponseWriter, format format, status int, r T) {
	var err error
	switch format {
	case formatJSON:
		err = encodeJSON(w, status, r)
	case formatXML:
		err = encodeXML(w, status, r)
	}
	if err != nil {
		// we were not able to return an application level error (OperationOutcome)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
