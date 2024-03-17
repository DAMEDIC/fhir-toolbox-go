package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// An expression that describes applicability criteria, or start/stop conditions for the action.
type RequestGroupActionCondition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The kind of condition.
	Kind types.Code
	// An expression that returns true or false, indicating whether or not the condition is satisfied.
	Expression *types.Expression
}

func (s RequestGroupActionCondition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type RequestGroupActionRelatedAction struct {
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
	// The element id of the action this is related to.
	ActionId types.Id
	// The relationship of this action to the related action.
	Relationship types.Code
}

func (s RequestGroupActionRelatedAction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The actions, if any, produced by the evaluation of the artifact.
type RequestGroupAction struct {
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *types.Code
	// The resource that is the target of the action (e.g. CommunicationRequest).
	Resource *types.Reference
	// Sub actions.
	Action []RequestGroupAction
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The participant that should perform or be responsible for this action.
	Participant []types.Reference
	// A short description of the action used to provide a summary to display to the user.
	Description *types.String
	// An expression that describes applicability criteria, or start/stop conditions for the action.
	Condition []RequestGroupActionCondition
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []RequestGroupActionRelatedAction
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *types.Code
	// Defines expectations around whether an action is required.
	RequiredBehavior *types.Code
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *types.Code
	// A user-visible prefix for the action.
	Prefix *types.String
	// The title of the action displayed to a user.
	Title *types.String
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []types.RelatedArtifact
	// An optional value describing when the action should be performed.
	Timing r4.Element
	// The type of action to perform (create, update, remove).
	Type *types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *types.String
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for a section of a documentation template.
	Code []types.CodeableConcept
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *types.Code
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *types.Code
}

func (s RequestGroupAction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup struct {
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A URL referencing an externally defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesUri []types.Uri
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []types.Reference
	// The current state of the request. For request groups, the status reflects the status of all the requests in the group.
	Status types.Code
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
	// Describes the context of the request group, if any.
	Encounter *types.Reference
	// Indicates when the request group was created.
	AuthoredOn *types.DateTime
	// Describes the reason for the request group in coded or textual form.
	ReasonCode []types.CodeableConcept
	// Provides a mechanism to communicate additional information about the response.
	Note []types.Annotation
	// The actions, if any, produced by the evaluation of the artifact.
	Action []RequestGroupAction
	// Allows a service to provide a unique, business identifier for the request.
	Identifier []types.Identifier
	// A canonical URL referencing a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesCanonical []types.Canonical
	// A plan, proposal or order that is fulfilled in whole or in part by this request.
	BasedOn []types.Reference
	// Indicates how quickly the request should be addressed with respect to other requests.
	Priority *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Indicates the level of authority/intentionality associated with the request and where the request fits into the workflow chain.
	Intent types.Code
	// The subject for which the request group was created.
	Subject *types.Reference
	// Provides a reference to the author of the request group.
	Author *types.Reference
	// Indicates another resource whose existence justifies this request group.
	ReasonReference []types.Reference
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition, prescription or similar form.
	GroupIdentifier *types.Identifier
	// A code that identifies what the overall request group is.
	Code *types.CodeableConcept
}

func (s RequestGroup) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
