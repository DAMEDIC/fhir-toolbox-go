package rest

import (
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"log/slog"
	"net/http"
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
	base := strings.Trim(config.Base, "/ ")
	if base != "" {
		base = "/" + base
	}

	tz, err := time.LoadLocation(config.Timezone)
	if err != nil {
		return fmt.Errorf("unable to load timezone: %w", err)
	}

	mux.Handle(fmt.Sprintf("GET %s/{type}/{id}", base), handleRead(dispatch, backend))
	mux.Handle(fmt.Sprintf("GET %s/{type}", base), handleSearchType(dispatch, backend, base, tz))

	return nil
}

func handleRead(
	dispatch dispatch.Dispatcher,
	backend Backend,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		status, resource := read(r.Context(), dispatch, backend, resourceType, resourceID)
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

func handleSearchType(
	dispatch dispatch.Dispatcher,
	backend Backend,
	base string,
	tz *time.Location,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")

		status, resource := searchType(r.Context(), dispatch, backend, resourceType, r.URL.Query(), baseURL(r.URL.Scheme, r.Host, base), tz)
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

func baseURL(scheme, host, base string) string {
	if scheme == "" {
		scheme = "http"
	}
	return fmt.Sprintf("%s://%s%s", scheme, host, base)
}
