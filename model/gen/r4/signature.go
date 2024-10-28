package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for Signature Type: A signature along with supporting context. The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
//
// There are a number of places where content must be signed in healthcare.
type Signature struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	Type []Coding
	// When the digital signature was signed.
	When Instant
	// A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	Who Reference
	// A reference to an application-usable description of the identity that is represented by the signature.
	OnBehalfOf *Reference
	// A mime type that indicates the technical format of the target resources signed by the signature.
	TargetFormat *Code
	// A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	SigFormat *Code
	// The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Data *Base64Binary
}
type jsonSignature struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	Type                         []Coding          `json:"type,omitempty"`
	When                         Instant           `json:"when,omitempty"`
	WhenPrimitiveElement         *primitiveElement `json:"_when,omitempty"`
	Who                          Reference         `json:"who,omitempty"`
	OnBehalfOf                   *Reference        `json:"onBehalfOf,omitempty"`
	TargetFormat                 *Code             `json:"targetFormat,omitempty"`
	TargetFormatPrimitiveElement *primitiveElement `json:"_targetFormat,omitempty"`
	SigFormat                    *Code             `json:"sigFormat,omitempty"`
	SigFormatPrimitiveElement    *primitiveElement `json:"_sigFormat,omitempty"`
	Data                         *Base64Binary     `json:"data,omitempty"`
	DataPrimitiveElement         *primitiveElement `json:"_data,omitempty"`
}

func (r Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Signature) marshalJSON() jsonSignature {
	m := jsonSignature{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Type = r.Type
	if r.When.Value != nil {
		m.When = r.When
	}
	if r.When.Id != nil || r.When.Extension != nil {
		m.WhenPrimitiveElement = &primitiveElement{Id: r.When.Id, Extension: r.When.Extension}
	}
	m.Who = r.Who
	m.OnBehalfOf = r.OnBehalfOf
	if r.TargetFormat != nil && r.TargetFormat.Value != nil {
		m.TargetFormat = r.TargetFormat
	}
	if r.TargetFormat != nil && (r.TargetFormat.Id != nil || r.TargetFormat.Extension != nil) {
		m.TargetFormatPrimitiveElement = &primitiveElement{Id: r.TargetFormat.Id, Extension: r.TargetFormat.Extension}
	}
	if r.SigFormat != nil && r.SigFormat.Value != nil {
		m.SigFormat = r.SigFormat
	}
	if r.SigFormat != nil && (r.SigFormat.Id != nil || r.SigFormat.Extension != nil) {
		m.SigFormatPrimitiveElement = &primitiveElement{Id: r.SigFormat.Id, Extension: r.SigFormat.Extension}
	}
	if r.Data != nil && r.Data.Value != nil {
		m.Data = r.Data
	}
	if r.Data != nil && (r.Data.Id != nil || r.Data.Extension != nil) {
		m.DataPrimitiveElement = &primitiveElement{Id: r.Data.Id, Extension: r.Data.Extension}
	}
	return m
}
func (r *Signature) UnmarshalJSON(b []byte) error {
	var m jsonSignature
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Signature) unmarshalJSON(m jsonSignature) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Type = m.Type
	r.When = m.When
	if m.WhenPrimitiveElement != nil {
		r.When.Id = m.WhenPrimitiveElement.Id
		r.When.Extension = m.WhenPrimitiveElement.Extension
	}
	r.Who = m.Who
	r.OnBehalfOf = m.OnBehalfOf
	r.TargetFormat = m.TargetFormat
	if m.TargetFormatPrimitiveElement != nil {
		if r.TargetFormat == nil {
			r.TargetFormat = &Code{}
		}
		r.TargetFormat.Id = m.TargetFormatPrimitiveElement.Id
		r.TargetFormat.Extension = m.TargetFormatPrimitiveElement.Extension
	}
	r.SigFormat = m.SigFormat
	if m.SigFormatPrimitiveElement != nil {
		if r.SigFormat == nil {
			r.SigFormat = &Code{}
		}
		r.SigFormat.Id = m.SigFormatPrimitiveElement.Id
		r.SigFormat.Extension = m.SigFormatPrimitiveElement.Extension
	}
	r.Data = m.Data
	if m.DataPrimitiveElement != nil {
		if r.Data == nil {
			r.Data = &Base64Binary{}
		}
		r.Data.Id = m.DataPrimitiveElement.Id
		r.Data.Extension = m.DataPrimitiveElement.Extension
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
				r.When = v
			case "who":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Who = v
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
func (r Signature) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
