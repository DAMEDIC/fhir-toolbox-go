package r4b

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

// Base StructureDefinition for RatioRange Type: A range of ratios expressed as a low and high numerator and a denominator.
//
// Need to be able to specify ranges of ratios.
type RatioRange struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The value of the low limit numerator.
	LowNumerator *Quantity
	// The value of the high limit numerator.
	HighNumerator *Quantity
	// The value of the denominator.
	Denominator *Quantity
}

func (r RatioRange) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.LowNumerator != nil {
		s += r.LowNumerator.MemSize()
	}
	if r.HighNumerator != nil {
		s += r.HighNumerator.MemSize()
	}
	if r.Denominator != nil {
		s += r.Denominator.MemSize()
	}
	return s
}
func (r RatioRange) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r RatioRange) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r RatioRange) marshalJSON(w io.Writer) error {
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
	if r.LowNumerator != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lowNumerator\":"))
		if err != nil {
			return err
		}
		err = r.LowNumerator.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.HighNumerator != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"highNumerator\":"))
		if err != nil {
			return err
		}
		err = r.HighNumerator.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Denominator != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"denominator\":"))
		if err != nil {
			return err
		}
		err = r.Denominator.marshalJSON(w)
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
func (r *RatioRange) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in RatioRange element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in RatioRange element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in RatioRange element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in RatioRange element", t)
			}
		case "lowNumerator":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.LowNumerator = &v
		case "highNumerator":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.HighNumerator = &v
		case "denominator":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Denominator = &v
		default:
			return fmt.Errorf("invalid field: %s in RatioRange", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in RatioRange element", t)
	}
	return nil
}
func (r RatioRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.LowNumerator, xml.StartElement{Name: xml.Name{Local: "lowNumerator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HighNumerator, xml.StartElement{Name: xml.Name{Local: "highNumerator"}})
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
func (r *RatioRange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "lowNumerator":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LowNumerator = &v
			case "highNumerator":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HighNumerator = &v
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
func (r RatioRange) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "lowNumerator") {
		if r.LowNumerator != nil {
			children = append(children, *r.LowNumerator)
		}
	}
	if len(name) == 0 || slices.Contains(name, "highNumerator") {
		if r.HighNumerator != nil {
			children = append(children, *r.HighNumerator)
		}
	}
	if len(name) == 0 || slices.Contains(name, "denominator") {
		if r.Denominator != nil {
			children = append(children, *r.Denominator)
		}
	}
	return children
}
func (r RatioRange) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert RatioRange to Boolean")
}
func (r RatioRange) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert RatioRange to String")
}
func (r RatioRange) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert RatioRange to Integer")
}
func (r RatioRange) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert RatioRange to Decimal")
}
func (r RatioRange) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert RatioRange to Date")
}
func (r RatioRange) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert RatioRange to Time")
}
func (r RatioRange) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert RatioRange to DateTime")
}
func (r RatioRange) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert RatioRange to Quantity")
}
func (r RatioRange) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *RatioRange
	switch other := other.(type) {
	case RatioRange:
		o = &other
	case *RatioRange:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RatioRange) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *RatioRange
	switch other := other.(type) {
	case RatioRange:
		o = &other
	case *RatioRange:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RatioRange) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "LowNumerator",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "HighNumerator",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "Denominator",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "RatioRange",
			Namespace: "FHIR",
		},
	}
}
