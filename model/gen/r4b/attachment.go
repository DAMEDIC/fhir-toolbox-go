package r4b

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"unsafe"
)

// Base StructureDefinition for Attachment Type: For referring to data content defined in other formats.
//
// Many models need to include data defined in other specifications that is complex and opaque to the healthcare model. This includes documents, media recordings, structured data, etc.
type Attachment struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	ContentType *Code
	// The human language of the content. The value can be any valid value according to BCP 47.
	Language *Code
	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *Base64Binary
	// A location where the data can be accessed.
	Url *Url
	// The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	Size *UnsignedInt
	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash *Base64Binary
	// A label or set of text to display in place of the data.
	Title *String
	// The date that the attachment was first created.
	Creation *DateTime
}

func (r Attachment) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.ContentType != nil {
		s += r.ContentType.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Data != nil {
		s += r.Data.MemSize()
	}
	if r.Url != nil {
		s += r.Url.MemSize()
	}
	if r.Size != nil {
		s += r.Size.MemSize()
	}
	if r.Hash != nil {
		s += r.Hash.MemSize()
	}
	if r.Title != nil {
		s += r.Title.MemSize()
	}
	if r.Creation != nil {
		s += r.Creation.MemSize()
	}
	return s
}
func (r Attachment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Attachment) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Attachment) marshalJSON(w io.Writer) error {
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
	if r.ContentType != nil && r.ContentType.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contentType\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ContentType)
		if err != nil {
			return err
		}
	}
	if r.ContentType != nil && (r.ContentType.Id != nil || r.ContentType.Extension != nil) {
		p := primitiveElement{Id: r.ContentType.Id, Extension: r.ContentType.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_contentType\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
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
	if r.Url != nil && r.Url.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"url\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Url)
		if err != nil {
			return err
		}
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		p := primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_url\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Size != nil && r.Size.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"size\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Size)
		if err != nil {
			return err
		}
	}
	if r.Size != nil && (r.Size.Id != nil || r.Size.Extension != nil) {
		p := primitiveElement{Id: r.Size.Id, Extension: r.Size.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_size\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Hash != nil && r.Hash.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"hash\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Hash)
		if err != nil {
			return err
		}
	}
	if r.Hash != nil && (r.Hash.Id != nil || r.Hash.Extension != nil) {
		p := primitiveElement{Id: r.Hash.Id, Extension: r.Hash.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_hash\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Title != nil && r.Title.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"title\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Title)
		if err != nil {
			return err
		}
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		p := primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_title\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Creation != nil && r.Creation.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"creation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Creation)
		if err != nil {
			return err
		}
	}
	if r.Creation != nil && (r.Creation.Id != nil || r.Creation.Extension != nil) {
		p := primitiveElement{Id: r.Creation.Id, Extension: r.Creation.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_creation\":"))
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
func (r *Attachment) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Attachment element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Attachment element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Attachment element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Attachment element", t)
			}
		case "contentType":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ContentType == nil {
				r.ContentType = &Code{}
			}
			r.ContentType.Value = v.Value
		case "_contentType":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ContentType == nil {
				r.ContentType = &Code{}
			}
			r.ContentType.Id = v.Id
			r.ContentType.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
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
		case "url":
			var v Url
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Url{}
			}
			r.Url.Value = v.Value
		case "_url":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Url{}
			}
			r.Url.Id = v.Id
			r.Url.Extension = v.Extension
		case "size":
			var v UnsignedInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Size == nil {
				r.Size = &UnsignedInt{}
			}
			r.Size.Value = v.Value
		case "_size":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Size == nil {
				r.Size = &UnsignedInt{}
			}
			r.Size.Id = v.Id
			r.Size.Extension = v.Extension
		case "hash":
			var v Base64Binary
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Hash == nil {
				r.Hash = &Base64Binary{}
			}
			r.Hash.Value = v.Value
		case "_hash":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Hash == nil {
				r.Hash = &Base64Binary{}
			}
			r.Hash.Id = v.Id
			r.Hash.Extension = v.Extension
		case "title":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Title == nil {
				r.Title = &String{}
			}
			r.Title.Value = v.Value
		case "_title":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Title == nil {
				r.Title = &String{}
			}
			r.Title.Id = v.Id
			r.Title.Extension = v.Extension
		case "creation":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Creation == nil {
				r.Creation = &DateTime{}
			}
			r.Creation.Value = v.Value
		case "_creation":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Creation == nil {
				r.Creation = &DateTime{}
			}
			r.Creation.Id = v.Id
			r.Creation.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in Attachment", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Attachment element", t)
	}
	return nil
}
func (r Attachment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ContentType, xml.StartElement{Name: xml.Name{Local: "contentType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Data, xml.StartElement{Name: xml.Name{Local: "data"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Size, xml.StartElement{Name: xml.Name{Local: "size"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Hash, xml.StartElement{Name: xml.Name{Local: "hash"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Creation, xml.StartElement{Name: xml.Name{Local: "creation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Attachment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "contentType":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContentType = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "data":
				var v Base64Binary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Data = &v
			case "url":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "size":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Size = &v
			case "hash":
				var v Base64Binary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Hash = &v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "creation":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Creation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
