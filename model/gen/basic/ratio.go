package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"io"
	"reflect"
	"slices"
)

// Base StructureDefinition for Ratio Type: A relationship of two Quantity values - expressed as a numerator and a denominator.
//
// Need to able to capture ratios for some measurements (titers) and some rates (costs).
type Ratio struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The value of the numerator.
	Numerator *Quantity
	// The value of the denominator.
	Denominator *Quantity
}

func (r Ratio) MemSize() int {
	s := int(reflect.TypeOf(r).Size())
	if r.Id != nil {
		s += len(*r.Id) + int(reflect.TypeOf(*r.Id).Size())
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(reflect.TypeOf(Extension{}).Size())
	if r.Numerator != nil {
		s += r.Numerator.MemSize()
	}
	if r.Denominator != nil {
		s += r.Denominator.MemSize()
	}
	return s
}
func (r Ratio) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Ratio) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Ratio) marshalJSON(w io.Writer) error {
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
	if r.Numerator != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"numerator\":"))
		if err != nil {
			return err
		}
		err = r.Numerator.marshalJSON(w)
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
func (r Ratio) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Numerator, xml.StartElement{Name: xml.Name{Local: "numerator"}})
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
func (r Ratio) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "numerator") {
		if r.Numerator != nil {
			children = append(children, *r.Numerator)
		}
	}
	if len(name) == 0 || slices.Contains(name, "denominator") {
		if r.Denominator != nil {
			children = append(children, *r.Denominator)
		}
	}
	return children
}
func (r Ratio) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Ratio to Boolean")
}
func (r Ratio) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Ratio to String")
}
func (r Ratio) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Ratio to Integer")
}
func (r Ratio) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Ratio to Decimal")
}
func (r Ratio) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Ratio to Date")
}
func (r Ratio) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Ratio to Time")
}
func (r Ratio) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Ratio to DateTime")
}
func (r Ratio) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Ratio to Quantity")
}
func (r Ratio) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *Ratio
	switch other := other.(type) {
	case Ratio:
		o = &other
	case *Ratio:
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
func (r Ratio) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(Ratio)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r Ratio) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
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
			Name: "Numerator",
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
		Name:      "Ratio",
		Namespace: "FHIR",
	}
}
