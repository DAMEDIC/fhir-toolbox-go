package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// The EventDefinition resource provides a reusable description of when a particular event can occur.
type EventDefinition struct {
	// The base language in which the resource is written.
	Language *types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Explanation of why this event definition is needed and why it has been designed as it has.
	Purpose *types.Markdown
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []types.ContactDetail
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The identifier that is used to identify this version of the event definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the event definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *types.String
	// A short, descriptive, user-friendly title for the event definition.
	Title *types.String
	// Related resources such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []types.RelatedArtifact
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// An explanatory or alternate title for the event definition giving additional information about its content.
	Subtitle *types.String
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate event definition instances.
	UseContext []types.UsageContext
	// A legal or geographic region in which the event definition is intended to be used.
	Jurisdiction []types.CodeableConcept
	// The trigger element defines when the event occurs. If more than one trigger condition is specified, the event fires whenever any one of the trigger conditions is met.
	Trigger []types.TriggerDefinition
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// An absolute URI that is used to identify this event definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this event definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the event definition is stored on different servers.
	Url *types.Uri
	// A code or group definition that describes the intended subject of the event definition.
	Subject r4.Element
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// A detailed description of how the event definition is used from a clinical perspective.
	Usage *types.String
	// A Boolean value to indicate that this event definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *types.Boolean
	// The period during which the event definition content was or is planned to be in active use.
	EffectivePeriod *types.Period
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []types.ContactDetail
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []types.ContactDetail
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A natural language name identifying the event definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *types.String
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *types.Date
	// Descriptive topics related to the module. Topics provide a high-level categorization of the module that can be useful for filtering and searching.
	Topic []types.CodeableConcept
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []types.ContactDetail
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// A formal identifier that is used to identify this event definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// The status of this event definition. Enables tracking the life-cycle of the content.
	Status types.Code
	// The date  (and optionally time) when the event definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the event definition changes.
	Date *types.DateTime
	// A copyright statement relating to the event definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the event definition.
	Copyright *types.Markdown
	// The name of the organization or individual that published the event definition.
	Publisher *types.String
	// A free text natural language description of the event definition from a consumer's perspective.
	Description *types.Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *types.Date
}

func (s EventDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
