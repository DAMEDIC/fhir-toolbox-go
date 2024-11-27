package basic

import (
	"encoding/json"
	"fhir-toolbox/model"
)

type Bundle struct {
	Type  string        `json:"type,omitempty"`
	Link  []BundleLink  `json:"link,omitempty"`
	Entry []BundleEntry `json:"entry,omitempty"`
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
		ResourceType string `json:"resourceType,omitempty"`
		embedded
	}{b.ResourceType(), embedded(b)})
}

type BundleLink struct {
	Relation string `json:"relation,omitempty"`
	Url      string `json:"url,omitempty"`
}

type BundleEntry struct {
	FullUrl  string            `json:"fullUrl,omitempty"`
	Resource model.Resource    `json:"resource,omitempty"`
	Search   BundleEntrySearch `json:"search,omitempty"`
}

type BundleEntrySearch struct {
	Mode string `json:"mode,omitempty"`
}
