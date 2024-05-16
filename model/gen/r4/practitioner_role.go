package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
//
// Need to track services that a healthcare provider is able to provide at an organization's location, and the services that they can perform there.
type PractitionerRole struct {
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
	// Business Identifiers that are specific to a role/location.
	Identifier []Identifier
	// Whether this practitioner role record is in active use.
	Active *Boolean
	// The period during which the person is authorized to act as a practitioner in these role(s) for the organization.
	Period *Period
	// Practitioner that is able to provide the defined services for the organization.
	Practitioner *Reference
	// The organization where the Practitioner performs the roles associated.
	Organization *Reference
	// Roles which this practitioner is authorized to perform for the organization.
	Code []CodeableConcept
	// Specific specialty of the practitioner.
	Specialty []CodeableConcept
	// The location(s) at which this practitioner provides care.
	Location []Reference
	// The list of healthcare services that this worker provides for this role's Organization/Location(s).
	HealthcareService []Reference
	// Contact details that are specific to the role/location/service.
	Telecom []ContactPoint
	// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
	AvailableTime []PractitionerRoleAvailableTime
	// The practitioner is not available or performing this role during this period of time due to the provided reason.
	NotAvailable []PractitionerRoleNotAvailable
	// A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
	AvailabilityExceptions *String
	// Technical endpoints providing access to services operated for the practitioner with this role.
	Endpoint []Reference
}

func (r PractitionerRole) ResourceType() string {
	return "PractitionerRole"
}
func (r PractitionerRole) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonPractitionerRole struct {
	ResourceType                           string                          `json:"resourceType"`
	Id                                     *Id                             `json:"id,omitempty"`
	IdPrimitiveElement                     *primitiveElement               `json:"_id,omitempty"`
	Meta                                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules                          *Uri                            `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement          *primitiveElement               `json:"_implicitRules,omitempty"`
	Language                               *Code                           `json:"language,omitempty"`
	LanguagePrimitiveElement               *primitiveElement               `json:"_language,omitempty"`
	Text                                   *Narrative                      `json:"text,omitempty"`
	Contained                              []containedResource             `json:"contained,omitempty"`
	Extension                              []Extension                     `json:"extension,omitempty"`
	ModifierExtension                      []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                             []Identifier                    `json:"identifier,omitempty"`
	Active                                 *Boolean                        `json:"active,omitempty"`
	ActivePrimitiveElement                 *primitiveElement               `json:"_active,omitempty"`
	Period                                 *Period                         `json:"period,omitempty"`
	Practitioner                           *Reference                      `json:"practitioner,omitempty"`
	Organization                           *Reference                      `json:"organization,omitempty"`
	Code                                   []CodeableConcept               `json:"code,omitempty"`
	Specialty                              []CodeableConcept               `json:"specialty,omitempty"`
	Location                               []Reference                     `json:"location,omitempty"`
	HealthcareService                      []Reference                     `json:"healthcareService,omitempty"`
	Telecom                                []ContactPoint                  `json:"telecom,omitempty"`
	AvailableTime                          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
	NotAvailable                           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions                 *String                         `json:"availabilityExceptions,omitempty"`
	AvailabilityExceptionsPrimitiveElement *primitiveElement               `json:"_availabilityExceptions,omitempty"`
	Endpoint                               []Reference                     `json:"endpoint,omitempty"`
}

func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PractitionerRole) marshalJSON() jsonPractitionerRole {
	m := jsonPractitionerRole{}
	m.ResourceType = "PractitionerRole"
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
	m.Active = r.Active
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Period = r.Period
	m.Practitioner = r.Practitioner
	m.Organization = r.Organization
	m.Code = r.Code
	m.Specialty = r.Specialty
	m.Location = r.Location
	m.HealthcareService = r.HealthcareService
	m.Telecom = r.Telecom
	m.AvailableTime = r.AvailableTime
	m.NotAvailable = r.NotAvailable
	m.AvailabilityExceptions = r.AvailabilityExceptions
	if r.AvailabilityExceptions != nil && (r.AvailabilityExceptions.Id != nil || r.AvailabilityExceptions.Extension != nil) {
		m.AvailabilityExceptionsPrimitiveElement = &primitiveElement{Id: r.AvailabilityExceptions.Id, Extension: r.AvailabilityExceptions.Extension}
	}
	m.Endpoint = r.Endpoint
	return m
}
func (r *PractitionerRole) UnmarshalJSON(b []byte) error {
	var m jsonPractitionerRole
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PractitionerRole) unmarshalJSON(m jsonPractitionerRole) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Period = m.Period
	r.Practitioner = m.Practitioner
	r.Organization = m.Organization
	r.Code = m.Code
	r.Specialty = m.Specialty
	r.Location = m.Location
	r.HealthcareService = m.HealthcareService
	r.Telecom = m.Telecom
	r.AvailableTime = m.AvailableTime
	r.NotAvailable = m.NotAvailable
	r.AvailabilityExceptions = m.AvailabilityExceptions
	if m.AvailabilityExceptionsPrimitiveElement != nil {
		r.AvailabilityExceptions.Id = m.AvailabilityExceptionsPrimitiveElement.Id
		r.AvailabilityExceptions.Extension = m.AvailabilityExceptionsPrimitiveElement.Extension
	}
	r.Endpoint = m.Endpoint
	return nil
}
func (r PractitionerRole) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
type PractitionerRoleAvailableTime struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates which days of the week are available between the start and end Times.
	DaysOfWeek []Code
	// Is this always available? (hence times are irrelevant) e.g. 24 hour service.
	AllDay *Boolean
	// The opening time of day. Note: If the AllDay flag is set, then this time is ignored.
	AvailableStartTime *Time
	// The closing time of day. Note: If the AllDay flag is set, then this time is ignored.
	AvailableEndTime *Time
}
type jsonPractitionerRoleAvailableTime struct {
	Id                                 *string             `json:"id,omitempty"`
	Extension                          []Extension         `json:"extension,omitempty"`
	ModifierExtension                  []Extension         `json:"modifierExtension,omitempty"`
	DaysOfWeek                         []Code              `json:"daysOfWeek,omitempty"`
	DaysOfWeekPrimitiveElement         []*primitiveElement `json:"_daysOfWeek,omitempty"`
	AllDay                             *Boolean            `json:"allDay,omitempty"`
	AllDayPrimitiveElement             *primitiveElement   `json:"_allDay,omitempty"`
	AvailableStartTime                 *Time               `json:"availableStartTime,omitempty"`
	AvailableStartTimePrimitiveElement *primitiveElement   `json:"_availableStartTime,omitempty"`
	AvailableEndTime                   *Time               `json:"availableEndTime,omitempty"`
	AvailableEndTimePrimitiveElement   *primitiveElement   `json:"_availableEndTime,omitempty"`
}

func (r PractitionerRoleAvailableTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PractitionerRoleAvailableTime) marshalJSON() jsonPractitionerRoleAvailableTime {
	m := jsonPractitionerRoleAvailableTime{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.DaysOfWeek = r.DaysOfWeek
	anyDaysOfWeekIdOrExtension := false
	for _, e := range r.DaysOfWeek {
		if e.Id != nil || e.Extension != nil {
			anyDaysOfWeekIdOrExtension = true
			break
		}
	}
	if anyDaysOfWeekIdOrExtension {
		m.DaysOfWeekPrimitiveElement = make([]*primitiveElement, 0, len(r.DaysOfWeek))
		for _, e := range r.DaysOfWeek {
			if e.Id != nil || e.Extension != nil {
				m.DaysOfWeekPrimitiveElement = append(m.DaysOfWeekPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DaysOfWeekPrimitiveElement = append(m.DaysOfWeekPrimitiveElement, nil)
			}
		}
	}
	m.AllDay = r.AllDay
	if r.AllDay != nil && (r.AllDay.Id != nil || r.AllDay.Extension != nil) {
		m.AllDayPrimitiveElement = &primitiveElement{Id: r.AllDay.Id, Extension: r.AllDay.Extension}
	}
	m.AvailableStartTime = r.AvailableStartTime
	if r.AvailableStartTime != nil && (r.AvailableStartTime.Id != nil || r.AvailableStartTime.Extension != nil) {
		m.AvailableStartTimePrimitiveElement = &primitiveElement{Id: r.AvailableStartTime.Id, Extension: r.AvailableStartTime.Extension}
	}
	m.AvailableEndTime = r.AvailableEndTime
	if r.AvailableEndTime != nil && (r.AvailableEndTime.Id != nil || r.AvailableEndTime.Extension != nil) {
		m.AvailableEndTimePrimitiveElement = &primitiveElement{Id: r.AvailableEndTime.Id, Extension: r.AvailableEndTime.Extension}
	}
	return m
}
func (r *PractitionerRoleAvailableTime) UnmarshalJSON(b []byte) error {
	var m jsonPractitionerRoleAvailableTime
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PractitionerRoleAvailableTime) unmarshalJSON(m jsonPractitionerRoleAvailableTime) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DaysOfWeek = m.DaysOfWeek
	for i, e := range m.DaysOfWeekPrimitiveElement {
		if len(r.DaysOfWeek) > i {
			r.DaysOfWeek[i].Id = e.Id
			r.DaysOfWeek[i].Extension = e.Extension
		} else {
			r.DaysOfWeek = append(r.DaysOfWeek, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.AllDay = m.AllDay
	if m.AllDayPrimitiveElement != nil {
		r.AllDay.Id = m.AllDayPrimitiveElement.Id
		r.AllDay.Extension = m.AllDayPrimitiveElement.Extension
	}
	r.AvailableStartTime = m.AvailableStartTime
	if m.AvailableStartTimePrimitiveElement != nil {
		r.AvailableStartTime.Id = m.AvailableStartTimePrimitiveElement.Id
		r.AvailableStartTime.Extension = m.AvailableStartTimePrimitiveElement.Extension
	}
	r.AvailableEndTime = m.AvailableEndTime
	if m.AvailableEndTimePrimitiveElement != nil {
		r.AvailableEndTime.Id = m.AvailableEndTimePrimitiveElement.Id
		r.AvailableEndTime.Extension = m.AvailableEndTimePrimitiveElement.Extension
	}
	return nil
}
func (r PractitionerRoleAvailableTime) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The practitioner is not available or performing this role during this period of time due to the provided reason.
type PractitionerRoleNotAvailable struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The reason that can be presented to the user as to why this time is not available.
	Description String
	// Service is not available (seasonally or for a public holiday) from this date.
	During *Period
}
type jsonPractitionerRoleNotAvailable struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Description                 String            `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	During                      *Period           `json:"during,omitempty"`
}

func (r PractitionerRoleNotAvailable) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PractitionerRoleNotAvailable) marshalJSON() jsonPractitionerRoleNotAvailable {
	m := jsonPractitionerRoleNotAvailable{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description.Id != nil || r.Description.Extension != nil {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.During = r.During
	return m
}
func (r *PractitionerRoleNotAvailable) UnmarshalJSON(b []byte) error {
	var m jsonPractitionerRoleNotAvailable
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PractitionerRoleNotAvailable) unmarshalJSON(m jsonPractitionerRoleNotAvailable) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.During = m.During
	return nil
}
func (r PractitionerRoleNotAvailable) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
