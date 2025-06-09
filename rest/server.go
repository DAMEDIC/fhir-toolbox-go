// Package rest provides a FHIR REST API.
//
// Capabilities are detected by type assertion, a generated CapabilityStatement is served at the "/metadata" endpoint.
//
// Following interactions are currently supported:
//   - read
//   - create
//   - update
//   - delete
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
//   - create: "POST /{type}"
//   - read: "GET /{type}/{id}"
//   - update: "PUT /{type}/{id}"
//   - delete: "DELETE /{type}/{id}"
//   - search: "GET /{type}"
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
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/update"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/rest/internal/wrap"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
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
	mux.Handle("GET /metadata", metadataHandler[R](backend, config, config.Date))
	mux.Handle("POST /{type}", createHandler[R](backend, config))
	mux.Handle("GET /{type}/{id}", readHandler[R](backend, config))
	mux.Handle("PUT /{type}/{id}", updateHandler[R](backend, config))
	mux.Handle("DELETE /{type}/{id}", deleteHandler[R](backend, config))
	mux.Handle("GET /{type}", searchHandler[R](backend, config, config.Timezone))

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
			returnErr[R](w, responseFormat, err)
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

		if !checkInteractionImplemented[R](implemented, "create", w, responseFormat) {
			return
		}

		resource, err := dispatchCreate[R](r, backend, requestFormat, resourceType)
		if err != nil {
			slog.Error("error creating resource", "resourceType", resourceType, "outcome", resource)
			returnErr[R](w, responseFormat, err)
			return
		}

		// fall back to empty string if id is not set
		id, _ := resource.ResourceId()
		w.Header().Set("Location", config.Base.JoinPath(resourceType, id).String())

		returnResult(w, responseFormat, http.StatusCreated, resource)
	})
}

func dispatchCreate[R model.Release](
	r *http.Request,
	backend capabilities.GenericCreate,
	requestFormat Format,
	resourceType string,
) (model.Resource, error) {
	resource, err := decodeResource[R](r, requestFormat)
	if err != nil {
		return nil, err
	}

	if resourceType != resource.ResourceType() {
		return nil, unexpectedResourceError(resourceType, resource.ResourceType())
	}

	createdResource, err := backend.Create(r.Context(), resource)
	if err != nil {
		return nil, err
	}

	return createdResource, nil
}

func readHandler[R model.Release](
	anyBackend any,
	config Config,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericRead)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormat := detectFormat(r, "Accept", config.DefaultFormat)
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		if !checkInteractionImplemented[R](implemented, "read", w, responseFormat) {
			return
		}

		resource, err := dispatchRead(r.Context(), backend, resourceType, resourceID)
		if err != nil {
			slog.Error("error reading resource", "resourceType", resourceType, "outcome", resource)
			returnErr[R](w, responseFormat, err)
			return
		}

		returnResult(w, responseFormat, http.StatusOK, resource)
	})
}

func dispatchRead(
	ctx context.Context,
	backend capabilities.GenericRead,
	resourceType string,
	resourceID string,
) (model.Resource, error) {
	resource, err := backend.Read(ctx, resourceType, resourceID)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func updateHandler[R model.Release](
	anyBackend any,
	config Config,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericUpdate)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestFormat := detectFormat(r, "Content-Type", config.DefaultFormat)
		responseFormat := detectFormat(r, "Accept", requestFormat)
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		if !checkInteractionImplemented[R](implemented, "update", w, responseFormat) {
			return
		}

		result, err := dispatchUpdate[R](r, backend, requestFormat, resourceType, resourceID)
		if err != nil {
			slog.Error("error updating resource", "resourceType", resourceType, "id", resourceID, "outcome", result)
			returnErr[R](w, responseFormat, err)
			return
		}

		// set Location header with the resource's logical ID
		// the dispatchUpdate function checks that the path id matches the id included in the resource
		w.Header().Set("Location", config.Base.JoinPath(resourceType, resourceID).String())

		status := http.StatusOK
		if result.Created {
			status = http.StatusCreated
		}

		returnResult(w, responseFormat, status, result.Resource)
	})
}

func dispatchUpdate[R model.Release](
	r *http.Request,
	backend capabilities.GenericUpdate,
	requestFormat Format,
	resourceType string,
	resourceID string,
) (update.Result[model.Resource], error) {
	resource, err := decodeResource[R](r, requestFormat)
	if err != nil {
		return update.Result[model.Resource]{}, err
	}

	if resourceType != resource.ResourceType() {
		return update.Result[model.Resource]{}, unexpectedResourceError(resourceType, resource.ResourceType())
	}

	// Verify that the resource ID in the URL matches the resource ID in the body
	id, ok := resource.ResourceId()
	if !ok || id != resourceID {
		return update.Result[model.Resource]{}, createOperationOutcome(
			"fatal",
			"processing",
			fmt.Sprintf("resource ID in URL (%s) does not match resource ID in body (%s)", resourceID, id),
		)
	}

	result, err := backend.Update(r.Context(), resource)
	if err != nil {
		return update.Result[model.Resource]{}, err
	}

	return result, nil
}

func deleteHandler[R model.Release](
	anyBackend any,
	config Config,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericDelete)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormat := detectFormat(r, "Accept", config.DefaultFormat)
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		if !checkInteractionImplemented[R](implemented, "delete", w, responseFormat) {
			return
		}

		err := dispatchDelete(r.Context(), backend, resourceType, resourceID)
		if err != nil {
			slog.Error("error deleting resource", "resourceType", resourceType, "id", resourceID)
			returnErr[R](w, responseFormat, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func dispatchDelete(
	ctx context.Context,
	backend capabilities.GenericDelete,
	resourceType string,
	resourceID string,
) error {
	return backend.Delete(ctx, resourceType, resourceID)
}

func searchHandler[R model.Release](
	anyBackend any,
	config Config,
	tz *time.Location,
) http.Handler {
	backend, implemented := anyBackend.(capabilities.GenericSearch)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormat := detectFormat(r, "Accept", config.DefaultFormat)
		resourceType := r.PathValue("type")

		if !checkInteractionImplemented[R](implemented, "search-type", w, responseFormat) {
			return
		}

		resource, err := dispatchSearch(
			r.Context(),
			backend,
			config,
			resourceType,
			r.URL.Query(),
			tz,
		)
		if err != nil {
			slog.Error("error reading resource", "resourceType", resourceType, "outcome", resource)
			returnErr[R](w, responseFormat, err)
			return
		}

		returnResult(w, responseFormat, http.StatusOK, resource)
	})
}

func dispatchSearch(
	ctx context.Context,
	backend capabilities.GenericSearch,
	config Config,
	resourceType string,
	parameters url.Values,
	tz *time.Location,
) (model.Resource, error) {
	allCapabilities, err := backend.AllCapabilities(ctx)
	if err != nil {
		return nil, err
	}

	searchCapabilities := allCapabilities.Search[resourceType]

	options, err := parseSearchOptions(searchCapabilities, parameters, tz, config.MaxCount, config.DefaultCount)
	if err != nil {
		return nil, err
	}

	resources, err := backend.Search(ctx, resourceType, options)
	if err != nil {
		return nil, err
	}

	bundle, err := searchBundle(resourceType, resources, options, searchCapabilities, config.Base)
	if err != nil {
		return nil, err
	}

	return bundle, nil
}

func parseSearchOptions(
	searchCapabilities search.Capabilities,
	params url.Values,
	tz *time.Location,
	maxCount, defaultCount int) (search.Options, error) {
	options, err := search.ParseOptions(searchCapabilities, params, tz, maxCount, defaultCount)
	if err != nil {
		return search.Options{}, searchError(err.Error())
	}
	return options, nil
}

func returnErr[R model.Release](w http.ResponseWriter, format Format, err error) {
	status, oo := errToOperationOutcome[R](err)
	returnResult(w, format, status, oo)
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

func createOperationOutcome(severity, code, diagnostics string) basic.OperationOutcome {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{
				Severity:    basic.Code{Value: ptr.To(severity)},
				Code:        basic.Code{Value: ptr.To(code)},
				Diagnostics: &basic.String{Value: ptr.To(diagnostics)},
			},
		},
	}
}

func checkInteractionImplemented[R model.Release](
	implemented bool,
	interaction string,
	w http.ResponseWriter,
	format Format,
) bool {
	if !implemented {
		slog.Error("interaction not implemented by backend", "interaction", interaction)
		returnErr[R](w, format, notImplementedError(interaction))
		return false
	}
	return true
}

func notImplementedError(interaction string) error {
	return createOperationOutcome(
		"fatal",
		"not-supported",
		fmt.Sprintf("%s interaction not implemented", interaction),
	)
}

func unexpectedResourceError(expectedType string, gotType string) basic.OperationOutcome {
	return createOperationOutcome(
		"fatal",
		"processing",
		fmt.Sprintf("unexpected resource: expected %s, got %s", expectedType, gotType),
	)
}

func searchError(msg string) basic.OperationOutcome {
	return createOperationOutcome("fatal", "processing", msg)
}
