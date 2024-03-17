package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Indicates who should participate in performing the action described.
type ActivityDefinitionParticipant struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The type of participant in the action.
	Type types.Code
	// The role the participant should play in performing the described action.
	Role *types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s ActivityDefinitionParticipant) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The path to the element to be customized. This is the path on the resource that will hold the result of the calculation defined by the expression. The specified path SHALL be a FHIRPath resolveable on the specified target type of the ActivityDefinition, and SHALL consist only of identifiers, constant indexers, and a restricted subset of functions. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	Path types.String
	// An expression specifying the value of the customized element.
	Expression types.Expression
}

func (s ActivityDefinitionDynamicValue) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition struct {
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A legal or geographic region in which the activity definition is intended to be used.
	Jurisdiction []types.CodeableConcept
	// The period, timing or frequency upon which the described activity is to occur.
	Timing r4.Element
	// Defines specimen requirements for the action to be performed, such as required specimens for a lab test.
	SpecimenRequirement []types.Reference
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *types.Date
	// Identifies the quantity expected to be consumed at once (per dose, per meal, etc.).
	Quantity *types.Quantity
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A natural language name identifying the activity definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *types.String
	// Indicates who should participate in performing the action described.
	Participant []ActivityDefinitionParticipant
	// A free text natural language description of the activity definition from a consumer's perspective.
	Description *types.Markdown
	// Defines observation requirements for the action to be performed, such as body weight or surface area.
	ObservationRequirement []types.Reference
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// The period during which the activity definition content was or is planned to be in active use.
	EffectivePeriod *types.Period
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []types.ContactDetail
	// Detailed description of the type of activity; e.g. What lab test, what procedure, what kind of encounter.
	Code *types.CodeableConcept
	// Indicates the level of authority/intentionality associated with the activity and where the request should fit into the workflow chain.
	Intent *types.Code
	// Identifies the facility where the activity will occur; e.g. home, hospital, specific clinic, etc.
	Location *types.Reference
	// A short, descriptive, user-friendly title for the activity definition.
	Title *types.String
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate activity definition instances.
	UseContext []types.UsageContext
	// Explanation of why this activity definition is needed and why it has been designed as it has.
	Purpose *types.Markdown
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []types.ContactDetail
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []types.ContactDetail
	// A profile to which the target of the activity definition is expected to conform.
	Profile *types.Canonical
	// Defines the observations that are expected to be produced by the action.
	ObservationResultRequirement []types.Reference
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A copyright statement relating to the activity definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the activity definition.
	Copyright *types.Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *types.Date
	// A reference to a Library resource containing any formal logic used by the activity definition.
	Library []types.Canonical
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	Transform *types.Canonical
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// A formal identifier that is used to identify this activity definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// The identifier that is used to identify this version of the activity definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the activity definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active assets.
	Version *types.String
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A code or group definition that describes the intended subject of the activity being defined.
	Subject r4.Element
	// Indicates how quickly the activity  should be addressed with respect to other requests.
	Priority *types.Code
	// Identifies the food, drug or other product being consumed or supplied in the activity.
	Product r4.Element
	// An absolute URI that is used to identify this activity definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this activity definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the activity definition is stored on different servers.
	Url *types.Uri
	// An explanatory or alternate title for the activity definition giving additional information about its content.
	Subtitle *types.String
	// A Boolean value to indicate that this activity definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *types.Boolean
	// The date  (and optionally time) when the activity definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the activity definition changes.
	Date *types.DateTime
	// Set this to true if the definition is to indicate that a particular activity should NOT be performed. If true, this element should be interpreted to reinforce a negative coding. For example NPO as a code with a doNotPerform of true would still indicate to NOT perform the action.
	DoNotPerform *types.Boolean
	// Provides detailed dosage instructions in the same way that they are described for MedicationRequest resources.
	Dosage []types.Dosage
	// A description of the kind of resource the activity definition is representing. For example, a MedicationRequest, a ServiceRequest, or a CommunicationRequest. Typically, but not always, this is a Request resource.
	Kind *types.Code
	// The status of this activity definition. Enables tracking the life-cycle of the content.
	Status types.Code
	// Descriptive topics related to the content of the activity. Topics provide a high-level categorization of the activity that can be useful for filtering and searching.
	Topic []types.CodeableConcept
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []types.ContactDetail
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []types.RelatedArtifact
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The name of the organization or individual that published the activity definition.
	Publisher *types.String
	// A detailed description of how the activity definition is used from a clinical perspective.
	Usage *types.String
	// Indicates the sites on the subject's body where the procedure should be performed (I.e. the target sites).
	BodySite []types.CodeableConcept
	// Dynamic values that will be evaluated to produce values for elements of the resulting resource. For example, if the dosage of a medication must be computed based on the patient's weight, a dynamic value would be used to specify an expression that calculated the weight, and the path on the request resource that would contain the result.
	DynamicValue []ActivityDefinitionDynamicValue
}

func (s ActivityDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
