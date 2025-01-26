package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"io"
	"slices"
	"unsafe"
)

// Base StructureDefinition for ProdCharacteristic Type: The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type ProdCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Where applicable, the height can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Height *Quantity
	// Where applicable, the width can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Width *Quantity
	// Where applicable, the depth can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Depth *Quantity
	// Where applicable, the weight can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Weight *Quantity
	// Where applicable, the nominal volume can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	NominalVolume *Quantity
	// Where applicable, the external diameter can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	ExternalDiameter *Quantity
	// Where applicable, the shape can be specified An appropriate controlled vocabulary shall be used The term and the term identifier shall be used.
	Shape *String
	// Where applicable, the color can be specified An appropriate controlled vocabulary shall be used The term and the term identifier shall be used.
	Color []String
	// Where applicable, the imprint can be specified as text.
	Imprint []String
	// Where applicable, the image can be provided The format of the image attachment shall be specified by regional implementations.
	Image []Attachment
	// Where applicable, the scoring can be specified An appropriate controlled vocabulary shall be used The term and the term identifier shall be used.
	Scoring *CodeableConcept
}

func (r ProdCharacteristic) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Height != nil {
		s += r.Height.MemSize()
	}
	if r.Width != nil {
		s += r.Width.MemSize()
	}
	if r.Depth != nil {
		s += r.Depth.MemSize()
	}
	if r.Weight != nil {
		s += r.Weight.MemSize()
	}
	if r.NominalVolume != nil {
		s += r.NominalVolume.MemSize()
	}
	if r.ExternalDiameter != nil {
		s += r.ExternalDiameter.MemSize()
	}
	if r.Shape != nil {
		s += r.Shape.MemSize()
	}
	for _, i := range r.Color {
		s += i.MemSize()
	}
	s += (cap(r.Color) - len(r.Color)) * int(unsafe.Sizeof(String{}))
	for _, i := range r.Imprint {
		s += i.MemSize()
	}
	s += (cap(r.Imprint) - len(r.Imprint)) * int(unsafe.Sizeof(String{}))
	for _, i := range r.Image {
		s += i.MemSize()
	}
	s += (cap(r.Image) - len(r.Image)) * int(unsafe.Sizeof(Attachment{}))
	if r.Scoring != nil {
		s += r.Scoring.MemSize()
	}
	return s
}
func (r ProdCharacteristic) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ProdCharacteristic) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ProdCharacteristic) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Height != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"height\":"))
		if err != nil {
			return err
		}
		err = r.Height.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Width != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"width\":"))
		if err != nil {
			return err
		}
		err = r.Width.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Depth != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"depth\":"))
		if err != nil {
			return err
		}
		err = r.Depth.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Weight != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"weight\":"))
		if err != nil {
			return err
		}
		err = r.Weight.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.NominalVolume != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"nominalVolume\":"))
		if err != nil {
			return err
		}
		err = r.NominalVolume.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ExternalDiameter != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"externalDiameter\":"))
		if err != nil {
			return err
		}
		err = r.ExternalDiameter.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Shape != nil && r.Shape.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"shape\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Shape)
		if err != nil {
			return err
		}
	}
	if r.Shape != nil && (r.Shape.Id != nil || r.Shape.Extension != nil) {
		p := primitiveElement{Id: r.Shape.Id, Extension: r.Shape.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_shape\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.Color {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"color\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Color)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Color {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_color\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Color {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.Imprint {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"imprint\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Imprint)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Imprint {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_imprint\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Imprint {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if len(r.Image) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"image\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Image {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Scoring != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"scoring\":"))
		if err != nil {
			return err
		}
		err = r.Scoring.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *ProdCharacteristic) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ProdCharacteristic element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ProdCharacteristic element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "height":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Height = &v
		case "width":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Width = &v
		case "depth":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Depth = &v
		case "weight":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Weight = &v
		case "nominalVolume":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.NominalVolume = &v
		case "externalDiameter":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ExternalDiameter = &v
		case "shape":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Shape == nil {
				r.Shape = &String{}
			}
			r.Shape.Value = v.Value
		case "_shape":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Shape == nil {
				r.Shape = &String{}
			}
			r.Shape.Id = v.Id
			r.Shape.Extension = v.Extension
		case "color":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for i := 0; d.More(); i++ {
				var v String
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Color) <= i {
					r.Color = append(r.Color, String{})
				}
				r.Color[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "_color":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Color) <= i {
					r.Color = append(r.Color, String{})
				}
				r.Color[i].Id = v.Id
				r.Color[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "imprint":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for i := 0; d.More(); i++ {
				var v String
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Imprint) <= i {
					r.Imprint = append(r.Imprint, String{})
				}
				r.Imprint[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "_imprint":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Imprint) <= i {
					r.Imprint = append(r.Imprint, String{})
				}
				r.Imprint[i].Id = v.Id
				r.Imprint[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "image":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ProdCharacteristic element", t)
			}
			for d.More() {
				var v Attachment
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Image = append(r.Image, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ProdCharacteristic element", t)
			}
		case "scoring":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Scoring = &v
		default:
			return fmt.Errorf("invalid field: %s in ProdCharacteristic", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ProdCharacteristic element", t)
	}
	return nil
}
func (r ProdCharacteristic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Height, xml.StartElement{Name: xml.Name{Local: "height"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Width, xml.StartElement{Name: xml.Name{Local: "width"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Depth, xml.StartElement{Name: xml.Name{Local: "depth"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Weight, xml.StartElement{Name: xml.Name{Local: "weight"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NominalVolume, xml.StartElement{Name: xml.Name{Local: "nominalVolume"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExternalDiameter, xml.StartElement{Name: xml.Name{Local: "externalDiameter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Shape, xml.StartElement{Name: xml.Name{Local: "shape"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Color, xml.StartElement{Name: xml.Name{Local: "color"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Imprint, xml.StartElement{Name: xml.Name{Local: "imprint"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Image, xml.StartElement{Name: xml.Name{Local: "image"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Scoring, xml.StartElement{Name: xml.Name{Local: "scoring"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ProdCharacteristic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "height":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Height = &v
			case "width":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Width = &v
			case "depth":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Depth = &v
			case "weight":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Weight = &v
			case "nominalVolume":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NominalVolume = &v
			case "externalDiameter":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExternalDiameter = &v
			case "shape":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Shape = &v
			case "color":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Color = append(r.Color, v)
			case "imprint":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Imprint = append(r.Imprint, v)
			case "image":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Image = append(r.Image, v)
			case "scoring":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Scoring = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ProdCharacteristic) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "height") {
		if r.Height != nil {
			children = append(children, *r.Height)
		}
	}
	if len(name) == 0 || slices.Contains(name, "width") {
		if r.Width != nil {
			children = append(children, *r.Width)
		}
	}
	if len(name) == 0 || slices.Contains(name, "depth") {
		if r.Depth != nil {
			children = append(children, *r.Depth)
		}
	}
	if len(name) == 0 || slices.Contains(name, "weight") {
		if r.Weight != nil {
			children = append(children, *r.Weight)
		}
	}
	if len(name) == 0 || slices.Contains(name, "nominalVolume") {
		if r.NominalVolume != nil {
			children = append(children, *r.NominalVolume)
		}
	}
	if len(name) == 0 || slices.Contains(name, "externalDiameter") {
		if r.ExternalDiameter != nil {
			children = append(children, *r.ExternalDiameter)
		}
	}
	if len(name) == 0 || slices.Contains(name, "shape") {
		if r.Shape != nil {
			children = append(children, *r.Shape)
		}
	}
	if len(name) == 0 || slices.Contains(name, "color") {
		for _, v := range r.Color {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "imprint") {
		for _, v := range r.Imprint {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "image") {
		for _, v := range r.Image {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "scoring") {
		if r.Scoring != nil {
			children = append(children, *r.Scoring)
		}
	}
	return children
}
func (r ProdCharacteristic) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ProdCharacteristic to Boolean")
}
func (r ProdCharacteristic) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ProdCharacteristic to String")
}
func (r ProdCharacteristic) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ProdCharacteristic to Integer")
}
func (r ProdCharacteristic) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ProdCharacteristic to Decimal")
}
func (r ProdCharacteristic) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ProdCharacteristic to Date")
}
func (r ProdCharacteristic) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ProdCharacteristic to Time")
}
func (r ProdCharacteristic) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ProdCharacteristic to DateTime")
}
func (r ProdCharacteristic) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ProdCharacteristic to Quantity")
}
func (r ProdCharacteristic) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Height",
			Type: "FHIR.Quantity",
		}, {
			Name: "Width",
			Type: "FHIR.Quantity",
		}, {
			Name: "Depth",
			Type: "FHIR.Quantity",
		}, {
			Name: "Weight",
			Type: "FHIR.Quantity",
		}, {
			Name: "NominalVolume",
			Type: "FHIR.Quantity",
		}, {
			Name: "ExternalDiameter",
			Type: "FHIR.Quantity",
		}, {
			Name: "Shape",
			Type: "FHIR.String",
		}, {
			Name: "Color",
			Type: "List<FHIR.String>",
		}, {
			Name: "Imprint",
			Type: "List<FHIR.String>",
		}, {
			Name: "Image",
			Type: "List<FHIR.Attachment>",
		}, {
			Name: "Scoring",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ProdCharacteristic",
			Namespace: "FHIR",
		},
	}
}
