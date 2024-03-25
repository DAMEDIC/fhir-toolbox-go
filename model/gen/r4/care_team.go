package r4

import "encoding/json"

// The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient.
type CareTeam struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Business identifiers assigned to this care team by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// Indicates the current state of the care team.
	Status *Code
	// Identifies what kind of team.  This is to support differentiation between multiple co-existing teams, such as care plan team, episode of care team, longitudinal care team.
	Category []CodeableConcept
	// A label for human use intended to distinguish like teams.  E.g. the "red" vs. "green" trauma teams.
	Name *String
	// Identifies the patient or group whose intended care is handled by the team.
	Subject *Reference
	// The Encounter during which this CareTeam was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Indicates when the team did (or is intended to) come into effect and end.
	Period *Period
	// Identifies all people and organizations who are expected to be involved in the care team.
	Participant []CareTeamParticipant
	// Describes why the care team exists.
	ReasonCode []CodeableConcept
	// Condition(s) that this care team addresses.
	ReasonReference []Reference
	// The organization responsible for the care team.
	ManagingOrganization []Reference
	// A central contact detail for the care team (that applies to all members).
	Telecom []ContactPoint
	// Comments made about the CareTeam.
	Note []Annotation
}
type jsonCareTeam struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Text                          *Narrative            `json:"text,omitempty"`
	Contained                     []containedResource   `json:"contained,omitempty"`
	Extension                     []Extension           `json:"extension,omitempty"`
	ModifierExtension             []Extension           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier          `json:"identifier,omitempty"`
	Status                        *Code                 `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement     `json:"_status,omitempty"`
	Category                      []CodeableConcept     `json:"category,omitempty"`
	Name                          *String               `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement     `json:"_name,omitempty"`
	Subject                       *Reference            `json:"subject,omitempty"`
	Encounter                     *Reference            `json:"encounter,omitempty"`
	Period                        *Period               `json:"period,omitempty"`
	Participant                   []CareTeamParticipant `json:"participant,omitempty"`
	ReasonCode                    []CodeableConcept     `json:"reasonCode,omitempty"`
	ReasonReference               []Reference           `json:"reasonReference,omitempty"`
	ManagingOrganization          []Reference           `json:"managingOrganization,omitempty"`
	Telecom                       []ContactPoint        `json:"telecom,omitempty"`
	Note                          []Annotation          `json:"note,omitempty"`
}

func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CareTeam) marshalJSON() jsonCareTeam {
	m := jsonCareTeam{}
	m.ResourceType = "CareTeam"
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
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Category = r.Category
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.Period = r.Period
	m.Participant = r.Participant
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.ManagingOrganization = r.ManagingOrganization
	m.Telecom = r.Telecom
	m.Note = r.Note
	return m
}
func (r *CareTeam) UnmarshalJSON(b []byte) error {
	var m jsonCareTeam
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CareTeam) unmarshalJSON(m jsonCareTeam) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
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
	r.Category = m.Category
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Period = m.Period
	r.Participant = m.Participant
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.ManagingOrganization = m.ManagingOrganization
	r.Telecom = m.Telecom
	r.Note = m.Note
	return nil
}
func (r CareTeam) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies all people and organizations who are expected to be involved in the care team.
type CareTeamParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates specific responsibility of an individual within the care team, such as "Primary care physician", "Trained social worker counselor", "Caregiver", etc.
	Role []CodeableConcept
	// The specific person or organization who is participating/expected to participate in the care team.
	Member *Reference
	// The organization of the practitioner.
	OnBehalfOf *Reference
	// Indicates when the specific member or organization did (or is intended to) come into effect and end.
	Period *Period
}
type jsonCareTeamParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Member            *Reference        `json:"member,omitempty"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
	Period            *Period           `json:"period,omitempty"`
}

func (r CareTeamParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CareTeamParticipant) marshalJSON() jsonCareTeamParticipant {
	m := jsonCareTeamParticipant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Role = r.Role
	m.Member = r.Member
	m.OnBehalfOf = r.OnBehalfOf
	m.Period = r.Period
	return m
}
func (r *CareTeamParticipant) UnmarshalJSON(b []byte) error {
	var m jsonCareTeamParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CareTeamParticipant) unmarshalJSON(m jsonCareTeamParticipant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Role = m.Role
	r.Member = m.Member
	r.OnBehalfOf = m.OnBehalfOf
	r.Period = m.Period
	return nil
}
func (r CareTeamParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
