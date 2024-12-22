package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"unsafe"
)

// Base StructureDefinition for xhtml Type
type Xhtml struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// Actual xhtml
	Value string
}

func (r Xhtml) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	s += len(r.Value)
	return s
}
func (r Xhtml) MarshalJSON() ([]byte, error) {
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
func (r Xhtml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Space = "http://www.w3.org/1999/xhtml"
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	var v struct {
		V *string `xml:",innerxml"`
	}
	v.V = &r.Value
	err := e.EncodeElement(v, start)
	if err != nil {
		return err
	}
	return nil
}
