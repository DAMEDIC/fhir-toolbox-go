package raw

import "encoding/json"

// raw.Resource is a raw encoded value.
// This can be used to pass-through FHIR resources without actually deserializing.
type Resource struct {
	typ      string
	id       *string
	Resource json.RawMessage
}

func (r Resource) ResourceType() string {
	return r.typ
}

func (r Resource) ResourceId() (string, bool) {
	if r.id == nil {
		return "", false
	}
	return *r.id, true
}

func (r *Resource) UnmarshalJSON(b []byte) error {
	var t struct {
		ResourceType string  `json:"resourceType"`
		Id           *string `json:"id,omitempty"`
	}
	var raw json.RawMessage

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	*r = Resource{
		typ:      t.ResourceType,
		Resource: raw,
	}
	if t.Id != nil {
		r.id = t.Id
	}

	return nil
}
