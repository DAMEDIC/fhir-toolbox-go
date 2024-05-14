package capabilities

import (
	"encoding/json"
	"fhir-toolbox/model"
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
	return json.Marshal(struct {
		ResourceType string `json:"resourceType"`
		BasicOperationOutcome
	}{r.ResourceType(), r})
}
