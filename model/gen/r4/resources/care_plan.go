package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
type CarePlanActivityDetail struct {
	// The URL pointing to an externally maintained protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan activity.
	InstantiatesUri []types.Uri
	// Indicates another resource, such as the health condition(s), whose existence justifies this request and drove the inclusion of this particular activity as part of the plan.
	ReasonReference []types.Reference
	// Identifies who's expected to be involved in the activity.
	Performer []types.Reference
	// Identifies the food, drug or other product to be consumed or supplied in the activity.
	Product r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Identifies the quantity expected to be supplied, administered or consumed by the subject.
	Quantity *types.Quantity
	// This provides a textual description of constraints on the intended activity occurrence, including relation to other activities.  It may also include objectives, pre-conditions and end-conditions.  Finally, it may convey specifics about the activity such as body site, method, route, etc.
	Description *types.String
	// A description of the kind of resource the in-line definition of a care plan activity is representing.  The CarePlan.activity.detail is an in-line definition when a resource is not referenced using CarePlan.activity.reference.  For example, a MedicationRequest, a ServiceRequest, or a CommunicationRequest.
	Kind *types.Code
	// Provides the rationale that drove the inclusion of this particular activity as part of the plan or the reason why the activity was prohibited.
	ReasonCode []types.CodeableConcept
	// Internal reference that identifies the goals that this activity is intended to contribute towards meeting.
	Goal []types.Reference
	// Identifies what progress is being made for the specific activity.
	Status types.Code
	// Identifies the quantity expected to be consumed in a given day.
	DailyAmount *types.Quantity
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan activity.
	InstantiatesCanonical []types.Canonical
	// Detailed description of the type of planned activity; e.g. what lab test, what procedure, what kind of encounter.
	Code *types.CodeableConcept
	// Provides reason why the activity isn't yet started, is on hold, was cancelled, etc.
	StatusReason *types.CodeableConcept
	// If true, indicates that the described activity is one that must NOT be engaged in when following the plan.  If false, or missing, indicates that the described activity is one that should be engaged in when following the plan.
	DoNotPerform *types.Boolean
	// The period, timing or frequency upon which the described activity is to occur.
	Scheduled r4.Element
	// Identifies the facility where the activity will occur; e.g. home, hospital, specific clinic, etc.
	Location *types.Reference
}

func (s CarePlanActivityDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
type CarePlanActivity struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identifies the outcome at the point when the status of the activity is assessed.  For example, the outcome of an education activity could be patient understands (or not).
	OutcomeCodeableConcept []types.CodeableConcept
	// Details of the outcome or action resulting from the activity.  The reference to an "event" resource, such as Procedure or Encounter or Observation, is the result/outcome of the activity itself.  The activity can be conveyed using CarePlan.activity.detail OR using the CarePlan.activity.reference (a reference to a “request” resource).
	OutcomeReference []types.Reference
	// Notes about the adherence/status/progress of the activity.
	Progress []types.Annotation
	// The details of the proposed activity represented in a specific resource.
	Reference *types.Reference
	// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
	Detail *CarePlanActivityDetail
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s CarePlanActivity) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan struct {
	// A care plan that is fulfilled in whole or in part by this care plan.
	BasedOn []types.Reference
	// Human-friendly name for the care plan.
	Title *types.String
	// Describes the intended objective(s) of carrying out the care plan.
	Goal []types.Reference
	// The URL pointing to an externally maintained protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan.
	InstantiatesUri []types.Uri
	// A description of the scope and nature of the plan.
	Description *types.String
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identifies all people and organizations who are expected to be involved in the care envisioned by this plan.
	CareTeam []types.Reference
	// General notes about the care plan not covered elsewhere.
	Note []types.Annotation
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// Indicates whether the plan is currently being acted upon, represents future intentions or is now a historical record.
	Status types.Code
	// Indicates the level of authority/intentionality associated with the care plan and where the care plan fits into the workflow chain.
	Intent types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Business identifiers assigned to this care plan by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []types.Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan.
	InstantiatesCanonical []types.Canonical
	// When populated, the author is responsible for the care plan.  The care plan is attributed to the author.
	Author *types.Reference
	// Identifies the patient or group whose intended care is described by the plan.
	Subject types.Reference
	// Represents when this particular CarePlan record was created in the system, which is often a system-generated date.
	Created *types.DateTime
	// Identifies portions of the patient's record that specifically influenced the formation of the plan.  These might include comorbidities, recent procedures, limitations, recent assessments, etc.
	SupportingInfo []types.Reference
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A larger care plan of which this particular care plan is a component or step.
	PartOf []types.Reference
	// The Encounter during which this CarePlan was created or to which the creation of this record is tightly associated.
	Encounter *types.Reference
	// Identifies the conditions/problems/concerns/diagnoses/etc. whose management and/or mitigation are handled by this plan.
	Addresses []types.Reference
	// The base language in which the resource is written.
	Language *types.Code
	// Completed or terminated care plan whose function is taken by this new care plan.
	Replaces []types.Reference
	// Indicates when the plan did (or is intended to) come into effect and end.
	Period *types.Period
	// Identifies the individual(s) or organization who provided the contents of the care plan.
	Contributor []types.Reference
	// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
	Activity []CarePlanActivity
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// Identifies what "kind" of plan this is to support differentiation between multiple co-existing plans; e.g. "Home health", "psychiatric", "asthma", "disease management", "wellness plan", etc.
	Category []types.CodeableConcept
}

func (s CarePlan) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
