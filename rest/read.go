package rest

import (
	"context"
	"fhir-toolbox/backend"
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"net/http"
)

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
