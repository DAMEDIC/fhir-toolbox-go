package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for Money Type: An amount of economic utility in some recognized currency.
type Money struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Numerical value (with implicit precision).
	Value *Decimal
	// ISO 4217 Currency Code.
	Currency *Code
}
type jsonMoney struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	Value                    *Decimal          `json:"value,omitempty"`
	ValuePrimitiveElement    *primitiveElement `json:"_value,omitempty"`
	Currency                 *Code             `json:"currency,omitempty"`
	CurrencyPrimitiveElement *primitiveElement `json:"_currency,omitempty"`
}

func (r Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Money) marshalJSON() jsonMoney {
	m := jsonMoney{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Value != nil && r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	if r.Currency != nil && r.Currency.Value != nil {
		m.Currency = r.Currency
	}
	if r.Currency != nil && (r.Currency.Id != nil || r.Currency.Extension != nil) {
		m.CurrencyPrimitiveElement = &primitiveElement{Id: r.Currency.Id, Extension: r.Currency.Extension}
	}
	return m
}
func (r *Money) UnmarshalJSON(b []byte) error {
	var m jsonMoney
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Money) unmarshalJSON(m jsonMoney) error {
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
	r.Currency = m.Currency
	if m.CurrencyPrimitiveElement != nil {
		if r.Currency == nil {
			r.Currency = &Code{}
		}
		r.Currency.Id = m.CurrencyPrimitiveElement.Id
		r.Currency.Extension = m.CurrencyPrimitiveElement.Extension
	}
	return nil
}
func (r Money) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Currency, xml.StartElement{Name: xml.Name{Local: "currency"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Money) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "currency":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Currency = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Money) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
