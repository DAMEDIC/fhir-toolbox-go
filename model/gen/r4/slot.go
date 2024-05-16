package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A slot of time on a schedule that may be available for booking appointments.
type Slot struct {
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
	// External Ids for this item.
	Identifier []Identifier
	// A broad categorization of the service that is to be performed during this appointment.
	ServiceCategory []CodeableConcept
	// The type of appointments that can be booked into this slot (ideally this would be an identifiable service - which is at a location, rather than the location itself). If provided then this overrides the value provided on the availability resource.
	ServiceType []CodeableConcept
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []CodeableConcept
	// The style of appointment or patient that may be booked in the slot (not service type).
	AppointmentType *CodeableConcept
	// The schedule resource that this slot defines an interval of status information.
	Schedule Reference
	// busy | free | busy-unavailable | busy-tentative | entered-in-error.
	Status Code
	// Date/Time that the slot is to begin.
	Start Instant
	// Date/Time that the slot is to conclude.
	End Instant
	// This slot has already been overbooked, appointments are unlikely to be accepted for this time.
	Overbooked *Boolean
	// Comments on the slot to describe any extended information. Such as custom constraints on the slot.
	Comment *String
}

func (r Slot) ResourceType() string {
	return "Slot"
}
func (r Slot) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonSlot struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []containedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	ServiceCategory               []CodeableConcept   `json:"serviceCategory,omitempty"`
	ServiceType                   []CodeableConcept   `json:"serviceType,omitempty"`
	Specialty                     []CodeableConcept   `json:"specialty,omitempty"`
	AppointmentType               *CodeableConcept    `json:"appointmentType,omitempty"`
	Schedule                      Reference           `json:"schedule,omitempty"`
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Start                         Instant             `json:"start,omitempty"`
	StartPrimitiveElement         *primitiveElement   `json:"_start,omitempty"`
	End                           Instant             `json:"end,omitempty"`
	EndPrimitiveElement           *primitiveElement   `json:"_end,omitempty"`
	Overbooked                    *Boolean            `json:"overbooked,omitempty"`
	OverbookedPrimitiveElement    *primitiveElement   `json:"_overbooked,omitempty"`
	Comment                       *String             `json:"comment,omitempty"`
	CommentPrimitiveElement       *primitiveElement   `json:"_comment,omitempty"`
}

func (r Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Slot) marshalJSON() jsonSlot {
	m := jsonSlot{}
	m.ResourceType = "Slot"
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
	m.ServiceCategory = r.ServiceCategory
	m.ServiceType = r.ServiceType
	m.Specialty = r.Specialty
	m.AppointmentType = r.AppointmentType
	m.Schedule = r.Schedule
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Start = r.Start
	if r.Start.Id != nil || r.Start.Extension != nil {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	m.End = r.End
	if r.End.Id != nil || r.End.Extension != nil {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	m.Overbooked = r.Overbooked
	if r.Overbooked != nil && (r.Overbooked.Id != nil || r.Overbooked.Extension != nil) {
		m.OverbookedPrimitiveElement = &primitiveElement{Id: r.Overbooked.Id, Extension: r.Overbooked.Extension}
	}
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	return m
}
func (r *Slot) UnmarshalJSON(b []byte) error {
	var m jsonSlot
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Slot) unmarshalJSON(m jsonSlot) error {
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
	r.ServiceCategory = m.ServiceCategory
	r.ServiceType = m.ServiceType
	r.Specialty = m.Specialty
	r.AppointmentType = m.AppointmentType
	r.Schedule = m.Schedule
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
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
	r.Overbooked = m.Overbooked
	if m.OverbookedPrimitiveElement != nil {
		r.Overbooked.Id = m.OverbookedPrimitiveElement.Id
		r.Overbooked.Extension = m.OverbookedPrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	return nil
}
func (r Slot) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
