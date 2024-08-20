package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A task to be performed.
type Task struct {
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
	// The business identifier for this task.
	Identifier []Identifier
	// The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
	InstantiatesCanonical *Canonical
	// The URL pointing to an *externally* maintained  protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task.
	InstantiatesUri *Uri
	// BasedOn refers to a higher-level authorization that triggered the creation of the task.  It references a "request" resource such as a ServiceRequest, MedicationRequest, ServiceRequest, CarePlan, etc. which is distinct from the "request" resource the task is seeking to fulfill.  This latter resource is referenced by FocusOn.  For example, based on a ServiceRequest (= BasedOn), a task is created to fulfill a procedureRequest ( = FocusOn ) to collect a specimen from a patient.
	BasedOn []Reference
	// An identifier that links together multiple tasks and other requests that were created in the same context.
	GroupIdentifier *Identifier
	// Task that this particular task is part of.
	PartOf []Reference
	// The current status of the task.
	Status Code
	// An explanation as to why this task is held, failed, was refused, etc.
	StatusReason *CodeableConcept
	// Contains business-specific nuances of the business state.
	BusinessStatus *CodeableConcept
	// Indicates the "level" of actionability associated with the Task, i.e. i+R[9]Cs this a proposed task, a planned task, an actionable task, etc.
	Intent Code
	// Indicates how quickly the Task should be addressed with respect to other requests.
	Priority *Code
	// A name or code (or both) briefly describing what the task involves.
	Code *CodeableConcept
	// A free-text description of what is to be performed.
	Description *String
	// The request being actioned or the resource being manipulated by this task.
	Focus *Reference
	// The entity who benefits from the performance of the service specified in the task (e.g., the patient).
	For *Reference
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this task was created.
	Encounter *Reference
	// Identifies the time action was first taken against the task (start) and/or the time final action was taken against the task prior to marking it as completed (end).
	ExecutionPeriod *Period
	// The date and time this task was created.
	AuthoredOn *DateTime
	// The date and time of last modification to this task.
	LastModified *DateTime
	// The creator of the task.
	Requester *Reference
	// The kind of participant that should perform the task.
	PerformerType []CodeableConcept
	// Individual organization or Device currently responsible for task execution.
	Owner *Reference
	// Principal physical location where the this task is performed.
	Location *Reference
	// A description or code indicating why this task needs to be performed.
	ReasonCode *CodeableConcept
	// A resource reference indicating why this task needs to be performed.
	ReasonReference *Reference
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Task.
	Insurance []Reference
	// Free-text information captured about the task as it progresses.
	Note []Annotation
	// Links to Provenance records for past versions of this Task that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the task.
	RelevantHistory []Reference
	// If the Task.focus is a request resource and the task is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
	Restriction *TaskRestriction
	// Additional information that may be needed in the execution of the task.
	Input []TaskInput
	// Outputs produced by the Task.
	Output []TaskOutput
}

func (r Task) ResourceType() string {
	return "Task"
}
func (r Task) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonTask struct {
	ResourceType                          string              `json:"resourceType"`
	Id                                    *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement   `json:"_id,omitempty"`
	Meta                                  *Meta               `json:"meta,omitempty"`
	ImplicitRules                         *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                              *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement   `json:"_language,omitempty"`
	Text                                  *Narrative          `json:"text,omitempty"`
	Contained                             []ContainedResource `json:"contained,omitempty"`
	Extension                             []Extension         `json:"extension,omitempty"`
	ModifierExtension                     []Extension         `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier        `json:"identifier,omitempty"`
	InstantiatesCanonical                 *Canonical          `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement *primitiveElement   `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       *Uri                `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       *primitiveElement   `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference         `json:"basedOn,omitempty"`
	GroupIdentifier                       *Identifier         `json:"groupIdentifier,omitempty"`
	PartOf                                []Reference         `json:"partOf,omitempty"`
	Status                                Code                `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement   `json:"_status,omitempty"`
	StatusReason                          *CodeableConcept    `json:"statusReason,omitempty"`
	BusinessStatus                        *CodeableConcept    `json:"businessStatus,omitempty"`
	Intent                                Code                `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement   `json:"_intent,omitempty"`
	Priority                              *Code               `json:"priority,omitempty"`
	PriorityPrimitiveElement              *primitiveElement   `json:"_priority,omitempty"`
	Code                                  *CodeableConcept    `json:"code,omitempty"`
	Description                           *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement           *primitiveElement   `json:"_description,omitempty"`
	Focus                                 *Reference          `json:"focus,omitempty"`
	For                                   *Reference          `json:"for,omitempty"`
	Encounter                             *Reference          `json:"encounter,omitempty"`
	ExecutionPeriod                       *Period             `json:"executionPeriod,omitempty"`
	AuthoredOn                            *DateTime           `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement            *primitiveElement   `json:"_authoredOn,omitempty"`
	LastModified                          *DateTime           `json:"lastModified,omitempty"`
	LastModifiedPrimitiveElement          *primitiveElement   `json:"_lastModified,omitempty"`
	Requester                             *Reference          `json:"requester,omitempty"`
	PerformerType                         []CodeableConcept   `json:"performerType,omitempty"`
	Owner                                 *Reference          `json:"owner,omitempty"`
	Location                              *Reference          `json:"location,omitempty"`
	ReasonCode                            *CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference                       *Reference          `json:"reasonReference,omitempty"`
	Insurance                             []Reference         `json:"insurance,omitempty"`
	Note                                  []Annotation        `json:"note,omitempty"`
	RelevantHistory                       []Reference         `json:"relevantHistory,omitempty"`
	Restriction                           *TaskRestriction    `json:"restriction,omitempty"`
	Input                                 []TaskInput         `json:"input,omitempty"`
	Output                                []TaskOutput        `json:"output,omitempty"`
}

func (r Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Task) marshalJSON() jsonTask {
	m := jsonTask{}
	m.ResourceType = "Task"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.InstantiatesCanonical = r.InstantiatesCanonical
	if r.InstantiatesCanonical != nil && (r.InstantiatesCanonical.Id != nil || r.InstantiatesCanonical.Extension != nil) {
		m.InstantiatesCanonicalPrimitiveElement = &primitiveElement{Id: r.InstantiatesCanonical.Id, Extension: r.InstantiatesCanonical.Extension}
	}
	m.InstantiatesUri = r.InstantiatesUri
	if r.InstantiatesUri != nil && (r.InstantiatesUri.Id != nil || r.InstantiatesUri.Extension != nil) {
		m.InstantiatesUriPrimitiveElement = &primitiveElement{Id: r.InstantiatesUri.Id, Extension: r.InstantiatesUri.Extension}
	}
	m.BasedOn = r.BasedOn
	m.GroupIdentifier = r.GroupIdentifier
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.BusinessStatus = r.BusinessStatus
	m.Intent = r.Intent
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Focus = r.Focus
	m.For = r.For
	m.Encounter = r.Encounter
	m.ExecutionPeriod = r.ExecutionPeriod
	m.AuthoredOn = r.AuthoredOn
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		m.AuthoredOnPrimitiveElement = &primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
	}
	m.LastModified = r.LastModified
	if r.LastModified != nil && (r.LastModified.Id != nil || r.LastModified.Extension != nil) {
		m.LastModifiedPrimitiveElement = &primitiveElement{Id: r.LastModified.Id, Extension: r.LastModified.Extension}
	}
	m.Requester = r.Requester
	m.PerformerType = r.PerformerType
	m.Owner = r.Owner
	m.Location = r.Location
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Insurance = r.Insurance
	m.Note = r.Note
	m.RelevantHistory = r.RelevantHistory
	m.Restriction = r.Restriction
	m.Input = r.Input
	m.Output = r.Output
	return m
}
func (r *Task) UnmarshalJSON(b []byte) error {
	var m jsonTask
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Task) unmarshalJSON(m jsonTask) error {
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
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.InstantiatesCanonical = m.InstantiatesCanonical
	if m.InstantiatesCanonicalPrimitiveElement != nil {
		r.InstantiatesCanonical.Id = m.InstantiatesCanonicalPrimitiveElement.Id
		r.InstantiatesCanonical.Extension = m.InstantiatesCanonicalPrimitiveElement.Extension
	}
	r.InstantiatesUri = m.InstantiatesUri
	if m.InstantiatesUriPrimitiveElement != nil {
		r.InstantiatesUri.Id = m.InstantiatesUriPrimitiveElement.Id
		r.InstantiatesUri.Extension = m.InstantiatesUriPrimitiveElement.Extension
	}
	r.BasedOn = m.BasedOn
	r.GroupIdentifier = m.GroupIdentifier
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.BusinessStatus = m.BusinessStatus
	r.Intent = m.Intent
	if m.IntentPrimitiveElement != nil {
		r.Intent.Id = m.IntentPrimitiveElement.Id
		r.Intent.Extension = m.IntentPrimitiveElement.Extension
	}
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Focus = m.Focus
	r.For = m.For
	r.Encounter = m.Encounter
	r.ExecutionPeriod = m.ExecutionPeriod
	r.AuthoredOn = m.AuthoredOn
	if m.AuthoredOnPrimitiveElement != nil {
		r.AuthoredOn.Id = m.AuthoredOnPrimitiveElement.Id
		r.AuthoredOn.Extension = m.AuthoredOnPrimitiveElement.Extension
	}
	r.LastModified = m.LastModified
	if m.LastModifiedPrimitiveElement != nil {
		r.LastModified.Id = m.LastModifiedPrimitiveElement.Id
		r.LastModified.Extension = m.LastModifiedPrimitiveElement.Extension
	}
	r.Requester = m.Requester
	r.PerformerType = m.PerformerType
	r.Owner = m.Owner
	r.Location = m.Location
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Insurance = m.Insurance
	r.Note = m.Note
	r.RelevantHistory = m.RelevantHistory
	r.Restriction = m.Restriction
	r.Input = m.Input
	r.Output = m.Output
	return nil
}
func (r Task) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
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
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates the number of times the requested action should occur.
	Repetitions *PositiveInt
	// Over what time-period is fulfillment sought.
	Period *Period
	// For requests that are targeted to more than on potential recipient/target, for whom is fulfillment sought?
	Recipient []Reference
}
type jsonTaskRestriction struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Repetitions                 *PositiveInt      `json:"repetitions,omitempty"`
	RepetitionsPrimitiveElement *primitiveElement `json:"_repetitions,omitempty"`
	Period                      *Period           `json:"period,omitempty"`
	Recipient                   []Reference       `json:"recipient,omitempty"`
}

func (r TaskRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TaskRestriction) marshalJSON() jsonTaskRestriction {
	m := jsonTaskRestriction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Repetitions = r.Repetitions
	if r.Repetitions != nil && (r.Repetitions.Id != nil || r.Repetitions.Extension != nil) {
		m.RepetitionsPrimitiveElement = &primitiveElement{Id: r.Repetitions.Id, Extension: r.Repetitions.Extension}
	}
	m.Period = r.Period
	m.Recipient = r.Recipient
	return m
}
func (r *TaskRestriction) UnmarshalJSON(b []byte) error {
	var m jsonTaskRestriction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TaskRestriction) unmarshalJSON(m jsonTaskRestriction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Repetitions = m.Repetitions
	if m.RepetitionsPrimitiveElement != nil {
		r.Repetitions.Id = m.RepetitionsPrimitiveElement.Id
		r.Repetitions.Extension = m.RepetitionsPrimitiveElement.Extension
	}
	r.Period = m.Period
	r.Recipient = m.Recipient
	return nil
}
func (r TaskRestriction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional information that may be needed in the execution of the task.
type TaskInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code or description indicating how the input is intended to be used as part of the task execution.
	Type CodeableConcept
	// The value of the input parameter as a basic type.
	Value isTaskInputValue
}
type isTaskInputValue interface {
	isTaskInputValue()
}

func (r Base64Binary) isTaskInputValue()        {}
func (r Boolean) isTaskInputValue()             {}
func (r Canonical) isTaskInputValue()           {}
func (r Code) isTaskInputValue()                {}
func (r Date) isTaskInputValue()                {}
func (r DateTime) isTaskInputValue()            {}
func (r Decimal) isTaskInputValue()             {}
func (r Id) isTaskInputValue()                  {}
func (r Instant) isTaskInputValue()             {}
func (r Integer) isTaskInputValue()             {}
func (r Markdown) isTaskInputValue()            {}
func (r Oid) isTaskInputValue()                 {}
func (r PositiveInt) isTaskInputValue()         {}
func (r String) isTaskInputValue()              {}
func (r Time) isTaskInputValue()                {}
func (r UnsignedInt) isTaskInputValue()         {}
func (r Uri) isTaskInputValue()                 {}
func (r Url) isTaskInputValue()                 {}
func (r Uuid) isTaskInputValue()                {}
func (r Address) isTaskInputValue()             {}
func (r Age) isTaskInputValue()                 {}
func (r Annotation) isTaskInputValue()          {}
func (r Attachment) isTaskInputValue()          {}
func (r CodeableConcept) isTaskInputValue()     {}
func (r Coding) isTaskInputValue()              {}
func (r ContactPoint) isTaskInputValue()        {}
func (r Count) isTaskInputValue()               {}
func (r Distance) isTaskInputValue()            {}
func (r Duration) isTaskInputValue()            {}
func (r HumanName) isTaskInputValue()           {}
func (r Identifier) isTaskInputValue()          {}
func (r Money) isTaskInputValue()               {}
func (r Period) isTaskInputValue()              {}
func (r Quantity) isTaskInputValue()            {}
func (r Range) isTaskInputValue()               {}
func (r Ratio) isTaskInputValue()               {}
func (r Reference) isTaskInputValue()           {}
func (r SampledData) isTaskInputValue()         {}
func (r Signature) isTaskInputValue()           {}
func (r Timing) isTaskInputValue()              {}
func (r ContactDetail) isTaskInputValue()       {}
func (r Contributor) isTaskInputValue()         {}
func (r DataRequirement) isTaskInputValue()     {}
func (r Expression) isTaskInputValue()          {}
func (r ParameterDefinition) isTaskInputValue() {}
func (r RelatedArtifact) isTaskInputValue()     {}
func (r TriggerDefinition) isTaskInputValue()   {}
func (r UsageContext) isTaskInputValue()        {}
func (r Dosage) isTaskInputValue()              {}
func (r Meta) isTaskInputValue()                {}

type jsonTaskInput struct {
	Id                                *string              `json:"id,omitempty"`
	Extension                         []Extension          `json:"extension,omitempty"`
	ModifierExtension                 []Extension          `json:"modifierExtension,omitempty"`
	Type                              CodeableConcept      `json:"type,omitempty"`
	ValueBase64Binary                 *Base64Binary        `json:"valueBase64Binary,omitempty"`
	ValueBase64BinaryPrimitiveElement *primitiveElement    `json:"_valueBase64Binary,omitempty"`
	ValueBoolean                      *Boolean             `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement      *primitiveElement    `json:"_valueBoolean,omitempty"`
	ValueCanonical                    *Canonical           `json:"valueCanonical,omitempty"`
	ValueCanonicalPrimitiveElement    *primitiveElement    `json:"_valueCanonical,omitempty"`
	ValueCode                         *Code                `json:"valueCode,omitempty"`
	ValueCodePrimitiveElement         *primitiveElement    `json:"_valueCode,omitempty"`
	ValueDate                         *Date                `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement         *primitiveElement    `json:"_valueDate,omitempty"`
	ValueDateTime                     *DateTime            `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement     *primitiveElement    `json:"_valueDateTime,omitempty"`
	ValueDecimal                      *Decimal             `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement      *primitiveElement    `json:"_valueDecimal,omitempty"`
	ValueId                           *Id                  `json:"valueId,omitempty"`
	ValueIdPrimitiveElement           *primitiveElement    `json:"_valueId,omitempty"`
	ValueInstant                      *Instant             `json:"valueInstant,omitempty"`
	ValueInstantPrimitiveElement      *primitiveElement    `json:"_valueInstant,omitempty"`
	ValueInteger                      *Integer             `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement      *primitiveElement    `json:"_valueInteger,omitempty"`
	ValueMarkdown                     *Markdown            `json:"valueMarkdown,omitempty"`
	ValueMarkdownPrimitiveElement     *primitiveElement    `json:"_valueMarkdown,omitempty"`
	ValueOid                          *Oid                 `json:"valueOid,omitempty"`
	ValueOidPrimitiveElement          *primitiveElement    `json:"_valueOid,omitempty"`
	ValuePositiveInt                  *PositiveInt         `json:"valuePositiveInt,omitempty"`
	ValuePositiveIntPrimitiveElement  *primitiveElement    `json:"_valuePositiveInt,omitempty"`
	ValueString                       *String              `json:"valueString,omitempty"`
	ValueStringPrimitiveElement       *primitiveElement    `json:"_valueString,omitempty"`
	ValueTime                         *Time                `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement         *primitiveElement    `json:"_valueTime,omitempty"`
	ValueUnsignedInt                  *UnsignedInt         `json:"valueUnsignedInt,omitempty"`
	ValueUnsignedIntPrimitiveElement  *primitiveElement    `json:"_valueUnsignedInt,omitempty"`
	ValueUri                          *Uri                 `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement          *primitiveElement    `json:"_valueUri,omitempty"`
	ValueUrl                          *Url                 `json:"valueUrl,omitempty"`
	ValueUrlPrimitiveElement          *primitiveElement    `json:"_valueUrl,omitempty"`
	ValueUuid                         *Uuid                `json:"valueUuid,omitempty"`
	ValueUuidPrimitiveElement         *primitiveElement    `json:"_valueUuid,omitempty"`
	ValueAddress                      *Address             `json:"valueAddress,omitempty"`
	ValueAge                          *Age                 `json:"valueAge,omitempty"`
	ValueAnnotation                   *Annotation          `json:"valueAnnotation,omitempty"`
	ValueAttachment                   *Attachment          `json:"valueAttachment,omitempty"`
	ValueCodeableConcept              *CodeableConcept     `json:"valueCodeableConcept,omitempty"`
	ValueCoding                       *Coding              `json:"valueCoding,omitempty"`
	ValueContactPoint                 *ContactPoint        `json:"valueContactPoint,omitempty"`
	ValueCount                        *Count               `json:"valueCount,omitempty"`
	ValueDistance                     *Distance            `json:"valueDistance,omitempty"`
	ValueDuration                     *Duration            `json:"valueDuration,omitempty"`
	ValueHumanName                    *HumanName           `json:"valueHumanName,omitempty"`
	ValueIdentifier                   *Identifier          `json:"valueIdentifier,omitempty"`
	ValueMoney                        *Money               `json:"valueMoney,omitempty"`
	ValuePeriod                       *Period              `json:"valuePeriod,omitempty"`
	ValueQuantity                     *Quantity            `json:"valueQuantity,omitempty"`
	ValueRange                        *Range               `json:"valueRange,omitempty"`
	ValueRatio                        *Ratio               `json:"valueRatio,omitempty"`
	ValueReference                    *Reference           `json:"valueReference,omitempty"`
	ValueSampledData                  *SampledData         `json:"valueSampledData,omitempty"`
	ValueSignature                    *Signature           `json:"valueSignature,omitempty"`
	ValueTiming                       *Timing              `json:"valueTiming,omitempty"`
	ValueContactDetail                *ContactDetail       `json:"valueContactDetail,omitempty"`
	ValueContributor                  *Contributor         `json:"valueContributor,omitempty"`
	ValueDataRequirement              *DataRequirement     `json:"valueDataRequirement,omitempty"`
	ValueExpression                   *Expression          `json:"valueExpression,omitempty"`
	ValueParameterDefinition          *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact              *RelatedArtifact     `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition            *TriggerDefinition   `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext                 *UsageContext        `json:"valueUsageContext,omitempty"`
	ValueDosage                       *Dosage              `json:"valueDosage,omitempty"`
	ValueMeta                         *Meta                `json:"valueMeta,omitempty"`
}

func (r TaskInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TaskInput) marshalJSON() jsonTaskInput {
	m := jsonTaskInput{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	switch v := r.Value.(type) {
	case Base64Binary:
		m.ValueBase64Binary = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Base64Binary:
		m.ValueBase64Binary = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
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
	case Canonical:
		m.ValueCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.ValueCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.ValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Code:
		m.ValueCode = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Code:
		m.ValueCode = v
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case Id:
		m.ValueId = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Id:
		m.ValueId = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Instant:
		m.ValueInstant = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Instant:
		m.ValueInstant = v
		if v.Id != nil || v.Extension != nil {
			m.ValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case Markdown:
		m.ValueMarkdown = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Markdown:
		m.ValueMarkdown = v
		if v.Id != nil || v.Extension != nil {
			m.ValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Oid:
		m.ValueOid = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Oid:
		m.ValueOid = v
		if v.Id != nil || v.Extension != nil {
			m.ValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case PositiveInt:
		m.ValuePositiveInt = &v
		if v.Id != nil || v.Extension != nil {
			m.ValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *PositiveInt:
		m.ValuePositiveInt = v
		if v.Id != nil || v.Extension != nil {
			m.ValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case UnsignedInt:
		m.ValueUnsignedInt = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *UnsignedInt:
		m.ValueUnsignedInt = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case Url:
		m.ValueUrl = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Url:
		m.ValueUrl = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uuid:
		m.ValueUuid = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uuid:
		m.ValueUuid = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Address:
		m.ValueAddress = &v
	case *Address:
		m.ValueAddress = v
	case Age:
		m.ValueAge = &v
	case *Age:
		m.ValueAge = v
	case Annotation:
		m.ValueAnnotation = &v
	case *Annotation:
		m.ValueAnnotation = v
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case ContactPoint:
		m.ValueContactPoint = &v
	case *ContactPoint:
		m.ValueContactPoint = v
	case Count:
		m.ValueCount = &v
	case *Count:
		m.ValueCount = v
	case Distance:
		m.ValueDistance = &v
	case *Distance:
		m.ValueDistance = v
	case Duration:
		m.ValueDuration = &v
	case *Duration:
		m.ValueDuration = v
	case HumanName:
		m.ValueHumanName = &v
	case *HumanName:
		m.ValueHumanName = v
	case Identifier:
		m.ValueIdentifier = &v
	case *Identifier:
		m.ValueIdentifier = v
	case Money:
		m.ValueMoney = &v
	case *Money:
		m.ValueMoney = v
	case Period:
		m.ValuePeriod = &v
	case *Period:
		m.ValuePeriod = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Ratio:
		m.ValueRatio = &v
	case *Ratio:
		m.ValueRatio = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	case SampledData:
		m.ValueSampledData = &v
	case *SampledData:
		m.ValueSampledData = v
	case Signature:
		m.ValueSignature = &v
	case *Signature:
		m.ValueSignature = v
	case Timing:
		m.ValueTiming = &v
	case *Timing:
		m.ValueTiming = v
	case ContactDetail:
		m.ValueContactDetail = &v
	case *ContactDetail:
		m.ValueContactDetail = v
	case Contributor:
		m.ValueContributor = &v
	case *Contributor:
		m.ValueContributor = v
	case DataRequirement:
		m.ValueDataRequirement = &v
	case *DataRequirement:
		m.ValueDataRequirement = v
	case Expression:
		m.ValueExpression = &v
	case *Expression:
		m.ValueExpression = v
	case ParameterDefinition:
		m.ValueParameterDefinition = &v
	case *ParameterDefinition:
		m.ValueParameterDefinition = v
	case RelatedArtifact:
		m.ValueRelatedArtifact = &v
	case *RelatedArtifact:
		m.ValueRelatedArtifact = v
	case TriggerDefinition:
		m.ValueTriggerDefinition = &v
	case *TriggerDefinition:
		m.ValueTriggerDefinition = v
	case UsageContext:
		m.ValueUsageContext = &v
	case *UsageContext:
		m.ValueUsageContext = v
	case Dosage:
		m.ValueDosage = &v
	case *Dosage:
		m.ValueDosage = v
	case Meta:
		m.ValueMeta = &v
	case *Meta:
		m.ValueMeta = v
	}
	return m
}
func (r *TaskInput) UnmarshalJSON(b []byte) error {
	var m jsonTaskInput
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TaskInput) unmarshalJSON(m jsonTaskInput) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.ValueBase64Binary != nil || m.ValueBase64BinaryPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBase64Binary
		if m.ValueBase64BinaryPrimitiveElement != nil {
			if v == nil {
				v = &Base64Binary{}
			}
			v.Id = m.ValueBase64BinaryPrimitiveElement.Id
			v.Extension = m.ValueBase64BinaryPrimitiveElement.Extension
		}
		r.Value = v
	}
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
	if m.ValueCanonical != nil || m.ValueCanonicalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCanonical
		if m.ValueCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.ValueCanonicalPrimitiveElement.Id
			v.Extension = m.ValueCanonicalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueCode != nil || m.ValueCodePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCode
		if m.ValueCodePrimitiveElement != nil {
			if v == nil {
				v = &Code{}
			}
			v.Id = m.ValueCodePrimitiveElement.Id
			v.Extension = m.ValueCodePrimitiveElement.Extension
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
	if m.ValueId != nil || m.ValueIdPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueId
		if m.ValueIdPrimitiveElement != nil {
			if v == nil {
				v = &Id{}
			}
			v.Id = m.ValueIdPrimitiveElement.Id
			v.Extension = m.ValueIdPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInstant != nil || m.ValueInstantPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInstant
		if m.ValueInstantPrimitiveElement != nil {
			if v == nil {
				v = &Instant{}
			}
			v.Id = m.ValueInstantPrimitiveElement.Id
			v.Extension = m.ValueInstantPrimitiveElement.Extension
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
	if m.ValueMarkdown != nil || m.ValueMarkdownPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMarkdown
		if m.ValueMarkdownPrimitiveElement != nil {
			if v == nil {
				v = &Markdown{}
			}
			v.Id = m.ValueMarkdownPrimitiveElement.Id
			v.Extension = m.ValueMarkdownPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueOid != nil || m.ValueOidPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueOid
		if m.ValueOidPrimitiveElement != nil {
			if v == nil {
				v = &Oid{}
			}
			v.Id = m.ValueOidPrimitiveElement.Id
			v.Extension = m.ValueOidPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValuePositiveInt != nil || m.ValuePositiveIntPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePositiveInt
		if m.ValuePositiveIntPrimitiveElement != nil {
			if v == nil {
				v = &PositiveInt{}
			}
			v.Id = m.ValuePositiveIntPrimitiveElement.Id
			v.Extension = m.ValuePositiveIntPrimitiveElement.Extension
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
	if m.ValueUnsignedInt != nil || m.ValueUnsignedIntPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUnsignedInt
		if m.ValueUnsignedIntPrimitiveElement != nil {
			if v == nil {
				v = &UnsignedInt{}
			}
			v.Id = m.ValueUnsignedIntPrimitiveElement.Id
			v.Extension = m.ValueUnsignedIntPrimitiveElement.Extension
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
	if m.ValueUrl != nil || m.ValueUrlPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUrl
		if m.ValueUrlPrimitiveElement != nil {
			if v == nil {
				v = &Url{}
			}
			v.Id = m.ValueUrlPrimitiveElement.Id
			v.Extension = m.ValueUrlPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUuid != nil || m.ValueUuidPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUuid
		if m.ValueUuidPrimitiveElement != nil {
			if v == nil {
				v = &Uuid{}
			}
			v.Id = m.ValueUuidPrimitiveElement.Id
			v.Extension = m.ValueUuidPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueAddress != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAddress
		r.Value = v
	}
	if m.ValueAge != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAge
		r.Value = v
	}
	if m.ValueAnnotation != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAnnotation
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueContactPoint != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContactPoint
		r.Value = v
	}
	if m.ValueCount != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCount
		r.Value = v
	}
	if m.ValueDistance != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDistance
		r.Value = v
	}
	if m.ValueDuration != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDuration
		r.Value = v
	}
	if m.ValueHumanName != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueHumanName
		r.Value = v
	}
	if m.ValueIdentifier != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueIdentifier
		r.Value = v
	}
	if m.ValueMoney != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMoney
		r.Value = v
	}
	if m.ValuePeriod != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePeriod
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueRange != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRange
		r.Value = v
	}
	if m.ValueRatio != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRatio
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	if m.ValueSampledData != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSampledData
		r.Value = v
	}
	if m.ValueSignature != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSignature
		r.Value = v
	}
	if m.ValueTiming != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTiming
		r.Value = v
	}
	if m.ValueContactDetail != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContactDetail
		r.Value = v
	}
	if m.ValueContributor != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContributor
		r.Value = v
	}
	if m.ValueDataRequirement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDataRequirement
		r.Value = v
	}
	if m.ValueExpression != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueExpression
		r.Value = v
	}
	if m.ValueParameterDefinition != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueParameterDefinition
		r.Value = v
	}
	if m.ValueRelatedArtifact != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRelatedArtifact
		r.Value = v
	}
	if m.ValueTriggerDefinition != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTriggerDefinition
		r.Value = v
	}
	if m.ValueUsageContext != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUsageContext
		r.Value = v
	}
	if m.ValueDosage != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDosage
		r.Value = v
	}
	if m.ValueMeta != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMeta
		r.Value = v
	}
	return nil
}
func (r TaskInput) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
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
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the Output parameter.
	Type CodeableConcept
	// The value of the Output parameter as a basic type.
	Value isTaskOutputValue
}
type isTaskOutputValue interface {
	isTaskOutputValue()
}

func (r Base64Binary) isTaskOutputValue()        {}
func (r Boolean) isTaskOutputValue()             {}
func (r Canonical) isTaskOutputValue()           {}
func (r Code) isTaskOutputValue()                {}
func (r Date) isTaskOutputValue()                {}
func (r DateTime) isTaskOutputValue()            {}
func (r Decimal) isTaskOutputValue()             {}
func (r Id) isTaskOutputValue()                  {}
func (r Instant) isTaskOutputValue()             {}
func (r Integer) isTaskOutputValue()             {}
func (r Markdown) isTaskOutputValue()            {}
func (r Oid) isTaskOutputValue()                 {}
func (r PositiveInt) isTaskOutputValue()         {}
func (r String) isTaskOutputValue()              {}
func (r Time) isTaskOutputValue()                {}
func (r UnsignedInt) isTaskOutputValue()         {}
func (r Uri) isTaskOutputValue()                 {}
func (r Url) isTaskOutputValue()                 {}
func (r Uuid) isTaskOutputValue()                {}
func (r Address) isTaskOutputValue()             {}
func (r Age) isTaskOutputValue()                 {}
func (r Annotation) isTaskOutputValue()          {}
func (r Attachment) isTaskOutputValue()          {}
func (r CodeableConcept) isTaskOutputValue()     {}
func (r Coding) isTaskOutputValue()              {}
func (r ContactPoint) isTaskOutputValue()        {}
func (r Count) isTaskOutputValue()               {}
func (r Distance) isTaskOutputValue()            {}
func (r Duration) isTaskOutputValue()            {}
func (r HumanName) isTaskOutputValue()           {}
func (r Identifier) isTaskOutputValue()          {}
func (r Money) isTaskOutputValue()               {}
func (r Period) isTaskOutputValue()              {}
func (r Quantity) isTaskOutputValue()            {}
func (r Range) isTaskOutputValue()               {}
func (r Ratio) isTaskOutputValue()               {}
func (r Reference) isTaskOutputValue()           {}
func (r SampledData) isTaskOutputValue()         {}
func (r Signature) isTaskOutputValue()           {}
func (r Timing) isTaskOutputValue()              {}
func (r ContactDetail) isTaskOutputValue()       {}
func (r Contributor) isTaskOutputValue()         {}
func (r DataRequirement) isTaskOutputValue()     {}
func (r Expression) isTaskOutputValue()          {}
func (r ParameterDefinition) isTaskOutputValue() {}
func (r RelatedArtifact) isTaskOutputValue()     {}
func (r TriggerDefinition) isTaskOutputValue()   {}
func (r UsageContext) isTaskOutputValue()        {}
func (r Dosage) isTaskOutputValue()              {}
func (r Meta) isTaskOutputValue()                {}

type jsonTaskOutput struct {
	Id                                *string              `json:"id,omitempty"`
	Extension                         []Extension          `json:"extension,omitempty"`
	ModifierExtension                 []Extension          `json:"modifierExtension,omitempty"`
	Type                              CodeableConcept      `json:"type,omitempty"`
	ValueBase64Binary                 *Base64Binary        `json:"valueBase64Binary,omitempty"`
	ValueBase64BinaryPrimitiveElement *primitiveElement    `json:"_valueBase64Binary,omitempty"`
	ValueBoolean                      *Boolean             `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement      *primitiveElement    `json:"_valueBoolean,omitempty"`
	ValueCanonical                    *Canonical           `json:"valueCanonical,omitempty"`
	ValueCanonicalPrimitiveElement    *primitiveElement    `json:"_valueCanonical,omitempty"`
	ValueCode                         *Code                `json:"valueCode,omitempty"`
	ValueCodePrimitiveElement         *primitiveElement    `json:"_valueCode,omitempty"`
	ValueDate                         *Date                `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement         *primitiveElement    `json:"_valueDate,omitempty"`
	ValueDateTime                     *DateTime            `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement     *primitiveElement    `json:"_valueDateTime,omitempty"`
	ValueDecimal                      *Decimal             `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement      *primitiveElement    `json:"_valueDecimal,omitempty"`
	ValueId                           *Id                  `json:"valueId,omitempty"`
	ValueIdPrimitiveElement           *primitiveElement    `json:"_valueId,omitempty"`
	ValueInstant                      *Instant             `json:"valueInstant,omitempty"`
	ValueInstantPrimitiveElement      *primitiveElement    `json:"_valueInstant,omitempty"`
	ValueInteger                      *Integer             `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement      *primitiveElement    `json:"_valueInteger,omitempty"`
	ValueMarkdown                     *Markdown            `json:"valueMarkdown,omitempty"`
	ValueMarkdownPrimitiveElement     *primitiveElement    `json:"_valueMarkdown,omitempty"`
	ValueOid                          *Oid                 `json:"valueOid,omitempty"`
	ValueOidPrimitiveElement          *primitiveElement    `json:"_valueOid,omitempty"`
	ValuePositiveInt                  *PositiveInt         `json:"valuePositiveInt,omitempty"`
	ValuePositiveIntPrimitiveElement  *primitiveElement    `json:"_valuePositiveInt,omitempty"`
	ValueString                       *String              `json:"valueString,omitempty"`
	ValueStringPrimitiveElement       *primitiveElement    `json:"_valueString,omitempty"`
	ValueTime                         *Time                `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement         *primitiveElement    `json:"_valueTime,omitempty"`
	ValueUnsignedInt                  *UnsignedInt         `json:"valueUnsignedInt,omitempty"`
	ValueUnsignedIntPrimitiveElement  *primitiveElement    `json:"_valueUnsignedInt,omitempty"`
	ValueUri                          *Uri                 `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement          *primitiveElement    `json:"_valueUri,omitempty"`
	ValueUrl                          *Url                 `json:"valueUrl,omitempty"`
	ValueUrlPrimitiveElement          *primitiveElement    `json:"_valueUrl,omitempty"`
	ValueUuid                         *Uuid                `json:"valueUuid,omitempty"`
	ValueUuidPrimitiveElement         *primitiveElement    `json:"_valueUuid,omitempty"`
	ValueAddress                      *Address             `json:"valueAddress,omitempty"`
	ValueAge                          *Age                 `json:"valueAge,omitempty"`
	ValueAnnotation                   *Annotation          `json:"valueAnnotation,omitempty"`
	ValueAttachment                   *Attachment          `json:"valueAttachment,omitempty"`
	ValueCodeableConcept              *CodeableConcept     `json:"valueCodeableConcept,omitempty"`
	ValueCoding                       *Coding              `json:"valueCoding,omitempty"`
	ValueContactPoint                 *ContactPoint        `json:"valueContactPoint,omitempty"`
	ValueCount                        *Count               `json:"valueCount,omitempty"`
	ValueDistance                     *Distance            `json:"valueDistance,omitempty"`
	ValueDuration                     *Duration            `json:"valueDuration,omitempty"`
	ValueHumanName                    *HumanName           `json:"valueHumanName,omitempty"`
	ValueIdentifier                   *Identifier          `json:"valueIdentifier,omitempty"`
	ValueMoney                        *Money               `json:"valueMoney,omitempty"`
	ValuePeriod                       *Period              `json:"valuePeriod,omitempty"`
	ValueQuantity                     *Quantity            `json:"valueQuantity,omitempty"`
	ValueRange                        *Range               `json:"valueRange,omitempty"`
	ValueRatio                        *Ratio               `json:"valueRatio,omitempty"`
	ValueReference                    *Reference           `json:"valueReference,omitempty"`
	ValueSampledData                  *SampledData         `json:"valueSampledData,omitempty"`
	ValueSignature                    *Signature           `json:"valueSignature,omitempty"`
	ValueTiming                       *Timing              `json:"valueTiming,omitempty"`
	ValueContactDetail                *ContactDetail       `json:"valueContactDetail,omitempty"`
	ValueContributor                  *Contributor         `json:"valueContributor,omitempty"`
	ValueDataRequirement              *DataRequirement     `json:"valueDataRequirement,omitempty"`
	ValueExpression                   *Expression          `json:"valueExpression,omitempty"`
	ValueParameterDefinition          *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact              *RelatedArtifact     `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition            *TriggerDefinition   `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext                 *UsageContext        `json:"valueUsageContext,omitempty"`
	ValueDosage                       *Dosage              `json:"valueDosage,omitempty"`
	ValueMeta                         *Meta                `json:"valueMeta,omitempty"`
}

func (r TaskOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TaskOutput) marshalJSON() jsonTaskOutput {
	m := jsonTaskOutput{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	switch v := r.Value.(type) {
	case Base64Binary:
		m.ValueBase64Binary = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Base64Binary:
		m.ValueBase64Binary = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
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
	case Canonical:
		m.ValueCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.ValueCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.ValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Code:
		m.ValueCode = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Code:
		m.ValueCode = v
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case Id:
		m.ValueId = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Id:
		m.ValueId = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Instant:
		m.ValueInstant = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Instant:
		m.ValueInstant = v
		if v.Id != nil || v.Extension != nil {
			m.ValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case Markdown:
		m.ValueMarkdown = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Markdown:
		m.ValueMarkdown = v
		if v.Id != nil || v.Extension != nil {
			m.ValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Oid:
		m.ValueOid = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Oid:
		m.ValueOid = v
		if v.Id != nil || v.Extension != nil {
			m.ValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case PositiveInt:
		m.ValuePositiveInt = &v
		if v.Id != nil || v.Extension != nil {
			m.ValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *PositiveInt:
		m.ValuePositiveInt = v
		if v.Id != nil || v.Extension != nil {
			m.ValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case UnsignedInt:
		m.ValueUnsignedInt = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *UnsignedInt:
		m.ValueUnsignedInt = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
	case Url:
		m.ValueUrl = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Url:
		m.ValueUrl = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uuid:
		m.ValueUuid = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uuid:
		m.ValueUuid = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Address:
		m.ValueAddress = &v
	case *Address:
		m.ValueAddress = v
	case Age:
		m.ValueAge = &v
	case *Age:
		m.ValueAge = v
	case Annotation:
		m.ValueAnnotation = &v
	case *Annotation:
		m.ValueAnnotation = v
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case ContactPoint:
		m.ValueContactPoint = &v
	case *ContactPoint:
		m.ValueContactPoint = v
	case Count:
		m.ValueCount = &v
	case *Count:
		m.ValueCount = v
	case Distance:
		m.ValueDistance = &v
	case *Distance:
		m.ValueDistance = v
	case Duration:
		m.ValueDuration = &v
	case *Duration:
		m.ValueDuration = v
	case HumanName:
		m.ValueHumanName = &v
	case *HumanName:
		m.ValueHumanName = v
	case Identifier:
		m.ValueIdentifier = &v
	case *Identifier:
		m.ValueIdentifier = v
	case Money:
		m.ValueMoney = &v
	case *Money:
		m.ValueMoney = v
	case Period:
		m.ValuePeriod = &v
	case *Period:
		m.ValuePeriod = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Ratio:
		m.ValueRatio = &v
	case *Ratio:
		m.ValueRatio = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	case SampledData:
		m.ValueSampledData = &v
	case *SampledData:
		m.ValueSampledData = v
	case Signature:
		m.ValueSignature = &v
	case *Signature:
		m.ValueSignature = v
	case Timing:
		m.ValueTiming = &v
	case *Timing:
		m.ValueTiming = v
	case ContactDetail:
		m.ValueContactDetail = &v
	case *ContactDetail:
		m.ValueContactDetail = v
	case Contributor:
		m.ValueContributor = &v
	case *Contributor:
		m.ValueContributor = v
	case DataRequirement:
		m.ValueDataRequirement = &v
	case *DataRequirement:
		m.ValueDataRequirement = v
	case Expression:
		m.ValueExpression = &v
	case *Expression:
		m.ValueExpression = v
	case ParameterDefinition:
		m.ValueParameterDefinition = &v
	case *ParameterDefinition:
		m.ValueParameterDefinition = v
	case RelatedArtifact:
		m.ValueRelatedArtifact = &v
	case *RelatedArtifact:
		m.ValueRelatedArtifact = v
	case TriggerDefinition:
		m.ValueTriggerDefinition = &v
	case *TriggerDefinition:
		m.ValueTriggerDefinition = v
	case UsageContext:
		m.ValueUsageContext = &v
	case *UsageContext:
		m.ValueUsageContext = v
	case Dosage:
		m.ValueDosage = &v
	case *Dosage:
		m.ValueDosage = v
	case Meta:
		m.ValueMeta = &v
	case *Meta:
		m.ValueMeta = v
	}
	return m
}
func (r *TaskOutput) UnmarshalJSON(b []byte) error {
	var m jsonTaskOutput
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TaskOutput) unmarshalJSON(m jsonTaskOutput) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.ValueBase64Binary != nil || m.ValueBase64BinaryPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBase64Binary
		if m.ValueBase64BinaryPrimitiveElement != nil {
			if v == nil {
				v = &Base64Binary{}
			}
			v.Id = m.ValueBase64BinaryPrimitiveElement.Id
			v.Extension = m.ValueBase64BinaryPrimitiveElement.Extension
		}
		r.Value = v
	}
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
	if m.ValueCanonical != nil || m.ValueCanonicalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCanonical
		if m.ValueCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.ValueCanonicalPrimitiveElement.Id
			v.Extension = m.ValueCanonicalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueCode != nil || m.ValueCodePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCode
		if m.ValueCodePrimitiveElement != nil {
			if v == nil {
				v = &Code{}
			}
			v.Id = m.ValueCodePrimitiveElement.Id
			v.Extension = m.ValueCodePrimitiveElement.Extension
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
	if m.ValueId != nil || m.ValueIdPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueId
		if m.ValueIdPrimitiveElement != nil {
			if v == nil {
				v = &Id{}
			}
			v.Id = m.ValueIdPrimitiveElement.Id
			v.Extension = m.ValueIdPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInstant != nil || m.ValueInstantPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInstant
		if m.ValueInstantPrimitiveElement != nil {
			if v == nil {
				v = &Instant{}
			}
			v.Id = m.ValueInstantPrimitiveElement.Id
			v.Extension = m.ValueInstantPrimitiveElement.Extension
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
	if m.ValueMarkdown != nil || m.ValueMarkdownPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMarkdown
		if m.ValueMarkdownPrimitiveElement != nil {
			if v == nil {
				v = &Markdown{}
			}
			v.Id = m.ValueMarkdownPrimitiveElement.Id
			v.Extension = m.ValueMarkdownPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueOid != nil || m.ValueOidPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueOid
		if m.ValueOidPrimitiveElement != nil {
			if v == nil {
				v = &Oid{}
			}
			v.Id = m.ValueOidPrimitiveElement.Id
			v.Extension = m.ValueOidPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValuePositiveInt != nil || m.ValuePositiveIntPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePositiveInt
		if m.ValuePositiveIntPrimitiveElement != nil {
			if v == nil {
				v = &PositiveInt{}
			}
			v.Id = m.ValuePositiveIntPrimitiveElement.Id
			v.Extension = m.ValuePositiveIntPrimitiveElement.Extension
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
	if m.ValueUnsignedInt != nil || m.ValueUnsignedIntPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUnsignedInt
		if m.ValueUnsignedIntPrimitiveElement != nil {
			if v == nil {
				v = &UnsignedInt{}
			}
			v.Id = m.ValueUnsignedIntPrimitiveElement.Id
			v.Extension = m.ValueUnsignedIntPrimitiveElement.Extension
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
	if m.ValueUrl != nil || m.ValueUrlPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUrl
		if m.ValueUrlPrimitiveElement != nil {
			if v == nil {
				v = &Url{}
			}
			v.Id = m.ValueUrlPrimitiveElement.Id
			v.Extension = m.ValueUrlPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUuid != nil || m.ValueUuidPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUuid
		if m.ValueUuidPrimitiveElement != nil {
			if v == nil {
				v = &Uuid{}
			}
			v.Id = m.ValueUuidPrimitiveElement.Id
			v.Extension = m.ValueUuidPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueAddress != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAddress
		r.Value = v
	}
	if m.ValueAge != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAge
		r.Value = v
	}
	if m.ValueAnnotation != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAnnotation
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueContactPoint != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContactPoint
		r.Value = v
	}
	if m.ValueCount != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCount
		r.Value = v
	}
	if m.ValueDistance != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDistance
		r.Value = v
	}
	if m.ValueDuration != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDuration
		r.Value = v
	}
	if m.ValueHumanName != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueHumanName
		r.Value = v
	}
	if m.ValueIdentifier != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueIdentifier
		r.Value = v
	}
	if m.ValueMoney != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMoney
		r.Value = v
	}
	if m.ValuePeriod != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePeriod
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueRange != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRange
		r.Value = v
	}
	if m.ValueRatio != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRatio
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	if m.ValueSampledData != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSampledData
		r.Value = v
	}
	if m.ValueSignature != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSignature
		r.Value = v
	}
	if m.ValueTiming != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTiming
		r.Value = v
	}
	if m.ValueContactDetail != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContactDetail
		r.Value = v
	}
	if m.ValueContributor != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContributor
		r.Value = v
	}
	if m.ValueDataRequirement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDataRequirement
		r.Value = v
	}
	if m.ValueExpression != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueExpression
		r.Value = v
	}
	if m.ValueParameterDefinition != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueParameterDefinition
		r.Value = v
	}
	if m.ValueRelatedArtifact != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRelatedArtifact
		r.Value = v
	}
	if m.ValueTriggerDefinition != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTriggerDefinition
		r.Value = v
	}
	if m.ValueUsageContext != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUsageContext
		r.Value = v
	}
	if m.ValueDosage != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDosage
		r.Value = v
	}
	if m.ValueMeta != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMeta
		r.Value = v
	}
	return nil
}
func (r TaskOutput) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
