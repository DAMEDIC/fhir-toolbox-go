package r5

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

// Signature Type: A signature along with supporting context. The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
//
// There are a number of places where content must be signed in healthcare.
type Signature struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	Type []Coding
	// When the digital signature was signed.
	When *Instant
	// A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	Who *Reference
	// A reference to an application-usable description of the identity that is represented by the signature.
	OnBehalfOf *Reference
	// A mime type that indicates the technical format of the target resources signed by the signature.
	TargetFormat *Code
	// A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	SigFormat *Code
	// The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Data *Base64Binary
}

func (r Signature) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Type {
		s += i.MemSize()
	}
	s += (cap(r.Type) - len(r.Type)) * int(unsafe.Sizeof(Coding{}))
	if r.When != nil {
		s += r.When.MemSize()
	}
	if r.Who != nil {
		s += r.Who.MemSize()
	}
	if r.OnBehalfOf != nil {
		s += r.OnBehalfOf.MemSize()
	}
	if r.TargetFormat != nil {
		s += r.TargetFormat.MemSize()
	}
	if r.SigFormat != nil {
		s += r.SigFormat.MemSize()
	}
	if r.Data != nil {
		s += r.Data.MemSize()
	}
	return s
}
func (r Signature) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Signature) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Signature) marshalJSON(w io.Writer) error {
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
	if len(r.Type) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Type {
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
	if r.When != nil && r.When.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"when\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.When)
		if err != nil {
			return err
		}
	}
	if r.When != nil && (r.When.Id != nil || r.When.Extension != nil) {
		p := primitiveElement{Id: r.When.Id, Extension: r.When.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_when\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Who != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"who\":"))
		if err != nil {
			return err
		}
		err = r.Who.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OnBehalfOf != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"onBehalfOf\":"))
		if err != nil {
			return err
		}
		err = r.OnBehalfOf.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.TargetFormat != nil && r.TargetFormat.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetFormat\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.TargetFormat)
		if err != nil {
			return err
		}
	}
	if r.TargetFormat != nil && (r.TargetFormat.Id != nil || r.TargetFormat.Extension != nil) {
		p := primitiveElement{Id: r.TargetFormat.Id, Extension: r.TargetFormat.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_targetFormat\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SigFormat != nil && r.SigFormat.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sigFormat\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SigFormat)
		if err != nil {
			return err
		}
	}
	if r.SigFormat != nil && (r.SigFormat.Id != nil || r.SigFormat.Extension != nil) {
		p := primitiveElement{Id: r.SigFormat.Id, Extension: r.SigFormat.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_sigFormat\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Data != nil && r.Data.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"data\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Data)
		if err != nil {
			return err
		}
	}
	if r.Data != nil && (r.Data.Id != nil || r.Data.Extension != nil) {
		p := primitiveElement{Id: r.Data.Id, Extension: r.Data.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_data\":"))
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
func (r *Signature) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Signature element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Signature element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Signature element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Signature element", t)
			}
		case "type":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Signature element", t)
			}
			for d.More() {
				var v Coding
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Signature element", t)
			}
		case "when":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.When == nil {
				r.When = &Instant{}
			}
			r.When.Value = v.Value
		case "_when":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.When == nil {
				r.When = &Instant{}
			}
			r.When.Id = v.Id
			r.When.Extension = v.Extension
		case "who":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Who = &v
		case "onBehalfOf":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.OnBehalfOf = &v
		case "targetFormat":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.TargetFormat == nil {
				r.TargetFormat = &Code{}
			}
			r.TargetFormat.Value = v.Value
		case "_targetFormat":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.TargetFormat == nil {
				r.TargetFormat = &Code{}
			}
			r.TargetFormat.Id = v.Id
			r.TargetFormat.Extension = v.Extension
		case "sigFormat":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SigFormat == nil {
				r.SigFormat = &Code{}
			}
			r.SigFormat.Value = v.Value
		case "_sigFormat":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SigFormat == nil {
				r.SigFormat = &Code{}
			}
			r.SigFormat.Id = v.Id
			r.SigFormat.Extension = v.Extension
		case "data":
			var v Base64Binary
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Data == nil {
				r.Data = &Base64Binary{}
			}
			r.Data.Value = v.Value
		case "_data":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Data == nil {
				r.Data = &Base64Binary{}
			}
			r.Data.Id = v.Id
			r.Data.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in Signature", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Signature element", t)
	}
	return nil
}
func (r Signature) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.When, xml.StartElement{Name: xml.Name{Local: "when"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Who, xml.StartElement{Name: xml.Name{Local: "who"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OnBehalfOf, xml.StartElement{Name: xml.Name{Local: "onBehalfOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetFormat, xml.StartElement{Name: xml.Name{Local: "targetFormat"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SigFormat, xml.StartElement{Name: xml.Name{Local: "sigFormat"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Data, xml.StartElement{Name: xml.Name{Local: "data"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Signature) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "when":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.When = &v
			case "who":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Who = &v
			case "onBehalfOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OnBehalfOf = &v
			case "targetFormat":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetFormat = &v
			case "sigFormat":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SigFormat = &v
			case "data":
				var v Base64Binary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Data = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Signature) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		for _, v := range r.Type {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "when") {
		if r.When != nil {
			children = append(children, *r.When)
		}
	}
	if len(name) == 0 || slices.Contains(name, "who") {
		if r.Who != nil {
			children = append(children, *r.Who)
		}
	}
	if len(name) == 0 || slices.Contains(name, "onBehalfOf") {
		if r.OnBehalfOf != nil {
			children = append(children, *r.OnBehalfOf)
		}
	}
	if len(name) == 0 || slices.Contains(name, "targetFormat") {
		if r.TargetFormat != nil {
			children = append(children, *r.TargetFormat)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sigFormat") {
		if r.SigFormat != nil {
			children = append(children, *r.SigFormat)
		}
	}
	if len(name) == 0 || slices.Contains(name, "data") {
		if r.Data != nil {
			children = append(children, *r.Data)
		}
	}
	return children
}
func (r Signature) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Signature to Boolean")
}
func (r Signature) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Signature to String")
}
func (r Signature) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Signature to Integer")
}
func (r Signature) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Signature to Decimal")
}
func (r Signature) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Signature to Date")
}
func (r Signature) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Signature to Time")
}
func (r Signature) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Signature to DateTime")
}
func (r Signature) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Signature to Quantity")
}
func (r Signature) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *Signature
	switch other := other.(type) {
	case Signature:
		o = &other
	case *Signature:
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
func (r Signature) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(Signature)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r Signature) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Coding",
				Namespace: "FHIR",
			},
		}, {
			Name: "When",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Instant",
				Namespace: "FHIR",
			},
		}, {
			Name: "Who",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "OnBehalfOf",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "TargetFormat",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "SigFormat",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Data",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Base64Binary",
				Namespace: "FHIR",
			},
		}},
		Name:      "Signature",
		Namespace: "FHIR",
	}
}
