package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
type EvidenceVariableCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Define members of the evidence element using Codes (such as condition, medication, or observation), Expressions ( using an expression language such as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in the last year).
	Definition r4.Element
	// Indicates what effective period the study covers.
	ParticipantEffective r4.Element
	// Indicates how elements are aggregated within the study effective period.
	GroupMeasure *types.Code
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A short, natural language description of the characteristic that could be used to communicate the criteria to an end-user.
	Description *types.String
	// Use UsageContext to define the members of the population, such as Age Ranges, Genders, Settings.
	UsageContext []types.UsageContext
	// When true, members with this characteristic are excluded from the element.
	Exclude *types.Boolean
	// Indicates duration from the participant's study entry.
	TimeFromStart *types.Duration
}

func (s EvidenceVariableCharacteristic) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The EvidenceVariable resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
//
// Need to be able to define and reuse the definition of individual elements of a research question.
type EvidenceVariable struct {
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A short, descriptive, user-friendly title for the evidence variable.
	Title *types.String
	// A human-readable string to clarify or explain concepts about the resource.
	Note []types.Annotation
	// A legal or geographic region in which the evidence variable is intended to be used.
	Jurisdiction []types.CodeableConcept
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []types.ContactDetail
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []types.RelatedArtifact
	// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
	Characteristic []EvidenceVariableCharacteristic
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A formal identifier that is used to identify this evidence variable when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// A free text natural language description of the evidence variable from a consumer's perspective.
	Description *types.Markdown
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *types.Date
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate evidence variable instances.
	UseContext []types.UsageContext
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *types.Date
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []types.ContactDetail
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []types.ContactDetail
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// The period during which the evidence variable content was or is planned to be in active use.
	EffectivePeriod *types.Period
	// The name of the organization or individual that published the evidence variable.
	Publisher *types.String
	// A copyright statement relating to the evidence variable and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the evidence variable.
	Copyright *types.Markdown
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []types.ContactDetail
	// The type of evidence element, a population, an exposure, or an outcome.
	Type *types.Code
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// A natural language name identifying the evidence variable. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *types.String
	// The short title provides an alternate title for use in informal descriptive contexts where the full, formal title is not necessary.
	ShortTitle *types.String
	// The date  (and optionally time) when the evidence variable was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the evidence variable changes.
	Date *types.DateTime
	// The base language in which the resource is written.
	Language *types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// An absolute URI that is used to identify this evidence variable when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this evidence variable is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the evidence variable is stored on different servers.
	Url *types.Uri
	// An explanatory or alternate title for the EvidenceVariable giving additional information about its content.
	Subtitle *types.String
	// The identifier that is used to identify this version of the evidence variable when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the evidence variable author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *types.String
	// The status of this evidence variable. Enables tracking the life-cycle of the content.
	Status types.Code
	// Descriptive topics related to the content of the EvidenceVariable. Topics provide a high-level categorization grouping types of EvidenceVariables that can be useful for filtering and searching.
	Topic []types.CodeableConcept
}

func (s EvidenceVariable) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
