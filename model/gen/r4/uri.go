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

// Base StructureDefinition for uri Type: String of characters used to identify a name or a resource
type Uri struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *string
}

func (r Uri) MemSize() int {
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
func (r Uri) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Uri) MarshalJSON() ([]byte, error) {
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
func (r *Uri) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*r = Uri{Value: &v}
	return nil
}
func (r Uri) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *Uri) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (r Uri) Children(name ...string) fhirpath.Collection {
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
func (r Uri) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Uri to Boolean")
}
func (r Uri) ToString(explicit bool) (*fhirpath.String, error) {
	if r.Value != nil {
		v := fhirpath.String(*r.Value)
		return &v, nil
	} else {
		return nil, nil
	}
}
func (r Uri) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Uri to Integer")
}
func (r Uri) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Uri to Decimal")
}
func (r Uri) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Uri to Date")
}
func (r Uri) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Uri to Time")
}
func (r Uri) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Uri to DateTime")
}
func (r Uri) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Uri to Quantity")
}
func (r Uri) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o Uri
	switch other := other.(type) {
	case Uri:
		o = other
	case *Uri:
		o = *other
	default:
		return false
	}
	a, err := r.ToString(false)
	if err != nil {
		return false
	}
	b, err := o.ToString(false)
	if err != nil {
		return false
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	if a != nil && b != nil && *a != *b {
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r Uri) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o Uri
	switch other := other.(type) {
	case Uri:
		o = other
	case *Uri:
		o = *other
	default:
		return false
	}
	a, err := r.ToString(false)
	if err != nil {
		return false
	}
	b, err := o.ToString(false)
	if err != nil {
		return false
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	if a != nil && b != nil && *a != *b {
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r Uri) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "PrimitiveType",
				Namespace: "FHIR",
			},
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}
}
