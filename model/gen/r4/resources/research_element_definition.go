package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// A characteristic that defines the members of the research element. Multiple characteristics are applied with "and" semantics.
type ResearchElementDefinitionCharacteristic struct {
	// Indicates how elements are aggregated within the study effective period.
	ParticipantEffectiveGroupMeasure *types.Code
	// A narrative description of the time period the study covers.
	StudyEffectiveDescription *types.String
	// Indicates how elements are aggregated within the study effective period.
	StudyEffectiveGroupMeasure *types.Code
	// Indicates what effective period the study covers.
	ParticipantEffective r4.Element
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates duration from the participant's study entry.
	ParticipantEffectiveTimeFromStart *types.Duration
	// When true, members with this characteristic are excluded from the element.
	Exclude *types.Boolean
	// Indicates what effective period the study covers.
	StudyEffective r4.Element
	// A narrative description of the time period the study covers.
	ParticipantEffectiveDescription *types.String
	// Use UsageContext to define the members of the population, such as Age Ranges, Genders, Settings.
	UsageContext []types.UsageContext
	// Specifies the UCUM unit for the outcome.
	UnitOfMeasure *types.CodeableConcept
	// Indicates duration from the study initiation.
	StudyEffectiveTimeFromStart *types.Duration
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Define members of the research element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
	Definition r4.Element
}

func (s ResearchElementDefinitionCharacteristic) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The ResearchElementDefinition resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
//
// Need to be able to define and reuse the definition of individual elements of a research question.
type ResearchElementDefinition struct {
	// A free text natural language description of the research element definition from a consumer's perspective.
	Description *types.Markdown
	// Explanation of why this research element definition is needed and why it has been designed as it has.
	Purpose *types.Markdown
	// The type of research element, a population, an exposure, or an outcome.
	Type types.Code
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []types.ContactDetail
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// An absolute URI that is used to identify this research element definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this research element definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the research element definition is stored on different servers.
	Url *types.Uri
	// A legal or geographic region in which the research element definition is intended to be used.
	Jurisdiction []types.CodeableConcept
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *types.Date
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []types.ContactDetail
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The name of the organization or individual that published the research element definition.
	Publisher *types.String
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate research element definition instances.
	UseContext []types.UsageContext
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []types.ContactDetail
	// A formal identifier that is used to identify this research element definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// The short title provides an alternate title for use in informal descriptive contexts where the full, formal title is not necessary.
	ShortTitle *types.String
	// A detailed description, from a clinical perspective, of how the ResearchElementDefinition is used.
	Usage *types.String
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []types.RelatedArtifact
	// The date  (and optionally time) when the research element definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the research element definition changes.
	Date *types.DateTime
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// A copyright statement relating to the research element definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the research element definition.
	Copyright *types.Markdown
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A natural language name identifying the research element definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *types.String
	// A Boolean value to indicate that this research element definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *types.Boolean
	// The intended subjects for the ResearchElementDefinition. If this element is not provided, a Patient subject is assumed, but the subject of the ResearchElementDefinition can be anything.
	Subject r4.Element
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *types.Date
	// A reference to a Library resource containing the formal logic used by the ResearchElementDefinition.
	Library []types.Canonical
	// The type of the outcome (e.g. Dichotomous, Continuous, or Descriptive).
	VariableType *types.Code
	// Descriptive topics related to the content of the ResearchElementDefinition. Topics provide a high-level categorization grouping types of ResearchElementDefinitions that can be useful for filtering and searching.
	Topic []types.CodeableConcept
	// A human-readable string to clarify or explain concepts about the resource.
	Comment []types.String
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A short, descriptive, user-friendly title for the research element definition.
	Title *types.String
	// An explanatory or alternate title for the ResearchElementDefinition giving additional information about its content.
	Subtitle *types.String
	// The status of this research element definition. Enables tracking the life-cycle of the content.
	Status types.Code
	// A characteristic that defines the members of the research element. Multiple characteristics are applied with "and" semantics.
	Characteristic []ResearchElementDefinitionCharacteristic
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// The identifier that is used to identify this version of the research element definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the research element definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *types.String
	// The period during which the research element definition content was or is planned to be in active use.
	EffectivePeriod *types.Period
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []types.ContactDetail
}

func (s ResearchElementDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
