package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A container for a collection of resources.
type Bundle struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A persistent identifier for the bundle that won't change as a bundle is copied from server to server.
	Identifier *Identifier
	// Indicates the purpose of this bundle - how it is intended to be used.
	Type Code
	// The date/time that the bundle was assembled - i.e. when the resources were placed in the bundle.
	Timestamp *Instant
	// If a set of search matches, this is the total number of entries of type 'match' across all pages in the search.  It does not include search.mode = 'include' or 'outcome' entries and it does not provide a count of the number of entries in the Bundle.
	Total *UnsignedInt
	// A series of links that provide context to this bundle.
	Link []BundleLink
	// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
	Entry []BundleEntry
	// Digital Signature - base64 encoded. XML-DSig or a JWT.
	Signature *Signature
}

func (r Bundle) ResourceType() string {
	return "Bundle"
}
func (r Bundle) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonBundle struct {
	ResourceType                  string            `json:"resourceType"`
	Id                            *Id               `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement `json:"_id,omitempty"`
	Meta                          *Meta             `json:"meta,omitempty"`
	ImplicitRules                 *Uri              `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement `json:"_implicitRules,omitempty"`
	Language                      *Code             `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement `json:"_language,omitempty"`
	Identifier                    *Identifier       `json:"identifier,omitempty"`
	Type                          Code              `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement `json:"_type,omitempty"`
	Timestamp                     *Instant          `json:"timestamp,omitempty"`
	TimestampPrimitiveElement     *primitiveElement `json:"_timestamp,omitempty"`
	Total                         *UnsignedInt      `json:"total,omitempty"`
	TotalPrimitiveElement         *primitiveElement `json:"_total,omitempty"`
	Link                          []BundleLink      `json:"link,omitempty"`
	Entry                         []BundleEntry     `json:"entry,omitempty"`
	Signature                     *Signature        `json:"signature,omitempty"`
}

func (r Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Bundle) marshalJSON() jsonBundle {
	m := jsonBundle{}
	m.ResourceType = "Bundle"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Identifier = r.Identifier
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Timestamp != nil && r.Timestamp.Value != nil {
		m.Timestamp = r.Timestamp
	}
	if r.Timestamp != nil && (r.Timestamp.Id != nil || r.Timestamp.Extension != nil) {
		m.TimestampPrimitiveElement = &primitiveElement{Id: r.Timestamp.Id, Extension: r.Timestamp.Extension}
	}
	if r.Total != nil && r.Total.Value != nil {
		m.Total = r.Total
	}
	if r.Total != nil && (r.Total.Id != nil || r.Total.Extension != nil) {
		m.TotalPrimitiveElement = &primitiveElement{Id: r.Total.Id, Extension: r.Total.Extension}
	}
	m.Link = r.Link
	m.Entry = r.Entry
	m.Signature = r.Signature
	return m
}
func (r *Bundle) UnmarshalJSON(b []byte) error {
	var m jsonBundle
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Bundle) unmarshalJSON(m jsonBundle) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Timestamp = m.Timestamp
	if m.TimestampPrimitiveElement != nil {
		if r.Timestamp == nil {
			r.Timestamp = &Instant{}
		}
		r.Timestamp.Id = m.TimestampPrimitiveElement.Id
		r.Timestamp.Extension = m.TimestampPrimitiveElement.Extension
	}
	r.Total = m.Total
	if m.TotalPrimitiveElement != nil {
		if r.Total == nil {
			r.Total = &UnsignedInt{}
		}
		r.Total.Id = m.TotalPrimitiveElement.Id
		r.Total.Extension = m.TotalPrimitiveElement.Extension
	}
	r.Link = m.Link
	r.Entry = m.Entry
	r.Signature = m.Signature
	return nil
}
func (r Bundle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Bundle"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Timestamp, xml.StartElement{Name: xml.Name{Local: "timestamp"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Total, xml.StartElement{Name: xml.Name{Local: "total"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Entry, xml.StartElement{Name: xml.Name{Local: "entry"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Signature, xml.StartElement{Name: xml.Name{Local: "signature"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Bundle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "timestamp":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timestamp = &v
			case "total":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Total = &v
			case "link":
				var v BundleLink
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Link = append(r.Link, v)
			case "entry":
				var v BundleEntry
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Entry = append(r.Entry, v)
			case "signature":
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Signature = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Bundle) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A series of links that provide context to this bundle.
type BundleLink struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A name which details the functional use for this link - see [http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1](http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1).
	Relation String
	// The reference details for the link.
	Url Uri
}
type jsonBundleLink struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Relation                 String            `json:"relation,omitempty"`
	RelationPrimitiveElement *primitiveElement `json:"_relation,omitempty"`
	Url                      Uri               `json:"url,omitempty"`
	UrlPrimitiveElement      *primitiveElement `json:"_url,omitempty"`
}

func (r BundleLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BundleLink) marshalJSON() jsonBundleLink {
	m := jsonBundleLink{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Relation.Value != nil {
		m.Relation = r.Relation
	}
	if r.Relation.Id != nil || r.Relation.Extension != nil {
		m.RelationPrimitiveElement = &primitiveElement{Id: r.Relation.Id, Extension: r.Relation.Extension}
	}
	if r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	return m
}
func (r *BundleLink) UnmarshalJSON(b []byte) error {
	var m jsonBundleLink
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BundleLink) unmarshalJSON(m jsonBundleLink) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Relation = m.Relation
	if m.RelationPrimitiveElement != nil {
		r.Relation.Id = m.RelationPrimitiveElement.Id
		r.Relation.Extension = m.RelationPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	return nil
}
func (r BundleLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Relation, xml.StartElement{Name: xml.Name{Local: "relation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BundleLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "relation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relation = v
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BundleLink) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
type BundleEntry struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A series of links that provide context to this entry.
	Link []BundleLink
	// The Absolute URL for the resource.  The fullUrl SHALL NOT disagree with the id in the resource - i.e. if the fullUrl is not a urn:uuid, the URL shall be version-independent URL consistent with the Resource.id. The fullUrl is a version independent reference to the resource. The fullUrl element SHALL have a value except that:
	// * fullUrl can be empty on a POST (although it does not need to when specifying a temporary id for reference in the bundle)
	// * Results from operations might involve resources that are not identified.
	FullUrl *Uri
	// The Resource for the entry. The purpose/meaning of the resource is determined by the Bundle.type.
	Resource *model.Resource
	// Information about the search process that lead to the creation of this entry.
	Search *BundleEntrySearch
	// Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
	Request *BundleEntryRequest
	// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
	Response *BundleEntryResponse
}
type jsonBundleEntry struct {
	Id                      *string              `json:"id,omitempty"`
	Extension               []Extension          `json:"extension,omitempty"`
	ModifierExtension       []Extension          `json:"modifierExtension,omitempty"`
	Link                    []BundleLink         `json:"link,omitempty"`
	FullUrl                 *Uri                 `json:"fullUrl,omitempty"`
	FullUrlPrimitiveElement *primitiveElement    `json:"_fullUrl,omitempty"`
	Resource                *ContainedResource   `json:"resource,omitempty"`
	Search                  *BundleEntrySearch   `json:"search,omitempty"`
	Request                 *BundleEntryRequest  `json:"request,omitempty"`
	Response                *BundleEntryResponse `json:"response,omitempty"`
}

func (r BundleEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BundleEntry) marshalJSON() jsonBundleEntry {
	m := jsonBundleEntry{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Link = r.Link
	if r.FullUrl != nil && r.FullUrl.Value != nil {
		m.FullUrl = r.FullUrl
	}
	if r.FullUrl != nil && (r.FullUrl.Id != nil || r.FullUrl.Extension != nil) {
		m.FullUrlPrimitiveElement = &primitiveElement{Id: r.FullUrl.Id, Extension: r.FullUrl.Extension}
	}
	if r.Resource != nil {
		m.Resource = &ContainedResource{*r.Resource}
	}
	m.Search = r.Search
	m.Request = r.Request
	m.Response = r.Response
	return m
}
func (r *BundleEntry) UnmarshalJSON(b []byte) error {
	var m jsonBundleEntry
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BundleEntry) unmarshalJSON(m jsonBundleEntry) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Link = m.Link
	r.FullUrl = m.FullUrl
	if m.FullUrlPrimitiveElement != nil {
		if r.FullUrl == nil {
			r.FullUrl = &Uri{}
		}
		r.FullUrl.Id = m.FullUrlPrimitiveElement.Id
		r.FullUrl.Extension = m.FullUrlPrimitiveElement.Extension
	}
	if m.Resource != nil {
		r.Resource = &m.Resource.Resource
	}
	r.Search = m.Search
	r.Request = m.Request
	r.Response = m.Response
	return nil
}
func (r BundleEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FullUrl, xml.StartElement{Name: xml.Name{Local: "fullUrl"}})
	if err != nil {
		return err
	}
	if r.Resource != nil {
		v := ContainedResource{*r.Resource}
		err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "resource"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Search, xml.StartElement{Name: xml.Name{Local: "search"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Response, xml.StartElement{Name: xml.Name{Local: "response"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BundleEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "link":
				var v BundleLink
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Link = append(r.Link, v)
			case "fullUrl":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FullUrl = &v
			case "resource":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Resource = &c.Resource
			case "search":
				var v BundleEntrySearch
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Search = &v
			case "request":
				var v BundleEntryRequest
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = &v
			case "response":
				var v BundleEntryResponse
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Response = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BundleEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information about the search process that lead to the creation of this entry.
type BundleEntrySearch struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Why this entry is in the result set - whether it's included as a match or because of an _include requirement, or to convey information or warning information about the search process.
	Mode *Code
	// When searching, the server's search ranking score for the entry.
	Score *Decimal
}
type jsonBundleEntrySearch struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Mode                  *Code             `json:"mode,omitempty"`
	ModePrimitiveElement  *primitiveElement `json:"_mode,omitempty"`
	Score                 *Decimal          `json:"score,omitempty"`
	ScorePrimitiveElement *primitiveElement `json:"_score,omitempty"`
}

func (r BundleEntrySearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BundleEntrySearch) marshalJSON() jsonBundleEntrySearch {
	m := jsonBundleEntrySearch{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Mode != nil && r.Mode.Value != nil {
		m.Mode = r.Mode
	}
	if r.Mode != nil && (r.Mode.Id != nil || r.Mode.Extension != nil) {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	if r.Score != nil && r.Score.Value != nil {
		m.Score = r.Score
	}
	if r.Score != nil && (r.Score.Id != nil || r.Score.Extension != nil) {
		m.ScorePrimitiveElement = &primitiveElement{Id: r.Score.Id, Extension: r.Score.Extension}
	}
	return m
}
func (r *BundleEntrySearch) UnmarshalJSON(b []byte) error {
	var m jsonBundleEntrySearch
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BundleEntrySearch) unmarshalJSON(m jsonBundleEntrySearch) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		if r.Mode == nil {
			r.Mode = &Code{}
		}
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Score = m.Score
	if m.ScorePrimitiveElement != nil {
		if r.Score == nil {
			r.Score = &Decimal{}
		}
		r.Score.Id = m.ScorePrimitiveElement.Id
		r.Score.Extension = m.ScorePrimitiveElement.Extension
	}
	return nil
}
func (r BundleEntrySearch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Mode, xml.StartElement{Name: xml.Name{Local: "mode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Score, xml.StartElement{Name: xml.Name{Local: "score"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BundleEntrySearch) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "mode":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Mode = &v
			case "score":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Score = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BundleEntrySearch) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
type BundleEntryRequest struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// In a transaction or batch, this is the HTTP action to be executed for this entry. In a history bundle, this indicates the HTTP action that occurred.
	Method Code
	// The URL for this entry, relative to the root (the address to which the request is posted).
	Url Uri
	// If the ETag values match, return a 304 Not Modified status. See the API documentation for ["Conditional Read"](http.html#cread).
	IfNoneMatch *String
	// Only perform the operation if the last updated date matches. See the API documentation for ["Conditional Read"](http.html#cread).
	IfModifiedSince *Instant
	// Only perform the operation if the Etag value matches. For more information, see the API section ["Managing Resource Contention"](http.html#concurrency).
	IfMatch *String
	// Instruct the server not to perform the create if a specified resource already exists. For further information, see the API documentation for ["Conditional Create"](http.html#ccreate). This is just the query portion of the URL - what follows the "?" (not including the "?").
	IfNoneExist *String
}
type jsonBundleEntryRequest struct {
	Id                              *string           `json:"id,omitempty"`
	Extension                       []Extension       `json:"extension,omitempty"`
	ModifierExtension               []Extension       `json:"modifierExtension,omitempty"`
	Method                          Code              `json:"method,omitempty"`
	MethodPrimitiveElement          *primitiveElement `json:"_method,omitempty"`
	Url                             Uri               `json:"url,omitempty"`
	UrlPrimitiveElement             *primitiveElement `json:"_url,omitempty"`
	IfNoneMatch                     *String           `json:"ifNoneMatch,omitempty"`
	IfNoneMatchPrimitiveElement     *primitiveElement `json:"_ifNoneMatch,omitempty"`
	IfModifiedSince                 *Instant          `json:"ifModifiedSince,omitempty"`
	IfModifiedSincePrimitiveElement *primitiveElement `json:"_ifModifiedSince,omitempty"`
	IfMatch                         *String           `json:"ifMatch,omitempty"`
	IfMatchPrimitiveElement         *primitiveElement `json:"_ifMatch,omitempty"`
	IfNoneExist                     *String           `json:"ifNoneExist,omitempty"`
	IfNoneExistPrimitiveElement     *primitiveElement `json:"_ifNoneExist,omitempty"`
}

func (r BundleEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BundleEntryRequest) marshalJSON() jsonBundleEntryRequest {
	m := jsonBundleEntryRequest{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Method.Value != nil {
		m.Method = r.Method
	}
	if r.Method.Id != nil || r.Method.Extension != nil {
		m.MethodPrimitiveElement = &primitiveElement{Id: r.Method.Id, Extension: r.Method.Extension}
	}
	if r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.IfNoneMatch != nil && r.IfNoneMatch.Value != nil {
		m.IfNoneMatch = r.IfNoneMatch
	}
	if r.IfNoneMatch != nil && (r.IfNoneMatch.Id != nil || r.IfNoneMatch.Extension != nil) {
		m.IfNoneMatchPrimitiveElement = &primitiveElement{Id: r.IfNoneMatch.Id, Extension: r.IfNoneMatch.Extension}
	}
	if r.IfModifiedSince != nil && r.IfModifiedSince.Value != nil {
		m.IfModifiedSince = r.IfModifiedSince
	}
	if r.IfModifiedSince != nil && (r.IfModifiedSince.Id != nil || r.IfModifiedSince.Extension != nil) {
		m.IfModifiedSincePrimitiveElement = &primitiveElement{Id: r.IfModifiedSince.Id, Extension: r.IfModifiedSince.Extension}
	}
	if r.IfMatch != nil && r.IfMatch.Value != nil {
		m.IfMatch = r.IfMatch
	}
	if r.IfMatch != nil && (r.IfMatch.Id != nil || r.IfMatch.Extension != nil) {
		m.IfMatchPrimitiveElement = &primitiveElement{Id: r.IfMatch.Id, Extension: r.IfMatch.Extension}
	}
	if r.IfNoneExist != nil && r.IfNoneExist.Value != nil {
		m.IfNoneExist = r.IfNoneExist
	}
	if r.IfNoneExist != nil && (r.IfNoneExist.Id != nil || r.IfNoneExist.Extension != nil) {
		m.IfNoneExistPrimitiveElement = &primitiveElement{Id: r.IfNoneExist.Id, Extension: r.IfNoneExist.Extension}
	}
	return m
}
func (r *BundleEntryRequest) UnmarshalJSON(b []byte) error {
	var m jsonBundleEntryRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BundleEntryRequest) unmarshalJSON(m jsonBundleEntryRequest) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Method = m.Method
	if m.MethodPrimitiveElement != nil {
		r.Method.Id = m.MethodPrimitiveElement.Id
		r.Method.Extension = m.MethodPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.IfNoneMatch = m.IfNoneMatch
	if m.IfNoneMatchPrimitiveElement != nil {
		if r.IfNoneMatch == nil {
			r.IfNoneMatch = &String{}
		}
		r.IfNoneMatch.Id = m.IfNoneMatchPrimitiveElement.Id
		r.IfNoneMatch.Extension = m.IfNoneMatchPrimitiveElement.Extension
	}
	r.IfModifiedSince = m.IfModifiedSince
	if m.IfModifiedSincePrimitiveElement != nil {
		if r.IfModifiedSince == nil {
			r.IfModifiedSince = &Instant{}
		}
		r.IfModifiedSince.Id = m.IfModifiedSincePrimitiveElement.Id
		r.IfModifiedSince.Extension = m.IfModifiedSincePrimitiveElement.Extension
	}
	r.IfMatch = m.IfMatch
	if m.IfMatchPrimitiveElement != nil {
		if r.IfMatch == nil {
			r.IfMatch = &String{}
		}
		r.IfMatch.Id = m.IfMatchPrimitiveElement.Id
		r.IfMatch.Extension = m.IfMatchPrimitiveElement.Extension
	}
	r.IfNoneExist = m.IfNoneExist
	if m.IfNoneExistPrimitiveElement != nil {
		if r.IfNoneExist == nil {
			r.IfNoneExist = &String{}
		}
		r.IfNoneExist.Id = m.IfNoneExistPrimitiveElement.Id
		r.IfNoneExist.Extension = m.IfNoneExistPrimitiveElement.Extension
	}
	return nil
}
func (r BundleEntryRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfNoneMatch, xml.StartElement{Name: xml.Name{Local: "ifNoneMatch"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfModifiedSince, xml.StartElement{Name: xml.Name{Local: "ifModifiedSince"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfMatch, xml.StartElement{Name: xml.Name{Local: "ifMatch"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IfNoneExist, xml.StartElement{Name: xml.Name{Local: "ifNoneExist"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BundleEntryRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "method":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = v
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = v
			case "ifNoneMatch":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IfNoneMatch = &v
			case "ifModifiedSince":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IfModifiedSince = &v
			case "ifMatch":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IfMatch = &v
			case "ifNoneExist":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IfNoneExist = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BundleEntryRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
type BundleEntryResponse struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The status code returned by processing this entry. The status SHALL start with a 3 digit HTTP code (e.g. 404) and may contain the standard HTTP description associated with the status code.
	Status String
	// The location header created by processing this operation, populated if the operation returns a location.
	Location *Uri
	// The Etag for the resource, if the operation for the entry produced a versioned resource (see [Resource Metadata and Versioning](http.html#versioning) and [Managing Resource Contention](http.html#concurrency)).
	Etag *String
	// The date/time that the resource was modified on the server.
	LastModified *Instant
	// An OperationOutcome containing hints and warnings produced as part of processing this entry in a batch or transaction.
	Outcome *model.Resource
}
type jsonBundleEntryResponse struct {
	Id                           *string            `json:"id,omitempty"`
	Extension                    []Extension        `json:"extension,omitempty"`
	ModifierExtension            []Extension        `json:"modifierExtension,omitempty"`
	Status                       String             `json:"status,omitempty"`
	StatusPrimitiveElement       *primitiveElement  `json:"_status,omitempty"`
	Location                     *Uri               `json:"location,omitempty"`
	LocationPrimitiveElement     *primitiveElement  `json:"_location,omitempty"`
	Etag                         *String            `json:"etag,omitempty"`
	EtagPrimitiveElement         *primitiveElement  `json:"_etag,omitempty"`
	LastModified                 *Instant           `json:"lastModified,omitempty"`
	LastModifiedPrimitiveElement *primitiveElement  `json:"_lastModified,omitempty"`
	Outcome                      *ContainedResource `json:"outcome,omitempty"`
}

func (r BundleEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BundleEntryResponse) marshalJSON() jsonBundleEntryResponse {
	m := jsonBundleEntryResponse{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Location != nil && r.Location.Value != nil {
		m.Location = r.Location
	}
	if r.Location != nil && (r.Location.Id != nil || r.Location.Extension != nil) {
		m.LocationPrimitiveElement = &primitiveElement{Id: r.Location.Id, Extension: r.Location.Extension}
	}
	if r.Etag != nil && r.Etag.Value != nil {
		m.Etag = r.Etag
	}
	if r.Etag != nil && (r.Etag.Id != nil || r.Etag.Extension != nil) {
		m.EtagPrimitiveElement = &primitiveElement{Id: r.Etag.Id, Extension: r.Etag.Extension}
	}
	if r.LastModified != nil && r.LastModified.Value != nil {
		m.LastModified = r.LastModified
	}
	if r.LastModified != nil && (r.LastModified.Id != nil || r.LastModified.Extension != nil) {
		m.LastModifiedPrimitiveElement = &primitiveElement{Id: r.LastModified.Id, Extension: r.LastModified.Extension}
	}
	if r.Outcome != nil {
		m.Outcome = &ContainedResource{*r.Outcome}
	}
	return m
}
func (r *BundleEntryResponse) UnmarshalJSON(b []byte) error {
	var m jsonBundleEntryResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BundleEntryResponse) unmarshalJSON(m jsonBundleEntryResponse) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Location = m.Location
	if m.LocationPrimitiveElement != nil {
		if r.Location == nil {
			r.Location = &Uri{}
		}
		r.Location.Id = m.LocationPrimitiveElement.Id
		r.Location.Extension = m.LocationPrimitiveElement.Extension
	}
	r.Etag = m.Etag
	if m.EtagPrimitiveElement != nil {
		if r.Etag == nil {
			r.Etag = &String{}
		}
		r.Etag.Id = m.EtagPrimitiveElement.Id
		r.Etag.Extension = m.EtagPrimitiveElement.Extension
	}
	r.LastModified = m.LastModified
	if m.LastModifiedPrimitiveElement != nil {
		if r.LastModified == nil {
			r.LastModified = &Instant{}
		}
		r.LastModified.Id = m.LastModifiedPrimitiveElement.Id
		r.LastModified.Extension = m.LastModifiedPrimitiveElement.Extension
	}
	if m.Outcome != nil {
		r.Outcome = &m.Outcome.Resource
	}
	return nil
}
func (r BundleEntryResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Etag, xml.StartElement{Name: xml.Name{Local: "etag"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastModified, xml.StartElement{Name: xml.Name{Local: "lastModified"}})
	if err != nil {
		return err
	}
	if r.Outcome != nil {
		v := ContainedResource{*r.Outcome}
		err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "outcome"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BundleEntryResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "status":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "location":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "etag":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Etag = &v
			case "lastModified":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastModified = &v
			case "outcome":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Outcome = &c.Resource
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BundleEntryResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
