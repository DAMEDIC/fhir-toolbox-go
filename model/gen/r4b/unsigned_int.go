package r4b

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"slices"
	"strconv"
	"unsafe"
)

// Base StructureDefinition for unsignedInt type: An integer with a value that is not negative (e.g. >= 0)
type UnsignedInt struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Primitive value for unsignedInt
	Value *uint32
}

func (r UnsignedInt) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Value != nil {
		s += int(unsafe.Sizeof(*r.Value))
	}
	return s
}
func (r UnsignedInt) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r UnsignedInt) MarshalJSON() ([]byte, error) {
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
func (r *UnsignedInt) UnmarshalJSON(b []byte) error {
	var v uint32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*r = UnsignedInt{Value: &v}
	return nil
}
func (r UnsignedInt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	if r.Value != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: strconv.FormatInt(int64(*r.Value), 10),
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
func (r *UnsignedInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			i, err := strconv.ParseInt(a.Value, 10, 0)
			if err != nil {
				return err
			}
			v := uint32(i)
			r.Value = &v
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
func (r UnsignedInt) Children(name ...string) fhirpath.Collection {
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
func (r UnsignedInt) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert UnsignedInt to Boolean")
}
func (r UnsignedInt) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert UnsignedInt to String")
}
func (r UnsignedInt) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert UnsignedInt to Integer")
}
func (r UnsignedInt) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert UnsignedInt to Decimal")
}
func (r UnsignedInt) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert UnsignedInt to Date")
}
func (r UnsignedInt) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert UnsignedInt to Time")
}
func (r UnsignedInt) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert UnsignedInt to DateTime")
}
func (r UnsignedInt) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert UnsignedInt to Quantity")
}
func (r UnsignedInt) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o UnsignedInt
	switch other := other.(type) {
	case UnsignedInt:
		o = other
	case *UnsignedInt:
		o = *other
	default:
		return false
	}
	a := r.Value
	b := o.Value
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
func (r UnsignedInt) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o UnsignedInt
	switch other := other.(type) {
	case UnsignedInt:
		o = other
	case *UnsignedInt:
		o = *other
	default:
		return false
	}
	a := r.Value
	b := o.Value
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
func (r UnsignedInt) TypeInfo() fhirpath.TypeInfo {
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
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}
}
