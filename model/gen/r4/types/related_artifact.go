package types

import "encoding/json"

// Base StructureDefinition for RelatedArtifact Type: Related artifacts such as additional documentation, justification, or bibliographic references.
//
// Knowledge resources must be able to provide enough information for consumers of the content (and/or interventions or results produced by the content) to be able to determine and understand the justification for and evidence in support of the content.
type RelatedArtifact struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A short label that can be used to reference the citation from elsewhere in the containing artifact, such as a footnote index.
	Label *String
	// A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Display *String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// The type of relationship to the related artifact.
	Type Code
	// A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
	Citation *Markdown
	// A url for the artifact that can be followed to access the actual content.
	Url *Url
	// The document being referenced, represented as an attachment. This is exclusive with the resource element.
	Document *Attachment
	// The related resource, such as a library, value set, profile, or other knowledge resource.
	Resource *Canonical
}

func (s RelatedArtifact) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
