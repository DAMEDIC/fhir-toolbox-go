package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"slices"
	"unsafe"
)

// Base StructureDefinition for instant Type: An instant in time - known at least to the second
type Instant struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *string
}

func (r Instant) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Value != nil {
		s += len(*r.Value) + int(unsafe.Sizeof(*r.Value))
	}
	return s
}
func (r Instant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Instant) MarshalJSON() ([]byte, error) {
	v := r.Value
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r *Instant) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*r = Instant{Value: &v}
	return nil
}
func (r Instant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	if r.Value != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: *r.Value,
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
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Instant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "value":
			r.Value = &a.Value
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
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Instant) Children(name ...string) fhirpath.Collection {
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
	return children
}
func (r Instant) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Instant to Boolean")
}
func (r Instant) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Instant to String")
}
func (r Instant) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Instant to Integer")
}
func (r Instant) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Instant to Decimal")
}
func (r Instant) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Instant to Date")
}
func (r Instant) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Instant to Time")
}
func (r Instant) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	if r.Value != nil {
		v, err := fhirpath.ParseDateTime(*r.Value)
		return v, true, err
	} else {
		return fhirpath.DateTime{}, false, nil
	}
}
func (r Instant) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Instant to Quantity")
}
func (r Instant) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	a, ok, err := r.ToDateTime(false)
	if err != nil || !ok {
		return false, true
	}
	b, ok, err := other.ToDateTime(false)
	if err != nil || !ok {
		return false, true
	}
	return a.Equal(b)
}
func (r Instant) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r Instant) TypeInfo() fhirpath.TypeInfo {
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
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "PrimitiveType",
				Namespace: "FHIR",
			},
			Name:      "instant",
			Namespace: "FHIR",
		},
	}
}
