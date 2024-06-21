package rest

import (
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Backend any

func NewServer[R model.Release](backend Backend, config Config) (http.Handler, error) {
	dispatcher := dispatch.DispatcherFor[R]()
	mux := http.NewServeMux()

	err := registerRoutes(mux, dispatcher, backend, config)
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
	dispatch dispatch.Dispatcher,
	backend Backend,
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

	mux.Handle(fmt.Sprintf("GET %s/{type}/{id}", basePath), readHandler(dispatch, backend))
	mux.Handle(fmt.Sprintf("GET %s/{type}", basePath), searchHandler(dispatch, backend, baseURL, tz, config.MaxCount, config.DefaultCount))

	return nil
}

func readHandler(
	dispatch dispatch.Dispatcher,
	backend Backend,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		status, resource := dispatchRead(r.Context(), dispatch, backend, resourceType, resourceID)
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
	dispatch dispatch.Dispatcher,
	backend Backend,
	baseURL *url.URL,
	tz *time.Location,
	maxCount,
	defaultCount int,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")

		status, resource := dispatchSearch(
			r.Context(),
			dispatch,
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
