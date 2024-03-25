package r4

import "encoding/json"

// Base StructureDefinition for ContactPoint Type: Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
//
// Need to track phone, fax, mobile, sms numbers, email addresses, twitter tags, etc.
type ContactPoint struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Telecommunications form for contact point - what communications system is required to make use of the contact.
	System *Code
	// The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Value *String
	// Identifies the purpose for the contact point.
	Use *Code
	// Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	Rank *PositiveInt
	// Time period when the contact point was/is in use.
	Period *Period
}
type jsonContactPoint struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	System                 *Code             `json:"system,omitempty"`
	SystemPrimitiveElement *primitiveElement `json:"_system,omitempty"`
	Value                  *String           `json:"value,omitempty"`
	ValuePrimitiveElement  *primitiveElement `json:"_value,omitempty"`
	Use                    *Code             `json:"use,omitempty"`
	UsePrimitiveElement    *primitiveElement `json:"_use,omitempty"`
	Rank                   *PositiveInt      `json:"rank,omitempty"`
	RankPrimitiveElement   *primitiveElement `json:"_rank,omitempty"`
	Period                 *Period           `json:"period,omitempty"`
}

func (r ContactPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContactPoint) marshalJSON() jsonContactPoint {
	m := jsonContactPoint{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.System = r.System
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	m.Value = r.Value
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.Use = r.Use
	if r.Use != nil && (r.Use.Id != nil || r.Use.Extension != nil) {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	m.Rank = r.Rank
	if r.Rank != nil && (r.Rank.Id != nil || r.Rank.Extension != nil) {
		m.RankPrimitiveElement = &primitiveElement{Id: r.Rank.Id, Extension: r.Rank.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *ContactPoint) UnmarshalJSON(b []byte) error {
	var m jsonContactPoint
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContactPoint) unmarshalJSON(m jsonContactPoint) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Rank = m.Rank
	if m.RankPrimitiveElement != nil {
		r.Rank.Id = m.RankPrimitiveElement.Id
		r.Rank.Extension = m.RankPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r ContactPoint) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
