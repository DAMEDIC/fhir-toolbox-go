package r4

import "encoding/json"

// Base StructureDefinition for Ratio Type: A relationship of two Quantity values - expressed as a numerator and a denominator.
//
// Need to able to capture ratios for some measurements (titers) and some rates (costs).
type Ratio struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The value of the numerator.
	Numerator *Quantity
	// The value of the denominator.
	Denominator *Quantity
}
type jsonRatio struct {
	Id          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Numerator   *Quantity   `json:"numerator,omitempty"`
	Denominator *Quantity   `json:"denominator,omitempty"`
}

func (r Ratio) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Ratio) marshalJSON() jsonRatio {
	m := jsonRatio{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Numerator = r.Numerator
	m.Denominator = r.Denominator
	return m
}
func (r *Ratio) UnmarshalJSON(b []byte) error {
	var m jsonRatio
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Ratio) unmarshalJSON(m jsonRatio) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Numerator = m.Numerator
	r.Denominator = m.Denominator
	return nil
}
func (r Ratio) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
