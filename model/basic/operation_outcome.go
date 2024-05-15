package basic

import "encoding/json"

type OperationOutcome struct {
	Issue []OperationOutcomeIssue `json:"issue"`
}

type OperationOutcomeIssue struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
}

func (r OperationOutcome) ResourceType() string {
	return "OperationOutcome"
}

func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	type embedded OperationOutcome
	return json.Marshal(struct {
		ResourceType string `json:"resourceType"`
		embedded
	}{r.ResourceType(), embedded(r)})
}
