package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
//
// To support structured, hierarchical registration of data gathered using digital forms and other questionnaires.  Questionnaires provide greater control over presentation and allow capture of data in a domain-independent way (i.e. capturing information that would otherwise require multiple distinct types of resources).
type Questionnaire struct {
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
	// An absolute URI that is used to identify this questionnaire when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this questionnaire is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the questionnaire is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this questionnaire when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the questionnaire when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the questionnaire author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the questionnaire. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the questionnaire.
	Title *String
	// The URL of a Questionnaire that this Questionnaire is based on.
	DerivedFrom []Canonical
	// The status of this questionnaire. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this questionnaire is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The types of subjects that can be the subject of responses created for the questionnaire.
	SubjectType []Code
	// The date  (and optionally time) when the questionnaire was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the questionnaire changes.
	Date *DateTime
	// The name of the organization or individual that published the questionnaire.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the questionnaire from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate questionnaire instances.
	UseContext []UsageContext
	// A legal or geographic region in which the questionnaire is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this questionnaire is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the questionnaire and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the questionnaire.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the questionnaire content was or is planned to be in active use.
	EffectivePeriod *Period
	// An identifier for this question or group of questions in a particular terminology such as LOINC.
	Code []Coding
	// A particular question, question grouping or display text that is part of the questionnaire.
	Item []QuestionnaireItem
}

func (r Questionnaire) ResourceType() string {
	return "Questionnaire"
}
func (r Questionnaire) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonQuestionnaire struct {
	ResourceType                   string              `json:"resourceType"`
	Id                             *Id                 `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement   `json:"_id,omitempty"`
	Meta                           *Meta               `json:"meta,omitempty"`
	ImplicitRules                  *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                       *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement   `json:"_language,omitempty"`
	Text                           *Narrative          `json:"text,omitempty"`
	Contained                      []containedResource `json:"contained,omitempty"`
	Extension                      []Extension         `json:"extension,omitempty"`
	ModifierExtension              []Extension         `json:"modifierExtension,omitempty"`
	Url                            *Uri                `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement   `json:"_url,omitempty"`
	Identifier                     []Identifier        `json:"identifier,omitempty"`
	Version                        *String             `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement   `json:"_version,omitempty"`
	Name                           *String             `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement   `json:"_name,omitempty"`
	Title                          *String             `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement   `json:"_title,omitempty"`
	DerivedFrom                    []Canonical         `json:"derivedFrom,omitempty"`
	DerivedFromPrimitiveElement    []*primitiveElement `json:"_derivedFrom,omitempty"`
	Status                         Code                `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement   `json:"_status,omitempty"`
	Experimental                   *Boolean            `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement   *primitiveElement   `json:"_experimental,omitempty"`
	SubjectType                    []Code              `json:"subjectType,omitempty"`
	SubjectTypePrimitiveElement    []*primitiveElement `json:"_subjectType,omitempty"`
	Date                           *DateTime           `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement   `json:"_date,omitempty"`
	Publisher                      *String             `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement   `json:"_publisher,omitempty"`
	Contact                        []ContactDetail     `json:"contact,omitempty"`
	Description                    *Markdown           `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement   `json:"_description,omitempty"`
	UseContext                     []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                        *Markdown           `json:"purpose,omitempty"`
	PurposePrimitiveElement        *primitiveElement   `json:"_purpose,omitempty"`
	Copyright                      *Markdown           `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement   `json:"_copyright,omitempty"`
	ApprovalDate                   *Date               `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement   `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date               `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement   `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period             `json:"effectivePeriod,omitempty"`
	Code                           []Coding            `json:"code,omitempty"`
	Item                           []QuestionnaireItem `json:"item,omitempty"`
}

func (r Questionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Questionnaire) marshalJSON() jsonQuestionnaire {
	m := jsonQuestionnaire{}
	m.ResourceType = "Questionnaire"
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
	m.DerivedFrom = r.DerivedFrom
	anyDerivedFromIdOrExtension := false
	for _, e := range r.DerivedFrom {
		if e.Id != nil || e.Extension != nil {
			anyDerivedFromIdOrExtension = true
			break
		}
	}
	if anyDerivedFromIdOrExtension {
		m.DerivedFromPrimitiveElement = make([]*primitiveElement, 0, len(r.DerivedFrom))
		for _, e := range r.DerivedFrom {
			if e.Id != nil || e.Extension != nil {
				m.DerivedFromPrimitiveElement = append(m.DerivedFromPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DerivedFromPrimitiveElement = append(m.DerivedFromPrimitiveElement, nil)
			}
		}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Experimental = r.Experimental
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	m.SubjectType = r.SubjectType
	anySubjectTypeIdOrExtension := false
	for _, e := range r.SubjectType {
		if e.Id != nil || e.Extension != nil {
			anySubjectTypeIdOrExtension = true
			break
		}
	}
	if anySubjectTypeIdOrExtension {
		m.SubjectTypePrimitiveElement = make([]*primitiveElement, 0, len(r.SubjectType))
		for _, e := range r.SubjectType {
			if e.Id != nil || e.Extension != nil {
				m.SubjectTypePrimitiveElement = append(m.SubjectTypePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SubjectTypePrimitiveElement = append(m.SubjectTypePrimitiveElement, nil)
			}
		}
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
	m.Code = r.Code
	m.Item = r.Item
	return m
}
func (r *Questionnaire) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaire
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Questionnaire) unmarshalJSON(m jsonQuestionnaire) error {
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
	r.DerivedFrom = m.DerivedFrom
	for i, e := range m.DerivedFromPrimitiveElement {
		if len(r.DerivedFrom) > i {
			r.DerivedFrom[i].Id = e.Id
			r.DerivedFrom[i].Extension = e.Extension
		} else {
			r.DerivedFrom = append(r.DerivedFrom, Canonical{Id: e.Id, Extension: e.Extension})
		}
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
	r.SubjectType = m.SubjectType
	for i, e := range m.SubjectTypePrimitiveElement {
		if len(r.SubjectType) > i {
			r.SubjectType[i].Id = e.Id
			r.SubjectType[i].Extension = e.Extension
		} else {
			r.SubjectType = append(r.SubjectType, Code{Id: e.Id, Extension: e.Extension})
		}
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
	r.Code = m.Code
	r.Item = m.Item
	return nil
}
func (r Questionnaire) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A particular question, question grouping or display text that is part of the questionnaire.
type QuestionnaireItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An identifier that is unique within the Questionnaire allowing linkage to the equivalent item in a QuestionnaireResponse resource.
	LinkId String
	// This element is a URI that refers to an [ElementDefinition](elementdefinition.html) that provides information about this item, including information that might otherwise be included in the instance of the Questionnaire resource. A detailed description of the construction of the URI is shown in Comments, below. If this element is present then the following element values MAY be derived from the Element Definition if the corresponding elements of this Questionnaire resource instance have no value:
	//
	// * code (ElementDefinition.code)
	// * type (ElementDefinition.type)
	// * required (ElementDefinition.min)
	// * repeats (ElementDefinition.max)
	// * maxLength (ElementDefinition.maxLength)
	// * answerValueSet (ElementDefinition.binding)
	// * options (ElementDefinition.binding).
	Definition *Uri
	// A terminology code that corresponds to this group or question (e.g. a code from LOINC, which defines many questions and answers).
	Code []Coding
	// A short label for a particular group, question or set of display text within the questionnaire used for reference by the individual completing the questionnaire.
	Prefix *String
	// The name of a section, the text of a question or text content for a display item.
	Text *String
	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured (string, integer, coded choice, etc.).
	Type Code
	// A constraint indicating that this item should only be enabled (displayed/allow answers to be captured) when the specified condition is true.
	EnableWhen []QuestionnaireItemEnableWhen
	// Controls how multiple enableWhen values are interpreted -  whether all or any must be true.
	EnableBehavior *Code
	// An indication, if true, that the item must be present in a "completed" QuestionnaireResponse.  If false, the item may be skipped when answering the questionnaire.
	Required *Boolean
	// An indication, if true, that the item may occur multiple times in the response, collecting multiple answers for questions or multiple sets of answers for groups.
	Repeats *Boolean
	// An indication, when true, that the value cannot be changed by a human respondent to the Questionnaire.
	ReadOnly *Boolean
	// The maximum number of characters that are permitted in the answer to be considered a "valid" QuestionnaireResponse.
	MaxLength *Integer
	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question.
	AnswerValueSet *Canonical
	// One of the permitted answers for a "choice" or "open-choice" question.
	AnswerOption []QuestionnaireItemAnswerOption
	// One or more values that should be pre-populated in the answer when initially rendering the questionnaire for user input.
	Initial []QuestionnaireItemInitial
	// Text, questions and other groups to be nested beneath a question or group.
	Item []QuestionnaireItem
}
type jsonQuestionnaireItem struct {
	Id                             *string                         `json:"id,omitempty"`
	Extension                      []Extension                     `json:"extension,omitempty"`
	ModifierExtension              []Extension                     `json:"modifierExtension,omitempty"`
	LinkId                         String                          `json:"linkId,omitempty"`
	LinkIdPrimitiveElement         *primitiveElement               `json:"_linkId,omitempty"`
	Definition                     *Uri                            `json:"definition,omitempty"`
	DefinitionPrimitiveElement     *primitiveElement               `json:"_definition,omitempty"`
	Code                           []Coding                        `json:"code,omitempty"`
	Prefix                         *String                         `json:"prefix,omitempty"`
	PrefixPrimitiveElement         *primitiveElement               `json:"_prefix,omitempty"`
	Text                           *String                         `json:"text,omitempty"`
	TextPrimitiveElement           *primitiveElement               `json:"_text,omitempty"`
	Type                           Code                            `json:"type,omitempty"`
	TypePrimitiveElement           *primitiveElement               `json:"_type,omitempty"`
	EnableWhen                     []QuestionnaireItemEnableWhen   `json:"enableWhen,omitempty"`
	EnableBehavior                 *Code                           `json:"enableBehavior,omitempty"`
	EnableBehaviorPrimitiveElement *primitiveElement               `json:"_enableBehavior,omitempty"`
	Required                       *Boolean                        `json:"required,omitempty"`
	RequiredPrimitiveElement       *primitiveElement               `json:"_required,omitempty"`
	Repeats                        *Boolean                        `json:"repeats,omitempty"`
	RepeatsPrimitiveElement        *primitiveElement               `json:"_repeats,omitempty"`
	ReadOnly                       *Boolean                        `json:"readOnly,omitempty"`
	ReadOnlyPrimitiveElement       *primitiveElement               `json:"_readOnly,omitempty"`
	MaxLength                      *Integer                        `json:"maxLength,omitempty"`
	MaxLengthPrimitiveElement      *primitiveElement               `json:"_maxLength,omitempty"`
	AnswerValueSet                 *Canonical                      `json:"answerValueSet,omitempty"`
	AnswerValueSetPrimitiveElement *primitiveElement               `json:"_answerValueSet,omitempty"`
	AnswerOption                   []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`
	Initial                        []QuestionnaireItemInitial      `json:"initial,omitempty"`
	Item                           []QuestionnaireItem             `json:"item,omitempty"`
}

func (r QuestionnaireItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireItem) marshalJSON() jsonQuestionnaireItem {
	m := jsonQuestionnaireItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.LinkId = r.LinkId
	if r.LinkId.Id != nil || r.LinkId.Extension != nil {
		m.LinkIdPrimitiveElement = &primitiveElement{Id: r.LinkId.Id, Extension: r.LinkId.Extension}
	}
	m.Definition = r.Definition
	if r.Definition != nil && (r.Definition.Id != nil || r.Definition.Extension != nil) {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	m.Code = r.Code
	m.Prefix = r.Prefix
	if r.Prefix != nil && (r.Prefix.Id != nil || r.Prefix.Extension != nil) {
		m.PrefixPrimitiveElement = &primitiveElement{Id: r.Prefix.Id, Extension: r.Prefix.Extension}
	}
	m.Text = r.Text
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.EnableWhen = r.EnableWhen
	m.EnableBehavior = r.EnableBehavior
	if r.EnableBehavior != nil && (r.EnableBehavior.Id != nil || r.EnableBehavior.Extension != nil) {
		m.EnableBehaviorPrimitiveElement = &primitiveElement{Id: r.EnableBehavior.Id, Extension: r.EnableBehavior.Extension}
	}
	m.Required = r.Required
	if r.Required != nil && (r.Required.Id != nil || r.Required.Extension != nil) {
		m.RequiredPrimitiveElement = &primitiveElement{Id: r.Required.Id, Extension: r.Required.Extension}
	}
	m.Repeats = r.Repeats
	if r.Repeats != nil && (r.Repeats.Id != nil || r.Repeats.Extension != nil) {
		m.RepeatsPrimitiveElement = &primitiveElement{Id: r.Repeats.Id, Extension: r.Repeats.Extension}
	}
	m.ReadOnly = r.ReadOnly
	if r.ReadOnly != nil && (r.ReadOnly.Id != nil || r.ReadOnly.Extension != nil) {
		m.ReadOnlyPrimitiveElement = &primitiveElement{Id: r.ReadOnly.Id, Extension: r.ReadOnly.Extension}
	}
	m.MaxLength = r.MaxLength
	if r.MaxLength != nil && (r.MaxLength.Id != nil || r.MaxLength.Extension != nil) {
		m.MaxLengthPrimitiveElement = &primitiveElement{Id: r.MaxLength.Id, Extension: r.MaxLength.Extension}
	}
	m.AnswerValueSet = r.AnswerValueSet
	if r.AnswerValueSet != nil && (r.AnswerValueSet.Id != nil || r.AnswerValueSet.Extension != nil) {
		m.AnswerValueSetPrimitiveElement = &primitiveElement{Id: r.AnswerValueSet.Id, Extension: r.AnswerValueSet.Extension}
	}
	m.AnswerOption = r.AnswerOption
	m.Initial = r.Initial
	m.Item = r.Item
	return m
}
func (r *QuestionnaireItem) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireItem) unmarshalJSON(m jsonQuestionnaireItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.LinkId = m.LinkId
	if m.LinkIdPrimitiveElement != nil {
		r.LinkId.Id = m.LinkIdPrimitiveElement.Id
		r.LinkId.Extension = m.LinkIdPrimitiveElement.Extension
	}
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Prefix = m.Prefix
	if m.PrefixPrimitiveElement != nil {
		r.Prefix.Id = m.PrefixPrimitiveElement.Id
		r.Prefix.Extension = m.PrefixPrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.EnableWhen = m.EnableWhen
	r.EnableBehavior = m.EnableBehavior
	if m.EnableBehaviorPrimitiveElement != nil {
		r.EnableBehavior.Id = m.EnableBehaviorPrimitiveElement.Id
		r.EnableBehavior.Extension = m.EnableBehaviorPrimitiveElement.Extension
	}
	r.Required = m.Required
	if m.RequiredPrimitiveElement != nil {
		r.Required.Id = m.RequiredPrimitiveElement.Id
		r.Required.Extension = m.RequiredPrimitiveElement.Extension
	}
	r.Repeats = m.Repeats
	if m.RepeatsPrimitiveElement != nil {
		r.Repeats.Id = m.RepeatsPrimitiveElement.Id
		r.Repeats.Extension = m.RepeatsPrimitiveElement.Extension
	}
	r.ReadOnly = m.ReadOnly
	if m.ReadOnlyPrimitiveElement != nil {
		r.ReadOnly.Id = m.ReadOnlyPrimitiveElement.Id
		r.ReadOnly.Extension = m.ReadOnlyPrimitiveElement.Extension
	}
	r.MaxLength = m.MaxLength
	if m.MaxLengthPrimitiveElement != nil {
		r.MaxLength.Id = m.MaxLengthPrimitiveElement.Id
		r.MaxLength.Extension = m.MaxLengthPrimitiveElement.Extension
	}
	r.AnswerValueSet = m.AnswerValueSet
	if m.AnswerValueSetPrimitiveElement != nil {
		r.AnswerValueSet.Id = m.AnswerValueSetPrimitiveElement.Id
		r.AnswerValueSet.Extension = m.AnswerValueSetPrimitiveElement.Extension
	}
	r.AnswerOption = m.AnswerOption
	r.Initial = m.Initial
	r.Item = m.Item
	return nil
}
func (r QuestionnaireItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A constraint indicating that this item should only be enabled (displayed/allow answers to be captured) when the specified condition is true.
type QuestionnaireItemEnableWhen struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The linkId for the question whose answer (or lack of answer) governs whether this item is enabled.
	Question String
	// Specifies the criteria by which the question is enabled.
	Operator Code
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	Answer isQuestionnaireItemEnableWhenAnswer
}
type isQuestionnaireItemEnableWhenAnswer interface {
	isQuestionnaireItemEnableWhenAnswer()
}

func (r Boolean) isQuestionnaireItemEnableWhenAnswer()   {}
func (r Decimal) isQuestionnaireItemEnableWhenAnswer()   {}
func (r Integer) isQuestionnaireItemEnableWhenAnswer()   {}
func (r Date) isQuestionnaireItemEnableWhenAnswer()      {}
func (r DateTime) isQuestionnaireItemEnableWhenAnswer()  {}
func (r Time) isQuestionnaireItemEnableWhenAnswer()      {}
func (r String) isQuestionnaireItemEnableWhenAnswer()    {}
func (r Coding) isQuestionnaireItemEnableWhenAnswer()    {}
func (r Quantity) isQuestionnaireItemEnableWhenAnswer()  {}
func (r Reference) isQuestionnaireItemEnableWhenAnswer() {}

type jsonQuestionnaireItemEnableWhen struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	Question                       String            `json:"question,omitempty"`
	QuestionPrimitiveElement       *primitiveElement `json:"_question,omitempty"`
	Operator                       Code              `json:"operator,omitempty"`
	OperatorPrimitiveElement       *primitiveElement `json:"_operator,omitempty"`
	AnswerBoolean                  *Boolean          `json:"answerBoolean,omitempty"`
	AnswerBooleanPrimitiveElement  *primitiveElement `json:"_answerBoolean,omitempty"`
	AnswerDecimal                  *Decimal          `json:"answerDecimal,omitempty"`
	AnswerDecimalPrimitiveElement  *primitiveElement `json:"_answerDecimal,omitempty"`
	AnswerInteger                  *Integer          `json:"answerInteger,omitempty"`
	AnswerIntegerPrimitiveElement  *primitiveElement `json:"_answerInteger,omitempty"`
	AnswerDate                     *Date             `json:"answerDate,omitempty"`
	AnswerDatePrimitiveElement     *primitiveElement `json:"_answerDate,omitempty"`
	AnswerDateTime                 *DateTime         `json:"answerDateTime,omitempty"`
	AnswerDateTimePrimitiveElement *primitiveElement `json:"_answerDateTime,omitempty"`
	AnswerTime                     *Time             `json:"answerTime,omitempty"`
	AnswerTimePrimitiveElement     *primitiveElement `json:"_answerTime,omitempty"`
	AnswerString                   *String           `json:"answerString,omitempty"`
	AnswerStringPrimitiveElement   *primitiveElement `json:"_answerString,omitempty"`
	AnswerCoding                   *Coding           `json:"answerCoding,omitempty"`
	AnswerQuantity                 *Quantity         `json:"answerQuantity,omitempty"`
	AnswerReference                *Reference        `json:"answerReference,omitempty"`
}

func (r QuestionnaireItemEnableWhen) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireItemEnableWhen) marshalJSON() jsonQuestionnaireItemEnableWhen {
	m := jsonQuestionnaireItemEnableWhen{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Question = r.Question
	if r.Question.Id != nil || r.Question.Extension != nil {
		m.QuestionPrimitiveElement = &primitiveElement{Id: r.Question.Id, Extension: r.Question.Extension}
	}
	m.Operator = r.Operator
	if r.Operator.Id != nil || r.Operator.Extension != nil {
		m.OperatorPrimitiveElement = &primitiveElement{Id: r.Operator.Id, Extension: r.Operator.Extension}
	}
	switch v := r.Answer.(type) {
	case Boolean:
		m.AnswerBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.AnswerBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Decimal:
		m.AnswerDecimal = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		m.AnswerDecimal = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.AnswerInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.AnswerInteger = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		m.AnswerDate = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.AnswerDate = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		m.AnswerDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.AnswerDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		m.AnswerTime = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		m.AnswerTime = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.AnswerString = &v
		if v.Id != nil || v.Extension != nil {
			m.AnswerStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.AnswerString = v
		if v.Id != nil || v.Extension != nil {
			m.AnswerStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Coding:
		m.AnswerCoding = &v
	case *Coding:
		m.AnswerCoding = v
	case Quantity:
		m.AnswerQuantity = &v
	case *Quantity:
		m.AnswerQuantity = v
	case Reference:
		m.AnswerReference = &v
	case *Reference:
		m.AnswerReference = v
	}
	return m
}
func (r *QuestionnaireItemEnableWhen) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireItemEnableWhen
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireItemEnableWhen) unmarshalJSON(m jsonQuestionnaireItemEnableWhen) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Question = m.Question
	if m.QuestionPrimitiveElement != nil {
		r.Question.Id = m.QuestionPrimitiveElement.Id
		r.Question.Extension = m.QuestionPrimitiveElement.Extension
	}
	r.Operator = m.Operator
	if m.OperatorPrimitiveElement != nil {
		r.Operator.Id = m.OperatorPrimitiveElement.Id
		r.Operator.Extension = m.OperatorPrimitiveElement.Extension
	}
	if m.AnswerBoolean != nil || m.AnswerBooleanPrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerBoolean
		if m.AnswerBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.AnswerBooleanPrimitiveElement.Id
			v.Extension = m.AnswerBooleanPrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerDecimal != nil || m.AnswerDecimalPrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerDecimal
		if m.AnswerDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.AnswerDecimalPrimitiveElement.Id
			v.Extension = m.AnswerDecimalPrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerInteger != nil || m.AnswerIntegerPrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerInteger
		if m.AnswerIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.AnswerIntegerPrimitiveElement.Id
			v.Extension = m.AnswerIntegerPrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerDate != nil || m.AnswerDatePrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerDate
		if m.AnswerDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.AnswerDatePrimitiveElement.Id
			v.Extension = m.AnswerDatePrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerDateTime != nil || m.AnswerDateTimePrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerDateTime
		if m.AnswerDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.AnswerDateTimePrimitiveElement.Id
			v.Extension = m.AnswerDateTimePrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerTime != nil || m.AnswerTimePrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerTime
		if m.AnswerTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.AnswerTimePrimitiveElement.Id
			v.Extension = m.AnswerTimePrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerString != nil || m.AnswerStringPrimitiveElement != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerString
		if m.AnswerStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AnswerStringPrimitiveElement.Id
			v.Extension = m.AnswerStringPrimitiveElement.Extension
		}
		r.Answer = v
	}
	if m.AnswerCoding != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerCoding
		r.Answer = v
	}
	if m.AnswerQuantity != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerQuantity
		r.Answer = v
	}
	if m.AnswerReference != nil {
		if r.Answer != nil {
			return fmt.Errorf("multiple values for field \"Answer\"")
		}
		v := m.AnswerReference
		r.Answer = v
	}
	return nil
}
func (r QuestionnaireItemEnableWhen) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// One of the permitted answers for a "choice" or "open-choice" question.
type QuestionnaireItemAnswerOption struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A potential answer that's allowed as the answer to this question.
	Value isQuestionnaireItemAnswerOptionValue
	// Indicates whether the answer value is selected when the list of possible answers is initially shown.
	InitialSelected *Boolean
}
type isQuestionnaireItemAnswerOptionValue interface {
	isQuestionnaireItemAnswerOptionValue()
}

func (r Integer) isQuestionnaireItemAnswerOptionValue()   {}
func (r Date) isQuestionnaireItemAnswerOptionValue()      {}
func (r Time) isQuestionnaireItemAnswerOptionValue()      {}
func (r String) isQuestionnaireItemAnswerOptionValue()    {}
func (r Coding) isQuestionnaireItemAnswerOptionValue()    {}
func (r Reference) isQuestionnaireItemAnswerOptionValue() {}

type jsonQuestionnaireItemAnswerOption struct {
	Id                              *string           `json:"id,omitempty"`
	Extension                       []Extension       `json:"extension,omitempty"`
	ModifierExtension               []Extension       `json:"modifierExtension,omitempty"`
	ValueInteger                    *Integer          `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement    *primitiveElement `json:"_valueInteger,omitempty"`
	ValueDate                       *Date             `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement       *primitiveElement `json:"_valueDate,omitempty"`
	ValueTime                       *Time             `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement       *primitiveElement `json:"_valueTime,omitempty"`
	ValueString                     *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement     *primitiveElement `json:"_valueString,omitempty"`
	ValueCoding                     *Coding           `json:"valueCoding,omitempty"`
	ValueReference                  *Reference        `json:"valueReference,omitempty"`
	InitialSelected                 *Boolean          `json:"initialSelected,omitempty"`
	InitialSelectedPrimitiveElement *primitiveElement `json:"_initialSelected,omitempty"`
}

func (r QuestionnaireItemAnswerOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireItemAnswerOption) marshalJSON() jsonQuestionnaireItemAnswerOption {
	m := jsonQuestionnaireItemAnswerOption{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Value.(type) {
	case Integer:
		m.ValueInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.ValueInteger = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		m.ValueDate = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.ValueDate = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		m.ValueTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		m.ValueTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.ValueString = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ValueString = v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	m.InitialSelected = r.InitialSelected
	if r.InitialSelected != nil && (r.InitialSelected.Id != nil || r.InitialSelected.Extension != nil) {
		m.InitialSelectedPrimitiveElement = &primitiveElement{Id: r.InitialSelected.Id, Extension: r.InitialSelected.Extension}
	}
	return m
}
func (r *QuestionnaireItemAnswerOption) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireItemAnswerOption
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireItemAnswerOption) unmarshalJSON(m jsonQuestionnaireItemAnswerOption) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ValueInteger != nil || m.ValueIntegerPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInteger
		if m.ValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.ValueIntegerPrimitiveElement.Id
			v.Extension = m.ValueIntegerPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDate != nil || m.ValueDatePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDate
		if m.ValueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ValueDatePrimitiveElement.Id
			v.Extension = m.ValueDatePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueTime != nil || m.ValueTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTime
		if m.ValueTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.ValueTimePrimitiveElement.Id
			v.Extension = m.ValueTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueString != nil || m.ValueStringPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueString
		if m.ValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ValueStringPrimitiveElement.Id
			v.Extension = m.ValueStringPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	r.InitialSelected = m.InitialSelected
	if m.InitialSelectedPrimitiveElement != nil {
		r.InitialSelected.Id = m.InitialSelectedPrimitiveElement.Id
		r.InitialSelected.Extension = m.InitialSelectedPrimitiveElement.Extension
	}
	return nil
}
func (r QuestionnaireItemAnswerOption) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// One or more values that should be pre-populated in the answer when initially rendering the questionnaire for user input.
type QuestionnaireItemInitial struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The actual value to for an initial answer.
	Value isQuestionnaireItemInitialValue
}
type isQuestionnaireItemInitialValue interface {
	isQuestionnaireItemInitialValue()
}

func (r Boolean) isQuestionnaireItemInitialValue()    {}
func (r Decimal) isQuestionnaireItemInitialValue()    {}
func (r Integer) isQuestionnaireItemInitialValue()    {}
func (r Date) isQuestionnaireItemInitialValue()       {}
func (r DateTime) isQuestionnaireItemInitialValue()   {}
func (r Time) isQuestionnaireItemInitialValue()       {}
func (r String) isQuestionnaireItemInitialValue()     {}
func (r Uri) isQuestionnaireItemInitialValue()        {}
func (r Attachment) isQuestionnaireItemInitialValue() {}
func (r Coding) isQuestionnaireItemInitialValue()     {}
func (r Quantity) isQuestionnaireItemInitialValue()   {}
func (r Reference) isQuestionnaireItemInitialValue()  {}

type jsonQuestionnaireItemInitial struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	ValueBoolean                  *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement  *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueDecimal                  *Decimal          `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement  *primitiveElement `json:"_valueDecimal,omitempty"`
	ValueInteger                  *Integer          `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement  *primitiveElement `json:"_valueInteger,omitempty"`
	ValueDate                     *Date             `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement     *primitiveElement `json:"_valueDate,omitempty"`
	ValueDateTime                 *DateTime         `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement `json:"_valueDateTime,omitempty"`
	ValueTime                     *Time             `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement     *primitiveElement `json:"_valueTime,omitempty"`
	ValueString                   *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement   *primitiveElement `json:"_valueString,omitempty"`
	ValueUri                      *Uri              `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement      *primitiveElement `json:"_valueUri,omitempty"`
	ValueAttachment               *Attachment       `json:"valueAttachment,omitempty"`
	ValueCoding                   *Coding           `json:"valueCoding,omitempty"`
	ValueQuantity                 *Quantity         `json:"valueQuantity,omitempty"`
	ValueReference                *Reference        `json:"valueReference,omitempty"`
}

func (r QuestionnaireItemInitial) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireItemInitial) marshalJSON() jsonQuestionnaireItemInitial {
	m := jsonQuestionnaireItemInitial{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Value.(type) {
	case Boolean:
		m.ValueBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.ValueBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Decimal:
		m.ValueDecimal = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		m.ValueDecimal = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.ValueInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.ValueInteger = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		m.ValueDate = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.ValueDate = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		m.ValueDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.ValueDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		m.ValueTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		m.ValueTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.ValueString = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ValueString = v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uri:
		m.ValueUri = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.ValueUri = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	return m
}
func (r *QuestionnaireItemInitial) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireItemInitial
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireItemInitial) unmarshalJSON(m jsonQuestionnaireItemInitial) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ValueBoolean != nil || m.ValueBooleanPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBoolean
		if m.ValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ValueBooleanPrimitiveElement.Id
			v.Extension = m.ValueBooleanPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDecimal != nil || m.ValueDecimalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDecimal
		if m.ValueDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.ValueDecimalPrimitiveElement.Id
			v.Extension = m.ValueDecimalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInteger != nil || m.ValueIntegerPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInteger
		if m.ValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.ValueIntegerPrimitiveElement.Id
			v.Extension = m.ValueIntegerPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDate != nil || m.ValueDatePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDate
		if m.ValueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ValueDatePrimitiveElement.Id
			v.Extension = m.ValueDatePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDateTime != nil || m.ValueDateTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDateTime
		if m.ValueDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.ValueDateTimePrimitiveElement.Id
			v.Extension = m.ValueDateTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueTime != nil || m.ValueTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTime
		if m.ValueTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.ValueTimePrimitiveElement.Id
			v.Extension = m.ValueTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueString != nil || m.ValueStringPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueString
		if m.ValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ValueStringPrimitiveElement.Id
			v.Extension = m.ValueStringPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUri != nil || m.ValueUriPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUri
		if m.ValueUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.ValueUriPrimitiveElement.Id
			v.Extension = m.ValueUriPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	return nil
}
func (r QuestionnaireItemInitial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
