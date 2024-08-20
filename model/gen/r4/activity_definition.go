package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition struct {
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
	// An absolute URI that is used to identify this activity definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this activity definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the activity definition is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this activity definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the activity definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the activity definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active assets.
	Version *String
	// A natural language name identifying the activity definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the activity definition.
	Title *String
	// An explanatory or alternate title for the activity definition giving additional information about its content.
	Subtitle *String
	// The status of this activity definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this activity definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// A code or group definition that describes the intended subject of the activity being defined.
	Subject isActivityDefinitionSubject
	// The date  (and optionally time) when the activity definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the activity definition changes.
	Date *DateTime
	// The name of the organization or individual that published the activity definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the activity definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate activity definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the activity definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this activity definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// A detailed description of how the activity definition is used from a clinical perspective.
	Usage *String
	// A copyright statement relating to the activity definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the activity definition.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the activity definition content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the activity. Topics provide a high-level categorization of the activity that can be useful for filtering and searching.
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
	// A reference to a Library resource containing any formal logic used by the activity definition.
	Library []Canonical
	// A description of the kind of resource the activity definition is representing. For example, a MedicationRequest, a ServiceRequest, or a CommunicationRequest. Typically, but not always, this is a Request resource.
	Kind *Code
	// A profile to which the target of the activity definition is expected to conform.
	Profile *Canonical
	// Detailed description of the type of activity; e.g. What lab test, what procedure, what kind of encounter.
	Code *CodeableConcept
	// Indicates the level of authority/intentionality associated with the activity and where the request should fit into the workflow chain.
	Intent *Code
	// Indicates how quickly the activity  should be addressed with respect to other requests.
	Priority *Code
	// Set this to true if the definition is to indicate that a particular activity should NOT be performed. If true, this element should be interpreted to reinforce a negative coding. For example NPO as a code with a doNotPerform of true would still indicate to NOT perform the action.
	DoNotPerform *Boolean
	// The period, timing or frequency upon which the described activity is to occur.
	Timing isActivityDefinitionTiming
	// Identifies the facility where the activity will occur; e.g. home, hospital, specific clinic, etc.
	Location *Reference
	// Indicates who should participate in performing the action described.
	Participant []ActivityDefinitionParticipant
	// Identifies the food, drug or other product being consumed or supplied in the activity.
	Product isActivityDefinitionProduct
	// Identifies the quantity expected to be consumed at once (per dose, per meal, etc.).
	Quantity *Quantity
	// Provides detailed dosage instructions in the same way that they are described for MedicationRequest resources.
	Dosage []Dosage
	// Indicates the sites on the subject's body where the procedure should be performed (I.e. the target sites).
	BodySite []CodeableConcept
	// Defines specimen requirements for the action to be performed, such as required specimens for a lab test.
	SpecimenRequirement []Reference
	// Defines observation requirements for the action to be performed, such as body weight or surface area.
	ObservationRequirement []Reference
	// Defines the observations that are expected to be produced by the action.
	ObservationResultRequirement []Reference
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	Transform *Canonical
	// Dynamic values that will be evaluated to produce values for elements of the resulting resource. For example, if the dosage of a medication must be computed based on the patient's weight, a dynamic value would be used to specify an expression that calculated the weight, and the path on the request resource that would contain the result.
	DynamicValue []ActivityDefinitionDynamicValue
}

func (r ActivityDefinition) ResourceType() string {
	return "ActivityDefinition"
}
func (r ActivityDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isActivityDefinitionSubject interface {
	isActivityDefinitionSubject()
}

func (r CodeableConcept) isActivityDefinitionSubject() {}
func (r Reference) isActivityDefinitionSubject()       {}

type isActivityDefinitionTiming interface {
	isActivityDefinitionTiming()
}

func (r Timing) isActivityDefinitionTiming()   {}
func (r DateTime) isActivityDefinitionTiming() {}
func (r Age) isActivityDefinitionTiming()      {}
func (r Period) isActivityDefinitionTiming()   {}
func (r Range) isActivityDefinitionTiming()    {}
func (r Duration) isActivityDefinitionTiming() {}

type isActivityDefinitionProduct interface {
	isActivityDefinitionProduct()
}

func (r Reference) isActivityDefinitionProduct()       {}
func (r CodeableConcept) isActivityDefinitionProduct() {}

type jsonActivityDefinition struct {
	ResourceType                   string                           `json:"resourceType"`
	Id                             *Id                              `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                `json:"_id,omitempty"`
	Meta                           *Meta                            `json:"meta,omitempty"`
	ImplicitRules                  *Uri                             `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                `json:"_implicitRules,omitempty"`
	Language                       *Code                            `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                `json:"_language,omitempty"`
	Text                           *Narrative                       `json:"text,omitempty"`
	Contained                      []ContainedResource              `json:"contained,omitempty"`
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
	Subtitle                       *String                          `json:"subtitle,omitempty"`
	SubtitlePrimitiveElement       *primitiveElement                `json:"_subtitle,omitempty"`
	Status                         Code                             `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                `json:"_status,omitempty"`
	Experimental                   *Boolean                         `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement   *primitiveElement                `json:"_experimental,omitempty"`
	SubjectCodeableConcept         *CodeableConcept                 `json:"subjectCodeableConcept,omitempty"`
	SubjectReference               *Reference                       `json:"subjectReference,omitempty"`
	Date                           *DateTime                        `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement                `json:"_date,omitempty"`
	Publisher                      *String                          `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement                `json:"_publisher,omitempty"`
	Contact                        []ContactDetail                  `json:"contact,omitempty"`
	Description                    *Markdown                        `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement                `json:"_description,omitempty"`
	UseContext                     []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                        *Markdown                        `json:"purpose,omitempty"`
	PurposePrimitiveElement        *primitiveElement                `json:"_purpose,omitempty"`
	Usage                          *String                          `json:"usage,omitempty"`
	UsagePrimitiveElement          *primitiveElement                `json:"_usage,omitempty"`
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
	Library                        []Canonical                      `json:"library,omitempty"`
	LibraryPrimitiveElement        []*primitiveElement              `json:"_library,omitempty"`
	Kind                           *Code                            `json:"kind,omitempty"`
	KindPrimitiveElement           *primitiveElement                `json:"_kind,omitempty"`
	Profile                        *Canonical                       `json:"profile,omitempty"`
	ProfilePrimitiveElement        *primitiveElement                `json:"_profile,omitempty"`
	Code                           *CodeableConcept                 `json:"code,omitempty"`
	Intent                         *Code                            `json:"intent,omitempty"`
	IntentPrimitiveElement         *primitiveElement                `json:"_intent,omitempty"`
	Priority                       *Code                            `json:"priority,omitempty"`
	PriorityPrimitiveElement       *primitiveElement                `json:"_priority,omitempty"`
	DoNotPerform                   *Boolean                         `json:"doNotPerform,omitempty"`
	DoNotPerformPrimitiveElement   *primitiveElement                `json:"_doNotPerform,omitempty"`
	TimingTiming                   *Timing                          `json:"timingTiming,omitempty"`
	TimingDateTime                 *DateTime                        `json:"timingDateTime,omitempty"`
	TimingDateTimePrimitiveElement *primitiveElement                `json:"_timingDateTime,omitempty"`
	TimingAge                      *Age                             `json:"timingAge,omitempty"`
	TimingPeriod                   *Period                          `json:"timingPeriod,omitempty"`
	TimingRange                    *Range                           `json:"timingRange,omitempty"`
	TimingDuration                 *Duration                        `json:"timingDuration,omitempty"`
	Location                       *Reference                       `json:"location,omitempty"`
	Participant                    []ActivityDefinitionParticipant  `json:"participant,omitempty"`
	ProductReference               *Reference                       `json:"productReference,omitempty"`
	ProductCodeableConcept         *CodeableConcept                 `json:"productCodeableConcept,omitempty"`
	Quantity                       *Quantity                        `json:"quantity,omitempty"`
	Dosage                         []Dosage                         `json:"dosage,omitempty"`
	BodySite                       []CodeableConcept                `json:"bodySite,omitempty"`
	SpecimenRequirement            []Reference                      `json:"specimenRequirement,omitempty"`
	ObservationRequirement         []Reference                      `json:"observationRequirement,omitempty"`
	ObservationResultRequirement   []Reference                      `json:"observationResultRequirement,omitempty"`
	Transform                      *Canonical                       `json:"transform,omitempty"`
	TransformPrimitiveElement      *primitiveElement                `json:"_transform,omitempty"`
	DynamicValue                   []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`
}

func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ActivityDefinition) marshalJSON() jsonActivityDefinition {
	m := jsonActivityDefinition{}
	m.ResourceType = "ActivityDefinition"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
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
	m.Kind = r.Kind
	if r.Kind != nil && (r.Kind.Id != nil || r.Kind.Extension != nil) {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.Profile = r.Profile
	if r.Profile != nil && (r.Profile.Id != nil || r.Profile.Extension != nil) {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	m.Code = r.Code
	m.Intent = r.Intent
	if r.Intent != nil && (r.Intent.Id != nil || r.Intent.Extension != nil) {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.DoNotPerform = r.DoNotPerform
	if r.DoNotPerform != nil && (r.DoNotPerform.Id != nil || r.DoNotPerform.Extension != nil) {
		m.DoNotPerformPrimitiveElement = &primitiveElement{Id: r.DoNotPerform.Id, Extension: r.DoNotPerform.Extension}
	}
	switch v := r.Timing.(type) {
	case Timing:
		m.TimingTiming = &v
	case *Timing:
		m.TimingTiming = v
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
	case Range:
		m.TimingRange = &v
	case *Range:
		m.TimingRange = v
	case Duration:
		m.TimingDuration = &v
	case *Duration:
		m.TimingDuration = v
	}
	m.Location = r.Location
	m.Participant = r.Participant
	switch v := r.Product.(type) {
	case Reference:
		m.ProductReference = &v
	case *Reference:
		m.ProductReference = v
	case CodeableConcept:
		m.ProductCodeableConcept = &v
	case *CodeableConcept:
		m.ProductCodeableConcept = v
	}
	m.Quantity = r.Quantity
	m.Dosage = r.Dosage
	m.BodySite = r.BodySite
	m.SpecimenRequirement = r.SpecimenRequirement
	m.ObservationRequirement = r.ObservationRequirement
	m.ObservationResultRequirement = r.ObservationResultRequirement
	m.Transform = r.Transform
	if r.Transform != nil && (r.Transform.Id != nil || r.Transform.Extension != nil) {
		m.TransformPrimitiveElement = &primitiveElement{Id: r.Transform.Id, Extension: r.Transform.Extension}
	}
	m.DynamicValue = r.DynamicValue
	return m
}
func (r *ActivityDefinition) UnmarshalJSON(b []byte) error {
	var m jsonActivityDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ActivityDefinition) unmarshalJSON(m jsonActivityDefinition) error {
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
		r.Contained = append(r.Contained, v.Resource)
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
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Intent = m.Intent
	if m.IntentPrimitiveElement != nil {
		r.Intent.Id = m.IntentPrimitiveElement.Id
		r.Intent.Extension = m.IntentPrimitiveElement.Extension
	}
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.DoNotPerform = m.DoNotPerform
	if m.DoNotPerformPrimitiveElement != nil {
		r.DoNotPerform.Id = m.DoNotPerformPrimitiveElement.Id
		r.DoNotPerform.Extension = m.DoNotPerformPrimitiveElement.Extension
	}
	if m.TimingTiming != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingTiming
		r.Timing = v
	}
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
	if m.TimingRange != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingRange
		r.Timing = v
	}
	if m.TimingDuration != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDuration
		r.Timing = v
	}
	r.Location = m.Location
	r.Participant = m.Participant
	if m.ProductReference != nil {
		if r.Product != nil {
			return fmt.Errorf("multiple values for field \"Product\"")
		}
		v := m.ProductReference
		r.Product = v
	}
	if m.ProductCodeableConcept != nil {
		if r.Product != nil {
			return fmt.Errorf("multiple values for field \"Product\"")
		}
		v := m.ProductCodeableConcept
		r.Product = v
	}
	r.Quantity = m.Quantity
	r.Dosage = m.Dosage
	r.BodySite = m.BodySite
	r.SpecimenRequirement = m.SpecimenRequirement
	r.ObservationRequirement = m.ObservationRequirement
	r.ObservationResultRequirement = m.ObservationResultRequirement
	r.Transform = m.Transform
	if m.TransformPrimitiveElement != nil {
		r.Transform.Id = m.TransformPrimitiveElement.Id
		r.Transform.Extension = m.TransformPrimitiveElement.Extension
	}
	r.DynamicValue = m.DynamicValue
	return nil
}
func (r ActivityDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who should participate in performing the action described.
type ActivityDefinitionParticipant struct {
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
type jsonActivityDefinitionParticipant struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
	Role                 *CodeableConcept  `json:"role,omitempty"`
}

func (r ActivityDefinitionParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ActivityDefinitionParticipant) marshalJSON() jsonActivityDefinitionParticipant {
	m := jsonActivityDefinitionParticipant{}
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
func (r *ActivityDefinitionParticipant) UnmarshalJSON(b []byte) error {
	var m jsonActivityDefinitionParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ActivityDefinitionParticipant) unmarshalJSON(m jsonActivityDefinitionParticipant) error {
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
func (r ActivityDefinitionParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Dynamic values that will be evaluated to produce values for elements of the resulting resource. For example, if the dosage of a medication must be computed based on the patient's weight, a dynamic value would be used to specify an expression that calculated the weight, and the path on the request resource that would contain the result.
type ActivityDefinitionDynamicValue struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The path to the element to be customized. This is the path on the resource that will hold the result of the calculation defined by the expression. The specified path SHALL be a FHIRPath resolveable on the specified target type of the ActivityDefinition, and SHALL consist only of identifiers, constant indexers, and a restricted subset of functions. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	Path String
	// An expression specifying the value of the customized element.
	Expression Expression
}
type jsonActivityDefinitionDynamicValue struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Path                 String            `json:"path,omitempty"`
	PathPrimitiveElement *primitiveElement `json:"_path,omitempty"`
	Expression           Expression        `json:"expression,omitempty"`
}

func (r ActivityDefinitionDynamicValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ActivityDefinitionDynamicValue) marshalJSON() jsonActivityDefinitionDynamicValue {
	m := jsonActivityDefinitionDynamicValue{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Path = r.Path
	if r.Path.Id != nil || r.Path.Extension != nil {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	m.Expression = r.Expression
	return m
}
func (r *ActivityDefinitionDynamicValue) UnmarshalJSON(b []byte) error {
	var m jsonActivityDefinitionDynamicValue
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ActivityDefinitionDynamicValue) unmarshalJSON(m jsonActivityDefinitionDynamicValue) error {
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
func (r ActivityDefinitionDynamicValue) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
