package rest

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/capabilities/wrap"
	"fhir-toolbox/model"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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
	baseURL, err := url.Parse(config.Base)
	if err != nil {
		return fmt.Errorf("unable to parse base URL: %w", err)
	}

	basePath := strings.Trim(baseURL.Path, "/ ")
	if basePath != "" {
		basePath = "/" + basePath
	}

	tz, err := time.LoadLocation(config.Timezone)
	if err != nil {
		return fmt.Errorf("unable to load timezone: %w", err)
	}

	date, _, err := search.ParseDate(config.Date, tz)
	if err != nil {
		return fmt.Errorf("error parsing date '%s': %w", config.Date, err)
	}

	mux.Handle(fmt.Sprintf("GET %s/metadata", basePath),
		metadataHandler[R](backend, config.DefaultFormat, baseURL, date))
	mux.Handle(fmt.Sprintf("GET %s/{type}/{id}", basePath),
		readHandler(backend, config.DefaultFormat))
	mux.Handle(fmt.Sprintf("GET %s/{type}", basePath),
		searchHandler(backend, config.DefaultFormat, baseURL, tz, config.MaxCount, config.DefaultCount))

	return nil
}

func metadataHandler[R model.Release](
	backend capabilities.GenericAPI,
	defaultFormat Format,
	baseURL *url.URL,
	date time.Time,
) http.Handler {
	capabilityStatement := CapabilityStatement[R](baseURL, backend.AllCapabilities(), date)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := detectFormat(r, defaultFormat)
		returnResult(w, format, http.StatusOK, capabilityStatement)
	})
}

func readHandler(
	backend capabilities.GenericAPI,
	defaultFormat Format,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := detectFormat(r, defaultFormat)
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
	defaultFormat Format,
	baseURL *url.URL,
	tz *time.Location,
	maxCount,
	defaultCount int,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := detectFormat(r, defaultFormat)
		resourceType := r.PathValue("type")

		status, resource := dispatchSearch(
			r.Context(),
			backend,
			resourceType,
			r.URL.Query(),
			baseURL,
			tz,
			maxCount,
			defaultCount,
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
	resourceType string,
	parameters url.Values,
	baseURL *url.URL,
	tz *time.Location,
	maxCount,
	defaultCount int,
) (int, model.Resource) {
	searchCapabilities, err := backend.SearchCapabilities(resourceType)

	options, err := parseSearchOptions(searchCapabilities, parameters, tz, maxCount, defaultCount)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	resources, err := backend.Search(context, resourceType, options)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	bundle, err := SearchBundle(resourceType, resources, options, searchCapabilities, baseURL)
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
