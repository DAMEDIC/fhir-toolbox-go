package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"slices"
	"unsafe"
)

// Base StructureDefinition for base64Binary Type: A stream of bytes
type Base64Binary struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *string
}

func (r Base64Binary) MemSize() int {
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
func (r Base64Binary) MarshalJSON() ([]byte, error) {
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
func (r Base64Binary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r Base64Binary) Children(name ...string) fhirpath.Collection {
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
func (r Base64Binary) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Base64Binary to Boolean")
}
func (r Base64Binary) ToString(explicit bool) (*fhirpath.String, error) {
	if r.Value != nil {
		v := fhirpath.String(*r.Value)
		return &v, nil
	} else {
		return nil, nil
	}
}
func (r Base64Binary) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Base64Binary to Integer")
}
func (r Base64Binary) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Base64Binary to Decimal")
}
func (r Base64Binary) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Base64Binary to Date")
}
func (r Base64Binary) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Base64Binary to Time")
}
func (r Base64Binary) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Base64Binary to DateTime")
}
func (r Base64Binary) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Base64Binary to Quantity")
}
func (r Base64Binary) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Value",
			Type: "FHIR.string",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "PrimitiveType",
				Namespace: "FHIR",
			},
			Name:      "Base64Binary",
			Namespace: "FHIR",
		},
	}
}
