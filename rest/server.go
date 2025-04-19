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
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/wrap"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
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
	backend capabilities.GenericAPI,
	config Config,
) error {
	date, _, err := search.ParseDate(config.Date, config.Timezone)
	if err != nil {
		return fmt.Errorf("error parsing date '%s': %w", config.Date, err)
	}

	mux.Handle("GET /metadata", metadataHandler[R](backend, config, date))
	mux.Handle("GET /{type}/{id}", readHandler(backend, config))
	mux.Handle("GET /{type}", searchHandler(backend, config, config.Timezone))

	return nil
}

func metadataHandler[R model.Release](
	backend capabilities.GenericAPI,
	config Config,
	date time.Time,
) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := detectFormat(r, config.DefaultFormat)

		capabilities, err := backend.AllCapabilities()
		if err != nil {
			returnResult(w, format, err.StatusCode(), err.OperationOutcome())
			return
		}

		capabilityStatement := capabilityStatement[R](config.Base, capabilities, date)
		returnResult(w, format, http.StatusOK, capabilityStatement)
	})
}

func readHandler(
	backend capabilities.GenericAPI,
	config Config,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := detectFormat(r, config.DefaultFormat)
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		status, resource := dispatchRead(r.Context(), backend, resourceType, resourceID)
		if status != http.StatusOK {
			slog.Error("error reading resource", "resourceType", resourceType, "outcome", resource)
		}

		returnResult(w, format, status, resource)
	})
}

func dispatchRead(
	context context.Context,
	backend capabilities.GenericAPI,
	resourceType string,
	resourceID string,
) (int, model.Resource) {
	resource, err := backend.Read(context, resourceType, resourceID)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}
	return http.StatusOK, resource
}

func searchHandler(
	backend capabilities.GenericAPI,
	config Config,
	tz *time.Location,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := detectFormat(r, config.DefaultFormat)
		resourceType := r.PathValue("type")

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

		returnResult(w, format, status, resource)
	})
}

func dispatchSearch(
	context context.Context,
	backend capabilities.GenericAPI,
	config Config,
	resourceType string,
	parameters url.Values,
	tz *time.Location,
) (int, model.Resource) {
	searchCapabilities, err := backend.SearchCapabilities(resourceType)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	options, err := parseSearchOptions(searchCapabilities, parameters, tz, config.MaxCount, config.DefaultCount)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	resources, err := backend.Search(context, resourceType, options)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	bundle, err := SearchBundle(resourceType, resources, options, searchCapabilities, config.Base)
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

func returnResult[T any](w http.ResponseWriter, format Format, status int, r T) {
	var err error
	switch format {
	case FormatJSON:
		err = encodeJSON(w, status, r)
	case FormatXML:
		err = encodeXML(w, status, r)
	}
	if err != nil {
		// we were not able to return an application level error (OperationOutcome)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
