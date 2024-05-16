package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {
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
	// This records identifiers associated with this appointment concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []Identifier
	// The overall status of the Appointment. Each of the participants has their own participation status which indicates their involvement in the process, however this status indicates the shared status.
	Status Code
	// The coded reason for the appointment being cancelled. This is often used in reporting/billing/futher processing to determine if further actions are required, or specific fees apply.
	CancelationReason *CodeableConcept
	// A broad categorization of the service that is to be performed during this appointment.
	ServiceCategory []CodeableConcept
	// The specific service that is to be performed during this appointment.
	ServiceType []CodeableConcept
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []CodeableConcept
	// The style of appointment or patient that has been booked in the slot (not service type).
	AppointmentType *CodeableConcept
	// The coded reason that this appointment is being scheduled. This is more clinical than administrative.
	ReasonCode []CodeableConcept
	// Reason the appointment has been scheduled to take place, as specified using information from another resource. When the patient arrives and the encounter begins it may be used as the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
	ReasonReference []Reference
	// The priority of the appointment. Can be used to make informed decisions if needing to re-prioritize appointments. (The iCal Standard specifies 0 as undefined, 1 as highest, 9 as lowest priority).
	Priority *UnsignedInt
	// The brief description of the appointment as would be shown on a subject line in a meeting request, or appointment list. Detailed or expanded information should be put in the comment field.
	Description *String
	// Additional information to support the appointment provided when making the appointment.
	SupportingInformation []Reference
	// Date/Time that the appointment is to take place.
	Start *Instant
	// Date/Time that the appointment is to conclude.
	End *Instant
	// Number of minutes that the appointment is to take. This can be less than the duration between the start and end times.  For example, where the actual time of appointment is only an estimate or if a 30 minute appointment is being requested, but any time would work.  Also, if there is, for example, a planned 15 minute break in the middle of a long appointment, the duration may be 15 minutes less than the difference between the start and end.
	MinutesDuration *PositiveInt
	// The slots from the participants' schedules that will be filled by the appointment.
	Slot []Reference
	// The date that this appointment was initially created. This could be different to the meta.lastModified value on the initial entry, as this could have been before the resource was created on the FHIR server, and should remain unchanged over the lifespan of the appointment.
	Created *DateTime
	// Additional comments about the appointment.
	Comment *String
	// While Appointment.comment contains information for internal use, Appointment.patientInstructions is used to capture patient facing information about the Appointment (e.g. please bring your referral or fast from 8pm night before).
	PatientInstruction *String
	// The service request this appointment is allocated to assess (e.g. incoming referral or procedure request).
	BasedOn []Reference
	// List of participants involved in the appointment.
	Participant []AppointmentParticipant
	// A set of date ranges (potentially including times) that the appointment is preferred to be scheduled within.
	//
	// The duration (usually in minutes) could also be provided to indicate the length of the appointment to fill and populate the start/end times for the actual allocated time. However, in other situations the duration may be calculated by the scheduling system.
	RequestedPeriod []Period
}

func (r Appointment) ResourceType() string {
	return "Appointment"
}
func (r Appointment) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonAppointment struct {
	ResourceType                       string                   `json:"resourceType"`
	Id                                 *Id                      `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement        `json:"_id,omitempty"`
	Meta                               *Meta                    `json:"meta,omitempty"`
	ImplicitRules                      *Uri                     `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement        `json:"_implicitRules,omitempty"`
	Language                           *Code                    `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement        `json:"_language,omitempty"`
	Text                               *Narrative               `json:"text,omitempty"`
	Contained                          []containedResource      `json:"contained,omitempty"`
	Extension                          []Extension              `json:"extension,omitempty"`
	ModifierExtension                  []Extension              `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier             `json:"identifier,omitempty"`
	Status                             Code                     `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement        `json:"_status,omitempty"`
	CancelationReason                  *CodeableConcept         `json:"cancelationReason,omitempty"`
	ServiceCategory                    []CodeableConcept        `json:"serviceCategory,omitempty"`
	ServiceType                        []CodeableConcept        `json:"serviceType,omitempty"`
	Specialty                          []CodeableConcept        `json:"specialty,omitempty"`
	AppointmentType                    *CodeableConcept         `json:"appointmentType,omitempty"`
	ReasonCode                         []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference                    []Reference              `json:"reasonReference,omitempty"`
	Priority                           *UnsignedInt             `json:"priority,omitempty"`
	PriorityPrimitiveElement           *primitiveElement        `json:"_priority,omitempty"`
	Description                        *String                  `json:"description,omitempty"`
	DescriptionPrimitiveElement        *primitiveElement        `json:"_description,omitempty"`
	SupportingInformation              []Reference              `json:"supportingInformation,omitempty"`
	Start                              *Instant                 `json:"start,omitempty"`
	StartPrimitiveElement              *primitiveElement        `json:"_start,omitempty"`
	End                                *Instant                 `json:"end,omitempty"`
	EndPrimitiveElement                *primitiveElement        `json:"_end,omitempty"`
	MinutesDuration                    *PositiveInt             `json:"minutesDuration,omitempty"`
	MinutesDurationPrimitiveElement    *primitiveElement        `json:"_minutesDuration,omitempty"`
	Slot                               []Reference              `json:"slot,omitempty"`
	Created                            *DateTime                `json:"created,omitempty"`
	CreatedPrimitiveElement            *primitiveElement        `json:"_created,omitempty"`
	Comment                            *String                  `json:"comment,omitempty"`
	CommentPrimitiveElement            *primitiveElement        `json:"_comment,omitempty"`
	PatientInstruction                 *String                  `json:"patientInstruction,omitempty"`
	PatientInstructionPrimitiveElement *primitiveElement        `json:"_patientInstruction,omitempty"`
	BasedOn                            []Reference              `json:"basedOn,omitempty"`
	Participant                        []AppointmentParticipant `json:"participant,omitempty"`
	RequestedPeriod                    []Period                 `json:"requestedPeriod,omitempty"`
}

func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Appointment) marshalJSON() jsonAppointment {
	m := jsonAppointment{}
	m.ResourceType = "Appointment"
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
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.CancelationReason = r.CancelationReason
	m.ServiceCategory = r.ServiceCategory
	m.ServiceType = r.ServiceType
	m.Specialty = r.Specialty
	m.AppointmentType = r.AppointmentType
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.SupportingInformation = r.SupportingInformation
	m.Start = r.Start
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	m.End = r.End
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	m.MinutesDuration = r.MinutesDuration
	if r.MinutesDuration != nil && (r.MinutesDuration.Id != nil || r.MinutesDuration.Extension != nil) {
		m.MinutesDurationPrimitiveElement = &primitiveElement{Id: r.MinutesDuration.Id, Extension: r.MinutesDuration.Extension}
	}
	m.Slot = r.Slot
	m.Created = r.Created
	if r.Created != nil && (r.Created.Id != nil || r.Created.Extension != nil) {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.PatientInstruction = r.PatientInstruction
	if r.PatientInstruction != nil && (r.PatientInstruction.Id != nil || r.PatientInstruction.Extension != nil) {
		m.PatientInstructionPrimitiveElement = &primitiveElement{Id: r.PatientInstruction.Id, Extension: r.PatientInstruction.Extension}
	}
	m.BasedOn = r.BasedOn
	m.Participant = r.Participant
	m.RequestedPeriod = r.RequestedPeriod
	return m
}
func (r *Appointment) UnmarshalJSON(b []byte) error {
	var m jsonAppointment
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Appointment) unmarshalJSON(m jsonAppointment) error {
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
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.CancelationReason = m.CancelationReason
	r.ServiceCategory = m.ServiceCategory
	r.ServiceType = m.ServiceType
	r.Specialty = m.Specialty
	r.AppointmentType = m.AppointmentType
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.SupportingInformation = m.SupportingInformation
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	r.MinutesDuration = m.MinutesDuration
	if m.MinutesDurationPrimitiveElement != nil {
		r.MinutesDuration.Id = m.MinutesDurationPrimitiveElement.Id
		r.MinutesDuration.Extension = m.MinutesDurationPrimitiveElement.Extension
	}
	r.Slot = m.Slot
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.PatientInstruction = m.PatientInstruction
	if m.PatientInstructionPrimitiveElement != nil {
		r.PatientInstruction.Id = m.PatientInstructionPrimitiveElement.Id
		r.PatientInstruction.Extension = m.PatientInstructionPrimitiveElement.Extension
	}
	r.BasedOn = m.BasedOn
	r.Participant = m.Participant
	r.RequestedPeriod = m.RequestedPeriod
	return nil
}
func (r Appointment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// List of participants involved in the appointment.
type AppointmentParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Role of participant in the appointment.
	Type []CodeableConcept
	// A Person, Location/HealthcareService or Device that is participating in the appointment.
	Actor *Reference
	// Whether this participant is required to be present at the meeting. This covers a use-case where two doctors need to meet to discuss the results for a specific patient, and the patient is not required to be present.
	Required *Code
	// Participation status of the actor.
	Status Code
	// Participation period of the actor.
	Period *Period
}
type jsonAppointmentParticipant struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Actor                    *Reference        `json:"actor,omitempty"`
	Required                 *Code             `json:"required,omitempty"`
	RequiredPrimitiveElement *primitiveElement `json:"_required,omitempty"`
	Status                   Code              `json:"status,omitempty"`
	StatusPrimitiveElement   *primitiveElement `json:"_status,omitempty"`
	Period                   *Period           `json:"period,omitempty"`
}

func (r AppointmentParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AppointmentParticipant) marshalJSON() jsonAppointmentParticipant {
	m := jsonAppointmentParticipant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Actor = r.Actor
	m.Required = r.Required
	if r.Required != nil && (r.Required.Id != nil || r.Required.Extension != nil) {
		m.RequiredPrimitiveElement = &primitiveElement{Id: r.Required.Id, Extension: r.Required.Extension}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *AppointmentParticipant) UnmarshalJSON(b []byte) error {
	var m jsonAppointmentParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AppointmentParticipant) unmarshalJSON(m jsonAppointmentParticipant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Actor = m.Actor
	r.Required = m.Required
	if m.RequiredPrimitiveElement != nil {
		r.Required.Id = m.RequiredPrimitiveElement.Id
		r.Required.Extension = m.RequiredPrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r AppointmentParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
