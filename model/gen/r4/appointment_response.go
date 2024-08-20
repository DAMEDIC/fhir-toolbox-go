package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
type AppointmentResponse struct {
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
	// This records identifiers associated with this appointment response concern that are defined by business processes and/ or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier []Identifier
	// Appointment that this response is replying to.
	Appointment Reference
	// Date/Time that the appointment is to take place, or requested new start time.
	Start *Instant
	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the end time.
	End *Instant
	// Role of participant in the appointment.
	ParticipantType []CodeableConcept
	// A Person, Location, HealthcareService, or Device that is participating in the appointment.
	Actor *Reference
	// Participation status of the participant. When the status is declined or tentative if the start/end times are different to the appointment, then these times should be interpreted as a requested time change. When the status is accepted, the times can either be the time of the appointment (as a confirmation of the time) or can be empty.
	ParticipantStatus Code
	// Additional comments about the appointment.
	Comment *String
}

func (r AppointmentResponse) ResourceType() string {
	return "AppointmentResponse"
}
func (r AppointmentResponse) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonAppointmentResponse struct {
	ResourceType                      string              `json:"resourceType"`
	Id                                *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement   `json:"_id,omitempty"`
	Meta                              *Meta               `json:"meta,omitempty"`
	ImplicitRules                     *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                          *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement   `json:"_language,omitempty"`
	Text                              *Narrative          `json:"text,omitempty"`
	Contained                         []ContainedResource `json:"contained,omitempty"`
	Extension                         []Extension         `json:"extension,omitempty"`
	ModifierExtension                 []Extension         `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier        `json:"identifier,omitempty"`
	Appointment                       Reference           `json:"appointment,omitempty"`
	Start                             *Instant            `json:"start,omitempty"`
	StartPrimitiveElement             *primitiveElement   `json:"_start,omitempty"`
	End                               *Instant            `json:"end,omitempty"`
	EndPrimitiveElement               *primitiveElement   `json:"_end,omitempty"`
	ParticipantType                   []CodeableConcept   `json:"participantType,omitempty"`
	Actor                             *Reference          `json:"actor,omitempty"`
	ParticipantStatus                 Code                `json:"participantStatus,omitempty"`
	ParticipantStatusPrimitiveElement *primitiveElement   `json:"_participantStatus,omitempty"`
	Comment                           *String             `json:"comment,omitempty"`
	CommentPrimitiveElement           *primitiveElement   `json:"_comment,omitempty"`
}

func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AppointmentResponse) marshalJSON() jsonAppointmentResponse {
	m := jsonAppointmentResponse{}
	m.ResourceType = "AppointmentResponse"
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
	m.Appointment = r.Appointment
	m.Start = r.Start
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	m.End = r.End
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	m.ParticipantType = r.ParticipantType
	m.Actor = r.Actor
	m.ParticipantStatus = r.ParticipantStatus
	if r.ParticipantStatus.Id != nil || r.ParticipantStatus.Extension != nil {
		m.ParticipantStatusPrimitiveElement = &primitiveElement{Id: r.ParticipantStatus.Id, Extension: r.ParticipantStatus.Extension}
	}
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	return m
}
func (r *AppointmentResponse) UnmarshalJSON(b []byte) error {
	var m jsonAppointmentResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AppointmentResponse) unmarshalJSON(m jsonAppointmentResponse) error {
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
	r.Appointment = m.Appointment
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
	r.ParticipantType = m.ParticipantType
	r.Actor = m.Actor
	r.ParticipantStatus = m.ParticipantStatus
	if m.ParticipantStatusPrimitiveElement != nil {
		r.ParticipantStatus.Id = m.ParticipantStatusPrimitiveElement.Id
		r.ParticipantStatus.Extension = m.ParticipantStatusPrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	return nil
}
func (r AppointmentResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
