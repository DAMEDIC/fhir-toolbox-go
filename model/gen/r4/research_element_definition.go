package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The ResearchElementDefinition resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
//
// Need to be able to define and reuse the definition of individual elements of a research question.
type ResearchElementDefinition struct {
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
	// An absolute URI that is used to identify this research element definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this research element definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the research element definition is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this research element definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the research element definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the research element definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *String
	// A natural language name identifying the research element definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the research element definition.
	Title *String
	// The short title provides an alternate title for use in informal descriptive contexts where the full, formal title is not necessary.
	ShortTitle *String
	// An explanatory or alternate title for the ResearchElementDefinition giving additional information about its content.
	Subtitle *String
	// The status of this research element definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this research element definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The intended subjects for the ResearchElementDefinition. If this element is not provided, a Patient subject is assumed, but the subject of the ResearchElementDefinition can be anything.
	Subject isResearchElementDefinitionSubject
	// The date  (and optionally time) when the research element definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the research element definition changes.
	Date *DateTime
	// The name of the organization or individual that published the research element definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the research element definition from a consumer's perspective.
	Description *Markdown
	// A human-readable string to clarify or explain concepts about the resource.
	Comment []String
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate research element definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the research element definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this research element definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// A detailed description, from a clinical perspective, of how the ResearchElementDefinition is used.
	Usage *String
	// A copyright statement relating to the research element definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the research element definition.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the research element definition content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the ResearchElementDefinition. Topics provide a high-level categorization grouping types of ResearchElementDefinitions that can be useful for filtering and searching.
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
	// A reference to a Library resource containing the formal logic used by the ResearchElementDefinition.
	Library []Canonical
	// The type of research element, a population, an exposure, or an outcome.
	Type Code
	// The type of the outcome (e.g. Dichotomous, Continuous, or Descriptive).
	VariableType *Code
	// A characteristic that defines the members of the research element. Multiple characteristics are applied with "and" semantics.
	Characteristic []ResearchElementDefinitionCharacteristic
}

func (r ResearchElementDefinition) ResourceType() string {
	return "ResearchElementDefinition"
}
func (r ResearchElementDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isResearchElementDefinitionSubject interface {
	isResearchElementDefinitionSubject()
}

func (r CodeableConcept) isResearchElementDefinitionSubject() {}
func (r Reference) isResearchElementDefinitionSubject()       {}

type jsonResearchElementDefinition struct {
	ResourceType                   string                                    `json:"resourceType"`
	Id                             *Id                                       `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                         `json:"_id,omitempty"`
	Meta                           *Meta                                     `json:"meta,omitempty"`
	ImplicitRules                  *Uri                                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                         `json:"_implicitRules,omitempty"`
	Language                       *Code                                     `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                         `json:"_language,omitempty"`
	Text                           *Narrative                                `json:"text,omitempty"`
	Contained                      []ContainedResource                       `json:"contained,omitempty"`
	Extension                      []Extension                               `json:"extension,omitempty"`
	ModifierExtension              []Extension                               `json:"modifierExtension,omitempty"`
	Url                            *Uri                                      `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement                         `json:"_url,omitempty"`
	Identifier                     []Identifier                              `json:"identifier,omitempty"`
	Version                        *String                                   `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement                         `json:"_version,omitempty"`
	Name                           *String                                   `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement                         `json:"_name,omitempty"`
	Title                          *String                                   `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement                         `json:"_title,omitempty"`
	ShortTitle                     *String                                   `json:"shortTitle,omitempty"`
	ShortTitlePrimitiveElement     *primitiveElement                         `json:"_shortTitle,omitempty"`
	Subtitle                       *String                                   `json:"subtitle,omitempty"`
	SubtitlePrimitiveElement       *primitiveElement                         `json:"_subtitle,omitempty"`
	Status                         Code                                      `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                         `json:"_status,omitempty"`
	Experimental                   *Boolean                                  `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement   *primitiveElement                         `json:"_experimental,omitempty"`
	SubjectCodeableConcept         *CodeableConcept                          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference               *Reference                                `json:"subjectReference,omitempty"`
	Date                           *DateTime                                 `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement                         `json:"_date,omitempty"`
	Publisher                      *String                                   `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement                         `json:"_publisher,omitempty"`
	Contact                        []ContactDetail                           `json:"contact,omitempty"`
	Description                    *Markdown                                 `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement                         `json:"_description,omitempty"`
	Comment                        []String                                  `json:"comment,omitempty"`
	CommentPrimitiveElement        []*primitiveElement                       `json:"_comment,omitempty"`
	UseContext                     []UsageContext                            `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept                         `json:"jurisdiction,omitempty"`
	Purpose                        *Markdown                                 `json:"purpose,omitempty"`
	PurposePrimitiveElement        *primitiveElement                         `json:"_purpose,omitempty"`
	Usage                          *String                                   `json:"usage,omitempty"`
	UsagePrimitiveElement          *primitiveElement                         `json:"_usage,omitempty"`
	Copyright                      *Markdown                                 `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement                         `json:"_copyright,omitempty"`
	ApprovalDate                   *Date                                     `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement                         `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date                                     `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement                         `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period                                   `json:"effectivePeriod,omitempty"`
	Topic                          []CodeableConcept                         `json:"topic,omitempty"`
	Author                         []ContactDetail                           `json:"author,omitempty"`
	Editor                         []ContactDetail                           `json:"editor,omitempty"`
	Reviewer                       []ContactDetail                           `json:"reviewer,omitempty"`
	Endorser                       []ContactDetail                           `json:"endorser,omitempty"`
	RelatedArtifact                []RelatedArtifact                         `json:"relatedArtifact,omitempty"`
	Library                        []Canonical                               `json:"library,omitempty"`
	LibraryPrimitiveElement        []*primitiveElement                       `json:"_library,omitempty"`
	Type                           Code                                      `json:"type,omitempty"`
	TypePrimitiveElement           *primitiveElement                         `json:"_type,omitempty"`
	VariableType                   *Code                                     `json:"variableType,omitempty"`
	VariableTypePrimitiveElement   *primitiveElement                         `json:"_variableType,omitempty"`
	Characteristic                 []ResearchElementDefinitionCharacteristic `json:"characteristic,omitempty"`
}

func (r ResearchElementDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ResearchElementDefinition) marshalJSON() jsonResearchElementDefinition {
	m := jsonResearchElementDefinition{}
	m.ResourceType = "ResearchElementDefinition"
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
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.ShortTitle != nil && r.ShortTitle.Value != nil {
		m.ShortTitle = r.ShortTitle
	}
	if r.ShortTitle != nil && (r.ShortTitle.Id != nil || r.ShortTitle.Extension != nil) {
		m.ShortTitlePrimitiveElement = &primitiveElement{Id: r.ShortTitle.Id, Extension: r.ShortTitle.Extension}
	}
	if r.Subtitle != nil && r.Subtitle.Value != nil {
		m.Subtitle = r.Subtitle
	}
	if r.Subtitle != nil && (r.Subtitle.Id != nil || r.Subtitle.Extension != nil) {
		m.SubtitlePrimitiveElement = &primitiveElement{Id: r.Subtitle.Id, Extension: r.Subtitle.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Experimental != nil && r.Experimental.Value != nil {
		m.Experimental = r.Experimental
	}
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	switch v := r.Subject.(type) {
	case CodeableConcept:
		m.SubjectCodeableConcept = &v
	case *CodeableConcept:
		m.SubjectCodeableConcept = v
	case Reference:
		m.SubjectReference = &v
	case *Reference:
		m.SubjectReference = v
	}
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	if r.Publisher != nil && r.Publisher.Value != nil {
		m.Publisher = r.Publisher
	}
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	anyCommentValue := false
	for _, e := range r.Comment {
		if e.Value != nil {
			anyCommentValue = true
			break
		}
	}
	if anyCommentValue {
		m.Comment = r.Comment
	}
	anyCommentIdOrExtension := false
	for _, e := range r.Comment {
		if e.Id != nil || e.Extension != nil {
			anyCommentIdOrExtension = true
			break
		}
	}
	if anyCommentIdOrExtension {
		m.CommentPrimitiveElement = make([]*primitiveElement, 0, len(r.Comment))
		for _, e := range r.Comment {
			if e.Id != nil || e.Extension != nil {
				m.CommentPrimitiveElement = append(m.CommentPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.CommentPrimitiveElement = append(m.CommentPrimitiveElement, nil)
			}
		}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	if r.Purpose != nil && r.Purpose.Value != nil {
		m.Purpose = r.Purpose
	}
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	if r.Usage != nil && r.Usage.Value != nil {
		m.Usage = r.Usage
	}
	if r.Usage != nil && (r.Usage.Id != nil || r.Usage.Extension != nil) {
		m.UsagePrimitiveElement = &primitiveElement{Id: r.Usage.Id, Extension: r.Usage.Extension}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	if r.ApprovalDate != nil && r.ApprovalDate.Value != nil {
		m.ApprovalDate = r.ApprovalDate
	}
	if r.ApprovalDate != nil && (r.ApprovalDate.Id != nil || r.ApprovalDate.Extension != nil) {
		m.ApprovalDatePrimitiveElement = &primitiveElement{Id: r.ApprovalDate.Id, Extension: r.ApprovalDate.Extension}
	}
	if r.LastReviewDate != nil && r.LastReviewDate.Value != nil {
		m.LastReviewDate = r.LastReviewDate
	}
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
	anyLibraryValue := false
	for _, e := range r.Library {
		if e.Value != nil {
			anyLibraryValue = true
			break
		}
	}
	if anyLibraryValue {
		m.Library = r.Library
	}
	anyLibraryIdOrExtension := false
	for _, e := range r.Library {
		if e.Id != nil || e.Extension != nil {
			anyLibraryIdOrExtension = true
			break
		}
	}
	if anyLibraryIdOrExtension {
		m.LibraryPrimitiveElement = make([]*primitiveElement, 0, len(r.Library))
		for _, e := range r.Library {
			if e.Id != nil || e.Extension != nil {
				m.LibraryPrimitiveElement = append(m.LibraryPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LibraryPrimitiveElement = append(m.LibraryPrimitiveElement, nil)
			}
		}
	}
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.VariableType != nil && r.VariableType.Value != nil {
		m.VariableType = r.VariableType
	}
	if r.VariableType != nil && (r.VariableType.Id != nil || r.VariableType.Extension != nil) {
		m.VariableTypePrimitiveElement = &primitiveElement{Id: r.VariableType.Id, Extension: r.VariableType.Extension}
	}
	m.Characteristic = r.Characteristic
	return m
}
func (r *ResearchElementDefinition) UnmarshalJSON(b []byte) error {
	var m jsonResearchElementDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ResearchElementDefinition) unmarshalJSON(m jsonResearchElementDefinition) error {
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
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.ShortTitle = m.ShortTitle
	if m.ShortTitlePrimitiveElement != nil {
		if r.ShortTitle == nil {
			r.ShortTitle = &String{}
		}
		r.ShortTitle.Id = m.ShortTitlePrimitiveElement.Id
		r.ShortTitle.Extension = m.ShortTitlePrimitiveElement.Extension
	}
	r.Subtitle = m.Subtitle
	if m.SubtitlePrimitiveElement != nil {
		if r.Subtitle == nil {
			r.Subtitle = &String{}
		}
		r.Subtitle.Id = m.SubtitlePrimitiveElement.Id
		r.Subtitle.Extension = m.SubtitlePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
		if r.Experimental == nil {
			r.Experimental = &Boolean{}
		}
		r.Experimental.Id = m.ExperimentalPrimitiveElement.Id
		r.Experimental.Extension = m.ExperimentalPrimitiveElement.Extension
	}
	if m.SubjectCodeableConcept != nil {
		if r.Subject != nil {
			return fmt.Errorf("multiple values for field \"Subject\"")
		}
		v := m.SubjectCodeableConcept
		r.Subject = v
	}
	if m.SubjectReference != nil {
		if r.Subject != nil {
			return fmt.Errorf("multiple values for field \"Subject\"")
		}
		v := m.SubjectReference
		r.Subject = v
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		if r.Publisher == nil {
			r.Publisher = &String{}
		}
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Comment = m.Comment
	for i, e := range m.CommentPrimitiveElement {
		if len(r.Comment) <= i {
			r.Comment = append(r.Comment, String{})
		}
		if e != nil {
			r.Comment[i].Id = e.Id
			r.Comment[i].Extension = e.Extension
		}
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		if r.Purpose == nil {
			r.Purpose = &Markdown{}
		}
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Usage = m.Usage
	if m.UsagePrimitiveElement != nil {
		if r.Usage == nil {
			r.Usage = &String{}
		}
		r.Usage.Id = m.UsagePrimitiveElement.Id
		r.Usage.Extension = m.UsagePrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.ApprovalDate = m.ApprovalDate
	if m.ApprovalDatePrimitiveElement != nil {
		if r.ApprovalDate == nil {
			r.ApprovalDate = &Date{}
		}
		r.ApprovalDate.Id = m.ApprovalDatePrimitiveElement.Id
		r.ApprovalDate.Extension = m.ApprovalDatePrimitiveElement.Extension
	}
	r.LastReviewDate = m.LastReviewDate
	if m.LastReviewDatePrimitiveElement != nil {
		if r.LastReviewDate == nil {
			r.LastReviewDate = &Date{}
		}
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
	r.Library = m.Library
	for i, e := range m.LibraryPrimitiveElement {
		if len(r.Library) <= i {
			r.Library = append(r.Library, Canonical{})
		}
		if e != nil {
			r.Library[i].Id = e.Id
			r.Library[i].Extension = e.Extension
		}
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.VariableType = m.VariableType
	if m.VariableTypePrimitiveElement != nil {
		if r.VariableType == nil {
			r.VariableType = &Code{}
		}
		r.VariableType.Id = m.VariableTypePrimitiveElement.Id
		r.VariableType.Extension = m.VariableTypePrimitiveElement.Extension
	}
	r.Characteristic = m.Characteristic
	return nil
}
func (r ResearchElementDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A characteristic that defines the members of the research element. Multiple characteristics are applied with "and" semantics.
type ResearchElementDefinitionCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Define members of the research element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
	Definition isResearchElementDefinitionCharacteristicDefinition
	// Use UsageContext to define the members of the population, such as Age Ranges, Genders, Settings.
	UsageContext []UsageContext
	// When true, members with this characteristic are excluded from the element.
	Exclude *Boolean
	// Specifies the UCUM unit for the outcome.
	UnitOfMeasure *CodeableConcept
	// A narrative description of the time period the study covers.
	StudyEffectiveDescription *String
	// Indicates what effective period the study covers.
	StudyEffective isResearchElementDefinitionCharacteristicStudyEffective
	// Indicates duration from the study initiation.
	StudyEffectiveTimeFromStart *Duration
	// Indicates how elements are aggregated within the study effective period.
	StudyEffectiveGroupMeasure *Code
	// A narrative description of the time period the study covers.
	ParticipantEffectiveDescription *String
	// Indicates what effective period the study covers.
	ParticipantEffective isResearchElementDefinitionCharacteristicParticipantEffective
	// Indicates duration from the participant's study entry.
	ParticipantEffectiveTimeFromStart *Duration
	// Indicates how elements are aggregated within the study effective period.
	ParticipantEffectiveGroupMeasure *Code
}
type isResearchElementDefinitionCharacteristicDefinition interface {
	isResearchElementDefinitionCharacteristicDefinition()
}

func (r CodeableConcept) isResearchElementDefinitionCharacteristicDefinition() {}
func (r Canonical) isResearchElementDefinitionCharacteristicDefinition()       {}
func (r Expression) isResearchElementDefinitionCharacteristicDefinition()      {}
func (r DataRequirement) isResearchElementDefinitionCharacteristicDefinition() {}

type isResearchElementDefinitionCharacteristicStudyEffective interface {
	isResearchElementDefinitionCharacteristicStudyEffective()
}

func (r DateTime) isResearchElementDefinitionCharacteristicStudyEffective() {}
func (r Period) isResearchElementDefinitionCharacteristicStudyEffective()   {}
func (r Duration) isResearchElementDefinitionCharacteristicStudyEffective() {}
func (r Timing) isResearchElementDefinitionCharacteristicStudyEffective()   {}

type isResearchElementDefinitionCharacteristicParticipantEffective interface {
	isResearchElementDefinitionCharacteristicParticipantEffective()
}

func (r DateTime) isResearchElementDefinitionCharacteristicParticipantEffective() {}
func (r Period) isResearchElementDefinitionCharacteristicParticipantEffective()   {}
func (r Duration) isResearchElementDefinitionCharacteristicParticipantEffective() {}
func (r Timing) isResearchElementDefinitionCharacteristicParticipantEffective()   {}

type jsonResearchElementDefinitionCharacteristic struct {
	Id                                               *string           `json:"id,omitempty"`
	Extension                                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                                []Extension       `json:"modifierExtension,omitempty"`
	DefinitionCodeableConcept                        *CodeableConcept  `json:"definitionCodeableConcept,omitempty"`
	DefinitionCanonical                              *Canonical        `json:"definitionCanonical,omitempty"`
	DefinitionCanonicalPrimitiveElement              *primitiveElement `json:"_definitionCanonical,omitempty"`
	DefinitionExpression                             *Expression       `json:"definitionExpression,omitempty"`
	DefinitionDataRequirement                        *DataRequirement  `json:"definitionDataRequirement,omitempty"`
	UsageContext                                     []UsageContext    `json:"usageContext,omitempty"`
	Exclude                                          *Boolean          `json:"exclude,omitempty"`
	ExcludePrimitiveElement                          *primitiveElement `json:"_exclude,omitempty"`
	UnitOfMeasure                                    *CodeableConcept  `json:"unitOfMeasure,omitempty"`
	StudyEffectiveDescription                        *String           `json:"studyEffectiveDescription,omitempty"`
	StudyEffectiveDescriptionPrimitiveElement        *primitiveElement `json:"_studyEffectiveDescription,omitempty"`
	StudyEffectiveDateTime                           *DateTime         `json:"studyEffectiveDateTime,omitempty"`
	StudyEffectiveDateTimePrimitiveElement           *primitiveElement `json:"_studyEffectiveDateTime,omitempty"`
	StudyEffectivePeriod                             *Period           `json:"studyEffectivePeriod,omitempty"`
	StudyEffectiveDuration                           *Duration         `json:"studyEffectiveDuration,omitempty"`
	StudyEffectiveTiming                             *Timing           `json:"studyEffectiveTiming,omitempty"`
	StudyEffectiveTimeFromStart                      *Duration         `json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure                       *Code             `json:"studyEffectiveGroupMeasure,omitempty"`
	StudyEffectiveGroupMeasurePrimitiveElement       *primitiveElement `json:"_studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription                  *String           `json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDescriptionPrimitiveElement  *primitiveElement `json:"_participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime                     *DateTime         `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectiveDateTimePrimitiveElement     *primitiveElement `json:"_participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod                       *Period           `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration                     *Duration         `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming                       *Timing           `json:"participantEffectiveTiming,omitempty"`
	ParticipantEffectiveTimeFromStart                *Duration         `json:"participantEffectiveTimeFromStart,omitempty"`
	ParticipantEffectiveGroupMeasure                 *Code             `json:"participantEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveGroupMeasurePrimitiveElement *primitiveElement `json:"_participantEffectiveGroupMeasure,omitempty"`
}

func (r ResearchElementDefinitionCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ResearchElementDefinitionCharacteristic) marshalJSON() jsonResearchElementDefinitionCharacteristic {
	m := jsonResearchElementDefinitionCharacteristic{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Definition.(type) {
	case CodeableConcept:
		m.DefinitionCodeableConcept = &v
	case *CodeableConcept:
		m.DefinitionCodeableConcept = v
	case Canonical:
		if v.Value != nil {
			m.DefinitionCanonical = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DefinitionCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		if v.Value != nil {
			m.DefinitionCanonical = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DefinitionCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Expression:
		m.DefinitionExpression = &v
	case *Expression:
		m.DefinitionExpression = v
	case DataRequirement:
		m.DefinitionDataRequirement = &v
	case *DataRequirement:
		m.DefinitionDataRequirement = v
	}
	m.UsageContext = r.UsageContext
	if r.Exclude != nil && r.Exclude.Value != nil {
		m.Exclude = r.Exclude
	}
	if r.Exclude != nil && (r.Exclude.Id != nil || r.Exclude.Extension != nil) {
		m.ExcludePrimitiveElement = &primitiveElement{Id: r.Exclude.Id, Extension: r.Exclude.Extension}
	}
	m.UnitOfMeasure = r.UnitOfMeasure
	if r.StudyEffectiveDescription != nil && r.StudyEffectiveDescription.Value != nil {
		m.StudyEffectiveDescription = r.StudyEffectiveDescription
	}
	if r.StudyEffectiveDescription != nil && (r.StudyEffectiveDescription.Id != nil || r.StudyEffectiveDescription.Extension != nil) {
		m.StudyEffectiveDescriptionPrimitiveElement = &primitiveElement{Id: r.StudyEffectiveDescription.Id, Extension: r.StudyEffectiveDescription.Extension}
	}
	switch v := r.StudyEffective.(type) {
	case DateTime:
		if v.Value != nil {
			m.StudyEffectiveDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.StudyEffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.StudyEffectiveDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.StudyEffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.StudyEffectivePeriod = &v
	case *Period:
		m.StudyEffectivePeriod = v
	case Duration:
		m.StudyEffectiveDuration = &v
	case *Duration:
		m.StudyEffectiveDuration = v
	case Timing:
		m.StudyEffectiveTiming = &v
	case *Timing:
		m.StudyEffectiveTiming = v
	}
	m.StudyEffectiveTimeFromStart = r.StudyEffectiveTimeFromStart
	if r.StudyEffectiveGroupMeasure != nil && r.StudyEffectiveGroupMeasure.Value != nil {
		m.StudyEffectiveGroupMeasure = r.StudyEffectiveGroupMeasure
	}
	if r.StudyEffectiveGroupMeasure != nil && (r.StudyEffectiveGroupMeasure.Id != nil || r.StudyEffectiveGroupMeasure.Extension != nil) {
		m.StudyEffectiveGroupMeasurePrimitiveElement = &primitiveElement{Id: r.StudyEffectiveGroupMeasure.Id, Extension: r.StudyEffectiveGroupMeasure.Extension}
	}
	if r.ParticipantEffectiveDescription != nil && r.ParticipantEffectiveDescription.Value != nil {
		m.ParticipantEffectiveDescription = r.ParticipantEffectiveDescription
	}
	if r.ParticipantEffectiveDescription != nil && (r.ParticipantEffectiveDescription.Id != nil || r.ParticipantEffectiveDescription.Extension != nil) {
		m.ParticipantEffectiveDescriptionPrimitiveElement = &primitiveElement{Id: r.ParticipantEffectiveDescription.Id, Extension: r.ParticipantEffectiveDescription.Extension}
	}
	switch v := r.ParticipantEffective.(type) {
	case DateTime:
		if v.Value != nil {
			m.ParticipantEffectiveDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ParticipantEffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.ParticipantEffectiveDateTime = v
		}
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
	m.ParticipantEffectiveTimeFromStart = r.ParticipantEffectiveTimeFromStart
	if r.ParticipantEffectiveGroupMeasure != nil && r.ParticipantEffectiveGroupMeasure.Value != nil {
		m.ParticipantEffectiveGroupMeasure = r.ParticipantEffectiveGroupMeasure
	}
	if r.ParticipantEffectiveGroupMeasure != nil && (r.ParticipantEffectiveGroupMeasure.Id != nil || r.ParticipantEffectiveGroupMeasure.Extension != nil) {
		m.ParticipantEffectiveGroupMeasurePrimitiveElement = &primitiveElement{Id: r.ParticipantEffectiveGroupMeasure.Id, Extension: r.ParticipantEffectiveGroupMeasure.Extension}
	}
	return m
}
func (r *ResearchElementDefinitionCharacteristic) UnmarshalJSON(b []byte) error {
	var m jsonResearchElementDefinitionCharacteristic
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ResearchElementDefinitionCharacteristic) unmarshalJSON(m jsonResearchElementDefinitionCharacteristic) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.DefinitionCodeableConcept != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionCodeableConcept
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
	r.UsageContext = m.UsageContext
	r.Exclude = m.Exclude
	if m.ExcludePrimitiveElement != nil {
		if r.Exclude == nil {
			r.Exclude = &Boolean{}
		}
		r.Exclude.Id = m.ExcludePrimitiveElement.Id
		r.Exclude.Extension = m.ExcludePrimitiveElement.Extension
	}
	r.UnitOfMeasure = m.UnitOfMeasure
	r.StudyEffectiveDescription = m.StudyEffectiveDescription
	if m.StudyEffectiveDescriptionPrimitiveElement != nil {
		if r.StudyEffectiveDescription == nil {
			r.StudyEffectiveDescription = &String{}
		}
		r.StudyEffectiveDescription.Id = m.StudyEffectiveDescriptionPrimitiveElement.Id
		r.StudyEffectiveDescription.Extension = m.StudyEffectiveDescriptionPrimitiveElement.Extension
	}
	if m.StudyEffectiveDateTime != nil || m.StudyEffectiveDateTimePrimitiveElement != nil {
		if r.StudyEffective != nil {
			return fmt.Errorf("multiple values for field \"StudyEffective\"")
		}
		v := m.StudyEffectiveDateTime
		if m.StudyEffectiveDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.StudyEffectiveDateTimePrimitiveElement.Id
			v.Extension = m.StudyEffectiveDateTimePrimitiveElement.Extension
		}
		r.StudyEffective = v
	}
	if m.StudyEffectivePeriod != nil {
		if r.StudyEffective != nil {
			return fmt.Errorf("multiple values for field \"StudyEffective\"")
		}
		v := m.StudyEffectivePeriod
		r.StudyEffective = v
	}
	if m.StudyEffectiveDuration != nil {
		if r.StudyEffective != nil {
			return fmt.Errorf("multiple values for field \"StudyEffective\"")
		}
		v := m.StudyEffectiveDuration
		r.StudyEffective = v
	}
	if m.StudyEffectiveTiming != nil {
		if r.StudyEffective != nil {
			return fmt.Errorf("multiple values for field \"StudyEffective\"")
		}
		v := m.StudyEffectiveTiming
		r.StudyEffective = v
	}
	r.StudyEffectiveTimeFromStart = m.StudyEffectiveTimeFromStart
	r.StudyEffectiveGroupMeasure = m.StudyEffectiveGroupMeasure
	if m.StudyEffectiveGroupMeasurePrimitiveElement != nil {
		if r.StudyEffectiveGroupMeasure == nil {
			r.StudyEffectiveGroupMeasure = &Code{}
		}
		r.StudyEffectiveGroupMeasure.Id = m.StudyEffectiveGroupMeasurePrimitiveElement.Id
		r.StudyEffectiveGroupMeasure.Extension = m.StudyEffectiveGroupMeasurePrimitiveElement.Extension
	}
	r.ParticipantEffectiveDescription = m.ParticipantEffectiveDescription
	if m.ParticipantEffectiveDescriptionPrimitiveElement != nil {
		if r.ParticipantEffectiveDescription == nil {
			r.ParticipantEffectiveDescription = &String{}
		}
		r.ParticipantEffectiveDescription.Id = m.ParticipantEffectiveDescriptionPrimitiveElement.Id
		r.ParticipantEffectiveDescription.Extension = m.ParticipantEffectiveDescriptionPrimitiveElement.Extension
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
	r.ParticipantEffectiveTimeFromStart = m.ParticipantEffectiveTimeFromStart
	r.ParticipantEffectiveGroupMeasure = m.ParticipantEffectiveGroupMeasure
	if m.ParticipantEffectiveGroupMeasurePrimitiveElement != nil {
		if r.ParticipantEffectiveGroupMeasure == nil {
			r.ParticipantEffectiveGroupMeasure = &Code{}
		}
		r.ParticipantEffectiveGroupMeasure.Id = m.ParticipantEffectiveGroupMeasurePrimitiveElement.Id
		r.ParticipantEffectiveGroupMeasure.Extension = m.ParticipantEffectiveGroupMeasurePrimitiveElement.Extension
	}
	return nil
}
func (r ResearchElementDefinitionCharacteristic) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
