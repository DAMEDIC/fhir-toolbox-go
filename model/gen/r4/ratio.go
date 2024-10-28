package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

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
func (r Ratio) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Numerator, xml.StartElement{Name: xml.Name{Local: "numerator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Denominator, xml.StartElement{Name: xml.Name{Local: "denominator"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Ratio) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "numerator":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Numerator = &v
			case "denominator":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Denominator = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Ratio) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
