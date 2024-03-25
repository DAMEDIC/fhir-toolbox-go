package r4

import "encoding/json"

// Base StructureDefinition for Narrative Type: A human-readable summary of the resource conveying the essential clinical and business information for the resource.
type Narrative struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	Status Code
	// The actual narrative content, a stripped down version of XHTML.
	Div Xhtml
}
type jsonNarrative struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	Status                 Code              `json:"status,omitempty"`
	StatusPrimitiveElement *primitiveElement `json:"_status,omitempty"`
	Div                    Xhtml             `json:"div,omitempty"`
	DivPrimitiveElement    *primitiveElement `json:"_div,omitempty"`
}

func (r Narrative) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Narrative) marshalJSON() jsonNarrative {
	m := jsonNarrative{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Div = r.Div
	if r.Div.Id != nil {
		m.DivPrimitiveElement = &primitiveElement{Id: r.Div.Id}
	}
	return m
}
func (r *Narrative) UnmarshalJSON(b []byte) error {
	var m jsonNarrative
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Narrative) unmarshalJSON(m jsonNarrative) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Div = m.Div
	if m.DivPrimitiveElement != nil {
		r.Div.Id = m.DivPrimitiveElement.Id
	}
	return nil
}
func (r Narrative) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
