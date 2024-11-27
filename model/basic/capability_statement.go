package basic

import "encoding/json"

type CapabilityStatement struct {
	Status         string                            `json:"status,omitempty"`
	Date           string                            `json:"date,omitempty"`
	Kind           string                            `json:"kind,omitempty"`
	Software       CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation CapabilityStatementImplementation `json:"implementation,omitempty"`
	FHIRVersion    string                            `json:"fhirVersion,omitempty"`
	Format         []string                          `json:"format,omitempty"`
	Rest           []CapabilityStatementRest         `json:"rest,omitempty"`
}

func (r CapabilityStatement) ResourceType() string {
	return "CapabilityStatement"
}

func (r CapabilityStatement) ResourceId() (string, bool) {
	return "", false
}

func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	type embedded CapabilityStatement
	return json.Marshal(struct {
		ResourceType string `json:"resourceType,omitempty"`
		embedded
	}{r.ResourceType(), embedded(r)})
}

type CapabilityStatementSoftware struct {
	Name string `json:"name,omitempty"`
}

type CapabilityStatementImplementation struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

type CapabilityStatementRest struct {
	Mode     string                            `json:"mode,omitempty"`
	Resource []CapabilityStatementRestResource `json:"resource,omitempty"`
}

type CapabilityStatementRestResource struct {
	Type          string                                       `json:"type,omitempty"`
	Interaction   []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	SearchInclude []string                                     `json:"searchInclude,omitempty"`
	SearchParam   []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
}

type CapabilityStatementRestResourceInteraction struct {
	Code string `json:"code,omitempty"`
}

type CapabilityStatementRestResourceSearchParam struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
