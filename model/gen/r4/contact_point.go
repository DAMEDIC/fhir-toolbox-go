package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

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
	if r.Use != nil && r.Use.Value != nil {
		m.Use = r.Use
	}
	if r.Use != nil && (r.Use.Id != nil || r.Use.Extension != nil) {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	if r.Rank != nil && r.Rank.Value != nil {
		m.Rank = r.Rank
	}
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
		if r.System == nil {
			r.System = &Code{}
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
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		if r.Use == nil {
			r.Use = &Code{}
		}
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Rank = m.Rank
	if m.RankPrimitiveElement != nil {
		if r.Rank == nil {
			r.Rank = &PositiveInt{}
		}
		r.Rank.Id = m.RankPrimitiveElement.Id
		r.Rank.Extension = m.RankPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r ContactPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.System, xml.StartElement{Name: xml.Name{Local: "system"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Use, xml.StartElement{Name: xml.Name{Local: "use"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Rank, xml.StartElement{Name: xml.Name{Local: "rank"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContactPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "system":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.System = &v
			case "value":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = &v
			case "use":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Use = &v
			case "rank":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rank = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContactPoint) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
