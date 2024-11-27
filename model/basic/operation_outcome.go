package basic

import "encoding/json"

type OperationOutcome struct {
	Issue []OperationOutcomeIssue `json:"issue,omitempty"`
}

func (r OperationOutcome) ResourceType() string {
	return "OperationOutcome"
}

func (r OperationOutcome) ResourceId() (string, bool) {
	return "", false
}

func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	type embedded OperationOutcome
	return json.Marshal(struct {
		ResourceType string `json:"resourceType,omitempty"`
		embedded
	}{r.ResourceType(), embedded(r)})
}

type OperationOutcomeIssue struct {
	Severity    string `json:"severity,omitempty"`
	Code        string `json:"code,omitempty"`
	Diagnostics string `json:"diagnostics,omitempty"`
}
