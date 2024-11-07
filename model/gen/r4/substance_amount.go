package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for SubstanceAmount Type: Chemical substances are a single substance type whose primary defining element is the molecular structure. Chemical substances shall be defined on the basis of their complete covalent molecular structure; the presence of a salt (counter-ion) and/or solvates (water, alcohols) is also captured. Purity, grade, physical form or particle size are not taken into account in the definition of a chemical substance or in the assignment of a Substance ID.
type SubstanceAmount struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Used to capture quantitative values for a variety of elements. If only limits are given, the arithmetic mean would be the average. If only a single definite value for a given element is given, it would be captured in this field.
	Amount isSubstanceAmountAmount
	// Most elements that require a quantitative value will also have a field called amount type. Amount type should always be specified because the actual value of the amount is often dependent on it. EXAMPLE: In capturing the actual relative amounts of substances or molecular fragments it is essential to indicate whether the amount refers to a mole ratio or weight ratio. For any given element an effort should be made to use same the amount type for all related definitional elements.
	AmountType *CodeableConcept
	// A textual comment on a numeric value.
	AmountText *String
	// Reference range of possible or expected values.
	ReferenceRange *SubstanceAmountReferenceRange
}
type isSubstanceAmountAmount interface {
	isSubstanceAmountAmount()
}

func (r Quantity) isSubstanceAmountAmount() {}
func (r Range) isSubstanceAmountAmount()    {}
func (r String) isSubstanceAmountAmount()   {}

type jsonSubstanceAmount struct {
	Id                           *string                        `json:"id,omitempty"`
	Extension                    []Extension                    `json:"extension,omitempty"`
	ModifierExtension            []Extension                    `json:"modifierExtension,omitempty"`
	AmountQuantity               *Quantity                      `json:"amountQuantity,omitempty"`
	AmountRange                  *Range                         `json:"amountRange,omitempty"`
	AmountString                 *String                        `json:"amountString,omitempty"`
	AmountStringPrimitiveElement *primitiveElement              `json:"_amountString,omitempty"`
	AmountType                   *CodeableConcept               `json:"amountType,omitempty"`
	AmountText                   *String                        `json:"amountText,omitempty"`
	AmountTextPrimitiveElement   *primitiveElement              `json:"_amountText,omitempty"`
	ReferenceRange               *SubstanceAmountReferenceRange `json:"referenceRange,omitempty"`
}

func (r SubstanceAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceAmount) marshalJSON() jsonSubstanceAmount {
	m := jsonSubstanceAmount{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Amount.(type) {
	case Quantity:
		m.AmountQuantity = &v
	case *Quantity:
		m.AmountQuantity = v
	case Range:
		m.AmountRange = &v
	case *Range:
		m.AmountRange = v
	case String:
		if v.Value != nil {
			m.AmountString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.AmountString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.AmountType = r.AmountType
	if r.AmountText != nil && r.AmountText.Value != nil {
		m.AmountText = r.AmountText
	}
	if r.AmountText != nil && (r.AmountText.Id != nil || r.AmountText.Extension != nil) {
		m.AmountTextPrimitiveElement = &primitiveElement{Id: r.AmountText.Id, Extension: r.AmountText.Extension}
	}
	m.ReferenceRange = r.ReferenceRange
	return m
}
func (r *SubstanceAmount) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceAmount
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceAmount) unmarshalJSON(m jsonSubstanceAmount) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.AmountQuantity != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountQuantity
		r.Amount = v
	}
	if m.AmountRange != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountRange
		r.Amount = v
	}
	if m.AmountString != nil || m.AmountStringPrimitiveElement != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountString
		if m.AmountStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AmountStringPrimitiveElement.Id
			v.Extension = m.AmountStringPrimitiveElement.Extension
		}
		r.Amount = v
	}
	r.AmountType = m.AmountType
	r.AmountText = m.AmountText
	if m.AmountTextPrimitiveElement != nil {
		if r.AmountText == nil {
			r.AmountText = &String{}
		}
		r.AmountText.Id = m.AmountTextPrimitiveElement.Id
		r.AmountText.Extension = m.AmountTextPrimitiveElement.Extension
	}
	r.ReferenceRange = m.ReferenceRange
	return nil
}
func (r SubstanceAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	switch v := r.Amount.(type) {
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountQuantity"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountRange"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.AmountType, xml.StartElement{Name: xml.Name{Local: "amountType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AmountText, xml.StartElement{Name: xml.Name{Local: "amountText"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceRange, xml.StartElement{Name: xml.Name{Local: "referenceRange"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceAmount) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "amountQuantity":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountRange":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountString":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AmountType = &v
			case "amountText":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AmountText = &v
			case "referenceRange":
				var v SubstanceAmountReferenceRange
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceRange = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceAmount) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Reference range of possible or expected values.
type SubstanceAmountReferenceRange struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Lower limit possible or expected.
	LowLimit *Quantity
	// Upper limit possible or expected.
	HighLimit *Quantity
}
type jsonSubstanceAmountReferenceRange struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	LowLimit  *Quantity   `json:"lowLimit,omitempty"`
	HighLimit *Quantity   `json:"highLimit,omitempty"`
}

func (r SubstanceAmountReferenceRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceAmountReferenceRange) marshalJSON() jsonSubstanceAmountReferenceRange {
	m := jsonSubstanceAmountReferenceRange{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.LowLimit = r.LowLimit
	m.HighLimit = r.HighLimit
	return m
}
func (r *SubstanceAmountReferenceRange) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceAmountReferenceRange
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceAmountReferenceRange) unmarshalJSON(m jsonSubstanceAmountReferenceRange) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.LowLimit = m.LowLimit
	r.HighLimit = m.HighLimit
	return nil
}
func (r SubstanceAmountReferenceRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.LowLimit, xml.StartElement{Name: xml.Name{Local: "lowLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HighLimit, xml.StartElement{Name: xml.Name{Local: "highLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceAmountReferenceRange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "lowLimit":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LowLimit = &v
			case "highLimit":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HighLimit = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceAmountReferenceRange) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
