package r4

import "encoding/json"

// An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare struct {
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
	// The EpisodeOfCare may be known by different identifiers for different contexts of use, such as when an external agency is tracking the Episode for funding purposes.
	Identifier []Identifier
	// planned | waitlist | active | onhold | finished | cancelled.
	Status Code
	// The history of statuses that the EpisodeOfCare has been through (without requiring processing the history of the resource).
	StatusHistory []EpisodeOfCareStatusHistory
	// A classification of the type of episode of care; e.g. specialist referral, disease management, type of funded care.
	Type []CodeableConcept
	// The list of diagnosis relevant to this episode of care.
	Diagnosis []EpisodeOfCareDiagnosis
	// The patient who is the focus of this episode of care.
	Patient Reference
	// The organization that has assumed the specific responsibilities for the specified duration.
	ManagingOrganization *Reference
	// The interval during which the managing organization assumes the defined responsibility.
	Period *Period
	// Referral Request(s) that are fulfilled by this EpisodeOfCare, incoming referrals.
	ReferralRequest []Reference
	// The practitioner that is the care manager/care coordinator for this patient.
	CareManager *Reference
	// The list of practitioners that may be facilitating this episode of care for specific purposes.
	Team []Reference
	// The set of accounts that may be used for billing for this EpisodeOfCare.
	Account []Reference
}
type jsonEpisodeOfCare struct {
	ResourceType                  string                       `json:"resourceType"`
	Id                            *Id                          `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement            `json:"_id,omitempty"`
	Meta                          *Meta                        `json:"meta,omitempty"`
	ImplicitRules                 *Uri                         `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement            `json:"_implicitRules,omitempty"`
	Language                      *Code                        `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement            `json:"_language,omitempty"`
	Text                          *Narrative                   `json:"text,omitempty"`
	Contained                     []containedResource          `json:"contained,omitempty"`
	Extension                     []Extension                  `json:"extension,omitempty"`
	ModifierExtension             []Extension                  `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                 `json:"identifier,omitempty"`
	Status                        Code                         `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement            `json:"_status,omitempty"`
	StatusHistory                 []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty"`
	Type                          []CodeableConcept            `json:"type,omitempty"`
	Diagnosis                     []EpisodeOfCareDiagnosis     `json:"diagnosis,omitempty"`
	Patient                       Reference                    `json:"patient,omitempty"`
	ManagingOrganization          *Reference                   `json:"managingOrganization,omitempty"`
	Period                        *Period                      `json:"period,omitempty"`
	ReferralRequest               []Reference                  `json:"referralRequest,omitempty"`
	CareManager                   *Reference                   `json:"careManager,omitempty"`
	Team                          []Reference                  `json:"team,omitempty"`
	Account                       []Reference                  `json:"account,omitempty"`
}

func (r EpisodeOfCare) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EpisodeOfCare) marshalJSON() jsonEpisodeOfCare {
	m := jsonEpisodeOfCare{}
	m.ResourceType = "EpisodeOfCare"
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
	m.Type = r.Type
	m.Diagnosis = r.Diagnosis
	m.Patient = r.Patient
	m.ManagingOrganization = r.ManagingOrganization
	m.Period = r.Period
	m.ReferralRequest = r.ReferralRequest
	m.CareManager = r.CareManager
	m.Team = r.Team
	m.Account = r.Account
	return m
}
func (r *EpisodeOfCare) UnmarshalJSON(b []byte) error {
	var m jsonEpisodeOfCare
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EpisodeOfCare) unmarshalJSON(m jsonEpisodeOfCare) error {
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
	r.StatusHistory = m.StatusHistory
	r.Type = m.Type
	r.Diagnosis = m.Diagnosis
	r.Patient = m.Patient
	r.ManagingOrganization = m.ManagingOrganization
	r.Period = m.Period
	r.ReferralRequest = m.ReferralRequest
	r.CareManager = m.CareManager
	r.Team = m.Team
	r.Account = m.Account
	return nil
}
func (r EpisodeOfCare) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The history of statuses that the EpisodeOfCare has been through (without requiring processing the history of the resource).
type EpisodeOfCareStatusHistory struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// planned | waitlist | active | onhold | finished | cancelled.
	Status Code
	// The period during this EpisodeOfCare that the specific status applied.
	Period Period
}
type jsonEpisodeOfCareStatusHistory struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Status                 Code              `json:"status,omitempty"`
	StatusPrimitiveElement *primitiveElement `json:"_status,omitempty"`
	Period                 Period            `json:"period,omitempty"`
}

func (r EpisodeOfCareStatusHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EpisodeOfCareStatusHistory) marshalJSON() jsonEpisodeOfCareStatusHistory {
	m := jsonEpisodeOfCareStatusHistory{}
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
func (r *EpisodeOfCareStatusHistory) UnmarshalJSON(b []byte) error {
	var m jsonEpisodeOfCareStatusHistory
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EpisodeOfCareStatusHistory) unmarshalJSON(m jsonEpisodeOfCareStatusHistory) error {
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
func (r EpisodeOfCareStatusHistory) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The list of diagnosis relevant to this episode of care.
type EpisodeOfCareDiagnosis struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A list of conditions/problems/diagnoses that this episode of care is intended to be providing care for.
	Condition Reference
	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge …).
	Role *CodeableConcept
	// Ranking of the diagnosis (for each role type).
	Rank *PositiveInt
}
type jsonEpisodeOfCareDiagnosis struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Condition            Reference         `json:"condition,omitempty"`
	Role                 *CodeableConcept  `json:"role,omitempty"`
	Rank                 *PositiveInt      `json:"rank,omitempty"`
	RankPrimitiveElement *primitiveElement `json:"_rank,omitempty"`
}

func (r EpisodeOfCareDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EpisodeOfCareDiagnosis) marshalJSON() jsonEpisodeOfCareDiagnosis {
	m := jsonEpisodeOfCareDiagnosis{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Condition = r.Condition
	m.Role = r.Role
	m.Rank = r.Rank
	if r.Rank != nil && (r.Rank.Id != nil || r.Rank.Extension != nil) {
		m.RankPrimitiveElement = &primitiveElement{Id: r.Rank.Id, Extension: r.Rank.Extension}
	}
	return m
}
func (r *EpisodeOfCareDiagnosis) UnmarshalJSON(b []byte) error {
	var m jsonEpisodeOfCareDiagnosis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EpisodeOfCareDiagnosis) unmarshalJSON(m jsonEpisodeOfCareDiagnosis) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Condition = m.Condition
	r.Role = m.Role
	r.Rank = m.Rank
	if m.RankPrimitiveElement != nil {
		r.Rank.Id = m.RankPrimitiveElement.Id
		r.Rank.Extension = m.RankPrimitiveElement.Extension
	}
	return nil
}
func (r EpisodeOfCareDiagnosis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
