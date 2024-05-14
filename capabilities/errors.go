package capabilities

import (
	"encoding/json"
	"fhir-toolbox/model"
	"fmt"
)

type FHIRError interface {
	error
	StatusCode() int
	OperationOutcome() model.Resource
}

type BasicOperationOutcome struct {
	Issue []BasicOperationOutcomeIssue `json:"issue"`
}

type BasicOperationOutcomeIssue struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
}

func (r BasicOperationOutcome) ResourceType() string {
	return "OperationOutcome"
}

func (r BasicOperationOutcome) MarshalJSON() ([]byte, error) {
	type embedded BasicOperationOutcome
	return json.Marshal(struct {
		ResourceType string `json:"resourceType"`
		embedded
	}{r.ResourceType(), embedded(r)})
}

type NotFoundError struct {
	ResourceType string
	ID           string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("resource type %s with ID %s not found", e.ResourceType, e.ID)
}

func (e NotFoundError) StatusCode() int {
	return 404
}

func (e NotFoundError) OperationOutcome() model.Resource {
	return BasicOperationOutcome{
		Issue: []BasicOperationOutcomeIssue{
			{Severity: "error", Code: "not-found"},
		},
	}
}
