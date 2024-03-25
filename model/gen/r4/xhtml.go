package r4

import "encoding/json"

// Base StructureDefinition for xhtml Type
type Xhtml struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// Actual xhtml
	Value string
}
type jsonXhtml struct {
	Id    *string `json:"id,omitempty"`
	Value string  `json:"value,omitempty"`
}

func (r Xhtml) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *Xhtml) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = Xhtml{Value: value}
	return nil
}
