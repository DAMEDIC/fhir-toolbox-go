package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Indicates who should participate in performing the action described.
type PlanDefinitionActionParticipant struct {
	// The role the participant should play in performing the described action.
	Role *types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The type of participant in the action.
	Type types.Code
}

func (s PlanDefinitionActionParticipant) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The path to the element to be customized. This is the path on the resource that will hold the result of the calculation defined by the expression. The specified path SHALL be a FHIRPath resolveable on the specified target type of the ActivityDefinition, and SHALL consist only of identifiers, constant indexers, and a restricted subset of functions. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	Path *types.String
	// An expression specifying the value of the customized element.
	Expression *types.Expression
}

func (s PlanDefinitionActionDynamicValue) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An expression that describes applicability criteria or start/stop conditions for the action.
type PlanDefinitionActionCondition struct {
	// The kind of condition.
	Kind types.Code
	// An expression that returns true or false, indicating whether the condition is satisfied.
	Expression *types.Expression
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s PlanDefinitionActionCondition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type PlanDefinitionActionRelatedAction struct {
	// The element id of the related action.
	ActionId types.Id
	// The relationship of this action to the related action.
	Relationship types.Code
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	Offset r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s PlanDefinitionActionRelatedAction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An action or group of actions to be taken as part of the plan.
type PlanDefinitionAction struct {
	// A description of when the action should be triggered.
	Trigger []types.TriggerDefinition
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *types.Code
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// A brief description of the action used to provide a summary to display to the user.
	Description *types.String
	// A code or group definition that describes the intended subject of the action and its children, if any.
	Subject r4.Element
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *types.Code
	// A reference to an ActivityDefinition that describes the action to be taken in detail, or a PlanDefinition that describes a series of actions to be taken.
	Definition r4.Element
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []types.RelatedArtifact
	// Identifies goals that this action supports. The reference must be to a goal element defined within this plan definition.
	GoalId []types.Id
	// Indicates who should participate in performing the action described.
	Participant []PlanDefinitionActionParticipant
	// Customizations that should be applied to the statically defined resource. For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
	DynamicValue []PlanDefinitionActionDynamicValue
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for the section of a documentation template.
	Code []types.CodeableConcept
	// An optional value describing when the action should be performed.
	Timing r4.Element
	// Defines the required behavior for the action.
	RequiredBehavior *types.Code
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *types.Code
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	Transform *types.Canonical
	// Sub actions that are contained within the action. The behavior of this action determines the functionality of the sub-actions. For example, a selection behavior of at-most-one indicates that of the sub-actions, at most one may be chosen as part of realizing the action definition.
	Action []PlanDefinitionAction
	// A user-visible prefix for the action.
	Prefix *types.String
	// The title of the action displayed to a user.
	Title *types.String
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *types.Code
	// Defines the outputs of the action, if any.
	Output []types.DataRequirement
	// The type of action to perform (create, update, remove).
	Type *types.CodeableConcept
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *types.String
	// A description of why this action is necessary or appropriate.
	Reason []types.CodeableConcept
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *types.Code
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// An expression that describes applicability criteria or start/stop conditions for the action.
	Condition []PlanDefinitionActionCondition
	// Defines input data requirements for the action.
	Input []types.DataRequirement
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []PlanDefinitionActionRelatedAction
}

func (s PlanDefinitionAction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The parameter whose value is to be tracked, e.g. body weight, blood pressure, or hemoglobin A1c level.
	Measure *types.CodeableConcept
	// The target value of the measure to be achieved to signify fulfillment of the goal, e.g. 150 pounds or 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any value at or above the low value.
	Detail r4.Element
	// Indicates the timeframe after the start of the goal in which the goal should be met.
	Due *types.Duration
}

func (s PlanDefinitionGoalTarget) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Goals that describe what the activities within the plan are intended to achieve. For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
type PlanDefinitionGoal struct {
	// Indicates what should be done and within what timeframe.
	Target []PlanDefinitionGoalTarget
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates a category the goal falls within.
	Category *types.CodeableConcept
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	Description types.CodeableConcept
	// The event after which the goal should begin being pursued.
	Start *types.CodeableConcept
	// Identifies problems, conditions, issues, or concerns the goal is intended to address.
	Addresses []types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Identifies the expected level of importance associated with reaching/sustaining the defined goal.
	Priority *types.CodeableConcept
	// Didactic or other informational resources associated with the goal that provide further supporting information about the goal. Information resources can include inline text commentary and links to web resources.
	Documentation []types.RelatedArtifact
}

func (s PlanDefinitionGoal) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition struct {
	// The base language in which the resource is written.
	Language *types.Code
	// A free text natural language description of the plan definition from a consumer's perspective.
	Description *types.Markdown
	// A legal or geographic region in which the plan definition is intended to be used.
	Jurisdiction []types.CodeableConcept
	// Descriptive topics related to the content of the plan definition. Topics provide a high-level categorization of the definition that can be useful for filtering and searching.
	Topic []types.CodeableConcept
	// An action or group of actions to be taken as part of the plan.
	Action []PlanDefinitionAction
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The period during which the plan definition content was or is planned to be in active use.
	EffectivePeriod *types.Period
	// The identifier that is used to identify this version of the plan definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the plan definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *types.String
	// A Boolean value to indicate that this plan definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *types.Boolean
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate plan definition instances.
	UseContext []types.UsageContext
	// A copyright statement relating to the plan definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the plan definition.
	Copyright *types.Markdown
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The date  (and optionally time) when the plan definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the plan definition changes.
	Date *types.DateTime
	// A detailed description of how the plan definition is used from a clinical perspective.
	Usage *types.String
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []types.RelatedArtifact
	// The status of this plan definition. Enables tracking the life-cycle of the content.
	Status types.Code
	// A code or group definition that describes the intended subject of the plan definition.
	Subject r4.Element
	// The name of the organization or individual that published the plan definition.
	Publisher *types.String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []types.ContactDetail
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A short, descriptive, user-friendly title for the plan definition.
	Title *types.String
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *types.Date
	// Goals that describe what the activities within the plan are intended to achieve. For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
	Goal []PlanDefinitionGoal
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A high-level category for the plan definition that distinguishes the kinds of systems that would be interested in the plan definition.
	Type *types.CodeableConcept
	// Explanation of why this plan definition is needed and why it has been designed as it has.
	Purpose *types.Markdown
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *types.Date
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []types.ContactDetail
	// A reference to a Library resource containing any formal logic used by the plan definition.
	Library []types.Canonical
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// An absolute URI that is used to identify this plan definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this plan definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the plan definition is stored on different servers.
	Url *types.Uri
	// A formal identifier that is used to identify this plan definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// A natural language name identifying the plan definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *types.String
	// An explanatory or alternate title for the plan definition giving additional information about its content.
	Subtitle *types.String
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []types.ContactDetail
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []types.ContactDetail
}

func (s PlanDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
