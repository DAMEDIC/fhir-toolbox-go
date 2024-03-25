package r4

import "encoding/json"

// Base StructureDefinition for Reference Type: A reference from one resource to another.
type Reference struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	Reference *String
	// The expected type of the target of the reference. If both Reference.type and Reference.reference are populated and Reference.reference is a FHIR URL, both SHALL be consistent.
	//
	// The type is the Canonical URL of Resource Definition that is the type this reference refers to. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	Type *Uri
	// An identifier for the target resource. This is used when there is no way to reference the other resource directly, either because the entity it represents is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	Identifier *Identifier
	// Plain text narrative that identifies the resource in addition to the resource reference.
	Display *String
}
type jsonReference struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	Reference                 *String           `json:"reference,omitempty"`
	ReferencePrimitiveElement *primitiveElement `json:"_reference,omitempty"`
	Type                      *Uri              `json:"type,omitempty"`
	TypePrimitiveElement      *primitiveElement `json:"_type,omitempty"`
	Identifier                *Identifier       `json:"identifier,omitempty"`
	Display                   *String           `json:"display,omitempty"`
	DisplayPrimitiveElement   *primitiveElement `json:"_display,omitempty"`
}

func (r Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Reference) marshalJSON() jsonReference {
	m := jsonReference{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Reference = r.Reference
	if r.Reference != nil && (r.Reference.Id != nil || r.Reference.Extension != nil) {
		m.ReferencePrimitiveElement = &primitiveElement{Id: r.Reference.Id, Extension: r.Reference.Extension}
	}
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Identifier = r.Identifier
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	return m
}
func (r *Reference) UnmarshalJSON(b []byte) error {
	var m jsonReference
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Reference) unmarshalJSON(m jsonReference) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Reference = m.Reference
	if m.ReferencePrimitiveElement != nil {
		r.Reference.Id = m.ReferencePrimitiveElement.Id
		r.Reference.Extension = m.ReferencePrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	return nil
}
func (r Reference) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
