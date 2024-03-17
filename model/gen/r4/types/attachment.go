package types

import "encoding/json"

// Base StructureDefinition for Attachment Type: For referring to data content defined in other formats.
//
// Many models need to include data defined in other specifications that is complex and opaque to the healthcare model. This includes documents, media recordings, structured data, etc.
type Attachment struct {
	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	ContentType *Code
	// A label or set of text to display in place of the data.
	Title *String
	// The date that the attachment was first created.
	Creation *DateTime
	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *Base64Binary
	// A location where the data can be accessed.
	Url *Url
	// The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	Size *UnsignedInt
	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash *Base64Binary
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The human language of the content. The value can be any valid value according to BCP 47.
	Language *Code
}

func (s Attachment) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
