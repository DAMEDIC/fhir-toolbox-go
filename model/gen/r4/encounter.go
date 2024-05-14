package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter struct {
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
	// Identifier(s) by which this encounter is known.
	Identifier []Identifier
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status Code
	// The status history permits the encounter resource to contain the status history without needing to read through the historical versions of the resource, or even have the server store them.
	StatusHistory []EncounterStatusHistory
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class Coding
	// The class history permits the tracking of the encounters transitions without needing to go  through the resource history.  This would be used for a case where an admission starts of as an emergency encounter, then transitions into an inpatient scenario. Doing this and not restarting a new encounter ensures that any lab/diagnostic results can more easily follow the patient and not require re-processing and not get lost or cancelled during a kind of discharge from emergency to inpatient.
	ClassHistory []EncounterClassHistory
	// Specific type of encounter (e.g. e-mail consultation, surgical day-care, skilled nursing, rehabilitation).
	Type []CodeableConcept
	// Broad categorization of the service that is to be provided (e.g. cardiology).
	ServiceType *CodeableConcept
	// Indicates the urgency of the encounter.
	Priority *CodeableConcept
	// The patient or group present at the encounter.
	Subject *Reference
	// Where a specific encounter should be classified as a part of a specific episode(s) of care this field should be used. This association can facilitate grouping of related encounters together for a specific purpose, such as government reporting, issue tracking, association via a common problem.  The association is recorded on the encounter as these are typically created after the episode of care and grouped on entry rather than editing the episode of care to append another encounter to it (the episode of care could span years).
	EpisodeOfCare []Reference
	// The request this encounter satisfies (e.g. incoming referral or procedure request).
	BasedOn []Reference
	// The list of people responsible for providing the service.
	Participant []EncounterParticipant
	// The appointment that scheduled this encounter.
	Appointment []Reference
	// The start and end time of the encounter.
	Period *Period
	// Quantity of time the encounter lasted. This excludes the time during leaves of absence.
	Length *Duration
	// Reason the encounter takes place, expressed as a code. For admissions, this can be used for a coded admission diagnosis.
	ReasonCode []CodeableConcept
	// Reason the encounter takes place, expressed as a code. For admissions, this can be used for a coded admission diagnosis.
	ReasonReference []Reference
	// The list of diagnosis relevant to this encounter.
	Diagnosis []EncounterDiagnosis
	// The set of accounts that may be used for billing for this Encounter.
	Account []Reference
	// Details about the admission to a healthcare service.
	Hospitalization *EncounterHospitalization
	// List of locations where  the patient has been during this encounter.
	Location []EncounterLocation
	// The organization that is primarily responsible for this Encounter's services. This MAY be the same as the organization on the Patient record, however it could be different, such as if the actor performing the services was from an external organization (which may be billed seperately) for an external consultation.  Refer to the example bundle showing an abbreviated set of Encounters for a colonoscopy.
	ServiceProvider *Reference
	// Another Encounter of which this encounter is a part of (administratively or in time).
	PartOf *Reference
}

func (r Encounter) ResourceType() string {
	return "Encounter"
}

type jsonEncounter struct {
	ResourceType                  string                    `json:"resourceType"`
	Id                            *Id                       `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement         `json:"_id,omitempty"`
	Meta                          *Meta                     `json:"meta,omitempty"`
	ImplicitRules                 *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                      *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement         `json:"_language,omitempty"`
	Text                          *Narrative                `json:"text,omitempty"`
	Contained                     []containedResource       `json:"contained,omitempty"`
	Extension                     []Extension               `json:"extension,omitempty"`
	ModifierExtension             []Extension               `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier              `json:"identifier,omitempty"`
	Status                        Code                      `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement         `json:"_status,omitempty"`
	StatusHistory                 []EncounterStatusHistory  `json:"statusHistory,omitempty"`
	Class                         Coding                    `json:"class,omitempty"`
	ClassHistory                  []EncounterClassHistory   `json:"classHistory,omitempty"`
	Type                          []CodeableConcept         `json:"type,omitempty"`
	ServiceType                   *CodeableConcept          `json:"serviceType,omitempty"`
	Priority                      *CodeableConcept          `json:"priority,omitempty"`
	Subject                       *Reference                `json:"subject,omitempty"`
	EpisodeOfCare                 []Reference               `json:"episodeOfCare,omitempty"`
	BasedOn                       []Reference               `json:"basedOn,omitempty"`
	Participant                   []EncounterParticipant    `json:"participant,omitempty"`
	Appointment                   []Reference               `json:"appointment,omitempty"`
	Period                        *Period                   `json:"period,omitempty"`
	Length                        *Duration                 `json:"length,omitempty"`
	ReasonCode                    []CodeableConcept         `json:"reasonCode,omitempty"`
	ReasonReference               []Reference               `json:"reasonReference,omitempty"`
	Diagnosis                     []EncounterDiagnosis      `json:"diagnosis,omitempty"`
	Account                       []Reference               `json:"account,omitempty"`
	Hospitalization               *EncounterHospitalization `json:"hospitalization,omitempty"`
	Location                      []EncounterLocation       `json:"location,omitempty"`
	ServiceProvider               *Reference                `json:"serviceProvider,omitempty"`
	PartOf                        *Reference                `json:"partOf,omitempty"`
}

func (r Encounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Encounter) marshalJSON() jsonEncounter {
	m := jsonEncounter{}
	m.ResourceType = "Encounter"
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
	m.StatusHistory = r.StatusHistory
	m.Class = r.Class
	m.ClassHistory = r.ClassHistory
	m.Type = r.Type
	m.ServiceType = r.ServiceType
	m.Priority = r.Priority
	m.Subject = r.Subject
	m.EpisodeOfCare = r.EpisodeOfCare
	m.BasedOn = r.BasedOn
	m.Participant = r.Participant
	m.Appointment = r.Appointment
	m.Period = r.Period
	m.Length = r.Length
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Diagnosis = r.Diagnosis
	m.Account = r.Account
	m.Hospitalization = r.Hospitalization
	m.Location = r.Location
	m.ServiceProvider = r.ServiceProvider
	m.PartOf = r.PartOf
	return m
}
func (r *Encounter) UnmarshalJSON(b []byte) error {
	var m jsonEncounter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Encounter) unmarshalJSON(m jsonEncounter) error {
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
	r.StatusHistory = m.StatusHistory
	r.Class = m.Class
	r.ClassHistory = m.ClassHistory
	r.Type = m.Type
	r.ServiceType = m.ServiceType
	r.Priority = m.Priority
	r.Subject = m.Subject
	r.EpisodeOfCare = m.EpisodeOfCare
	r.BasedOn = m.BasedOn
	r.Participant = m.Participant
	r.Appointment = m.Appointment
	r.Period = m.Period
	r.Length = m.Length
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Diagnosis = m.Diagnosis
	r.Account = m.Account
	r.Hospitalization = m.Hospitalization
	r.Location = m.Location
	r.ServiceProvider = m.ServiceProvider
	r.PartOf = m.PartOf
	return nil
}
func (r Encounter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The status history permits the encounter resource to contain the status history without needing to read through the historical versions of the resource, or even have the server store them.
type EncounterStatusHistory struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status Code
	// The time that the episode was in the specified status.
	Period Period
}
type jsonEncounterStatusHistory struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Status                 Code              `json:"status,omitempty"`
	StatusPrimitiveElement *primitiveElement `json:"_status,omitempty"`
	Period                 Period            `json:"period,omitempty"`
}

func (r EncounterStatusHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EncounterStatusHistory) marshalJSON() jsonEncounterStatusHistory {
	m := jsonEncounterStatusHistory{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *EncounterStatusHistory) UnmarshalJSON(b []byte) error {
	var m jsonEncounterStatusHistory
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EncounterStatusHistory) unmarshalJSON(m jsonEncounterStatusHistory) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r EncounterStatusHistory) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The class history permits the tracking of the encounters transitions without needing to go  through the resource history.  This would be used for a case where an admission starts of as an emergency encounter, then transitions into an inpatient scenario. Doing this and not restarting a new encounter ensures that any lab/diagnostic results can more easily follow the patient and not require re-processing and not get lost or cancelled during a kind of discharge from emergency to inpatient.
type EncounterClassHistory struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// inpatient | outpatient | ambulatory | emergency +.
	Class Coding
	// The time that the episode was in the specified class.
	Period Period
}
type jsonEncounterClassHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Class             Coding      `json:"class,omitempty"`
	Period            Period      `json:"period,omitempty"`
}

func (r EncounterClassHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EncounterClassHistory) marshalJSON() jsonEncounterClassHistory {
	m := jsonEncounterClassHistory{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Class = r.Class
	m.Period = r.Period
	return m
}
func (r *EncounterClassHistory) UnmarshalJSON(b []byte) error {
	var m jsonEncounterClassHistory
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EncounterClassHistory) unmarshalJSON(m jsonEncounterClassHistory) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Class = m.Class
	r.Period = m.Period
	return nil
}
func (r EncounterClassHistory) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The list of people responsible for providing the service.
type EncounterParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Role of participant in encounter.
	Type []CodeableConcept
	// The period of time that the specified participant participated in the encounter. These can overlap or be sub-sets of the overall encounter's period.
	Period *Period
	// Persons involved in the encounter other than the patient.
	Individual *Reference
}
type jsonEncounterParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Individual        *Reference        `json:"individual,omitempty"`
}

func (r EncounterParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EncounterParticipant) marshalJSON() jsonEncounterParticipant {
	m := jsonEncounterParticipant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Period = r.Period
	m.Individual = r.Individual
	return m
}
func (r *EncounterParticipant) UnmarshalJSON(b []byte) error {
	var m jsonEncounterParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EncounterParticipant) unmarshalJSON(m jsonEncounterParticipant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Period = m.Period
	r.Individual = m.Individual
	return nil
}
func (r EncounterParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The list of diagnosis relevant to this encounter.
type EncounterDiagnosis struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Reason the encounter takes place, as specified using information from another resource. For admissions, this is the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
	Condition Reference
	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge …).
	Use *CodeableConcept
	// Ranking of the diagnosis (for each role type).
	Rank *PositiveInt
}
type jsonEncounterDiagnosis struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Condition            Reference         `json:"condition,omitempty"`
	Use                  *CodeableConcept  `json:"use,omitempty"`
	Rank                 *PositiveInt      `json:"rank,omitempty"`
	RankPrimitiveElement *primitiveElement `json:"_rank,omitempty"`
}

func (r EncounterDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EncounterDiagnosis) marshalJSON() jsonEncounterDiagnosis {
	m := jsonEncounterDiagnosis{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Condition = r.Condition
	m.Use = r.Use
	m.Rank = r.Rank
	if r.Rank != nil && (r.Rank.Id != nil || r.Rank.Extension != nil) {
		m.RankPrimitiveElement = &primitiveElement{Id: r.Rank.Id, Extension: r.Rank.Extension}
	}
	return m
}
func (r *EncounterDiagnosis) UnmarshalJSON(b []byte) error {
	var m jsonEncounterDiagnosis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EncounterDiagnosis) unmarshalJSON(m jsonEncounterDiagnosis) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Condition = m.Condition
	r.Use = m.Use
	r.Rank = m.Rank
	if m.RankPrimitiveElement != nil {
		r.Rank.Id = m.RankPrimitiveElement.Id
		r.Rank.Extension = m.RankPrimitiveElement.Extension
	}
	return nil
}
func (r EncounterDiagnosis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details about the admission to a healthcare service.
type EncounterHospitalization struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Pre-admission identifier.
	PreAdmissionIdentifier *Identifier
	// The location/organization from which the patient came before admission.
	Origin *Reference
	// From where patient was admitted (physician referral, transfer).
	AdmitSource *CodeableConcept
	// Whether this hospitalization is a readmission and why if known.
	ReAdmission *CodeableConcept
	// Diet preferences reported by the patient.
	DietPreference []CodeableConcept
	// Special courtesies (VIP, board member).
	SpecialCourtesy []CodeableConcept
	// Any special requests that have been made for this hospitalization encounter, such as the provision of specific equipment or other things.
	SpecialArrangement []CodeableConcept
	// Location/organization to which the patient is discharged.
	Destination *Reference
	// Category or kind of location after discharge.
	DischargeDisposition *CodeableConcept
}
type jsonEncounterHospitalization struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier       `json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept  `json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `json:"specialArrangement,omitempty"`
	Destination            *Reference        `json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `json:"dischargeDisposition,omitempty"`
}

func (r EncounterHospitalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EncounterHospitalization) marshalJSON() jsonEncounterHospitalization {
	m := jsonEncounterHospitalization{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.PreAdmissionIdentifier = r.PreAdmissionIdentifier
	m.Origin = r.Origin
	m.AdmitSource = r.AdmitSource
	m.ReAdmission = r.ReAdmission
	m.DietPreference = r.DietPreference
	m.SpecialCourtesy = r.SpecialCourtesy
	m.SpecialArrangement = r.SpecialArrangement
	m.Destination = r.Destination
	m.DischargeDisposition = r.DischargeDisposition
	return m
}
func (r *EncounterHospitalization) UnmarshalJSON(b []byte) error {
	var m jsonEncounterHospitalization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EncounterHospitalization) unmarshalJSON(m jsonEncounterHospitalization) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.PreAdmissionIdentifier = m.PreAdmissionIdentifier
	r.Origin = m.Origin
	r.AdmitSource = m.AdmitSource
	r.ReAdmission = m.ReAdmission
	r.DietPreference = m.DietPreference
	r.SpecialCourtesy = m.SpecialCourtesy
	r.SpecialArrangement = m.SpecialArrangement
	r.Destination = m.Destination
	r.DischargeDisposition = m.DischargeDisposition
	return nil
}
func (r EncounterHospitalization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// List of locations where  the patient has been during this encounter.
type EncounterLocation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The location where the encounter takes place.
	Location Reference
	// The status of the participants' presence at the specified location during the period specified. If the participant is no longer at the location, then the period will have an end date/time.
	Status *Code
	// This will be used to specify the required levels (bed/ward/room/etc.) desired to be recorded to simplify either messaging or query.
	PhysicalType *CodeableConcept
	// Time period during which the patient was present at the location.
	Period *Period
}
type jsonEncounterLocation struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Location               Reference         `json:"location,omitempty"`
	Status                 *Code             `json:"status,omitempty"`
	StatusPrimitiveElement *primitiveElement `json:"_status,omitempty"`
	PhysicalType           *CodeableConcept  `json:"physicalType,omitempty"`
	Period                 *Period           `json:"period,omitempty"`
}

func (r EncounterLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EncounterLocation) marshalJSON() jsonEncounterLocation {
	m := jsonEncounterLocation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Location = r.Location
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.PhysicalType = r.PhysicalType
	m.Period = r.Period
	return m
}
func (r *EncounterLocation) UnmarshalJSON(b []byte) error {
	var m jsonEncounterLocation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EncounterLocation) unmarshalJSON(m jsonEncounterLocation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Location = m.Location
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.PhysicalType = m.PhysicalType
	r.Period = m.Period
	return nil
}
func (r EncounterLocation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
