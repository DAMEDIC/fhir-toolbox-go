package rest

import (
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

		resource, err := dispatch.Read(r.Context(), backend, resourceType, resourceID)
		if err != nil {
			// TODO: return appropiate OperationOutcome
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = encodeJSON(w, http.StatusOK, resource)
		if err != nil {
			// TODO: return appropiate OperationOutcome
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
