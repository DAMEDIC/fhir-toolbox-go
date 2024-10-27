package r4

import "encoding/json"

// Base StructureDefinition for Period Type: A time period defined by a start and end date and optionally time.
type Period struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The start of the period. The boundary is inclusive.
	Start *DateTime
	// The end of the period. If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	End *DateTime
}
type jsonPeriod struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	Start                 *DateTime         `json:"start,omitempty"`
	StartPrimitiveElement *primitiveElement `json:"_start,omitempty"`
	End                   *DateTime         `json:"end,omitempty"`
	EndPrimitiveElement   *primitiveElement `json:"_end,omitempty"`
}

func (r Period) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Period) marshalJSON() jsonPeriod {
	m := jsonPeriod{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Start != nil && r.Start.Value != nil {
		m.Start = r.Start
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	if r.End != nil && r.End.Value != nil {
		m.End = r.End
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	return m
}
func (r *Period) UnmarshalJSON(b []byte) error {
	var m jsonPeriod
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Period) unmarshalJSON(m jsonPeriod) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		if r.Start == nil {
			r.Start = &DateTime{}
		}
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		if r.End == nil {
			r.End = &DateTime{}
		}
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	return nil
}
func (r Period) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
