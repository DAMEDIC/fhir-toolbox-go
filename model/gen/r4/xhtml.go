package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for xhtml Type
type Xhtml struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// Actual xhtml
	Value string
}

func (r Xhtml) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(r.Value)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r *Xhtml) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*r = Xhtml{Value: v}
	return nil
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
func (r *Xhtml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://www.w3.org/1999/xhtml" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://www.w3.org/1999/xhtml\"", start.Name.Space)
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
	var v struct {
		V string `xml:",innerxml"`
	}
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	r.Value = v.V
	return nil
}
