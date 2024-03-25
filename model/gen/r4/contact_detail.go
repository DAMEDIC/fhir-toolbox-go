package r4

import "encoding/json"

// Base StructureDefinition for ContactDetail Type: Specifies contact information for a person or organization.
//
// Need to track contact information in the same way across multiple resources.
type ContactDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The name of an individual to contact.
	Name *String
	// The contact details for the individual (if a name was provided) or the organization.
	Telecom []ContactPoint
}
type jsonContactDetail struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	Name                 *String           `json:"name,omitempty"`
	NamePrimitiveElement *primitiveElement `json:"_name,omitempty"`
	Telecom              []ContactPoint    `json:"telecom,omitempty"`
}

func (r ContactDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContactDetail) marshalJSON() jsonContactDetail {
	m := jsonContactDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Telecom = r.Telecom
	return m
}
func (r *ContactDetail) UnmarshalJSON(b []byte) error {
	var m jsonContactDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContactDetail) unmarshalJSON(m jsonContactDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Telecom = m.Telecom
	return nil
}
func (r ContactDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
