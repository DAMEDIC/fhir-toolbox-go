package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// This records identifiers associated with this appointment concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []Identifier
	// The overall status of the Appointment. Each of the participants has their own participation status which indicates their involvement in the process, however this status indicates the shared status.
	Status Code
	// The coded reason for the appointment being cancelled. This is often used in reporting/billing/futher processing to determine if further actions are required, or specific fees apply.
	CancellationReason *CodeableConcept
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class []CodeableConcept
	// A broad categorization of the service that is to be performed during this appointment.
	ServiceCategory []CodeableConcept
	// The specific service that is to be performed during this appointment.
	ServiceType []CodeableReference
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []CodeableConcept
	// The style of appointment or patient that has been booked in the slot (not service type).
	AppointmentType *CodeableConcept
	// The reason that this appointment is being scheduled. This is more clinical than administrative. This can be coded, or as specified using information from another resource. When the patient arrives and the encounter begins it may be used as the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
	Reason []CodeableReference
	// The priority of the appointment. Can be used to make informed decisions if needing to re-prioritize appointments. (The iCal Standard specifies 0 as undefined, 1 as highest, 9 as lowest priority).
	Priority *CodeableConcept
	// The brief description of the appointment as would be shown on a subject line in a meeting request, or appointment list. Detailed or expanded information should be put in the note field.
	Description *String
	// Appointment replaced by this Appointment in cases where there is a cancellation, the details of the cancellation can be found in the cancellationReason property (on the referenced resource).
	Replaces []Reference
	// Connection details of a virtual service (e.g. conference call).
	VirtualService []VirtualServiceDetail
	// Additional information to support the appointment provided when making the appointment.
	SupportingInformation []Reference
	// The previous appointment in a series of related appointments.
	PreviousAppointment *Reference
	// The originating appointment in a recurring set of related appointments.
	OriginatingAppointment *Reference
	// Date/Time that the appointment is to take place.
	Start *Instant
	// Date/Time that the appointment is to conclude.
	End *Instant
	// Number of minutes that the appointment is to take. This can be less than the duration between the start and end times.  For example, where the actual time of appointment is only an estimate or if a 30 minute appointment is being requested, but any time would work.  Also, if there is, for example, a planned 15 minute break in the middle of a long appointment, the duration may be 15 minutes less than the difference between the start and end.
	MinutesDuration *PositiveInt
	// A set of date ranges (potentially including times) that the appointment is preferred to be scheduled within.
	//
	// The duration (usually in minutes) could also be provided to indicate the length of the appointment to fill and populate the start/end times for the actual allocated time. However, in other situations the duration may be calculated by the scheduling system.
	RequestedPeriod []Period
	// The slots from the participants' schedules that will be filled by the appointment.
	Slot []Reference
	// The set of accounts that is expected to be used for billing the activities that result from this Appointment.
	Account []Reference
	// The date that this appointment was initially created. This could be different to the meta.lastModified value on the initial entry, as this could have been before the resource was created on the FHIR server, and should remain unchanged over the lifespan of the appointment.
	Created *DateTime
	// The date/time describing when the appointment was cancelled.
	CancellationDate *DateTime
	// Additional notes/comments about the appointment.
	Note []Annotation
	// While Appointment.note contains information for internal use, Appointment.patientInstructions is used to capture patient facing information about the Appointment (e.g. please bring your referral or fast from 8pm night before).
	PatientInstruction []CodeableReference
	// The request this appointment is allocated to assess (e.g. incoming referral or procedure request).
	BasedOn []Reference
	// The patient or group associated with the appointment, if they are to be present (usually) then they should also be included in the participant backbone element.
	Subject *Reference
	// List of participants involved in the appointment.
	Participant []AppointmentParticipant
	// The sequence number that identifies a specific appointment in a recurring pattern.
	RecurrenceId *PositiveInt
	// This appointment varies from the recurring pattern.
	OccurrenceChanged *Boolean
	// The details of the recurrence pattern or template that is used to generate recurring appointments.
	RecurrenceTemplate []AppointmentRecurrenceTemplate
}

// List of participants involved in the appointment.
type AppointmentParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Role of participant in the appointment.
	Type []CodeableConcept
	// Participation period of the actor.
	Period *Period
	// The individual, device, location, or service participating in the appointment.
	Actor *Reference
	// Whether this participant is required to be present at the meeting. If false, the participant is optional.
	Required *Boolean
	// Participation status of the actor.
	Status Code
}

// The details of the recurrence pattern or template that is used to generate recurring appointments.
type AppointmentRecurrenceTemplate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The timezone of the recurring appointment occurrences.
	Timezone *CodeableConcept
	// How often the appointment series should recur.
	RecurrenceType CodeableConcept
	// Recurring appointments will not occur after this date.
	LastOccurrenceDate *Date
	// How many appointments are planned in the recurrence.
	OccurrenceCount *PositiveInt
	// The list of specific dates that will have appointments generated.
	OccurrenceDate []Date
	// Information about weekly recurring appointments.
	WeeklyTemplate *AppointmentRecurrenceTemplateWeeklyTemplate
	// Information about monthly recurring appointments.
	MonthlyTemplate *AppointmentRecurrenceTemplateMonthlyTemplate
	// Information about yearly recurring appointments.
	YearlyTemplate *AppointmentRecurrenceTemplateYearlyTemplate
	// Any dates, such as holidays, that should be excluded from the recurrence.
	ExcludingDate []Date
	// Any dates, such as holidays, that should be excluded from the recurrence.
	ExcludingRecurrenceId []PositiveInt
}

// Information about weekly recurring appointments.
type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates that recurring appointments should occur on Mondays.
	Monday *Boolean
	// Indicates that recurring appointments should occur on Tuesdays.
	Tuesday *Boolean
	// Indicates that recurring appointments should occur on Wednesdays.
	Wednesday *Boolean
	// Indicates that recurring appointments should occur on Thursdays.
	Thursday *Boolean
	// Indicates that recurring appointments should occur on Fridays.
	Friday *Boolean
	// Indicates that recurring appointments should occur on Saturdays.
	Saturday *Boolean
	// Indicates that recurring appointments should occur on Sundays.
	Sunday *Boolean
	// The interval defines if the recurrence is every nth week. The default is every week, so it is expected that this value will be 2 or more.e.g. For recurring every second week this interval would be 2, or every third week the interval would be 3.
	WeekInterval *PositiveInt
}

// Information about monthly recurring appointments.
type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates that appointments in the series of recurring appointments should occur on a specific day of the month.
	DayOfMonth *PositiveInt
	// Indicates which week within a month the appointments in the series of recurring appointments should occur on.
	NthWeekOfMonth *Coding
	// Indicates which day of the week the recurring appointments should occur each nth week.
	DayOfWeek *Coding
	// Indicates that recurring appointments should occur every nth month.
	MonthInterval PositiveInt
}

// Information about yearly recurring appointments.
type AppointmentRecurrenceTemplateYearlyTemplate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Appointment recurs every nth year.
	YearInterval PositiveInt
}

func (r Appointment) ResourceType() string {
	return "Appointment"
}
func (r Appointment) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r Appointment) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	if r.CancellationReason != nil {
		s += r.CancellationReason.MemSize()
	}
	for _, i := range r.Class {
		s += i.MemSize()
	}
	s += (cap(r.Class) - len(r.Class)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ServiceCategory {
		s += i.MemSize()
	}
	s += (cap(r.ServiceCategory) - len(r.ServiceCategory)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ServiceType {
		s += i.MemSize()
	}
	s += (cap(r.ServiceType) - len(r.ServiceType)) * int(unsafe.Sizeof(CodeableReference{}))
	for _, i := range r.Specialty {
		s += i.MemSize()
	}
	s += (cap(r.Specialty) - len(r.Specialty)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.AppointmentType != nil {
		s += r.AppointmentType.MemSize()
	}
	for _, i := range r.Reason {
		s += i.MemSize()
	}
	s += (cap(r.Reason) - len(r.Reason)) * int(unsafe.Sizeof(CodeableReference{}))
	if r.Priority != nil {
		s += r.Priority.MemSize()
	}
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	for _, i := range r.Replaces {
		s += i.MemSize()
	}
	s += (cap(r.Replaces) - len(r.Replaces)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.VirtualService {
		s += i.MemSize()
	}
	s += (cap(r.VirtualService) - len(r.VirtualService)) * int(unsafe.Sizeof(VirtualServiceDetail{}))
	for _, i := range r.SupportingInformation {
		s += i.MemSize()
	}
	s += (cap(r.SupportingInformation) - len(r.SupportingInformation)) * int(unsafe.Sizeof(Reference{}))
	if r.PreviousAppointment != nil {
		s += r.PreviousAppointment.MemSize()
	}
	if r.OriginatingAppointment != nil {
		s += r.OriginatingAppointment.MemSize()
	}
	if r.Start != nil {
		s += r.Start.MemSize()
	}
	if r.End != nil {
		s += r.End.MemSize()
	}
	if r.MinutesDuration != nil {
		s += r.MinutesDuration.MemSize()
	}
	for _, i := range r.RequestedPeriod {
		s += i.MemSize()
	}
	s += (cap(r.RequestedPeriod) - len(r.RequestedPeriod)) * int(unsafe.Sizeof(Period{}))
	for _, i := range r.Slot {
		s += i.MemSize()
	}
	s += (cap(r.Slot) - len(r.Slot)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Account {
		s += i.MemSize()
	}
	s += (cap(r.Account) - len(r.Account)) * int(unsafe.Sizeof(Reference{}))
	if r.Created != nil {
		s += r.Created.MemSize()
	}
	if r.CancellationDate != nil {
		s += r.CancellationDate.MemSize()
	}
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	for _, i := range r.PatientInstruction {
		s += i.MemSize()
	}
	s += (cap(r.PatientInstruction) - len(r.PatientInstruction)) * int(unsafe.Sizeof(CodeableReference{}))
	for _, i := range r.BasedOn {
		s += i.MemSize()
	}
	s += (cap(r.BasedOn) - len(r.BasedOn)) * int(unsafe.Sizeof(Reference{}))
	if r.Subject != nil {
		s += r.Subject.MemSize()
	}
	for _, i := range r.Participant {
		s += i.MemSize()
	}
	s += (cap(r.Participant) - len(r.Participant)) * int(unsafe.Sizeof(AppointmentParticipant{}))
	if r.RecurrenceId != nil {
		s += r.RecurrenceId.MemSize()
	}
	if r.OccurrenceChanged != nil {
		s += r.OccurrenceChanged.MemSize()
	}
	for _, i := range r.RecurrenceTemplate {
		s += i.MemSize()
	}
	s += (cap(r.RecurrenceTemplate) - len(r.RecurrenceTemplate)) * int(unsafe.Sizeof(AppointmentRecurrenceTemplate{}))
	return s
}
func (r AppointmentParticipant) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Type {
		s += i.MemSize()
	}
	s += (cap(r.Type) - len(r.Type)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Period != nil {
		s += r.Period.MemSize()
	}
	if r.Actor != nil {
		s += r.Actor.MemSize()
	}
	if r.Required != nil {
		s += r.Required.MemSize()
	}
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	return s
}
func (r AppointmentRecurrenceTemplate) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Timezone != nil {
		s += r.Timezone.MemSize()
	}
	s += r.RecurrenceType.MemSize() - int(unsafe.Sizeof(r.RecurrenceType))
	if r.LastOccurrenceDate != nil {
		s += r.LastOccurrenceDate.MemSize()
	}
	if r.OccurrenceCount != nil {
		s += r.OccurrenceCount.MemSize()
	}
	for _, i := range r.OccurrenceDate {
		s += i.MemSize()
	}
	s += (cap(r.OccurrenceDate) - len(r.OccurrenceDate)) * int(unsafe.Sizeof(Date{}))
	if r.WeeklyTemplate != nil {
		s += r.WeeklyTemplate.MemSize()
	}
	if r.MonthlyTemplate != nil {
		s += r.MonthlyTemplate.MemSize()
	}
	if r.YearlyTemplate != nil {
		s += r.YearlyTemplate.MemSize()
	}
	for _, i := range r.ExcludingDate {
		s += i.MemSize()
	}
	s += (cap(r.ExcludingDate) - len(r.ExcludingDate)) * int(unsafe.Sizeof(Date{}))
	for _, i := range r.ExcludingRecurrenceId {
		s += i.MemSize()
	}
	s += (cap(r.ExcludingRecurrenceId) - len(r.ExcludingRecurrenceId)) * int(unsafe.Sizeof(PositiveInt{}))
	return s
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Monday != nil {
		s += r.Monday.MemSize()
	}
	if r.Tuesday != nil {
		s += r.Tuesday.MemSize()
	}
	if r.Wednesday != nil {
		s += r.Wednesday.MemSize()
	}
	if r.Thursday != nil {
		s += r.Thursday.MemSize()
	}
	if r.Friday != nil {
		s += r.Friday.MemSize()
	}
	if r.Saturday != nil {
		s += r.Saturday.MemSize()
	}
	if r.Sunday != nil {
		s += r.Sunday.MemSize()
	}
	if r.WeekInterval != nil {
		s += r.WeekInterval.MemSize()
	}
	return s
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.DayOfMonth != nil {
		s += r.DayOfMonth.MemSize()
	}
	if r.NthWeekOfMonth != nil {
		s += r.NthWeekOfMonth.MemSize()
	}
	if r.DayOfWeek != nil {
		s += r.DayOfWeek.MemSize()
	}
	s += r.MonthInterval.MemSize() - int(unsafe.Sizeof(r.MonthInterval))
	return s
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.YearInterval.MemSize() - int(unsafe.Sizeof(r.YearInterval))
	return s
}
func (r Appointment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AppointmentParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AppointmentRecurrenceTemplate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Appointment) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Appointment) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"Appointment\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CancellationReason != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"cancellationReason\":"))
		if err != nil {
			return err
		}
		err = r.CancellationReason.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Class) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"class\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Class {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ServiceCategory) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"serviceCategory\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ServiceCategory {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ServiceType) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"serviceType\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ServiceType {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Specialty) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"specialty\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Specialty {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.AppointmentType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"appointmentType\":"))
		if err != nil {
			return err
		}
		err = r.AppointmentType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Reason) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reason\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Reason {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Priority != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"priority\":"))
		if err != nil {
			return err
		}
		err = r.Priority.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Replaces) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"replaces\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Replaces {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.VirtualService) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"virtualService\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.VirtualService {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.SupportingInformation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"supportingInformation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SupportingInformation {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.PreviousAppointment != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"previousAppointment\":"))
		if err != nil {
			return err
		}
		err = r.PreviousAppointment.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OriginatingAppointment != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"originatingAppointment\":"))
		if err != nil {
			return err
		}
		err = r.OriginatingAppointment.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && r.Start.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"start\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Start)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		p := primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_start\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.End != nil && r.End.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"end\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.End)
		if err != nil {
			return err
		}
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		p := primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_end\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.MinutesDuration != nil && r.MinutesDuration.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"minutesDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.MinutesDuration)
		if err != nil {
			return err
		}
	}
	if r.MinutesDuration != nil && (r.MinutesDuration.Id != nil || r.MinutesDuration.Extension != nil) {
		p := primitiveElement{Id: r.MinutesDuration.Id, Extension: r.MinutesDuration.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_minutesDuration\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.RequestedPeriod) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"requestedPeriod\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RequestedPeriod {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Slot) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"slot\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Slot {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Account) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"account\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Account {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Created != nil && r.Created.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"created\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Created)
		if err != nil {
			return err
		}
	}
	if r.Created != nil && (r.Created.Id != nil || r.Created.Extension != nil) {
		p := primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_created\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CancellationDate != nil && r.CancellationDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"cancellationDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CancellationDate)
		if err != nil {
			return err
		}
	}
	if r.CancellationDate != nil && (r.CancellationDate.Id != nil || r.CancellationDate.Extension != nil) {
		p := primitiveElement{Id: r.CancellationDate.Id, Extension: r.CancellationDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_cancellationDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Note) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"note\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Note {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.PatientInstruction) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"patientInstruction\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PatientInstruction {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.BasedOn) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"basedOn\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.BasedOn {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Subject != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subject\":"))
		if err != nil {
			return err
		}
		err = r.Subject.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Participant) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"participant\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Participant {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.RecurrenceId != nil && r.RecurrenceId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"recurrenceId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.RecurrenceId)
		if err != nil {
			return err
		}
	}
	if r.RecurrenceId != nil && (r.RecurrenceId.Id != nil || r.RecurrenceId.Extension != nil) {
		p := primitiveElement{Id: r.RecurrenceId.Id, Extension: r.RecurrenceId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_recurrenceId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OccurrenceChanged != nil && r.OccurrenceChanged.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"occurrenceChanged\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.OccurrenceChanged)
		if err != nil {
			return err
		}
	}
	if r.OccurrenceChanged != nil && (r.OccurrenceChanged.Id != nil || r.OccurrenceChanged.Extension != nil) {
		p := primitiveElement{Id: r.OccurrenceChanged.Id, Extension: r.OccurrenceChanged.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_occurrenceChanged\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.RecurrenceTemplate) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"recurrenceTemplate\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RecurrenceTemplate {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentParticipant) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AppointmentParticipant) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Type) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Type {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Period != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"period\":"))
		if err != nil {
			return err
		}
		err = r.Period.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Actor != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"actor\":"))
		if err != nil {
			return err
		}
		err = r.Actor.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Required != nil && r.Required.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"required\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Required)
		if err != nil {
			return err
		}
	}
	if r.Required != nil && (r.Required.Id != nil || r.Required.Extension != nil) {
		p := primitiveElement{Id: r.Required.Id, Extension: r.Required.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_required\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AppointmentRecurrenceTemplate) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Timezone != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timezone\":"))
		if err != nil {
			return err
		}
		err = r.Timezone.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"recurrenceType\":"))
	if err != nil {
		return err
	}
	err = r.RecurrenceType.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.LastOccurrenceDate != nil && r.LastOccurrenceDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastOccurrenceDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastOccurrenceDate)
		if err != nil {
			return err
		}
	}
	if r.LastOccurrenceDate != nil && (r.LastOccurrenceDate.Id != nil || r.LastOccurrenceDate.Extension != nil) {
		p := primitiveElement{Id: r.LastOccurrenceDate.Id, Extension: r.LastOccurrenceDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastOccurrenceDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OccurrenceCount != nil && r.OccurrenceCount.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"occurrenceCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.OccurrenceCount)
		if err != nil {
			return err
		}
	}
	if r.OccurrenceCount != nil && (r.OccurrenceCount.Id != nil || r.OccurrenceCount.Extension != nil) {
		p := primitiveElement{Id: r.OccurrenceCount.Id, Extension: r.OccurrenceCount.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_occurrenceCount\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.OccurrenceDate {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"occurrenceDate\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.OccurrenceDate)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.OccurrenceDate {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_occurrenceDate\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.OccurrenceDate {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if r.WeeklyTemplate != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"weeklyTemplate\":"))
		if err != nil {
			return err
		}
		err = r.WeeklyTemplate.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.MonthlyTemplate != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"monthlyTemplate\":"))
		if err != nil {
			return err
		}
		err = r.MonthlyTemplate.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.YearlyTemplate != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"yearlyTemplate\":"))
		if err != nil {
			return err
		}
		err = r.YearlyTemplate.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.ExcludingDate {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"excludingDate\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.ExcludingDate)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.ExcludingDate {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_excludingDate\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.ExcludingDate {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.ExcludingRecurrenceId {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"excludingRecurrenceId\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.ExcludingRecurrenceId)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.ExcludingRecurrenceId {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_excludingRecurrenceId\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.ExcludingRecurrenceId {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Monday != nil && r.Monday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"monday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Monday)
		if err != nil {
			return err
		}
	}
	if r.Monday != nil && (r.Monday.Id != nil || r.Monday.Extension != nil) {
		p := primitiveElement{Id: r.Monday.Id, Extension: r.Monday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_monday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Tuesday != nil && r.Tuesday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"tuesday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Tuesday)
		if err != nil {
			return err
		}
	}
	if r.Tuesday != nil && (r.Tuesday.Id != nil || r.Tuesday.Extension != nil) {
		p := primitiveElement{Id: r.Tuesday.Id, Extension: r.Tuesday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_tuesday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Wednesday != nil && r.Wednesday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"wednesday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Wednesday)
		if err != nil {
			return err
		}
	}
	if r.Wednesday != nil && (r.Wednesday.Id != nil || r.Wednesday.Extension != nil) {
		p := primitiveElement{Id: r.Wednesday.Id, Extension: r.Wednesday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_wednesday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Thursday != nil && r.Thursday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"thursday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Thursday)
		if err != nil {
			return err
		}
	}
	if r.Thursday != nil && (r.Thursday.Id != nil || r.Thursday.Extension != nil) {
		p := primitiveElement{Id: r.Thursday.Id, Extension: r.Thursday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_thursday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Friday != nil && r.Friday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"friday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Friday)
		if err != nil {
			return err
		}
	}
	if r.Friday != nil && (r.Friday.Id != nil || r.Friday.Extension != nil) {
		p := primitiveElement{Id: r.Friday.Id, Extension: r.Friday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_friday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Saturday != nil && r.Saturday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"saturday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Saturday)
		if err != nil {
			return err
		}
	}
	if r.Saturday != nil && (r.Saturday.Id != nil || r.Saturday.Extension != nil) {
		p := primitiveElement{Id: r.Saturday.Id, Extension: r.Saturday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_saturday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Sunday != nil && r.Sunday.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sunday\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Sunday)
		if err != nil {
			return err
		}
	}
	if r.Sunday != nil && (r.Sunday.Id != nil || r.Sunday.Extension != nil) {
		p := primitiveElement{Id: r.Sunday.Id, Extension: r.Sunday.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_sunday\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.WeekInterval != nil && r.WeekInterval.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"weekInterval\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.WeekInterval)
		if err != nil {
			return err
		}
	}
	if r.WeekInterval != nil && (r.WeekInterval.Id != nil || r.WeekInterval.Extension != nil) {
		p := primitiveElement{Id: r.WeekInterval.Id, Extension: r.WeekInterval.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_weekInterval\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.DayOfMonth != nil && r.DayOfMonth.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"dayOfMonth\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DayOfMonth)
		if err != nil {
			return err
		}
	}
	if r.DayOfMonth != nil && (r.DayOfMonth.Id != nil || r.DayOfMonth.Extension != nil) {
		p := primitiveElement{Id: r.DayOfMonth.Id, Extension: r.DayOfMonth.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_dayOfMonth\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.NthWeekOfMonth != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"nthWeekOfMonth\":"))
		if err != nil {
			return err
		}
		err = r.NthWeekOfMonth.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DayOfWeek != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"dayOfWeek\":"))
		if err != nil {
			return err
		}
		err = r.DayOfWeek.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"monthInterval\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.MonthInterval)
		if err != nil {
			return err
		}
	}
	if r.MonthInterval.Id != nil || r.MonthInterval.Extension != nil {
		p := primitiveElement{Id: r.MonthInterval.Id, Extension: r.MonthInterval.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_monthInterval\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"yearInterval\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.YearInterval)
		if err != nil {
			return err
		}
	}
	if r.YearInterval.Id != nil || r.YearInterval.Extension != nil {
		p := primitiveElement{Id: r.YearInterval.Id, Extension: r.YearInterval.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_yearInterval\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *Appointment) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *Appointment) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Appointment element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Appointment element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "cancellationReason":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CancellationReason = &v
		case "class":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Class = append(r.Class, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "serviceCategory":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ServiceCategory = append(r.ServiceCategory, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "serviceType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ServiceType = append(r.ServiceType, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "specialty":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Specialty = append(r.Specialty, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "appointmentType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AppointmentType = &v
		case "reason":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "priority":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Priority = &v
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "replaces":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Replaces = append(r.Replaces, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "virtualService":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v VirtualServiceDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.VirtualService = append(r.VirtualService, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "supportingInformation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SupportingInformation = append(r.SupportingInformation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "previousAppointment":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PreviousAppointment = &v
		case "originatingAppointment":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.OriginatingAppointment = &v
		case "start":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Instant{}
			}
			r.Start.Value = v.Value
		case "_start":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Instant{}
			}
			r.Start.Id = v.Id
			r.Start.Extension = v.Extension
		case "end":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Instant{}
			}
			r.End.Value = v.Value
		case "_end":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Instant{}
			}
			r.End.Id = v.Id
			r.End.Extension = v.Extension
		case "minutesDuration":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.MinutesDuration == nil {
				r.MinutesDuration = &PositiveInt{}
			}
			r.MinutesDuration.Value = v.Value
		case "_minutesDuration":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.MinutesDuration == nil {
				r.MinutesDuration = &PositiveInt{}
			}
			r.MinutesDuration.Id = v.Id
			r.MinutesDuration.Extension = v.Extension
		case "requestedPeriod":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Period
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RequestedPeriod = append(r.RequestedPeriod, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "slot":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Slot = append(r.Slot, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "account":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Account = append(r.Account, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "created":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Created == nil {
				r.Created = &DateTime{}
			}
			r.Created.Value = v.Value
		case "_created":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Created == nil {
				r.Created = &DateTime{}
			}
			r.Created.Id = v.Id
			r.Created.Extension = v.Extension
		case "cancellationDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.CancellationDate == nil {
				r.CancellationDate = &DateTime{}
			}
			r.CancellationDate.Value = v.Value
		case "_cancellationDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.CancellationDate == nil {
				r.CancellationDate = &DateTime{}
			}
			r.CancellationDate.Id = v.Id
			r.CancellationDate.Extension = v.Extension
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Annotation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "patientInstruction":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PatientInstruction = append(r.PatientInstruction, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "basedOn":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "subject":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Subject = &v
		case "participant":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v AppointmentParticipant
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Participant = append(r.Participant, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		case "recurrenceId":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.RecurrenceId == nil {
				r.RecurrenceId = &PositiveInt{}
			}
			r.RecurrenceId.Value = v.Value
		case "_recurrenceId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.RecurrenceId == nil {
				r.RecurrenceId = &PositiveInt{}
			}
			r.RecurrenceId.Id = v.Id
			r.RecurrenceId.Extension = v.Extension
		case "occurrenceChanged":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.OccurrenceChanged == nil {
				r.OccurrenceChanged = &Boolean{}
			}
			r.OccurrenceChanged.Value = v.Value
		case "_occurrenceChanged":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.OccurrenceChanged == nil {
				r.OccurrenceChanged = &Boolean{}
			}
			r.OccurrenceChanged.Id = v.Id
			r.OccurrenceChanged.Extension = v.Extension
		case "recurrenceTemplate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Appointment element", t)
			}
			for d.More() {
				var v AppointmentRecurrenceTemplate
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RecurrenceTemplate = append(r.RecurrenceTemplate, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Appointment element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in Appointment", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Appointment element", t)
	}
	return nil
}
func (r *AppointmentParticipant) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AppointmentParticipant element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AppointmentParticipant element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentParticipant element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentParticipant element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentParticipant element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentParticipant element", t)
			}
		case "type":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentParticipant element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentParticipant element", t)
			}
		case "period":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Period = &v
		case "actor":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Actor = &v
		case "required":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Required == nil {
				r.Required = &Boolean{}
			}
			r.Required.Value = v.Value
		case "_required":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Required == nil {
				r.Required = &Boolean{}
			}
			r.Required.Id = v.Id
			r.Required.Extension = v.Extension
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in AppointmentParticipant", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AppointmentParticipant element", t)
	}
	return nil
}
func (r *AppointmentRecurrenceTemplate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AppointmentRecurrenceTemplate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AppointmentRecurrenceTemplate element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "timezone":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timezone = &v
		case "recurrenceType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.RecurrenceType = v
		case "lastOccurrenceDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastOccurrenceDate == nil {
				r.LastOccurrenceDate = &Date{}
			}
			r.LastOccurrenceDate.Value = v.Value
		case "_lastOccurrenceDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastOccurrenceDate == nil {
				r.LastOccurrenceDate = &Date{}
			}
			r.LastOccurrenceDate.Id = v.Id
			r.LastOccurrenceDate.Extension = v.Extension
		case "occurrenceCount":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.OccurrenceCount == nil {
				r.OccurrenceCount = &PositiveInt{}
			}
			r.OccurrenceCount.Value = v.Value
		case "_occurrenceCount":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.OccurrenceCount == nil {
				r.OccurrenceCount = &PositiveInt{}
			}
			r.OccurrenceCount.Id = v.Id
			r.OccurrenceCount.Extension = v.Extension
		case "occurrenceDate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for i := 0; d.More(); i++ {
				var v Date
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.OccurrenceDate) <= i {
					r.OccurrenceDate = append(r.OccurrenceDate, Date{})
				}
				r.OccurrenceDate[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "_occurrenceDate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.OccurrenceDate) <= i {
					r.OccurrenceDate = append(r.OccurrenceDate, Date{})
				}
				r.OccurrenceDate[i].Id = v.Id
				r.OccurrenceDate[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "weeklyTemplate":
			var v AppointmentRecurrenceTemplateWeeklyTemplate
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.WeeklyTemplate = &v
		case "monthlyTemplate":
			var v AppointmentRecurrenceTemplateMonthlyTemplate
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MonthlyTemplate = &v
		case "yearlyTemplate":
			var v AppointmentRecurrenceTemplateYearlyTemplate
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.YearlyTemplate = &v
		case "excludingDate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for i := 0; d.More(); i++ {
				var v Date
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.ExcludingDate) <= i {
					r.ExcludingDate = append(r.ExcludingDate, Date{})
				}
				r.ExcludingDate[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "_excludingDate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.ExcludingDate) <= i {
					r.ExcludingDate = append(r.ExcludingDate, Date{})
				}
				r.ExcludingDate[i].Id = v.Id
				r.ExcludingDate[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "excludingRecurrenceId":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.ExcludingRecurrenceId) <= i {
					r.ExcludingRecurrenceId = append(r.ExcludingRecurrenceId, PositiveInt{})
				}
				r.ExcludingRecurrenceId[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		case "_excludingRecurrenceId":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplate element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.ExcludingRecurrenceId) <= i {
					r.ExcludingRecurrenceId = append(r.ExcludingRecurrenceId, PositiveInt{})
				}
				r.ExcludingRecurrenceId[i].Id = v.Id
				r.ExcludingRecurrenceId[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplate element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in AppointmentRecurrenceTemplate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AppointmentRecurrenceTemplate element", t)
	}
	return nil
}
func (r *AppointmentRecurrenceTemplateWeeklyTemplate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
			}
		case "monday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Monday == nil {
				r.Monday = &Boolean{}
			}
			r.Monday.Value = v.Value
		case "_monday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Monday == nil {
				r.Monday = &Boolean{}
			}
			r.Monday.Id = v.Id
			r.Monday.Extension = v.Extension
		case "tuesday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Tuesday == nil {
				r.Tuesday = &Boolean{}
			}
			r.Tuesday.Value = v.Value
		case "_tuesday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Tuesday == nil {
				r.Tuesday = &Boolean{}
			}
			r.Tuesday.Id = v.Id
			r.Tuesday.Extension = v.Extension
		case "wednesday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Wednesday == nil {
				r.Wednesday = &Boolean{}
			}
			r.Wednesday.Value = v.Value
		case "_wednesday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Wednesday == nil {
				r.Wednesday = &Boolean{}
			}
			r.Wednesday.Id = v.Id
			r.Wednesday.Extension = v.Extension
		case "thursday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Thursday == nil {
				r.Thursday = &Boolean{}
			}
			r.Thursday.Value = v.Value
		case "_thursday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Thursday == nil {
				r.Thursday = &Boolean{}
			}
			r.Thursday.Id = v.Id
			r.Thursday.Extension = v.Extension
		case "friday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Friday == nil {
				r.Friday = &Boolean{}
			}
			r.Friday.Value = v.Value
		case "_friday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Friday == nil {
				r.Friday = &Boolean{}
			}
			r.Friday.Id = v.Id
			r.Friday.Extension = v.Extension
		case "saturday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Saturday == nil {
				r.Saturday = &Boolean{}
			}
			r.Saturday.Value = v.Value
		case "_saturday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Saturday == nil {
				r.Saturday = &Boolean{}
			}
			r.Saturday.Id = v.Id
			r.Saturday.Extension = v.Extension
		case "sunday":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Sunday == nil {
				r.Sunday = &Boolean{}
			}
			r.Sunday.Value = v.Value
		case "_sunday":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Sunday == nil {
				r.Sunday = &Boolean{}
			}
			r.Sunday.Id = v.Id
			r.Sunday.Extension = v.Extension
		case "weekInterval":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.WeekInterval == nil {
				r.WeekInterval = &PositiveInt{}
			}
			r.WeekInterval.Value = v.Value
		case "_weekInterval":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.WeekInterval == nil {
				r.WeekInterval = &PositiveInt{}
			}
			r.WeekInterval.Id = v.Id
			r.WeekInterval.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in AppointmentRecurrenceTemplateWeeklyTemplate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AppointmentRecurrenceTemplateWeeklyTemplate element", t)
	}
	return nil
}
func (r *AppointmentRecurrenceTemplateMonthlyTemplate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
			}
		case "dayOfMonth":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DayOfMonth == nil {
				r.DayOfMonth = &PositiveInt{}
			}
			r.DayOfMonth.Value = v.Value
		case "_dayOfMonth":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DayOfMonth == nil {
				r.DayOfMonth = &PositiveInt{}
			}
			r.DayOfMonth.Id = v.Id
			r.DayOfMonth.Extension = v.Extension
		case "nthWeekOfMonth":
			var v Coding
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.NthWeekOfMonth = &v
		case "dayOfWeek":
			var v Coding
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DayOfWeek = &v
		case "monthInterval":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.MonthInterval.Value = v.Value
		case "_monthInterval":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MonthInterval.Id = v.Id
			r.MonthInterval.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in AppointmentRecurrenceTemplateMonthlyTemplate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AppointmentRecurrenceTemplateMonthlyTemplate element", t)
	}
	return nil
}
func (r *AppointmentRecurrenceTemplateYearlyTemplate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AppointmentRecurrenceTemplateYearlyTemplate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AppointmentRecurrenceTemplateYearlyTemplate element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplateYearlyTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplateYearlyTemplate element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AppointmentRecurrenceTemplateYearlyTemplate element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AppointmentRecurrenceTemplateYearlyTemplate element", t)
			}
		case "yearInterval":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.YearInterval.Value = v.Value
		case "_yearInterval":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.YearInterval.Id = v.Id
			r.YearInterval.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in AppointmentRecurrenceTemplateYearlyTemplate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AppointmentRecurrenceTemplateYearlyTemplate element", t)
	}
	return nil
}
func (r Appointment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "Appointment"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CancellationReason, xml.StartElement{Name: xml.Name{Local: "cancellationReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Class, xml.StartElement{Name: xml.Name{Local: "class"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ServiceCategory, xml.StartElement{Name: xml.Name{Local: "serviceCategory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ServiceType, xml.StartElement{Name: xml.Name{Local: "serviceType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Specialty, xml.StartElement{Name: xml.Name{Local: "specialty"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AppointmentType, xml.StartElement{Name: xml.Name{Local: "appointmentType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Replaces, xml.StartElement{Name: xml.Name{Local: "replaces"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VirtualService, xml.StartElement{Name: xml.Name{Local: "virtualService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingInformation, xml.StartElement{Name: xml.Name{Local: "supportingInformation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreviousAppointment, xml.StartElement{Name: xml.Name{Local: "previousAppointment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OriginatingAppointment, xml.StartElement{Name: xml.Name{Local: "originatingAppointment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Start, xml.StartElement{Name: xml.Name{Local: "start"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.End, xml.StartElement{Name: xml.Name{Local: "end"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MinutesDuration, xml.StartElement{Name: xml.Name{Local: "minutesDuration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequestedPeriod, xml.StartElement{Name: xml.Name{Local: "requestedPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Slot, xml.StartElement{Name: xml.Name{Local: "slot"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Account, xml.StartElement{Name: xml.Name{Local: "account"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Created, xml.StartElement{Name: xml.Name{Local: "created"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CancellationDate, xml.StartElement{Name: xml.Name{Local: "cancellationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PatientInstruction, xml.StartElement{Name: xml.Name{Local: "patientInstruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BasedOn, xml.StartElement{Name: xml.Name{Local: "basedOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Participant, xml.StartElement{Name: xml.Name{Local: "participant"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RecurrenceId, xml.StartElement{Name: xml.Name{Local: "recurrenceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OccurrenceChanged, xml.StartElement{Name: xml.Name{Local: "occurrenceChanged"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RecurrenceTemplate, xml.StartElement{Name: xml.Name{Local: "recurrenceTemplate"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentParticipant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Actor, xml.StartElement{Name: xml.Name{Local: "actor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Required, xml.StartElement{Name: xml.Name{Local: "required"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Timezone, xml.StartElement{Name: xml.Name{Local: "timezone"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RecurrenceType, xml.StartElement{Name: xml.Name{Local: "recurrenceType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastOccurrenceDate, xml.StartElement{Name: xml.Name{Local: "lastOccurrenceDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OccurrenceCount, xml.StartElement{Name: xml.Name{Local: "occurrenceCount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OccurrenceDate, xml.StartElement{Name: xml.Name{Local: "occurrenceDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WeeklyTemplate, xml.StartElement{Name: xml.Name{Local: "weeklyTemplate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MonthlyTemplate, xml.StartElement{Name: xml.Name{Local: "monthlyTemplate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.YearlyTemplate, xml.StartElement{Name: xml.Name{Local: "yearlyTemplate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExcludingDate, xml.StartElement{Name: xml.Name{Local: "excludingDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExcludingRecurrenceId, xml.StartElement{Name: xml.Name{Local: "excludingRecurrenceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Monday, xml.StartElement{Name: xml.Name{Local: "monday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Tuesday, xml.StartElement{Name: xml.Name{Local: "tuesday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Wednesday, xml.StartElement{Name: xml.Name{Local: "wednesday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Thursday, xml.StartElement{Name: xml.Name{Local: "thursday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Friday, xml.StartElement{Name: xml.Name{Local: "friday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Saturday, xml.StartElement{Name: xml.Name{Local: "saturday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sunday, xml.StartElement{Name: xml.Name{Local: "sunday"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WeekInterval, xml.StartElement{Name: xml.Name{Local: "weekInterval"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DayOfMonth, xml.StartElement{Name: xml.Name{Local: "dayOfMonth"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NthWeekOfMonth, xml.StartElement{Name: xml.Name{Local: "nthWeekOfMonth"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DayOfWeek, xml.StartElement{Name: xml.Name{Local: "dayOfWeek"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MonthInterval, xml.StartElement{Name: xml.Name{Local: "monthInterval"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.YearInterval, xml.StartElement{Name: xml.Name{Local: "yearInterval"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Appointment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Narrative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "contained":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, c.Resource)
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "cancellationReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CancellationReason = &v
			case "class":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Class = append(r.Class, v)
			case "serviceCategory":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ServiceCategory = append(r.ServiceCategory, v)
			case "serviceType":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ServiceType = append(r.ServiceType, v)
			case "specialty":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Specialty = append(r.Specialty, v)
			case "appointmentType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AppointmentType = &v
			case "reason":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			case "priority":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Priority = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "replaces":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Replaces = append(r.Replaces, v)
			case "virtualService":
				var v VirtualServiceDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VirtualService = append(r.VirtualService, v)
			case "supportingInformation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingInformation = append(r.SupportingInformation, v)
			case "previousAppointment":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreviousAppointment = &v
			case "originatingAppointment":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OriginatingAppointment = &v
			case "start":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = &v
			case "end":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.End = &v
			case "minutesDuration":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MinutesDuration = &v
			case "requestedPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequestedPeriod = append(r.RequestedPeriod, v)
			case "slot":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Slot = append(r.Slot, v)
			case "account":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Account = append(r.Account, v)
			case "created":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Created = &v
			case "cancellationDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CancellationDate = &v
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "patientInstruction":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PatientInstruction = append(r.PatientInstruction, v)
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "participant":
				var v AppointmentParticipant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Participant = append(r.Participant, v)
			case "recurrenceId":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RecurrenceId = &v
			case "occurrenceChanged":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OccurrenceChanged = &v
			case "recurrenceTemplate":
				var v AppointmentRecurrenceTemplate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RecurrenceTemplate = append(r.RecurrenceTemplate, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AppointmentParticipant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			case "actor":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Actor = &v
			case "required":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Required = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AppointmentRecurrenceTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "timezone":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timezone = &v
			case "recurrenceType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RecurrenceType = v
			case "lastOccurrenceDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastOccurrenceDate = &v
			case "occurrenceCount":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OccurrenceCount = &v
			case "occurrenceDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OccurrenceDate = append(r.OccurrenceDate, v)
			case "weeklyTemplate":
				var v AppointmentRecurrenceTemplateWeeklyTemplate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WeeklyTemplate = &v
			case "monthlyTemplate":
				var v AppointmentRecurrenceTemplateMonthlyTemplate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MonthlyTemplate = &v
			case "yearlyTemplate":
				var v AppointmentRecurrenceTemplateYearlyTemplate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.YearlyTemplate = &v
			case "excludingDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExcludingDate = append(r.ExcludingDate, v)
			case "excludingRecurrenceId":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExcludingRecurrenceId = append(r.ExcludingRecurrenceId, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AppointmentRecurrenceTemplateWeeklyTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "monday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Monday = &v
			case "tuesday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tuesday = &v
			case "wednesday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Wednesday = &v
			case "thursday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Thursday = &v
			case "friday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Friday = &v
			case "saturday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Saturday = &v
			case "sunday":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sunday = &v
			case "weekInterval":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WeekInterval = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AppointmentRecurrenceTemplateMonthlyTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "dayOfMonth":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DayOfMonth = &v
			case "nthWeekOfMonth":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NthWeekOfMonth = &v
			case "dayOfWeek":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DayOfWeek = &v
			case "monthInterval":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MonthInterval = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AppointmentRecurrenceTemplateYearlyTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "yearInterval":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.YearInterval = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Appointment) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "cancellationReason") {
		if r.CancellationReason != nil {
			children = append(children, *r.CancellationReason)
		}
	}
	if len(name) == 0 || slices.Contains(name, "class") {
		for _, v := range r.Class {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "serviceCategory") {
		for _, v := range r.ServiceCategory {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "serviceType") {
		for _, v := range r.ServiceType {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "specialty") {
		for _, v := range r.Specialty {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "appointmentType") {
		if r.AppointmentType != nil {
			children = append(children, *r.AppointmentType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reason") {
		for _, v := range r.Reason {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "priority") {
		if r.Priority != nil {
			children = append(children, *r.Priority)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "replaces") {
		for _, v := range r.Replaces {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "virtualService") {
		for _, v := range r.VirtualService {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "supportingInformation") {
		for _, v := range r.SupportingInformation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "previousAppointment") {
		if r.PreviousAppointment != nil {
			children = append(children, *r.PreviousAppointment)
		}
	}
	if len(name) == 0 || slices.Contains(name, "originatingAppointment") {
		if r.OriginatingAppointment != nil {
			children = append(children, *r.OriginatingAppointment)
		}
	}
	if len(name) == 0 || slices.Contains(name, "start") {
		if r.Start != nil {
			children = append(children, *r.Start)
		}
	}
	if len(name) == 0 || slices.Contains(name, "end") {
		if r.End != nil {
			children = append(children, *r.End)
		}
	}
	if len(name) == 0 || slices.Contains(name, "minutesDuration") {
		if r.MinutesDuration != nil {
			children = append(children, *r.MinutesDuration)
		}
	}
	if len(name) == 0 || slices.Contains(name, "requestedPeriod") {
		for _, v := range r.RequestedPeriod {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "slot") {
		for _, v := range r.Slot {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "account") {
		for _, v := range r.Account {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "created") {
		if r.Created != nil {
			children = append(children, *r.Created)
		}
	}
	if len(name) == 0 || slices.Contains(name, "cancellationDate") {
		if r.CancellationDate != nil {
			children = append(children, *r.CancellationDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "note") {
		for _, v := range r.Note {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "patientInstruction") {
		for _, v := range r.PatientInstruction {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "basedOn") {
		for _, v := range r.BasedOn {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subject") {
		if r.Subject != nil {
			children = append(children, *r.Subject)
		}
	}
	if len(name) == 0 || slices.Contains(name, "participant") {
		for _, v := range r.Participant {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "recurrenceId") {
		if r.RecurrenceId != nil {
			children = append(children, *r.RecurrenceId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "occurrenceChanged") {
		if r.OccurrenceChanged != nil {
			children = append(children, *r.OccurrenceChanged)
		}
	}
	if len(name) == 0 || slices.Contains(name, "recurrenceTemplate") {
		for _, v := range r.RecurrenceTemplate {
			children = append(children, v)
		}
	}
	return children
}
func (r Appointment) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Appointment to Boolean")
}
func (r Appointment) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Appointment to String")
}
func (r Appointment) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Appointment to Integer")
}
func (r Appointment) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Appointment to Decimal")
}
func (r Appointment) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Appointment to Date")
}
func (r Appointment) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Appointment to Time")
}
func (r Appointment) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Appointment to DateTime")
}
func (r Appointment) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Appointment to Quantity")
}
func (r Appointment) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o Appointment
	switch other := other.(type) {
	case Appointment:
		o = other
	case *Appointment:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r Appointment) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o Appointment
	switch other := other.(type) {
	case Appointment:
		o = other
	case *Appointment:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r Appointment) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Narrative",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contained",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Status",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "CancellationReason",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Class",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ServiceCategory",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ServiceType",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableReference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Specialty",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "AppointmentType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Reason",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableReference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Priority",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Description",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Replaces",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "VirtualService",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "VirtualServiceDetail",
				Namespace: "FHIR",
			},
		}, {
			Name: "SupportingInformation",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "PreviousAppointment",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "OriginatingAppointment",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Start",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Instant",
				Namespace: "FHIR",
			},
		}, {
			Name: "End",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Instant",
				Namespace: "FHIR",
			},
		}, {
			Name: "MinutesDuration",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "RequestedPeriod",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}, {
			Name: "Slot",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Account",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Created",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "CancellationDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "Note",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Annotation",
				Namespace: "FHIR",
			},
		}, {
			Name: "PatientInstruction",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableReference",
				Namespace: "FHIR",
			},
		}, {
			Name: "BasedOn",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Subject",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Participant",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "AppointmentParticipant",
				Namespace: "FHIR",
			},
		}, {
			Name: "RecurrenceId",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "OccurrenceChanged",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "RecurrenceTemplate",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "AppointmentRecurrenceTemplate",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "Appointment",
			Namespace: "FHIR",
		},
	}
}
func (r AppointmentParticipant) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		for _, v := range r.Type {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "period") {
		if r.Period != nil {
			children = append(children, *r.Period)
		}
	}
	if len(name) == 0 || slices.Contains(name, "actor") {
		if r.Actor != nil {
			children = append(children, *r.Actor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "required") {
		if r.Required != nil {
			children = append(children, *r.Required)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	return children
}
func (r AppointmentParticipant) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert AppointmentParticipant to Boolean")
}
func (r AppointmentParticipant) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert AppointmentParticipant to String")
}
func (r AppointmentParticipant) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert AppointmentParticipant to Integer")
}
func (r AppointmentParticipant) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert AppointmentParticipant to Decimal")
}
func (r AppointmentParticipant) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert AppointmentParticipant to Date")
}
func (r AppointmentParticipant) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert AppointmentParticipant to Time")
}
func (r AppointmentParticipant) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert AppointmentParticipant to DateTime")
}
func (r AppointmentParticipant) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert AppointmentParticipant to Quantity")
}
func (r AppointmentParticipant) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentParticipant
	switch other := other.(type) {
	case AppointmentParticipant:
		o = other
	case *AppointmentParticipant:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentParticipant) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentParticipant
	switch other := other.(type) {
	case AppointmentParticipant:
		o = other
	case *AppointmentParticipant:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentParticipant) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Period",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}, {
			Name: "Actor",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Required",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Status",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "AppointmentParticipant",
			Namespace: "FHIR",
		},
	}
}
func (r AppointmentRecurrenceTemplate) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timezone") {
		if r.Timezone != nil {
			children = append(children, *r.Timezone)
		}
	}
	if len(name) == 0 || slices.Contains(name, "recurrenceType") {
		children = append(children, r.RecurrenceType)
	}
	if len(name) == 0 || slices.Contains(name, "lastOccurrenceDate") {
		if r.LastOccurrenceDate != nil {
			children = append(children, *r.LastOccurrenceDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "occurrenceCount") {
		if r.OccurrenceCount != nil {
			children = append(children, *r.OccurrenceCount)
		}
	}
	if len(name) == 0 || slices.Contains(name, "occurrenceDate") {
		for _, v := range r.OccurrenceDate {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "weeklyTemplate") {
		if r.WeeklyTemplate != nil {
			children = append(children, *r.WeeklyTemplate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "monthlyTemplate") {
		if r.MonthlyTemplate != nil {
			children = append(children, *r.MonthlyTemplate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "yearlyTemplate") {
		if r.YearlyTemplate != nil {
			children = append(children, *r.YearlyTemplate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "excludingDate") {
		for _, v := range r.ExcludingDate {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "excludingRecurrenceId") {
		for _, v := range r.ExcludingRecurrenceId {
			children = append(children, v)
		}
	}
	return children
}
func (r AppointmentRecurrenceTemplate) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to Boolean")
}
func (r AppointmentRecurrenceTemplate) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to String")
}
func (r AppointmentRecurrenceTemplate) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to Integer")
}
func (r AppointmentRecurrenceTemplate) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to Decimal")
}
func (r AppointmentRecurrenceTemplate) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to Date")
}
func (r AppointmentRecurrenceTemplate) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to Time")
}
func (r AppointmentRecurrenceTemplate) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to DateTime")
}
func (r AppointmentRecurrenceTemplate) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplate to Quantity")
}
func (r AppointmentRecurrenceTemplate) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplate:
		o = other
	case *AppointmentRecurrenceTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplate) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplate:
		o = other
	case *AppointmentRecurrenceTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Timezone",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "RecurrenceType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "LastOccurrenceDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "OccurrenceCount",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "OccurrenceDate",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "WeeklyTemplate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "AppointmentRecurrenceTemplateWeeklyTemplate",
				Namespace: "FHIR",
			},
		}, {
			Name: "MonthlyTemplate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "AppointmentRecurrenceTemplateMonthlyTemplate",
				Namespace: "FHIR",
			},
		}, {
			Name: "YearlyTemplate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "AppointmentRecurrenceTemplateYearlyTemplate",
				Namespace: "FHIR",
			},
		}, {
			Name: "ExcludingDate",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "ExcludingRecurrenceId",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "AppointmentRecurrenceTemplate",
			Namespace: "FHIR",
		},
	}
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "monday") {
		if r.Monday != nil {
			children = append(children, *r.Monday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "tuesday") {
		if r.Tuesday != nil {
			children = append(children, *r.Tuesday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "wednesday") {
		if r.Wednesday != nil {
			children = append(children, *r.Wednesday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "thursday") {
		if r.Thursday != nil {
			children = append(children, *r.Thursday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "friday") {
		if r.Friday != nil {
			children = append(children, *r.Friday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "saturday") {
		if r.Saturday != nil {
			children = append(children, *r.Saturday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sunday") {
		if r.Sunday != nil {
			children = append(children, *r.Sunday)
		}
	}
	if len(name) == 0 || slices.Contains(name, "weekInterval") {
		if r.WeekInterval != nil {
			children = append(children, *r.WeekInterval)
		}
	}
	return children
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to Boolean")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to String")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to Integer")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to Decimal")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to Date")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to Time")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to DateTime")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateWeeklyTemplate to Quantity")
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplateWeeklyTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplateWeeklyTemplate:
		o = other
	case *AppointmentRecurrenceTemplateWeeklyTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplateWeeklyTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplateWeeklyTemplate:
		o = other
	case *AppointmentRecurrenceTemplateWeeklyTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplateWeeklyTemplate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Monday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Tuesday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Wednesday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Thursday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Friday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Saturday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Sunday",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "WeekInterval",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "AppointmentRecurrenceTemplateWeeklyTemplate",
			Namespace: "FHIR",
		},
	}
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dayOfMonth") {
		if r.DayOfMonth != nil {
			children = append(children, *r.DayOfMonth)
		}
	}
	if len(name) == 0 || slices.Contains(name, "nthWeekOfMonth") {
		if r.NthWeekOfMonth != nil {
			children = append(children, *r.NthWeekOfMonth)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dayOfWeek") {
		if r.DayOfWeek != nil {
			children = append(children, *r.DayOfWeek)
		}
	}
	if len(name) == 0 || slices.Contains(name, "monthInterval") {
		children = append(children, r.MonthInterval)
	}
	return children
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to Boolean")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to String")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to Integer")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to Decimal")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to Date")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to Time")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to DateTime")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateMonthlyTemplate to Quantity")
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplateMonthlyTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplateMonthlyTemplate:
		o = other
	case *AppointmentRecurrenceTemplateMonthlyTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplateMonthlyTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplateMonthlyTemplate:
		o = other
	case *AppointmentRecurrenceTemplateMonthlyTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplateMonthlyTemplate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "DayOfMonth",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "NthWeekOfMonth",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Coding",
				Namespace: "FHIR",
			},
		}, {
			Name: "DayOfWeek",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Coding",
				Namespace: "FHIR",
			},
		}, {
			Name: "MonthInterval",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "AppointmentRecurrenceTemplateMonthlyTemplate",
			Namespace: "FHIR",
		},
	}
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "yearInterval") {
		children = append(children, r.YearInterval)
	}
	return children
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to Boolean")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to String")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to Integer")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to Decimal")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to Date")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to Time")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to DateTime")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert AppointmentRecurrenceTemplateYearlyTemplate to Quantity")
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplateYearlyTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplateYearlyTemplate:
		o = other
	case *AppointmentRecurrenceTemplateYearlyTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o AppointmentRecurrenceTemplateYearlyTemplate
	switch other := other.(type) {
	case AppointmentRecurrenceTemplateYearlyTemplate:
		o = other
	case *AppointmentRecurrenceTemplateYearlyTemplate:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r AppointmentRecurrenceTemplateYearlyTemplate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "YearInterval",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "AppointmentRecurrenceTemplateYearlyTemplate",
			Namespace: "FHIR",
		},
	}
}
