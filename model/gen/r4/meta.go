package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

// Base StructureDefinition for Meta Type: The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
type Meta struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The version specific identifier, as it appears in the version portion of the URL. This value changes when the resource is created, updated, or deleted.
	VersionId *Id
	// When the resource last changed - e.g. when the version changed.
	LastUpdated *Instant
	// A uri that identifies the source system of the resource. This provides a minimal amount of [Provenance](provenance.html#) information that can be used to track or differentiate the source of information in the resource. The source may identify another FHIR server, document, message, database, etc.
	Source *Uri
	// A list of profiles (references to [StructureDefinition](structuredefinition.html#) resources) that this resource claims to conform to. The URL is a reference to [StructureDefinition.url](structuredefinition-definitions.html#StructureDefinition.url).
	Profile []Canonical
	// Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Security []Coding
	// Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	Tag []Coding
}

func (r Meta) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Meta) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Meta) marshalJSON(w io.Writer) error {
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
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
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
	if r.VersionId != nil && r.VersionId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"versionId\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.VersionId)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.VersionId != nil && (r.VersionId.Id != nil || r.VersionId.Extension != nil) {
		p := primitiveElement{Id: r.VersionId.Id, Extension: r.VersionId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_versionId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LastUpdated != nil && r.LastUpdated.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastUpdated\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastUpdated)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.LastUpdated != nil && (r.LastUpdated.Id != nil || r.LastUpdated.Extension != nil) {
		p := primitiveElement{Id: r.LastUpdated.Id, Extension: r.LastUpdated.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastUpdated\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Source != nil && r.Source.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"source\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Source)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Source != nil && (r.Source.Id != nil || r.Source.Extension != nil) {
		p := primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_source\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.Profile {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"profile\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Profile)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Profile {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_profile\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Profile {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if len(r.Security) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"security\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Security {
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
	if len(r.Tag) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"tag\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Tag {
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *Meta) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Meta element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Meta element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Meta element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Meta element", t)
			}
		case "versionId":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.VersionId == nil {
				r.VersionId = &Id{}
			}
			r.VersionId.Value = v.Value
		case "_versionId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.VersionId == nil {
				r.VersionId = &Id{}
			}
			r.VersionId.Id = v.Id
			r.VersionId.Extension = v.Extension
		case "lastUpdated":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastUpdated == nil {
				r.LastUpdated = &Instant{}
			}
			r.LastUpdated.Value = v.Value
		case "_lastUpdated":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastUpdated == nil {
				r.LastUpdated = &Instant{}
			}
			r.LastUpdated.Id = v.Id
			r.LastUpdated.Extension = v.Extension
		case "source":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Source == nil {
				r.Source = &Uri{}
			}
			r.Source.Value = v.Value
		case "_source":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Source == nil {
				r.Source = &Uri{}
			}
			r.Source.Id = v.Id
			r.Source.Extension = v.Extension
		case "profile":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Meta element", t)
			}
			for i := 0; d.More(); i++ {
				var v Canonical
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Profile) <= i {
					r.Profile = append(r.Profile, Canonical{})
				}
				r.Profile[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Meta element", t)
			}
		case "_profile":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Meta element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Profile) <= i {
					r.Profile = append(r.Profile, Canonical{})
				}
				r.Profile[i].Id = v.Id
				r.Profile[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Meta element", t)
			}
		case "security":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Meta element", t)
			}
			for d.More() {
				var v Coding
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Security = append(r.Security, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Meta element", t)
			}
		case "tag":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Meta element", t)
			}
			for d.More() {
				var v Coding
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Tag = append(r.Tag, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Meta element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in Meta", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Meta element", t)
	}
	return nil
}
func (r Meta) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.VersionId, xml.StartElement{Name: xml.Name{Local: "versionId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastUpdated, xml.StartElement{Name: xml.Name{Local: "lastUpdated"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Profile, xml.StartElement{Name: xml.Name{Local: "profile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Security, xml.StartElement{Name: xml.Name{Local: "security"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Tag, xml.StartElement{Name: xml.Name{Local: "tag"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Meta) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "versionId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VersionId = &v
			case "lastUpdated":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastUpdated = &v
			case "source":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = &v
			case "profile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Profile = append(r.Profile, v)
			case "security":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Security = append(r.Security, v)
			case "tag":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tag = append(r.Tag, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
