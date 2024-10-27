package r4

import "encoding/json"

// Base StructureDefinition for Identifier Type: An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
//
// Need to be able to identify things with confidence and be sure that the identification is not subject to misinterpretation.
type Identifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The purpose of this identifier.
	Use *Code
	// A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	Type *CodeableConcept
	// Establishes the namespace for the value - that is, a URL that describes a set values that are unique.
	System *Uri
	// The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Value *String
	// Time period during which identifier is/was valid for use.
	Period *Period
	// Organization that issued/manages the identifier.
	Assigner *Reference
}
type jsonIdentifier struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	Use                    *Code             `json:"use,omitempty"`
	UsePrimitiveElement    *primitiveElement `json:"_use,omitempty"`
	Type                   *CodeableConcept  `json:"type,omitempty"`
	System                 *Uri              `json:"system,omitempty"`
	SystemPrimitiveElement *primitiveElement `json:"_system,omitempty"`
	Value                  *String           `json:"value,omitempty"`
	ValuePrimitiveElement  *primitiveElement `json:"_value,omitempty"`
	Period                 *Period           `json:"period,omitempty"`
	Assigner               *Reference        `json:"assigner,omitempty"`
}

func (r Identifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Identifier) marshalJSON() jsonIdentifier {
	m := jsonIdentifier{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Use != nil && r.Use.Value != nil {
		m.Use = r.Use
	}
	if r.Use != nil && (r.Use.Id != nil || r.Use.Extension != nil) {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	m.Type = r.Type
	if r.System != nil && r.System.Value != nil {
		m.System = r.System
	}
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	if r.Value != nil && r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.Period = r.Period
	m.Assigner = r.Assigner
	return m
}
func (r *Identifier) UnmarshalJSON(b []byte) error {
	var m jsonIdentifier
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Identifier) unmarshalJSON(m jsonIdentifier) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		if r.Use == nil {
			r.Use = &Code{}
		}
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		if r.System == nil {
			r.System = &Uri{}
		}
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		if r.Value == nil {
			r.Value = &String{}
		}
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Period = m.Period
	r.Assigner = m.Assigner
	return nil
}
func (r Identifier) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
