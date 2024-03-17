package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Additional information that may be needed in the execution of the task.
type TaskInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A code or description indicating how the input is intended to be used as part of the task execution.
	Type types.CodeableConcept
	// The value of the input parameter as a basic type.
	Value r4.Element
}

func (s TaskInput) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// If the Task.focus is a request resource and the task is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
type TaskRestriction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates the number of times the requested action should occur.
	Repetitions *types.PositiveInt
	// Over what time-period is fulfillment sought.
	Period *types.Period
	// For requests that are targeted to more than on potential recipient/target, for whom is fulfillment sought?
	Recipient []types.Reference
}

func (s TaskRestriction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Outputs produced by the Task.
type TaskOutput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The name of the Output parameter.
	Type types.CodeableConcept
	// The value of the Output parameter as a basic type.
	Value r4.Element
}

func (s TaskOutput) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A task to be performed.
type Task struct {
	// A name or code (or both) briefly describing what the task involves.
	Code *types.CodeableConcept
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Task.
	Insurance []types.Reference
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// An explanation as to why this task is held, failed, was refused, etc.
	StatusReason *types.CodeableConcept
	// The business identifier for this task.
	Identifier []types.Identifier
	// The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
	InstantiatesCanonical *types.Canonical
	// The request being actioned or the resource being manipulated by this task.
	Focus *types.Reference
	// Identifies the time action was first taken against the task (start) and/or the time final action was taken against the task prior to marking it as completed (end).
	ExecutionPeriod *types.Period
	// Additional information that may be needed in the execution of the task.
	Input []TaskInput
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A free-text description of what is to be performed.
	Description *types.String
	// The creator of the task.
	Requester *types.Reference
	// A resource reference indicating why this task needs to be performed.
	ReasonReference *types.Reference
	// The base language in which the resource is written.
	Language *types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Contains business-specific nuances of the business state.
	BusinessStatus *types.CodeableConcept
	// Task that this particular task is part of.
	PartOf []types.Reference
	// The date and time this task was created.
	AuthoredOn *types.DateTime
	// Principal physical location where the this task is performed.
	Location *types.Reference
	// A description or code indicating why this task needs to be performed.
	ReasonCode *types.CodeableConcept
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// An identifier that links together multiple tasks and other requests that were created in the same context.
	GroupIdentifier *types.Identifier
	// Links to Provenance records for past versions of this Task that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the task.
	RelevantHistory []types.Reference
	// If the Task.focus is a request resource and the task is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
	Restriction *TaskRestriction
	// The date and time of last modification to this task.
	LastModified *types.DateTime
	// The kind of participant that should perform the task.
	PerformerType []types.CodeableConcept
	// The entity who benefits from the performance of the service specified in the task (e.g., the patient).
	For *types.Reference
	// Individual organization or Device currently responsible for task execution.
	Owner *types.Reference
	// The URL pointing to an *externally* maintained  protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
	InstantiatesUri *types.Uri
	// BasedOn refers to a higher-level authorization that triggered the creation of the task.  It references a "request" resource such as a ServiceRequest, MedicationRequest, ServiceRequest, CarePlan, etc. which is distinct from the "request" resource the task is seeking to fulfill.  This latter resource is referenced by FocusOn.  For example, based on a ServiceRequest (= BasedOn), a task is created to fulfill a procedureRequest ( = FocusOn ) to collect a specimen from a patient.
	BasedOn []types.Reference
	// Indicates the "level" of actionability associated with the Task, i.e. i+R[9]Cs this a proposed task, a planned task, an actionable task, etc.
	Intent types.Code
	// Indicates how quickly the Task should be addressed with respect to other requests.
	Priority *types.Code
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this task was created.
	Encounter *types.Reference
	// Free-text information captured about the task as it progresses.
	Note []types.Annotation
	// Outputs produced by the Task.
	Output []TaskOutput
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The current status of the task.
	Status types.Code
}

func (s Task) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
