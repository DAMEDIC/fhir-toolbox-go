package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The individual responsible for making the annotation.
	Author isAnnotationAuthor
	// Indicates when this particular annotation was made.
	Time *DateTime
	// The text of the annotation in markdown format.
	Text Markdown
}
type isAnnotationAuthor interface {
	isAnnotationAuthor()
}

func (r Reference) isAnnotationAuthor() {}
func (r String) isAnnotationAuthor()    {}

type jsonAnnotation struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	AuthorReference              *Reference        `json:"authorReference,omitempty"`
	AuthorString                 *String           `json:"authorString,omitempty"`
	AuthorStringPrimitiveElement *primitiveElement `json:"_authorString,omitempty"`
	Time                         *DateTime         `json:"time,omitempty"`
	TimePrimitiveElement         *primitiveElement `json:"_time,omitempty"`
	Text                         Markdown          `json:"text,omitempty"`
	TextPrimitiveElement         *primitiveElement `json:"_text,omitempty"`
}

func (r Annotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Annotation) marshalJSON() jsonAnnotation {
	m := jsonAnnotation{}
	m.Id = r.Id
	m.Extension = r.Extension
	switch v := r.Author.(type) {
	case Reference:
		m.AuthorReference = &v
	case *Reference:
		m.AuthorReference = v
	case String:
		if v.Value != nil {
			m.AuthorString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AuthorStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.AuthorString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AuthorStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	if r.Time != nil && r.Time.Value != nil {
		m.Time = r.Time
	}
	if r.Time != nil && (r.Time.Id != nil || r.Time.Extension != nil) {
		m.TimePrimitiveElement = &primitiveElement{Id: r.Time.Id, Extension: r.Time.Extension}
	}
	if r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text.Id != nil || r.Text.Extension != nil {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	return m
}
func (r *Annotation) UnmarshalJSON(b []byte) error {
	var m jsonAnnotation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Annotation) unmarshalJSON(m jsonAnnotation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	if m.AuthorReference != nil {
		if r.Author != nil {
			return fmt.Errorf("multiple values for field \"Author\"")
		}
		v := m.AuthorReference
		r.Author = v
	}
	if m.AuthorString != nil || m.AuthorStringPrimitiveElement != nil {
		if r.Author != nil {
			return fmt.Errorf("multiple values for field \"Author\"")
		}
		v := m.AuthorString
		if m.AuthorStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AuthorStringPrimitiveElement.Id
			v.Extension = m.AuthorStringPrimitiveElement.Extension
		}
		r.Author = v
	}
	r.Time = m.Time
	if m.TimePrimitiveElement != nil {
		if r.Time == nil {
			r.Time = &DateTime{}
		}
		r.Time.Id = m.TimePrimitiveElement.Id
		r.Time.Extension = m.TimePrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	return nil
}
func (r Annotation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Author.(type) {
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "authorReference"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "authorString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Time, xml.StartElement{Name: xml.Name{Local: "time"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Annotation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "authorReference":
				if r.Author != nil {
					return fmt.Errorf("multiple values for field \"Author\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = v
			case "authorString":
				if r.Author != nil {
					return fmt.Errorf("multiple values for field \"Author\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = v
			case "time":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Time = &v
			case "text":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Annotation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
