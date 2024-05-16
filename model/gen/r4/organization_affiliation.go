package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Defines an affiliation/assotiation/relationship between 2 distinct oganizations, that is not a part-of relationship/sub-division relationship.
//
// Need to define relationships between organizations that are not sub-divisions of the same organization (part-of relationships).
type OrganizationAffiliation struct {
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
	// Business identifiers that are specific to this role.
	Identifier []Identifier
	// Whether this organization affiliation record is in active use.
	Active *Boolean
	// The period during which the participatingOrganization is affiliated with the primary organization.
	Period *Period
	// Organization where the role is available (primary organization/has members).
	Organization *Reference
	// The Participating Organization provides/performs the role(s) defined by the code to the Primary Organization (e.g. providing services or is a member of).
	ParticipatingOrganization *Reference
	// Health insurance provider network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined).
	Network []Reference
	// Definition of the role the participatingOrganization plays in the association.
	Code []CodeableConcept
	// Specific specialty of the participatingOrganization in the context of the role.
	Specialty []CodeableConcept
	// The location(s) at which the role occurs.
	Location []Reference
	// Healthcare services provided through the role.
	HealthcareService []Reference
	// Contact details at the participatingOrganization relevant to this Affiliation.
	Telecom []ContactPoint
	// Technical endpoints providing access to services operated for this role.
	Endpoint []Reference
}

func (r OrganizationAffiliation) ResourceType() string {
	return "OrganizationAffiliation"
}
func (r OrganizationAffiliation) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonOrganizationAffiliation struct {
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
	Active                        *Boolean            `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement   `json:"_active,omitempty"`
	Period                        *Period             `json:"period,omitempty"`
	Organization                  *Reference          `json:"organization,omitempty"`
	ParticipatingOrganization     *Reference          `json:"participatingOrganization,omitempty"`
	Network                       []Reference         `json:"network,omitempty"`
	Code                          []CodeableConcept   `json:"code,omitempty"`
	Specialty                     []CodeableConcept   `json:"specialty,omitempty"`
	Location                      []Reference         `json:"location,omitempty"`
	HealthcareService             []Reference         `json:"healthcareService,omitempty"`
	Telecom                       []ContactPoint      `json:"telecom,omitempty"`
	Endpoint                      []Reference         `json:"endpoint,omitempty"`
}

func (r OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OrganizationAffiliation) marshalJSON() jsonOrganizationAffiliation {
	m := jsonOrganizationAffiliation{}
	m.ResourceType = "OrganizationAffiliation"
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
	m.Organization = r.Organization
	m.ParticipatingOrganization = r.ParticipatingOrganization
	m.Network = r.Network
	m.Code = r.Code
	m.Specialty = r.Specialty
	m.Location = r.Location
	m.HealthcareService = r.HealthcareService
	m.Telecom = r.Telecom
	m.Endpoint = r.Endpoint
	return m
}
func (r *OrganizationAffiliation) UnmarshalJSON(b []byte) error {
	var m jsonOrganizationAffiliation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OrganizationAffiliation) unmarshalJSON(m jsonOrganizationAffiliation) error {
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
	r.Organization = m.Organization
	r.ParticipatingOrganization = m.ParticipatingOrganization
	r.Network = m.Network
	r.Code = m.Code
	r.Specialty = m.Specialty
	r.Location = m.Location
	r.HealthcareService = m.HealthcareService
	r.Telecom = m.Telecom
	r.Endpoint = m.Endpoint
	return nil
}
func (r OrganizationAffiliation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
