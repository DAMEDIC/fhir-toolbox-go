package r4

import (
	"encoding/json"
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
		m.AuthorString = &v
		if v.Id != nil || v.Extension != nil {
			m.AuthorStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.AuthorString = v
		if v.Id != nil || v.Extension != nil {
			m.AuthorStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Time = r.Time
	if r.Time != nil && (r.Time.Id != nil || r.Time.Extension != nil) {
		m.TimePrimitiveElement = &primitiveElement{Id: r.Time.Id, Extension: r.Time.Extension}
	}
	m.Text = r.Text
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
func (r Annotation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
