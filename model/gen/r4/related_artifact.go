package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for RelatedArtifact Type: Related artifacts such as additional documentation, justification, or bibliographic references.
//
// Knowledge resources must be able to provide enough information for consumers of the content (and/or interventions or results produced by the content) to be able to determine and understand the justification for and evidence in support of the content.
type RelatedArtifact struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The type of relationship to the related artifact.
	Type Code
	// A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
	Label *String
	// A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Display *String
	// A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
	Citation *Markdown
	// A url for the artifact that can be followed to access the actual content.
	Url *Url
	// The document being referenced, represented as an attachment. This is exclusive with the resource element.
	Document *Attachment
	// The related resource, such as a library, value set, profile, or other knowledge resource.
	Resource *Canonical
}
type jsonRelatedArtifact struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	Type                     Code              `json:"type,omitempty"`
	TypePrimitiveElement     *primitiveElement `json:"_type,omitempty"`
	Label                    *String           `json:"label,omitempty"`
	LabelPrimitiveElement    *primitiveElement `json:"_label,omitempty"`
	Display                  *String           `json:"display,omitempty"`
	DisplayPrimitiveElement  *primitiveElement `json:"_display,omitempty"`
	Citation                 *Markdown         `json:"citation,omitempty"`
	CitationPrimitiveElement *primitiveElement `json:"_citation,omitempty"`
	Url                      *Url              `json:"url,omitempty"`
	UrlPrimitiveElement      *primitiveElement `json:"_url,omitempty"`
	Document                 *Attachment       `json:"document,omitempty"`
	Resource                 *Canonical        `json:"resource,omitempty"`
	ResourcePrimitiveElement *primitiveElement `json:"_resource,omitempty"`
}

func (r RelatedArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RelatedArtifact) marshalJSON() jsonRelatedArtifact {
	m := jsonRelatedArtifact{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Label != nil && r.Label.Value != nil {
		m.Label = r.Label
	}
	if r.Label != nil && (r.Label.Id != nil || r.Label.Extension != nil) {
		m.LabelPrimitiveElement = &primitiveElement{Id: r.Label.Id, Extension: r.Label.Extension}
	}
	if r.Display != nil && r.Display.Value != nil {
		m.Display = r.Display
	}
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	if r.Citation != nil && r.Citation.Value != nil {
		m.Citation = r.Citation
	}
	if r.Citation != nil && (r.Citation.Id != nil || r.Citation.Extension != nil) {
		m.CitationPrimitiveElement = &primitiveElement{Id: r.Citation.Id, Extension: r.Citation.Extension}
	}
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Document = r.Document
	if r.Resource != nil && r.Resource.Value != nil {
		m.Resource = r.Resource
	}
	if r.Resource != nil && (r.Resource.Id != nil || r.Resource.Extension != nil) {
		m.ResourcePrimitiveElement = &primitiveElement{Id: r.Resource.Id, Extension: r.Resource.Extension}
	}
	return m
}
func (r *RelatedArtifact) UnmarshalJSON(b []byte) error {
	var m jsonRelatedArtifact
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RelatedArtifact) unmarshalJSON(m jsonRelatedArtifact) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Label = m.Label
	if m.LabelPrimitiveElement != nil {
		if r.Label == nil {
			r.Label = &String{}
		}
		r.Label.Id = m.LabelPrimitiveElement.Id
		r.Label.Extension = m.LabelPrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		if r.Display == nil {
			r.Display = &String{}
		}
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	r.Citation = m.Citation
	if m.CitationPrimitiveElement != nil {
		if r.Citation == nil {
			r.Citation = &Markdown{}
		}
		r.Citation.Id = m.CitationPrimitiveElement.Id
		r.Citation.Extension = m.CitationPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Url{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Document = m.Document
	r.Resource = m.Resource
	if m.ResourcePrimitiveElement != nil {
		if r.Resource == nil {
			r.Resource = &Canonical{}
		}
		r.Resource.Id = m.ResourcePrimitiveElement.Id
		r.Resource.Extension = m.ResourcePrimitiveElement.Extension
	}
	return nil
}
func (r RelatedArtifact) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Label, xml.StartElement{Name: xml.Name{Local: "label"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Display, xml.StartElement{Name: xml.Name{Local: "display"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Citation, xml.StartElement{Name: xml.Name{Local: "citation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Document, xml.StartElement{Name: xml.Name{Local: "document"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *RelatedArtifact) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "label":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Label = &v
			case "display":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Display = &v
			case "citation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Citation = &v
			case "url":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "document":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Document = &v
			case "resource":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r RelatedArtifact) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
