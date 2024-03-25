package r4

import "encoding/json"

// Base StructureDefinition for CodeableConcept Type: A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
//
// This is a common pattern in healthcare - a concept that may be defined by one or more codes from formal definitions including LOINC and SNOMED CT, and/or defined by the provision of text that captures a human sense of the concept.
type CodeableConcept struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A reference to a code defined by a terminology system.
	Coding []Coding
	// A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	Text *String
}
type jsonCodeableConcept struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	Coding               []Coding          `json:"coding,omitempty"`
	Text                 *String           `json:"text,omitempty"`
	TextPrimitiveElement *primitiveElement `json:"_text,omitempty"`
}

func (r CodeableConcept) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeableConcept) marshalJSON() jsonCodeableConcept {
	m := jsonCodeableConcept{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Coding = r.Coding
	m.Text = r.Text
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	return m
}
func (r *CodeableConcept) UnmarshalJSON(b []byte) error {
	var m jsonCodeableConcept
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeableConcept) unmarshalJSON(m jsonCodeableConcept) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Coding = m.Coding
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	return nil
}
func (r CodeableConcept) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
