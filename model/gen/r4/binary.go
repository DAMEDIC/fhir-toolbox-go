package r4

import "encoding/json"

// A resource that represents the data of a single raw artifact as digital content accessible in its native format.  A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
//
// There are situations where it is useful or required to handle pure binary content using the same framework as other resources.
type Binary struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// MimeType of the binary content represented as a standard MimeType (BCP 13).
	ContentType Code
	// This element identifies another resource that can be used as a proxy of the security sensitivity to use when deciding and enforcing access control rules for the Binary resource. Given that the Binary resource contains very few elements that can be used to determine the sensitivity of the data and relationships to individuals, the referenced resource stands in as a proxy equivalent for this purpose. This referenced resource may be related to the Binary (e.g. Media, DocumentReference), or may be some non-related Resource purely as a security proxy. E.g. to identify that the binary resource relates to a patient, and access should only be granted to applications that have access to the patient.
	SecurityContext *Reference
	// The actual content, base64 encoded.
	Data *Base64Binary
}
type jsonBinary struct {
	ResourceType                  string            `json:"resourceType"`
	Id                            *Id               `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement `json:"_id,omitempty"`
	Meta                          *Meta             `json:"meta,omitempty"`
	ImplicitRules                 *Uri              `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement `json:"_implicitRules,omitempty"`
	Language                      *Code             `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement `json:"_language,omitempty"`
	ContentType                   Code              `json:"contentType,omitempty"`
	ContentTypePrimitiveElement   *primitiveElement `json:"_contentType,omitempty"`
	SecurityContext               *Reference        `json:"securityContext,omitempty"`
	Data                          *Base64Binary     `json:"data,omitempty"`
	DataPrimitiveElement          *primitiveElement `json:"_data,omitempty"`
}

func (r Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Binary) marshalJSON() jsonBinary {
	m := jsonBinary{}
	m.ResourceType = "Binary"
	m.Id = r.Id
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	m.ImplicitRules = r.ImplicitRules
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.ContentType = r.ContentType
	if r.ContentType.Id != nil || r.ContentType.Extension != nil {
		m.ContentTypePrimitiveElement = &primitiveElement{Id: r.ContentType.Id, Extension: r.ContentType.Extension}
	}
	m.SecurityContext = r.SecurityContext
	m.Data = r.Data
	if r.Data != nil && (r.Data.Id != nil || r.Data.Extension != nil) {
		m.DataPrimitiveElement = &primitiveElement{Id: r.Data.Id, Extension: r.Data.Extension}
	}
	return m
}
func (r *Binary) UnmarshalJSON(b []byte) error {
	var m jsonBinary
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Binary) unmarshalJSON(m jsonBinary) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.ContentType = m.ContentType
	if m.ContentTypePrimitiveElement != nil {
		r.ContentType.Id = m.ContentTypePrimitiveElement.Id
		r.ContentType.Extension = m.ContentTypePrimitiveElement.Extension
	}
	r.SecurityContext = m.SecurityContext
	r.Data = m.Data
	if m.DataPrimitiveElement != nil {
		r.Data.Id = m.DataPrimitiveElement.Id
		r.Data.Extension = m.DataPrimitiveElement.Extension
	}
	return nil
}
func (r Binary) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
