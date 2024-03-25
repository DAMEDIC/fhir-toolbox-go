package r4

import "encoding/json"

// Base StructureDefinition for ParameterDefinition Type: The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
type ParameterDefinition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
	Name *Code
	// Whether the parameter is input or output for the module.
	Use Code
	// The minimum number of times this parameter SHALL appear in the request or response.
	Min *Integer
	// The maximum number of times this element is permitted to appear in the request or response.
	Max *String
	// A brief discussion of what the parameter is for and how it is used by the module.
	Documentation *String
	// The type of the parameter.
	Type Code
	// If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
	Profile *Canonical
}
type jsonParameterDefinition struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	Name                          *Code             `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement `json:"_name,omitempty"`
	Use                           Code              `json:"use,omitempty"`
	UsePrimitiveElement           *primitiveElement `json:"_use,omitempty"`
	Min                           *Integer          `json:"min,omitempty"`
	MinPrimitiveElement           *primitiveElement `json:"_min,omitempty"`
	Max                           *String           `json:"max,omitempty"`
	MaxPrimitiveElement           *primitiveElement `json:"_max,omitempty"`
	Documentation                 *String           `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
	Type                          Code              `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement `json:"_type,omitempty"`
	Profile                       *Canonical        `json:"profile,omitempty"`
	ProfilePrimitiveElement       *primitiveElement `json:"_profile,omitempty"`
}

func (r ParameterDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ParameterDefinition) marshalJSON() jsonParameterDefinition {
	m := jsonParameterDefinition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Use = r.Use
	if r.Use.Id != nil || r.Use.Extension != nil {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	m.Min = r.Min
	if r.Min != nil && (r.Min.Id != nil || r.Min.Extension != nil) {
		m.MinPrimitiveElement = &primitiveElement{Id: r.Min.Id, Extension: r.Min.Extension}
	}
	m.Max = r.Max
	if r.Max != nil && (r.Max.Id != nil || r.Max.Extension != nil) {
		m.MaxPrimitiveElement = &primitiveElement{Id: r.Max.Id, Extension: r.Max.Extension}
	}
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Profile = r.Profile
	if r.Profile != nil && (r.Profile.Id != nil || r.Profile.Extension != nil) {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	return m
}
func (r *ParameterDefinition) UnmarshalJSON(b []byte) error {
	var m jsonParameterDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ParameterDefinition) unmarshalJSON(m jsonParameterDefinition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Min = m.Min
	if m.MinPrimitiveElement != nil {
		r.Min.Id = m.MinPrimitiveElement.Id
		r.Min.Extension = m.MinPrimitiveElement.Extension
	}
	r.Max = m.Max
	if m.MaxPrimitiveElement != nil {
		r.Max.Id = m.MaxPrimitiveElement.Id
		r.Max.Extension = m.MaxPrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	return nil
}
func (r ParameterDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
