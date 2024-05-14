package rest

import (
	"context"
	"fhir-toolbox/backend"
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"fmt"
	"net/http"
)

func NewServer[R model.Release](backend backend.Backend, config Config) http.Handler {
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
	backend backend.Backend,
	config Config,
) {
	mux.Handle(fmt.Sprintf("GET %s/{type}/{id}", config.Base), handleRead(dispatch, backend))
}

func handleRead(
	dispatch dispatch.Dispatcher,
	backend backend.Backend,
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

func read(
	context context.Context,
	dispatch dispatch.Dispatcher,
	backend backend.Backend,
	resourceType string,
	resourceID string,
) (int, model.Resource) {
	resource, err := dispatch.Read(context, backend, resourceType, resourceID)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}
	return http.StatusOK, resource
}
