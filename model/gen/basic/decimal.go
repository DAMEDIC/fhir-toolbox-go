package basic

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/cockroachdb/apd/v3"
	"slices"
	"unsafe"
)

// Base StructureDefinition for decimal Type: A rational number with implicit precision
type Decimal struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *apd.Decimal
}

func (r Decimal) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Value != nil {
		s += int(r.Value.Size())
	}
	return s
}
func (r Decimal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Decimal) MarshalJSON() ([]byte, error) {
	if r.Value == nil {
		return []byte("null"), nil
	}
	return []byte(r.Value.Text('G')), nil
}
func (r Decimal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	if r.Value != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: r.Value.Text('G'),
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
func (r Decimal) Children(name ...string) fhirpath.Collection {
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
func (r Decimal) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Decimal to Boolean")
}
func (r Decimal) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Decimal to String")
}
func (r Decimal) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Decimal to Integer")
}
func (r Decimal) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	if r.Value != nil {
		v := fhirpath.Decimal{Value: r.Value}
		return &v, nil
	} else {
		return nil, nil
	}
}
func (r Decimal) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Decimal to Date")
}
func (r Decimal) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Decimal to Time")
}
func (r Decimal) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Decimal to DateTime")
}
func (r Decimal) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Decimal to Quantity")
}
func (r Decimal) TypeInfo() fhirpath.TypeInfo {
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
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}
}
