package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// One of the permitted answers for a "choice" or "open-choice" question.
type QuestionnaireItemAnswerOption struct {
	// Indicates whether the answer value is selected when the list of possible answers is initially shown.
	InitialSelected *types.Boolean
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A potential answer that's allowed as the answer to this question.
	Value r4.Element
}

func (s QuestionnaireItemAnswerOption) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// One or more values that should be pre-populated in the answer when initially rendering the questionnaire for user input.
type QuestionnaireItemInitial struct {
	// The actual value to for an initial answer.
	Value r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s QuestionnaireItemInitial) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A constraint indicating that this item should only be enabled (displayed/allow answers to be captured) when the specified condition is true.
type QuestionnaireItemEnableWhen struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The linkId for the question whose answer (or lack of answer) governs whether this item is enabled.
	Question types.String
	// Specifies the criteria by which the question is enabled.
	Operator types.Code
	// A value that the referenced question is tested using the specified operator in order for the item to be enabled.
	Answer r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s QuestionnaireItemEnableWhen) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A particular question, question grouping or display text that is part of the questionnaire.
type QuestionnaireItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A terminology code that corresponds to this group or question (e.g. a code from LOINC, which defines many questions and answers).
	Code []types.Coding
	// The maximum number of characters that are permitted in the answer to be considered a "valid" QuestionnaireResponse.
	MaxLength *types.Integer
	// One of the permitted answers for a "choice" or "open-choice" question.
	AnswerOption []QuestionnaireItemAnswerOption
	// One or more values that should be pre-populated in the answer when initially rendering the questionnaire for user input.
	Initial []QuestionnaireItemInitial
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A constraint indicating that this item should only be enabled (displayed/allow answers to be captured) when the specified condition is true.
	EnableWhen []QuestionnaireItemEnableWhen
	// An indication, if true, that the item must be present in a "completed" QuestionnaireResponse.  If false, the item may be skipped when answering the questionnaire.
	Required *types.Boolean
	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question.
	AnswerValueSet *types.Canonical
	// The name of a section, the text of a question or text content for a display item.
	Text *types.String
	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured (string, integer, coded choice, etc.).
	Type types.Code
	// Controls how multiple enableWhen values are interpreted -  whether all or any must be true.
	EnableBehavior *types.Code
	// An indication, when true, that the value cannot be changed by a human respondent to the Questionnaire.
	ReadOnly *types.Boolean
	// Text, questions and other groups to be nested beneath a question or group.
	Item []QuestionnaireItem
	// An identifier that is unique within the Questionnaire allowing linkage to the equivalent item in a QuestionnaireResponse resource.
	LinkId types.String
	// This element is a URI that refers to an [ElementDefinition](elementdefinition.html) that provides information about this item, including information that might otherwise be included in the instance of the Questionnaire resource. A detailed description of the construction of the URI is shown in Comments, below. If this element is present then the following element values MAY be derived from the Element Definition if the corresponding elements of this Questionnaire resource instance have no value:
	//
	// * code (ElementDefinition.code)
	// * type (ElementDefinition.type)
	// * required (ElementDefinition.min)
	// * repeats (ElementDefinition.max)
	// * maxLength (ElementDefinition.maxLength)
	// * answerValueSet (ElementDefinition.binding)
	// * options (ElementDefinition.binding).
	Definition *types.Uri
	// A short label for a particular group, question or set of display text within the questionnaire used for reference by the individual completing the questionnaire.
	Prefix *types.String
	// An indication, if true, that the item may occur multiple times in the response, collecting multiple answers for questions or multiple sets of answers for groups.
	Repeats *types.Boolean
}

func (s QuestionnaireItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
//
// To support structured, hierarchical registration of data gathered using digital forms and other questionnaires.  Questionnaires provide greater control over presentation and allow capture of data in a domain-independent way (i.e. capturing information that would otherwise require multiple distinct types of resources).
type Questionnaire struct {
	// An absolute URI that is used to identify this questionnaire when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this questionnaire is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the questionnaire is stored on different servers.
	Url *types.Uri
	// The URL of a Questionnaire that this Questionnaire is based on.
	DerivedFrom []types.Canonical
	// The status of this questionnaire. Enables tracking the life-cycle of the content.
	Status types.Code
	// A Boolean value to indicate that this questionnaire is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *types.Boolean
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The date  (and optionally time) when the questionnaire was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the questionnaire changes.
	Date *types.DateTime
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate questionnaire instances.
	UseContext []types.UsageContext
	// A copyright statement relating to the questionnaire and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the questionnaire.
	Copyright *types.Markdown
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// Explanation of why this questionnaire is needed and why it has been designed as it has.
	Purpose *types.Markdown
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *types.Date
	// A particular question, question grouping or display text that is part of the questionnaire.
	Item []QuestionnaireItem
	// The identifier that is used to identify this version of the questionnaire when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the questionnaire author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *types.String
	// The name of the organization or individual that published the questionnaire.
	Publisher *types.String
	// A formal identifier that is used to identify this questionnaire when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// A free text natural language description of the questionnaire from a consumer's perspective.
	Description *types.Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *types.Date
	// A natural language name identifying the questionnaire. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *types.String
	// A legal or geographic region in which the questionnaire is intended to be used.
	Jurisdiction []types.CodeableConcept
	// The base language in which the resource is written.
	Language *types.Code
	// The types of subjects that can be the subject of responses created for the questionnaire.
	SubjectType []types.Code
	// An identifier for this question or group of questions in a particular terminology such as LOINC.
	Code []types.Coding
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A short, descriptive, user-friendly title for the questionnaire.
	Title *types.String
	// The period during which the questionnaire content was or is planned to be in active use.
	EffectivePeriod *types.Period
}

func (s Questionnaire) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
