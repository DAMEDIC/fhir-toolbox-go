package r4b

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

func (r RelatedArtifact) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.Label != nil {
		s += r.Label.MemSize()
	}
	if r.Display != nil {
		s += r.Display.MemSize()
	}
	if r.Citation != nil {
		s += r.Citation.MemSize()
	}
	if r.Url != nil {
		s += r.Url.MemSize()
	}
	if r.Document != nil {
		s += r.Document.MemSize()
	}
	if r.Resource != nil {
		s += r.Resource.MemSize()
	}
	return s
}
func (r RelatedArtifact) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r RelatedArtifact) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r RelatedArtifact) marshalJSON(w io.Writer) error {
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
	{
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Type)
		if err != nil {
			return err
		}
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		p := primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_type\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Label != nil && r.Label.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"label\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Label)
		if err != nil {
			return err
		}
	}
	if r.Label != nil && (r.Label.Id != nil || r.Label.Extension != nil) {
		p := primitiveElement{Id: r.Label.Id, Extension: r.Label.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_label\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Display != nil && r.Display.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"display\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Display)
		if err != nil {
			return err
		}
	}
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		p := primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_display\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Citation != nil && r.Citation.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"citation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Citation)
		if err != nil {
			return err
		}
	}
	if r.Citation != nil && (r.Citation.Id != nil || r.Citation.Extension != nil) {
		p := primitiveElement{Id: r.Citation.Id, Extension: r.Citation.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_citation\":"))
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
	if r.Document != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"document\":"))
		if err != nil {
			return err
		}
		err = r.Document.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Resource != nil && r.Resource.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"resource\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Resource)
		if err != nil {
			return err
		}
	}
	if r.Resource != nil && (r.Resource.Id != nil || r.Resource.Extension != nil) {
		p := primitiveElement{Id: r.Resource.Id, Extension: r.Resource.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_resource\":"))
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
func (r *RelatedArtifact) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in RelatedArtifact element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in RelatedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in RelatedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in RelatedArtifact element", t)
			}
		case "type":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Type.Value = v.Value
		case "_type":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type.Id = v.Id
			r.Type.Extension = v.Extension
		case "label":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Label == nil {
				r.Label = &String{}
			}
			r.Label.Value = v.Value
		case "_label":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Label == nil {
				r.Label = &String{}
			}
			r.Label.Id = v.Id
			r.Label.Extension = v.Extension
		case "display":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Display == nil {
				r.Display = &String{}
			}
			r.Display.Value = v.Value
		case "_display":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Display == nil {
				r.Display = &String{}
			}
			r.Display.Id = v.Id
			r.Display.Extension = v.Extension
		case "citation":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Citation == nil {
				r.Citation = &Markdown{}
			}
			r.Citation.Value = v.Value
		case "_citation":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Citation == nil {
				r.Citation = &Markdown{}
			}
			r.Citation.Id = v.Id
			r.Citation.Extension = v.Extension
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
		case "document":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Document = &v
		case "resource":
			var v Canonical
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Resource == nil {
				r.Resource = &Canonical{}
			}
			r.Resource.Value = v.Value
		case "_resource":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Resource == nil {
				r.Resource = &Canonical{}
			}
			r.Resource.Id = v.Id
			r.Resource.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in RelatedArtifact", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in RelatedArtifact element", t)
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
func (r RelatedArtifact) Children(name ...string) fhirpath.Collection {
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
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "label") {
		if r.Label != nil {
			children = append(children, *r.Label)
		}
	}
	if len(name) == 0 || slices.Contains(name, "display") {
		if r.Display != nil {
			children = append(children, *r.Display)
		}
	}
	if len(name) == 0 || slices.Contains(name, "citation") {
		if r.Citation != nil {
			children = append(children, *r.Citation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "url") {
		if r.Url != nil {
			children = append(children, *r.Url)
		}
	}
	if len(name) == 0 || slices.Contains(name, "document") {
		if r.Document != nil {
			children = append(children, *r.Document)
		}
	}
	if len(name) == 0 || slices.Contains(name, "resource") {
		if r.Resource != nil {
			children = append(children, *r.Resource)
		}
	}
	return children
}
func (r RelatedArtifact) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert RelatedArtifact to Boolean")
}
func (r RelatedArtifact) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert RelatedArtifact to String")
}
func (r RelatedArtifact) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert RelatedArtifact to Integer")
}
func (r RelatedArtifact) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert RelatedArtifact to Decimal")
}
func (r RelatedArtifact) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert RelatedArtifact to Date")
}
func (r RelatedArtifact) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert RelatedArtifact to Time")
}
func (r RelatedArtifact) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert RelatedArtifact to DateTime")
}
func (r RelatedArtifact) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert RelatedArtifact to Quantity")
}
func (r RelatedArtifact) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *RelatedArtifact
	switch other := other.(type) {
	case RelatedArtifact:
		o = &other
	case *RelatedArtifact:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RelatedArtifact) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *RelatedArtifact
	switch other := other.(type) {
	case RelatedArtifact:
		o = &other
	case *RelatedArtifact:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RelatedArtifact) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
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
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Label",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Display",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Citation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}, {
			Name: "Url",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Url",
				Namespace: "FHIR",
			},
		}, {
			Name: "Document",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Attachment",
				Namespace: "FHIR",
			},
		}, {
			Name: "Resource",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Canonical",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}
}
