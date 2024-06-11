package rest

import (
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"fmt"
	"net/http"
	"strings"
)

type Backend any

func NewServer[R model.Release](backend Backend, config Config) http.Handler {
	dispatcher := dispatch.DispatcherFor[R]()
	mux := http.NewServeMux()

	registerRoutes(mux, dispatcher, backend, config)

	var handler http.Handler = mux
	handler = withLogging(handler)
	handler = withRequestContext(handler)

	return handler
}

func registerRoutes(
	mux *http.ServeMux,
	dispatch dispatch.Dispatcher,
	backend Backend,
	config Config,
) {
	base := strings.Trim(config.Base, "/ ")
	if base != "" {
		base = "/" + base
	}

	mux.Handle(fmt.Sprintf("GET %s/{type}/{id}", base), handleRead(dispatch, backend))
	mux.Handle(fmt.Sprintf("GET %s/{type}", base), handleSearchType(dispatch, backend, base))
}

func handleRead(
	dispatch dispatch.Dispatcher,
	backend Backend,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")
		resourceID := r.PathValue("id")

		status, resource := read(r.Context(), dispatch, backend, resourceType, resourceID)

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
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.PathValue("type")

		status, resource := searchType(r.Context(), dispatch, backend, resourceType, r.URL.Query(), baseURL(r.URL.Scheme, r.Host, base))

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
