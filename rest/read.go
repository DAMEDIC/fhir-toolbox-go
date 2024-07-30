package rest

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"net/http"
)

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
