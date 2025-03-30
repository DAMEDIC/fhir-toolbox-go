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

// Base StructureDefinition for xhtml Type
type Xhtml struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// Actual xhtml
	Value string
}

func (r Xhtml) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	s += len(r.Value)
	return s
}
func (r Xhtml) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Xhtml) MarshalJSON() ([]byte, error) {
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
func (r *Xhtml) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*r = Xhtml{Value: v}
	return nil
}
func (r Xhtml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Space = "http://www.w3.org/1999/xhtml"
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	var v struct {
		V *string `xml:",innerxml"`
	}
	v.V = &r.Value
	err := e.EncodeElement(v, start)
	if err != nil {
		return err
	}
	return nil
}
func (r *Xhtml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://www.w3.org/1999/xhtml" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://www.w3.org/1999/xhtml\"", start.Name.Space)
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
	var v struct {
		V string `xml:",innerxml"`
	}
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	r.Value = v.V
	return nil
}
func (r Xhtml) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	return children
}
func (r Xhtml) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Xhtml to Boolean")
}
func (r Xhtml) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Xhtml to String")
}
func (r Xhtml) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Xhtml to Integer")
}
func (r Xhtml) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Xhtml to Decimal")
}
func (r Xhtml) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Xhtml to Date")
}
func (r Xhtml) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Xhtml to Time")
}
func (r Xhtml) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Xhtml to DateTime")
}
func (r Xhtml) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Xhtml to Quantity")
}
func (r Xhtml) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	o, ok := other.(Xhtml)
	if !ok {
		return false, true
	}
	return r.Value == o.Value, true
}
func (r Xhtml) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r Xhtml) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}},
		Name:      "xhtml",
		Namespace: "FHIR",
	}
}
