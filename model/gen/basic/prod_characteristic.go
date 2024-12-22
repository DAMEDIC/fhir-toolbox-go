package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
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
