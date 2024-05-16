package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Demographics and administrative information about a person independent of a specific health-related context.
//
// Need to track persons potentially across multiple roles.
type Person struct {
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
	// Identifier for a person within a particular scope.
	Identifier []Identifier
	// A name associated with the person.
	Name []HumanName
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []ContactPoint
	// Administrative Gender.
	Gender *Code
	// The birth date for the person.
	BirthDate *Date
	// One or more addresses for the person.
	Address []Address
	// An image that can be displayed as a thumbnail of the person to enhance the identification of the individual.
	Photo *Attachment
	// The organization that is the custodian of the person record.
	ManagingOrganization *Reference
	// Whether this person's record is in active use.
	Active *Boolean
	// Link to a resource that concerns the same actual person.
	Link []PersonLink
}

func (r Person) ResourceType() string {
	return "Person"
}
func (r Person) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonPerson struct {
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
	Name                          []HumanName         `json:"name,omitempty"`
	Telecom                       []ContactPoint      `json:"telecom,omitempty"`
	Gender                        *Code               `json:"gender,omitempty"`
	GenderPrimitiveElement        *primitiveElement   `json:"_gender,omitempty"`
	BirthDate                     *Date               `json:"birthDate,omitempty"`
	BirthDatePrimitiveElement     *primitiveElement   `json:"_birthDate,omitempty"`
	Address                       []Address           `json:"address,omitempty"`
	Photo                         *Attachment         `json:"photo,omitempty"`
	ManagingOrganization          *Reference          `json:"managingOrganization,omitempty"`
	Active                        *Boolean            `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement   `json:"_active,omitempty"`
	Link                          []PersonLink        `json:"link,omitempty"`
}

func (r Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Person) marshalJSON() jsonPerson {
	m := jsonPerson{}
	m.ResourceType = "Person"
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
	m.Name = r.Name
	m.Telecom = r.Telecom
	m.Gender = r.Gender
	if r.Gender != nil && (r.Gender.Id != nil || r.Gender.Extension != nil) {
		m.GenderPrimitiveElement = &primitiveElement{Id: r.Gender.Id, Extension: r.Gender.Extension}
	}
	m.BirthDate = r.BirthDate
	if r.BirthDate != nil && (r.BirthDate.Id != nil || r.BirthDate.Extension != nil) {
		m.BirthDatePrimitiveElement = &primitiveElement{Id: r.BirthDate.Id, Extension: r.BirthDate.Extension}
	}
	m.Address = r.Address
	m.Photo = r.Photo
	m.ManagingOrganization = r.ManagingOrganization
	m.Active = r.Active
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Link = r.Link
	return m
}
func (r *Person) UnmarshalJSON(b []byte) error {
	var m jsonPerson
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Person) unmarshalJSON(m jsonPerson) error {
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
	r.Name = m.Name
	r.Telecom = m.Telecom
	r.Gender = m.Gender
	if m.GenderPrimitiveElement != nil {
		r.Gender.Id = m.GenderPrimitiveElement.Id
		r.Gender.Extension = m.GenderPrimitiveElement.Extension
	}
	r.BirthDate = m.BirthDate
	if m.BirthDatePrimitiveElement != nil {
		r.BirthDate.Id = m.BirthDatePrimitiveElement.Id
		r.BirthDate.Extension = m.BirthDatePrimitiveElement.Extension
	}
	r.Address = m.Address
	r.Photo = m.Photo
	r.ManagingOrganization = m.ManagingOrganization
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Link = m.Link
	return nil
}
func (r Person) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Link to a resource that concerns the same actual person.
type PersonLink struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The resource to which this actual person is associated.
	Target Reference
	// Level of assurance that this link is associated with the target resource.
	Assurance *Code
}
type jsonPersonLink struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Target                    Reference         `json:"target,omitempty"`
	Assurance                 *Code             `json:"assurance,omitempty"`
	AssurancePrimitiveElement *primitiveElement `json:"_assurance,omitempty"`
}

func (r PersonLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PersonLink) marshalJSON() jsonPersonLink {
	m := jsonPersonLink{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Target = r.Target
	m.Assurance = r.Assurance
	if r.Assurance != nil && (r.Assurance.Id != nil || r.Assurance.Extension != nil) {
		m.AssurancePrimitiveElement = &primitiveElement{Id: r.Assurance.Id, Extension: r.Assurance.Extension}
	}
	return m
}
func (r *PersonLink) UnmarshalJSON(b []byte) error {
	var m jsonPersonLink
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PersonLink) unmarshalJSON(m jsonPersonLink) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Target = m.Target
	r.Assurance = m.Assurance
	if m.AssurancePrimitiveElement != nil {
		r.Assurance.Id = m.AssurancePrimitiveElement.Id
		r.Assurance.Extension = m.AssurancePrimitiveElement.Extension
	}
	return nil
}
func (r PersonLink) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
