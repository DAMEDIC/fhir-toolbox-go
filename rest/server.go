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
	"cmp"
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/update"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/rest/internal/encoding"
	"github.com/DAMEDIC/fhir-toolbox-go/rest/internal/wrap"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Format = encoding.Format

const (
	FormatJSON = encoding.FormatJSON
	FormatXML  = encoding.FormatXML
)

var (
	defaultTimezone      = time.Local
	defaultMaxCount      = 500
	defaultDefaultCount  = 500
	defaultDefaultFormat = FormatJSON
)

// Server is a generic FHIR server type that registers and serves HTTP endpoints for FHIR resource interactions.
// It is parameterized with the FHIR Release version R as a type constraint.
// It supports configuration of timezone, default and maximum search result counts, and strict parameter handling.
//
// The Server uses a dynamic backend system that detects capabilities through type assertions.
// The Backend field can accept two different types of implementations:
//
// 1. Generic API implementations (resource-agnostic):
//   - capabilities.GenericCapabilities: For capability statement generation
//   - capabilities.GenericCreate: For resource creation operations
//   - capabilities.GenericRead: For resource retrieval operations
//   - capabilities.GenericUpdate: For resource update operations
//   - capabilities.GenericDelete: For resource deletion operations
//   - capabilities.GenericSearch: For resource search operations
//
// 2. Concrete API implementations (resource-specific):
//   - capabilities.ConcreteCapabilities: Required for concrete backends
//   - Resource-specific methods like PatientRead, PatientSearch, etc.
//
// The Server automatically detects which approach your backend uses and handles
// the appropriate conversions. For concrete implementations, it will wrap them
// in a generic interface adapter using wrap.Generic[R].
//
// If a backend doesn't implement a specific capability interface, the corresponding
// HTTP endpoint will return a "not-supported" OperationOutcome.
//
// Backends can implement only the capabilities they need, giving flexibility while
// maintaining type safety. The server handles validation and proper error responses.
//
// NOTE: When using the concrete API, you must implement the CapabilityBase method
// which provides the base CapabilityStatement that will be enhanced with the
// capabilities detected from your concrete implementation.
type Server[R model.Release] struct {
	// Backend is the actual concrete or generic FHIR handler.
	// This field can hold any implementation that satisfies at least one of the capability
	// interfaces from the capabilities package. The server will detect which operations
	// are supported through type assertions at runtime.
	Backend any

	// Timezone used for parsing date parameters without timezones.
	// Defaults to current server timezone.
	Timezone *time.Location
	// MaxCount of search bundle entries.
	// Defaults to 500.
	MaxCount int
	// DefaultCount of search bundle entries.
	// Defaults to 500.
	DefaultCount int
	// DefaultFormat of the server.
	// Defaults to JSON.
	DefaultFormat Format
	// StrictSearchParameters when true causes the server to return an error
	// if unsupported search parameters are used. When false (default),
	// unsupported search parameters are silently ignored.
	StrictSearchParameters bool

	// internal fields
	muxMu sync.Mutex
	mux   *http.ServeMux
}

func (s *Server[R]) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if s.mux == nil {
		s.registerRoutes()
	}
	s.mux.ServeHTTP(writer, request)
}

func (s *Server[R]) registerRoutes() {
	s.muxMu.Lock()
	defer s.muxMu.Unlock()

	// double check, mux might have been set in the background while waiting
	if s.mux != nil {
		return
	}

	s.mux = http.NewServeMux()
	s.mux.Handle("GET /metadata", http.HandlerFunc(s.handleMetadata))
	s.mux.Handle("POST /{type}", http.HandlerFunc(s.handlerCreate))
	s.mux.Handle("GET /{type}/{id}", http.HandlerFunc(s.handlerRead))
	s.mux.Handle("PUT /{type}/{id}", http.HandlerFunc(s.handlerUpdate))
	s.mux.Handle("DELETE /{type}/{id}", http.HandlerFunc(s.handlerDelete))
	s.mux.Handle("GET /{type}", http.HandlerFunc(s.handlerSearch))

}

func (s *Server[R]) genericBackend() (capabilities.GenericCapabilities, error) {
	genericBackend, ok := s.Backend.(capabilities.GenericCapabilities)
	if ok {
		return genericBackend, nil
	}

	concreteBackend, ok := s.Backend.(capabilities.ConcreteCapabilities)
	if !ok {
		return nil, fmt.Errorf("backend does not implement capabilities.GenericCapabilities or capabilities.ConcreteCapabilities")
	}

	return wrap.Generic[R](concreteBackend)
}

func (s *Server[R]) handleMetadata(w http.ResponseWriter, r *http.Request) {
	responseFormat := s.detectFormat(r, "Accept")

	backend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	capabilityStatement, err := backend.CapabilityStatement(r.Context())
	if err != nil {
		slog.Error("error getting metadata", "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	// The CapabilityStatement comes fully configured from the backend
	// No additional modification needed
	returnResult(w, capabilityStatement, http.StatusOK, responseFormat)

}

func (s *Server[R]) handlerCreate(w http.ResponseWriter, r *http.Request) {
	requestFormat := s.detectFormat(r, "Content-Type")
	responseFormat := s.detectFormat(r, "Accept")
	resourceType := r.PathValue("type")

	anyBackend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		returnErr(w, err, responseFormat)
		return
	}
	backend, impl := anyBackend.(capabilities.GenericCreate)

	if !checkInteractionImplemented[R](impl, "create", w, responseFormat) {
		return
	}

	resource, err := dispatchCreate[R](r, backend, requestFormat, resourceType)
	if err != nil {
		slog.Error("error creating resource", "resourceType", resourceType, "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	// fall back to empty string if id is not set
	id, _ := resource.ResourceId()
	baseURL := getBaseURL(r)
	w.Header().Set("Location", baseURL.JoinPath(resourceType, id).String())

	returnResult(w, resource, http.StatusCreated, responseFormat)
}

func dispatchCreate[R model.Release](
	r *http.Request,
	backend capabilities.GenericCreate,
	requestFormat encoding.Format,
	resourceType string,
) (model.Resource, error) {
	resource, err := encoding.DecodeResource[R](r.Body, requestFormat)
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

func (s *Server[R]) handlerRead(w http.ResponseWriter, r *http.Request) {
	responseFormat := s.detectFormat(r, "Accept")
	resourceType := r.PathValue("type")
	resourceID := r.PathValue("id")

	anyBackend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		returnErr(w, err, responseFormat)
		return
	}
	backend, impl := anyBackend.(capabilities.GenericRead)

	if !checkInteractionImplemented[R](impl, "read", w, responseFormat) {
		return
	}

	resource, err := dispatchRead(r.Context(), backend, resourceType, resourceID)
	if err != nil {
		slog.Error("error reading resource", "resourceType", resourceType, "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	returnResult(w, resource, http.StatusOK, responseFormat)
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

func (s *Server[R]) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	requestFormat := s.detectFormat(r, "Content-Type")
	responseFormat := s.detectFormat(r, "Accept")
	resourceType := r.PathValue("type")
	resourceID := r.PathValue("id")

	anyBackend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		returnErr(w, err, responseFormat)
		return
	}
	backend, impl := anyBackend.(capabilities.GenericUpdate)

	if !checkInteractionImplemented[R](impl, "update", w, responseFormat) {
		return
	}

	result, err := dispatchUpdate[R](r, backend, requestFormat, resourceType, resourceID)
	if err != nil {
		slog.Error("error updating resource", "resourceType", resourceType, "id", resourceID, "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	// set Location header with the resource's logical ID
	// the dispatchUpdate function checks that the path id matches the id included in the resource
	baseURL := getBaseURL(r)
	w.Header().Set("Location", baseURL.JoinPath(resourceType, resourceID).String())

	status := http.StatusOK
	if result.Created {
		status = http.StatusCreated
	}

	returnResult(w, result.Resource, status, responseFormat)
}

func dispatchUpdate[R model.Release](
	r *http.Request,
	backend capabilities.GenericUpdate,
	requestFormat encoding.Format,
	resourceType string,
	resourceID string,
) (update.Result[model.Resource], error) {
	resource, err := encoding.DecodeResource[R](r.Body, requestFormat)
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

func (s *Server[R]) handlerDelete(w http.ResponseWriter, r *http.Request) {
	responseFormat := s.detectFormat(r, "Accept")
	resourceType := r.PathValue("type")
	resourceID := r.PathValue("id")

	anyBackend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		returnErr(w, err, responseFormat)
		return
	}
	backend, impl := anyBackend.(capabilities.GenericDelete)

	if !checkInteractionImplemented[R](impl, "delete", w, responseFormat) {
		return
	}

	err = dispatchDelete(r.Context(), backend, resourceType, resourceID)
	if err != nil {
		slog.Error("error deleting resource", "resourceType", resourceType, "id", resourceID, "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func dispatchDelete(
	ctx context.Context,
	backend capabilities.GenericDelete,
	resourceType string,
	resourceID string,
) error {
	return backend.Delete(ctx, resourceType, resourceID)
}

func (s *Server[R]) handlerSearch(w http.ResponseWriter, r *http.Request) {
	responseFormat := s.detectFormat(r, "Accept")
	resourceType := r.PathValue("type")

	anyBackend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		returnErr(w, err, responseFormat)
		return
	}
	backend, impl := anyBackend.(capabilities.GenericSearch)

	if !checkInteractionImplemented[R](impl, "search-type", w, responseFormat) {
		return
	}

	resource, err := dispatchSearch(
		r,
		backend,
		resourceType,
		r.URL.Query(),
		cmp.Or(s.Timezone, defaultTimezone),
		cmp.Or(s.MaxCount, defaultMaxCount),
		cmp.Or(s.DefaultCount, defaultDefaultCount),
		s.StrictSearchParameters,
	)
	if err != nil {
		slog.Error("error reading searching", "resourceType", resourceType, "err", err)
		returnErr(w, err, responseFormat)
		return
	}

	returnResult(w, resource, http.StatusOK, responseFormat)
}

func dispatchSearch(
	r *http.Request,
	backend capabilities.GenericSearch,
	resourceType string,
	parameters url.Values,
	tz *time.Location,
	maxCount, defaultCount int,
	strictSearchParameters bool,
) (model.Resource, error) {
	ctx := r.Context()
	// Get CapabilityStatement to extract SearchParameter information
	capabilityStatement, err := backend.CapabilityStatement(ctx)
	if err != nil {
		return nil, err
	}

	// Create a SearchParameter resolver function
	resolveSearchParameter := func(canonical string) (model.Element, error) {
		// Try to resolve SearchParameter using Read operation if available
		if readBackend, ok := backend.(capabilities.GenericRead); ok {
			// Extract SearchParameter ID from canonical URL
			lastSlash := strings.LastIndex(canonical, "/")
			if lastSlash != -1 && lastSlash < len(canonical)-1 {
				searchParamId := canonical[lastSlash+1:]
				resource, err := readBackend.Read(ctx, "SearchParameter", searchParamId)
				if err != nil {
					return nil, err
				}
				// Return the SearchParameter resource as a model.Element
				return resource, nil
			}
		}
		// Return error if SearchParameter cannot be resolved
		return nil, fmt.Errorf("cannot resolve SearchParameter from canonical URL: %s", canonical)
	}

	searchParameters, options, err := parseSearchOptions(
		capabilityStatement,
		resourceType,
		resolveSearchParameter,
		parameters,
		tz, maxCount, defaultCount,
		strictSearchParameters,
	)
	if err != nil {
		return nil, err
	}

	resources, err := backend.Search(ctx, resourceType, searchParameters, options)
	if err != nil {
		return nil, err
	}

	bundle, err := buildSearchBundle(resourceType, resources, searchParameters, options, capabilityStatement, resolveSearchParameter)
	if err != nil {
		return nil, err
	}

	return bundle, nil
}

func parseSearchOptions(
	capabilityStatement basic.CapabilityStatement,
	resourceType string,
	resolveSearchParameter func(canonical string) (model.Element, error),
	params url.Values,
	tz *time.Location,
	maxCount, defaultCount int,
	strict bool) (search.Parameters, search.Options, error) {
	parameters, options, err := search.ParseQuery(capabilityStatement, resourceType, resolveSearchParameter, params, tz, maxCount, defaultCount, strict)
	if err != nil {
		return nil, search.Options{}, searchError(err.Error())
	}
	return parameters, options, nil
}

func (s *Server[R]) detectFormat(r *http.Request, headerName string) encoding.Format {
	// url parameter overrides the Accept header
	formatQuery := r.URL.Query()["_format"]
	if len(formatQuery) > 0 {
		format := encoding.MatchFormat(formatQuery[0])
		if format != "" {
			return format
		}
	}

	for _, accept := range r.Header[headerName] {
		format := encoding.MatchFormat(accept)
		if format != "" {
			return format
		}
	}
	return cmp.Or(s.DefaultFormat, defaultDefaultFormat)
}

func returnErr(w http.ResponseWriter, err error, format encoding.Format) {
	status, oo := errToOperationOutcome(err)
	returnResult(w, oo, status, format)
}

func returnResult[T any](w http.ResponseWriter, r T, status int, format encoding.Format) {
	w.Header().Set("Content-Type", string(format))
	w.WriteHeader(status)

	err := encoding.Encode(w, r, format)
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
	format encoding.Format,
) bool {
	if !implemented {
		slog.Error("interaction not implemented by backend", "interaction", interaction)
		returnErr(w, notImplementedError(interaction), format)
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

// getBaseURL extracts the base URL from the request
func getBaseURL(r *http.Request) *url.URL {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	// Check for X-Forwarded-Proto header (common in load balancer setups)
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	}

	host := r.Host
	if host == "" {
		host = "localhost"
	}

	return &url.URL{
		Scheme: scheme,
		Host:   host,
	}
}
