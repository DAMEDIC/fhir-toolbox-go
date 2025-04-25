package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"io"
	"reflect"
	"slices"
)

// ExtendedContactDetail Type: Specifies contact information for a specific purpose over a period of time, might be handled/monitored by a specific named person or organization.
//
// Need to track contact and address information in the same way across multiple resources.
type ExtendedContactDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The purpose/type of contact.
	Purpose *CodeableConcept
	// The name of an individual to contact, some types of contact detail are usually blank.
	Name []HumanName
	// The contact details application for the purpose defined.
	Telecom []ContactPoint
	// Address for the contact.
	Address *Address
	// This contact detail is handled/monitored by a specific organization. If the name is provided in the contact, then it is referring to the named individual within this organization.
	Organization *Reference
	// Period that this contact was valid for usage.
	Period *Period
}

func (r ExtendedContactDetail) MemSize() int {
	s := int(reflect.TypeOf(r).Size())
	if r.Id != nil {
		s += len(*r.Id) + int(reflect.TypeOf(*r.Id).Size())
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(reflect.TypeOf(Extension{}).Size())
	if r.Purpose != nil {
		s += r.Purpose.MemSize()
	}
	for _, i := range r.Name {
		s += i.MemSize()
	}
	s += (cap(r.Name) - len(r.Name)) * int(reflect.TypeOf(HumanName{}).Size())
	for _, i := range r.Telecom {
		s += i.MemSize()
	}
	s += (cap(r.Telecom) - len(r.Telecom)) * int(reflect.TypeOf(ContactPoint{}).Size())
	if r.Address != nil {
		s += r.Address.MemSize()
	}
	if r.Organization != nil {
		s += r.Organization.MemSize()
	}
	if r.Period != nil {
		s += r.Period.MemSize()
	}
	return s
}
func (r ExtendedContactDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ExtendedContactDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ExtendedContactDetail) marshalJSON(w io.Writer) error {
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
	if r.Purpose != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"purpose\":"))
		if err != nil {
			return err
		}
		err = r.Purpose.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Name) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Name {
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
	if len(r.Telecom) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"telecom\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Telecom {
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
	if r.Address != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"address\":"))
		if err != nil {
			return err
		}
		err = r.Address.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Organization != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"organization\":"))
		if err != nil {
			return err
		}
		err = r.Organization.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Period != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"period\":"))
		if err != nil {
			return err
		}
		err = r.Period.marshalJSON(w)
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
func (r *ExtendedContactDetail) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ExtendedContactDetail element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ExtendedContactDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ExtendedContactDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ExtendedContactDetail element", t)
			}
		case "purpose":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Purpose = &v
		case "name":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ExtendedContactDetail element", t)
			}
			for d.More() {
				var v HumanName
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Name = append(r.Name, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ExtendedContactDetail element", t)
			}
		case "telecom":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ExtendedContactDetail element", t)
			}
			for d.More() {
				var v ContactPoint
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ExtendedContactDetail element", t)
			}
		case "address":
			var v Address
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Address = &v
		case "organization":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Organization = &v
		case "period":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Period = &v
		default:
			return fmt.Errorf("invalid field: %s in ExtendedContactDetail", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ExtendedContactDetail element", t)
	}
	return nil
}
func (r ExtendedContactDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Purpose, xml.StartElement{Name: xml.Name{Local: "purpose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Telecom, xml.StartElement{Name: xml.Name{Local: "telecom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Organization, xml.StartElement{Name: xml.Name{Local: "organization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ExtendedContactDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "purpose":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Purpose = &v
			case "name":
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = append(r.Name, v)
			case "telecom":
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			case "address":
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = &v
			case "organization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Organization = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ExtendedContactDetail) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "purpose") {
		if r.Purpose != nil {
			children = append(children, *r.Purpose)
		}
	}
	if len(name) == 0 || slices.Contains(name, "name") {
		for _, v := range r.Name {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "telecom") {
		for _, v := range r.Telecom {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "address") {
		if r.Address != nil {
			children = append(children, *r.Address)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organization") {
		if r.Organization != nil {
			children = append(children, *r.Organization)
		}
	}
	if len(name) == 0 || slices.Contains(name, "period") {
		if r.Period != nil {
			children = append(children, *r.Period)
		}
	}
	return children
}
func (r ExtendedContactDetail) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ExtendedContactDetail to Boolean")
}
func (r ExtendedContactDetail) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ExtendedContactDetail to String")
}
func (r ExtendedContactDetail) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ExtendedContactDetail to Integer")
}
func (r ExtendedContactDetail) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ExtendedContactDetail to Decimal")
}
func (r ExtendedContactDetail) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ExtendedContactDetail to Date")
}
func (r ExtendedContactDetail) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ExtendedContactDetail to Time")
}
func (r ExtendedContactDetail) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ExtendedContactDetail to DateTime")
}
func (r ExtendedContactDetail) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ExtendedContactDetail to Quantity")
}
func (r ExtendedContactDetail) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ExtendedContactDetail
	switch other := other.(type) {
	case ExtendedContactDetail:
		o = &other
	case *ExtendedContactDetail:
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
func (r ExtendedContactDetail) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(ExtendedContactDetail)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r ExtendedContactDetail) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Purpose",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Name",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "HumanName",
				Namespace: "FHIR",
			},
		}, {
			Name: "Telecom",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactPoint",
				Namespace: "FHIR",
			},
		}, {
			Name: "Address",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Address",
				Namespace: "FHIR",
			},
		}, {
			Name: "Organization",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Period",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}},
		Name:      "ExtendedContactDetail",
		Namespace: "FHIR",
	}
}
