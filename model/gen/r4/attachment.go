package r4

import "encoding/json"

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
type jsonAttachment struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ContentType                 *Code             `json:"contentType,omitempty"`
	ContentTypePrimitiveElement *primitiveElement `json:"_contentType,omitempty"`
	Language                    *Code             `json:"language,omitempty"`
	LanguagePrimitiveElement    *primitiveElement `json:"_language,omitempty"`
	Data                        *Base64Binary     `json:"data,omitempty"`
	DataPrimitiveElement        *primitiveElement `json:"_data,omitempty"`
	Url                         *Url              `json:"url,omitempty"`
	UrlPrimitiveElement         *primitiveElement `json:"_url,omitempty"`
	Size                        *UnsignedInt      `json:"size,omitempty"`
	SizePrimitiveElement        *primitiveElement `json:"_size,omitempty"`
	Hash                        *Base64Binary     `json:"hash,omitempty"`
	HashPrimitiveElement        *primitiveElement `json:"_hash,omitempty"`
	Title                       *String           `json:"title,omitempty"`
	TitlePrimitiveElement       *primitiveElement `json:"_title,omitempty"`
	Creation                    *DateTime         `json:"creation,omitempty"`
	CreationPrimitiveElement    *primitiveElement `json:"_creation,omitempty"`
}

func (r Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Attachment) marshalJSON() jsonAttachment {
	m := jsonAttachment{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.ContentType != nil && r.ContentType.Value != nil {
		m.ContentType = r.ContentType
	}
	if r.ContentType != nil && (r.ContentType.Id != nil || r.ContentType.Extension != nil) {
		m.ContentTypePrimitiveElement = &primitiveElement{Id: r.ContentType.Id, Extension: r.ContentType.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	if r.Data != nil && r.Data.Value != nil {
		m.Data = r.Data
	}
	if r.Data != nil && (r.Data.Id != nil || r.Data.Extension != nil) {
		m.DataPrimitiveElement = &primitiveElement{Id: r.Data.Id, Extension: r.Data.Extension}
	}
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.Size != nil && r.Size.Value != nil {
		m.Size = r.Size
	}
	if r.Size != nil && (r.Size.Id != nil || r.Size.Extension != nil) {
		m.SizePrimitiveElement = &primitiveElement{Id: r.Size.Id, Extension: r.Size.Extension}
	}
	if r.Hash != nil && r.Hash.Value != nil {
		m.Hash = r.Hash
	}
	if r.Hash != nil && (r.Hash.Id != nil || r.Hash.Extension != nil) {
		m.HashPrimitiveElement = &primitiveElement{Id: r.Hash.Id, Extension: r.Hash.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Creation != nil && r.Creation.Value != nil {
		m.Creation = r.Creation
	}
	if r.Creation != nil && (r.Creation.Id != nil || r.Creation.Extension != nil) {
		m.CreationPrimitiveElement = &primitiveElement{Id: r.Creation.Id, Extension: r.Creation.Extension}
	}
	return m
}
func (r *Attachment) UnmarshalJSON(b []byte) error {
	var m jsonAttachment
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Attachment) unmarshalJSON(m jsonAttachment) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ContentType = m.ContentType
	if m.ContentTypePrimitiveElement != nil {
		if r.ContentType == nil {
			r.ContentType = &Code{}
		}
		r.ContentType.Id = m.ContentTypePrimitiveElement.Id
		r.ContentType.Extension = m.ContentTypePrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Data = m.Data
	if m.DataPrimitiveElement != nil {
		if r.Data == nil {
			r.Data = &Base64Binary{}
		}
		r.Data.Id = m.DataPrimitiveElement.Id
		r.Data.Extension = m.DataPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Url{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Size = m.Size
	if m.SizePrimitiveElement != nil {
		if r.Size == nil {
			r.Size = &UnsignedInt{}
		}
		r.Size.Id = m.SizePrimitiveElement.Id
		r.Size.Extension = m.SizePrimitiveElement.Extension
	}
	r.Hash = m.Hash
	if m.HashPrimitiveElement != nil {
		if r.Hash == nil {
			r.Hash = &Base64Binary{}
		}
		r.Hash.Id = m.HashPrimitiveElement.Id
		r.Hash.Extension = m.HashPrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Creation = m.Creation
	if m.CreationPrimitiveElement != nil {
		if r.Creation == nil {
			r.Creation = &DateTime{}
		}
		r.Creation.Id = m.CreationPrimitiveElement.Id
		r.Creation.Extension = m.CreationPrimitiveElement.Extension
	}
	return nil
}
func (r Attachment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
