package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.  Includes companies, institutions, corporations, departments, community groups, healthcare practice groups, payer/insurer, etc.
type Organization struct {
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
	// Identifier for the organization that is used to identify the organization across multiple disparate systems.
	Identifier []Identifier
	// Whether the organization's record is still in active use.
	Active *Boolean
	// The kind(s) of organization that this is.
	Type []CodeableConcept
	// A name associated with the organization.
	Name *String
	// A list of alternate names that the organization is known as, or was known as in the past.
	Alias []String
	// A contact detail for the organization.
	Telecom []ContactPoint
	// An address for the organization.
	Address []Address
	// The organization of which this organization forms a part.
	PartOf *Reference
	// Contact for the organization for a certain purpose.
	Contact []OrganizationContact
	// Technical endpoints providing access to services operated for the organization.
	Endpoint []Reference
}
type jsonOrganization struct {
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
	Active                        *Boolean              `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement     `json:"_active,omitempty"`
	Type                          []CodeableConcept     `json:"type,omitempty"`
	Name                          *String               `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement     `json:"_name,omitempty"`
	Alias                         []String              `json:"alias,omitempty"`
	AliasPrimitiveElement         []*primitiveElement   `json:"_alias,omitempty"`
	Telecom                       []ContactPoint        `json:"telecom,omitempty"`
	Address                       []Address             `json:"address,omitempty"`
	PartOf                        *Reference            `json:"partOf,omitempty"`
	Contact                       []OrganizationContact `json:"contact,omitempty"`
	Endpoint                      []Reference           `json:"endpoint,omitempty"`
}

func (r Organization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Organization) marshalJSON() jsonOrganization {
	m := jsonOrganization{}
	m.ResourceType = "Organization"
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
	m.Type = r.Type
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Alias = r.Alias
	anyAliasIdOrExtension := false
	for _, e := range r.Alias {
		if e.Id != nil || e.Extension != nil {
			anyAliasIdOrExtension = true
			break
		}
	}
	if anyAliasIdOrExtension {
		m.AliasPrimitiveElement = make([]*primitiveElement, 0, len(r.Alias))
		for _, e := range r.Alias {
			if e.Id != nil || e.Extension != nil {
				m.AliasPrimitiveElement = append(m.AliasPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.AliasPrimitiveElement = append(m.AliasPrimitiveElement, nil)
			}
		}
	}
	m.Telecom = r.Telecom
	m.Address = r.Address
	m.PartOf = r.PartOf
	m.Contact = r.Contact
	m.Endpoint = r.Endpoint
	return m
}
func (r *Organization) UnmarshalJSON(b []byte) error {
	var m jsonOrganization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Organization) unmarshalJSON(m jsonOrganization) error {
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
	r.Type = m.Type
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Alias = m.Alias
	for i, e := range m.AliasPrimitiveElement {
		if len(r.Alias) > i {
			r.Alias[i].Id = e.Id
			r.Alias[i].Extension = e.Extension
		} else {
			r.Alias = append(r.Alias, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Telecom = m.Telecom
	r.Address = m.Address
	r.PartOf = m.PartOf
	r.Contact = m.Contact
	r.Endpoint = m.Endpoint
	return nil
}
func (r Organization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Contact for the organization for a certain purpose.
type OrganizationContact struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a purpose for which the contact can be reached.
	Purpose *CodeableConcept
	// A name associated with the contact.
	Name *HumanName
	// A contact detail (e.g. a telephone number or an email address) by which the party may be contacted.
	Telecom []ContactPoint
	// Visiting or postal addresses for the contact.
	Address *Address
}
type jsonOrganizationContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `json:"purpose,omitempty"`
	Name              *HumanName       `json:"name,omitempty"`
	Telecom           []ContactPoint   `json:"telecom,omitempty"`
	Address           *Address         `json:"address,omitempty"`
}

func (r OrganizationContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OrganizationContact) marshalJSON() jsonOrganizationContact {
	m := jsonOrganizationContact{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Purpose = r.Purpose
	m.Name = r.Name
	m.Telecom = r.Telecom
	m.Address = r.Address
	return m
}
func (r *OrganizationContact) UnmarshalJSON(b []byte) error {
	var m jsonOrganizationContact
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OrganizationContact) unmarshalJSON(m jsonOrganizationContact) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Purpose = m.Purpose
	r.Name = m.Name
	r.Telecom = m.Telecom
	r.Address = m.Address
	return nil
}
func (r OrganizationContact) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
