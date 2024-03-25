package r4

import "encoding/json"

// Base StructureDefinition for Range Type: A set of ordered Quantities defined by a low and high limit.
//
// Need to be able to specify ranges of values.
type Range struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The low limit. The boundary is inclusive.
	Low *Quantity
	// The high limit. The boundary is inclusive.
	High *Quantity
}
type jsonRange struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Low       *Quantity   `json:"low,omitempty"`
	High      *Quantity   `json:"high,omitempty"`
}

func (r Range) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Range) marshalJSON() jsonRange {
	m := jsonRange{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Low = r.Low
	m.High = r.High
	return m
}
func (r *Range) UnmarshalJSON(b []byte) error {
	var m jsonRange
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Range) unmarshalJSON(m jsonRange) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Low = m.Low
	r.High = m.High
	return nil
}
func (r Range) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
