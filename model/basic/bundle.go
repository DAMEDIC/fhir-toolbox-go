package basic

import (
	"encoding/json"
	"fhir-toolbox/generate/model"
)

type Bundle struct {
	Type  string        `json:"type"`
	Link  []BundleLink  `json:"link"`
	Entry []BundleEntry `json:"entry"`
}

func (b Bundle) ResourceType() string {
	return "Bundle"
}

func (b Bundle) ResourceId() (string, bool) {
	return "", false
}

func (b Bundle) MarshalJSON() ([]byte, error) {
	type embedded Bundle
	return json.Marshal(struct {
		ResourceType string `json:"resourceType"`
		embedded
	}{b.ResourceType(), embedded(b)})
}

type BundleLink struct {
	Relation string `json:"relation"`
	Url      string `json:"url"`
}

type BundleEntry struct {
	FullUrl  string            `json:"fullUrl"`
	Resource model.Resource    `json:"resource"`
	Search   BundleEntrySearch `json:"search"`
}

type BundleEntrySearch struct {
	Mode string `json:"mode"`
}
