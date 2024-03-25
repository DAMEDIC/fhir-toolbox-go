package r4

import "encoding/json"

// Base StructureDefinition for Contributor Type: A contributor to the content of a knowledge asset, including authors, editors, reviewers, and endorsers.
//
// Need to track contributor information in the same way across multiple resources.
type Contributor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The type of contributor.
	Type Code
	// The name of the individual or organization responsible for the contribution.
	Name String
	// Contact details to assist a user in finding and communicating with the contributor.
	Contact []ContactDetail
}
type jsonContributor struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
	Name                 String            `json:"name,omitempty"`
	NamePrimitiveElement *primitiveElement `json:"_name,omitempty"`
	Contact              []ContactDetail   `json:"contact,omitempty"`
}

func (r Contributor) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Contributor) marshalJSON() jsonContributor {
	m := jsonContributor{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Contact = r.Contact
	return m
}
func (r *Contributor) UnmarshalJSON(b []byte) error {
	var m jsonContributor
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Contributor) unmarshalJSON(m jsonContributor) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Contact = m.Contact
	return nil
}
func (r Contributor) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
