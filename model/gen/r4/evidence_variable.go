package r4

import (
	"encoding/json"
	"fmt"
)

// The EvidenceVariable resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
//
// Need to be able to define and reuse the definition of individual elements of a research question.
type EvidenceVariable struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An absolute URI that is used to identify this evidence variable when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this evidence variable is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the evidence variable is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this evidence variable when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the evidence variable when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the evidence variable author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *String
	// A natural language name identifying the evidence variable. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the evidence variable.
	Title *String
	// The short title provides an alternate title for use in informal descriptive contexts where the full, formal title is not necessary.
	ShortTitle *String
	// An explanatory or alternate title for the EvidenceVariable giving additional information about its content.
	Subtitle *String
	// The status of this evidence variable. Enables tracking the life-cycle of the content.
	Status Code
	// The date  (and optionally time) when the evidence variable was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the evidence variable changes.
	Date *DateTime
	// The name of the organization or individual that published the evidence variable.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the evidence variable from a consumer's perspective.
	Description *Markdown
	// A human-readable string to clarify or explain concepts about the resource.
	Note []Annotation
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate evidence variable instances.
	UseContext []UsageContext
	// A legal or geographic region in which the evidence variable is intended to be used.
	Jurisdiction []CodeableConcept
	// A copyright statement relating to the evidence variable and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the evidence variable.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the evidence variable content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the EvidenceVariable. Topics provide a high-level categorization grouping types of EvidenceVariables that can be useful for filtering and searching.
	Topic []CodeableConcept
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []ContactDetail
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []ContactDetail
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []ContactDetail
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []ContactDetail
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []RelatedArtifact
	// The type of evidence element, a population, an exposure, or an outcome.
	Type *Code
	// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
	Characteristic []EvidenceVariableCharacteristic
}
type jsonEvidenceVariable struct {
	ResourceType                   string                           `json:"resourceType"`
	Id                             *Id                              `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                `json:"_id,omitempty"`
	Meta                           *Meta                            `json:"meta,omitempty"`
	ImplicitRules                  *Uri                             `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                `json:"_implicitRules,omitempty"`
	Language                       *Code                            `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                `json:"_language,omitempty"`
	Text                           *Narrative                       `json:"text,omitempty"`
	Contained                      []containedResource              `json:"contained,omitempty"`
	Extension                      []Extension                      `json:"extension,omitempty"`
	ModifierExtension              []Extension                      `json:"modifierExtension,omitempty"`
	Url                            *Uri                             `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement                `json:"_url,omitempty"`
	Identifier                     []Identifier                     `json:"identifier,omitempty"`
	Version                        *String                          `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement                `json:"_version,omitempty"`
	Name                           *String                          `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement                `json:"_name,omitempty"`
	Title                          *String                          `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement                `json:"_title,omitempty"`
	ShortTitle                     *String                          `json:"shortTitle,omitempty"`
	ShortTitlePrimitiveElement     *primitiveElement                `json:"_shortTitle,omitempty"`
	Subtitle                       *String                          `json:"subtitle,omitempty"`
	SubtitlePrimitiveElement       *primitiveElement                `json:"_subtitle,omitempty"`
	Status                         Code                             `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                `json:"_status,omitempty"`
	Date                           *DateTime                        `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement                `json:"_date,omitempty"`
	Publisher                      *String                          `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement                `json:"_publisher,omitempty"`
	Contact                        []ContactDetail                  `json:"contact,omitempty"`
	Description                    *Markdown                        `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement                `json:"_description,omitempty"`
	Note                           []Annotation                     `json:"note,omitempty"`
	UseContext                     []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept                `json:"jurisdiction,omitempty"`
	Copyright                      *Markdown                        `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement                `json:"_copyright,omitempty"`
	ApprovalDate                   *Date                            `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement                `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date                            `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement                `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period                          `json:"effectivePeriod,omitempty"`
	Topic                          []CodeableConcept                `json:"topic,omitempty"`
	Author                         []ContactDetail                  `json:"author,omitempty"`
	Editor                         []ContactDetail                  `json:"editor,omitempty"`
	Reviewer                       []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser                       []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact                []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Type                           *Code                            `json:"type,omitempty"`
	TypePrimitiveElement           *primitiveElement                `json:"_type,omitempty"`
	Characteristic                 []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
}

func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EvidenceVariable) marshalJSON() jsonEvidenceVariable {
	m := jsonEvidenceVariable{}
	m.ResourceType = "EvidenceVariable"
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
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.ShortTitle = r.ShortTitle
	if r.ShortTitle != nil && (r.ShortTitle.Id != nil || r.ShortTitle.Extension != nil) {
		m.ShortTitlePrimitiveElement = &primitiveElement{Id: r.ShortTitle.Id, Extension: r.ShortTitle.Extension}
	}
	m.Subtitle = r.Subtitle
	if r.Subtitle != nil && (r.Subtitle.Id != nil || r.Subtitle.Extension != nil) {
		m.SubtitlePrimitiveElement = &primitiveElement{Id: r.Subtitle.Id, Extension: r.Subtitle.Extension}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Publisher = r.Publisher
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Note = r.Note
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.ApprovalDate = r.ApprovalDate
	if r.ApprovalDate != nil && (r.ApprovalDate.Id != nil || r.ApprovalDate.Extension != nil) {
		m.ApprovalDatePrimitiveElement = &primitiveElement{Id: r.ApprovalDate.Id, Extension: r.ApprovalDate.Extension}
	}
	m.LastReviewDate = r.LastReviewDate
	if r.LastReviewDate != nil && (r.LastReviewDate.Id != nil || r.LastReviewDate.Extension != nil) {
		m.LastReviewDatePrimitiveElement = &primitiveElement{Id: r.LastReviewDate.Id, Extension: r.LastReviewDate.Extension}
	}
	m.EffectivePeriod = r.EffectivePeriod
	m.Topic = r.Topic
	m.Author = r.Author
	m.Editor = r.Editor
	m.Reviewer = r.Reviewer
	m.Endorser = r.Endorser
	m.RelatedArtifact = r.RelatedArtifact
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Characteristic = r.Characteristic
	return m
}
func (r *EvidenceVariable) UnmarshalJSON(b []byte) error {
	var m jsonEvidenceVariable
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EvidenceVariable) unmarshalJSON(m jsonEvidenceVariable) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.ShortTitle = m.ShortTitle
	if m.ShortTitlePrimitiveElement != nil {
		r.ShortTitle.Id = m.ShortTitlePrimitiveElement.Id
		r.ShortTitle.Extension = m.ShortTitlePrimitiveElement.Extension
	}
	r.Subtitle = m.Subtitle
	if m.SubtitlePrimitiveElement != nil {
		r.Subtitle.Id = m.SubtitlePrimitiveElement.Id
		r.Subtitle.Extension = m.SubtitlePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Note = m.Note
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.ApprovalDate = m.ApprovalDate
	if m.ApprovalDatePrimitiveElement != nil {
		r.ApprovalDate.Id = m.ApprovalDatePrimitiveElement.Id
		r.ApprovalDate.Extension = m.ApprovalDatePrimitiveElement.Extension
	}
	r.LastReviewDate = m.LastReviewDate
	if m.LastReviewDatePrimitiveElement != nil {
		r.LastReviewDate.Id = m.LastReviewDatePrimitiveElement.Id
		r.LastReviewDate.Extension = m.LastReviewDatePrimitiveElement.Extension
	}
	r.EffectivePeriod = m.EffectivePeriod
	r.Topic = m.Topic
	r.Author = m.Author
	r.Editor = m.Editor
	r.Reviewer = m.Reviewer
	r.Endorser = m.Endorser
	r.RelatedArtifact = m.RelatedArtifact
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Characteristic = m.Characteristic
	return nil
}
func (r EvidenceVariable) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
type EvidenceVariableCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A short, natural language description of the characteristic that could be used to communicate the criteria to an end-user.
	Description *String
	// Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
	Definition isEvidenceVariableCharacteristicDefinition
	// Use UsageContext to define the members of the population, such as Age Ranges, Genders, Settings.
	UsageContext []UsageContext
	// When true, members with this characteristic are excluded from the element.
	Exclude *Boolean
	// Indicates what effective period the study covers.
	ParticipantEffective isEvidenceVariableCharacteristicParticipantEffective
	// Indicates duration from the participant's study entry.
	TimeFromStart *Duration
	// Indicates how elements are aggregated within the study effective period.
	GroupMeasure *Code
}
type isEvidenceVariableCharacteristicDefinition interface {
	isEvidenceVariableCharacteristicDefinition()
}

func (r Reference) isEvidenceVariableCharacteristicDefinition()         {}
func (r Canonical) isEvidenceVariableCharacteristicDefinition()         {}
func (r CodeableConcept) isEvidenceVariableCharacteristicDefinition()   {}
func (r Expression) isEvidenceVariableCharacteristicDefinition()        {}
func (r DataRequirement) isEvidenceVariableCharacteristicDefinition()   {}
func (r TriggerDefinition) isEvidenceVariableCharacteristicDefinition() {}

type isEvidenceVariableCharacteristicParticipantEffective interface {
	isEvidenceVariableCharacteristicParticipantEffective()
}

func (r DateTime) isEvidenceVariableCharacteristicParticipantEffective() {}
func (r Period) isEvidenceVariableCharacteristicParticipantEffective()   {}
func (r Duration) isEvidenceVariableCharacteristicParticipantEffective() {}
func (r Timing) isEvidenceVariableCharacteristicParticipantEffective()   {}

type jsonEvidenceVariableCharacteristic struct {
	Id                                           *string            `json:"id,omitempty"`
	Extension                                    []Extension        `json:"extension,omitempty"`
	ModifierExtension                            []Extension        `json:"modifierExtension,omitempty"`
	Description                                  *String            `json:"description,omitempty"`
	DescriptionPrimitiveElement                  *primitiveElement  `json:"_description,omitempty"`
	DefinitionReference                          *Reference         `json:"definitionReference,omitempty"`
	DefinitionCanonical                          *Canonical         `json:"definitionCanonical,omitempty"`
	DefinitionCanonicalPrimitiveElement          *primitiveElement  `json:"_definitionCanonical,omitempty"`
	DefinitionCodeableConcept                    *CodeableConcept   `json:"definitionCodeableConcept,omitempty"`
	DefinitionExpression                         *Expression        `json:"definitionExpression,omitempty"`
	DefinitionDataRequirement                    *DataRequirement   `json:"definitionDataRequirement,omitempty"`
	DefinitionTriggerDefinition                  *TriggerDefinition `json:"definitionTriggerDefinition,omitempty"`
	UsageContext                                 []UsageContext     `json:"usageContext,omitempty"`
	Exclude                                      *Boolean           `json:"exclude,omitempty"`
	ExcludePrimitiveElement                      *primitiveElement  `json:"_exclude,omitempty"`
	ParticipantEffectiveDateTime                 *DateTime          `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectiveDateTimePrimitiveElement *primitiveElement  `json:"_participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod                   *Period            `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration                 *Duration          `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming                   *Timing            `json:"participantEffectiveTiming,omitempty"`
	TimeFromStart                                *Duration          `json:"timeFromStart,omitempty"`
	GroupMeasure                                 *Code              `json:"groupMeasure,omitempty"`
	GroupMeasurePrimitiveElement                 *primitiveElement  `json:"_groupMeasure,omitempty"`
}

func (r EvidenceVariableCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EvidenceVariableCharacteristic) marshalJSON() jsonEvidenceVariableCharacteristic {
	m := jsonEvidenceVariableCharacteristic{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	switch v := r.Definition.(type) {
	case Reference:
		m.DefinitionReference = &v
	case *Reference:
		m.DefinitionReference = v
	case Canonical:
		m.DefinitionCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.DefinitionCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.DefinitionCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.DefinitionCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.DefinitionCodeableConcept = &v
	case *CodeableConcept:
		m.DefinitionCodeableConcept = v
	case Expression:
		m.DefinitionExpression = &v
	case *Expression:
		m.DefinitionExpression = v
	case DataRequirement:
		m.DefinitionDataRequirement = &v
	case *DataRequirement:
		m.DefinitionDataRequirement = v
	case TriggerDefinition:
		m.DefinitionTriggerDefinition = &v
	case *TriggerDefinition:
		m.DefinitionTriggerDefinition = v
	}
	m.UsageContext = r.UsageContext
	m.Exclude = r.Exclude
	if r.Exclude != nil && (r.Exclude.Id != nil || r.Exclude.Extension != nil) {
		m.ExcludePrimitiveElement = &primitiveElement{Id: r.Exclude.Id, Extension: r.Exclude.Extension}
	}
	switch v := r.ParticipantEffective.(type) {
	case DateTime:
		m.ParticipantEffectiveDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ParticipantEffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.ParticipantEffectiveDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.ParticipantEffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.ParticipantEffectivePeriod = &v
	case *Period:
		m.ParticipantEffectivePeriod = v
	case Duration:
		m.ParticipantEffectiveDuration = &v
	case *Duration:
		m.ParticipantEffectiveDuration = v
	case Timing:
		m.ParticipantEffectiveTiming = &v
	case *Timing:
		m.ParticipantEffectiveTiming = v
	}
	m.TimeFromStart = r.TimeFromStart
	m.GroupMeasure = r.GroupMeasure
	if r.GroupMeasure != nil && (r.GroupMeasure.Id != nil || r.GroupMeasure.Extension != nil) {
		m.GroupMeasurePrimitiveElement = &primitiveElement{Id: r.GroupMeasure.Id, Extension: r.GroupMeasure.Extension}
	}
	return m
}
func (r *EvidenceVariableCharacteristic) UnmarshalJSON(b []byte) error {
	var m jsonEvidenceVariableCharacteristic
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EvidenceVariableCharacteristic) unmarshalJSON(m jsonEvidenceVariableCharacteristic) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	if m.DefinitionReference != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionReference
		r.Definition = v
	}
	if m.DefinitionCanonical != nil || m.DefinitionCanonicalPrimitiveElement != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionCanonical
		if m.DefinitionCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.DefinitionCanonicalPrimitiveElement.Id
			v.Extension = m.DefinitionCanonicalPrimitiveElement.Extension
		}
		r.Definition = v
	}
	if m.DefinitionCodeableConcept != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionCodeableConcept
		r.Definition = v
	}
	if m.DefinitionExpression != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionExpression
		r.Definition = v
	}
	if m.DefinitionDataRequirement != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionDataRequirement
		r.Definition = v
	}
	if m.DefinitionTriggerDefinition != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionTriggerDefinition
		r.Definition = v
	}
	r.UsageContext = m.UsageContext
	r.Exclude = m.Exclude
	if m.ExcludePrimitiveElement != nil {
		r.Exclude.Id = m.ExcludePrimitiveElement.Id
		r.Exclude.Extension = m.ExcludePrimitiveElement.Extension
	}
	if m.ParticipantEffectiveDateTime != nil || m.ParticipantEffectiveDateTimePrimitiveElement != nil {
		if r.ParticipantEffective != nil {
			return fmt.Errorf("multiple values for field \"ParticipantEffective\"")
		}
		v := m.ParticipantEffectiveDateTime
		if m.ParticipantEffectiveDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.ParticipantEffectiveDateTimePrimitiveElement.Id
			v.Extension = m.ParticipantEffectiveDateTimePrimitiveElement.Extension
		}
		r.ParticipantEffective = v
	}
	if m.ParticipantEffectivePeriod != nil {
		if r.ParticipantEffective != nil {
			return fmt.Errorf("multiple values for field \"ParticipantEffective\"")
		}
		v := m.ParticipantEffectivePeriod
		r.ParticipantEffective = v
	}
	if m.ParticipantEffectiveDuration != nil {
		if r.ParticipantEffective != nil {
			return fmt.Errorf("multiple values for field \"ParticipantEffective\"")
		}
		v := m.ParticipantEffectiveDuration
		r.ParticipantEffective = v
	}
	if m.ParticipantEffectiveTiming != nil {
		if r.ParticipantEffective != nil {
			return fmt.Errorf("multiple values for field \"ParticipantEffective\"")
		}
		v := m.ParticipantEffectiveTiming
		r.ParticipantEffective = v
	}
	r.TimeFromStart = m.TimeFromStart
	r.GroupMeasure = m.GroupMeasure
	if m.GroupMeasurePrimitiveElement != nil {
		r.GroupMeasure.Id = m.GroupMeasurePrimitiveElement.Id
		r.GroupMeasure.Extension = m.GroupMeasurePrimitiveElement.Extension
	}
	return nil
}
func (r EvidenceVariableCharacteristic) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
