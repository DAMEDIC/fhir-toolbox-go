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

func (r Money) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Value != nil {
		s += r.Value.MemSize()
	}
	if r.Currency != nil {
		s += r.Currency.MemSize()
	}
	return s
}
func (r Money) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Money) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Money) marshalJSON(w io.Writer) error {
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
	if r.Value != nil && r.Value.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"value\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Value)
		if err != nil {
			return err
		}
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		p := primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_value\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Currency != nil && r.Currency.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"currency\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Currency)
		if err != nil {
			return err
		}
	}
	if r.Currency != nil && (r.Currency.Id != nil || r.Currency.Extension != nil) {
		p := primitiveElement{Id: r.Currency.Id, Extension: r.Currency.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_currency\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r *Money) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Money element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Money element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Money element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Money element", t)
			}
		case "value":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value == nil {
				r.Value = &Decimal{}
			}
			r.Value.Value = v.Value
		case "_value":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value == nil {
				r.Value = &Decimal{}
			}
			r.Value.Id = v.Id
			r.Value.Extension = v.Extension
		case "currency":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Currency == nil {
				r.Currency = &Code{}
			}
			r.Currency.Value = v.Value
		case "_currency":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Currency == nil {
				r.Currency = &Code{}
			}
			r.Currency.Id = v.Id
			r.Currency.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in Money", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Money element", t)
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
func (r Money) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "value") {
		if r.Value != nil {
			children = append(children, *r.Value)
		}
	}
	if len(name) == 0 || slices.Contains(name, "currency") {
		if r.Currency != nil {
			children = append(children, *r.Currency)
		}
	}
	return children
}
func (r Money) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Money to Boolean")
}
func (r Money) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Money to String")
}
func (r Money) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Money to Integer")
}
func (r Money) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Money to Decimal")
}
func (r Money) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Money to Date")
}
func (r Money) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Money to Time")
}
func (r Money) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Money to DateTime")
}
func (r Money) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Money to Quantity")
}
func (r Money) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *Money
	switch other := other.(type) {
	case Money:
		o = &other
	case *Money:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r Money) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(Money)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r Money) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Value",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Currency",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "Money",
			Namespace: "FHIR",
		},
	}
}
