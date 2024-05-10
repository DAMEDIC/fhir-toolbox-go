package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition struct {
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
	// An absolute URI that is used to identify this plan definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this plan definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the plan definition is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this plan definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the plan definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the plan definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *String
	// A natural language name identifying the plan definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the plan definition.
	Title *String
	// An explanatory or alternate title for the plan definition giving additional information about its content.
	Subtitle *String
	// A high-level category for the plan definition that distinguishes the kinds of systems that would be interested in the plan definition.
	Type *CodeableConcept
	// The status of this plan definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this plan definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// A code or group definition that describes the intended subject of the plan definition.
	Subject isPlanDefinitionSubject
	// The date  (and optionally time) when the plan definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the plan definition changes.
	Date *DateTime
	// The name of the organization or individual that published the plan definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the plan definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate plan definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the plan definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this plan definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// A detailed description of how the plan definition is used from a clinical perspective.
	Usage *String
	// A copyright statement relating to the plan definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the plan definition.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the plan definition content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the plan definition. Topics provide a high-level categorization of the definition that can be useful for filtering and searching.
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
	// A reference to a Library resource containing any formal logic used by the plan definition.
	Library []Canonical
	// Goals that describe what the activities within the plan are intended to achieve. For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
	Goal []PlanDefinitionGoal
	// An action or group of actions to be taken as part of the plan.
	Action []PlanDefinitionAction
}
type isPlanDefinitionSubject interface {
	isPlanDefinitionSubject()
}

func (r CodeableConcept) isPlanDefinitionSubject() {}
func (r Reference) isPlanDefinitionSubject()       {}

type jsonPlanDefinition struct {
	ResourceType                   string                 `json:"resourceType"`
	Id                             *Id                    `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement      `json:"_id,omitempty"`
	Meta                           *Meta                  `json:"meta,omitempty"`
	ImplicitRules                  *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                       *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement      `json:"_language,omitempty"`
	Text                           *Narrative             `json:"text,omitempty"`
	Contained                      []containedResource    `json:"contained,omitempty"`
	Extension                      []Extension            `json:"extension,omitempty"`
	ModifierExtension              []Extension            `json:"modifierExtension,omitempty"`
	Url                            *Uri                   `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement      `json:"_url,omitempty"`
	Identifier                     []Identifier           `json:"identifier,omitempty"`
	Version                        *String                `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement      `json:"_version,omitempty"`
	Name                           *String                `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement      `json:"_name,omitempty"`
	Title                          *String                `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement      `json:"_title,omitempty"`
	Subtitle                       *String                `json:"subtitle,omitempty"`
	SubtitlePrimitiveElement       *primitiveElement      `json:"_subtitle,omitempty"`
	Type                           *CodeableConcept       `json:"type,omitempty"`
	Status                         Code                   `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement      `json:"_status,omitempty"`
	Experimental                   *Boolean               `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement   *primitiveElement      `json:"_experimental,omitempty"`
	SubjectCodeableConcept         *CodeableConcept       `json:"subjectCodeableConcept,omitempty"`
	SubjectReference               *Reference             `json:"subjectReference,omitempty"`
	Date                           *DateTime              `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement      `json:"_date,omitempty"`
	Publisher                      *String                `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement      `json:"_publisher,omitempty"`
	Contact                        []ContactDetail        `json:"contact,omitempty"`
	Description                    *Markdown              `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement      `json:"_description,omitempty"`
	UseContext                     []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                        *Markdown              `json:"purpose,omitempty"`
	PurposePrimitiveElement        *primitiveElement      `json:"_purpose,omitempty"`
	Usage                          *String                `json:"usage,omitempty"`
	UsagePrimitiveElement          *primitiveElement      `json:"_usage,omitempty"`
	Copyright                      *Markdown              `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement      `json:"_copyright,omitempty"`
	ApprovalDate                   *Date                  `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement      `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date                  `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement      `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period                `json:"effectivePeriod,omitempty"`
	Topic                          []CodeableConcept      `json:"topic,omitempty"`
	Author                         []ContactDetail        `json:"author,omitempty"`
	Editor                         []ContactDetail        `json:"editor,omitempty"`
	Reviewer                       []ContactDetail        `json:"reviewer,omitempty"`
	Endorser                       []ContactDetail        `json:"endorser,omitempty"`
	RelatedArtifact                []RelatedArtifact      `json:"relatedArtifact,omitempty"`
	Library                        []Canonical            `json:"library,omitempty"`
	LibraryPrimitiveElement        []*primitiveElement    `json:"_library,omitempty"`
	Goal                           []PlanDefinitionGoal   `json:"goal,omitempty"`
	Action                         []PlanDefinitionAction `json:"action,omitempty"`
}

func (r PlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinition) marshalJSON() jsonPlanDefinition {
	m := jsonPlanDefinition{}
	m.ResourceType = "PlanDefinition"
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
	m.Subtitle = r.Subtitle
	if r.Subtitle != nil && (r.Subtitle.Id != nil || r.Subtitle.Extension != nil) {
		m.SubtitlePrimitiveElement = &primitiveElement{Id: r.Subtitle.Id, Extension: r.Subtitle.Extension}
	}
	m.Type = r.Type
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Experimental = r.Experimental
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
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Usage = r.Usage
	if r.Usage != nil && (r.Usage.Id != nil || r.Usage.Extension != nil) {
		m.UsagePrimitiveElement = &primitiveElement{Id: r.Usage.Id, Extension: r.Usage.Extension}
	}
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
	m.Library = r.Library
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
	m.Goal = r.Goal
	m.Action = r.Action
	return m
}
func (r *PlanDefinition) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinition) unmarshalJSON(m jsonPlanDefinition) error {
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
	r.Subtitle = m.Subtitle
	if m.SubtitlePrimitiveElement != nil {
		r.Subtitle.Id = m.SubtitlePrimitiveElement.Id
		r.Subtitle.Extension = m.SubtitlePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
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
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Usage = m.Usage
	if m.UsagePrimitiveElement != nil {
		r.Usage.Id = m.UsagePrimitiveElement.Id
		r.Usage.Extension = m.UsagePrimitiveElement.Extension
	}
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
	r.Library = m.Library
	for i, e := range m.LibraryPrimitiveElement {
		if len(r.Library) > i {
			r.Library[i].Id = e.Id
			r.Library[i].Extension = e.Extension
		} else {
			r.Library = append(r.Library, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Goal = m.Goal
	r.Action = m.Action
	return nil
}
func (r PlanDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Goals that describe what the activities within the plan are intended to achieve. For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
type PlanDefinitionGoal struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a category the goal falls within.
	Category *CodeableConcept
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	Description CodeableConcept
	// Identifies the expected level of importance associated with reaching/sustaining the defined goal.
	Priority *CodeableConcept
	// The event after which the goal should begin being pursued.
	Start *CodeableConcept
	// Identifies problems, conditions, issues, or concerns the goal is intended to address.
	Addresses []CodeableConcept
	// Didactic or other informational resources associated with the goal that provide further supporting information about the goal. Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact
	// Indicates what should be done and within what timeframe.
	Target []PlanDefinitionGoalTarget
}
type jsonPlanDefinitionGoal struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept           `json:"category,omitempty"`
	Description       CodeableConcept            `json:"description,omitempty"`
	Priority          *CodeableConcept           `json:"priority,omitempty"`
	Start             *CodeableConcept           `json:"start,omitempty"`
	Addresses         []CodeableConcept          `json:"addresses,omitempty"`
	Documentation     []RelatedArtifact          `json:"documentation,omitempty"`
	Target            []PlanDefinitionGoalTarget `json:"target,omitempty"`
}

func (r PlanDefinitionGoal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionGoal) marshalJSON() jsonPlanDefinitionGoal {
	m := jsonPlanDefinitionGoal{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Description = r.Description
	m.Priority = r.Priority
	m.Start = r.Start
	m.Addresses = r.Addresses
	m.Documentation = r.Documentation
	m.Target = r.Target
	return m
}
func (r *PlanDefinitionGoal) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionGoal
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionGoal) unmarshalJSON(m jsonPlanDefinitionGoal) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Description = m.Description
	r.Priority = m.Priority
	r.Start = m.Start
	r.Addresses = m.Addresses
	r.Documentation = m.Documentation
	r.Target = m.Target
	return nil
}
func (r PlanDefinitionGoal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates what should be done and within what timeframe.
type PlanDefinitionGoalTarget struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The parameter whose value is to be tracked, e.g. body weight, blood pressure, or hemoglobin A1c level.
	Measure *CodeableConcept
	// The target value of the measure to be achieved to signify fulfillment of the goal, e.g. 150 pounds or 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any value at or above the low value.
	Detail isPlanDefinitionGoalTargetDetail
	// Indicates the timeframe after the start of the goal in which the goal should be met.
	Due *Duration
}
type isPlanDefinitionGoalTargetDetail interface {
	isPlanDefinitionGoalTargetDetail()
}

func (r Quantity) isPlanDefinitionGoalTargetDetail()        {}
func (r Range) isPlanDefinitionGoalTargetDetail()           {}
func (r CodeableConcept) isPlanDefinitionGoalTargetDetail() {}

type jsonPlanDefinitionGoalTarget struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `json:"measure,omitempty"`
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range           `json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
	Due                   *Duration        `json:"due,omitempty"`
}

func (r PlanDefinitionGoalTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionGoalTarget) marshalJSON() jsonPlanDefinitionGoalTarget {
	m := jsonPlanDefinitionGoalTarget{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Measure = r.Measure
	switch v := r.Detail.(type) {
	case Quantity:
		m.DetailQuantity = &v
	case *Quantity:
		m.DetailQuantity = v
	case Range:
		m.DetailRange = &v
	case *Range:
		m.DetailRange = v
	case CodeableConcept:
		m.DetailCodeableConcept = &v
	case *CodeableConcept:
		m.DetailCodeableConcept = v
	}
	m.Due = r.Due
	return m
}
func (r *PlanDefinitionGoalTarget) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionGoalTarget
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionGoalTarget) unmarshalJSON(m jsonPlanDefinitionGoalTarget) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Measure = m.Measure
	if m.DetailQuantity != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailQuantity
		r.Detail = v
	}
	if m.DetailRange != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailRange
		r.Detail = v
	}
	if m.DetailCodeableConcept != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailCodeableConcept
		r.Detail = v
	}
	r.Due = m.Due
	return nil
}
func (r PlanDefinitionGoalTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An action or group of actions to be taken as part of the plan.
type PlanDefinitionAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A user-visible prefix for the action.
	Prefix *String
	// The title of the action displayed to a user.
	Title *String
	// A brief description of the action used to provide a summary to display to the user.
	Description *String
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *String
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *Code
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for the section of a documentation template.
	Code []CodeableConcept
	// A description of why this action is necessary or appropriate.
	Reason []CodeableConcept
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact
	// Identifies goals that this action supports. The reference must be to a goal element defined within this plan definition.
	GoalId []Id
	// A code or group definition that describes the intended subject of the action and its children, if any.
	Subject isPlanDefinitionActionSubject
	// A description of when the action should be triggered.
	Trigger []TriggerDefinition
	// An expression that describes applicability criteria or start/stop conditions for the action.
	Condition []PlanDefinitionActionCondition
	// Defines input data requirements for the action.
	Input []DataRequirement
	// Defines the outputs of the action, if any.
	Output []DataRequirement
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []PlanDefinitionActionRelatedAction
	// An optional value describing when the action should be performed.
	Timing isPlanDefinitionActionTiming
	// Indicates who should participate in performing the action described.
	Participant []PlanDefinitionActionParticipant
	// The type of action to perform (create, update, remove).
	Type *CodeableConcept
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *Code
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *Code
	// Defines the required behavior for the action.
	RequiredBehavior *Code
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *Code
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *Code
	// A reference to an ActivityDefinition that describes the action to be taken in detail, or a PlanDefinition that describes a series of actions to be taken.
	Definition isPlanDefinitionActionDefinition
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	Transform *Canonical
	// Customizations that should be applied to the statically defined resource. For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
	DynamicValue []PlanDefinitionActionDynamicValue
	// Sub actions that are contained within the action. The behavior of this action determines the functionality of the sub-actions. For example, a selection behavior of at-most-one indicates that of the sub-actions, at most one may be chosen as part of realizing the action definition.
	Action []PlanDefinitionAction
}
type isPlanDefinitionActionSubject interface {
	isPlanDefinitionActionSubject()
}

func (r CodeableConcept) isPlanDefinitionActionSubject() {}
func (r Reference) isPlanDefinitionActionSubject()       {}

type isPlanDefinitionActionTiming interface {
	isPlanDefinitionActionTiming()
}

func (r DateTime) isPlanDefinitionActionTiming() {}
func (r Age) isPlanDefinitionActionTiming()      {}
func (r Period) isPlanDefinitionActionTiming()   {}
func (r Duration) isPlanDefinitionActionTiming() {}
func (r Range) isPlanDefinitionActionTiming()    {}
func (r Timing) isPlanDefinitionActionTiming()   {}

type isPlanDefinitionActionDefinition interface {
	isPlanDefinitionActionDefinition()
}

func (r Canonical) isPlanDefinitionActionDefinition() {}
func (r Uri) isPlanDefinitionActionDefinition()       {}

type jsonPlanDefinitionAction struct {
	Id                                  *string                             `json:"id,omitempty"`
	Extension                           []Extension                         `json:"extension,omitempty"`
	ModifierExtension                   []Extension                         `json:"modifierExtension,omitempty"`
	Prefix                              *String                             `json:"prefix,omitempty"`
	PrefixPrimitiveElement              *primitiveElement                   `json:"_prefix,omitempty"`
	Title                               *String                             `json:"title,omitempty"`
	TitlePrimitiveElement               *primitiveElement                   `json:"_title,omitempty"`
	Description                         *String                             `json:"description,omitempty"`
	DescriptionPrimitiveElement         *primitiveElement                   `json:"_description,omitempty"`
	TextEquivalent                      *String                             `json:"textEquivalent,omitempty"`
	TextEquivalentPrimitiveElement      *primitiveElement                   `json:"_textEquivalent,omitempty"`
	Priority                            *Code                               `json:"priority,omitempty"`
	PriorityPrimitiveElement            *primitiveElement                   `json:"_priority,omitempty"`
	Code                                []CodeableConcept                   `json:"code,omitempty"`
	Reason                              []CodeableConcept                   `json:"reason,omitempty"`
	Documentation                       []RelatedArtifact                   `json:"documentation,omitempty"`
	GoalId                              []Id                                `json:"goalId,omitempty"`
	GoalIdPrimitiveElement              []*primitiveElement                 `json:"_goalId,omitempty"`
	SubjectCodeableConcept              *CodeableConcept                    `json:"subjectCodeableConcept,omitempty"`
	SubjectReference                    *Reference                          `json:"subjectReference,omitempty"`
	Trigger                             []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition                           []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                               []DataRequirement                   `json:"input,omitempty"`
	Output                              []DataRequirement                   `json:"output,omitempty"`
	RelatedAction                       []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime                      *DateTime                           `json:"timingDateTime,omitempty"`
	TimingDateTimePrimitiveElement      *primitiveElement                   `json:"_timingDateTime,omitempty"`
	TimingAge                           *Age                                `json:"timingAge,omitempty"`
	TimingPeriod                        *Period                             `json:"timingPeriod,omitempty"`
	TimingDuration                      *Duration                           `json:"timingDuration,omitempty"`
	TimingRange                         *Range                              `json:"timingRange,omitempty"`
	TimingTiming                        *Timing                             `json:"timingTiming,omitempty"`
	Participant                         []PlanDefinitionActionParticipant   `json:"participant,omitempty"`
	Type                                *CodeableConcept                    `json:"type,omitempty"`
	GroupingBehavior                    *Code                               `json:"groupingBehavior,omitempty"`
	GroupingBehaviorPrimitiveElement    *primitiveElement                   `json:"_groupingBehavior,omitempty"`
	SelectionBehavior                   *Code                               `json:"selectionBehavior,omitempty"`
	SelectionBehaviorPrimitiveElement   *primitiveElement                   `json:"_selectionBehavior,omitempty"`
	RequiredBehavior                    *Code                               `json:"requiredBehavior,omitempty"`
	RequiredBehaviorPrimitiveElement    *primitiveElement                   `json:"_requiredBehavior,omitempty"`
	PrecheckBehavior                    *Code                               `json:"precheckBehavior,omitempty"`
	PrecheckBehaviorPrimitiveElement    *primitiveElement                   `json:"_precheckBehavior,omitempty"`
	CardinalityBehavior                 *Code                               `json:"cardinalityBehavior,omitempty"`
	CardinalityBehaviorPrimitiveElement *primitiveElement                   `json:"_cardinalityBehavior,omitempty"`
	DefinitionCanonical                 *Canonical                          `json:"definitionCanonical,omitempty"`
	DefinitionCanonicalPrimitiveElement *primitiveElement                   `json:"_definitionCanonical,omitempty"`
	DefinitionUri                       *Uri                                `json:"definitionUri,omitempty"`
	DefinitionUriPrimitiveElement       *primitiveElement                   `json:"_definitionUri,omitempty"`
	Transform                           *Canonical                          `json:"transform,omitempty"`
	TransformPrimitiveElement           *primitiveElement                   `json:"_transform,omitempty"`
	DynamicValue                        []PlanDefinitionActionDynamicValue  `json:"dynamicValue,omitempty"`
	Action                              []PlanDefinitionAction              `json:"action,omitempty"`
}

func (r PlanDefinitionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionAction) marshalJSON() jsonPlanDefinitionAction {
	m := jsonPlanDefinitionAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Prefix = r.Prefix
	if r.Prefix != nil && (r.Prefix.Id != nil || r.Prefix.Extension != nil) {
		m.PrefixPrimitiveElement = &primitiveElement{Id: r.Prefix.Id, Extension: r.Prefix.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.TextEquivalent = r.TextEquivalent
	if r.TextEquivalent != nil && (r.TextEquivalent.Id != nil || r.TextEquivalent.Extension != nil) {
		m.TextEquivalentPrimitiveElement = &primitiveElement{Id: r.TextEquivalent.Id, Extension: r.TextEquivalent.Extension}
	}
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.Code = r.Code
	m.Reason = r.Reason
	m.Documentation = r.Documentation
	m.GoalId = r.GoalId
	anyGoalIdIdOrExtension := false
	for _, e := range r.GoalId {
		if e.Id != nil || e.Extension != nil {
			anyGoalIdIdOrExtension = true
			break
		}
	}
	if anyGoalIdIdOrExtension {
		m.GoalIdPrimitiveElement = make([]*primitiveElement, 0, len(r.GoalId))
		for _, e := range r.GoalId {
			if e.Id != nil || e.Extension != nil {
				m.GoalIdPrimitiveElement = append(m.GoalIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.GoalIdPrimitiveElement = append(m.GoalIdPrimitiveElement, nil)
			}
		}
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
	m.Trigger = r.Trigger
	m.Condition = r.Condition
	m.Input = r.Input
	m.Output = r.Output
	m.RelatedAction = r.RelatedAction
	switch v := r.Timing.(type) {
	case DateTime:
		m.TimingDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.TimingDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.TimingDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.TimingDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Age:
		m.TimingAge = &v
	case *Age:
		m.TimingAge = v
	case Period:
		m.TimingPeriod = &v
	case *Period:
		m.TimingPeriod = v
	case Duration:
		m.TimingDuration = &v
	case *Duration:
		m.TimingDuration = v
	case Range:
		m.TimingRange = &v
	case *Range:
		m.TimingRange = v
	case Timing:
		m.TimingTiming = &v
	case *Timing:
		m.TimingTiming = v
	}
	m.Participant = r.Participant
	m.Type = r.Type
	m.GroupingBehavior = r.GroupingBehavior
	if r.GroupingBehavior != nil && (r.GroupingBehavior.Id != nil || r.GroupingBehavior.Extension != nil) {
		m.GroupingBehaviorPrimitiveElement = &primitiveElement{Id: r.GroupingBehavior.Id, Extension: r.GroupingBehavior.Extension}
	}
	m.SelectionBehavior = r.SelectionBehavior
	if r.SelectionBehavior != nil && (r.SelectionBehavior.Id != nil || r.SelectionBehavior.Extension != nil) {
		m.SelectionBehaviorPrimitiveElement = &primitiveElement{Id: r.SelectionBehavior.Id, Extension: r.SelectionBehavior.Extension}
	}
	m.RequiredBehavior = r.RequiredBehavior
	if r.RequiredBehavior != nil && (r.RequiredBehavior.Id != nil || r.RequiredBehavior.Extension != nil) {
		m.RequiredBehaviorPrimitiveElement = &primitiveElement{Id: r.RequiredBehavior.Id, Extension: r.RequiredBehavior.Extension}
	}
	m.PrecheckBehavior = r.PrecheckBehavior
	if r.PrecheckBehavior != nil && (r.PrecheckBehavior.Id != nil || r.PrecheckBehavior.Extension != nil) {
		m.PrecheckBehaviorPrimitiveElement = &primitiveElement{Id: r.PrecheckBehavior.Id, Extension: r.PrecheckBehavior.Extension}
	}
	m.CardinalityBehavior = r.CardinalityBehavior
	if r.CardinalityBehavior != nil && (r.CardinalityBehavior.Id != nil || r.CardinalityBehavior.Extension != nil) {
		m.CardinalityBehaviorPrimitiveElement = &primitiveElement{Id: r.CardinalityBehavior.Id, Extension: r.CardinalityBehavior.Extension}
	}
	switch v := r.Definition.(type) {
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
	case Uri:
		m.DefinitionUri = &v
		if v.Id != nil || v.Extension != nil {
			m.DefinitionUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.DefinitionUri = v
		if v.Id != nil || v.Extension != nil {
			m.DefinitionUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Transform = r.Transform
	if r.Transform != nil && (r.Transform.Id != nil || r.Transform.Extension != nil) {
		m.TransformPrimitiveElement = &primitiveElement{Id: r.Transform.Id, Extension: r.Transform.Extension}
	}
	m.DynamicValue = r.DynamicValue
	m.Action = r.Action
	return m
}
func (r *PlanDefinitionAction) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionAction) unmarshalJSON(m jsonPlanDefinitionAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Prefix = m.Prefix
	if m.PrefixPrimitiveElement != nil {
		r.Prefix.Id = m.PrefixPrimitiveElement.Id
		r.Prefix.Extension = m.PrefixPrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.TextEquivalent = m.TextEquivalent
	if m.TextEquivalentPrimitiveElement != nil {
		r.TextEquivalent.Id = m.TextEquivalentPrimitiveElement.Id
		r.TextEquivalent.Extension = m.TextEquivalentPrimitiveElement.Extension
	}
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Reason = m.Reason
	r.Documentation = m.Documentation
	r.GoalId = m.GoalId
	for i, e := range m.GoalIdPrimitiveElement {
		if len(r.GoalId) > i {
			r.GoalId[i].Id = e.Id
			r.GoalId[i].Extension = e.Extension
		} else {
			r.GoalId = append(r.GoalId, Id{Id: e.Id, Extension: e.Extension})
		}
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
	r.Trigger = m.Trigger
	r.Condition = m.Condition
	r.Input = m.Input
	r.Output = m.Output
	r.RelatedAction = m.RelatedAction
	if m.TimingDateTime != nil || m.TimingDateTimePrimitiveElement != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDateTime
		if m.TimingDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.TimingDateTimePrimitiveElement.Id
			v.Extension = m.TimingDateTimePrimitiveElement.Extension
		}
		r.Timing = v
	}
	if m.TimingAge != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingAge
		r.Timing = v
	}
	if m.TimingPeriod != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingPeriod
		r.Timing = v
	}
	if m.TimingDuration != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDuration
		r.Timing = v
	}
	if m.TimingRange != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingRange
		r.Timing = v
	}
	if m.TimingTiming != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingTiming
		r.Timing = v
	}
	r.Participant = m.Participant
	r.Type = m.Type
	r.GroupingBehavior = m.GroupingBehavior
	if m.GroupingBehaviorPrimitiveElement != nil {
		r.GroupingBehavior.Id = m.GroupingBehaviorPrimitiveElement.Id
		r.GroupingBehavior.Extension = m.GroupingBehaviorPrimitiveElement.Extension
	}
	r.SelectionBehavior = m.SelectionBehavior
	if m.SelectionBehaviorPrimitiveElement != nil {
		r.SelectionBehavior.Id = m.SelectionBehaviorPrimitiveElement.Id
		r.SelectionBehavior.Extension = m.SelectionBehaviorPrimitiveElement.Extension
	}
	r.RequiredBehavior = m.RequiredBehavior
	if m.RequiredBehaviorPrimitiveElement != nil {
		r.RequiredBehavior.Id = m.RequiredBehaviorPrimitiveElement.Id
		r.RequiredBehavior.Extension = m.RequiredBehaviorPrimitiveElement.Extension
	}
	r.PrecheckBehavior = m.PrecheckBehavior
	if m.PrecheckBehaviorPrimitiveElement != nil {
		r.PrecheckBehavior.Id = m.PrecheckBehaviorPrimitiveElement.Id
		r.PrecheckBehavior.Extension = m.PrecheckBehaviorPrimitiveElement.Extension
	}
	r.CardinalityBehavior = m.CardinalityBehavior
	if m.CardinalityBehaviorPrimitiveElement != nil {
		r.CardinalityBehavior.Id = m.CardinalityBehaviorPrimitiveElement.Id
		r.CardinalityBehavior.Extension = m.CardinalityBehaviorPrimitiveElement.Extension
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
	if m.DefinitionUri != nil || m.DefinitionUriPrimitiveElement != nil {
		if r.Definition != nil {
			return fmt.Errorf("multiple values for field \"Definition\"")
		}
		v := m.DefinitionUri
		if m.DefinitionUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.DefinitionUriPrimitiveElement.Id
			v.Extension = m.DefinitionUriPrimitiveElement.Extension
		}
		r.Definition = v
	}
	r.Transform = m.Transform
	if m.TransformPrimitiveElement != nil {
		r.Transform.Id = m.TransformPrimitiveElement.Id
		r.Transform.Extension = m.TransformPrimitiveElement.Extension
	}
	r.DynamicValue = m.DynamicValue
	r.Action = m.Action
	return nil
}
func (r PlanDefinitionAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An expression that describes applicability criteria or start/stop conditions for the action.
type PlanDefinitionActionCondition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of condition.
	Kind Code
	// An expression that returns true or false, indicating whether the condition is satisfied.
	Expression *Expression
}
type jsonPlanDefinitionActionCondition struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Kind                 Code              `json:"kind,omitempty"`
	KindPrimitiveElement *primitiveElement `json:"_kind,omitempty"`
	Expression           *Expression       `json:"expression,omitempty"`
}

func (r PlanDefinitionActionCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionActionCondition) marshalJSON() jsonPlanDefinitionActionCondition {
	m := jsonPlanDefinitionActionCondition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Kind = r.Kind
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.Expression = r.Expression
	return m
}
func (r *PlanDefinitionActionCondition) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionActionCondition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionActionCondition) unmarshalJSON(m jsonPlanDefinitionActionCondition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Expression = m.Expression
	return nil
}
func (r PlanDefinitionActionCondition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type PlanDefinitionActionRelatedAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The element id of the related action.
	ActionId Id
	// The relationship of this action to the related action.
	Relationship Code
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	Offset isPlanDefinitionActionRelatedActionOffset
}
type isPlanDefinitionActionRelatedActionOffset interface {
	isPlanDefinitionActionRelatedActionOffset()
}

func (r Duration) isPlanDefinitionActionRelatedActionOffset() {}
func (r Range) isPlanDefinitionActionRelatedActionOffset()    {}

type jsonPlanDefinitionActionRelatedAction struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	ActionId                     Id                `json:"actionId,omitempty"`
	ActionIdPrimitiveElement     *primitiveElement `json:"_actionId,omitempty"`
	Relationship                 Code              `json:"relationship,omitempty"`
	RelationshipPrimitiveElement *primitiveElement `json:"_relationship,omitempty"`
	OffsetDuration               *Duration         `json:"offsetDuration,omitempty"`
	OffsetRange                  *Range            `json:"offsetRange,omitempty"`
}

func (r PlanDefinitionActionRelatedAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionActionRelatedAction) marshalJSON() jsonPlanDefinitionActionRelatedAction {
	m := jsonPlanDefinitionActionRelatedAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ActionId = r.ActionId
	if r.ActionId.Id != nil || r.ActionId.Extension != nil {
		m.ActionIdPrimitiveElement = &primitiveElement{Id: r.ActionId.Id, Extension: r.ActionId.Extension}
	}
	m.Relationship = r.Relationship
	if r.Relationship.Id != nil || r.Relationship.Extension != nil {
		m.RelationshipPrimitiveElement = &primitiveElement{Id: r.Relationship.Id, Extension: r.Relationship.Extension}
	}
	switch v := r.Offset.(type) {
	case Duration:
		m.OffsetDuration = &v
	case *Duration:
		m.OffsetDuration = v
	case Range:
		m.OffsetRange = &v
	case *Range:
		m.OffsetRange = v
	}
	return m
}
func (r *PlanDefinitionActionRelatedAction) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionActionRelatedAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionActionRelatedAction) unmarshalJSON(m jsonPlanDefinitionActionRelatedAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ActionId = m.ActionId
	if m.ActionIdPrimitiveElement != nil {
		r.ActionId.Id = m.ActionIdPrimitiveElement.Id
		r.ActionId.Extension = m.ActionIdPrimitiveElement.Extension
	}
	r.Relationship = m.Relationship
	if m.RelationshipPrimitiveElement != nil {
		r.Relationship.Id = m.RelationshipPrimitiveElement.Id
		r.Relationship.Extension = m.RelationshipPrimitiveElement.Extension
	}
	if m.OffsetDuration != nil {
		if r.Offset != nil {
			return fmt.Errorf("multiple values for field \"Offset\"")
		}
		v := m.OffsetDuration
		r.Offset = v
	}
	if m.OffsetRange != nil {
		if r.Offset != nil {
			return fmt.Errorf("multiple values for field \"Offset\"")
		}
		v := m.OffsetRange
		r.Offset = v
	}
	return nil
}
func (r PlanDefinitionActionRelatedAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who should participate in performing the action described.
type PlanDefinitionActionParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of participant in the action.
	Type Code
	// The role the participant should play in performing the described action.
	Role *CodeableConcept
}
type jsonPlanDefinitionActionParticipant struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
	Role                 *CodeableConcept  `json:"role,omitempty"`
}

func (r PlanDefinitionActionParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionActionParticipant) marshalJSON() jsonPlanDefinitionActionParticipant {
	m := jsonPlanDefinitionActionParticipant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Role = r.Role
	return m
}
func (r *PlanDefinitionActionParticipant) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionActionParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionActionParticipant) unmarshalJSON(m jsonPlanDefinitionActionParticipant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Role = m.Role
	return nil
}
func (r PlanDefinitionActionParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Customizations that should be applied to the statically defined resource. For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
type PlanDefinitionActionDynamicValue struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The path to the element to be customized. This is the path on the resource that will hold the result of the calculation defined by the expression. The specified path SHALL be a FHIRPath resolveable on the specified target type of the ActivityDefinition, and SHALL consist only of identifiers, constant indexers, and a restricted subset of functions. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	Path *String
	// An expression specifying the value of the customized element.
	Expression *Expression
}
type jsonPlanDefinitionActionDynamicValue struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Path                 *String           `json:"path,omitempty"`
	PathPrimitiveElement *primitiveElement `json:"_path,omitempty"`
	Expression           *Expression       `json:"expression,omitempty"`
}

func (r PlanDefinitionActionDynamicValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PlanDefinitionActionDynamicValue) marshalJSON() jsonPlanDefinitionActionDynamicValue {
	m := jsonPlanDefinitionActionDynamicValue{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Path = r.Path
	if r.Path != nil && (r.Path.Id != nil || r.Path.Extension != nil) {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	m.Expression = r.Expression
	return m
}
func (r *PlanDefinitionActionDynamicValue) UnmarshalJSON(b []byte) error {
	var m jsonPlanDefinitionActionDynamicValue
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PlanDefinitionActionDynamicValue) unmarshalJSON(m jsonPlanDefinitionActionDynamicValue) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.Expression = m.Expression
	return nil
}
func (r PlanDefinitionActionDynamicValue) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
