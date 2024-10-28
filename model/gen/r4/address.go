package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
//
// Need to be able to record postal addresses, along with notes about their use.
type Address struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The purpose of this address.
	Use *Code
	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type *Code
	// Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	Text *String
	// This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	Line []String
	// The name of the city, town, suburb, village or other community or delivery center.
	City *String
	// The name of the administrative area (county).
	District *String
	// Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	State *String
	// A postal code designating a region defined by the postal service.
	PostalCode *String
	// Country - a nation as commonly understood or generally accepted.
	Country *String
	// Time period when address was/is in use.
	Period *Period
}
type jsonAddress struct {
	Id                         *string             `json:"id,omitempty"`
	Extension                  []Extension         `json:"extension,omitempty"`
	Use                        *Code               `json:"use,omitempty"`
	UsePrimitiveElement        *primitiveElement   `json:"_use,omitempty"`
	Type                       *Code               `json:"type,omitempty"`
	TypePrimitiveElement       *primitiveElement   `json:"_type,omitempty"`
	Text                       *String             `json:"text,omitempty"`
	TextPrimitiveElement       *primitiveElement   `json:"_text,omitempty"`
	Line                       []String            `json:"line,omitempty"`
	LinePrimitiveElement       []*primitiveElement `json:"_line,omitempty"`
	City                       *String             `json:"city,omitempty"`
	CityPrimitiveElement       *primitiveElement   `json:"_city,omitempty"`
	District                   *String             `json:"district,omitempty"`
	DistrictPrimitiveElement   *primitiveElement   `json:"_district,omitempty"`
	State                      *String             `json:"state,omitempty"`
	StatePrimitiveElement      *primitiveElement   `json:"_state,omitempty"`
	PostalCode                 *String             `json:"postalCode,omitempty"`
	PostalCodePrimitiveElement *primitiveElement   `json:"_postalCode,omitempty"`
	Country                    *String             `json:"country,omitempty"`
	CountryPrimitiveElement    *primitiveElement   `json:"_country,omitempty"`
	Period                     *Period             `json:"period,omitempty"`
}

func (r Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Address) marshalJSON() jsonAddress {
	m := jsonAddress{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Use != nil && r.Use.Value != nil {
		m.Use = r.Use
	}
	if r.Use != nil && (r.Use.Id != nil || r.Use.Extension != nil) {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	anyLineValue := false
	for _, e := range r.Line {
		if e.Value != nil {
			anyLineValue = true
			break
		}
	}
	if anyLineValue {
		m.Line = r.Line
	}
	anyLineIdOrExtension := false
	for _, e := range r.Line {
		if e.Id != nil || e.Extension != nil {
			anyLineIdOrExtension = true
			break
		}
	}
	if anyLineIdOrExtension {
		m.LinePrimitiveElement = make([]*primitiveElement, 0, len(r.Line))
		for _, e := range r.Line {
			if e.Id != nil || e.Extension != nil {
				m.LinePrimitiveElement = append(m.LinePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LinePrimitiveElement = append(m.LinePrimitiveElement, nil)
			}
		}
	}
	if r.City != nil && r.City.Value != nil {
		m.City = r.City
	}
	if r.City != nil && (r.City.Id != nil || r.City.Extension != nil) {
		m.CityPrimitiveElement = &primitiveElement{Id: r.City.Id, Extension: r.City.Extension}
	}
	if r.District != nil && r.District.Value != nil {
		m.District = r.District
	}
	if r.District != nil && (r.District.Id != nil || r.District.Extension != nil) {
		m.DistrictPrimitiveElement = &primitiveElement{Id: r.District.Id, Extension: r.District.Extension}
	}
	if r.State != nil && r.State.Value != nil {
		m.State = r.State
	}
	if r.State != nil && (r.State.Id != nil || r.State.Extension != nil) {
		m.StatePrimitiveElement = &primitiveElement{Id: r.State.Id, Extension: r.State.Extension}
	}
	if r.PostalCode != nil && r.PostalCode.Value != nil {
		m.PostalCode = r.PostalCode
	}
	if r.PostalCode != nil && (r.PostalCode.Id != nil || r.PostalCode.Extension != nil) {
		m.PostalCodePrimitiveElement = &primitiveElement{Id: r.PostalCode.Id, Extension: r.PostalCode.Extension}
	}
	if r.Country != nil && r.Country.Value != nil {
		m.Country = r.Country
	}
	if r.Country != nil && (r.Country.Id != nil || r.Country.Extension != nil) {
		m.CountryPrimitiveElement = &primitiveElement{Id: r.Country.Id, Extension: r.Country.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *Address) UnmarshalJSON(b []byte) error {
	var m jsonAddress
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Address) unmarshalJSON(m jsonAddress) error {
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
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &Code{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Line = m.Line
	for i, e := range m.LinePrimitiveElement {
		if len(r.Line) <= i {
			r.Line = append(r.Line, String{})
		}
		if e != nil {
			r.Line[i].Id = e.Id
			r.Line[i].Extension = e.Extension
		}
	}
	r.City = m.City
	if m.CityPrimitiveElement != nil {
		if r.City == nil {
			r.City = &String{}
		}
		r.City.Id = m.CityPrimitiveElement.Id
		r.City.Extension = m.CityPrimitiveElement.Extension
	}
	r.District = m.District
	if m.DistrictPrimitiveElement != nil {
		if r.District == nil {
			r.District = &String{}
		}
		r.District.Id = m.DistrictPrimitiveElement.Id
		r.District.Extension = m.DistrictPrimitiveElement.Extension
	}
	r.State = m.State
	if m.StatePrimitiveElement != nil {
		if r.State == nil {
			r.State = &String{}
		}
		r.State.Id = m.StatePrimitiveElement.Id
		r.State.Extension = m.StatePrimitiveElement.Extension
	}
	r.PostalCode = m.PostalCode
	if m.PostalCodePrimitiveElement != nil {
		if r.PostalCode == nil {
			r.PostalCode = &String{}
		}
		r.PostalCode.Id = m.PostalCodePrimitiveElement.Id
		r.PostalCode.Extension = m.PostalCodePrimitiveElement.Extension
	}
	r.Country = m.Country
	if m.CountryPrimitiveElement != nil {
		if r.Country == nil {
			r.Country = &String{}
		}
		r.Country.Id = m.CountryPrimitiveElement.Id
		r.Country.Extension = m.CountryPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r Address) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Use, xml.StartElement{Name: xml.Name{Local: "use"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Line, xml.StartElement{Name: xml.Name{Local: "line"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.City, xml.StartElement{Name: xml.Name{Local: "city"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.District, xml.StartElement{Name: xml.Name{Local: "district"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.State, xml.StartElement{Name: xml.Name{Local: "state"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PostalCode, xml.StartElement{Name: xml.Name{Local: "postalCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Country, xml.StartElement{Name: xml.Name{Local: "country"}})
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
func (r *Address) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "use":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Use = &v
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "line":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Line = append(r.Line, v)
			case "city":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.City = &v
			case "district":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.District = &v
			case "state":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.State = &v
			case "postalCode":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PostalCode = &v
			case "country":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Country = &v
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
func (r Address) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
