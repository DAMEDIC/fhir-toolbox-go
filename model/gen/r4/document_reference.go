package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A reference to a document of any kind for any purpose. Provides metadata about the document so that the document can be discovered and managed. The scope of a document is any seralized object with a mime-type, so includes formal patient centric documents (CDA), cliical notes, scanned paper, and non-patient specific documents like policy text.
type DocumentReference struct {
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
	// Document identifier as assigned by the source of the document. This identifier is specific to this version of the document. This unique identifier may be used elsewhere to identify this version of the document.
	MasterIdentifier *Identifier
	// Other identifiers associated with the document, including version independent identifiers.
	Identifier []Identifier
	// The status of this document reference.
	Status Code
	// The status of the underlying document.
	DocStatus *Code
	// Specifies the particular kind of document referenced  (e.g. History and Physical, Discharge Summary, Progress Note). This usually equates to the purpose of making the document referenced.
	Type *CodeableConcept
	// A categorization for the type of document referenced - helps for indexing and searching. This may be implied by or derived from the code specified in the DocumentReference.type.
	Category []CodeableConcept
	// Who or what the document is about. The document can be about a person, (patient or healthcare practitioner), a device (e.g. a machine) or even a group of subjects (such as a document about a herd of farm animals, or a set of patients that share a common exposure).
	Subject *Reference
	// When the document reference was created.
	Date *Instant
	// Identifies who is responsible for adding the information to the document.
	Author []Reference
	// Which person or organization authenticates that this document is valid.
	Authenticator *Reference
	// Identifies the organization or group who is responsible for ongoing maintenance of and access to the document.
	Custodian *Reference
	// Relationships that this document has with other document references that already exist.
	RelatesTo []DocumentReferenceRelatesTo
	// Human-readable description of the source document.
	Description *String
	// A set of Security-Tag codes specifying the level of privacy/security of the Document. Note that DocumentReference.meta.security contains the security labels of the "reference" to the document, while DocumentReference.securityLabel contains a snapshot of the security labels on the document the reference refers to.
	SecurityLabel []CodeableConcept
	// The document and format referenced. There may be multiple content element repetitions, each with a different format.
	Content []DocumentReferenceContent
	// The clinical context in which the document was prepared.
	Context *DocumentReferenceContext
}

func (r DocumentReference) ResourceType() string {
	return "DocumentReference"
}

type jsonDocumentReference struct {
	ResourceType                  string                       `json:"resourceType"`
	Id                            *Id                          `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement            `json:"_id,omitempty"`
	Meta                          *Meta                        `json:"meta,omitempty"`
	ImplicitRules                 *Uri                         `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement            `json:"_implicitRules,omitempty"`
	Language                      *Code                        `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement            `json:"_language,omitempty"`
	Text                          *Narrative                   `json:"text,omitempty"`
	Contained                     []containedResource          `json:"contained,omitempty"`
	Extension                     []Extension                  `json:"extension,omitempty"`
	ModifierExtension             []Extension                  `json:"modifierExtension,omitempty"`
	MasterIdentifier              *Identifier                  `json:"masterIdentifier,omitempty"`
	Identifier                    []Identifier                 `json:"identifier,omitempty"`
	Status                        Code                         `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement            `json:"_status,omitempty"`
	DocStatus                     *Code                        `json:"docStatus,omitempty"`
	DocStatusPrimitiveElement     *primitiveElement            `json:"_docStatus,omitempty"`
	Type                          *CodeableConcept             `json:"type,omitempty"`
	Category                      []CodeableConcept            `json:"category,omitempty"`
	Subject                       *Reference                   `json:"subject,omitempty"`
	Date                          *Instant                     `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement            `json:"_date,omitempty"`
	Author                        []Reference                  `json:"author,omitempty"`
	Authenticator                 *Reference                   `json:"authenticator,omitempty"`
	Custodian                     *Reference                   `json:"custodian,omitempty"`
	RelatesTo                     []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Description                   *String                      `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement            `json:"_description,omitempty"`
	SecurityLabel                 []CodeableConcept            `json:"securityLabel,omitempty"`
	Content                       []DocumentReferenceContent   `json:"content,omitempty"`
	Context                       *DocumentReferenceContext    `json:"context,omitempty"`
}

func (r DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DocumentReference) marshalJSON() jsonDocumentReference {
	m := jsonDocumentReference{}
	m.ResourceType = "DocumentReference"
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
	m.DocStatus = r.DocStatus
	if r.DocStatus != nil && (r.DocStatus.Id != nil || r.DocStatus.Extension != nil) {
		m.DocStatusPrimitiveElement = &primitiveElement{Id: r.DocStatus.Id, Extension: r.DocStatus.Extension}
	}
	m.Type = r.Type
	m.Category = r.Category
	m.Subject = r.Subject
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Author = r.Author
	m.Authenticator = r.Authenticator
	m.Custodian = r.Custodian
	m.RelatesTo = r.RelatesTo
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.SecurityLabel = r.SecurityLabel
	m.Content = r.Content
	m.Context = r.Context
	return m
}
func (r *DocumentReference) UnmarshalJSON(b []byte) error {
	var m jsonDocumentReference
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DocumentReference) unmarshalJSON(m jsonDocumentReference) error {
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
	r.DocStatus = m.DocStatus
	if m.DocStatusPrimitiveElement != nil {
		r.DocStatus.Id = m.DocStatusPrimitiveElement.Id
		r.DocStatus.Extension = m.DocStatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Category = m.Category
	r.Subject = m.Subject
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Authenticator = m.Authenticator
	r.Custodian = m.Custodian
	r.RelatesTo = m.RelatesTo
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.SecurityLabel = m.SecurityLabel
	r.Content = m.Content
	r.Context = m.Context
	return nil
}
func (r DocumentReference) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Relationships that this document has with other document references that already exist.
type DocumentReferenceRelatesTo struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of relationship that this document has with anther document.
	Code Code
	// The target document of this relationship.
	Target Reference
}
type jsonDocumentReferenceRelatesTo struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Code                 Code              `json:"code,omitempty"`
	CodePrimitiveElement *primitiveElement `json:"_code,omitempty"`
	Target               Reference         `json:"target,omitempty"`
}

func (r DocumentReferenceRelatesTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DocumentReferenceRelatesTo) marshalJSON() jsonDocumentReferenceRelatesTo {
	m := jsonDocumentReferenceRelatesTo{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Target = r.Target
	return m
}
func (r *DocumentReferenceRelatesTo) UnmarshalJSON(b []byte) error {
	var m jsonDocumentReferenceRelatesTo
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DocumentReferenceRelatesTo) unmarshalJSON(m jsonDocumentReferenceRelatesTo) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Target = m.Target
	return nil
}
func (r DocumentReferenceRelatesTo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The document and format referenced. There may be multiple content element repetitions, each with a different format.
type DocumentReferenceContent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The document or URL of the document along with critical metadata to prove content has integrity.
	Attachment Attachment
	// An identifier of the document encoding, structure, and template that the document conforms to beyond the base format indicated in the mimeType.
	Format *Coding
}
type jsonDocumentReferenceContent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Attachment        Attachment  `json:"attachment,omitempty"`
	Format            *Coding     `json:"format,omitempty"`
}

func (r DocumentReferenceContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DocumentReferenceContent) marshalJSON() jsonDocumentReferenceContent {
	m := jsonDocumentReferenceContent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Attachment = r.Attachment
	m.Format = r.Format
	return m
}
func (r *DocumentReferenceContent) UnmarshalJSON(b []byte) error {
	var m jsonDocumentReferenceContent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DocumentReferenceContent) unmarshalJSON(m jsonDocumentReferenceContent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Attachment = m.Attachment
	r.Format = m.Format
	return nil
}
func (r DocumentReferenceContent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The clinical context in which the document was prepared.
type DocumentReferenceContext struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the clinical encounter or type of care that the document content is associated with.
	Encounter []Reference
	// This list of codes represents the main clinical acts, such as a colonoscopy or an appendectomy, being documented. In some cases, the event is inherent in the type Code, such as a "History and Physical Report" in which the procedure being documented is necessarily a "History and Physical" act.
	Event []CodeableConcept
	// The time period over which the service that is described by the document was provided.
	Period *Period
	// The kind of facility where the patient was seen.
	FacilityType *CodeableConcept
	// This property may convey specifics about the practice setting where the content was created, often reflecting the clinical specialty.
	PracticeSetting *CodeableConcept
	// The Patient Information as known when the document was published. May be a reference to a version specific, or contained.
	SourcePatientInfo *Reference
	// Related identifiers or resources associated with the DocumentReference.
	Related []Reference
}
type jsonDocumentReferenceContext struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Encounter         []Reference       `json:"encounter,omitempty"`
	Event             []CodeableConcept `json:"event,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	FacilityType      *CodeableConcept  `json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept  `json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference        `json:"sourcePatientInfo,omitempty"`
	Related           []Reference       `json:"related,omitempty"`
}

func (r DocumentReferenceContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DocumentReferenceContext) marshalJSON() jsonDocumentReferenceContext {
	m := jsonDocumentReferenceContext{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Encounter = r.Encounter
	m.Event = r.Event
	m.Period = r.Period
	m.FacilityType = r.FacilityType
	m.PracticeSetting = r.PracticeSetting
	m.SourcePatientInfo = r.SourcePatientInfo
	m.Related = r.Related
	return m
}
func (r *DocumentReferenceContext) UnmarshalJSON(b []byte) error {
	var m jsonDocumentReferenceContext
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DocumentReferenceContext) unmarshalJSON(m jsonDocumentReferenceContext) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Encounter = m.Encounter
	r.Event = m.Event
	r.Period = m.Period
	r.FacilityType = m.FacilityType
	r.PracticeSetting = m.PracticeSetting
	r.SourcePatientInfo = m.SourcePatientInfo
	r.Related = m.Related
	return nil
}
func (r DocumentReferenceContext) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
