package rest

import (
	"fhir-toolbox/capabilities"
	"fhir-toolbox/generic"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewServer[R model.Release](backend any, config Config) (http.Handler, error) {
	mux := http.NewServeMux()

	err := registerRoutes(mux, generic.Wrap[R](backend), config)
	if err != nil {
		return nil, err
	}

	var handler http.Handler = mux
	handler = withLogging(handler)
	handler = withRequestContext(handler)

	return handler, nil
}

func registerRoutes(
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

	mux.Handle(fmt.Sprintf("GET %s/{type}/{id}", basePath), readHandler(backend))
	mux.Handle(fmt.Sprintf("GET %s/{type}", basePath), searchHandler(backend, baseURL, tz, config.MaxCount, config.DefaultCount))

	return nil
}

func readHandler(
	backend capabilities.GenericAPI,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		status, resource := dispatchRead(r.Context(), backend, resourceType, resourceID)
		if outcome, ok := resource.(basic.OperationOutcome); ok {
			slog.Error("error reading resource", "resourceType", resourceType, "OperationOutcome", outcome)
		}

		err := encodeJSON(w, status, resource)
		if err != nil {
			// we were not able to return an application level error (OperationOutcome)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func searchHandler(
	backend capabilities.GenericAPI,
	baseURL *url.URL,
	tz *time.Location,
	maxCount,
	defaultCount int,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if outcome, ok := resource.(basic.OperationOutcome); ok {
			slog.Error("error searching resource", "resourceType", resourceType, "OperationOutcome", outcome)
		}

		err := encodeJSON(w, status, resource)
		if err != nil {
			// we were not able to return an application level error (OperationOutcome)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
