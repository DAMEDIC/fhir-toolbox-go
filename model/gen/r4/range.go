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

// Base StructureDefinition for Range Type: A set of ordered Quantities defined by a low and high limit.
//
// Need to be able to specify ranges of values.
type Range struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The low limit. The boundary is inclusive.
	Low *Quantity
	// The high limit. The boundary is inclusive.
	High *Quantity
}

func (r Range) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Low != nil {
		s += r.Low.MemSize()
	}
	if r.High != nil {
		s += r.High.MemSize()
	}
	return s
}
func (r Range) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Range) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Range) marshalJSON(w io.Writer) error {
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
	if r.Low != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"low\":"))
		if err != nil {
			return err
		}
		err = r.Low.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.High != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"high\":"))
		if err != nil {
			return err
		}
		err = r.High.marshalJSON(w)
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
func (r *Range) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Range element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Range element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Range element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Range element", t)
			}
		case "low":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Low = &v
		case "high":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.High = &v
		default:
			return fmt.Errorf("invalid field: %s in Range", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Range element", t)
	}
	return nil
}
func (r Range) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Low, xml.StartElement{Name: xml.Name{Local: "low"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.High, xml.StartElement{Name: xml.Name{Local: "high"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Range) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "low":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Low = &v
			case "high":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.High = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Range) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "low") {
		if r.Low != nil {
			children = append(children, *r.Low)
		}
	}
	if len(name) == 0 || slices.Contains(name, "high") {
		if r.High != nil {
			children = append(children, *r.High)
		}
	}
	return children
}
func (r Range) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Range to Boolean")
}
func (r Range) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Range to String")
}
func (r Range) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Range to Integer")
}
func (r Range) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Range to Decimal")
}
func (r Range) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Range to Date")
}
func (r Range) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Range to Time")
}
func (r Range) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Range to DateTime")
}
func (r Range) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Range to Quantity")
}
func (r Range) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Low",
			Type: "FHIR.Quantity",
		}, {
			Name: "High",
			Type: "FHIR.Quantity",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "Range",
			Namespace: "FHIR",
		},
	}
}
