package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for Age Type: A duration of time during which an organism (or a process) has existed.
type Age struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *Decimal
	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *Code
	// A human-readable form of the unit.
	Unit *String
	// The identification of the system that provides the coded form of the unit.
	System *Uri
	// A computer processable form of the unit in some unit representation system.
	Code *Code
}
type jsonAge struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	Value                      *Decimal          `json:"value,omitempty"`
	ValuePrimitiveElement      *primitiveElement `json:"_value,omitempty"`
	Comparator                 *Code             `json:"comparator,omitempty"`
	ComparatorPrimitiveElement *primitiveElement `json:"_comparator,omitempty"`
	Unit                       *String           `json:"unit,omitempty"`
	UnitPrimitiveElement       *primitiveElement `json:"_unit,omitempty"`
	System                     *Uri              `json:"system,omitempty"`
	SystemPrimitiveElement     *primitiveElement `json:"_system,omitempty"`
	Code                       *Code             `json:"code,omitempty"`
	CodePrimitiveElement       *primitiveElement `json:"_code,omitempty"`
}

func (r Age) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Age) marshalJSON() jsonAge {
	m := jsonAge{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Value != nil && r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	if r.Comparator != nil && r.Comparator.Value != nil {
		m.Comparator = r.Comparator
	}
	if r.Comparator != nil && (r.Comparator.Id != nil || r.Comparator.Extension != nil) {
		m.ComparatorPrimitiveElement = &primitiveElement{Id: r.Comparator.Id, Extension: r.Comparator.Extension}
	}
	if r.Unit != nil && r.Unit.Value != nil {
		m.Unit = r.Unit
	}
	if r.Unit != nil && (r.Unit.Id != nil || r.Unit.Extension != nil) {
		m.UnitPrimitiveElement = &primitiveElement{Id: r.Unit.Id, Extension: r.Unit.Extension}
	}
	if r.System != nil && r.System.Value != nil {
		m.System = r.System
	}
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	if r.Code != nil && r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	return m
}
func (r *Age) UnmarshalJSON(b []byte) error {
	var m jsonAge
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Age) unmarshalJSON(m jsonAge) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		if r.Value == nil {
			r.Value = &Decimal{}
		}
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Comparator = m.Comparator
	if m.ComparatorPrimitiveElement != nil {
		if r.Comparator == nil {
			r.Comparator = &Code{}
		}
		r.Comparator.Id = m.ComparatorPrimitiveElement.Id
		r.Comparator.Extension = m.ComparatorPrimitiveElement.Extension
	}
	r.Unit = m.Unit
	if m.UnitPrimitiveElement != nil {
		if r.Unit == nil {
			r.Unit = &String{}
		}
		r.Unit.Id = m.UnitPrimitiveElement.Id
		r.Unit.Extension = m.UnitPrimitiveElement.Extension
	}
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		if r.System == nil {
			r.System = &Uri{}
		}
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		if r.Code == nil {
			r.Code = &Code{}
		}
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	return nil
}
func (r Age) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comparator, xml.StartElement{Name: xml.Name{Local: "comparator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Unit, xml.StartElement{Name: xml.Name{Local: "unit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.System, xml.StartElement{Name: xml.Name{Local: "system"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Age) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "value":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = &v
			case "comparator":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comparator = &v
			case "unit":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Unit = &v
			case "system":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.System = &v
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Age) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
