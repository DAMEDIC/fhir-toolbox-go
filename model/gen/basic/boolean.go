package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"slices"
	"strconv"
	"unsafe"
)

// Base StructureDefinition for boolean Type: Value of "true" or "false"
type Boolean struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *bool
}

func (r Boolean) MemSize() int {
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
func (r Boolean) MarshalJSON() ([]byte, error) {
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
func (r Boolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	if r.Value != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: strconv.FormatBool(*r.Value),
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
func (r Boolean) Children(name ...string) fhirpath.Collection {
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
func (r Boolean) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	if r.Value != nil {
		v := fhirpath.Boolean(*r.Value)
		return &v, nil
	} else {
		return nil, nil
	}
}
func (r Boolean) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Boolean to String")
}
func (r Boolean) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Boolean to Integer")
}
func (r Boolean) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Boolean to Decimal")
}
func (r Boolean) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Boolean to Date")
}
func (r Boolean) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Boolean to Time")
}
func (r Boolean) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Boolean to DateTime")
}
func (r Boolean) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Boolean to Quantity")
}
func (r Boolean) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Value",
			Type: "FHIR.bool",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "PrimitiveType",
				Namespace: "FHIR",
			},
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}
}
