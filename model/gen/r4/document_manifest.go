package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A single identifier that uniquely identifies this manifest. Principally used to refer to the manifest in non-FHIR contexts.
	MasterIdentifier *Identifier
	// Other identifiers associated with the document manifest, including version independent  identifiers.
	Identifier []Identifier
	// The status of this document manifest.
	Status Code
	// The code specifying the type of clinical activity that resulted in placing the associated content into the DocumentManifest.
	Type *CodeableConcept
	// Who or what the set of documents is about. The documents can be about a person, (patient or healthcare practitioner), a device (i.e. machine) or even a group of subjects (such as a document about a herd of farm animals, or a set of patients that share a common exposure). If the documents cross more than one subject, then more than one subject is allowed here (unusual use case).
	Subject *Reference
	// When the document manifest was created for submission to the server (not necessarily the same thing as the actual resource last modified time, since it may be modified, replicated, etc.).
	Created *DateTime
	// Identifies who is the author of the manifest. Manifest author is not necessarly the author of the references included.
	Author []Reference
	// A patient, practitioner, or organization for which this set of documents is intended.
	Recipient []Reference
	// Identifies the source system, application, or software that produced the document manifest.
	Source *Uri
	// Human-readable description of the source document. This is sometimes known as the "title".
	Description *String
	// The list of Resources that consist of the parts of this manifest.
	Content []Reference
	// Related identifiers or resources associated with the DocumentManifest.
	Related []DocumentManifestRelated
}

func (r DocumentManifest) ResourceType() string {
	return "DocumentManifest"
}
func (r DocumentManifest) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonDocumentManifest struct {
	ResourceType                  string                    `json:"resourceType"`
	Id                            *Id                       `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement         `json:"_id,omitempty"`
	Meta                          *Meta                     `json:"meta,omitempty"`
	ImplicitRules                 *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                      *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement         `json:"_language,omitempty"`
	Text                          *Narrative                `json:"text,omitempty"`
	Contained                     []containedResource       `json:"contained,omitempty"`
	Extension                     []Extension               `json:"extension,omitempty"`
	ModifierExtension             []Extension               `json:"modifierExtension,omitempty"`
	MasterIdentifier              *Identifier               `json:"masterIdentifier,omitempty"`
	Identifier                    []Identifier              `json:"identifier,omitempty"`
	Status                        Code                      `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement         `json:"_status,omitempty"`
	Type                          *CodeableConcept          `json:"type,omitempty"`
	Subject                       *Reference                `json:"subject,omitempty"`
	Created                       *DateTime                 `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement         `json:"_created,omitempty"`
	Author                        []Reference               `json:"author,omitempty"`
	Recipient                     []Reference               `json:"recipient,omitempty"`
	Source                        *Uri                      `json:"source,omitempty"`
	SourcePrimitiveElement        *primitiveElement         `json:"_source,omitempty"`
	Description                   *String                   `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement         `json:"_description,omitempty"`
	Content                       []Reference               `json:"content,omitempty"`
	Related                       []DocumentManifestRelated `json:"related,omitempty"`
}

func (r DocumentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DocumentManifest) marshalJSON() jsonDocumentManifest {
	m := jsonDocumentManifest{}
	m.ResourceType = "DocumentManifest"
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
	m.Text = r.Text
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.MasterIdentifier = r.MasterIdentifier
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Type = r.Type
	m.Subject = r.Subject
	m.Created = r.Created
	if r.Created != nil && (r.Created.Id != nil || r.Created.Extension != nil) {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Author = r.Author
	m.Recipient = r.Recipient
	m.Source = r.Source
	if r.Source != nil && (r.Source.Id != nil || r.Source.Extension != nil) {
		m.SourcePrimitiveElement = &primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Content = r.Content
	m.Related = r.Related
	return m
}
func (r *DocumentManifest) UnmarshalJSON(b []byte) error {
	var m jsonDocumentManifest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DocumentManifest) unmarshalJSON(m jsonDocumentManifest) error {
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
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.MasterIdentifier = m.MasterIdentifier
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Subject = m.Subject
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Recipient = m.Recipient
	r.Source = m.Source
	if m.SourcePrimitiveElement != nil {
		r.Source.Id = m.SourcePrimitiveElement.Id
		r.Source.Extension = m.SourcePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Content = m.Content
	r.Related = m.Related
	return nil
}
func (r DocumentManifest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Related identifiers or resources associated with the DocumentManifest.
type DocumentManifestRelated struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Related identifier to this DocumentManifest.  For example, Order numbers, accession numbers, XDW workflow numbers.
	Identifier *Identifier
	// Related Resource to this DocumentManifest. For example, Order, ServiceRequest,  Procedure, EligibilityRequest, etc.
	Ref *Reference
}
type jsonDocumentManifestRelated struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Ref               *Reference  `json:"ref,omitempty"`
}

func (r DocumentManifestRelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DocumentManifestRelated) marshalJSON() jsonDocumentManifestRelated {
	m := jsonDocumentManifestRelated{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Ref = r.Ref
	return m
}
func (r *DocumentManifestRelated) UnmarshalJSON(b []byte) error {
	var m jsonDocumentManifestRelated
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DocumentManifestRelated) unmarshalJSON(m jsonDocumentManifestRelated) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Ref = m.Ref
	return nil
}
func (r DocumentManifestRelated) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
